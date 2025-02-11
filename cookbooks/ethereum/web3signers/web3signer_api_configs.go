package web3signer_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

var (
	Web3SignerExternalAPIClusterClassName        = "web3SignerAPI"
	Web3SignerExternalAPIClusterBaseName         = "web3signerAPI"
	Web3SignerExternalAPIClusterSkeletonBaseName = "web3signerAPI"

	Web3SignerExternalAPIClusterIngressBaseName         = "web3SignerIngress"
	Web3SignerExternalAPIClusterIngressSkeletonBaseName = "web3SignerIngress"
)

var Web3SignerExternalAPIClusterDefinition = zeus_req_types.ClusterTopologyDeployRequest{
	ClusterClassName: Web3SignerExternalAPIClusterClassName,
	SkeletonBaseOptions: []string{
		Web3SignerExternalAPIClusterSkeletonBaseName, Web3SignerExternalAPIClusterIngressSkeletonBaseName,
	},
	CloudCtxNs: Web3SignerExternalAPICloudCtxNs,
}

var Web3SignerExternalAPICloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "nyc1",
	Context:       "do-nyc1-do-nyc1-zeus-demo",
	Namespace:     "web3signer",
	Env:           "dev",
}

var DeployWeb3SignerExternalAPIKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: Web3SignerExternalAPICloudCtxNs,
}

var Web3SignerAPIChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      Web3SignerExternalAPIClusterClassName,
	ChartName:         Web3SignerExternalAPIClusterClassName,
	ChartDescription:  Web3SignerExternalAPIClusterClassName,
	Version:           fmt.Sprintf("web3SignerAPI-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  Web3SignerExternalAPIClusterSkeletonBaseName,
	ComponentBaseName: Web3SignerExternalAPIClusterBaseName,
	ClusterClassName:  Web3SignerExternalAPIClusterClassName,
	Tag:               "latest",
}

var Web3SignerAPIChartChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/web3signers/infra/consensys_web3signer_custom_config",
	DirOut:      "./ethereum/web3signers/infra/processed_consensys_web3signer",
	FnIn:        Web3SignerExternalAPIClusterClassName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

var Web3SignerIngressChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "web3SignerIngress",
	ChartName:         "web3SignerIngress",
	ChartDescription:  "web3SignerIngress",
	Version:           fmt.Sprintf("web3SignerIngress-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  Web3SignerExternalAPIClusterIngressSkeletonBaseName,
	ComponentBaseName: Web3SignerExternalAPIClusterIngressBaseName,
	ClusterClassName:  Web3SignerExternalAPIClusterClassName,
	Tag:               "latest",
}

var Web3SignerIngressChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/web3signers/infra/ingress",
	DirOut:      "./ethereum/web3signers/infra/processed_consensys_web3signer",
	FnIn:        Web3SignerExternalAPIClusterIngressSkeletonBaseName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}
