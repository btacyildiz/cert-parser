package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
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

func parseBundle(params certParserParams) (certList, error) {
	data, err := os.ReadFile(params.bundlePath)
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

			if params.searchText != "" && !doesCertContains(cert, params.searchText) {
				continue
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

func doesCertContains(cert *x509.Certificate, searchText string) bool {
	return doesContain(cert.Subject.String(), searchText) ||
		doesContain(cert.Issuer.String(), searchText) ||
		doesContain(cert.Issuer.SerialNumber, searchText) ||
		doesContain(cert.Subject.SerialNumber, searchText)
}

func doesContain(text, searchText string) bool {
	return strings.Contains(strings.ToLower(text), strings.ToLower(searchText))
}
