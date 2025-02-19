/*
Copyright 2025.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	foodv1alpha1 "github.com/joshbrgs/taco-operator/api/v1alpha1"
	"github.com/prometheus/client_golang/prometheus"
)

// TacoOrderReconciler reconciles a TacoOrder object
type TacoOrderReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Define a Prometheus counter to track tacos served
var tacosServedCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "tacos_served_total",
		Help: "Total number of tacos served",
	},
	[]string{"type"}, // Label by taco type
)

func init() {
	metrics.Registry.MustRegister(tacosServedCounter)
}

// +kubebuilder:rbac:groups=food.tacos.io,resources=tacoorders,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=food.tacos.io,resources=tacoorders/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=food.tacos.io,resources=tacoorders/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TacoOrder object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *TacoOrderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	var order foodv1alpha1.TacoOrder
	if err := r.Get(ctx, req.NamespacedName, &order); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// "Cooking" the tacos by updating the status
	order.Status.Served = order.Spec.Quantity
	fmt.Printf("🌮 Serving %d %s tacos with extras: %v\n", order.Spec.Quantity, order.Spec.Type, order.Spec.Extras)

	// Increment the Prometheus counter
	tacosServedCounter.WithLabelValues(order.Spec.Type).Add(float64(order.Spec.Quantity))

	if err := r.Status().Update(ctx, &order); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TacoOrderReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&foodv1alpha1.TacoOrder{}).
		Complete(r)
}
