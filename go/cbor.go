package main

import (
	"fmt"

	cbg "github.com/whyrusleeping/cbor-gen"
)

type TransactionType uint8
type Hash32 []byte
type Address Hash32
type Signature = Hash32

type Transaction struct {
	Type  TransactionType `refmt:"T" json:"type"`  // Required
	From  Address         `refmt:"f" json:"from"`  // Required
	Nonce uint64          `refmt:"n" json:"nonce"` // Required
	To    Address         `refmt:"t,omitempty" json:"to,omitempty"`
	Data  []byte          `refmt:"d,omitempty" json:"data,omitempty"`
	Sig   Signature       `refmt:" ,omitempty" json:"signature,omitempty"` // Omitted before signing, use space in CBOR so that it always come first
}

var TestBytes = []byte{
	0x26, 0x34, 0x82, 0x3b, 0xc8, 0xb6, 0xcf, 0xd2, 0x66, 0xed, 0xf1, 0xce, 0x82, 0xcb, 0xac, 0xf8,
	0x1e, 0x77, 0x79, 0x0c, 0xec, 0x0f, 0x71, 0x57, 0x00, 0xea, 0xb5, 0x6c, 0xd5, 0xc6, 0xc3, 0xeb,
}

func main() {
	Gen()

	//cbor.RegisterCborType(Transaction{})
	//cbor.RegisterCborType(big.Int{})

	//fin, err := os.Open("java.txt")
	//if err != nil {
	//	print(err)
	//	return
	//}
	//defer fin.Close()

	tx := &Transaction{
		Type:  1,
		From:  TestBytes,
		Nonce: 52,
		To:    TestBytes,
		Data:  []byte{0x4, 0x13, 0x52},
		Sig:   TestBytes,
	}
	//cbor.DecodeReader(fin, tx)
	//if err != nil {
	//	print(err)
	//	return
	//}
	fmt.Println(tx.Type)
	fmt.Printf("%02x\n", tx.From)
	fmt.Println(tx.Nonce)
	fmt.Printf("%02x\n", tx.To)
	fmt.Printf("%02x\n", tx.Data)
	fmt.Printf("%02x\n", tx.Sig)
	fmt.Printf("%02x\n", tx.Sig)

	//out, err := cbor.DumpObject(*tx)
	//if err != nil {
	//	print(err)
	//	return
	//}
	//fmt.Printf("%02x\n", out)
	//err = ioutil.WriteFile("go.txt", out, 0644)
	//if err != nil {
	//	print(err)
	//	return
	//}

	//f, _ := os.Create("go_gen.txt")
	//tx.MarshalCBOR(f)
	//f.Sync()
	//f.Close()
}

func Gen() {
	if err := cbg.WriteTupleEncodersToFile("go/cbor_gen.go", "main",
		Transaction{},
	); err != nil {
		panic(err)
	}

	if err := cbg.WriteMapEncodersToFile("go/cbor_map_gen.go", "main",
		Transaction{},
	); err != nil {
		panic(err)
	}
}
