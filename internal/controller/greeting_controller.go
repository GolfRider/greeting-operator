/*
Copyright 2026.

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
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "greeting-operator/api/v1"
)

// GreetingReconciler reconciles a Greeting object
type GreetingReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// RBAC markers - generate permissions for the controller
// +kubebuilder:rbac:groups=apps.example.com,resources=greetings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.example.com,resources=greetings/status,verbs=get;update;patch

// Reconcile - called whenever a Greeting changes (or periodically)
//
// The goal: make reality match the desired state
// This function should be IDEMPOTENT - running it twice = same result
func (r *GreetingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := logf.FromContext(ctx)

	// -------------------------------------------------------------------------
	// STEP 1: Fetch the Greeting resource
	// -------------------------------------------------------------------------
	var greeting appsv1.Greeting
	if err := r.Get(ctx, req.NamespacedName, &greeting); err != nil {
		// Not found = deleted, ignore
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Reconciling", "name", greeting.Spec.Name)

	// -------------------------------------------------------------------------
	// STEP 2: Do our "business logic" (just build a message)
	// -------------------------------------------------------------------------
	message := fmt.Sprintf("%s, %s!", greeting.Spec.Greeting, greeting.Spec.Name)

	// -------------------------------------------------------------------------
	// STEP 3: Update the status
	// -----------------------------------------------------------------

	greeting.Status.Message = message
	greeting.Status.Ready = true
	// Status().Update() only updates the status subresource
	if err := r.Status().Update(ctx, &greeting); err != nil {
		logger.Error(err, "Failed to update status")
		return ctrl.Result{}, err
	}

	logger.Info("Reconciled successfully", "message", message)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GreetingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Greeting{}). // Watch Greeting resources
		Named("greeting").
		Complete(r)
}
