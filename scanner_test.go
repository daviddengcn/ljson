// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ljson

import (
	"encoding/json"
	"math"
	"math/rand"
	"testing"
)

var ex1 = `[true,false,null,"x",1,1.5,0,-5e+2]`

var ex1i = `[
	true,
	false,
	null,
	"x",
	1,
	1.5,
	0,
	-5e+2
]`

func TestNextValueBig(t *testing.T) {
	initBig()
	var scan scanner
	item, rest, err := nextValue(jsonBig, &scan)
	if err != nil {
		t.Fatalf("nextValue: %s", err)
	}
	if len(item) != len(jsonBig) || &item[0] != &jsonBig[0] {
		t.Errorf("invalid item: %d %d", len(item), len(jsonBig))
	}
	if len(rest) != 0 {
		t.Errorf("invalid rest: %d", len(rest))
	}

	item, rest, err = nextValue(append(jsonBig, "HELLO WORLD"...), &scan)
	if err != nil {
		t.Fatalf("nextValue extra: %s", err)
	}
	if len(item) != len(jsonBig) {
		t.Errorf("invalid item: %d %d", len(item), len(jsonBig))
	}
	if string(rest) != "HELLO WORLD" {
		t.Errorf("invalid rest: %d", len(rest))
	}
}

var benchScan scanner

func BenchmarkSkipValue(b *testing.B) {
	initBig()
	for i := 0; i < b.N; i++ {
		nextValue(jsonBig, &benchScan)
	}
	b.SetBytes(int64(len(jsonBig)))
}

func diff(t *testing.T, a, b []byte) {
	for i := 0; ; i++ {
		if i >= len(a) || i >= len(b) || a[i] != b[i] {
			j := i - 10
			if j < 0 {
				j = 0
			}
			t.Errorf("diverge at %d: «%s» vs «%s»", i, trim(a[j:]), trim(b[j:]))
			return
		}
	}
}

func trim(b []byte) []byte {
	if len(b) > 20 {
		return b[0:20]
	}
	return b
}

// Generate a random JSON object.

var jsonBig []byte

const (
	big   = 10000
	small = 100
)

func initBig() {
	n := big
	if testing.Short() {
		n = small
	}
	if len(jsonBig) != n {
		b, err := json.Marshal(genValue(n))
		if err != nil {
			panic(err)
		}
		jsonBig = b
	}
}

func genValue(n int) interface{} {
	if n > 1 {
		switch rand.Intn(2) {
		case 0:
			return genArray(n)
		case 1:
			return genMap(n)
		}
	}
	switch rand.Intn(3) {
	case 0:
		return rand.Intn(2) == 0
	case 1:
		return rand.NormFloat64()
	case 2:
		return genString(30)
	}
	panic("unreachable")
}

func genString(stddev float64) string {
	n := int(math.Abs(rand.NormFloat64()*stddev + stddev/2))
	c := make([]rune, n)
	for i := range c {
		f := math.Abs(rand.NormFloat64()*64 + 32)
		if f > 0x10ffff {
			f = 0x10ffff
		}
		c[i] = rune(f)
	}
	return string(c)
}

func genArray(n int) []interface{} {
	f := int(math.Abs(rand.NormFloat64()) * math.Min(10, float64(n/2)))
	if f > n {
		f = n
	}
	if n > 0 && f == 0 {
		f = 1
	}
	x := make([]interface{}, f)
	for i := range x {
		x[i] = genValue(((i+1)*n)/f - (i*n)/f)
	}
	return x
}

func genMap(n int) map[string]interface{} {
	f := int(math.Abs(rand.NormFloat64()) * math.Min(10, float64(n/2)))
	if f > n {
		f = n
	}
	if n > 0 && f == 0 {
		f = 1
	}
	x := make(map[string]interface{})
	for i := 0; i < f; i++ {
		x[genString(10)] = genValue(((i+1)*n)/f - (i*n)/f)
	}
	return x
}
