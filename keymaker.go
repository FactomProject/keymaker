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

	if len(args) == 1 || len(args) == 12 {

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
		fmt.Printf("Private key (keep secret): %x\nPublic key: %x\nFactoid Address: %v       <---- Use this to check your balance at http://explorer.factom.org/\n", priv, pub, adr)
	} else {
		if len(args) != 0 {
			fmt.Printf("\n\nError: Invalid number of arguments passed.\n")
		}
		fmt.Printf("\n\n\tFactom Mnemonic Converter\n")
		fmt.Printf("\nUsage: Run this program with your 12 word mnemonic private key passed in the command line.\n" +
			"It will print out your private and public keys as well as your Factom address.\n" +
			"You can use that address to check your account balance at http://explorer.factom.org/ .\n")
	}

}
