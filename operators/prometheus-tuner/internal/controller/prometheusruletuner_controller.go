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
	"time"

	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	srev1alpha1 "github.com/taikun-cloud/poseidon/operators/prometheus-tuner/api/v1alpha1"
)

// PrometheusRuleTunerReconciler reconciles a PrometheusRuleTuner object
type PrometheusRuleTunerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=sre.taikun.cloud,resources=prometheusruletuners,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch
// +kubebuilder:rbac:groups=sre.taikun.cloud,resources=prometheusruletuners/finalizers,verbs=update
func (r *PrometheusRuleTunerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the PrometheusRuleTuner instance
	var tuner srev1alpha1.PrometheusRuleTuner
	if err := r.Get(ctx, req.NamespacedName, &tuner); err != nil {
		log.Error(err, "unable to fetch PrometheusRuleTuner")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Setup metrics client
	cfg, err := ctrl.GetConfig()
	if err != nil {
		log.Error(err, "unable to get Kubernetes config")
		return ctrl.Result{}, err
	}

	metricsClient, err := metrics.NewForConfig(cfg)
	if err != nil {
		log.Error(err, "unable to create metrics client")
		return ctrl.Result{}, err
	}

	// Get all pod metrics in the namespace
	podMetricsList, err := metricsClient.MetricsV1beta1().PodMetricses(tuner.Spec.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Error(err, "unable to list pod metrics")
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	var totalCPU, totalMem int64
	var matchedPods int

	for _, podMetric := range podMetricsList.Items {
		// Match pods by owner deployment name in labels
		if podMetric.Labels["app"] == tuner.Spec.DeploymentName {
			for _, container := range podMetric.Containers {
				totalCPU += container.Usage.Cpu().MilliValue()
				totalMem += container.Usage.Memory().Value()
			}
			matchedPods++
		}
	}

	log.Info("Resource usage collected",
		"deployment", tuner.Spec.DeploymentName,
		"cpu(m)", totalCPU,
		"memory(bytes)", totalMem)

	// Update status
	tuner.Status.CPUUsage = fmt.Sprintf("%dm", totalCPU)
	tuner.Status.MemoryUsage = fmt.Sprintf("%dMi", totalMem/1024/1024)

	if err := r.Status().Update(ctx, &tuner); err != nil {
		log.Error(err, "unable to update PrometheusRuleTuner status")
		return ctrl.Result{}, err
	}

	// Requeue after 60 seconds to refresh metrics
	return ctrl.Result{RequeueAfter: 60 * time.Second}, nil
}


// SetupWithManager sets up the controller with the Manager.
func (r *PrometheusRuleTunerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&srev1alpha1.PrometheusRuleTuner{}).
		Complete(r)
}
