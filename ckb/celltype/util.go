package celltype

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

/**
 * Copyright (C), 2019-2020
 * FileName: util
 * Author:   LinGuanHong
 * Date:     2020/12/18 2:57 下午
 * Description:
 */

// int64 4Byte
// func PackCellDataWithVersion(version uint32, cellData []byte) []byte {
// 	bytebuf := bytes.NewBuffer([]byte{})
// 	_ = binary.Write(bytebuf, binary.LittleEndian, version)
// 	return append(bytebuf.Bytes(), cellData...)
// }

// func UnpackCellDataWithVersion(cellData []byte) []byte {
// 	return cellData[CellVersionByteLen:]
// }

func GoHexToMoleculeHash(hexStr string) Hash {
	hexBytes, _ := hex.DecodeString(hexStr)
	return *HashFromSliceUnchecked(hexBytes)
}

func GoUint8ToMoleculeU8(i uint8) Uint8 {
	bytebuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuf, binary.LittleEndian, i)
	return *Uint8FromSliceUnchecked(bytebuf.Bytes())
}

func GoUint32ToMoleculeU32(i uint32) Uint32 {
	bytebuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuf, binary.LittleEndian, i)
	return *Uint32FromSliceUnchecked(bytebuf.Bytes())
}

func GoUint64ToMoleculeU64(i uint64) Uint64 {
	bytebuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuf, binary.LittleEndian, i)
	return *Uint64FromSliceUnchecked(bytebuf.Bytes())
}

func GoStrToMoleculeBytes(str string) Bytes {
	strBytes := []byte(str)
	return GoBytesToMoleculeBytes(strBytes)
}

func GoBytesToMoleculeBytes(bys []byte) Bytes {
	_bytesBuilder := NewBytesBuilder()
	for _, bye := range bys {
		_bytesBuilder.Push(*ByteFromSliceUnchecked([]byte{bye}))
	}
	return _bytesBuilder.Build()
}

func GoByteToMoleculeByte(byte byte) Byte {
	return NewByte(byte)
}

func GoTimeUnixToMoleculeBytes(timeSec int64) [8]Byte {
	bytebuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuf, binary.LittleEndian, timeSec)
	timestampByteArr := [8]Byte{}
	for index, bye := range bytebuf.Bytes() {
		timestampByteArr[index] = *ByteFromSliceUnchecked([]byte{bye})
	}
	return timestampByteArr
}

func GoCkbScriptToMoleculeScript(script types.Script) Script {
	// 这里 data 类型应该就是 0x00 ，type 就是 0x01
	ht := 0x00
	if script.HashType == types.HashTypeType {
		ht = 0x01
	}
	return NewScriptBuilder().
		CodeHash(GoHexToMoleculeHash(script.CodeHash.String())).
		HashType(GoByteToMoleculeByte(byte(ht))). // todo
		Args(GoBytesToMoleculeBytes(script.Args)).
		Build()
}

func MoleculeU32ToGo(bys []byte) (uint32, error) {
	var t uint32
	bytesBuffer := bytes.NewBuffer(bys)
	if err := binary.Read(bytesBuffer, binary.LittleEndian, &t); err != nil {
		return 0, err
	}
	return t, nil
}

func ParseTxWitnessToDasWitnessObj(rawData []byte) (*ParseDasWitnessBysDataObj, error) {
	ret := &ParseDasWitnessBysDataObj{}
	dasWitnessObj, err := NewDasWitnessDataFromSlice(rawData)
	if err != nil {
		return nil, fmt.Errorf("fail to parse dasWitness data: %s", err.Error())
	}
	ret.WitnessObj = dasWitnessObj
	data, err := DataFromSlice(dasWitnessObj.TableBys, false)
	if err != nil {
		return nil, fmt.Errorf("fail to parse data: %s", err.Error())
	}
	ret.MoleculeData = data
	if data.Old().IsNone() {
		ret.MoleculeOldDataEntity = nil
	} else {
		if oldData, err := DataEntityFromSlice(data.Old().AsSlice(), false); err != nil {
			return nil, fmt.Errorf("fail to parse old dataEntity: %s", err.Error())
		} else {
			ret.MoleculeOldDataEntity = oldData
		}
	}
	newData, err := DataEntityFromSlice(data.New().AsSlice(), false)
	if err != nil {
		return nil, fmt.Errorf("fail to parse new dataEntity: %s", err.Error())
	}
	ret.MoleculeNewDataEntity = newData
	return ret, nil
}

func buildDasCommonMoleculeDataObj(oldIndex, newIndex uint32, oldMolecule, newMolecule ICellData) *Data {
	var (
		oldData DataEntity
		newData = NewDataEntityBuilder().
			Index(GoUint32ToMoleculeU32(newIndex)).
			Version(GoUint32ToMoleculeU32(1)).
			Entity(*BytesFromSliceUnchecked(newMolecule.AsSlice())).
			Build()
		dataBuilder = NewDataBuilder().
				New(NewDataEntityOptBuilder().Set(newData).Build())
	)
	if oldMolecule != nil {
		oldData = NewDataEntityBuilder().
			Index(GoUint32ToMoleculeU32(oldIndex)).
			Version(GoUint32ToMoleculeU32(1)).
			Entity(*BytesFromSliceUnchecked(oldMolecule.AsSlice())).
			Build()
		dataBuilder.Old(NewDataEntityOptBuilder().Set(oldData).Build())
	}
	d := dataBuilder.Build()
	return &d
}

func FindTargetTypeScriptByInputList(ctx context.Context, rpcClient rpc.Client, inputList []*types.CellInput, codeHash types.Hash) (*types.Script, error) {
	for _, item := range inputList {
		tx, err := rpcClient.GetTransaction(ctx, item.PreviousOutput.TxHash)
		if err != nil {
			return nil, fmt.Errorf("FindSenderLockScriptByInputList err: %s", err.Error())
		}
		for _, output := range tx.Transaction.Outputs {
			if output.Type == nil && output.Lock.CodeHash == codeHash && output.Lock.HashType == types.HashTypeType {
				return &types.Script{
					CodeHash: codeHash,
					HashType: types.HashTypeType,
					Args:     output.Lock.Args,
				}, nil
			}
		}
	}
	return nil, errors.New("FindSenderLockScriptByInputList not found")
}
