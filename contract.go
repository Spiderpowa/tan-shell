// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tanshell

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TanshellABI is the input ABI used to generate the binding from.
const TanshellABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"clientIndex\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"}],\"name\":\"ClientAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"clientIndex\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"msgId\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"stream\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"eof\",\"type\":\"bool\"}],\"name\":\"Stdout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"clientIndex\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"msgId\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"stream\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"eof\",\"type\":\"bool\"}],\"name\":\"Stderr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"clientIndex\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"msgId\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"stream\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"eof\",\"type\":\"bool\"}],\"name\":\"Stdin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"client\",\"type\":\"address\"}],\"name\":\"addClient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"msgId\",\"type\":\"int256\"},{\"name\":\"stream\",\"type\":\"bytes\"},{\"name\":\"eof\",\"type\":\"bool\"}],\"name\":\"stdout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"msgId\",\"type\":\"int256\"},{\"name\":\"stream\",\"type\":\"bytes\"},{\"name\":\"eof\",\"type\":\"bool\"}],\"name\":\"stderr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"clientIndex\",\"type\":\"int256\"},{\"name\":\"msgId\",\"type\":\"int256\"},{\"name\":\"stream\",\"type\":\"bytes\"},{\"name\":\"eof\",\"type\":\"bool\"}],\"name\":\"stdin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Tanshell is an auto generated Go binding around an Ethereum contract.
type Tanshell struct {
	TanshellCaller     // Read-only binding to the contract
	TanshellTransactor // Write-only binding to the contract
	TanshellFilterer   // Log filterer for contract events
}

