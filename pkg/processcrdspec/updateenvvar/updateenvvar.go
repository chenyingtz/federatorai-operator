package updateenvvar

import (
	"strings"

	"github.com/containers-ai/federatorai-operator/pkg/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func AssignServiceToDeployment(dep *appsv1.Deployment, ns string) {
	if len(dep.Spec.Template.Spec.Containers[0].Env) > 0 {
		for index, value := range dep.Spec.Template.Spec.Containers[0].Env {
			if strings.Contains(value.String(), util.NamespaceService) {
				dep.Spec.Template.Spec.Containers[0].Env[index].Value = strings.Replace(dep.Spec.Template.Spec.Containers[0].Env[index].Value, util.NamespaceService, ns+".svc", -1)
			}
		}
	}

	for containerIdx, _ := range dep.Spec.Template.Spec.Containers {
		for index, value := range dep.Spec.Template.Spec.Containers[containerIdx].Args {
			if strings.Contains(value, util.NamespaceService) {
				newArg := strings.Replace(dep.Spec.Template.Spec.Containers[containerIdx].Args[index], util.NamespaceService, ns+".svc", -1)
				dep.Spec.Template.Spec.Containers[containerIdx].Args[index] = newArg
			}
		}
	}
}
func AssignServiceToDaemonSet(ds *appsv1.DaemonSet, ns string) {
	if len(ds.Spec.Template.Spec.Containers[0].Env) > 0 {
		for index, value := range ds.Spec.Template.Spec.Containers[0].Env {
			if strings.Contains(value.String(), util.NamespaceService) {
				ds.Spec.Template.Spec.Containers[0].Env[index].Value = strings.Replace(ds.Spec.Template.Spec.Containers[0].Env[index].Value, util.NamespaceService, ns+".svc", -1)
			}
		}
	}

	for containerIdx, _ := range ds.Spec.Template.Spec.Containers {
		for index, value := range ds.Spec.Template.Spec.Containers[containerIdx].Args {
			if strings.Contains(value, util.NamespaceService) {
				newArg := strings.Replace(ds.Spec.Template.Spec.Containers[containerIdx].Args[index], util.NamespaceService, ns+".svc", -1)
				ds.Spec.Template.Spec.Containers[containerIdx].Args[index] = newArg
			}
		}
	}
}
func AssignServiceToConfigMap(cm *corev1.ConfigMap, ns string) {
	if strings.Contains(cm.Data[util.OriginComfigMapPrometheusLocation], util.NamespaceService) {
		cm.Data[util.OriginComfigMapPrometheusLocation] = strings.Replace(cm.Data[util.OriginComfigMapPrometheusLocation], util.NamespaceService, ns+".svc", -1)
	}
}

func UpdateEnvVarsToDeployment(dep *appsv1.Deployment, envVars []corev1.EnvVar) {

	for containerIndex, container := range dep.Spec.Template.Spec.Containers {
		for _, envVar := range envVars {
			exist := false
			for envIndex, containerEnvVar := range container.Env {
				if envVar.Name == containerEnvVar.Name {
					exist = true
					dep.Spec.Template.Spec.Containers[containerIndex].Env[envIndex] = envVar
					break
				}
			}
			if !exist {
				dep.Spec.Template.Spec.Containers[containerIndex].Env = append(
					dep.Spec.Template.Spec.Containers[containerIndex].Env,
					envVar,
				)
			}
		}
	}
}
