// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package kyberswap_router_solana

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RemoveLiquiditySol is the `removeLiquiditySol` instruction.
type RemoveLiquiditySol struct {
	Liquidity      *ag_binary.Uint128
	AmountTokenMin *ag_binary.Uint128
	AmountSolMin   *ag_binary.Uint128

	// [0] = [] router
	//
	// [1] = [WRITE, SIGNER] user
	//
	// [2] = [WRITE] routerAuthority
	//
	// [3] = [WRITE] userBalanceLiquidity
	//
	// [4] = [WRITE] routerBalanceWsol
	//
	// [5] = [] wrappedSol
	//
	// [6] = [] poolProgram
	//
	// [7] = [] systemProgram
	//
	// [8] = [] rent
	//
	// [9] = [WRITE] poolState
	//
	// [10] = [] poolAuthority
	//
	// [11] = [WRITE] poolBalanceToken0
	//
	// [12] = [WRITE] poolBalanceToken1
	//
	// [13] = [WRITE] poolMintLiquidity
	//
	// [14] = [WRITE] poolBalanceLiquidity
	//
	// [15] = [WRITE] userBalanceToken
	//
	// [16] = [] factoryState
	//
	// [17] = [WRITE] factoryBalanceFeeTo
	//
	// [18] = [] systemTokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRemoveLiquiditySolInstructionBuilder creates a new `RemoveLiquiditySol` instruction builder.
func NewRemoveLiquiditySolInstructionBuilder() *RemoveLiquiditySol {
	nd := &RemoveLiquiditySol{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 19),
	}
	return nd
}

// SetLiquidity sets the "liquidity" parameter.
func (inst *RemoveLiquiditySol) SetLiquidity(liquidity ag_binary.Uint128) *RemoveLiquiditySol {
	inst.Liquidity = &liquidity
	return inst
}

// SetAmountTokenMin sets the "amountTokenMin" parameter.
func (inst *RemoveLiquiditySol) SetAmountTokenMin(amountTokenMin ag_binary.Uint128) *RemoveLiquiditySol {
	inst.AmountTokenMin = &amountTokenMin
	return inst
}

// SetAmountSolMin sets the "amountSolMin" parameter.
func (inst *RemoveLiquiditySol) SetAmountSolMin(amountSolMin ag_binary.Uint128) *RemoveLiquiditySol {
	inst.AmountSolMin = &amountSolMin
	return inst
}

// SetRouterAccount sets the "router" account.
func (inst *RemoveLiquiditySol) SetRouterAccount(router ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(router)
	return inst
}

// GetRouterAccount gets the "router" account.
func (inst *RemoveLiquiditySol) GetRouterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserAccount sets the "user" account.
func (inst *RemoveLiquiditySol) SetUserAccount(user ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *RemoveLiquiditySol) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRouterAuthorityAccount sets the "routerAuthority" account.
func (inst *RemoveLiquiditySol) SetRouterAuthorityAccount(routerAuthority ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(routerAuthority).WRITE()
	return inst
}

// GetRouterAuthorityAccount gets the "routerAuthority" account.
func (inst *RemoveLiquiditySol) GetRouterAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserBalanceLiquidityAccount sets the "userBalanceLiquidity" account.
func (inst *RemoveLiquiditySol) SetUserBalanceLiquidityAccount(userBalanceLiquidity ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userBalanceLiquidity).WRITE()
	return inst
}

// GetUserBalanceLiquidityAccount gets the "userBalanceLiquidity" account.
func (inst *RemoveLiquiditySol) GetUserBalanceLiquidityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRouterBalanceWsolAccount sets the "routerBalanceWsol" account.
func (inst *RemoveLiquiditySol) SetRouterBalanceWsolAccount(routerBalanceWsol ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(routerBalanceWsol).WRITE()
	return inst
}

// GetRouterBalanceWsolAccount gets the "routerBalanceWsol" account.
func (inst *RemoveLiquiditySol) GetRouterBalanceWsolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetWrappedSolAccount sets the "wrappedSol" account.
func (inst *RemoveLiquiditySol) SetWrappedSolAccount(wrappedSol ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(wrappedSol)
	return inst
}

// GetWrappedSolAccount gets the "wrappedSol" account.
func (inst *RemoveLiquiditySol) GetWrappedSolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPoolProgramAccount sets the "poolProgram" account.
func (inst *RemoveLiquiditySol) SetPoolProgramAccount(poolProgram ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(poolProgram)
	return inst
}

