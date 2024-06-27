package controllers

import (
    "context"

    "github.com/go-logr/logr"
    "k8s.io/apimachinery/pkg/api/errors"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/types"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/controller"
    "sigs.k8s.io/controller-runtime/pkg/handler"
    "sigs.k8s.io/controller-runtime/pkg/reconcile"
    "sigs.k8s.io/controller-runtime/pkg/source"
    "github.com/ihsanpr88/g42-storage-operator/api/v1alpha1"
    "github.com/ihsanpr88/g42-storage-operator/controllers"

)

// G42StorageReconciler reconciles a G42Storage object
type G42StorageReconciler struct {
    client.Client
    Log    logr.Logger
    Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=storage.example.com,resources=g42storages,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.example.com,resources=g42storages/status,verbs=get;update;patch

func (r *G42StorageReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    _ = r.Log.WithValues("g42storage", req.NamespacedName)

    // your logic here

    return ctrl.Result{}, nil
}

func (r *G42StorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&storagev1alpha1.G42Storage{}).
        Complete(r)
}
