// Copyright 2018 Alexander Poltoratskiy. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package fast_test

import (
	"bytes"
	"github.com/co11ter/goFAST"
	"os"
	"reflect"
	"testing"
)

var (
	encoder *fast.Encoder
	writer *bytes.Buffer
)

func init() {
	ftpl, err := os.Open("testdata/test.xml")
	if err != nil {
		panic(err)
	}
	defer ftpl.Close()
	tpls := fast.ParseXMLTemplate(ftpl)

	writer = &bytes.Buffer{}
	encoder = fast.NewEncoder(writer, tpls...)
}

func encode(msg interface{}, expect []byte, t *testing.T) {
	err := encoder.Encode(msg)
	if err != nil {
		t.Fatal("can not encode", err)
	}

	if !reflect.DeepEqual(writer.Bytes(), expect) {
		t.Fatal("data is not equal", writer.Bytes(), expect)
	}

	writer.Reset()
}

func TestDecimalEncode(t *testing.T) {
	encode(&decimalMessage1, decimalData1, t)
}