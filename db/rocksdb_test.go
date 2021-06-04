package db

import (
	"encoding/hex"
	"fmt"
	"github.com/tecbot/gorocksdb"
	"math/rand"
	"testing"
)

/**
 * Copyright (C), 2019-2021
 * FileName: rocksdb_test
 * Author:   LinGuanHong
 * Date:     2021/1/8 11:26
 * Description:
 */

func Test_Ascill(t *testing.T) {
	bys, _ := hex.DecodeString("40")
	t.Log(bys)
	t.Log(string(bys))
}

func Test_BatchWrite(t *testing.T) {
	db, err := NewDefaultRocksNormalDb("data")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		// db.Close()
	}()
	t.Log("finish init db")
	wb := gorocksdb.NewWriteBatch()
	for i := 0; i < 10; i++ {
		if i == 9 {
			wb.Put([]byte(fmt.Sprintf("111_a%dc", i)), []byte("cc"))
		} else {
			s := hex.EncodeToString([]byte(fmt.Sprintf("op99fs%d", rand.Int31n(10))))
			wb.Put([]byte(fmt.Sprintf("name_r%se%dww", s, i)), []byte("xx"))
		}
	}
	if err = db.Write(gorocksdb.NewDefaultWriteOptions(), wb); err != nil {
		t.Error(err)
	}
	reader := db.NewIterator(gorocksdb.NewDefaultReadOptions())
	keyPrefix := []byte("111_a")
	for reader.Seek(keyPrefix); ; reader.Next() {
		if valid := reader.ValidForPrefix(keyPrefix); !valid {
			break
		}
		t.Log(string(reader.Key().Data()), string(reader.Value().Data()))
	}
}
func Test_BatchRead(t *testing.T) {
	db, err := NewDefaultRocksNormalDb("data")
	if err != nil {
		t.Error(err)
	}
	// err = db.Delete(gorocksdb.NewDefaultWriteOptions(), []byte("name1111_0x0000"))
	// if err != nil {
	// 	t.Error(err)
	// }
	// err = db.Put(gorocksdb.NewDefaultWriteOptions(), []byte("name3_0x0000"), []byte("1"))
	// if err != nil {
	// 	t.Error(err)
	// }
	// else {
	// 	fmt.Println("Test_BatchRead ===> ",string(slice.Data()))
	// }
	readOpt := gorocksdb.NewDefaultReadOptions()
	reader := db.NewIterator(readOpt)
	keyPrefix := []byte("name3_")
	for reader.SeekForPrev([]byte("name3_0xb001")); ; reader.Next() {
		if valid := reader.ValidForPrefix(keyPrefix); !valid {
			break
		}
		t.Log(string(reader.Key().Data()), string(reader.Value().Data()))
	}
}

func Test_BatchRead2(t *testing.T) {
	db, err := NewDefaultRocksNormalDb("data")
	if err != nil {
		t.Error(err)
	}
	readOpt := gorocksdb.NewDefaultReadOptions()
	reader := db.NewIterator(readOpt)
	keyPrefix := []byte("name3_")
	for reader.Seek(keyPrefix); ; reader.Next() {
		if valid := reader.ValidForPrefix(keyPrefix); !valid {
			break
		}
		t.Log(string(reader.Key().Data()), string(reader.Value().Data()))
	}
}

func Test_BatchWrite2(t *testing.T) {
	db, err := NewDefaultRocksNormalDb("data")
	if err != nil {
		t.Error(err)
	}
	wb := gorocksdb.NewWriteBatch()
	keyStrs := []string{
		"name3_0x1000", "name3_0x2000", "name3_0x1100",
		"name3_0x1200", "name3_0xa000", "name3_0xb000", "name3_0xa100", "name3_0xb100",
		"name3_0x1234", "name3_0x0000", "name3_0x0001"}
	keys := [][]byte{}
	for _, item := range keyStrs {
		keys = append(keys, []byte(item))
	}
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		wb.Put(k, []byte(keyStrs[i]))
	}
	if err = db.Write(gorocksdb.NewDefaultWriteOptions(), wb); err != nil {
		t.Error(err)
	}
}

func Test_BatchDel(t *testing.T) {
	db, err := NewDefaultRocksNormalDb("data")
	if err != nil {
		t.Error(err)
	}
	wb := gorocksdb.NewWriteBatch()
	keyStrs := []string{"1234431", "name3_0x0000"}
	for i := 0; i < len(keyStrs); i++ {
		wb.Delete([]byte(keyStrs[i]))
	}
	if err = db.Write(gorocksdb.NewDefaultWriteOptions(), wb); err != nil {
		fmt.Println(err.Error())
	}
}
