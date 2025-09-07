package sglob_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/winebarrel/sglob"
)

func BenchmarkGlob(b *testing.B) {
	benchmarks := []struct {
		pat string
		str string
	}{
		{pat: "*@example.com", str: "scott@example.com"},
		{pat: "scott@exa?ple.com", str: "scott@example.com"},
		{pat: "*@*.com", str: "scott@example.com"},
		{pat: "scott@?*.com", str: "scott@example.com"},
	}

	for _, bm := range benchmarks {
		name := fmt.Sprintf("'%s' match '%s'", bm.pat, bm.str)

		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				m := sglob.Match(bm.pat, bm.str)

				if !m {
					b.Fail()
				}
			}
		})
	}
}

func BenchmarkRegexp(b *testing.B) {
	benchmarks := []struct {
		pat string
		str string
	}{
		{pat: `.*@example\.com`, str: "scott@example.com"},
		{pat: `scott@exa.ple\.com`, str: "scott@example.com"},
		{pat: `.*@.*\.com`, str: "scott@example.com"},
		{pat: `scott@.+\.com`, str: "scott@example.com"},
	}

	for _, bm := range benchmarks {
		name := fmt.Sprintf("'%s' match '%s'", bm.pat, bm.str)
		r := regexp.MustCompile(bm.pat)

		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				m := r.MatchString(bm.str)

				if !m {
					b.Fail()
				}
			}
		})
	}
}
