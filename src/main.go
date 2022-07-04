package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"test_go_client_solana/generated/kyberswap_router_solana"
	"test_go_client_solana/src/util"
)

func main() {

	cfg, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// Create a new RPC client:
	rpcClient := rpc.New(rpc.TestNet_RPC)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.TestNet_WS)
	if err != nil {
		panic(err)
	}

	// Load the account that you will send funds FROM:
	traderPrivateKey, err := solana.PrivateKeyFromSolanaKeygenFile(cfg.WALLET)
	if err != nil {
		panic(err)
	}

	// Hardcode accounts
	traderPublicKey := traderPrivateKey.PublicKey()
	routerProgram := solana.MustPublicKeyFromBase58(cfg.KYBERSWAP_ROUTER_PROGRAM)
	traderAccountKNC := solana.MustPublicKeyFromBase58("FvdmdqaxyBQGtmYthpQ93vsQw8CXHLDteji2P79rf7hW")
	routerStateAddress := solana.MustPublicKeyFromBase58(cfg.ROUTER_STATE_ADDRESS)
	poolProgram := solana.MustPublicKeyFromBase58(cfg.KYBERSWAP_POOL_PROGRAM)
	firstPoolBalanceTokenIn := solana.MustPublicKeyFromBase58("7grnTUEsdKbbd9TMRXHnuxAN914EkQTLijHLmpZtK2dK")
	routerAuthority := solana.MustPublicKeyFromBase58("2HaUCQNdNsu7r9Fd3R9yaHQjfP1QmwDw3taAxRXv3G9H")
	userBalanceTokenIn := solana.MustPublicKeyFromBase58("FvdmdqaxyBQGtmYthpQ93vsQw8CXHLDteji2P79rf7hW")
	routerBalanceWsol := solana.MustPublicKeyFromBase58("DboHptJk7mZX8S6S8EpYnD1jX5M8c5ptyCQYJsBz4c87")
	poolState := solana.MustPublicKeyFromBase58("AnvCMNZzxH6bCeiJUpU8yU4qQCHzC8xBgJQqgznHdDp6")
	poolAuthority := solana.MustPublicKeyFromBase58("DCyLoBVAUHBAPo8M1P72ActQabV6tjCLQF5JdfBu4H1w")
	poolBalanceToken0 := solana.MustPublicKeyFromBase58("7grnTUEsdKbbd9TMRXHnuxAN914EkQTLijHLmpZtK2dK")
	poolBalanceToken1 := solana.MustPublicKeyFromBase58("BHCyAMtj8oGZr9GPqYC1VCw5tyeA2KYthrm1dRPbf8MX")
	kncAmount := uint64(1000000)
	tradeDirection := []byte{0}
	fmt.Println(len(tradeDirection))
	recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	// To decode transaction data
	kyberswap_router_solana.SetProgramID(routerProgram)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			token.NewApproveInstruction(kncAmount, traderAccountKNC, routerAuthority, traderPublicKey, []solana.PublicKey{}).Build(),
			kyberswap_router_solana.NewPrepareForSwapExactTokensForTokensInstruction(
				bin.Uint128{kncAmount, 0, binary.BigEndian},
				bin.Uint128{0, 0, binary.BigEndian},
				tradeDirection,

				traderPublicKey,
				routerStateAddress,
				routerAuthority,
				poolProgram,
				userBalanceTokenIn,
				firstPoolBalanceTokenIn,
				solana.TokenProgramID,
			).Build(),
			kyberswap_router_solana.NewSwapToSolInstruction(
				traderPublicKey,
				routerStateAddress,
				routerAuthority,
				routerBalanceWsol,
				poolProgram,
				solana.WrappedSol,
				solana.SystemProgramID,
				solana.SysVarRentPubkey,
				poolState,
				poolAuthority,
				poolBalanceToken0,
				poolBalanceToken1,
				solana.TokenProgramID,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(traderPrivateKey.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if traderPrivateKey.PublicKey().Equals(key) {
				return &traderPrivateKey
			}
			return nil
		},
	)
	if err != nil {
		panic(fmt.Errorf("unable to sign transaction: %w", err))
	}
	spew.Dump(tx)
	// Pretty print the transaction:
	//tx.EncodeTree(text.NewTreeEncoder(os.Stdout, "Swap exact tokens for sol"))

	// Send transaction, and wait for confirmation:
	sig, err := confirm.SendAndConfirmTransaction(
		context.TODO(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(sig)

	// Or just send the transaction WITHOUT waiting for confirmation:
	// sig, err := rpcClient.SendTransactionWithOpts(
	//   context.TODO(),
	//   tx,
	//   false,
	//   rpc.CommitmentFinalized,
	// )
	// if err != nil {
	//   panic(err)
	// }
	// spew.Dump(sig)
}
