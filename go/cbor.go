package main

import (
	"fmt"
	"io/ioutil"
	"os"

	cbor "github.com/ipfs/go-ipld-cbor"
)

type Transaction struct {
	Type                 uint32  `json:"type,omitempty" refmt:"type,omitempty"`
	Nonce                uint64  `json:"nonce,omitempty" refmt:"nonce,omitempty"`
	SrcAddr /*[]byte*/   Address `json:"src_addr,omitempty" refmt:"src_addr,omitempty"`
	DstAddr /*[]byte*/   Address `json:"dst_addr,omitempty" refmt:"dst_addr,omitempty"`
	Data                 []byte  `json:"data,omitempty" refmt:"data,omitempty"`
	BlockHash /*[]byte*/ Hash    `json:"block_hash,omitempty" refmt:"block_hash,omitempty"`
}

type Address []byte
type Hash []byte

func main() {
	cbor.RegisterCborType(Transaction{})

	fin, err := os.Open("java.txt")
	if err != nil {
		print(err)
		return
	}
	defer fin.Close()

	tx := Transaction{
		//Type:      1,
		//Nonce:     52,
		//SrcAddr:   []byte{0x1, 0x2, 0x3, 0x4},
		//DstAddr:   []byte{0x5, 0x6, 0x7, 0x8},
		//Data:      []byte{0x1, 0x2, 0x3, 0x4},
		//BlockHash: []byte{0x5, 0x6, 0x7, 0x8},
	}
	//err = cborutil.ReadCborRPC(fin, &tx)
	cbor.DecodeReader(fin, &tx)
	if err != nil {
		print(err)
		return
	}
	fmt.Println(tx.Type)
	fmt.Println(tx.Nonce)
	fmt.Printf("%02x\n", tx.SrcAddr)
	fmt.Printf("%02x\n", tx.DstAddr)
	fmt.Printf("%02x\n", tx.Data)
	fmt.Printf("%02x\n", tx.BlockHash)

	out, err := cbor.DumpObject(tx)
	if err != nil {
		print(err)
		return
	}
	fmt.Printf("%02x\n", out)
	err = ioutil.WriteFile("go.txt", out, 0644)
	if err != nil {
		print(err)
		return
	}

}
