// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package sidecar

import "testing"

func TestExtractAddressFromReverse(t *testing.T) {
	tcases := []struct {
		Name     string
		Given    string
		Expected string
	}{
		{
			Name:     "IPv4",
			Given:    "1.0.0.10.in-addr.arpa",
			Expected: "10.0.0.1",
		},
		{
			Name:     "IPv4 loopback",
			Given:    "1.0.0.127.in-addr.arpa",
			Expected: "127.0.0.1",
		},
		{
			Name:     "IPv4 fully qualified",
			Given:    "1.0.0.10.in-addr.arpa.",
			Expected: "10.0.0.1",
		},
		{
			Name:     "IPv6",
			Given:    "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa",
			Expected: "2001:db8::1",
		},
		{
			Name:     "IPv6 uppercase nibbles",
			Given:    "F.E.D.C.B.A.9.8.7.6.5.4.3.2.1.0.0.0.0.0.0.0.0.0.8.B.D.0.1.0.0.2.ip6.arpa",
			Expected: "2001:db8::123:4567:89ab:cdef",
		},
		{
			Name:     "domain is left alone",
			Given:    "sentry.lenstra.fr",
			Expected: "sentry.lenstra.fr",
		},
		{
			Name:     "domain that merely mentions arpa is left alone",
			Given:    "in-addr.arpa.lenstra.fr",
			Expected: "in-addr.arpa.lenstra.fr",
		},
		{
			Name:     "too few IPv4 labels",
			Given:    "0.10.in-addr.arpa",
			Expected: "0.10.in-addr.arpa",
		},
		{
			Name:     "IPv4 label out of range",
			Given:    "1.0.0.256.in-addr.arpa",
			Expected: "1.0.0.256.in-addr.arpa",
		},
		{
			Name:     "IPv4 label is not a number",
			Given:    "1.0.0.zz.in-addr.arpa",
			Expected: "1.0.0.zz.in-addr.arpa",
		},
		{
			Name:     "too few IPv6 nibbles",
			Given:    "1.0.0.2.ip6.arpa",
			Expected: "1.0.0.2.ip6.arpa",
		},
		{
			Name:     "IPv6 nibble is not hexadecimal",
			Given:    "z.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa",
			Expected: "z.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa",
		},
		{
			Name:     "nothing but the suffix",
			Given:    ".in-addr.arpa",
			Expected: ".in-addr.arpa",
		},
		{
			Name:     "empty",
			Given:    "",
			Expected: "",
		},
	}

	for _, tc := range tcases {
		t.Run(tc.Name, func(t *testing.T) {
			if got := extractAddressFromReverse(tc.Given); got != tc.Expected {
				t.Fatalf("Was expecting '%s' but got '%s'", tc.Expected, got)
			}
		})
	}
}
