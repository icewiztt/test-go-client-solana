// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package kyberswap_router_solana

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AddLiquidity is the `addLiquidity` instruction.
type AddLiquidity struct {
	Amount0Desired *ag_binary.Uint128
	Amount1Desired *ag_binary.Uint128
	Amount0Min     *ag_binary.Uint128
	Amount1Min     *ag_binary.Uint128

	// [0] = [SIGNER] user
	//
	// [1] = [] router
	//
	// [2] = [] poolAuthority
	//
	// [3] = [WRITE] poolState
	//
	// [4] = [WRITE] poolBalanceToken0
	//
	// [5] = [WRITE] poolBalanceToken1
	//
	// [6] = [WRITE] poolMintLiquidity
	//
	// [7] = [WRITE] poolBalanceLockedLiquidity
	//
	// [8] = [WRITE] userBalanceLiquidity
	//
	// [9] = [] routerAuthority
	//
	// [10] = [WRITE] userBalanceToken0
	//
	// [11] = [WRITE] userBalanceToken1
	//
	// [12] = [] factoryState
	//
	// [13] = [WRITE] factoryBalanceFeeTo
	//
	// [14] = [] poolProgram
	//
	// [15] = [] systemTokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddLiquidityInstructionBuilder creates a new `AddLiquidity` instruction builder.
func NewAddLiquidityInstructionBuilder() *AddLiquidity {
	nd := &AddLiquidity{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	return nd
}

// SetAmount0Desired sets the "amount0Desired" parameter.
func (inst *AddLiquidity) SetAmount0Desired(amount0Desired ag_binary.Uint128) *AddLiquidity {
	inst.Amount0Desired = &amount0Desired
	return inst
}

// SetAmount1Desired sets the "amount1Desired" parameter.
func (inst *AddLiquidity) SetAmount1Desired(amount1Desired ag_binary.Uint128) *AddLiquidity {
	inst.Amount1Desired = &amount1Desired
	return inst
}

// SetAmount0Min sets the "amount0Min" parameter.
func (inst *AddLiquidity) SetAmount0Min(amount0Min ag_binary.Uint128) *AddLiquidity {
	inst.Amount0Min = &amount0Min
	return inst
}

// SetAmount1Min sets the "amount1Min" parameter.
func (inst *AddLiquidity) SetAmount1Min(amount1Min ag_binary.Uint128) *AddLiquidity {
	inst.Amount1Min = &amount1Min
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *AddLiquidity) SetUserAccount(user ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *AddLiquidity) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetRouterAccount sets the "router" account.
func (inst *AddLiquidity) SetRouterAccount(router ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(router)
	return inst
}

// GetRouterAccount gets the "router" account.
func (inst *AddLiquidity) GetRouterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPoolAuthorityAccount sets the "poolAuthority" account.
func (inst *AddLiquidity) SetPoolAuthorityAccount(poolAuthority ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(poolAuthority)
	return inst
}

// GetPoolAuthorityAccount gets the "poolAuthority" account.
func (inst *AddLiquidity) GetPoolAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolStateAccount sets the "poolState" account.
func (inst *AddLiquidity) SetPoolStateAccount(poolState ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "poolState" account.
func (inst *AddLiquidity) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPoolBalanceToken0Account sets the "poolBalanceToken0" account.
func (inst *AddLiquidity) SetPoolBalanceToken0Account(poolBalanceToken0 ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(poolBalanceToken0).WRITE()
	return inst
}

// GetPoolBalanceToken0Account gets the "poolBalanceToken0" account.
func (inst *AddLiquidity) GetPoolBalanceToken0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPoolBalanceToken1Account sets the "poolBalanceToken1" account.
func (inst *AddLiquidity) SetPoolBalanceToken1Account(poolBalanceToken1 ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(poolBalanceToken1).WRITE()
	return inst
}

// GetPoolBalanceToken1Account gets the "poolBalanceToken1" account.
func (inst *AddLiquidity) GetPoolBalanceToken1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPoolMintLiquidityAccount sets the "poolMintLiquidity" account.
func (inst *AddLiquidity) SetPoolMintLiquidityAccount(poolMintLiquidity ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(poolMintLiquidity).WRITE()
	return inst
}

// GetPoolMintLiquidityAccount gets the "poolMintLiquidity" account.
func (inst *AddLiquidity) GetPoolMintLiquidityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPoolBalanceLockedLiquidityAccount sets the "poolBalanceLockedLiquidity" account.
func (inst *AddLiquidity) SetPoolBalanceLockedLiquidityAccount(poolBalanceLockedLiquidity ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(poolBalanceLockedLiquidity).WRITE()
	return inst
}

// GetPoolBalanceLockedLiquidityAccount gets the "poolBalanceLockedLiquidity" account.
func (inst *AddLiquidity) GetPoolBalanceLockedLiquidityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetUserBalanceLiquidityAccount sets the "userBalanceLiquidity" account.
func (inst *AddLiquidity) SetUserBalanceLiquidityAccount(userBalanceLiquidity ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(userBalanceLiquidity).WRITE()
	return inst
}

// GetUserBalanceLiquidityAccount gets the "userBalanceLiquidity" account.
func (inst *AddLiquidity) GetUserBalanceLiquidityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRouterAuthorityAccount sets the "routerAuthority" account.
func (inst *AddLiquidity) SetRouterAuthorityAccount(routerAuthority ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(routerAuthority)
	return inst
}

// GetRouterAuthorityAccount gets the "routerAuthority" account.
func (inst *AddLiquidity) GetRouterAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetUserBalanceToken0Account sets the "userBalanceToken0" account.
func (inst *AddLiquidity) SetUserBalanceToken0Account(userBalanceToken0 ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(userBalanceToken0).WRITE()
	return inst
}

// GetUserBalanceToken0Account gets the "userBalanceToken0" account.
func (inst *AddLiquidity) GetUserBalanceToken0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetUserBalanceToken1Account sets the "userBalanceToken1" account.
func (inst *AddLiquidity) SetUserBalanceToken1Account(userBalanceToken1 ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(userBalanceToken1).WRITE()
	return inst
}

// GetUserBalanceToken1Account gets the "userBalanceToken1" account.
func (inst *AddLiquidity) GetUserBalanceToken1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetFactoryStateAccount sets the "factoryState" account.
func (inst *AddLiquidity) SetFactoryStateAccount(factoryState ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(factoryState)
	return inst
}

// GetFactoryStateAccount gets the "factoryState" account.
func (inst *AddLiquidity) GetFactoryStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetFactoryBalanceFeeToAccount sets the "factoryBalanceFeeTo" account.
func (inst *AddLiquidity) SetFactoryBalanceFeeToAccount(factoryBalanceFeeTo ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(factoryBalanceFeeTo).WRITE()
	return inst
}

// GetFactoryBalanceFeeToAccount gets the "factoryBalanceFeeTo" account.
func (inst *AddLiquidity) GetFactoryBalanceFeeToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetPoolProgramAccount sets the "poolProgram" account.
func (inst *AddLiquidity) SetPoolProgramAccount(poolProgram ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(poolProgram)
	return inst
}

// GetPoolProgramAccount gets the "poolProgram" account.
func (inst *AddLiquidity) GetPoolProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSystemTokenProgramAccount sets the "systemTokenProgram" account.
func (inst *AddLiquidity) SetSystemTokenProgramAccount(systemTokenProgram ag_solanago.PublicKey) *AddLiquidity {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(systemTokenProgram)
	return inst
}

// GetSystemTokenProgramAccount gets the "systemTokenProgram" account.
func (inst *AddLiquidity) GetSystemTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst AddLiquidity) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddLiquidity,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddLiquidity) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddLiquidity) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount0Desired == nil {
			return errors.New("Amount0Desired parameter is not set")
		}
		if inst.Amount1Desired == nil {
			return errors.New("Amount1Desired parameter is not set")
		}
		if inst.Amount0Min == nil {
			return errors.New("Amount0Min parameter is not set")
		}
		if inst.Amount1Min == nil {
			return errors.New("Amount1Min parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Router is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PoolAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolState is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PoolBalanceToken0 is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PoolBalanceToken1 is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PoolMintLiquidity is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.PoolBalanceLockedLiquidity is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.UserBalanceLiquidity is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.RouterAuthority is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.UserBalanceToken0 is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.UserBalanceToken1 is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.FactoryState is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.FactoryBalanceFeeTo is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.PoolProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SystemTokenProgram is not set")
		}
	}
	return nil
}

