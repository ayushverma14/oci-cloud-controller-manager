/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"log"
	"time"

	"go.uber.org/zap"

	"k8s.io/apimachinery/pkg/types"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	npnv1beta1 "github.com/oracle/oci-cloud-controller-manager/api/v1beta1"
	providercfg "github.com/oracle/oci-cloud-controller-manager/pkg/cloudprovider/providers/oci/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"k8s.io/client-go/tools/clientcmd"

	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	scheme         *runtime.Scheme = runtime.NewScheme()
	configFilePath                 = "/etc/oci/config.yaml"
)

func init() {
	utilruntime.Must(npnv1beta1.AddToScheme(scheme))
}

type NativePodNetworkNONOKEReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func AddController(mgr manager.Manager) error {

	logger := zap.L()
	// Create a new controller and set its parameters
	c, err := controller.New("NativePodNewtorkNONOKEReconciler-controller", mgr,
		controller.Options{Reconciler: &NativePodNetworkNONOKEReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		},
			MaxConcurrentReconciles: 20,
			CacheSyncTimeout:        time.Hour,
		})

	logger.Info("watching npn")
	if err != nil {
		log.Println(err, "err")
		return err
	}

	// Watch for changes to npn types
	err = c.Watch(
		&source.Kind{Type: &npnv1beta1.NativePodNetwork{}},
		&handler.EnqueueRequestForObject{})

	if err != nil {
		logger.Sugar().Error(err)
		return err
	}

	logger.Info("watching nodes")
	// Watch for changes to nodes  and trigger a Reconcile for the owner
	err = c.Watch(
		&source.Kind{Type: &v1.Node{}},
		&handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &npnv1beta1.NativePodNetwork{},
		})

	if err != nil {
		logger.Sugar().Error(err)
		return err
	}

	return nil
}

//+kubebuilder:rbac:groups=oci.oraclecloud.com,resources=nativepodnetworkings,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=oci.oraclecloud.com,resources=nativepodnetworkings/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=oci.oraclecloud.com,resources=nativepodnetworkings/finalizers,verbs=update

func (r *NativePodNetworkNONOKEReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {

	var npn = &npnv1beta1.NativePodNetwork{}

	login := zap.L()

	login.Info("Reconciling--------------------")
	login.Info("Generating Kubernetes clientset")
	// created a k8 clientset to fetch info about cluster
	kcs, err := NewK8sClient("https://129.146.114.63:6443", "/etc/kubernetes/admin.conf")

	login.Error("error", zap.Error(err))

	login.Info("Listing nodes")
	// List all nodes available on the cluster
	nodeName, err := ListNodes(kcs)

	login.Error("error", zap.Error(err))
	if err != nil {
		login.Error("error", zap.Error(err))
		return reconcile.Result{}, err
	}

	login.Info("fetched info about node")
	////////////////////////////////////
	// Loop to check whether the node has a lablel npn attached if yes check if the CR exist else create it
	for j := range nodeName {
		target_node := nodeName[j]
		label := getLabel((&target_node))
		login.Info(target_node.Name)
		login.Sugar().Info(label)

		if label {
			// checking if npn cr exists or not
			err = r.Get(ctx, types.NamespacedName{
				Name: nodeName[j].Name,
			}, npn)
			if err != nil {
				login.Info("npn not present on node ")
				login.Error("error", zap.Error(err))
				log.Println(err)

				if apierrors.IsNotFound(err) {
					// Object not found, return.  Created objects are automatically garbage collected.
					// For additional cleanup logic use finalizers.
					login.Info("creating npn cr on node")

					login.Info("Reading config")

					configPath := configFilePath

					cfg1 := providercfg.GetConfig(login.Sugar(), configPath)

					cgf_temp := *cfg1
					cfg := cgf_temp.Specs
					login.Sugar().Info(cfg)
					login.Info("creating subnets ids for cr on node")
					subnetIds := []*string{}
					for i := range cfg.PodSubnetId {
						subnetIds = append(subnetIds, &cfg.PodSubnetId[i])
					}
					login.Info("creating NSGids for cr  on node")
					nsgIds := []*string{}
					for i := range cfg.NetworkSecurityGroupIds {
						nsgIds = append(nsgIds, &cfg.NetworkSecurityGroupIds[i])
					}
					// initialising  NPN CR object
					var npn1 = &npnv1beta1.NativePodNetwork{
						Spec: npnv1beta1.NativePodNetworkSpec{
							MaxPodCount:             &cfg.MaxPodsperNode,
							PodSubnetIds:            subnetIds,
							Id:                      &cfg.Id,
							NetworkSecurityGroupIds: nsgIds,
						},
					}

					npn1.Name = target_node.Name

					login.Info("Creating the NPN CR ")
					err := r.Create(ctx, npn1)

					login.Error("error", zap.Error(err))
					login.Info("created the CR successfully")
					return reconcile.Result{}, nil
				}
				// Error reading the object - requeue the request.
				return reconcile.Result{}, err
			} else {
				login.Info("npn  already present on node")
			}
			login.Info("npn present on node")
		}
	}
	return reconcile.Result{}, err
}

// function to list all nodes availabe on cluster
func ListNodes(clientset kubernetes.Interface) ([]v1.Node, error) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	login := zap.L()

	if err != nil {
		login.Error("error", zap.Error(err))
		return nil, err

	}
	login.Debug("nodes", zap.Any("list", nodes.Items))

	return nodes.Items, nil
}

// function to generate a K8 client for accessing info og the cluster
func NewK8sClient(masterUrl, kubeconfigPath string) (kubernetes.Interface, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
	if err != nil {
		return nil, err
	}

	// create the clientset
	return kubernetes.NewForConfig(config)
}

// function to fetch npn label
func getLabel(node *v1.Node) bool {

	login := zap.L()
	if hh, ok := node.Labels["npn"]; ok {
		login.Info(hh)

		return true
	}

	return false
}
