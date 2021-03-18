package builder

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/DA-Services/das_commonlib/ckb/celltype"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/nervosnetwork/ckb-sdk-go/utils"
)

/**
 * Copyright (C), 2019-2020
 * FileName: transaction
 * Author:   LinGuanHong
 * Date:     2020/12/15 12:14 下午
 * Description:
 */

// systemScript info:
// https://github.com/nervosnetwork/ckb-sdk-js/blob/4921bfb1546467130898e942c2f262e3006c9ed8/packages/ckb-sdk-utils/__tests__/systemScripts/fixtures.json

var (
	EmptyWitnessArg = &types.WitnessArgs{
		Lock:       make([]byte, 65),
		InputType:  nil,
		OutputType: nil,
	}
	EmptyWitnessArgPlaceholder = make([]byte, 89)
	SignaturePlaceholder       = make([]byte, 65)
)

type TransactionBuilder struct {
	fromAddress       *types.Script
	totalInputCap     uint64
	totalOutputCap    uint64
	fee               uint64
	inputList         []celltype.TypeInputCell
	tx                *types.Transaction
	customWitnessList [][]byte
}

func NewTransactionBuilder0(action string, fromScript *types.Script, fee uint64) *TransactionBuilder {
	builder := NewTransactionBuilder2(fromScript, fee)
	actionData := celltype.NewActionDataBuilder().Action(celltype.GoStrToMoleculeBytes(action)).Build()
	witnessBys := celltype.NewDasWitnessData(celltype.TableType_ACTION, actionData.AsSlice()).ToWitness()
	builder.customWitnessList = append(builder.customWitnessList, witnessBys)
	return builder
}

func NewTransactionBuilder1(from string, fee uint64) (*TransactionBuilder, error) {
	if fromAddress, err := address.Parse(from); err != nil {
		return nil, fmt.Errorf("parse from address %s error: %v", from, err)
	} else {
		return NewTransactionBuilder2(fromAddress.Script, fee), nil
	}
}

func NewTransactionBuilder2(fromScript *types.Script, fee uint64) *TransactionBuilder {
	return &TransactionBuilder{
		totalOutputCap: 0,
		fromAddress:    fromScript,
		tx: &types.Transaction{
			Version:    0,
			HeaderDeps: []types.Hash{}, // without this, may cause param error
		},
		fee: fee,
	}
}

func (builder *TransactionBuilder) AddWitnessCellDep(cellDep *celltype.CellDepWithWitness) (*TransactionBuilder, error) {
	if cellDep == nil {
		return builder, nil
	}
	// 如果已经存在，那么不再重复添加
	for _, item := range builder.tx.CellDeps {
		if item.OutPoint.TxHash == cellDep.CellDep.OutPoint.TxHash && item.OutPoint.Index == cellDep.CellDep.OutPoint.Index {
			return builder, nil
		}
	}
	builder.tx.CellDeps = append(builder.tx.CellDeps, cellDep.CellDep)
	if cellDep.GetWitnessData != nil {
		cellDepIndex := uint32(len(builder.tx.CellDeps)) - 1
		witnessData, err := cellDep.GetWitnessData(cellDepIndex)
		if err != nil {
			return nil, fmt.Errorf("AddWitnessCellDep %s", err.Error())
		}
		builder.customWitnessList = append(builder.customWitnessList, witnessData)
	}
	return builder, nil
}

func (builder *TransactionBuilder) AddCellDep(cell *types.CellDep) *TransactionBuilder {
	_, _ = builder.AddWitnessCellDep(&celltype.CellDepWithWitness{
		CellDep:        cell,
		GetWitnessData: nil,
	})
	return builder
}

func (builder *TransactionBuilder) AddWitnessCellDeps(cellDeps []celltype.CellDepWithWitness) (*TransactionBuilder, error) {
	cellDepsSize := len(cellDeps)
	for i := 0; i < cellDepsSize; i++ {
		if _, err := builder.AddWitnessCellDep(&cellDeps[i]); err != nil {
			return nil, fmt.Errorf("AddWitnessCellDeps %s", err.Error())
		}
	}
	return builder, nil
}

