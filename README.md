## Zeus ##

Zeus is an evolution of web container orchestration into web ecosystems orchestration. It changes the paradigm into one that unifies configuration with the underlying orchestrator, instead of decoupling them into confusing messes like Helm, Terraform, and GitOps, while also reducing the operational complexity of building with Kubernetes significantly and is capable of building systems at the scale and complexity of AWS (without the terrible UX/UI) by unifying multi-network Kubernetes node pools with SQL driven relationship building & querying. 

We're not just a web3 company, we have a lot of experience in crypto cloud tech so that's why it's our first product line theme, we'll be offering many  web2 focused products by the end of the year in addition to advancing our web3 products.

## Upcoming Features Overview ##

### Event driven signature automation for Ethereum staking, and web3 dapp interactions

#### Timeline ~ End of January/Early Feb 23'

Event driven signature automation for Ethereum staking using our synthetic staking beacon, an in-house technology invention that dramatically lowers the infrastucture costs of traditional enterprise staking architectures by 100x+ and bundles middleware like slashing protection and mev into the service. Which also comes with the added benefit of letting you stake from your wallet without anyone having access to your signing or withdrawal keys and without any infrastructure setup, with only a few lines of code. For those who do want hosted cloud signers, you'll have flexible 1-click style deployable hosted signers.

#### How much will staking services cost for Ethereum?

$10/mo solo or large scale enterprise staking for Ethereum per validator, all inclusive, and unlike everyone else in the industry, we're not taking a cut of your staking yield, or locking you in to any contracts, smart contracts, or anything else. We'd rather give you a better experience in web3, one that is fun to interact with, helps you build communities, helps attracts new users to the web3 ecosystem, simply by making it user-friendly, and we're betting that is worth a lot more than extracting as much money as we can from our users.

#### Do you have any limits on how many validators you can support, are you also supporting direct to consumer?

Our system can comfortably scale to all of mainnet and far beyond, the cost difference for cloud, devops, etc costs to us between even 50k->100k->200k or more validators may as well be virtually zero to help illustrate how efficient it is. We're starting with wallet custodians, and B2B partners first since we don't have large enough CX support for large request volumes at this moment, but targeting direct to consumer before end of Q2, possibly much sooner. We'll setup an email signup list soon for you all to get notified over the next 1-2 months.

#### What if this results in centralization of staking through this platform?

Great question, and since that was a concern we've had on our minds early on, we've already come up with a with a few feasible designs that could turn this entire synthetic staking beacon technology into a fully decentralized synthetic beacon service. We'll advance our tech to that stage if it starts trending that way, we're a long ways from that though at this stage.

#### Can you share more about the native web2 interactions?

Both signer options come with powerful tx orchestration capabilites for smart contract & dapp interactions, allowing for truly native web2 experiences for dapp interactions for the first time (2fa, email login linked to your wallet, QR scanning, and more slated to follow this release, timelines and additional details to be announced in Feb). It will also allow developers to build smart contract automation flows for users that interact with their wallets behind the scenes, unlike the clunky browser extension setups used today that are entirely too frustrating to use.

### Automated web3 infrastructure setup

#### Timeline ~ Late January 23'

Completely orchestrated and automated web3 infra. Sets up infrastructure on demand, automates run books, sets up mev-boost, web3signer, adds snapshot chain download integration, automates devops that's done by hand today, automates interactions for seeding and withdrawing validators, automates upgrades, automates configuration setup and verification, automates notifications and rewards info. Enables web3 staking infrastructure to be portable across cloud, and for vendor switching on demand. Starting with Ethereum.

### Automated web3 network setup for large scale private network testing

#### Timeline ~ Q1/Q2 23'

Create network from scratch that can replicate the size of mainnet, starting with Ethereum.
```
Runs Genesis -> Seeds Validators -> Deploys Validator Infra -> Metrics/Data
```
### AI driven infrastructure configuration & devops

#### V0 Generation: Timeline ~ Q4 23'- Q1 24'

AI driven infrastructure that automates infrastructure config customization & handles devops via log ingestion. Generation v0, public access will be strictly limited to early users of Zeus and a small pool of API access will be allocated to those who request access in FIFO basis.

#### V1 Generation: Timeline ~ Q1 24'- Q1 25'

The next generation will be able to read helm charts & configure infrastructure automatically for medium size complexity applications, and small networks of related applications.

## Cookbooks ##

Contains common web2 & web3 building components like ethereum infra setups with customization driven through code, blurring the line between infra configuration and app development, and contains microservice designs & patterns like api servers and injectable choreography, their setups in kubernetes, golang, docker, and startup commands, and useful tools for debugging, interacting, and automating actions.

### Cookbook Structure ###

#### Ethereum ####

#### ```protocol/components```

Contains smart contract automation kits. This testcase shows a full end-end seed, create, and deposits validators on the Ethereum ephemery testnet.

