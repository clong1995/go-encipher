package json

import (
	"bytes"
	"reflect"
	"testing"
)

func TestJson_Success(t *testing.T) {
	// The data to be encoded
	original := &student{
		Name: "小明",
		Age:  18,
	}

	// Encode the data
	var buf bytes.Buffer
	if err := Encode(original, &buf); err != nil {
		t.Fatalf("Encode() failed: %v", err)
	}

	// Decode the data
	decoded := &student{}
	if err := Decode(&buf, decoded); err != nil {
		t.Fatalf("Decode() failed: %v", err)
	}

	// Verify that the decoded data matches the original
	if !reflect.DeepEqual(original, decoded) {
		t.Errorf("Decoded data does not match original. got %+v, want %+v", decoded, original)
	}
}

func TestJson_DecodeFailsWithInvalidInput(t *testing.T) {
	// Invalid json data
	invalidData := []byte("this is not valid json data")
	buf := bytes.NewBuffer(invalidData)

	// Try to decode, expecting an error
	decoded := &student{}
	err := Decode(buf, decoded)
	if err == nil {
		t.Error("Expected decoding to fail with invalid input, but it succeeded.")
	}
}

type student struct {
	Name string
	Age  int
}
