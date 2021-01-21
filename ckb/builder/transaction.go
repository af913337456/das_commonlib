package builder

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/DA-Services/das_commonlib/ckb/celltype"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/crypto"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
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

type TransactionBuilder struct {
	fromAddress    *types.Script
	totalInputCap  uint64
	totalOutputCap uint64
	fee            uint64
	tx             *types.Transaction
}

func NewTransactionBuilder0(action string, fromScript *types.Script, fee uint64) *TransactionBuilder {
	builder := NewTransactionBuilder2(fromScript, fee)
	actionData := celltype.NewActionDataBuilder().Action(celltype.GoStrToMoleculeBytes(action)).Build()
	witnessBys := celltype.NewDasWitnessData(celltype.TableType_ACTION, actionData.AsSlice()).ToWitness()
	builder.tx.Witnesses = append(builder.tx.Witnesses, witnessBys)
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

func (builder *TransactionBuilder) AddCellDep(cell *types.CellDep) *TransactionBuilder {
	if cell == nil {
		return builder
	}
	// 如果已经存在，那么不再重复添加
	for _, item := range builder.tx.CellDeps {
		if item.OutPoint.TxHash == cell.OutPoint.TxHash && item.OutPoint.Index == cell.OutPoint.Index {
			return builder
		}
	}
	builder.tx.CellDeps = append(builder.tx.CellDeps, cell)
	return builder
}

func (builder *TransactionBuilder) AddCellDeps(cellDeps []types.CellDep) *TransactionBuilder {
	cellDepsSize := len(cellDeps)
	for i := 0; i < cellDepsSize; i++ {
		builder.AddCellDep(&cellDeps[i])
	}
	return builder
}

func (builder *TransactionBuilder) AddInput(cell *types.CellInput, thisCellCap uint64) *TransactionBuilder {
	builder.totalInputCap = builder.totalInputCap + thisCellCap
	builder.tx.Inputs = append(builder.tx.Inputs, cell)
	return builder
}

// 自动计算需要的 input
func (builder *TransactionBuilder) AddInputAutoComputeItems(liveCellList *utils.LiveCellCollectResult) error {
	if needCap := builder.NeedCapacityValue(); liveCellList.Capacity < needCap {
		return fmt.Errorf("AddInputAutoComputeItems:not enough capacity, input: %d, want: %d", liveCellList.Capacity, needCap)
	} else {
		// 添加 input，只取需要的那么多个
		capCounter := uint64(0)
		for _, cell := range liveCellList.LiveCells {
			if capCounter < needCap {
				thisCellCap := cell.Output.Capacity
				input := &types.CellInput{
					Since: 0,
					PreviousOutput: &types.OutPoint{
						TxHash: cell.OutPoint.TxHash,
						Index:  cell.OutPoint.Index,
					},
				}
				builder.AddInput(input, thisCellCap)
				capCounter = capCounter + thisCellCap
			}
		}
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
	builder.AddCellDep(cell.LockDepCell())
	builder.AddCellDep(cell.TypeDepCell())
	dataBys, _ := cell.Data()
	witnessBys := celltype.NewDasWitnessData(cell.TableType(), cell.TableData()).ToWitness()
	builder.addOutputAutoComputeCap(cell.LockScript(), cell.TypeScript(), dataBys, witnessBys)
	return builder
}

func (builder *TransactionBuilder) addOutputAutoComputeCap(lockScript, typeScript *types.Script, data, witnessData []byte) *TransactionBuilder {
	output := &types.CellOutput{
		Lock: lockScript,
		Type: typeScript,
	}
	if data == nil {
		data = []byte{}
	}
	output.Capacity = output.OccupiedCapacity(data) * celltype.OneCkb
	builder.AddOutput(output, data)
	if witnessData != nil {
		builder.tx.Witnesses = append(builder.tx.Witnesses, witnessData)
	}
	return builder
}

// 强制每笔交易都要有找零
func (builder *TransactionBuilder) NeedCapacityValue() uint64 {
	if min := celltype.CkbTxMinOutputCKBValue + builder.fee; builder.totalOutputCap >= min {
		return builder.totalOutputCap + builder.fee
	} else {
		return min // 最少 61 kb + fee
	}
}

func (builder *TransactionBuilder) FromScript() *types.Script {
	return builder.fromAddress
}

// 在加完 input 和 output 后调用
func (builder *TransactionBuilder) AddChargeOutput(receiver *types.Script, signCell *utils.SystemScriptCell) *TransactionBuilder {
	builder.AddCellDep(&types.CellDep{
		OutPoint: signCell.OutPoint,
		DepType:  types.DepTypeDepGroup,
	})
	builder.tx.Outputs = append(builder.tx.Outputs, &types.CellOutput{
		Capacity: builder.totalInputCap - builder.NeedCapacityValue(),
		Lock:     receiver,
		Type:     nil,
	})
	// NOTE: need append data, or 'OutputsDataLengthMismatch: expected outputs data length (wrongLen) = outputs length (outputLen)'
	builder.tx.OutputsData = append(builder.tx.OutputsData, []byte{})
	return builder
}

// func (builder *TransactionBuilder) AddWitnessArgs(witnessData []byte) *TransactionBuilder {
// 	builder.tx.Witnesses = append(builder.tx.Witnesses, witnessData)
// 	return builder
// }

func (builder *TransactionBuilder) Log() string {
	depCellCou := len(builder.tx.CellDeps)
	inputCount := len(builder.tx.Inputs)
	outputCoun := len(builder.tx.Outputs)
	capInfo :=
		fmt.Sprintf("input cap: %d, output cap without charge: %d, need cap include fee: %d",
			builder.totalInputCap, builder.totalOutputCap, builder.NeedCapacityValue())
	return fmt.Sprintf("deps count: %d, input count: %d, output count: %d \ndata count: %d\nwitnesses count: %d\n%s",
		depCellCou, inputCount, outputCoun, len(builder.tx.OutputsData), len(builder.tx.Witnesses), capInfo)
}

func (builder *TransactionBuilder) TxHash() (string, error) {
	if needCap := builder.NeedCapacityValue(); builder.totalInputCap-needCap < 0 {
		return "", fmt.Errorf("TxHash:not enough capacity, input: %d, want: %d", builder.totalInputCap, needCap)
	}
	hash, err := builder.tx.ComputeHash()
	if err != nil {
		return "", err
	}
	return hash.String(), err
}

func (builder *TransactionBuilder) Tx() *types.Transaction {
	return builder.tx
}

func (builder *TransactionBuilder) SingleSignTransaction(key crypto.Key) error {
	data, _ := transaction.EmptyWitnessArg.Serialize() // 对应前 65 字节的签名信息
	length := make([]byte, 8)
	binary.LittleEndian.PutUint64(length, uint64(len(data)))
	hash, err := builder.tx.ComputeHash()
	if err != nil {
		return err
	}
	message := append(hash.Bytes(), length...)
	message = append(message, data...)

	// 从 1 开始，多个相同 input，填充空 []byte
	inputSize := len(builder.tx.Inputs)
	emptyWitnessList := make([][]byte, 0, inputSize-1)
	for i := 1; i < inputSize; i++ {
		emptyWitnessList = append(emptyWitnessList, []byte{})
	}

	// 添加自定义的 witness 见证数据到签名
	if len(emptyWitnessList) > 0 {
		emptyWitnessList = append(emptyWitnessList, builder.tx.Witnesses...)
		builder.tx.Witnesses = emptyWitnessList
	}
	witnessSize := len(builder.tx.Witnesses)
	for i := 0; i < witnessSize; i++ {
		_wData := builder.tx.Witnesses[i]
		length := make([]byte, 8)
		binary.LittleEndian.PutUint64(length, uint64(len(_wData)))
		message = append(message, length...) // 前 8 字节不变
		message = append(message, _wData...) // 实际数据
	}

	if message, err = blake2b.Blake256(message); err != nil {
		return err
	}
	if signed, err := key.Sign(message); err != nil {
		return err
	} else {
		wa := &types.WitnessArgs{
			Lock:       signed,
			InputType:  nil,
			OutputType: nil,
		}
		if wab, err := wa.Serialize(); err != nil {
			return err
		} else {
			if len(builder.tx.Witnesses) == 0 {
				builder.tx.Witnesses = append(builder.tx.Witnesses, wab)
			} else {
				builder.tx.Witnesses[0] = wab // 第一组放置签名的65字节
			}
		}
	}
	return nil
}

type DASTransactionBuilder struct {
	From        *types.Script
	To          *types.Script
	Amount      uint64
	Fee         uint64
	group       []int
	witnessArgs *types.WitnessArgs
	tx          *types.Transaction
}

func NewDASNormalTransactionBuilder(from, to string, amount, fee uint64) (*DASTransactionBuilder, error) {
	fromAddress, err := address.Parse(from)
	if err != nil {
		return nil, fmt.Errorf("parse from address %s error: %v", from, err)
	}
	toAddress, err := address.Parse(to)
	if err != nil {
		return nil, fmt.Errorf("parse to address %s error: %v", to, err)
	}
	if fromAddress.Mode != toAddress.Mode {
		return nil, fmt.Errorf("from address and to address with diffrent network: %v:%v", fromAddress.Mode, toAddress.Mode)
	}
	return &DASTransactionBuilder{
		From:   fromAddress.Script,
		To:     toAddress.Script,
		Amount: amount,
		Fee:    fee,
	}, nil
}

func (p *DASTransactionBuilder) GenerateTx(client rpc.Client) (*types.Transaction, error) {
	return generateTxWithIndexer(client, p)
}

func generateTxWithIndexer(client rpc.Client, p *DASTransactionBuilder) (*types.Transaction, error) {
	searchKey := &indexer.SearchKey{
		Script:     p.From,
		ScriptType: indexer.ScriptTypeLock,
	}
	collector := utils.NewLiveCellCollector(client, searchKey, indexer.SearchOrderAsc, 1000, "", utils.NewCapacityLiveCellProcessor(p.Amount+p.Fee))
	result, err := collector.Collect()
	if err != nil {
		return nil, fmt.Errorf("collect cell error: %v", err)
	}

	if result.Capacity < p.Amount+p.Fee {
		return nil, fmt.Errorf("insufficient balance: %d", result.Capacity)
	}

	systemScripts, err := utils.NewSystemScripts(client)
	if err != nil {
		return nil, fmt.Errorf("load system script error: %v", err)
	}

	tx := transaction.NewSecp256k1SingleSigTx(systemScripts)
	tx.Outputs = append(tx.Outputs, &types.CellOutput{
		Capacity: p.Amount,
		Lock:     p.To,
	})
	tx.OutputsData = [][]byte{{}}

	if result.Capacity-p.Amount-p.Fee > 0 {
		if result.Capacity-p.Amount-p.Fee >= 6100000000 {
			tx.Outputs = append(tx.Outputs, &types.CellOutput{ // 找零，如果钱包内还有 CKB 剩余的话，剩余 CKB 的数量不能小于 61 CKB
				Capacity: result.Capacity - p.Amount - p.Fee,
				Lock:     p.From,
			})
			tx.OutputsData = [][]byte{{}, {}}
		} else {
			tx.Outputs[0].Capacity = result.Capacity - p.Fee // 全部 CKB 都转出去
		}
	}
	var inputs []*types.CellInput
	for _, cell := range result.LiveCells {
		input := &types.CellInput{
			Since: 0,
			PreviousOutput: &types.OutPoint{
				TxHash: cell.OutPoint.TxHash,
				Index:  cell.OutPoint.Index,
			},
		}
		inputs = append(inputs, input)
	}
	group, witnessArgs, err := transaction.AddInputsForTransaction(tx, inputs)
	if err != nil {
		return nil, fmt.Errorf("add inputs to transaction error: %v", err)
	}

	p.group = group
	p.witnessArgs = witnessArgs
	p.tx = tx
	return tx, err
}

func (p *DASTransactionBuilder) Sign(key crypto.Key) (*types.Transaction, error) {
	err := transaction.SingleSignTransaction(p.tx, p.group, p.witnessArgs, key)
	if err != nil {
		return nil, fmt.Errorf("sign transaction error: %v", err)
	}

	return p.tx, err
}

func (p *DASTransactionBuilder) Send(client rpc.Client) (*types.Hash, error) {
	return client.SendTransaction(context.Background(), p.tx)
}