// TanshellCaller is an auto generated read-only Go binding around an Ethereum contract.
type TanshellCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TanshellTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TanshellTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TanshellFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TanshellFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TanshellSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TanshellSession struct {
	Contract     *Tanshell         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TanshellCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TanshellCallerSession struct {
	Contract *TanshellCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TanshellTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TanshellTransactorSession struct {
	Contract     *TanshellTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TanshellRaw is an auto generated low-level Go binding around an Ethereum contract.
type TanshellRaw struct {
	Contract *Tanshell // Generic contract binding to access the raw methods on
}

// TanshellCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TanshellCallerRaw struct {
	Contract *TanshellCaller // Generic read-only contract binding to access the raw methods on
}

// TanshellTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TanshellTransactorRaw struct {
	Contract *TanshellTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTanshell creates a new instance of Tanshell, bound to a specific deployed contract.
func NewTanshell(address common.Address, backend bind.ContractBackend) (*Tanshell, error) {
	contract, err := bindTanshell(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tanshell{TanshellCaller: TanshellCaller{contract: contract}, TanshellTransactor: TanshellTransactor{contract: contract}, TanshellFilterer: TanshellFilterer{contract: contract}}, nil
}

// NewTanshellCaller creates a new read-only instance of Tanshell, bound to a specific deployed contract.
func NewTanshellCaller(address common.Address, caller bind.ContractCaller) (*TanshellCaller, error) {
	contract, err := bindTanshell(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TanshellCaller{contract: contract}, nil
}

// NewTanshellTransactor creates a new write-only instance of Tanshell, bound to a specific deployed contract.
func NewTanshellTransactor(address common.Address, transactor bind.ContractTransactor) (*TanshellTransactor, error) {
	contract, err := bindTanshell(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TanshellTransactor{contract: contract}, nil
}

// NewTanshellFilterer creates a new log filterer instance of Tanshell, bound to a specific deployed contract.
func NewTanshellFilterer(address common.Address, filterer bind.ContractFilterer) (*TanshellFilterer, error) {
	contract, err := bindTanshell(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TanshellFilterer{contract: contract}, nil
}

// bindTanshell binds a generic wrapper to an already deployed contract.
func bindTanshell(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TanshellABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tanshell *TanshellRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tanshell.Contract.TanshellCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tanshell *TanshellRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tanshell.Contract.TanshellTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tanshell *TanshellRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tanshell.Contract.TanshellTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tanshell *TanshellCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tanshell.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tanshell *TanshellTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tanshell.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tanshell *TanshellTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tanshell.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Tanshell *TanshellCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tanshell.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Tanshell *TanshellSession) IsOwner() (bool, error) {
	return _Tanshell.Contract.IsOwner(&_Tanshell.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Tanshell *TanshellCallerSession) IsOwner() (bool, error) {
	return _Tanshell.Contract.IsOwner(&_Tanshell.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tanshell *TanshellCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tanshell.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tanshell *TanshellSession) Owner() (common.Address, error) {
	return _Tanshell.Contract.Owner(&_Tanshell.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tanshell *TanshellCallerSession) Owner() (common.Address, error) {
	return _Tanshell.Contract.Owner(&_Tanshell.CallOpts)
}

// AddClient is a paid mutator transaction binding the contract method 0x43928cfd.
//
// Solidity: function addClient(address client) returns()
func (_Tanshell *TanshellTransactor) AddClient(opts *bind.TransactOpts, client common.Address) (*types.Transaction, error) {
	return _Tanshell.contract.Transact(opts, "addClient", client)
}

// AddClient is a paid mutator transaction binding the contract method 0x43928cfd.
//
// Solidity: function addClient(address client) returns()
func (_Tanshell *TanshellSession) AddClient(client common.Address) (*types.Transaction, error) {
	return _Tanshell.Contract.AddClient(&_Tanshell.TransactOpts, client)
}

// AddClient is a paid mutator transaction binding the contract method 0x43928cfd.
//
// Solidity: function addClient(address client) returns()
func (_Tanshell *TanshellTransactorSession) AddClient(client common.Address) (*types.Transaction, error) {
	return _Tanshell.Contract.AddClient(&_Tanshell.TransactOpts, client)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tanshell *TanshellTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tanshell.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tanshell *TanshellSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tanshell.Contract.RenounceOwnership(&_Tanshell.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tanshell *TanshellTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tanshell.Contract.RenounceOwnership(&_Tanshell.TransactOpts)
}

// Stderr is a paid mutator transaction binding the contract method 0xd08d9612.
//
// Solidity: function stderr(int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellTransactor) Stderr(opts *bind.TransactOpts, msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.contract.Transact(opts, "stderr", msgId, stream, eof)
}

// Stderr is a paid mutator transaction binding the contract method 0xd08d9612.
//
// Solidity: function stderr(int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellSession) Stderr(msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.Contract.Stderr(&_Tanshell.TransactOpts, msgId, stream, eof)
}

// Stderr is a paid mutator transaction binding the contract method 0xd08d9612.
//
// Solidity: function stderr(int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellTransactorSession) Stderr(msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.Contract.Stderr(&_Tanshell.TransactOpts, msgId, stream, eof)
}

// Stdin is a paid mutator transaction binding the contract method 0xc6747d28.
//
// Solidity: function stdin(int256 clientIndex, int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellTransactor) Stdin(opts *bind.TransactOpts, clientIndex *big.Int, msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.contract.Transact(opts, "stdin", clientIndex, msgId, stream, eof)
}

// Stdin is a paid mutator transaction binding the contract method 0xc6747d28.
//
// Solidity: function stdin(int256 clientIndex, int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellSession) Stdin(clientIndex *big.Int, msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.Contract.Stdin(&_Tanshell.TransactOpts, clientIndex, msgId, stream, eof)
}

// Stdin is a paid mutator transaction binding the contract method 0xc6747d28.
//
// Solidity: function stdin(int256 clientIndex, int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellTransactorSession) Stdin(clientIndex *big.Int, msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.Contract.Stdin(&_Tanshell.TransactOpts, clientIndex, msgId, stream, eof)
}

// Stdout is a paid mutator transaction binding the contract method 0xe6c02d92.
//
// Solidity: function stdout(int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellTransactor) Stdout(opts *bind.TransactOpts, msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.contract.Transact(opts, "stdout", msgId, stream, eof)
}

// Stdout is a paid mutator transaction binding the contract method 0xe6c02d92.
//
// Solidity: function stdout(int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellSession) Stdout(msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.Contract.Stdout(&_Tanshell.TransactOpts, msgId, stream, eof)
}

// Stdout is a paid mutator transaction binding the contract method 0xe6c02d92.
//
// Solidity: function stdout(int256 msgId, bytes stream, bool eof) returns()
func (_Tanshell *TanshellTransactorSession) Stdout(msgId *big.Int, stream []byte, eof bool) (*types.Transaction, error) {
	return _Tanshell.Contract.Stdout(&_Tanshell.TransactOpts, msgId, stream, eof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tanshell *TanshellTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tanshell.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tanshell *TanshellSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tanshell.Contract.TransferOwnership(&_Tanshell.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tanshell *TanshellTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tanshell.Contract.TransferOwnership(&_Tanshell.TransactOpts, newOwner)
}

// TanshellClientAddedIterator is returned from FilterClientAdded and is used to iterate over the raw logs and unpacked data for ClientAdded events raised by the Tanshell contract.
type TanshellClientAddedIterator struct {
	Event *TanshellClientAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TanshellClientAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TanshellClientAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TanshellClientAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TanshellClientAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TanshellClientAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TanshellClientAdded represents a ClientAdded event raised by the Tanshell contract.
type TanshellClientAdded struct {
	ClientIndex *big.Int
	Client      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClientAdded is a free log retrieval operation binding the contract event 0xd919eebb70f12e12e2f3bb609eef2ffd7200273ff0ffac79fc360052cf16fca0.
//
// Solidity: event ClientAdded(int256 indexed clientIndex, address indexed client)
func (_Tanshell *TanshellFilterer) FilterClientAdded(opts *bind.FilterOpts, clientIndex []*big.Int, client []common.Address) (*TanshellClientAddedIterator, error) {

	var clientIndexRule []interface{}
	for _, clientIndexItem := range clientIndex {
		clientIndexRule = append(clientIndexRule, clientIndexItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Tanshell.contract.FilterLogs(opts, "ClientAdded", clientIndexRule, clientRule)
	if err != nil {
		return nil, err
	}
	return &TanshellClientAddedIterator{contract: _Tanshell.contract, event: "ClientAdded", logs: logs, sub: sub}, nil
}

// WatchClientAdded is a free log subscription operation binding the contract event 0xd919eebb70f12e12e2f3bb609eef2ffd7200273ff0ffac79fc360052cf16fca0.
//
// Solidity: event ClientAdded(int256 indexed clientIndex, address indexed client)
func (_Tanshell *TanshellFilterer) WatchClientAdded(opts *bind.WatchOpts, sink chan<- *TanshellClientAdded, clientIndex []*big.Int, client []common.Address) (event.Subscription, error) {

	var clientIndexRule []interface{}
	for _, clientIndexItem := range clientIndex {
		clientIndexRule = append(clientIndexRule, clientIndexItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Tanshell.contract.WatchLogs(opts, "ClientAdded", clientIndexRule, clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TanshellClientAdded)
				if err := _Tanshell.contract.UnpackLog(event, "ClientAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClientAdded is a log parse operation binding the contract event 0xd919eebb70f12e12e2f3bb609eef2ffd7200273ff0ffac79fc360052cf16fca0.
//
// Solidity: event ClientAdded(int256 indexed clientIndex, address indexed client)
func (_Tanshell *TanshellFilterer) ParseClientAdded(log types.Log) (*TanshellClientAdded, error) {
	event := new(TanshellClientAdded)
	if err := _Tanshell.contract.UnpackLog(event, "ClientAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TanshellOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tanshell contract.
type TanshellOwnershipTransferredIterator struct {
	Event *TanshellOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TanshellOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TanshellOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TanshellOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TanshellOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TanshellOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TanshellOwnershipTransferred represents a OwnershipTransferred event raised by the Tanshell contract.
type TanshellOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tanshell *TanshellFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TanshellOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tanshell.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TanshellOwnershipTransferredIterator{contract: _Tanshell.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tanshell *TanshellFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TanshellOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tanshell.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TanshellOwnershipTransferred)
				if err := _Tanshell.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tanshell *TanshellFilterer) ParseOwnershipTransferred(log types.Log) (*TanshellOwnershipTransferred, error) {
	event := new(TanshellOwnershipTransferred)
	if err := _Tanshell.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TanshellStderrIterator is returned from FilterStderr and is used to iterate over the raw logs and unpacked data for Stderr events raised by the Tanshell contract.
type TanshellStderrIterator struct {
	Event *TanshellStderr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TanshellStderrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TanshellStderr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TanshellStderr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TanshellStderrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TanshellStderrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TanshellStderr represents a Stderr event raised by the Tanshell contract.
type TanshellStderr struct {
	ClientIndex *big.Int
	MsgId       *big.Int
	Stream      []byte
	Eof         bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStderr is a free log retrieval operation binding the contract event 0x6141a2da3d119ea247441372887addada26a2f3d99dec71c9c7c198a990d6a0f.
//
// Solidity: event Stderr(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) FilterStderr(opts *bind.FilterOpts, clientIndex []*big.Int, msgId []*big.Int) (*TanshellStderrIterator, error) {

	var clientIndexRule []interface{}
	for _, clientIndexItem := range clientIndex {
		clientIndexRule = append(clientIndexRule, clientIndexItem)
	}
	var msgIdRule []interface{}
	for _, msgIdItem := range msgId {
		msgIdRule = append(msgIdRule, msgIdItem)
	}

	logs, sub, err := _Tanshell.contract.FilterLogs(opts, "Stderr", clientIndexRule, msgIdRule)
	if err != nil {
		return nil, err
	}
	return &TanshellStderrIterator{contract: _Tanshell.contract, event: "Stderr", logs: logs, sub: sub}, nil
}

// WatchStderr is a free log subscription operation binding the contract event 0x6141a2da3d119ea247441372887addada26a2f3d99dec71c9c7c198a990d6a0f.
//
// Solidity: event Stderr(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) WatchStderr(opts *bind.WatchOpts, sink chan<- *TanshellStderr, clientIndex []*big.Int, msgId []*big.Int) (event.Subscription, error) {

	var clientIndexRule []interface{}
	for _, clientIndexItem := range clientIndex {
		clientIndexRule = append(clientIndexRule, clientIndexItem)
	}
	var msgIdRule []interface{}
	for _, msgIdItem := range msgId {
		msgIdRule = append(msgIdRule, msgIdItem)
	}

	logs, sub, err := _Tanshell.contract.WatchLogs(opts, "Stderr", clientIndexRule, msgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TanshellStderr)
				if err := _Tanshell.contract.UnpackLog(event, "Stderr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStderr is a log parse operation binding the contract event 0x6141a2da3d119ea247441372887addada26a2f3d99dec71c9c7c198a990d6a0f.
//
// Solidity: event Stderr(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) ParseStderr(log types.Log) (*TanshellStderr, error) {
	event := new(TanshellStderr)
	if err := _Tanshell.contract.UnpackLog(event, "Stderr", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TanshellStdinIterator is returned from FilterStdin and is used to iterate over the raw logs and unpacked data for Stdin events raised by the Tanshell contract.
type TanshellStdinIterator struct {
	Event *TanshellStdin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TanshellStdinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TanshellStdin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TanshellStdin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TanshellStdinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TanshellStdinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TanshellStdin represents a Stdin event raised by the Tanshell contract.
type TanshellStdin struct {
	ClientIndex *big.Int
	MsgId       *big.Int
	Stream      []byte
	Eof         bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStdin is a free log retrieval operation binding the contract event 0x17e77dfd3fe1fa2d425285a427fbe76d0d39d4a1358c27ac8e5a5b07bd8a37d6.
//
// Solidity: event Stdin(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) FilterStdin(opts *bind.FilterOpts, clientIndex []*big.Int, msgId []*big.Int) (*TanshellStdinIterator, error) {

	var clientIndexRule []interface{}
	for _, clientIndexItem := range clientIndex {
		clientIndexRule = append(clientIndexRule, clientIndexItem)
	}
	var msgIdRule []interface{}
	for _, msgIdItem := range msgId {
		msgIdRule = append(msgIdRule, msgIdItem)
	}

	logs, sub, err := _Tanshell.contract.FilterLogs(opts, "Stdin", clientIndexRule, msgIdRule)
	if err != nil {
		return nil, err
	}
	return &TanshellStdinIterator{contract: _Tanshell.contract, event: "Stdin", logs: logs, sub: sub}, nil
}

// WatchStdin is a free log subscription operation binding the contract event 0x17e77dfd3fe1fa2d425285a427fbe76d0d39d4a1358c27ac8e5a5b07bd8a37d6.
//
// Solidity: event Stdin(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) WatchStdin(opts *bind.WatchOpts, sink chan<- *TanshellStdin, clientIndex []*big.Int, msgId []*big.Int) (event.Subscription, error) {

	var clientIndexRule []interface{}
	for _, clientIndexItem := range clientIndex {
		clientIndexRule = append(clientIndexRule, clientIndexItem)
	}
	var msgIdRule []interface{}
	for _, msgIdItem := range msgId {
		msgIdRule = append(msgIdRule, msgIdItem)
	}

	logs, sub, err := _Tanshell.contract.WatchLogs(opts, "Stdin", clientIndexRule, msgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TanshellStdin)
				if err := _Tanshell.contract.UnpackLog(event, "Stdin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStdin is a log parse operation binding the contract event 0x17e77dfd3fe1fa2d425285a427fbe76d0d39d4a1358c27ac8e5a5b07bd8a37d6.
//
// Solidity: event Stdin(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) ParseStdin(log types.Log) (*TanshellStdin, error) {
	event := new(TanshellStdin)
	if err := _Tanshell.contract.UnpackLog(event, "Stdin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TanshellStdoutIterator is returned from FilterStdout and is used to iterate over the raw logs and unpacked data for Stdout events raised by the Tanshell contract.
type TanshellStdoutIterator struct {
	Event *TanshellStdout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TanshellStdoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TanshellStdout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TanshellStdout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TanshellStdoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TanshellStdoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TanshellStdout represents a Stdout event raised by the Tanshell contract.
type TanshellStdout struct {
	ClientIndex *big.Int
	MsgId       *big.Int
	Stream      []byte
	Eof         bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStdout is a free log retrieval operation binding the contract event 0x9c7bf136cee05f1b5e2dbafc14a6a07080745aeda47fa726801011a380f5a93c.
//
// Solidity: event Stdout(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) FilterStdout(opts *bind.FilterOpts, clientIndex []*big.Int, msgId []*big.Int) (*TanshellStdoutIterator, error) {

	var clientIndexRule []interface{}
	for _, clientIndexItem := range clientIndex {
		clientIndexRule = append(clientIndexRule, clientIndexItem)
	}
	var msgIdRule []interface{}
	for _, msgIdItem := range msgId {
		msgIdRule = append(msgIdRule, msgIdItem)
	}

	logs, sub, err := _Tanshell.contract.FilterLogs(opts, "Stdout", clientIndexRule, msgIdRule)
	if err != nil {
		return nil, err
	}
	return &TanshellStdoutIterator{contract: _Tanshell.contract, event: "Stdout", logs: logs, sub: sub}, nil
}

// WatchStdout is a free log subscription operation binding the contract event 0x9c7bf136cee05f1b5e2dbafc14a6a07080745aeda47fa726801011a380f5a93c.
//
// Solidity: event Stdout(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) WatchStdout(opts *bind.WatchOpts, sink chan<- *TanshellStdout, clientIndex []*big.Int, msgId []*big.Int) (event.Subscription, error) {

	var clientIndexRule []interface{}
	for _, clientIndexItem := range clientIndex {
		clientIndexRule = append(clientIndexRule, clientIndexItem)
	}
	var msgIdRule []interface{}
	for _, msgIdItem := range msgId {
		msgIdRule = append(msgIdRule, msgIdItem)
	}

	logs, sub, err := _Tanshell.contract.WatchLogs(opts, "Stdout", clientIndexRule, msgIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TanshellStdout)
				if err := _Tanshell.contract.UnpackLog(event, "Stdout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStdout is a log parse operation binding the contract event 0x9c7bf136cee05f1b5e2dbafc14a6a07080745aeda47fa726801011a380f5a93c.
//
// Solidity: event Stdout(int256 indexed clientIndex, int256 indexed msgId, bytes stream, bool eof)
func (_Tanshell *TanshellFilterer) ParseStdout(log types.Log) (*TanshellStdout, error) {
	event := new(TanshellStdout)
	if err := _Tanshell.contract.UnpackLog(event, "Stdout", log); err != nil {
		return nil, err
	}
	return event, nil
}
