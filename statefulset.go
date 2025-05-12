package debuggerkubernetesssa

import (
	"context"

	"k8s.io/apimachinery/pkg/api/equality"
	appsv1apply "k8s.io/client-go/applyconfigurations/apps/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

func DebugStatefulsetApplyConfiguration(ctx context.Context, oldStatefulsetApplyConfig *appsv1apply.StatefulSetApplyConfiguration, newStatefulsetApplyConfig *appsv1apply.StatefulSetApplyConfiguration) {
	logger := logf.FromContext(ctx)

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Replicas, newStatefulsetApplyConfig.Spec.Replicas) {
		logger.Info("statefulset spec.replicas are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Selector, newStatefulsetApplyConfig.Spec.Selector) {
		logger.Info("statefulset spec.selector are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template, newStatefulsetApplyConfig.Spec.Template) {
		logger.Info("statefulset spec.template.are not equal")

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Name, newStatefulsetApplyConfig.Spec.Template.Name) {
			logger.Info("statefulset spec.template.name are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.GenerateName, newStatefulsetApplyConfig.Spec.Template.GenerateName) {
			logger.Info("statefulset spec.template.containers are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Namespace, newStatefulsetApplyConfig.Spec.Template.Namespace) {
			logger.Info("statefulset spec.template.namespace are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.UID, newStatefulsetApplyConfig.Spec.Template.UID) {
			logger.Info("statefulset spec.template.uid are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.ResourceVersion, newStatefulsetApplyConfig.Spec.Template.ResourceVersion) {
			logger.Info("statefulset spec.template.resource version are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Generation, newStatefulsetApplyConfig.Spec.Template.Generation) {
			logger.Info("statefulset spec.template.generation are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.CreationTimestamp, newStatefulsetApplyConfig.Spec.Template.CreationTimestamp) {
			logger.Info("statefulset spec.template.creation timestamp are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.DeletionTimestamp, newStatefulsetApplyConfig.Spec.Template.DeletionTimestamp) {
			logger.Info("statefulset spec.template.deletion timestamp are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.DeletionGracePeriodSeconds, newStatefulsetApplyConfig.Spec.Template.DeletionGracePeriodSeconds) {
			logger.Info("statefulset spec.template.deletion grace period seconds are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Labels, newStatefulsetApplyConfig.Spec.Template.Labels) {
			logger.Info("statefulset spec.template.labels are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Annotations, newStatefulsetApplyConfig.Spec.Template.Annotations) {
			logger.Info("statefulset spec.template.annotations are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.OwnerReferences, newStatefulsetApplyConfig.Spec.Template.OwnerReferences) {
			logger.Info("statefulset spec.template.owner references are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Finalizers, newStatefulsetApplyConfig.Spec.Template.Finalizers) {
			logger.Info("statefulset spec.template.finalizers are not equal")
		}

		if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec, newStatefulsetApplyConfig.Spec.Template.Spec) {
			logger.Info("statefulset spec.template.spec are not equal")

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes) {
				logger.Info("statefulset spec.template.spec.volumes are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.InitContainers, newStatefulsetApplyConfig.Spec.Template.Spec.InitContainers) {
				logger.Info("statefulset spec.template.spec.init containers are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Containers, newStatefulsetApplyConfig.Spec.Template.Spec.Containers) {
				logger.Info("statefulset spec.template.spec.containers are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.EphemeralContainers, newStatefulsetApplyConfig.Spec.Template.Spec.EphemeralContainers) {
				logger.Info("statefulset spec.template.spec.ephemeral containers are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.RestartPolicy, newStatefulsetApplyConfig.Spec.Template.Spec.RestartPolicy) {
				logger.Info("statefulset spec.template.spec.restart policy are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.TerminationGracePeriodSeconds, newStatefulsetApplyConfig.Spec.Template.Spec.TerminationGracePeriodSeconds) {
				logger.Info("statefulset spec.template.spec.termination grace period seconds are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.ActiveDeadlineSeconds, newStatefulsetApplyConfig.Spec.Template.Spec.ActiveDeadlineSeconds) {
				logger.Info("statefulset spec.template.spec.active deadline seconds are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.DNSPolicy, newStatefulsetApplyConfig.Spec.Template.Spec.DNSPolicy) {
				logger.Info("statefulset spec.template.spec.dns policy are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.NodeSelector, newStatefulsetApplyConfig.Spec.Template.Spec.NodeSelector) {
				logger.Info("statefulset spec.template.spec.node selector are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.ServiceAccountName, newStatefulsetApplyConfig.Spec.Template.Spec.ServiceAccountName) {
				logger.Info("statefulset spec.template.spec.service account name are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.DeprecatedServiceAccount, newStatefulsetApplyConfig.Spec.Template.Spec.DeprecatedServiceAccount) {
				logger.Info("statefulset spec.template.spec.deprecated service account are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.AutomountServiceAccountToken, newStatefulsetApplyConfig.Spec.Template.Spec.AutomountServiceAccountToken) {
				logger.Info("statefulset spec.template.spec.automount service account token are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.NodeName, newStatefulsetApplyConfig.Spec.Template.Spec.NodeName) {
				logger.Info("statefulset spec.template.spec.node name are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.HostNetwork, newStatefulsetApplyConfig.Spec.Template.Spec.HostNetwork) {
				logger.Info("statefulset spec.template.spec.host network are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.HostPID, newStatefulsetApplyConfig.Spec.Template.Spec.HostPID) {
				logger.Info("statefulset spec.template.spec.host pid are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.HostIPC, newStatefulsetApplyConfig.Spec.Template.Spec.HostIPC) {
				logger.Info("statefulset spec.template.spec.host ipc are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.ShareProcessNamespace, newStatefulsetApplyConfig.Spec.Template.Spec.ShareProcessNamespace) {
				logger.Info("statefulset spec.template.spec.share process namespace are not equal")
			}
			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.SecurityContext, newStatefulsetApplyConfig.Spec.Template.Spec.SecurityContext) {
				logger.Info("statefulset spec.template.spec.security context are not equal")
			}
			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.ImagePullSecrets, newStatefulsetApplyConfig.Spec.Template.Spec.ImagePullSecrets) {
				logger.Info("statefulset spec.template.spec.image pull secrets are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Hostname, newStatefulsetApplyConfig.Spec.Template.Spec.Hostname) {
				logger.Info("statefulset spec.template.spec.hostname are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Subdomain, newStatefulsetApplyConfig.Spec.Template.Spec.Subdomain) {
				logger.Info("statefulset spec.template.spec.subdomain are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Affinity, newStatefulsetApplyConfig.Spec.Template.Spec.Affinity) {
				logger.Info("statefulset spec.template.spec.affinity are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.SchedulerName, newStatefulsetApplyConfig.Spec.Template.Spec.SchedulerName) {
				logger.Info("statefulset spec.template.spec.scheduler name are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Tolerations, newStatefulsetApplyConfig.Spec.Template.Spec.Tolerations) {
				logger.Info("statefulset spec.template.spec.tolerations are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.HostAliases, newStatefulsetApplyConfig.Spec.Template.Spec.HostAliases) {
				logger.Info("statefulset spec.template.spec.host aliases are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.PriorityClassName, newStatefulsetApplyConfig.Spec.Template.Spec.PriorityClassName) {
				logger.Info("statefulset spec.template.spec.priority class name are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Priority, newStatefulsetApplyConfig.Spec.Template.Spec.Priority) {
				logger.Info("statefulset spec.template.spec.priority are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.DNSConfig, newStatefulsetApplyConfig.Spec.Template.Spec.DNSConfig) {
				logger.Info("statefulset spec.template.spec.dns config are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.ReadinessGates, newStatefulsetApplyConfig.Spec.Template.Spec.ReadinessGates) {
				logger.Info("statefulset spec.template.spec.readiness gates are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.RuntimeClassName, newStatefulsetApplyConfig.Spec.Template.Spec.RuntimeClassName) {
				logger.Info("statefulset spec.template.spec.runtime class name are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.EnableServiceLinks, newStatefulsetApplyConfig.Spec.Template.Spec.EnableServiceLinks) {
				logger.Info("statefulset spec.template.spec.enable service links are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.PreemptionPolicy, newStatefulsetApplyConfig.Spec.Template.Spec.PreemptionPolicy) {
				logger.Info("statefulset spec.template.spec.preemption policy are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Overhead, newStatefulsetApplyConfig.Spec.Template.Spec.Overhead) {
				logger.Info("statefulset spec.template.spec.overhead are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.TopologySpreadConstraints, newStatefulsetApplyConfig.Spec.Template.Spec.TopologySpreadConstraints) {
				logger.Info("statefulset spec.template.spec.topology spread constraints are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.SetHostnameAsFQDN, newStatefulsetApplyConfig.Spec.Template.Spec.SetHostnameAsFQDN) {
				logger.Info("statefulset spec.template.spec.set hostname as fqdn are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.OS, newStatefulsetApplyConfig.Spec.Template.Spec.OS) {
				logger.Info("statefulset spec.template.spec.os are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.HostUsers, newStatefulsetApplyConfig.Spec.Template.Spec.HostUsers) {
				logger.Info("statefulset spec.template.spec.host users are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.SchedulingGates, newStatefulsetApplyConfig.Spec.Template.Spec.SchedulingGates) {
				logger.Info("statefulset spec.template.spec.scheduling gates are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.ResourceClaims, newStatefulsetApplyConfig.Spec.Template.Spec.ResourceClaims) {
				logger.Info("statefulset spec.template.spec.resource claims are not equal")
			}

			if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Resources, newStatefulsetApplyConfig.Spec.Template.Spec.Resources) {
				logger.Info("statefulset spec.template.spec.resources are not equal")
			}
		}
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.VolumeClaimTemplates, newStatefulsetApplyConfig.Spec.VolumeClaimTemplates) {
		logger.Info("statefulset volume claim spec.templates are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.ServiceName, newStatefulsetApplyConfig.Spec.ServiceName) {
		logger.Info("statefulset service name are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.PodManagementPolicy, newStatefulsetApplyConfig.Spec.PodManagementPolicy) {
		logger.Info("statefulset pod management policy are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.UpdateStrategy, newStatefulsetApplyConfig.Spec.UpdateStrategy) {
		logger.Info("statefulset update strategy are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.RevisionHistoryLimit, newStatefulsetApplyConfig.Spec.RevisionHistoryLimit) {
		logger.Info("statefulset revision history limit are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.MinReadySeconds, newStatefulsetApplyConfig.Spec.MinReadySeconds) {
		logger.Info("statefulset min ready seconds are not equal")
	}
	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Finalizers, newStatefulsetApplyConfig.Spec.Template.Finalizers) {
		logger.Info("statefulset spec.template.finalizers are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.MinReadySeconds, newStatefulsetApplyConfig.Spec.MinReadySeconds) {
		logger.Info("statefulset min ready seconds are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.PersistentVolumeClaimRetentionPolicy, newStatefulsetApplyConfig.Spec.PersistentVolumeClaimRetentionPolicy) {
		logger.Info("statefulset persistent volume claim retention policy are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Ordinals, newStatefulsetApplyConfig.Spec.Ordinals) {
		logger.Info("statefulset volume claim spec.templates are not equal")
	}
}
