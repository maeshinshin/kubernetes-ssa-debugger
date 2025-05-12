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

				if len(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes) != len(newStatefulsetApplyConfig.Spec.Template.Spec.Volumes) {
					logger.Info("statefulset spec.template.spec.volumes length are not equal")
				} else {
					for i := range len(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes) {
						if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i], newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i]) {
							logger.Info("statefulset spec.template.spec.volumes are not equal")

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Name, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Name) {
								logger.Info("statefulset spec.template.spec.volumes.name are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Name, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Name) {
								logger.Info("statefulset spec.template.spec.volumes.name are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].HostPath, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].HostPath) {
								logger.Info("statefulset spec.template.spec.volumes.hostpath are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].EmptyDir, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].EmptyDir) {
								logger.Info("statefulset spec.template.spec.volumes.emptydir are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].GCEPersistentDisk, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].GCEPersistentDisk) {
								logger.Info("statefulset spec.template.spec.volumes.gcepersistentdisk are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].AWSElasticBlockStore, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].AWSElasticBlockStore) {
								logger.Info("statefulset spec.template.spec.volumes.awselasticblockstore are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].GitRepo, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].GitRepo) {
								logger.Info("statefulset spec.template.spec.volumes.gitrepo are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Secret, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Secret) {
								logger.Info("statefulset spec.template.spec.volumes.secret are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].NFS, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].NFS) {
								logger.Info("statefulset spec.template.spec.volumes.nfs are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].ISCSI, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].ISCSI) {
								logger.Info("statefulset spec.template.spec.volumes.iscsi are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Glusterfs, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Glusterfs) {
								logger.Info("statefulset spec.template.spec.volumes.glusterfs are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].PersistentVolumeClaim, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].PersistentVolumeClaim) {
								logger.Info("statefulset spec.template.spec.volumes.persistentvolumeclaim are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].RBD, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].RBD) {
								logger.Info("statefulset spec.template.spec.volumes.rbd are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].FlexVolume, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].FlexVolume) {
								logger.Info("statefulset spec.template.spec.volumes.flexvolume are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Cinder, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Cinder) {
								logger.Info("statefulset spec.template.spec.volumes.cinder are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].CephFS, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].CephFS) {
								logger.Info("statefulset spec.template.spec.volumes.cephfs are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Flocker, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Flocker) {
								logger.Info("statefulset spec.template.spec.volumes.flocker are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].DownwardAPI, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].DownwardAPI) {
								logger.Info("statefulset spec.template.spec.volumes.downwardapi are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].FC, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].FC) {
								logger.Info("statefulset spec.template.spec.volumes.fc are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].AzureFile, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].AzureFile) {
								logger.Info("statefulset spec.template.spec.volumes.azurefile are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].ConfigMap, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].ConfigMap) {
								logger.Info("statefulset spec.template.spec.volumes.configmap are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].VsphereVolume, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].VsphereVolume) {
								logger.Info("statefulset spec.template.spec.volumes.vspherevolume are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Quobyte, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Quobyte) {
								logger.Info("statefulset spec.template.spec.volumes.quobyte are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].AzureDisk, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].AzureDisk) {
								logger.Info("statefulset spec.template.spec.volumes.azuredisk are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].PhotonPersistentDisk, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].PhotonPersistentDisk) {
								logger.Info("statefulset spec.template.spec.volumes.photonpersistentdisk are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Projected, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Projected) {
								logger.Info("statefulset spec.template.spec.volumes.projected are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].PortworxVolume, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].PortworxVolume) {
								logger.Info("statefulset spec.template.spec.volumes.portworxvolume are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].ScaleIO, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].ScaleIO) {
								logger.Info("statefulset spec.template.spec.volumes.scaleio are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].StorageOS, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].StorageOS) {
								logger.Info("statefulset spec.template.spec.volumes.storageos are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].CSI, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].CSI) {
								logger.Info("statefulset spec.template.spec.volumes.csi are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Ephemeral, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Ephemeral) {
								logger.Info("statefulset spec.template.spec.volumes.ephemeral are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Image, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Image) {
								logger.Info("statefulset spec.template.spec.volumes.image are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Projected, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Projected) {
								logger.Info("statefulset spec.template.spec.volumes.projected are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].PortworxVolume, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].PortworxVolume) {
								logger.Info("statefulset spec.template.spec.volumes.portworxvolume are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].ScaleIO, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].ScaleIO) {
								logger.Info("statefulset spec.template.spec.volumes.scaleio are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].StorageOS, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].StorageOS) {
								logger.Info("statefulset spec.template.spec.volumes.storageos are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].CSI, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].CSI) {
								logger.Info("statefulset spec.template.spec.volumes.csi are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Ephemeral, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Ephemeral) {
								logger.Info("statefulset spec.template.spec.volumes.ephemeral are not equal")
							}

							if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Image, newStatefulsetApplyConfig.Spec.Template.Spec.Volumes[i].Image) {
								logger.Info("statefulset spec.template.spec.volumes.image are not equal")
							}
						}
					}
				}

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
		logger.Info("statefulset spec.volumeclametemplates are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.ServiceName, newStatefulsetApplyConfig.Spec.ServiceName) {
		logger.Info("statefulset spec.servicename are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.PodManagementPolicy, newStatefulsetApplyConfig.Spec.PodManagementPolicy) {
		logger.Info("statefulset spec.podmanagementpolicy are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.UpdateStrategy, newStatefulsetApplyConfig.Spec.UpdateStrategy) {
		logger.Info("statefulset spec.updatestrategy are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.RevisionHistoryLimit, newStatefulsetApplyConfig.Spec.RevisionHistoryLimit) {
		logger.Info("statefulset spec.revisionhistorylimit are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.MinReadySeconds, newStatefulsetApplyConfig.Spec.MinReadySeconds) {
		logger.Info("statefulset spec.minreadyseconds are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Template.Finalizers, newStatefulsetApplyConfig.Spec.Template.Finalizers) {
		logger.Info("statefulset spec.template.finalizers are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.MinReadySeconds, newStatefulsetApplyConfig.Spec.MinReadySeconds) {
		logger.Info("statefulset spec.minreadyseconds are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.PersistentVolumeClaimRetentionPolicy, newStatefulsetApplyConfig.Spec.PersistentVolumeClaimRetentionPolicy) {
		logger.Info("statefulset spec.persistentvolumeclaimretentionpolicy are not equal")
	}

	if !equality.Semantic.DeepEqual(oldStatefulsetApplyConfig.Spec.Ordinals, newStatefulsetApplyConfig.Spec.Ordinals) {
		logger.Info("statefulset spec.ordinals are not equal")
	}
}