func (builder *TransactionBuilder) AddCellDeps(cellDeps []types.CellDep) *TransactionBuilder {
	cellDepsSize := len(cellDeps)
	for i := 0; i < cellDepsSize; i++ {
		builder.AddCellDep(&cellDeps[i])
	}
	return builder
}

func (builder *TransactionBuilder) AddInput(typeInput celltype.TypeInputCell) *TransactionBuilder {
	builder.totalInputCap = builder.totalInputCap + typeInput.CellCap
	builder.inputList = append(builder.inputList, typeInput)
	return builder
}

func (builder *TransactionBuilder) AddWitnessInputs(cellInputs []*celltype.InputWithWitness) (*TransactionBuilder, error) {
	size := len(cellInputs)
	for i := 0; i < size; i++ {
		input := cellInputs[i]
		if _, err := builder.AddWitnessInput(input); err != nil {
			return nil, fmt.Errorf("AddWitnessInputs %s", err.Error())
		}
	}
	return builder, nil
}

func (builder *TransactionBuilder) AddWitnessInput(cellInput *celltype.InputWithWitness) (*TransactionBuilder, error) {
	builder.AddInput(cellInput.CellInput)
	if cellInput.GetWitnessData != nil {
		inputIndex := uint32(len(builder.inputList) - 1)
		witnessData, err := cellInput.GetWitnessData(inputIndex)
		if err != nil {
			return nil, fmt.Errorf("AddWitnessInput err: %s", err.Error())
		} else if witnessData != nil {
			builder.customWitnessList = append(builder.customWitnessList, witnessData)
		}
	}
	return builder, nil
}

func (builder *TransactionBuilder) OutputIndex() uint32 {
	return uint32(len(builder.tx.Outputs) - 1)
}

// 自动计算需要的 input
func (builder *TransactionBuilder) AddInputAutoComputeItems(liveCells []indexer.LiveCell, lockType celltype.LockScriptType) error {
	needCap := builder.NeedCapacityValue()
	// 添加 input，只取需要的那么多个
	capCounter := uint64(0)
	for _, cell := range liveCells {
		if capCounter < needCap {
			thisCellCap := cell.Output.Capacity
			input := celltype.TypeInputCell{
				Input: types.CellInput{
					Since: 0,
					PreviousOutput: &types.OutPoint{
						TxHash: cell.OutPoint.TxHash,
						Index:  cell.OutPoint.Index,
					},
				},
				LockType: lockType,
				CellCap:  thisCellCap,
			}
			builder.AddInput(input)
			capCounter = capCounter + thisCellCap
		}
	}
	if capCounter < needCap {
		return fmt.Errorf("AddInputAutoComputeItems:not enough capacity, input: %d, want: %d", capCounter, needCap)
	}
	return nil
}

func (builder *TransactionBuilder) AddOutput(cell *types.CellOutput, data []byte) *TransactionBuilder {
	builder.tx.Outputs = append(builder.tx.Outputs, cell)
	builder.tx.OutputsData = append(builder.tx.OutputsData, data)
	builder.totalOutputCap = builder.totalOutputCap + cell.Capacity
	return builder
}

func (builder *TransactionBuilder) AddDasSpecOutput(cell celltype.ICellType) *TransactionBuilder {
	return builder.AddDasSpecOutputWithCallback(cell, nil)
}

func (builder *TransactionBuilder) AddDasSpecOutputWithIncrementCellCap(cell celltype.ICellType, cellCap uint64) *TransactionBuilder {
	return builder.addDasSpecOutput(cell, nil, 0, cellCap)
}

func (builder *TransactionBuilder) AddDasSpecOutputWithCallback(cell celltype.ICellType, callback celltype.AddDasOutputCallback) *TransactionBuilder {
	return builder.addDasSpecOutput(cell, callback, 0, 0)
}

func (builder *TransactionBuilder) AddDasSpecOutputWithCustomCellCap(cell celltype.ICellType, cellCap uint64) *TransactionBuilder {
	return builder.addDasSpecOutput(cell, nil, cellCap, 0)
}

