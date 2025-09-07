package sglob_test

import (
	"testing"

	"github.com/motemen/go-testutil/dataloc"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/sglob"
)

func TestScannerNext(t *testing.T) {
	assert := assert.New(t)

	ss := sglob.NewScanner("abc")

	assert.True(ss.HasNext())
	c, ok := ss.Next()
	assert.Equal('a', c)
	assert.True(ok)
	assert.Equal("bc", ss.Rest())
	assert.Equal(2, ss.RestLen())

	assert.True(ss.HasNext())
	c, ok = ss.Next()
	assert.Equal('b', c)
	assert.True(ok)
	assert.Equal("c", ss.Rest())
	assert.Equal(1, ss.RestLen())

	assert.True(ss.HasNext())
	c, ok = ss.Next()
	assert.Equal('c', c)
	assert.True(ok)
	assert.Equal("", ss.Rest())
	assert.Equal(0, ss.RestLen())

	assert.False(ss.HasNext())
	_, ok = ss.Next()
	assert.False(ok)
}

func TestScannerPeek(t *testing.T) {
	assert := assert.New(t)

	ss := sglob.NewScanner("a")

	assert.True(ss.HasNext())
	c, ok := ss.Peek()
	assert.Equal('a', c)
	assert.True(ok)
	assert.Equal("a", ss.Rest())
	assert.Equal(1, ss.RestLen())
	assert.True(ss.HasNext())
}

func TestScannerPeekEmpty(t *testing.T) {
	assert := assert.New(t)

	ss := sglob.NewScanner("")

	_, ok := ss.Peek()
	assert.False(ok)
}

func TestScannerSkipWildcards(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		str     string
		atLeast int
		ok      bool
		rest    string
	}{
		{str: "*@example.com", atLeast: 0, ok: true, rest: "@example.com"},
		{str: "*@*.com", atLeast: 0, ok: true, rest: "@*.com"},
		{str: "**@example.com", atLeast: 0, ok: true, rest: "@example.com"},
		{str: "*?@example.com", atLeast: 1, ok: true, rest: "@example.com"},
		{str: "?*@example.com", atLeast: -1, ok: false, rest: "?*@example.com"},
		{str: "??@example.com", atLeast: -1, ok: false, rest: "??@example.com"},
		{str: "***@example.com", atLeast: 0, ok: true, rest: "@example.com"},
		{str: "**?@example.com", atLeast: 1, ok: true, rest: "@example.com"},
		{str: "*?*@example.com", atLeast: 1, ok: true, rest: "@example.com"},
		{str: "*??@example.com", atLeast: 2, ok: true, rest: "@example.com"},
		{str: "scott@*.com", atLeast: -1, ok: false, rest: "scott@*.com"},
		{str: "scott@example.com", atLeast: -1, ok: false, rest: "scott@example.com"},
		{str: "", atLeast: -1, ok: false, rest: ""},
	}

	for _, t := range tt {
		ss := sglob.NewScanner(t.str)
		atLeast, ok := ss.SkipWildcard()
		assert.Equal(t.atLeast, atLeast, "%s: 'atLeast' must be '%d'", dataloc.L(t.str), t.atLeast)
		assert.Equal(t.ok, ok, "%s: 'ok' must be '%v'", dataloc.L(t.str), t.ok)
		assert.Equal(t.rest, ss.Rest(), "%s: 'ss.Rest()' must be '%s'", dataloc.L(t.str), t.rest)
	}
}

func TestScannerSkipBefore(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		str  string
		c    rune
		n    int
		ok   bool
		rest string
	}{
		{str: "scott@example.com", c: '@', n: 5, ok: true, rest: "@example.com"},
		{str: "scott@scott@example.com", c: '@', n: 11, ok: true, rest: "@example.com"},
		{str: "scott.example.com", c: '@', n: -1, ok: false, rest: "scott.example.com"},
		{str: "scott@example.com", c: '.', n: 13, ok: true, rest: ".com"},
		{str: "scott@example.com.com", c: '.', n: 17, ok: true, rest: ".com"},
		{str: "scott@example@com", c: '.', n: -1, ok: false, rest: "scott@example@com"},
	}

	for _, t := range tt {
		ss := sglob.NewScanner(t.str)
		n, ok := ss.SkipBefore(t.c)
		assert.Equal(t.n, n, "%s: 'n' must be '%d'", dataloc.L(t.str), t.n)
		assert.Equal(t.ok, ok, "%s: 'ok' must be '%v'", dataloc.L(t.str), t.ok)
		assert.Equal(t.rest, ss.Rest(), "%s: 'ss.Rest()' must be '%s'", dataloc.L(t.str), t.rest)
	}
}
