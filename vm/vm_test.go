package vm_test

import (
	"testing"
	"github.com/annchain/OG/common"
	"github.com/annchain/OG/vm/ovm"
	"github.com/annchain/OG/types"
	vmtypes "github.com/annchain/OG/vm/types"
	"github.com/annchain/OG/common/math"
	"math/big"
	"github.com/annchain/OG/vm/eth/core/vm"
	"fmt"
	"io/ioutil"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/annchain/OG/vm/eth/common/hexutil"
	"os"
)

func readFile(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	bytes, err = hex.DecodeString(string(bytes))
	return bytes
}

func TestSmallContract(t *testing.T) {
	txContext := &ovm.TxContext{
		From: types.HexToAddress("0x01"),
		//To:       types.HexToAddress("0x02"),
		Value:    math.NewBigInt(0),
		Data:     readFile("./compiler/d/OwnedToken.bin"),
		GasPrice: math.NewBigInt(1),
		GasLimit: 600000,
	}
	coinBase := types.HexToAddress("0x1234567812345678AABBCCDDEEFF998877665544")

	mmdb := ovm.NewMemoryStateDB()
	ldb := ovm.NewLayerDB(mmdb)
	ldb.NewLayer()
	ldb.CreateAccount(txContext.From)
	ldb.AddBalance(txContext.From, big.NewInt(10000000))
	ldb.CreateAccount(coinBase)
	ldb.AddBalance(coinBase, big.NewInt(10000000))
	logrus.Info("Init accounts done")

	context := ovm.NewEVMContext(txContext, &ovm.DefaultChainContext{}, &coinBase, ldb)

	tracer := vm.NewStructLogger(&vm.LogConfig{
		Debug: true,
	})

	evmInterpreter := vm.NewEVMInterpreter(&context, &vm.InterpreterConfig{
		Debug:  true,
		Tracer: tracer,
	})

	ovm := ovm.NewOVM(context, []ovm.Interpreter{evmInterpreter}, &ovm.OVMConfig{NoRecursion: false})

	logrus.Info("Deploying contract")
	ret, contractAddr, leftOverGas, err := ovm.Create(&context, vmtypes.AccountRef(coinBase), txContext.Data, txContext.GasLimit, txContext.Value.Value)
	// make duplicate
	//ovm.StateDB.SetNonce(coinBase, 0)
	//ret, contractAddr, leftOverGas, err = ovm.Create(&context, vmtypes.AccountRef(coinBase), txContext.Data, txContext.GasLimit, txContext.Value.Value)
	logrus.Info("Deployed contract")
	fmt.Println("CP1", common.Bytes2Hex(ret), contractAddr.String(), leftOverGas, err)
	//fmt.Println(ldb.String())
	//vm.WriteTrace(os.Stdout, tracer.Logs)
	assert.NoError(t, err)

	txContext.Value = math.NewBigInt(0)

	logrus.Info("Calling contract")

	var name [32]byte
	copy(name[:], "abcdefghijklmnopqrstuvwxyz")

	var input []byte
	contractAddress, err := hexutil.Decode("0x898855ed")
	assert.NoError(t, err)
	input = append(input, contractAddress...)
	input = append(input, name[:]...)

	ret, leftOverGas, err = ovm.Call(&context, vmtypes.AccountRef(coinBase), contractAddr, input, txContext.GasLimit, txContext.Value.Value)
	logrus.Info("Called contract")
	fmt.Println("CP2", common.Bytes2Hex(ret), contractAddr.String(), leftOverGas, err)
	fmt.Println(ldb.String())
	vm.WriteTrace(os.Stdout, tracer.Logs)
	assert.NoError(t, err)
}
