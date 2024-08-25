package go_percent_encoding_test

import (
	"testing"

	go_percent_encoding "github.com/leoleaf/go-percent-encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func TestOther2Utf8_ValidGBK(t *testing.T) {
	input := "%D6%D0%CE%C4"          // Chinese characters in GBK encoding
	expected := "%E4%B8%AD%E6%96%87" // Expected UTF-8 percent-encoded string for "中文"

	decoder := simplifiedchinese.GBK.NewDecoder()
	result, err := go_percent_encoding.Other2Utf8(input, decoder)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestOther2Utf8_InvalidEncoding(t *testing.T) {
	input := "%D6%D0%ZZ%C4" // Invalid percent-encoded string

	decoder := simplifiedchinese.GBK.NewDecoder()
	_, err := go_percent_encoding.Other2Utf8(input, decoder)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

}

func TestUtf8ToOther_ValidGBK(t *testing.T) {
	input := "%E4%B8%AD%E6%96%87" // UTF-8 percent-encoded string for "中文"
	expected := "%D6%D0%CE%C4"    // Expected GBK percent-encoded string

	encoder := simplifiedchinese.GBK.NewEncoder()
	result, err := go_percent_encoding.Utf8ToOther(input, encoder)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestUtf8ToOther_InvalidUtf8(t *testing.T) {
	input := "%E4%B8%ZZ%E6%96" // Invalid UTF-8 percent-encoded string

	encoder := simplifiedchinese.GBK.NewEncoder()
	_, err := go_percent_encoding.Utf8ToOther(input, encoder)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestEncode(t *testing.T) {
	input := []byte("Go语言") // "Go语言" in UTF-8
	expected := "%47%6F%E8%AF%AD%E8%A8%80"

	result := go_percent_encoding.Encode(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
