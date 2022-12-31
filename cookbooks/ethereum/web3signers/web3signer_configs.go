package web3signer_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

var (
	EphemeryWeb3SignerClusterClassName = "ephemeryWeb3SignerCluster"

	web3SignerComponentBaseName = "web3signer"
	web3SignerSkeletonBaseName  = "web3signer"
	choreographySkeletonBase    = "choreography"
)

var cd = zeus_req_types.ClusterTopologyDeployRequest{
	ClusterClassName:    EphemeryWeb3SignerClusterClassName,
	SkeletonBaseOptions: []string{web3SignerSkeletonBaseName, choreographySkeletonBase},
	CloudCtxNs:          ValidatorCloudCtxNs,
}

var DeployConsensusValidatorClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: ValidatorCloudCtxNs,
}

var ValidatorCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "ephemery-web3signer", // set with your own namespace
	Env:           "production",
}

// chart workload metadata
var validatorsChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      web3SignerSkeletonBaseName,
	ChartName:         web3SignerSkeletonBaseName,
	ChartDescription:  web3SignerSkeletonBaseName,
	Version:           fmt.Sprintf("web3signerBase-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  web3SignerSkeletonBaseName,
	ComponentBaseName: web3SignerComponentBaseName,
	ClusterClassName:  EphemeryWeb3SignerClusterClassName,
	Tag:               "latest",
}

var validatorsChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/web3signer/infra",
	DirOut:      "./ethereum/validators/infra/processed_web3signers",
	FnIn:        web3SignerSkeletonBaseName, // filename for your gzip workload
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}
