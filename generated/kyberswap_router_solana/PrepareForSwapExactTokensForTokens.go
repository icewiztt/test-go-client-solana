// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package kyberswap_router_solana

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PrepareForSwapExactTokensForTokens is the `prepareForSwapExactTokensForTokens` instruction.
type PrepareForSwapExactTokensForTokens struct {
	AmountIn        *ag_binary.Uint128
	AmountOutMin    *ag_binary.Uint128
	TradeDirections *[]byte

	// [0] = [SIGNER] user
	//
	// [1] = [] router
	//
	// [2] = [] routerAuthority
	//
	// [3] = [] poolProgram
	//
	// [4] = [WRITE] userBalanceTokenIn
	//
	// [5] = [WRITE] firstPoolBalanceTokenIn
	//
	// [6] = [] systemTokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPrepareForSwapExactTokensForTokensInstructionBuilder creates a new `PrepareForSwapExactTokensForTokens` instruction builder.
func NewPrepareForSwapExactTokensForTokensInstructionBuilder() *PrepareForSwapExactTokensForTokens {
	nd := &PrepareForSwapExactTokensForTokens{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	nd.AccountMetaSlice[7] = ag_solanago.Meta(ag_solanago.MustPublicKeyFromBase58("AnvCMNZzxH6bCeiJUpU8yU4qQCHzC8xBgJQqgznHdDp6")).WRITE()
	return nd
}

// SetAmountIn sets the "amountIn" parameter.
func (inst *PrepareForSwapExactTokensForTokens) SetAmountIn(amountIn ag_binary.Uint128) *PrepareForSwapExactTokensForTokens {
	inst.AmountIn = &amountIn
	return inst
}

// SetAmountOutMin sets the "amountOutMin" parameter.
func (inst *PrepareForSwapExactTokensForTokens) SetAmountOutMin(amountOutMin ag_binary.Uint128) *PrepareForSwapExactTokensForTokens {
	inst.AmountOutMin = &amountOutMin
	return inst
}

// SetTradeDirections sets the "tradeDirections" parameter.
func (inst *PrepareForSwapExactTokensForTokens) SetTradeDirections(tradeDirections []byte) *PrepareForSwapExactTokensForTokens {
	inst.TradeDirections = &tradeDirections
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *PrepareForSwapExactTokensForTokens) SetUserAccount(user ag_solanago.PublicKey) *PrepareForSwapExactTokensForTokens {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *PrepareForSwapExactTokensForTokens) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetRouterAccount sets the "router" account.
func (inst *PrepareForSwapExactTokensForTokens) SetRouterAccount(router ag_solanago.PublicKey) *PrepareForSwapExactTokensForTokens {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(router)
	return inst
}

// GetRouterAccount gets the "router" account.
func (inst *PrepareForSwapExactTokensForTokens) GetRouterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRouterAuthorityAccount sets the "routerAuthority" account.
func (inst *PrepareForSwapExactTokensForTokens) SetRouterAuthorityAccount(routerAuthority ag_solanago.PublicKey) *PrepareForSwapExactTokensForTokens {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(routerAuthority)
	return inst
}

// GetRouterAuthorityAccount gets the "routerAuthority" account.
func (inst *PrepareForSwapExactTokensForTokens) GetRouterAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolProgramAccount sets the "poolProgram" account.
func (inst *PrepareForSwapExactTokensForTokens) SetPoolProgramAccount(poolProgram ag_solanago.PublicKey) *PrepareForSwapExactTokensForTokens {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolProgram)
	return inst
}

// GetPoolProgramAccount gets the "poolProgram" account.
func (inst *PrepareForSwapExactTokensForTokens) GetPoolProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserBalanceTokenInAccount sets the "userBalanceTokenIn" account.
func (inst *PrepareForSwapExactTokensForTokens) SetUserBalanceTokenInAccount(userBalanceTokenIn ag_solanago.PublicKey) *PrepareForSwapExactTokensForTokens {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userBalanceTokenIn).WRITE()
	return inst
}

// GetUserBalanceTokenInAccount gets the "userBalanceTokenIn" account.
func (inst *PrepareForSwapExactTokensForTokens) GetUserBalanceTokenInAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetFirstPoolBalanceTokenInAccount sets the "firstPoolBalanceTokenIn" account.
func (inst *PrepareForSwapExactTokensForTokens) SetFirstPoolBalanceTokenInAccount(firstPoolBalanceTokenIn ag_solanago.PublicKey) *PrepareForSwapExactTokensForTokens {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(firstPoolBalanceTokenIn).WRITE()
	return inst
}

