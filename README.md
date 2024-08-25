# go-percent-encoding 🚀

`go-percent-encoding` is a Go package designed to handle percent encoding with various character encodings. Although the URI standard (RFC 3986) doesn't specify a required character encoding for percent-encoded data, UTF-8 is widely used and recommended. However, some systems, especially older ones or those in specific regions, may use different encodings like GBK. This package allows developers to convert percent-encoded URIs between UTF-8 and other character encodings, ensuring compatibility with various systems.


## Features ✨

- **Convert Non-UTF-8 Encoded URIs to UTF-8**: Easily convert percent-encoded strings that were originally encoded using a different character encoding to a UTF-8 encoded string.
  
- **Convert UTF-8 Encoded URIs to Other Encodings**: Convert percent-encoded strings from UTF-8 to a different character encoding.

- **Custom Encoding Support**: Supports custom encoding standards via the Go text encoding package.

## Installation 📦

To install `go-percent-encoding`, run:

```bash
go get github.com/leoleaf/go-percent-encoding
```

## Usage 🛠️

### Converting to UTF-8

```go

import (
    "github.com/leoleaf/go-percent-encoding"
    "golang.org/x/text/encoding/simplifiedchinese"
)

func example() {
    input := "%D6%D0%CE%C4" // Chinese characters in GBK encoding
    decoder := simplifiedchinese.GBK.NewDecoder()
    result, err := go_percent_encoding.Other2Utf8(input, decoder)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println("Converted:", result)
}
```

### Converting from UTF-8 to Another Encoding

```go

import (
    "github.com/leoleaf/go-percent-encoding"
    "golang.org/x/text/encoding/simplifiedchinese"
)

func example() {
    input := "%E4%B8%AD%E6%96%87" // UTF-8 percent-encoded string for "中文"
    encoder := simplifiedchinese.GBK.NewEncoder()
    result, err := go_percent_encoding.Utf8ToOther(input, encoder)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println("Converted:", result)
}

```

### Encoding a Byte Slice

```go

import (
    "github.com/leoleaf/go-percent-encoding"
)

func example() {
    input := []byte("Go语言") // "Go语言" in UTF-8
    result := go_percent_encoding.Encode(input)
    fmt.Println("Encoded:", result)
}

```

## License ⚖️

This project is licensed under the MIT License - see the [LICENSE](https://github.com/leoleaf/go-percent-encoding/blob/main/LICENSE) file for details.


