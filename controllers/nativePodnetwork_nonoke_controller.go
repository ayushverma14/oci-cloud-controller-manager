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
	"errors"
	"math"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"github.com/oracle/oci-cloud-controller-manger/pkg/cloudprovider/providers/oci/config"
	npnv1beta1 "github.com/oracle/oci-cloud-controller-manager/api/v1beta1"
	"github.com/oracle/oci-cloud-controller-manager/pkg/metrics"
	ociclient "github.com/oracle/oci-cloud-controller-manager/pkg/oci/client"
	"github.com/oracle/oci-cloud-controller-manager/pkg/util"
	"github.com/oracle/oci-go-sdk/v49/core"
)

const (
	CREATE_PRIVATE_IP             = "CREATE_PRIVATE_IP"
	ATTACH_VNIC                   = "ATTACH_VNIC"
	INITIALIZE_NPN_NODE           = "INITIALIZE_NPN_NODE"
	maxSecondaryPrivateIPsPerVNIC = 31
	// GetNodeTimeout is the timeout for the node object to be created in Kubernetes
	GetNodeTimeout = 20 * time.Minute
	// RunningInstanceTimeout is the timeout for the instance to reach running state
	// before we try to attach VNIC(s) to them
	RunningInstanceTimeout                   = 5 * time.Minute
	FetchedExistingSecondaryVNICsForInstance = "Fetched existingSecondaryVNICs for instance"
)

var (
	STATE_SUCCESS     = "SUCCESS"
	STATE_IN_PROGRESS = "IN_PROGRESS"
	STATE_BACKOFF     = "BACKOFF"
	COMPLETED         = "COMPLETED"

	SKIP_SOURCE_DEST_CHECK = true
	errPrimaryVnicNotFound = errors.New("failed to get primary vnic for instance")
	errInstanceNotRunning  = errors.New("instance is not in running state")
)
type NativePodNetworkNONOKEReconciler struct {
	client.Client
	Scheme           *runtime.Scheme
	// MetricPusher     *metrics.MetricPusher
	// OCIClient        ociclient.Interface
	// TimeTakenTracker map[string]time.Time
    // Config           *config.NativepodNetwork
}
func Add(mgr manager.Manager) error () {
	// Create a new Controller
	c, err := controller.New("NativePodNewtorkNONOKEReconciler-controller", mgr,
	  controller.Options{Reconciler: &NativePodNetworkNONOKEController{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
		
	}})
	if err != nil {
	  return err
	}
  
	// Watch for changes to npn types
	err = c.Watch(
	  &source.Kind{Type:&npnv1beta1.NativePodNetwork{}},
		&handler.EnqueueRequestForObject{})
	if err != nil {
	  return err
	}
  
	  // Watch for changes to nodes created by a ContainerSet and trigger a Reconcile for the owner
	err = c.Watch(
	  &source.Kind{Type: &v1.Node{}},
		&handler.EnqueueRequestForOwner{
		  IsController: true,
		  OwnerType:    &npnv1beta1.NativePodNetwork{},
		})
	if err != nil {
	  return err
	}
  
	return nil
  }
  func (r *NativePodNetworkNONOKEReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	npn:= &v1beat1.NativePodNetwork{}

	nodeObject,err := getNodeObjectInCluster(context.TODO(),req.NamespacedName)
	if err != nil {
		return reconcile.Result{},err
	}


    err := r.Get(context.TODO(), request.NamespacedName, npn)
  if err != nil {
    if errors.IsNotFound(err) {
      // Object not found, return.  Created objects are automatically garbage collected.
      // For additional cleanup logic use finalizers.
	  s,err := Readspec(strings.NewReader(spec1))
 var npn1=&v1beat1.NativePodNetwork{
	MaxPodCount: s.Specs.maxPodsperNode,
	PodSubnetId: s.Specs.PodSubnetId,
	Id: s.Specs.id,
	NetworkSecurityGroupIds: s.Specs.NetworkSecurityGroupIds,

 }

      return reconcile.Result{}, nil
    }
    // Error reading the object - requeue the request.
    return reconcile.Result{}, err
  }
	
}
func (r NativePodNetworkNONOKEReconciler) getNodeObjectInCluster(ctx context.Context, cr types.NamespacedName) (*v1.Node, error) {
	
	nodeObject := v1.Node{}
	nodePresentInCluster := func() (bool, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Second*30)
		defer cancel()
		err := r.Client.Get(ctx, types.NamespacedName, &nodeObject)
		if err != nil {
			if apierrors.IsNotFound(err) {
				log.Error(err, "node object does not exist in cluster")
				return false, nil
			}
			log.Error(err, "failed to get node object")
			return false, err
		}
		return true, nil
	}

	err := wait.PollImmediate(time.Second*5, GetNodeTimeout, func() (bool, error) {
		present, err := nodePresentInCluster()
		if err != nil {
			log.Error(err, "failed to get node from cluster")
			return false, err
		}
		return present, nil
	})
	if err != nil {
		log.Error(err, "timed out waiting for node object to be present in the cluster")
	}
	return &nodeObject, err
}
