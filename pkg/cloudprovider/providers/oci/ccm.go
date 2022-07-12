// Copyright 2017 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package oci implements an external Kubernetes cloud-provider for Oracle Cloud
// Infrastructure.
package oci

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	npnv1beta1 "github.com/oracle/oci-cloud-controller-manager/api/v1beta1"
	"github.com/oracle/oci-cloud-controller-manager/controllers"
	providercfg "github.com/oracle/oci-cloud-controller-manager/pkg/cloudprovider/providers/oci/config"
	"github.com/oracle/oci-cloud-controller-manager/pkg/metrics"
	"github.com/oracle/oci-cloud-controller-manager/pkg/oci/client"
	"github.com/oracle/oci-cloud-controller-manager/pkg/oci/instance/metadata"
	"github.com/oracle/oci-go-sdk/v50/common"
	"github.com/oracle/oci-go-sdk/v50/core"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/go-logr/zapr"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	listersv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	cloudprovider "k8s.io/cloud-provider"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
)

const (
	// providerName uniquely identifies the Oracle Cloud Infrastructure
	// (OCI) cloud-provider.
	providerName   = "oci"
	providerPrefix = providerName + "://"
)

// ProviderName uniquely identifies the Oracle Bare Metal Cloud Services (OCI)
// cloud-provider.
func ProviderName() string {
	return providerName
}

// CloudProvider is an implementation of the cloud-provider interface for OCI.
type CloudProvider struct {
	// NodeLister provides a cache to lookup nodes for deleting a load balancer.
	// Due to limitations in the OCI API around going from an IP to a subnet
	// we use the node lister to go from IP -> node / provider id -> ... -> subnet
	NodeLister listersv1.NodeLister

	client     client.Interface
	kubeclient clientset.Interface

	securityListManagerFactory securityListManagerFactory
	config                     *providercfg.Config

	logger        *zap.SugaredLogger
	instanceCache cache.Store
	metricPusher  *metrics.MetricPusher
}

var (
	schemes        = runtime.NewScheme()
	npnSetupLog    = ctrl.Log.WithName("npn-controller-setup")
	configFilePath = "/etc/oci/config.yaml"
)

func (cp *CloudProvider) InstancesV2() (cloudprovider.InstancesV2, bool) {
	cp.logger.Debug("Claiming to not support instancesV2")
	return nil, false
}

// Compile time check that CloudProvider implements the cloudprovider.Interface
// interface.
var _ cloudprovider.Interface = &CloudProvider{}

// NewCloudProvider creates a new oci.CloudProvider.
func NewCloudProvider(config *providercfg.Config) (cloudprovider.Interface, error) {
	// The global logger has been replaced with the logger we constructed in
	// cloud_provider_oci.go so capture it here and then pass it into all components.
	logger := zap.L()
	logger = logger.With(zap.String("component", "cloud-controller-manager"))

	cp, err := providercfg.NewConfigurationProvider(config)
	if err != nil {
		return nil, err
	}

	rateLimiter := client.NewRateLimiter(logger.Sugar(), config.RateLimiter)

	c, err := client.New(logger.Sugar(), cp, &rateLimiter)
	if err != nil {
		return nil, err
	}

	if config.CompartmentID == "" {
		logger.Info("Compartment not supplied in config: attempting to infer from instance metadata")
		metadata, err := metadata.New().Get()
		if err != nil {
			return nil, err
		}
		config.CompartmentID = metadata.CompartmentID
	}

	if !config.LoadBalancer.Disabled && config.VCNID == "" {
		logger.Info("No VCN provided in cloud provider config. Falling back to looking up VCN via LB subnet.")
		subnet, err := c.Networking().GetSubnet(context.Background(), config.LoadBalancer.Subnet1)
		if err != nil {
			return nil, errors.Wrap(err, "get subnet for loadBalancer.subnet1")
		}
		config.VCNID = *subnet.VcnId
	}

	//metricPusher, err := metrics.NewMetricPusher(logger.Sugar())
	// if err != nil {
	// 	//logger.Sugar().With("error", err).Error("Metrics collection could not be enabled")
	// 	// disable metrics
	// 	metricPusher = nil
	// }

	// if metricPusher != nil {
	// 	logger.Info("Metrics collection has been enabled")
	// } else {
	// 	logger.Info("Metrics collection has not been enabled")
	// }

	return &CloudProvider{
		client:        c,
		config:        config,
		logger:        logger.Sugar(),
		instanceCache: cache.NewTTLStore(instanceCacheKeyFn, time.Duration(24)*time.Hour),
		//metricPusher:  metricPusher,
	}, nil
}

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName(), func(config io.Reader) (cloudprovider.Interface, error) {
		cfg, err := providercfg.ReadConfig(config)
		if err != nil {
			return nil, err
		}

		if err = cfg.Validate(); err != nil {
			return nil, err
		}

		return NewCloudProvider(cfg)
	})
	common.EnableInstanceMetadataServiceLookup()
}

