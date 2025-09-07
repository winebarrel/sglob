package sglob_test

import (
	"testing"

	"github.com/motemen/go-testutil/dataloc"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/sglob"
)

func TestMatch(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		pat   string
		str   string
		match bool
	}{
		{pat: "*", str: "", match: true},
		{pat: "*", str: "a", match: true},
		{pat: "*", str: "ab", match: true},

		{pat: "**", str: "", match: true},
		{pat: "**", str: "a", match: true},
		{pat: "**", str: "ab", match: true},

		{pat: "?", str: "", match: false},
		{pat: "?", str: "a", match: true},
		{pat: "?", str: "ab", match: false},

		{pat: "??", str: "", match: false},
		{pat: "??", str: "a", match: false},
		{pat: "??", str: "ab", match: true},
		{pat: "??", str: "abc", match: false},

		{pat: "scott@example.com", str: "scott@example.com", match: true},
		{pat: "scott@example.com", str: "cott@example.com", match: false},
		{pat: "scott@example.com", str: "scott@exmple.com", match: false},
		{pat: "scott@example.com", str: "scott@example.co", match: false},

		{pat: "*@example.com", str: "scott@example.com", match: true},
		{pat: "*@example.com", str: "scott@scott@example.com", match: true},
		{pat: "*@example.com", str: "@example.com", match: true},

		{pat: "**@example.com", str: "scott@example.com", match: true},
		{pat: "**@example.com", str: "scott@scott@example.com", match: true},
		{pat: "**@example.com", str: "@example.com", match: true},

		{pat: "?*@example.com", str: "scott@example.com", match: true},
		{pat: "?*@example.com", str: "scott@scott@example.com", match: true},
		{pat: "?*@example.com", str: "s@example.com", match: true},
		{pat: "?*@example.com", str: "@@example.com", match: true},
		{pat: "?*@example.com", str: "@example.com", match: false},

		{pat: "??*@example.com", str: "scott@example.com", match: true},
		{pat: "??*@example.com", str: "scott@scott@example.com", match: true},
		{pat: "??*@example.com", str: "sc@example.com", match: true},
		{pat: "??*@example.com", str: "@@@example.com", match: true},
		{pat: "??*@example.com", str: "s@example.com", match: false},
		{pat: "??*@example.com", str: "@example.com", match: false},

		{pat: "*?@example.com", str: "scott@example.com", match: true},
		{pat: "*?@example.com", str: "scott@scott@example.com", match: true},
		{pat: "*?@example.com", str: "s@example.com", match: true},
		{pat: "*?@example.com", str: "@@example.com", match: true},
		{pat: "*?@example.com", str: "@example.com", match: false},

		{pat: "*??@example.com", str: "scott@example.com", match: true},
		{pat: "*??@example.com", str: "scott@scott@example.com", match: true},
		{pat: "*??@example.com", str: "sc@example.com", match: true},
		{pat: "*??@example.com", str: "@@@example.com", match: true},
		{pat: "*??@example.com", str: "s@example.com", match: false},
		{pat: "*??@example.com", str: "@example.com", match: false},

		{pat: "?*?@example.com", str: "scott@example.com", match: true},
		{pat: "?*?@example.com", str: "scott@scott@example.com", match: true},
		{pat: "?*?@example.com", str: "sc@example.com", match: true},
		{pat: "?*?@example.com", str: "@@@example.com", match: true},
		{pat: "?*?@example.com", str: "s@example.com", match: false},
		{pat: "?*?@example.com", str: "@example.com", match: false},

		{pat: "scott@*.com", str: "scott@example.com", match: true},
		{pat: "scott@*.com", str: "scott@example.example.com", match: true},
		{pat: "scott@*.com", str: "scott@example.com.com", match: true},
		{pat: "scott@*.com", str: "scott@.com", match: true},

		{pat: "scott@**.com", str: "scott@example.com", match: true},
		{pat: "scott@**.com", str: "scott@example.example.com", match: true},
		{pat: "scott@**.com", str: "scott@example.com.com", match: true},
		{pat: "scott@**.com", str: "scott@.com", match: true},

		{pat: "scott@?*.com", str: "scott@example.com", match: true},
		{pat: "scott@?*.com", str: "scott@e.com", match: true},
		{pat: "scott@?*.com", str: "scott@..com", match: true},
		{pat: "scott@?*.com", str: "scott@.com", match: false},

		{pat: "scott@??*.com", str: "scott@example.com", match: true},
		{pat: "scott@??*.com", str: "scott@example.example.com", match: true},
		{pat: "scott@??*.com", str: "scott@example.com.com", match: true},
		{pat: "scott@??*.com", str: "scott@ex.com", match: true},
		{pat: "scott@??*.com", str: "scott@...com", match: true},
		{pat: "scott@??*.com", str: "scott@e.com", match: false},
		{pat: "scott@??*.com", str: "scott@.com", match: false},

		{pat: "scott@*?.com", str: "scott@example.com", match: true},
		{pat: "scott@*?.com", str: "scott@example.example.com", match: true},
		{pat: "scott@*?.com", str: "scott@example.com.com", match: true},
		{pat: "scott@*?.com", str: "scott@e.com", match: true},
		{pat: "scott@*?.com", str: "scott@..com", match: true},
		{pat: "scott@*?.com", str: "scott@.com", match: false},

		{pat: "scott@*??.com", str: "scott@example.com", match: true},
		{pat: "scott@*??.com", str: "scott@example.example.com", match: true},
		{pat: "scott@*??.com", str: "scott@example.com.com", match: true},
		{pat: "scott@*??.com", str: "scott@ex.com", match: true},
		{pat: "scott@*??.com", str: "scott@...com", match: true},
		{pat: "scott@*??.com", str: "scott@e.com", match: false},
		{pat: "scott@*??.com", str: "scott@.com", match: false},

		{pat: "scott@?*?.com", str: "scott@example.com", match: true},
		{pat: "scott@?*?.com", str: "scott@example.example.com", match: true},
		{pat: "scott@?*?.com", str: "scott@example.com.com", match: true},
		{pat: "scott@?*?.com", str: "scott@ex.com", match: true},
		{pat: "scott@?*?.com", str: "scott@...com", match: true},
		{pat: "scott@?*?.com", str: "scott@e.com", match: false},
		{pat: "scott@?*?.com", str: "scott@.com", match: false},

		{pat: "scott@example.*", str: "scott@example.com", match: true},
		{pat: "scott@example.*", str: "scott@example.com.com", match: true},
		{pat: "scott@example.*", str: "scott@example.", match: true},

		{pat: "scott@example.**", str: "scott@example.com", match: true},
		{pat: "scott@example.**", str: "scott@example.com.com", match: true},
		{pat: "scott@example.**", str: "scott@example.", match: true},

		{pat: "scott@example.?*", str: "scott@example.com", match: true},
		{pat: "scott@example.?*", str: "scott@example.com.com", match: true},
		{pat: "scott@example.?*", str: "scott@example.c", match: true},
		{pat: "scott@example.?*", str: "scott@example..", match: true},
		{pat: "scott@example.?*", str: "scott@example.", match: false},

		{pat: "scott@example.??*", str: "scott@example.com", match: true},
		{pat: "scott@example.??*", str: "scott@example.com.com", match: true},
		{pat: "scott@example.??*", str: "scott@example.co", match: true},
		{pat: "scott@example.??*", str: "scott@example...", match: true},
		{pat: "scott@example.??*", str: "scott@example.c", match: false},
		{pat: "scott@example.??*", str: "scott@example.", match: false},

		{pat: "scott@example.*?", str: "scott@example.com", match: true},
		{pat: "scott@example.*?", str: "scott@example.com.com", match: true},
		{pat: "scott@example.*?", str: "scott@example.c", match: true},
		{pat: "scott@example.*?", str: "scott@example..", match: true},
		{pat: "scott@example.*?", str: "scott@example.", match: false},

		{pat: "scott@example.*??", str: "scott@example.com", match: true},
		{pat: "scott@example.*??", str: "scott@example.com.com", match: true},
		{pat: "scott@example.*??", str: "scott@example.co", match: true},
		{pat: "scott@example.*??", str: "scott@example...", match: true},
		{pat: "scott@example.*??", str: "scott@example.c", match: false},
		{pat: "scott@example.*??", str: "scott@example.", match: false},

		{pat: "scott@example.?*?", str: "scott@example.com", match: true},
		{pat: "scott@example.?*?", str: "scott@example.com.com", match: true},
		{pat: "scott@example.?*?", str: "scott@example.co", match: true},
		{pat: "scott@example.?*?", str: "scott@example...", match: true},
		{pat: "scott@example.?*?", str: "scott@example.c", match: false},
		{pat: "scott@example.?*?", str: "scott@example.", match: false},

		{pat: "*@example.*", str: "scott@example.com", match: true},
		{pat: "*@example.*", str: "xscott@example.co", match: true},
		{pat: "*@example.*", str: "scott@scott@example.com.com", match: true},

		{pat: "*?@example.?*", str: "scott@example.com", match: true},
		{pat: "*?@example.?*", str: "scott@scott@example.com.com", match: true},
		{pat: "*?@example.?*", str: "s@example.c", match: true},
		{pat: "*?@example.?*", str: "@@example..", match: true},
		{pat: "*?@example.?*", str: "s@example.", match: false},
		{pat: "*?@example.?*", str: "@example.c", match: false},
	}

	for _, t := range tt {
		m := sglob.Match(t.pat, t.str)
		assert.Equal(t.match, m, "%s: '%s' ~= '%s' must be '%v'", dataloc.L(t.pat), t.pat, t.str, t.match)
	}
}
