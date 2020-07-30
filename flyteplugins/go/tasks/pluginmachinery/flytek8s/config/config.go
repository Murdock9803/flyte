// This package contains configuration for the flytek8s module.
// This config is under the subsection `k8s` and registered under the Plugin config
// All K8s based plugins can optionally use the flytek8s module and this configuration allows controlling the defaults
// For example if for every container execution if some default Environment Variables or Annotations should be used, then they can be configured here
// An important configuration is ResourceTolerations that are applied to every container execution that needs some resource on the cluster
package config

import (
	"time"

	config2 "github.com/lyft/flytestdlib/config"
	v1 "k8s.io/api/core/v1"

	"github.com/lyft/flyteplugins/go/tasks/config"
)

//go:generate pflags K8sPluginConfig --default-var=defaultK8sConfig

const k8sPluginConfigSectionKey = "k8s"
const defaultCPURequest = "1000m"
const defaultMemoryRequest = "1024Mi"

var (
	defaultK8sConfig = K8sPluginConfig{
		DefaultAnnotations: map[string]string{
			"cluster-autoscaler.kubernetes.io/safe-to-evict": "false",
		},
		CoPilot: FlyteCoPilotConfig{
			NamePrefix:           "flyte-copilot-",
			Image:                "docker.pkg.github.com/lyft/flyteplugins/operator:v0.4.0",
			DefaultInputDataPath: "/var/flyte/inputs",
			InputVolumeName:      "flyte-inputs",
			DefaultOutputPath:    "/var/flyte/outputs",
			OutputVolumeName:     "flyte-outputs",
			CPU:                  "500m",
			Memory:               "128Mi",
			StartTimeout: config2.Duration{
				Duration: time.Second * 60,
			},
		},
		DefaultCPURequest:    defaultCPURequest,
		DefaultMemoryRequest: defaultMemoryRequest,
	}

	// K8sPluginConfigSection provides a singular top level config section for all plugins.
	// If you are a plugin developer writing a k8s plugin, register your config section as a subsection to this.
	K8sPluginConfigSection = config.MustRegisterSubSection(k8sPluginConfigSectionKey, &defaultK8sConfig)
)

// Top level k8s plugin config.
type K8sPluginConfig struct {
	// Boolean flag that indicates if a finalizer should be injected into every K8s resource launched
	InjectFinalizer bool `json:"inject-finalizer" pflag:",Instructs the plugin to inject a finalizer on startTask and remove it on task termination."`
	// Provide default annotations that should be added to K8s resource
	DefaultAnnotations map[string]string `json:"default-annotations" pflag:"-,Defines a set of default annotations to add to the produced pods."`
	// Provide default labels that should be added to K8s resource
	DefaultLabels map[string]string `json:"default-labels" pflag:"-,Defines a set of default labels to add to the produced pods."`
	// Provide additional environment variable pairs that plugin authors will provide to containers
	DefaultEnvVars map[string]string `json:"default-env-vars" pflag:"-,Additional environment variable that should be injected into every resource"`
	// Provide additional environment variable pairs whose values resolve from the plugin's execution environment.
	DefaultEnvVarsFromEnv map[string]string `json:"default-env-vars-from-env" pflag:"-,Additional environment variable that should be injected into every resource"`
	// Tolerations in the cluster that should be applied for a specific resource
	// Currently we support simple resource based tolerations only
	ResourceTolerations map[v1.ResourceName][]v1.Toleration `json:"resource-tolerations"  pflag:"-,Default tolerations to be applied for resource of type 'key'"`
	// default cpu requests for a container
	DefaultCPURequest string `json:"default-cpus" pflag:",Defines a default value for cpu for containers if not specified."`
	// default memory requests for a container
	DefaultMemoryRequest string `json:"default-memory" pflag:",Defines a default value for memory for containers if not specified."`
	// Tolerations for interruptible k8s pods: These tolerations are added to the pods that can tolerate getting evicted from a node. We
	// can leverage this for better bin-packing and using low-reliability cheaper machines.
	InterruptibleTolerations []v1.Toleration `json:"interruptible-tolerations"  pflag:"-,Tolerations to be applied for interruptible pods"`
	// Node Selector Labels for interruptible pods: Similar to InterruptibleTolerations, these node selector labels are added for pods that can tolerate
	// eviction.
	InterruptibleNodeSelector map[string]string `json:"interruptible-node-selector" pflag:"-,Defines a set of node selector labels to add to the interruptible pods."`
	// Default scheduler that should be used for all pods or CRD that accept Scheduler name.
	SchedulerName string `json:"scheduler-name" pflag:",Defines scheduler name."`
	// Flyte CoPilot Configuration
	CoPilot FlyteCoPilotConfig `json:"co-pilot" pflag:",Co-Pilot Configuration"`
}

type FlyteCoPilotConfig struct {
	// Co-pilot sidecar container name
	NamePrefix string `json:"name" pflag:",Flyte co-pilot sidecar container name prefix. (additional bits will be added after this)"`
	// Docker image FQN where co-pilot binary is installed
	Image string `json:"image" pflag:",Flyte co-pilot Docker Image FQN"`
	// Default Input Path for every task execution that uses co-pilot. This is used only if a input path is not provided by the user and inputs are required for the task
	DefaultInputDataPath string `json:"default-input-path" pflag:",Default path where the volume should be mounted"`
	// Default Output Path for every task execution that uses co-pilot. This is used only if a output path is not provided by the user and outputs are required for the task
	DefaultOutputPath string `json:"default-output-path" pflag:",Default path where the volume should be mounted"`
	// Name of the input volume
	InputVolumeName string `json:"input-vol-name" pflag:",Name of the data volume that is created for storing inputs"`
	// Name of the output volume
	OutputVolumeName string `json:"output-vol-name" pflag:",Name of the data volume that is created for storing outputs"`
	// Time for which the sidecar container should wait after starting up, for the primary process to appear. If it does not show up in this time
	// the process will be assumed to be dead or in a terminal condition and will trigger an abort.
	StartTimeout config2.Duration `json:"start-timeout" pflag:"-,Time for which the sidecar should wait on startup before assuming the primary container to have failed startup."`
	// Resources for CoPilot Containers
	CPU     string `json:"cpu" pflag:",Used to set cpu for co-pilot containers"`
	Memory  string `json:"memory" pflag:",Used to set memory for co-pilot containers"`
	Storage string `json:"storage" pflag:",Default storage limit for individual inputs / outputs"`
}

// Retrieves the current k8s plugin config or default.
func GetK8sPluginConfig() *K8sPluginConfig {
	return K8sPluginConfigSection.GetConfig().(*K8sPluginConfig)
}

// [FOR TESTING ONLY] Sets current value for the config.
func SetK8sPluginConfig(cfg *K8sPluginConfig) error {
	return K8sPluginConfigSection.SetConfig(cfg)
}