// GetPoolProgramAccount gets the "poolProgram" account.
func (inst *RemoveLiquiditySol) GetPoolProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RemoveLiquiditySol) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RemoveLiquiditySol) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentAccount sets the "rent" account.
func (inst *RemoveLiquiditySol) SetRentAccount(rent ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *RemoveLiquiditySol) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetPoolStateAccount sets the "poolState" account.
func (inst *RemoveLiquiditySol) SetPoolStateAccount(poolState ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "poolState" account.
func (inst *RemoveLiquiditySol) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetPoolAuthorityAccount sets the "poolAuthority" account.
func (inst *RemoveLiquiditySol) SetPoolAuthorityAccount(poolAuthority ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(poolAuthority)
	return inst
}

// GetPoolAuthorityAccount gets the "poolAuthority" account.
func (inst *RemoveLiquiditySol) GetPoolAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetPoolBalanceToken0Account sets the "poolBalanceToken0" account.
func (inst *RemoveLiquiditySol) SetPoolBalanceToken0Account(poolBalanceToken0 ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(poolBalanceToken0).WRITE()
	return inst
}

// GetPoolBalanceToken0Account gets the "poolBalanceToken0" account.
func (inst *RemoveLiquiditySol) GetPoolBalanceToken0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetPoolBalanceToken1Account sets the "poolBalanceToken1" account.
func (inst *RemoveLiquiditySol) SetPoolBalanceToken1Account(poolBalanceToken1 ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(poolBalanceToken1).WRITE()
	return inst
}

// GetPoolBalanceToken1Account gets the "poolBalanceToken1" account.
func (inst *RemoveLiquiditySol) GetPoolBalanceToken1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetPoolMintLiquidityAccount sets the "poolMintLiquidity" account.
func (inst *RemoveLiquiditySol) SetPoolMintLiquidityAccount(poolMintLiquidity ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(poolMintLiquidity).WRITE()
	return inst
}

// GetPoolMintLiquidityAccount gets the "poolMintLiquidity" account.
func (inst *RemoveLiquiditySol) GetPoolMintLiquidityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetPoolBalanceLiquidityAccount sets the "poolBalanceLiquidity" account.
func (inst *RemoveLiquiditySol) SetPoolBalanceLiquidityAccount(poolBalanceLiquidity ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(poolBalanceLiquidity).WRITE()
	return inst
}

// GetPoolBalanceLiquidityAccount gets the "poolBalanceLiquidity" account.
func (inst *RemoveLiquiditySol) GetPoolBalanceLiquidityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetUserBalanceTokenAccount sets the "userBalanceToken" account.
func (inst *RemoveLiquiditySol) SetUserBalanceTokenAccount(userBalanceToken ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(userBalanceToken).WRITE()
	return inst
}

// GetUserBalanceTokenAccount gets the "userBalanceToken" account.
func (inst *RemoveLiquiditySol) GetUserBalanceTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetFactoryStateAccount sets the "factoryState" account.
func (inst *RemoveLiquiditySol) SetFactoryStateAccount(factoryState ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(factoryState)
	return inst
}

// GetFactoryStateAccount gets the "factoryState" account.
func (inst *RemoveLiquiditySol) GetFactoryStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetFactoryBalanceFeeToAccount sets the "factoryBalanceFeeTo" account.
func (inst *RemoveLiquiditySol) SetFactoryBalanceFeeToAccount(factoryBalanceFeeTo ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(factoryBalanceFeeTo).WRITE()
	return inst
}

// GetFactoryBalanceFeeToAccount gets the "factoryBalanceFeeTo" account.
func (inst *RemoveLiquiditySol) GetFactoryBalanceFeeToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetSystemTokenProgramAccount sets the "systemTokenProgram" account.
func (inst *RemoveLiquiditySol) SetSystemTokenProgramAccount(systemTokenProgram ag_solanago.PublicKey) *RemoveLiquiditySol {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(systemTokenProgram)
	return inst
}

// GetSystemTokenProgramAccount gets the "systemTokenProgram" account.
func (inst *RemoveLiquiditySol) GetSystemTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

func (inst RemoveLiquiditySol) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RemoveLiquiditySol,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RemoveLiquiditySol) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemoveLiquiditySol) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Liquidity == nil {
			return errors.New("Liquidity parameter is not set")
		}
		if inst.AmountTokenMin == nil {
			return errors.New("AmountTokenMin parameter is not set")
		}
		if inst.AmountSolMin == nil {
			return errors.New("AmountSolMin parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Router is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.RouterAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserBalanceLiquidity is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.RouterBalanceWsol is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.WrappedSol is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PoolProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.PoolState is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.PoolAuthority is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.PoolBalanceToken0 is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.PoolBalanceToken1 is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.PoolMintLiquidity is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.PoolBalanceLiquidity is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.UserBalanceToken is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.FactoryState is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.FactoryBalanceFeeTo is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.SystemTokenProgram is not set")
		}
	}
	return nil
}

