package main

import (
	"flag"
	"fmt"
	"github.com/FactomProject/factoid"
	"github.com/FactomProject/factoid/wallet"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	mnemonic := ""
	for _, v := range args {
		v = strings.ToLower(v)
		mnemonic = mnemonic + v + " "
	}
	mnemonic = strings.TrimSpace(mnemonic)
	private, err := wallet.MnemonicStringToPrivateKey(mnemonic)
	if err != nil {
		panic(err)
	}
	pub, priv, err := wallet.GenerateKeyFromPrivateKey(private)
	if err != nil {
		panic(err)
	}

	we := new(wallet.WalletEntry)
	we.AddKey(pub, priv)
	we.SetName([]byte("test"))
	we.SetRCD(factoid.NewRCD_1(pub))
	we.SetType("fct")

	address, _ := we.GetAddress()

	adr := factoid.ConvertFctAddressToUserStr(address)

	fmt.Printf("%v\n", adr)
	fmt.Printf("Private key (keep secret): %x\nPublic key: %x\nFactoid Address: %v       <---- Use this to check your balance\n", priv, pub, adr)
}
