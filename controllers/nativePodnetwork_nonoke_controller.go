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
	//"sync"
    
	"time"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"

	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	//"k8s.io/client-go/util/workqueue"
	//ctrl "sigs.k8s.io/controller-runtime"
	npnv1beta1 "github.com/oracle/oci-cloud-controller-manager/api/v1beta1"
	providercfg "github.com/oracle/oci-cloud-controller-manager/pkg/cloudprovider/providers/oci/config"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/log"

	//"github.com/oracle/oci-cloud-controller-manager/pkg/metrics"
	//	ociclient "github.com/oracle/oci-cloud-controller-manager/pkg/oci/client"
	//	"github.com/oracle/oci-cloud-controller-manager/pkg/util"

	//	"github.com/oracle/oci-go-sdk/v49/core"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const spec1 = `
apiVersion: ocicloud
kind: npn
metadata:
 name: aman
spec:
  maxPodCount: 3
  id: ocid1.instance.oc1.iad.anuwcljs2ahbgkyc7anhsitk7veikgxldc6ex7rsqo5hhjvslqxaa5nf62pa
  podSubnetIds: ["ocid1.subnet.oc1.iad.aaaaaaaame2kmeb2x3443s3kdcq27he2akbj67eijc3iar4l2atwnp5ttslq"]
  networkSecurityGroupIds: [ocid1.networksecuritygroup.oc1.iad.aaaaaaaa2ezkh44ul7yy2jioznmiydya4vcoedqxhpvhjjkl7a6rx2m267gq]
  `

var (
	scheme *runtime.Scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(npnv1beta1.AddToScheme(scheme))
}

type NativePodNetworkNONOKEReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	// MetricPusher     *metrics.MetricPusher
	// OCIClient        ociclient.Interface
	// TimeTakenTracker map[string]time.Time
	// Config           *config.NativepodNetwork
}

func Add(mgr manager.Manager) error {
	// Create a new Controller
	logger := zap.L()

	c, err := controller.New("NativePodNewtorkNONOKEReconciler-controller", mgr,
		controller.Options{Reconciler: &NativePodNetworkNONOKEReconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		}})
	logger.Info("watching npn")
	if err != nil {
		log.Log.Error(err,"err")
		return err
	}

	// Watch for changes to npn types
	err = c.Watch(
		&source.Kind{Type: &npnv1beta1.NativePodNetwork{}},
		&handler.EnqueueRequestForObject{})
	if err != nil {
		log.Log.Error(err,"err")
		return err
	}
	logger.Info("watching nodes")
	// Watch for changes to nodes  and trigger a Reconcile for the owner
	err = c.Watch(
		&source.Kind{Type: &v1.Node{}},
		&handler.EnqueueRequestForObject{})

	if err != nil {
		log.Log.Error(err,"err")
		return err
	}

	return nil
}
func (r NativePodNetworkNONOKEReconciler) getNodeObjectInCluster(ctx context.Context, cr types.NamespacedName) (*v1.Node, error) {

	nodeObject := v1.Node{}
	nodePresentInCluster := func() (bool, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Second*30)
		defer cancel()
		err := r.Client.Get(ctx, types.NamespacedName{
			Name: cr.Name,
		}, &nodeObject)

		if err != nil {
			if apierrors.IsNotFound(err) {
				log.Log.Error(err, "node object does not exist in cluster")
				return false, nil
			}
			log.Log.Error(err, "failed to get node object")
			return false, err
		}
		return true, nil
	}

	err := wait.PollImmediate(time.Second*5, GetNodeTimeout, func() (bool, error) {
		present, err := nodePresentInCluster()
		if err != nil {
			log.Log.Error(err, "failed to get node from cluster")
			return false, err
		}
		return present, nil
	})
	if err != nil {
		log.Log.Error(err, "timed out waiting for node object to be present in the cluster")
	}
	return &nodeObject, err
}

func (r *NativePodNetworkNONOKEReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := log.FromContext(ctx)
	var npn = &npnv1beta1.NativePodNetwork{}

	_, err := r.getNodeObjectInCluster(context.TODO(), request.NamespacedName)
	if err != nil {
		log.Error(err,"not able to fetch object")
		return reconcile.Result{}, err
	}

	err = r.Get(context.TODO(), request.NamespacedName, npn)
	if err != nil {
		log.Info("npn not present on node")
		if apierrors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.

			if err != nil {
				return reconcile.Result{}, err
			}
			var cfg providercfg.Spec
			CCEmails := []*string{}
			for i := range cfg.PodSubnetId {
				CCEmails = append(CCEmails, &cfg.PodSubnetId[i])
			}
			CCEmails1 := []*string{}
			for i := range cfg.NetworkSecurityGroupIds {
				CCEmails1 = append(CCEmails1, &cfg.NetworkSecurityGroupIds[i])
			}

			var npn1 = &npnv1beta1.NativePodNetwork{
				Spec: npnv1beta1.NativePodNetworkSpec{
					MaxPodCount:             &cfg.MaxPodsperNode,
					PodSubnetIds:            CCEmails,
					Id:                      &cfg.Id,
					NetworkSecurityGroupIds: CCEmails1,
				},
			}
			r.Create(context.TODO(), npn1)

			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}	else {
		log.Info("npn  already present on node")
	}
	log.Info("npn present on node")
	
	return reconcile.Result{}, err
}