func (inst *RemoveLiquiditySol) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveLiquiditySol")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     Liquidity", *inst.Liquidity))
						paramsBranch.Child(ag_format.Param("AmountTokenMin", *inst.AmountTokenMin))
						paramsBranch.Child(ag_format.Param("  AmountSolMin", *inst.AmountSolMin))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=19]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("              router", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                user", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     routerAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("userBalanceLiquidity", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   routerBalanceWsol", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          wrappedSol", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("         poolProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                rent", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           poolState", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       poolAuthority", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("   poolBalanceToken0", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("   poolBalanceToken1", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("   poolMintLiquidity", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("poolBalanceLiquidity", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("    userBalanceToken", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("        factoryState", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta(" factoryBalanceFeeTo", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("  systemTokenProgram", inst.AccountMetaSlice.Get(18)))
					})
				})
		})
}

func (obj RemoveLiquiditySol) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Liquidity` param:
	err = encoder.Encode(obj.Liquidity)
	if err != nil {
		return err
	}
	// Serialize `AmountTokenMin` param:
	err = encoder.Encode(obj.AmountTokenMin)
	if err != nil {
		return err
	}
	// Serialize `AmountSolMin` param:
	err = encoder.Encode(obj.AmountSolMin)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RemoveLiquiditySol) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Liquidity`:
	err = decoder.Decode(&obj.Liquidity)
	if err != nil {
		return err
	}
	// Deserialize `AmountTokenMin`:
	err = decoder.Decode(&obj.AmountTokenMin)
	if err != nil {
		return err
	}
	// Deserialize `AmountSolMin`:
	err = decoder.Decode(&obj.AmountSolMin)
	if err != nil {
		return err
	}
	return nil
}

// NewRemoveLiquiditySolInstruction declares a new RemoveLiquiditySol instruction with the provided parameters and accounts.
func NewRemoveLiquiditySolInstruction(
	// Parameters:
	liquidity ag_binary.Uint128,
	amountTokenMin ag_binary.Uint128,
	amountSolMin ag_binary.Uint128,
	// Accounts:
	router ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	routerAuthority ag_solanago.PublicKey,
	userBalanceLiquidity ag_solanago.PublicKey,
	routerBalanceWsol ag_solanago.PublicKey,
	wrappedSol ag_solanago.PublicKey,
	poolProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	poolAuthority ag_solanago.PublicKey,
	poolBalanceToken0 ag_solanago.PublicKey,
	poolBalanceToken1 ag_solanago.PublicKey,
	poolMintLiquidity ag_solanago.PublicKey,
	poolBalanceLiquidity ag_solanago.PublicKey,
	userBalanceToken ag_solanago.PublicKey,
	factoryState ag_solanago.PublicKey,
	factoryBalanceFeeTo ag_solanago.PublicKey,
	systemTokenProgram ag_solanago.PublicKey) *RemoveLiquiditySol {
	return NewRemoveLiquiditySolInstructionBuilder().
		SetLiquidity(liquidity).
		SetAmountTokenMin(amountTokenMin).
		SetAmountSolMin(amountSolMin).
		SetRouterAccount(router).
		SetUserAccount(user).
		SetRouterAuthorityAccount(routerAuthority).
		SetUserBalanceLiquidityAccount(userBalanceLiquidity).
		SetRouterBalanceWsolAccount(routerBalanceWsol).
		SetWrappedSolAccount(wrappedSol).
		SetPoolProgramAccount(poolProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetPoolStateAccount(poolState).
		SetPoolAuthorityAccount(poolAuthority).
		SetPoolBalanceToken0Account(poolBalanceToken0).
		SetPoolBalanceToken1Account(poolBalanceToken1).
		SetPoolMintLiquidityAccount(poolMintLiquidity).
		SetPoolBalanceLiquidityAccount(poolBalanceLiquidity).
		SetUserBalanceTokenAccount(userBalanceToken).
		SetFactoryStateAccount(factoryState).
		SetFactoryBalanceFeeToAccount(factoryBalanceFeeTo).
		SetSystemTokenProgramAccount(systemTokenProgram)
}