// Initialize passes a Kubernetes clientBuilder interface to the cloud provider.
func (cp *CloudProvider) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	var err error
	cp.kubeclient, err = clientBuilder.Client("cloud-controller-manager")
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to create kubeclient: %v", err))
	}

	factory := informers.NewSharedInformerFactory(cp.kubeclient, 5*time.Minute)

	nodeInfoController := NewNodeInfoController(
		factory.Core().V1().Nodes(),
		cp.kubeclient,
		cp,
		cp.logger,
		cp.instanceCache,
		cp.client)

	nodeInformer := factory.Core().V1().Nodes()
	go nodeInformer.Informer().Run(wait.NeverStop)
	serviceInformer := factory.Core().V1().Services()
	go serviceInformer.Informer().Run(wait.NeverStop)
	go nodeInfoController.Run(wait.NeverStop)

	cp.logger.Info("Waiting for node informer cache to sync")
	if !cache.WaitForCacheSync(wait.NeverStop, nodeInformer.Informer().HasSynced, serviceInformer.Informer().HasSynced) {
		utilruntime.HandleError(fmt.Errorf("Timed out waiting for informers to sync"))
	}
	cp.NodeLister = nodeInformer.Lister()

	cp.securityListManagerFactory = func(mode string) securityListManager {
		if cp.config.LoadBalancer.Disabled {
			return newSecurityListManagerNOOP()
		}
		if len(mode) == 0 {
			mode = cp.config.LoadBalancer.SecurityListManagementMode
		}
		return newSecurityListManager(cp.logger, cp.client, serviceInformer, cp.config.LoadBalancer.SecurityLists, mode)
	}
	var logger *zap.SugaredLogger
	var wg sync.WaitGroup
	enableNIC := true

	if providercfg.EnableNICController || enableNIC {
		wg.Add(1)
		logger = logger.With(zap.String("component", "npncr-controller"))
		ctrl.SetLogger(zapr.NewLogger(logger.Desugar()))
		logger.Info("NPN_CR controller is enabled.")
		defer wg.Done()
		mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
			Scheme:                  schemes,
			MetricsBindAddress:      ":8080",
			Port:                    9443,
			HealthProbeBindAddress:  ":8081",
			LeaderElection:          true,
			LeaderElectionID:        "npn.oci.oraclecloud.com",
			LeaderElectionNamespace: "kube-system",
		})
		if err != nil {
			npnSetupLog.Error(err, "unable to start manager")
			os.Exit(1)
		}
		err = controllers.Add(mgr)
		if err != nil {
			logger.Info(err)
		} else {
			logger.Info("controller is setup properly")
		}
	}

	if enableNIC || providercfg.EnableNICController {
		wg.Add(1)
		logger = logger.With(zap.String("component", "npn-controller"))
		ctrl.SetLogger(zapr.NewLogger(logger.Desugar()))
		logger.Info("NPN controller is enabled.")
		go func() {
			defer wg.Done()
			utilruntime.Must(clientgoscheme.AddToScheme(schemes))
			utilruntime.Must(npnv1beta1.AddToScheme(schemes))

			configPath, ok := os.LookupEnv("CONFIG_YAML_FILENAME")
			if !ok {
				configPath = configFilePath
			}
			cfg := providercfg.GetConfig(logger, configPath)
			ociClient := getOCIClient(logger, cfg)

			mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
				Scheme:                  schemes,
				MetricsBindAddress:      ":8080",
				Port:                    9443,
				HealthProbeBindAddress:  ":8081",
				LeaderElection:          true,
				LeaderElectionID:        "npn.oci.oraclecloud.com",
				LeaderElectionNamespace: "kube-system",
			})
			if err != nil {
				npnSetupLog.Error(err, "unable to start manager")
				os.Exit(1)
			}

			metricPusher, err := metrics.NewMetricPusher(logger)
			if err != nil {
				logger.With("error", err).Error("metrics collection could not be enabled")
				// disable metrics
				metricPusher = nil
			}

			if err = (&controllers.NativePodNetworkReconciler{
				Client:           mgr.GetClient(),
				Scheme:           mgr.GetScheme(),
				MetricPusher:     metricPusher,
				OCIClient:        ociClient,
				TimeTakenTracker: make(map[string]time.Time),
			}).SetupWithManager(mgr); err != nil {
				npnSetupLog.Error(err, "unable to create controller", "controller", "NativePodNetwork")
				os.Exit(1)
			}

			if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
				npnSetupLog.Error(err, "unable to set up health check")
				os.Exit(1)
			}
			if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
				npnSetupLog.Error(err, "unable to set up ready check")
				os.Exit(1)
			}

			npnSetupLog.Info("starting manager")
			if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
				npnSetupLog.Error(err, "problem running manager")
				// TODO: Handle the case of NPN controller not running more gracefully
				os.Exit(1)
			}
		}()
	}
	wg.Wait()
}

// ProviderName returns the cloud-provider ID.
func (cp *CloudProvider) ProviderName() string {
	return ProviderName()
}

// LoadBalancer returns a balancer interface. Also returns true if the interface
// is supported, false otherwise.
func (cp *CloudProvider) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	cp.logger.Debug("Claiming to support load balancers")
	return cp, !cp.config.LoadBalancer.Disabled
}

// Instances returns an instances interface. Also returns true if the interface
// is supported, false otherwise.
func (cp *CloudProvider) Instances() (cloudprovider.Instances, bool) {
	cp.logger.Debug("Claiming to support instances")
	return cp, true
}

// Zones returns a zones interface. Also returns true if the interface is
// supported, false otherwise.
func (cp *CloudProvider) Zones() (cloudprovider.Zones, bool) {
	cp.logger.Debug("Claiming to support zones")
	return cp, true
}

// Clusters returns a clusters interface.  Also returns true if the interface is
// supported, false otherwise.
func (cp *CloudProvider) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is
// supported.
func (cp *CloudProvider) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ScrubDNS provides an opportunity for cloud-provider-specific code to process
// DNS settings for pods.
func (cp *CloudProvider) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return nameservers, searches
}

// HasClusterID returns true if the cluster has a clusterID.
func (cp *CloudProvider) HasClusterID() bool {
	return true
}

func instanceCacheKeyFn(obj interface{}) (string, error) {
	return *obj.(*core.Instance).Id, nil
}
func getOCIClient(logger *zap.SugaredLogger, config *providercfg.Config) client.Interface {
	c, err := client.GetClient(logger, config)

	if err != nil {
		logger.With(zap.Error(err)).Fatal("client can not be generated.")
	}
	return c
}
