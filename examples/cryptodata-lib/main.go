//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"fmt"
	"github.com/ligato/cn-infra/db/cryptodata"
	"os"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
)

func main() {
	// Read private key
	bytes, err := ioutil.ReadFile("key.pem")
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(bytes)
	if block == nil {
		panic("failed to decode PEM for key key.pem")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// Read public key
	bytes, err = ioutil.ReadFile("key-pub.pem")
	if err != nil {
		panic(err)
	}
	block, _ = pem.Decode(bytes)
	if block == nil {
		panic("failed to decode PEM for key key-pub.pem")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := pubInterface.(*rsa.PublicKey)

	// Create cryptodata client
	client := cryptodata.NewClient(cryptodata.ClientConfig{
		PrivateKeys: []*rsa.PrivateKey{privateKey},
	})

	// Pass 1st argument from CLI as string to encrypt
	input := []byte(os.Args[1])
	fmt.Printf("Input %v\n", string(input))

	// Encrypt input string using public key
	encrypted, err := client.EncryptArbitrary(input, publicKey)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}
	fmt.Printf("Encrypted %v\n", encrypted)

	// Decrypt previously encrypted input string
	decrypted, err := client.DecryptArbitrary(encrypted)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}
	fmt.Printf("Decrypted %v\n", string(decrypted))
}