/*
Copyright 2019 The Kubernetes Authors.

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
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "kubevirt.io/client-go/api/v1"
	kubevirtv1alpha1 "sigs.k8s.io/cluster-api-provider-kubevirt/api/v1alpha1"
)

// MachineReconciler reconciles a Machine object
type MachineReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kubevirt.cluster.x-k8s.io,resources=machines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubevirt.cluster.x-k8s.io,resources=machines/status,verbs=get;update;patch

func (r *MachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("machine", req.NamespacedName)

	// your logic here

	bar := v1.VirtualMachineInstance{}
	fmt.Printf("%+v\n", bar)
	return ctrl.Result{}, nil
}

func (r *MachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubevirtv1alpha1.Machine{}).
		Complete(r)
}