func (inst *AddLiquidity) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddLiquidity")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount0Desired", *inst.Amount0Desired))
						paramsBranch.Child(ag_format.Param("Amount1Desired", *inst.Amount1Desired))
						paramsBranch.Child(ag_format.Param("    Amount0Min", *inst.Amount0Min))
						paramsBranch.Child(ag_format.Param("    Amount1Min", *inst.Amount1Min))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                      user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                    router", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             poolAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                 poolState", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         poolBalanceToken0", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         poolBalanceToken1", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("         poolMintLiquidity", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("poolBalanceLockedLiquidity", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      userBalanceLiquidity", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           routerAuthority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("         userBalanceToken0", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("         userBalanceToken1", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("              factoryState", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("       factoryBalanceFeeTo", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("               poolProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("        systemTokenProgram", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj AddLiquidity) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount0Desired` param:
	err = encoder.Encode(obj.Amount0Desired)
	if err != nil {
		return err
	}
	// Serialize `Amount1Desired` param:
	err = encoder.Encode(obj.Amount1Desired)
	if err != nil {
		return err
	}
	// Serialize `Amount0Min` param:
	err = encoder.Encode(obj.Amount0Min)
	if err != nil {
		return err
	}
	// Serialize `Amount1Min` param:
	err = encoder.Encode(obj.Amount1Min)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddLiquidity) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount0Desired`:
	err = decoder.Decode(&obj.Amount0Desired)
	if err != nil {
		return err
	}
	// Deserialize `Amount1Desired`:
	err = decoder.Decode(&obj.Amount1Desired)
	if err != nil {
		return err
	}
	// Deserialize `Amount0Min`:
	err = decoder.Decode(&obj.Amount0Min)
	if err != nil {
		return err
	}
	// Deserialize `Amount1Min`:
	err = decoder.Decode(&obj.Amount1Min)
	if err != nil {
		return err
	}
	return nil
}

// NewAddLiquidityInstruction declares a new AddLiquidity instruction with the provided parameters and accounts.
func NewAddLiquidityInstruction(
	// Parameters:
	amount0Desired ag_binary.Uint128,
	amount1Desired ag_binary.Uint128,
	amount0Min ag_binary.Uint128,
	amount1Min ag_binary.Uint128,
	// Accounts:
	user ag_solanago.PublicKey,
	router ag_solanago.PublicKey,
	poolAuthority ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	poolBalanceToken0 ag_solanago.PublicKey,
	poolBalanceToken1 ag_solanago.PublicKey,
	poolMintLiquidity ag_solanago.PublicKey,
	poolBalanceLockedLiquidity ag_solanago.PublicKey,
	userBalanceLiquidity ag_solanago.PublicKey,
	routerAuthority ag_solanago.PublicKey,
	userBalanceToken0 ag_solanago.PublicKey,
	userBalanceToken1 ag_solanago.PublicKey,
	factoryState ag_solanago.PublicKey,
	factoryBalanceFeeTo ag_solanago.PublicKey,
	poolProgram ag_solanago.PublicKey,
	systemTokenProgram ag_solanago.PublicKey) *AddLiquidity {
	return NewAddLiquidityInstructionBuilder().
		SetAmount0Desired(amount0Desired).
		SetAmount1Desired(amount1Desired).
		SetAmount0Min(amount0Min).
		SetAmount1Min(amount1Min).
		SetUserAccount(user).
		SetRouterAccount(router).
		SetPoolAuthorityAccount(poolAuthority).
		SetPoolStateAccount(poolState).
		SetPoolBalanceToken0Account(poolBalanceToken0).
		SetPoolBalanceToken1Account(poolBalanceToken1).
		SetPoolMintLiquidityAccount(poolMintLiquidity).
		SetPoolBalanceLockedLiquidityAccount(poolBalanceLockedLiquidity).
		SetUserBalanceLiquidityAccount(userBalanceLiquidity).
		SetRouterAuthorityAccount(routerAuthority).
		SetUserBalanceToken0Account(userBalanceToken0).
		SetUserBalanceToken1Account(userBalanceToken1).
		SetFactoryStateAccount(factoryState).
		SetFactoryBalanceFeeToAccount(factoryBalanceFeeTo).
		SetPoolProgramAccount(poolProgram).
		SetSystemTokenProgramAccount(systemTokenProgram)
}