package go_percent_encoding

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/text/encoding"
)

type Error struct {
	index int
	msg   string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error at index %d. %s", e.index, e.msg)
}

// Other2Utf8 converts a percent-encoded string from a specific encoding (e.g., GBK) to a UTF-8 encoded string.
// This function is useful when working with URIs that have been percent-encoded using non-UTF-8 encodings.
// It takes the original encoded string and a decoder for the source encoding, returning the UTF-8 encoded result
// or an error if the input is malformed.
func Other2Utf8(raw string, from *encoding.Decoder) (string, error) {
	return convert(raw, from)
}

// Utf8ToOther converts a UTF-8 encoded, percent-encoded string to a string encoded with a different character encoding (e.g., GBK).
// This function is useful for generating URIs expected to be interpreted by systems that use non-UTF-8 encodings.
// It takes the UTF-8 encoded string and an encoder for the target encoding, returning the newly encoded result
// or an error if the input is malformed.
func Utf8ToOther(raw string, to *encoding.Encoder) (string, error) {
	return convert(raw, to)
}

// Encode encodes a byte slice into a percent-encoded string.
// Each byte is converted to its corresponding percent-encoded representation using uppercase hexadecimal digits.
// This method does not perform any character encoding, but simply represents the raw bytes as percent-encoded.
func Encode(bs []byte) string {
	var buf strings.Builder
	buf.Grow(3 * len(bs))
	for _, b := range bs {
		buf.WriteByte('%')
		buf.WriteByte(upperhex[b>>4])
		buf.WriteByte(upperhex[b&0xf])
	}
	return buf.String()
}

// bytesConverter converts a slice of bytes in one encoding to another encoding.
type bytesConverter interface {
	Bytes([]byte) ([]byte, error)
}

func convert(raw string, converter bytesConverter) (string, error) {
	var n int
	for i := 0; i < len(raw); i++ {
		c := raw[i]
		if c == '%' {
			if i+2 < len(raw) && raw[i+1] != '%' && raw[i+2] != '%' {
				n++
			} else {
				return "", &Error{
					i,
					"invalid percent encoding",
				}
			}
		}
	}

	if n == 0 {
		return raw, nil
	}

	var buf strings.Builder
	buf.Grow(len(raw))
	var temp bytes.Buffer
	temp.Grow(3 * n)
	for i := 0; i < len(raw); {
		c := raw[i]
		if c == '%' {
			if i+2 < len(raw) && raw[i+1] != '%' && raw[i+2] != '%' && ishex(raw[i+1]) && ishex(raw[i+2]) {
				temp.WriteByte(unhex(raw[i+1])<<4 | unhex(raw[i+2]))
				i += 3
				continue
			} else {
				return "", &Error{
					i,
					"invalid hex character",
				}
			}
		}
		if temp.Len() > 0 {
			bs, err := converter.Bytes(temp.Bytes())
			if err != nil {
				return "", fmt.Errorf("%w. (%w)", &Error{
					i - temp.Len(),
					"failed to encode",
				}, err)
			}
			buf.WriteString(Encode(bs))
			temp.Reset()
		}
		buf.WriteByte(c)
		i++
	}

	if temp.Len() > 0 {
		bs, err := converter.Bytes(temp.Bytes())
		if err != nil {
			return "", fmt.Errorf("%w. (%w)", &Error{
				len(raw) - temp.Len(),
				"failed to encode",
			}, err)
		}
		buf.WriteString(Encode(bs))
	}

	return buf.String(), nil
}

const upperhex = "0123456789ABCDEF"

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