func (builder *TransactionBuilder) addDasSpecOutput(cell celltype.ICellType, callback celltype.AddDasOutputCallback, custom, increment uint64) *TransactionBuilder {
	builder.AddCellDep(cell.LockDepCell())
	builder.AddCellDep(cell.TypeDepCell())
	dataBys, _ := cell.Data()
	witnessBys := celltype.NewDasWitnessData(cell.TableType(), cell.TableData()).ToWitness()
	builder.addOutputAutoComputeCap(cell.LockScript(), cell.TypeScript(), dataBys, witnessBys, callback, custom, increment)
	return builder
}

func normalChargeOutputCellCap() uint64 {
	output := &types.CellOutput{
		Lock: &types.Script{
			CodeHash: types.HexToHash("0x82d76d1b75fe2fd9a27dfbaa65a039221a380d76c926f378d3f81cf3e7e13f2e"),
			HashType: types.HashTypeType,
			Args:     []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
		Type: nil,
	}
	return output.OccupiedCapacity(nil) * celltype.OneCkb
}

func (builder *TransactionBuilder) addOutputAutoComputeCap(lockScript, typeScript *types.Script,
	data, witnessData []byte, callback celltype.AddDasOutputCallback, customCellCap, incrementCellCap uint64) *TransactionBuilder {
	output := &types.CellOutput{
		Lock: lockScript,
		Type: typeScript,
	}
	if data == nil {
		data = []byte{}
	}
	if customCellCap == 0 {
		output.Capacity = output.OccupiedCapacity(data)*celltype.OneCkb + incrementCellCap
	} else {
		output.Capacity = customCellCap
	}
	if callback != nil {
		callback(output.Capacity)
	}
	builder.AddOutput(output, data)
	if witnessData != nil {
		builder.customWitnessList = append(builder.customWitnessList, witnessData)
	}
	return builder
}

func (builder *TransactionBuilder) NeedCapacityValue() uint64 {
	if min := celltype.CkbTxMinOutputCKBValue + builder.fee; builder.totalOutputCap >= min {
		return builder.totalOutputCap + builder.fee - builder.totalInputCap
	} else {
		return min // 最少 61 kb + fee
	}
}

func (builder *TransactionBuilder) FromScript() *types.Script {
	return builder.fromAddress
}

// 在加完 input 和 output 后调用
func (builder *TransactionBuilder) AddChargeOutput(receiver *types.Script, signCell *utils.SystemScriptCell) *TransactionBuilder {
	chargeCap := builder.totalInputCap - builder.totalOutputCap - builder.fee
	if chargeCap <= 0 {
		return builder
	}
	builder.AddCellDep(&types.CellDep{
		OutPoint: signCell.OutPoint,
		DepType:  types.DepTypeDepGroup,
	})
	builder.tx.Outputs = append(builder.tx.Outputs, &types.CellOutput{
		Capacity: chargeCap,
		Lock:     receiver,
		Type:     nil,
	})
	// NOTE: need append data, or 'OutputsDataLengthMismatch: expected outputs data length (wrongLen) = outputs length (outputLen)'
	builder.tx.OutputsData = append(builder.tx.OutputsData, []byte{})
	return builder
}

func (builder *TransactionBuilder) Log() string {
	depCellCou := len(builder.tx.CellDeps)
	inputCount := len(builder.inputList)
	outputCoun := len(builder.tx.Outputs)
	capInfo :=
		fmt.Sprintf("input cap: %d, output cap without charge: %d, need cap include fee: %d",
			builder.totalInputCap, builder.totalOutputCap, builder.NeedCapacityValue())
	return fmt.Sprintf("deps count: %d, input count: %d, output count: %d \ndata count: %d\nwitnesses count: %d\n%s",
		depCellCou, inputCount, outputCoun, len(builder.tx.OutputsData), len(builder.tx.Witnesses), capInfo)
}

func (builder *TransactionBuilder) TxHash() (string, error) {
	hash, err := builder.tx.ComputeHash()
	if err != nil {
		return "", err
	}
	return hash.String(), err
}

func (builder *TransactionBuilder) Tx() *types.Transaction {
	return builder.tx
}

func (builder *TransactionBuilder) addInputsForTransaction(inputs []*types.CellInput) ([]int, *types.WitnessArgs, error) {
	if len(inputs) == 0 {
		return nil, nil, errors.New("input cells empty")
	}
	group := make([]int, len(inputs))
	preInputSize := len(builder.tx.Inputs)
	start := preInputSize
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		builder.tx.Inputs = append(builder.tx.Inputs, input)
		builder.tx.Witnesses = append(builder.tx.Witnesses, []byte{})
		group[i] = start + i
	}
	builder.tx.Witnesses[start] = EmptyWitnessArgPlaceholder
	return group, EmptyWitnessArg, nil
}

