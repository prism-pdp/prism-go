package sol

import (
	"fmt"
	"testing"
	"math/big"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func setup(t *testing.T) (*backends.SimulatedBackend, *BaseCounter) {
	t.Helper()

	// generate key
	key, err := crypto.GenerateKey()
	assert.NoError(t, err)

	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	assert.NoError(t, err)

	alloc := core.GenesisAlloc{
		auth.From: {
			Balance: big.NewInt(5000000000000000000),
		},
	}
	sim := backends.NewSimulatedBackend(alloc, 10000000)
	// defer sim.Close()

	// deploy BaseCounter contract
	_, _, counter, err := DeployBaseCounter(auth, sim, big.NewInt(100))
	assert.NoError(t, err)

	sim.Commit()

	return sim, counter
}


func TestBaseCounter(t *testing.T) {
	_, counter := setup(t)

	callOpts := bind.CallOpts{
		Pending: true,
	}
	count, err := counter.Count(&callOpts)
	assert.NoError(t, err)

	fmt.Println(count)
}