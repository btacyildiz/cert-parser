package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

const (
	certificate   = "CERTIFICATE"
	privateKey    = "PRIVATE KEY"
	rsaPrivateKey = "RSA PRIVATE KEY"
)

type certInfo struct {
	subject    string
	issuer     string
	expiration string
	serial     string
}

type certList []certInfo

func (l certList) Print() {
	for _, cert := range l {
		fmt.Println("------------------------------")
		fmt.Println("Subject   : ", cert.subject)
		fmt.Println("Issuer    : ", cert.issuer)
		fmt.Println("Expiration: ", cert.expiration)
		fmt.Println("Serial    : ", cert.serial)
		fmt.Println("------------------------------")
	}
}

func parseBundle(path string) (certList, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read cert bundle %s", err)
	}
	list := certList{}
	for block, rest := pem.Decode(data); block != nil; block, rest = pem.Decode(rest) {
		switch block.Type {
		case certificate:
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("unexpected cert found")
			}
			list = append(list, certInfo{
				subject:    cert.Subject.String(),
				issuer:     cert.Issuer.String(),
				expiration: cert.NotAfter.String(),
				serial:     cert.Subject.SerialNumber})
		case privateKey, rsaPrivateKey:
			fmt.Println("Private key is found, skipping...")
		default:
			return nil, fmt.Errorf("unexpected block type: %s", block.Type)
		}
	}
	return list, nil
}
