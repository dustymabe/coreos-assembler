module github.com/coreos/coreos-assembler

go 1.19

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.3.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork v1.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.0.0
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210419045805-b50610722085
	github.com/IBM/ibm-cos-sdk-go v1.6.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1442
	github.com/aliyun/aliyun-oss-go-sdk v2.0.4+incompatible
	github.com/aws/aws-sdk-go v1.44.298
	github.com/coreos/butane v0.18.1-0.20230412230143-79c207705ee4
	github.com/coreos/go-semver v0.3.1
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
	github.com/coreos/go-systemd/v22 v22.5.0
	github.com/coreos/ignition/v2 v2.16.2
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/coreos/stream-metadata-go v0.4.3
	github.com/coreos/vcontext v0.0.0-20230201181013-d72178a18687
	github.com/digitalocean/go-qemu v0.0.0-20200529005954-1b453d036a9c
	github.com/digitalocean/godo v1.33.0
	github.com/frostschutz/go-fibmap v0.0.0-20160825162329-b32c231bfe6a
	github.com/gophercloud/gophercloud v0.22.0
	github.com/gophercloud/utils v0.0.0-20210323225332-7b186010c04f
	github.com/kballard/go-shellquote v0.0.0-20150810074751-d8ec1a69a250
	github.com/kylelemons/godebug v1.1.0
	github.com/onsi/gomega v1.23.0
	github.com/packethost/packngo v0.0.0-20180426081943-80f62d78849d
	github.com/pborman/uuid v1.2.1
	github.com/pin/tftp v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.6.0
	github.com/vincent-petithory/dataurl v1.0.0
	github.com/vishvananda/netlink v0.0.0-20150710184826-9cff81214893
	github.com/vishvananda/netns v0.0.0-20150710222425-604eaf189ee8
	github.com/vmware/govmomi v0.15.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/crypto v0.11.0
	golang.org/x/net v0.12.0
	golang.org/x/oauth2 v0.10.0
	golang.org/x/sys v0.10.0
	golang.org/x/term v0.10.0
	google.golang.org/api v0.130.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/api v0.27.1
	k8s.io/apimachinery v0.27.1
	kubevirt.io/api v0.0.0-20230706190111-5527663af491
	kubevirt.io/client-go v1.0.0
)

require (
	cloud.google.com/go/compute v1.20.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.1.1 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v0.8.1 // indirect
	github.com/clarketm/json v1.17.1 // indirect
	github.com/coreos/go-json v0.0.0-20230131223807-18775e0fb4fb // indirect
	github.com/coreos/prometheus-operator v0.38.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/digitalocean/go-libvirt v0.0.0-20200810224808-b9c702499bf7 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.1 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/godbus/dbus/v5 v5.0.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.5 // indirect
	github.com/googleapis/gax-go/v2 v2.11.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v0.0.0-20191119172530-79f836b90111 // indirect
	github.com/kubernetes-csi/external-snapshotter/client/v4 v4.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/openshift/api v0.0.0-20230503133300-8bbcb7ca7183 // indirect
	github.com/openshift/client-go v0.0.0-20210112165513-ebc401615f47 // indirect
	github.com/openshift/custom-resource-status v1.1.2 // indirect
	github.com/pkg/browser v0.0.0-20210115035449-ce105d075bb4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/spf13/pflag v1.0.6-0.20210604193023-d5e0c0615ace // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/grpc v1.56.1 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	k8s.io/apiextensions-apiserver v0.26.3 // indirect
	k8s.io/client-go v12.0.0+incompatible // indirect
	k8s.io/klog/v2 v2.90.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230501164219-8b0f38b5fd1f // indirect
	k8s.io/utils v0.0.0-20230505201702-9f6742963106 // indirect
	kubevirt.io/containerized-data-importer-api v1.57.0-alpha1 // indirect
	kubevirt.io/controller-lifecycle-operator-sdk/api v0.0.0-20220329064328-f3cc58c6ed90 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

// Perform replacements os we can include kubevirt.io/client-go code
// in mantle. This shouldn't be needed but there were problems
// importing the code from https://github.com/kubevirt/client-go.git
// without this:
//      go: warning: k8s.io/client-go@v12.0.0+incompatible: retracted by module author: pre-module versions are obsolete
// There are various open issues in https://github.com/kubevirt/client-go.git
// about this problem but no current solutions. These replacements
// were taken directly from https://github.com/kubevirt/client-go/blob/6a56296474b2b2ff4c65d38d2ce1bb2761501b8b/go.mod#L74-L99
replace (
	k8s.io/api => k8s.io/api v0.26.3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.26.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.26.3
	k8s.io/apiserver => k8s.io/apiserver v0.26.3
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.26.3
	k8s.io/client-go => k8s.io/client-go v0.26.3
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.26.3
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.26.3
	k8s.io/code-generator => k8s.io/code-generator v0.26.3
	k8s.io/component-base => k8s.io/component-base v0.26.3
	k8s.io/cri-api => k8s.io/cri-api v0.26.3
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.26.3
	k8s.io/klog => k8s.io/klog v0.4.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.26.3
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.26.3
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.26.3
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.26.3
	k8s.io/kubectl => k8s.io/kubectl v0.26.3
	k8s.io/kubelet => k8s.io/kubelet v0.26.3
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.26.3
	k8s.io/metrics => k8s.io/metrics v0.26.3
	k8s.io/node-api => k8s.io/node-api v0.26.3
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.26.3
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.26.3
	k8s.io/sample-controller => k8s.io/sample-controller v0.26.3
)
