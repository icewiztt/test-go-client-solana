// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package kyberswap_router_solana

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Router struct {
	FactoryAddress ag_solanago.PublicKey
}

var RouterDiscriminator = [8]byte{94, 226, 217, 169, 186, 4, 198, 7}

func (obj Router) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(RouterDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `FactoryAddress` param:
	err = encoder.Encode(obj.FactoryAddress)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Router) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(RouterDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[94 226 217 169 186 4 198 7]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `FactoryAddress`:
	err = decoder.Decode(&obj.FactoryAddress)
	if err != nil {
		return err
	}
	return nil
}