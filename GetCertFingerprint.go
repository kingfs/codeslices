package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func main() {
	certPEMBlock, err := ioutil.ReadFile("kingfs.pem")
	if err != nil {
		panic(err)
	}
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		panic(err)
	}
	publicKey, err := x509.ParsePKIXPublicKey(certDERBlock.Bytes)
	if publicKey == nil {
		panic(err)
	}
	pk, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(ssh.FingerprintLegacyMD5(pk))
}
