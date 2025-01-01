package prism

import (
	"testing"
	"math/big"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func setup(t *testing.T, _initVal *big.Int) (*backends.SimulatedBackend, *BaseCounterSession) {
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

	// deploy BaseCounter contract
	_, _, counter, err := DeployBaseCounter(auth, sim, _initVal)
	assert.NoError(t, err)

	sim.Commit()

	session := createSession(counter, auth)

	return sim, session
}

func teardown(t *testing.T, _sim *backends.SimulatedBackend) {
	_sim.Close()
}

func createSession(_counter *BaseCounter, _auth *bind.TransactOpts) *BaseCounterSession {
	session := BaseCounterSession{
		Contract: _counter,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *_auth,
	}
	return &session
}

func TestIncrementAndDecrement(t *testing.T) {
	// Variables
	testVal := big.NewInt(100)

	// Setup
	t.Log("Setup")
	sim, session := setup(t, testVal)

	// Main
	t.Log("Start Test")
	{
		t.Log("Check Init Value")
		count, err := session.Count()
		assert.NoError(t, err)
		assert.Equal(t, count.Cmp(testVal), 0)
		t.Logf("count: %+v\n", count)
	}

	{
		t.Log("Test Increment")
		// Exec increment
		_, err := session.Increment()
		assert.NoError(t, err)
		sim.Commit()
		// Test incremented value
		count, err := session.Count()
		expected := new(big.Int).Add(testVal, big.NewInt(1))
		assert.Equal(t, count.Cmp(expected), 0)
		t.Logf("count: %+v, expected: %+v\n", count, expected)
	}

	{
		t.Log("Test Decrement")
		// Exec increment
		_, err := session.Decrement()
		assert.NoError(t, err)
		sim.Commit()
		// Test decremented value
		count, err := session.Count()
		assert.Equal(t, count.Cmp(testVal), 0)
		t.Logf("count: %+v, expected: %+v\n", count, testVal)
	}

	// Teardown
	t.Log("Teardown")
	teardown(t, sim)
}

func TestOverflow(t *testing.T) {
	// Variables
	testVal := big.NewInt(2)
	testVal = testVal.Exp(testVal, big.NewInt(256), nil).Sub(testVal, big.NewInt(1))

	// Setup
	t.Log("Setup")
	sim, session := setup(t, testVal)

	// Main
	t.Log("Start Test")
	{
		t.Log("Test Overflow")
		// Exec increment
		_, err := session.Increment()
		assert.NoError(t, err)
		sim.Commit()
		// Test incremented value
		count, err := session.Count()
		assert.Equal(t, count.Cmp(testVal), 0)
		t.Logf("count: %+v, expected: %+v\n", count, testVal)
	}

	// Teardown
	t.Log("Teardown")
	teardown(t, sim)
}

func TestUnderflow(t *testing.T) {
	// Variables
	testVal := big.NewInt(0)

	// Setup
	t.Log("Setup")
	sim, session := setup(t, testVal)

	// Main
	t.Log("Start Test")
	{
		t.Log("Test Underflow")
		// Exec decrement
		_, err := session.Decrement()
		assert.NoError(t, err)
		sim.Commit()
		// Test decremented value
		count, err := session.Count()
		assert.Equal(t, count.Cmp(testVal), 0)
		t.Logf("count: %+v, expected: %+v\n", count, testVal)
	}

	// Teardown
	t.Log("Teardown")
	teardown(t, sim)
}
