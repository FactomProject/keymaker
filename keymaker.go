package main

import (
	"bufio"
	"fmt"
	"github.com/FactomProject/factoid"
	"github.com/FactomProject/factoid/wallet"
	"os"
	"strings"
)

func main() {

	r := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter 12 words from Koinify here: ")
	r.Scan()
	line := r.Text()
	args := strings.Fields(string(line))

	if len(args) == 12 {

		mnemonic := ""
		for _, v := range args {
			v = strings.ToLower(v)
			mnemonic = mnemonic + v + " "
		}
		mnemonic = strings.TrimSpace(mnemonic)
		private, err := wallet.MnemonicStringToPrivateKey(mnemonic)
		if err != nil {
			fmt.Print("\n\nThere was a problem with the 12 words you entered. Please check spelling against this list:\n")
			fmt.Print("https://github.com/FactomProject/go-bip39/blob/master/wordlist.go\n\n\n\n")
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

		fmt.Printf("\nFactoid Address: %v\n", adr)
		fmt.Printf("\nCheck your balance at http://explorer.factom.org/\n")
	} else {
		fmt.Printf("\n\nError: 12 and only 12 words are expected.\n")
	}

}
