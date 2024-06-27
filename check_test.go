package domain

import (
	"strings"
	"testing"
)

var longName = strings.Repeat(strings.Repeat("a", 63)+".", 4)

func TestCheck(t *testing.T) {
	tests := []struct {
		n string
		e string
	}{
		// 0
		{n: "", e: "domain name is empty"},
		{n: "example.com"},
		{n: "EXAMPLE.com"},
		{n: "foo-bar.com"},
		{n: "www1.foo-bar.com"},
		// 5
		{n: "192.168.1.1.example.com"},
		{n: strings.Repeat("a", 70) + ".com", e: "domain byte length of label 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa' is 70, can't exceed 63"},
		{n: "example.com" + strings.Repeat("a", 70), e: "domain's top level domain 'comaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa' has byte length 73, can't exceed 63"},
		{n: "?", e: "domain has invalid character '?' at offset 0"},
		{n: "\t", e: "domain has invalid character '\t' at offset 0"},
		// 10
		{n: "exàmple.com", e: "domain has invalid character 'à' at offset 2"},
		{n: "www.\xbd\xb2.com", e: "domain has invalid rune at offset 4"},
		{n: "-example.com", e: "domain label '-example' at offset 0 begins with a hyphen"},
		{n: "example-.com", e: "domain label 'example-' at offset 0 ends with a hyphen"},
		{n: "example.-com", e: "domain's top level domain '-com' at offset 8 begin with a hyphen"},
		// 15
		{n: "example.com-", e: "domain's top level domain 'com-' at offset 8 ends with a hyphen"},
		{n: "example.1com", e: "domain's top level domain '1com' at offset 8 begins with a digit"},
		{n: ".example.com", e: "domain has an empty label at offset 0"},
		{n: "example..com", e: "domain has an empty label at offset 8"},
		{n: "example.com."},
		// 20
		{n: longName, e: "domain name length is 256, can't exceed 254 with a trailing dot"},
		{n: longName[:253]},
		{n: longName[:253] + "."},
		{n: longName[:254], e: "domain name length is 254, can't exceed 253 without a trailing dot"},
		{n: longName[:255] + ".", e: "domain name length is 256, can't exceed 254 with a trailing dot"},
		// 25
		{n: ".", e: "domain name is a single dot"},
	}
	for i, test := range tests {
		err := Check(test.n)
		if (err == nil) != (test.e == "") {
			if err != nil {
				t.Errorf("%2d unexpected error: %q", i, err)
			} else {
				t.Errorf("%2d unexpected nil error", i)
			}
			continue
		}
		if err != nil && err.Error() != test.e {
			t.Errorf("%2d expect error %q, got %q", i, test.e, err)
		}
	}
}