func (builder *TransactionBuilder) BuildTransaction() ([]celltype.BuildTransactionRet, error) {
	size := len(builder.inputList)
	recordMap := map[celltype.LockScriptType][]*types.CellInput{}
	for i := 0; i < size; i++ {
		list := recordMap[builder.inputList[i].LockType]
		if list == nil {
			list = []*types.CellInput{}
		}
		list = append(list, &builder.inputList[i].Input)
		recordMap[builder.inputList[i].LockType] = list // same lockType is one group
	}
	retList := make([]celltype.BuildTransactionRet, 0, len(recordMap))
	for lockType, item := range recordMap {
		group, wArgs, err := builder.addInputsForTransaction(item)
		if err != nil {
			return nil, fmt.Errorf("BuildTransaction addInputsForTransaction err: %s", err.Error())
		}
		retList = append(retList, celltype.BuildTransactionRet{
			LockType:   lockType,
			Group:      group,
			WitnessArg: wArgs,
		})
	}
	if builder.customWitnessList != nil && len(builder.customWitnessList) > 0 {
		for _, witness := range builder.customWitnessList {
			builder.tx.Witnesses = append(builder.tx.Witnesses, witness)
		}
	}
	return retList, nil
}

func SingleCombineSignTransaction(tx *types.Transaction, list []celltype.BuildTransactionRet, key crypto.Key) error {
	size := len(list)
	for i := 0; i < size; i++ {
		item := list[i]
		if err := SingleSignTransaction(tx, item.Group, item.WitnessArg, key); err != nil {
			return err
		}
	}
	return nil
}

func BuildTxMessageWithoutSign(tx *types.Transaction, group []int, witnessArgs *types.WitnessArgs) ([]byte, error) {
	data, err := witnessArgs.Serialize()
	if err != nil {
		return nil, err
	}
	length := make([]byte, 8)
	binary.LittleEndian.PutUint64(length, uint64(len(data)))
	hash, err := tx.ComputeHash()
	if err != nil {
		return nil, err
	}

	message := append(hash.Bytes(), length...)
	message = append(message, data...)

	// hash the other witnesses in the group
	if len(group) > 1 {
		for i := 1; i < len(group); i++ {
			data := tx.Witnesses[i]
			length := make([]byte, 8)
			binary.LittleEndian.PutUint64(length, uint64(len(data)))
			message = append(message, length...)
			message = append(message, data...)
		}
	}
	// hash witnesses which do not in any input group
	for _, witness := range tx.Witnesses[len(tx.Inputs):] {
		length := make([]byte, 8)
		binary.LittleEndian.PutUint64(length, uint64(len(witness)))
		message = append(message, length...)
		message = append(message, witness...)
	}

	message, err = blake2b.Blake256(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func AppendSignedMsgToTx(tx *types.Transaction, group []int, witnessArgs *types.WitnessArgs, signed []byte) error {
	wa := &types.WitnessArgs{
		Lock:       signed,
		InputType:  witnessArgs.InputType,
		OutputType: witnessArgs.OutputType,
	}
	wab, err := wa.Serialize()
	if err != nil {
		return err
	}
	tx.Witnesses[group[0]] = wab
	return nil
}

func SingleSignTransaction(tx *types.Transaction, group []int, witnessArgs *types.WitnessArgs, key crypto.Key) error {
	message, err := BuildTxMessageWithoutSign(tx, group, witnessArgs)
	if err != nil {
		return err
	}
	return SignTransactionMessage(tx, group, witnessArgs, message, key)
}

func SignTransactionMessage(tx *types.Transaction, group []int, witnessArgs *types.WitnessArgs, message []byte, key crypto.Key) error {
	signed, err := key.Sign(message)
	if err != nil {
		return err
	}
	if err = AppendSignedMsgToTx(tx, group, witnessArgs, signed); err != nil {
		return err
	}
	return nil
}