#### ```cookbooks/ethereum/automation```
#### ```cookbooks/ethereum/automation/deposits_test.go ```

Cookbook items listed by protocol & component class. Eg. Ethereum has a beacon component group. Contains Kubernetes config setup templates. Here's a few example paths. Also contains an actions folder, which does log dumps, pod restarts, configuration changes on demand for k8s applications, and more.

#### ```cookbooks/ethereum/beacons/infra/consensus_client```
#### ```cookbooks/ethereum/beacons/infra/exec_client```
#### ```cookbooks/ethereum/validators/infra/validators```
#### ```cookbooks/ethereum/web3signers/infra/consensys_web3signer```

Complete, and powerful Ethereum infra automation 

#### ```zeus/cookbooks/ethereum/beacons/beacon_cluster_test.go ```
#### ```zeus/cookbooks/ethereum/validators/validator_cluster_test.go ```
#### ```zeus/cookbooks/ethereum/web3signers/web3signer_cluster_test.go ```

See this test case to see how a beacon cluster class was created, and then extended to support choreography to reset configs on a scheduled interval for the Ephemery testnet, and then added validator clients, and then again to add web3signing integration.

#### Microservices ###

Contains full kubernetes infra setup templates for microservices, injectable choreography for clusters, and more continually being added.

#### ```zeus/cookbooks/microservices/deployments ```
#### ```zeus/cookbooks/microservices/choreography ```

## Zeus Apps & Clients ##

Core Zeus Infra Automation Client
#### ```pkg/zeus/client```

Powerful Cluster Building, Allowing for Large Scale Infra Automation, Customization, Control

#### ```pkg/zeus/cluster_config_drivers ```
#### ```pkg/zeus/system_config_drivers ```
#### ```pkg/zeus/workload_config_drivers ```

#### API Endpoints 

Documentation and code examples are found here
[API_README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/API_README.md)

How to use the test suite to setup your own api calls
[README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/README.md)

The test directory contains useful mocks and tools for interacting with the API. It also contains a useful
config-sample.yaml, convert this to config.yaml and set your bearer token here, which then allows you to
use the demo code to create your first api request in seconds

### Hera Client

#### ```pkg/hera/client```

This client uses the OpenAI API to generate code with AI. This service is available at OpenAI cost, so just pay for the token cost, otherwise it is free to use.

### Artemis ###

#### ```artemis.zeus.fyi``` ####
#### ```pkg/artemis/client```

Artemis is a tx orchestrator. It reliably submits & confirms ethereum transactions and logs their receipts. Chain with 
the in memory db for storing web3 signer keys to build highly reliable web3 api actions with other users and smart contracts. You'll need
a bearer token to use this client. A more advanced orchestrator that can handle high volume DeFi trading, which manages nonce sequences, sets up event triggers & scheduling, and has queriable event artifacts is in works, targeted release by end of Feb.

### Aegis ###

#### ```pkg/aegis/inmemdbs```
#### ```pkg/crypto```

Aegis is a library for securely managing crypto keys. It currently contains an in memory database for storing ethereum validators
and for storing ecdsa, eth1 wallet keys. 

### Hercules ###

#### ```apps/hercules``` ####
#### ```pkg/hercules/client```

Hercules is web3 middleware that manages web3 infrastructure and connections to other middleware packages, such as web3signer, chain snapshot downloading, setting up mev-boost, orchestrating transactions and client switching, key generation and management, and much more coming soon.

It also contains useful apis to debug and troubleshoot web3 infrastructure.

### Snapshots ###

#### ```apps/snapshots``` ####

Snapshot app is embedded into the hercules docker app, and it can be used as an init container to download snapshot data on new node creation.

### Hades Library

#### ```pkg/hades```

Hades is used to interact with kubernetes workloads via API, and can apply saved Zeus workloads & cookbooks onto your own in house infrastructure.

## Zeus Users ##

### Beacon API ###

Our ephemeral ethereum testnet beacon is open to anyone, no auth is requried. Large Eth quantities for testing is available on request.

### https://eth.ephemeral.zeus.fyi

### Beacon Indexer ###

#### ```https://apollo.eth.zeus.fyi```
#### ```pkg/apollo```

Users with bearer tokens are able to access our common beacon balance & status indexer at https://apollo.eth.zeus.fyi, which indexes mainnet validator balances and statuses. It only indexes from epoch 169,000+. Once the DB reaches near capacity it removes the trailing 5k or so epochs, all the previous data is archived and accessible by request. It contains an updatedAt field for validator statuses, so you can tell how recent the status update was. It follows head behind ~10 epochs, tracking the finalized checkpoint range plus some small margin. The API is limited to 10k records per request. Requesting >10k, or requesting an epoch that isn't indexed yet will usually result in a null response. Better error messages will come soon though. See the pkg section for the apollo client which shows you how to use it, and makes it easy to integrate directly using the client.

