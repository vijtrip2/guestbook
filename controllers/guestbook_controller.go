/*
Copyright 2021.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	webappv1 "vijtrip2/guestbook/api/v1"
)

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webapp.vijtrip2,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.vijtrip2,resources=guestbooks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.vijtrip2,resources=guestbooks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Guestbook object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *GuestbookReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Hello from reconciler")
	var guestbook webappv1.Guestbook
	if err := r.Get(ctx, req.NamespacedName, &guestbook); err != nil {
		logger.Error(err, "Guestbook not found")
	}
	logger.Info("Guestbook:", "namespace:", guestbook.Namespace, "name:", guestbook.Name)
	logger.Info("Guestbook spec before the update is: ", "guestbook-spec", guestbook.Spec)
	logger.Info("Guestbook status before the update is: ", "guestbook-status", guestbook.Status)
	logger.Info("Guestbook annotations before the update is: ", "guestbook-annotations", guestbook.Annotations)
	guestbook.Spec.DefaultString = "defaultValue"
	guestbook.Spec.UpdatedString = "updatedValue"
	guestbook.Status.Res = "res"
	guestbook.Annotations["updatedAnnKey"] = "updatedAnnValue"
	logger.Info("Updated guestbook spec is:", "guestbook", guestbook.Spec)
	logger.Info("Updated guestbook status is:", "guestbook", guestbook.Status)
	r.Update(ctx, &guestbook)
	logger.Info("Updated the spec")
	r.Status().Update(ctx, &guestbook)
	logger.Info("Updated the status")

	var updatedGuestbook webappv1.Guestbook
	if err := r.Get(ctx, req.NamespacedName, &updatedGuestbook); err != nil {
		logger.Error(err, "Guestbook not found after update")
	}
	logger.Info("Guestbook spec after the update is: ", "guestbook-spec", updatedGuestbook.Spec)
	logger.Info("Guestbook status after the update is: ", "guestbook-status", updatedGuestbook.Status)
	logger.Info("Guestbook annotations after the update is: ", "guestbook-annotations", updatedGuestbook.Annotations)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Guestbook{}).
		Complete(r)
}
