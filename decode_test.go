package scsu

import (
	"bytes"
	"testing"
)

var (
	refEncoded = []byte{
		0x08, 0x00, 0x1B, 0x4C, 0xEA, 0x16, 0xCA, 0xD3, 0x94, 0x0F, 0x53, 0xEF, 0x61, 0x1B, 0xE5, 0x84,
		0xC4, 0x0F, 0x53, 0xEF, 0x61, 0x1B, 0xE5, 0x84, 0xC4, 0x16, 0xCA, 0xD3, 0x94, 0x08, 0x02, 0x0F,
		0x53, 0x4A, 0x4E, 0x16, 0x7D, 0x00, 0x30, 0x82, 0x52, 0x4D, 0x30, 0x6B, 0x6D, 0x41, 0x88, 0x4C,
		0xE5, 0x97, 0x9F, 0x08, 0x0C, 0x16, 0xCA, 0xD3, 0x94, 0x15, 0xAE, 0x0E, 0x6B, 0x4C, 0x08, 0x0D,
		0x8C, 0xB4, 0xA3, 0x9F, 0xCA, 0x99, 0xCB, 0x8B, 0xC2, 0x97, 0xCC, 0xAA, 0x84, 0x08, 0x02, 0x0E,
		0x7C, 0x73, 0xE2, 0x16, 0xA3, 0xB7, 0xCB, 0x93, 0xD3, 0xB4, 0xC5, 0xDC, 0x9F, 0x0E, 0x79, 0x3E,
		0x06, 0xAE, 0xB1, 0x9D, 0x93, 0xD3, 0x08, 0x0C, 0xBE, 0xA3, 0x8F, 0x08, 0x88, 0xBE, 0xA3, 0x8D,
		0xD3, 0xA8, 0xA3, 0x97, 0xC5, 0x17, 0x89, 0x08, 0x0D, 0x15, 0xD2, 0x08, 0x01, 0x93, 0xC8, 0xAA,
		0x8F, 0x0E, 0x61, 0x1B, 0x99, 0xCB, 0x0E, 0x4E, 0xBA, 0x9F, 0xA1, 0xAE, 0x93, 0xA8, 0xA0, 0x08,
		0x02, 0x08, 0x0C, 0xE2, 0x16, 0xA3, 0xB7, 0xCB, 0x0F, 0x4F, 0xE1, 0x80, 0x05, 0xEC, 0x60, 0x8D,
		0xEA, 0x06, 0xD3, 0xE6, 0x0F, 0x8A, 0x00, 0x30, 0x44, 0x65, 0xB9, 0xE4, 0xFE, 0xE7, 0xC2, 0x06,
		0xCB, 0x82,
	}
)

func TestDecode(t *testing.T) {
	r := bytes.NewBuffer([]byte{0x12, 0x9C, 0xBE, 0xC1, 0xBA, 0xB2, 0xB0})
	d := NewReader(r)
	s, err := d.ReadString()
	if err != nil {
		t.Fatal(err)
	}
	if s != "Москва" {
		t.Fatal(s)
	}

	s, err = Decode(refEncoded)
	if err != nil {
		t.Fatal(err)
	}
	if s != referenceString {
		t.Fatal(s)
	}
}

func TestReadRune(t *testing.T) {
	buf := bytes.NewBuffer([]byte{0x12, 0x9C, 0xBE, 0xC1, 0xBA, 0xB2, 0xB0})
	d := NewReader(buf)
	r, n, err := d.ReadRune()
	if err != nil {
		t.Fatal(err)
	}
	if r != '\u041C' {
		t.Fatalf("Unexpected rune: %c", r)
	}
	if n != 2 {
		t.Fatalf("Unexpected size: %d", n)
	}

	r, n, err = d.ReadRune()
	if err != nil {
		t.Fatal(err)
	}
	if r != '\u043E' {
		t.Fatalf("Unexpected rune: %c", r)
	}
	if n != 1 {
		t.Fatalf("Unexpected size: %d", n)
	}
}

func BenchmarkDecode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(refEncoded)
	}
}
