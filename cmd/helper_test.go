package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveHostsEntries(t *testing.T) {
	testCases := []struct {
		input    string
		ip       string
		domains  []string
		expected string
	}{

		{
			`
`,
			`1.2.3.4`,
			[]string{`example.com`},
			``,
		},

		{
			`1.2.3.4 example.com

`,
			`1.2.3.4`,
			[]string{`example.com`},
			``,
		},

		{
			`1.2.3.4 example.com test.com

`,
			`1.2.3.4`,
			[]string{`example.com`},
			`1.2.3.4         test.com
`,
		},

		{
			`1.2.3.4 example.com test.com

`,
			`1.2.3.4`,
			[]string{`example.com`, `stuff.com`},
			`1.2.3.4         test.com
`,
		},

		{
			`1.2.3.4 example.com # This is an example
`,
			`1.2.3.4`,
			[]string{`example.com`},
			``,
		},

		{
			`1.2.3.4 test.example.com    example.com # This is an example
`,
			`1.2.3.4`,
			[]string{`example.com`},
			`1.2.3.4         test.example.com # This is an example
`,
		},

		{
			`1.2.3.4 test.example.com    example.com      # This is an example
`,
			`1.2.3.4`,
			[]string{`stuff.example.com`},
			`1.2.3.4         test.example.com example.com # This is an example
`,
		},

		{
			`1.2.3.4 test.example.com    example.com      # This is an example


1.2.3.4 test.example.com example.com # This is an example
`,
			`2.3.4.5`,
			[]string{`stuff.example.com`},
			`1.2.3.4         test.example.com example.com # This is an example

1.2.3.4         test.example.com example.com # This is an example
`,
		},

		{
			`##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       localhost
255.255.255.255 broadcasthost
::1             localhost

1.2.3.4         hest.dk
`,
			`1.2.3.4`,
			[]string{`hest.dk`},
			`##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       localhost
255.255.255.255 broadcasthost
::1             localhost
`,
		},
	}

	for _, testCase := range testCases {

		actual := removeHostsEntries(testCase.ip, testCase.domains, testCase.input)

		assert.Equal(t, testCase.expected, actual, "remove %s", formatIpAndDomains(testCase.ip, testCase.domains))
	}
}

func TestAddHostsEntries(t *testing.T) {
	testCases := []struct {
		input    string
		ip       string
		domains  []string
		expected string
	}{

		{
			`
`,
			`1.2.3.4`,
			[]string{`example.com`},
			`1.2.3.4         example.com # added by go run github.com/mikkelricky/hosts-harker
`,
		},

		{
			`1.2.3.4 example.com

		`,
			`1.2.3.4`,
			[]string{`example.com`},
			`1.2.3.4         example.com # added by go run github.com/mikkelricky/hosts-harker
`,
		},

		{
			`1.2.3.4 example.com # This is an example
`,
			`1.2.3.4`,
			[]string{`example.com`},
			`1.2.3.4         example.com # added by go run github.com/mikkelricky/hosts-harker
`,
		},

		{
			`1.2.3.4 test.example.com    example.com # This is an example
`,
			`1.2.3.4`,
			[]string{`example.com`},
			`1.2.3.4         test.example.com # This is an example
1.2.3.4         example.com # added by go run github.com/mikkelricky/hosts-harker
`,
		},

		{
			`1.2.3.4 test.example.com    example.com      # This is an example
		`,
			`1.2.3.4`,
			[]string{`stuff.example.com`},
			`1.2.3.4         test.example.com example.com # This is an example
1.2.3.4         stuff.example.com # added by go run github.com/mikkelricky/hosts-harker
`,
		},

		{
			`1.2.3.4 test.example.com    example.com      # This is an example

1.2.3.4 test.example.com example.com # This is an example
`,
			`2.3.4.5`,
			[]string{`stuff.example.com`},
			`1.2.3.4         test.example.com example.com # This is an example
1.2.3.4         test.example.com example.com # This is an example
2.3.4.5         stuff.example.com # added by go run github.com/mikkelricky/hosts-harker
`,
		},

		{
			`##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       localhost
255.255.255.255 broadcasthost
::1             localhost

1.2.3.4         hest.dk
`,
			`1.2.3.4`,
			[]string{`hest.dk`},
			`##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       localhost
255.255.255.255 broadcasthost
::1             localhost
1.2.3.4         hest.dk # added by go run github.com/mikkelricky/hosts-harker
`,
		},

		{
			`1.2.3.4 test.dk
`,
			`1.2.3.4`,
			[]string{`test.dk`, `hest.dk`},
			`1.2.3.4         test.dk hest.dk # added by go run github.com/mikkelricky/hosts-harker
`,
		},

		{
			`##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       localhost
255.255.255.255 broadcasthost
::1             localhost

1.2.3.4         hest.dk
`,
			`1.2.3.4`,
			[]string{`hest.dk`},
			`##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1       localhost
255.255.255.255 broadcasthost
::1             localhost
1.2.3.4         hest.dk # added by go run github.com/mikkelricky/hosts-harker
`,
		},
	}

	for _, testCase := range testCases {

		actual := addHostsEntries(testCase.ip, testCase.domains, testCase.input)

		assert.Equal(t, testCase.expected, actual)
	}
}
