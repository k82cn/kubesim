module volcano.sh/kubesim

go 1.15

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/mock v1.3.1
	github.com/google/cadvisor v0.35.0
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.6.2-0.20200628121210-87a988cffbb9
	google.golang.org/grpc v1.26.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	k8s.io/component-base v0.18.6
	k8s.io/cri-api v0.18.6
	k8s.io/klog v1.0.0
	k8s.io/kubernetes v1.18.6
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
)

replace (
	k8s.io/api => k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
	k8s.io/apiserver => k8s.io/apiserver v0.18.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.6
	k8s.io/code-generator => k8s.io/code-generator v0.18.6
	k8s.io/component-base => k8s.io/component-base v0.18.6
	k8s.io/cri-api => k8s.io/cri-api v0.18.6
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.6
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.6
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.6
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.6
	k8s.io/kubectl => k8s.io/kubectl v0.18.6
	k8s.io/kubelet => k8s.io/kubelet v0.18.6
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.6
	k8s.io/metrics => k8s.io/metrics v0.18.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.6
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.18.6
	k8s.io/sample-controller => k8s.io/sample-controller v0.18.6
)
