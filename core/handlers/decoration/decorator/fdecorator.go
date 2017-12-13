package decorator

import (
	"fmt"

	"github.com/hyperledger/fabric/core/config"
	"github.com/hyperledger/fabric/core/handlers/decoration"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/spf13/viper"
)

// NewDecorator creates a new decorator
func NewFDecorator() decoration.Decorator {
	// Load decoration.yaml

	v := &viper.Viper{}
	err := config.InitViper(v, "decoration")
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error when initializing %s config : %s\n", "decoration", err))
	}
	keys := v.AllKeys()
	values := make(map[string][]byte)
	for _, k := range keys {
		values[k] = []byte(v.GetString(k))
	}

	return &fDecorator{values: values}
}

type fDecorator struct {
	values map[string][]byte
}

// Decorate decorates a chaincode input by changing it
func (d *fDecorator) Decorate(proposal *peer.Proposal, input *peer.ChaincodeInput) *peer.ChaincodeInput {
	for k, v := range d.values {
		input.Decorations[k] = v
	}

	return input
}
