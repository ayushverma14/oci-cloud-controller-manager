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
	"log"
	"time"

	"go.uber.org/zap"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"

	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	//"k8s.io/client-go/util/workqueue"
	//ctrl "sigs.k8s.io/controller-runtime"
	npnv1beta1 "github.com/oracle/oci-cloud-controller-manager/api/v1beta1"
	providercfg "github.com/oracle/oci-cloud-controller-manager/pkg/cloudprovider/providers/oci/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"k8s.io/client-go/tools/clientcmd"
	//"github.com/oracle/oci-cloud-controller-manager/pkg/oci/client"
	//	"sigs.k8s.io/controller-runtime/pkg/log"

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

func AddController(mgr manager.Manager) error {
	// Create a new Controller
	logger := zap.L()

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
		&handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &v1.Node{},
		})
	if err != nil {
		log.Println(err, "err")
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
		log.Println(err, "err")
		return err
	}

	return nil
}
func (r NativePodNetworkNONOKEReconciler) getNodeObjectInCluster(ctx context.Context, cr types.NamespacedName, nodeName string) (*v1.Node, error) {

	nodeObject := v1.Node{}
	login := zap.L()
	nodePresentInCluster := func() (bool, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Second*30)
		defer cancel()
		err := r.Client.Get(ctx, types.NamespacedName{
			Name: nodeName,
		}, &nodeObject)
		login.Info(cr.Name)
		if err != nil {
			if apierrors.IsNotFound(err) {
				login.Error("error", zap.Error(err))
				return false, nil
			}
			log.Println(err, "failed to get node object")
			return false, err
		}
		return true, nil
	}

	err := wait.PollImmediate(time.Second*5, GetNodeTimeout, func() (bool, error) {
		present, err := nodePresentInCluster()
		if err != nil {

			login.Error("error", zap.Error(err))
			return false, err
		}
		return present, nil
	})
	if err != nil {
		login.Error("error", zap.Error(err))
	}
	return &nodeObject, err
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

	login.Info(nodeName[0].Name)
	login.Info(nodeName[1].Name)
	login.Info(nodeName[2].Name)
	login.Sugar().Info(len(nodeName))
	login.Info("fetched info about node")
	////////////////////////////////////
	// Loop to check whether the node has a lablel npn attached if yes check if the CR exist else create it
	for j := range nodeName {
		target_node := nodeName[j]
		label := getLabel((&target_node))
		login.Info(target_node.Name)
		login.Sugar().Info(label)
		if label {

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
					// if err != nil {
					// 	login.Error("error", zap.Error(err))
					// 	return reconcile.Result{}, err
					// }
					login.Info("Reading config")
					configFilePath := "/etc/oci/config.yaml"

					configPath := configFilePath

					cfg1 := providercfg.GetConfig(login.Sugar(), configPath)

					ll := *cfg1
					cfg := ll.Specs
					login.Sugar().Info(cfg)

					CCEmails := []*string{}
					for i := range cfg.PodSubnetId {
						CCEmails = append(CCEmails, &cfg.PodSubnetId[i])
					}
					login.Info("creating subnets ids for cr on node")
					CCEmails1 := []*string{}
					for i := range cfg.NetworkSecurityGroupIds {
						CCEmails1 = append(CCEmails1, &cfg.NetworkSecurityGroupIds[i])
					}
					login.Info("creating nsfs for cr  on node")
					var npn1 = &npnv1beta1.NativePodNetwork{
						Spec: npnv1beta1.NativePodNetworkSpec{
							MaxPodCount:             &cfg.MaxPodsperNode,
							PodSubnetIds:            CCEmails,
							Id:                      &cfg.Id,
							NetworkSecurityGroupIds: CCEmails1,
						},
					}
					login.Info("creating npn1 for cr  on node ")
					npn1.Name = target_node.Name
					login.Debug("npn1", zap.Any("config", npn1))
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
func NewK8sClient(masterUrl, kubeconfigPath string) (kubernetes.Interface, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
	if err != nil {
		return nil, err
	}

	// create the clientset
	return kubernetes.NewForConfig(config)
}
func getLabel(node *v1.Node) bool {

	login := zap.L()
	if hh, ok := node.Labels["npn"]; ok {
		login.Info(hh)

		return true
	}

	return false
}
