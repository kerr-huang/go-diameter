// Copyright 2013-2014 go-diameter authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package format

import (
	"bytes"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	n := Time(time.Unix(1377093974, 0))
	b := []byte{0x52, 0x14, 0xc9, 0x56}
	if v := n.Serialize(); !bytes.Equal(v, b) {
		t.Fatalf("Unexpected value. Want 0x%x, have 0x%x", b, v)
	}
	t.Log(n)
}

func TestDecodeTime(t *testing.T) {
	b := []byte{0x52, 0x14, 0xc9, 0x56}
	v, err := DecodeTime(b)
	if err != nil {
		t.Fatal(err)
	}
	if n := time.Time(v.(Time)).Unix(); n != 1377093974 {
		t.Fatalf("Unexpected value. Want 1377093974, have %d", n)
	}
	t.Log(v)
}

func BenchmarkTime(b *testing.B) {
	v := Time(time.Unix(1377093974, 0))
	for n := 0; n < b.N; n++ {
		v.Serialize()
	}
}

func BenchmarkDecodeTime(b *testing.B) {
	v := []byte{0x52, 0x14, 0xc9, 0x56}
	for n := 0; n < b.N; n++ {
		DecodeTime(v)
	}
}