### Snapshot Downloads ###

Snapshot download urls for mainnet geth & lighthouse available on request.

## Infra configuration is in the 14th century. Introducing the 21st century printing press

By unifying the infra ecosystem tools of today it allows the acceleration of infra development exponentially overtime with each new open source cookbook & system template allowing you to glue sophisticated distributed systems together in zero time, which lets decentralized solo & small team builders compete directly against large incumbents at the edge of technology by erasing the need for large investment in infra & devops departments that are needed now to even get started.

Lastly, by solving infra configuration, mobility, and orchestration through remote API driven configuration, it allows you to run sophisticated web apps in virtually zero time on far more cost efficient bare metal cloud providers using commoditized managed kubernetes services and simple middleware we provide, which is up to 6x more cost efficient than AWS, GCP, and comparable cloud companies. Take a look for yourself, then ask yourself why you want to make Jeff Bezos richer than he already is?

#### You have alternatives

##### https://www.ibm.com/cloud/kubernetes-service
##### https://us.ovhcloud.com/public-cloud/kubernetes
##### https://www.linode.com/products/kubernetes
##### https://www.digitalocean.com/products/kubernetes
##### https://www.vultr.com/kubernetes

#### It doesn't take long to figure it out, here's some pricing links to help you out

##### https://www.ovhcloud.com/en/public-cloud/prices/
##### https://www.ibm.com/cloud/virtual-servers/pricing
##### https://instances.vantage.sh/

AWS, GCP, Azure, type cloud companies purposely have highly confusing pricing models to mislead you on costs using data transfer bills, among other sales tactics like free initial cloud usage until you're locked into their ecosystem complexity, exactly when you realize the cloud spend is outrageous. The vast majority of enterprise users simply need RAM, CPU, Bandwidth Traffic, and a few disk options like RAID setups, and HDD, SSD, NVMes. All major cloud companies have similar performance & online SLAs. Why would you want to spend 100+ engineering hours figuring out some obscure EC2 instance number that they deprecate in a year anyway? The 0.001% of people that have that need already know what they want.

### Overview 

1. Automates translation of kubernetes yaml configurations into representative SQL models
2. Users upload these infrastructure configurations via API where they are stored in the DB
3. Users can then query the contents of these infrastructure components, deploy, mutate or destroy them on demand
   
### Currently Supported Workload Types

1. Deployments
2. StatefulSets
3. Services
4. ConfigMaps
5. Ingresses

### Pods Controller 

1. GetPods
2. GetPodLogs
3. PortforwardReqToPods
4. DeletePods

Not every possible field type is supported, but the most common ones are, and even a decent amount of the uncommon ones.

See demos section for api calls, you'll need to get a bearer token from us first. More documentation to come.
Schedule a demo: https://calendly.com/alex-zeus-fyi/zeus-demo
Request a bearer token, email: alex@zeus.fyi

![Screenshot 2022-11-17 at 8 09 48 PM](https://user-images.githubusercontent.com/17446735/202614955-2708063e-1547-4dae-9332-f712102c287e.png)

### Cluster Topology Class Hierarchy Definitions ###

### Highest to Lowest Level ###

### Deployable Topologies ###

### Network Orchestration Topology ###

One or many network matrix system topologies that are combined with orchestration workflows from Artemis, and Zeus, to build complex control flows and/or sequenced network states. This could be an enterprise fleet of infrastructure, a complex devops operation done at scale, or a simulation of a large testnet, or a built-in orchstration flow for automating web3 smart contract interactions for users and sending them notifications and prompts.

### Matrix Topology ###

A multi-component cluster topology that accomplishes one or more system components on its own and combined with a Zeus injection deploys this topology onto the network. Some but not all of these topologies can be stacked with another cluster topology or a base topology to create a higher level system component.

It can be any combination of lower level system topologies or components.

### Cluster Topology ###

A fully working single component cluster topology that accomplishes one system component on its own and combined with a Zeus injection deploys this topology onto the network. Some but not all of these topologies can be stacked with another cluster topology or a base topology to create a higher level system component.

### Deployable but Incomplete Topologies ###

### Base Topology ###

A fully working single cluster topology that needs at least one other Base Topology to create a higher level component. An example would be deploying an execution client by itself post-merge on ethereum. It would be able to download chain data, but it wouldn’t be able to fulfill a useful purpose without another piece e.g. a consensus client component.  

Not deployable on its own, a mix of these is combined to create a deployable topology

### Infrastructure Base ###
An abstract atomic infrastructure base that needs a Skeleton and Configuration to create a Base Topology

### Configuration Base ###
An abstract atomic configuration base that needs an Infrastructure Base and Skeleton to create a Base Topology

### Skeleton Base ###
An abstract atomic component base that needs additional pieces to create deployable infrastructure like config map, docker image links, etc. Needs an Infrastructure and Configuration Base to create a Base Topology