// GetFirstPoolBalanceTokenInAccount gets the "firstPoolBalanceTokenIn" account.
func (inst *PrepareForSwapExactTokensForTokens) GetFirstPoolBalanceTokenInAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemTokenProgramAccount sets the "systemTokenProgram" account.
func (inst *PrepareForSwapExactTokensForTokens) SetSystemTokenProgramAccount(systemTokenProgram ag_solanago.PublicKey) *PrepareForSwapExactTokensForTokens {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemTokenProgram)
	return inst
}

// GetSystemTokenProgramAccount gets the "systemTokenProgram" account.
func (inst *PrepareForSwapExactTokensForTokens) GetSystemTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst PrepareForSwapExactTokensForTokens) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PrepareForSwapExactTokensForTokens,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PrepareForSwapExactTokensForTokens) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PrepareForSwapExactTokensForTokens) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AmountIn == nil {
			return errors.New("AmountIn parameter is not set")
		}
		if inst.AmountOutMin == nil {
			return errors.New("AmountOutMin parameter is not set")
		}
		if inst.TradeDirections == nil {
			return errors.New("TradeDirections parameter is not set")
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
			return errors.New("accounts.RouterAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserBalanceTokenIn is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FirstPoolBalanceTokenIn is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemTokenProgram is not set")
		}
	}
	return nil
}

func (inst *PrepareForSwapExactTokensForTokens) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PrepareForSwapExactTokensForTokens")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("       AmountIn", *inst.AmountIn))
						paramsBranch.Child(ag_format.Param("   AmountOutMin", *inst.AmountOutMin))
						paramsBranch.Child(ag_format.Param("TradeDirections", *inst.TradeDirections))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 router", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        routerAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            poolProgram", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     userBalanceTokenIn", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("firstPoolBalanceTokenIn", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     systemTokenProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj PrepareForSwapExactTokensForTokens) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AmountIn` param:
	err = encoder.Encode(obj.AmountIn)
	if err != nil {
		return err
	}
	// Serialize `AmountOutMin` param:
	err = encoder.Encode(obj.AmountOutMin)
	if err != nil {
		return err
	}
	// Serialize `TradeDirections` param:
	err = encoder.Encode(obj.TradeDirections)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PrepareForSwapExactTokensForTokens) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AmountIn`:
	err = decoder.Decode(&obj.AmountIn)
	if err != nil {
		return err
	}
	// Deserialize `AmountOutMin`:
	err = decoder.Decode(&obj.AmountOutMin)
	if err != nil {
		return err
	}
	// Deserialize `TradeDirections`:
	err = decoder.Decode(&obj.TradeDirections)
	if err != nil {
		return err
	}
	return nil
}

// NewPrepareForSwapExactTokensForTokensInstruction declares a new PrepareForSwapExactTokensForTokens instruction with the provided parameters and accounts.
func NewPrepareForSwapExactTokensForTokensInstruction(
	// Parameters:
	amountIn ag_binary.Uint128,
	amountOutMin ag_binary.Uint128,
	tradeDirections []byte,
	// Accounts:
	user ag_solanago.PublicKey,
	router ag_solanago.PublicKey,
	routerAuthority ag_solanago.PublicKey,
	poolProgram ag_solanago.PublicKey,
	userBalanceTokenIn ag_solanago.PublicKey,
	firstPoolBalanceTokenIn ag_solanago.PublicKey,
	systemTokenProgram ag_solanago.PublicKey) *PrepareForSwapExactTokensForTokens {
	return NewPrepareForSwapExactTokensForTokensInstructionBuilder().
		SetAmountIn(amountIn).
		SetAmountOutMin(amountOutMin).
		SetTradeDirections(tradeDirections).
		SetUserAccount(user).
		SetRouterAccount(router).
		SetRouterAuthorityAccount(routerAuthority).
		SetPoolProgramAccount(poolProgram).
		SetUserBalanceTokenInAccount(userBalanceTokenIn).
		SetFirstPoolBalanceTokenInAccount(firstPoolBalanceTokenIn).
		SetSystemTokenProgramAccount(systemTokenProgram)
}
