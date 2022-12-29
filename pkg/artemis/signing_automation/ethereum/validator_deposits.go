package signing_automation_ethereum

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/gochain/gochain/v4/common"
	"github.com/gochain/gochain/v4/core/types"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/web3_actions"
)

const (
	validatorDepositMethodName   = "deposit"
	validatorAbiFileLocation     = "smart_contracts/eth_deposit_contract.json"
	EphemeralDepositContractAddr = "0x4242424242424242424242424242424242424242"
	EphemeralBeacon              = "https://eth.ephemeral.zeus.fyi"
	BeaconGenesisPath            = "/eth/v1/beacon/genesis"
	BeaconForkPath               = "/eth/v1/beacon/states/head/fork"
)

func (w *Web3SignerClient) SignValidatorDepositTxToBroadcast(ctx context.Context, depositParams *DepositDataParams) (*types.Transaction, error) {
	ForceDirToEthSigningDirLocation()

	abiFile, err := ABIOpenFile(validatorAbiFileLocation)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDeposit: ABIOpenFile")
		return nil, err
	}

	pubkey, err := hex.DecodeString(strings.TrimPrefix(depositParams.PublicKey.String(), "0x"))
	if err != nil {
		panic(err)
	}
	sig, err := hex.DecodeString(strings.TrimPrefix(depositParams.Signature.String(), "0x"))
	if err != nil {
		panic(err)
	}

	params := web3_actions.SendContractTxPayload{
		SmartContractAddr: EphemeralDepositContractAddr,
		ContractABI:       abiFile,
		MethodName:        validatorDepositMethodName,
		SendEtherPayload: web3_actions.SendEtherPayload{
			TransferArgs: web3_actions.TransferArgs{
				Amount:    ValidatorDeposit32EthInGweiUnits,
				ToAddress: common.Address{},
			},
			GasPriceLimits: web3_actions.GasPriceLimits{},
		},

		Params: []interface{}{pubkey, depositParams.WithdrawalCredentials, sig, depositParams.DepositDataRoot},
	}
	signedTx, err := w.GetSignedTxToCallFunctionWithArgs(ctx, &params)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDeposit")
	}
	return signedTx, err
}
