package bin

import "testing"

func TestDecoder_ID(t *testing.T) {
	var b Buffer
	const id = 0x1234
	b.PutID(id)
	b.PutString("foo bar")
	b.PutBool(true)
	b.PutInt32(-150)
	gotID, err := b.ID()
	if err != nil {
		t.Fatal(err)
	}
	if gotID != id {
		t.Fatal("mismatch")
	}
	gotStr, err := b.String()
	if err != nil {
		t.Fatal(err)
	}
	if gotStr != "foo bar" {
		t.Fatal("mismatch")
	}
	gotBool, err := b.Bool()
	if err != nil {
		t.Fatal(err)
	}
	if !gotBool {
		t.Fatal("mismatch")
	}
	gotInt32, err := b.Int32()
	if err != nil {
		t.Fatal(err)
	}
	if gotInt32 != -150 {
		t.Fatal(gotInt32)
	}
}
