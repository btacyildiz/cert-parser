package main

import (
	"reflect"
	"testing"
)

func Test_parseBundle(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    certList
		wantErr bool
	}{
		{
			name: "only private key",
			path: "testfiles/test_key.pem",
			want: certList{},
		},
		{
			name: "cert",
			path: "testfiles/test_cert.pem",
			want: certList{
				certInfo{
					subject:    "CN=example.org,O=self,L=Amsterdam,ST=NL,C=NL",
					issuer:     "CN=example.org,O=self,L=Amsterdam,ST=NL,C=NL",
					expiration: "2027-04-15 22:07:30 +0000 UTC",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBundle(tt.path)
			got.Print()
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBundle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBundle() = %v, want %v", got, tt.want)
			}
		})
	}
}
