package core

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"golang.org/x/net/html/charset"

	"net/url"

	"github.com/kennygrant/sanitize"
)

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, ok := charset.DetermineEncoding(data, ""); ok {
			return name
		}
	}
	return "utf-8"
}

// randomBoundary was borrowed from
// github.com/golang/go/mime/multipart/writer.go#randomBoundary
func randomBoundary() string {
	var buf [30]byte
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf[:])
}

// SanitizeFileName replaces dangerous characters in a string
// so the return value can be used as a safe file name.
func SanitizeFileName(fileName string) string {
	ext := filepath.Ext(fileName)
	cleanExt := sanitize.BaseName(ext)
	if cleanExt == "" {
		cleanExt = ".unknown"
	}
	return strings.Replace(fmt.Sprintf(
		"%s.%s",
		sanitize.BaseName(fileName[:len(fileName)-len(ext)]),
		cleanExt[1:],
	), "-", "_", -1)
}

func createMultipartReader(boundary string, data map[string][]byte) io.Reader {
	dashBoundary := "--" + boundary

	body := []byte{}
	buffer := bytes.NewBuffer(body)

	buffer.WriteString("Content-type: multipart/form-data; boundary=" + boundary + "\n\n")
	for contentType, content := range data {
		buffer.WriteString(dashBoundary + "\n")
		buffer.WriteString("Content-Disposition: form-data; name=" + contentType + "\n")
		buffer.WriteString(fmt.Sprintf("Content-Length: %d \n\n", len(content)))
		buffer.Write(content)
		buffer.WriteString("\n")
	}
	buffer.WriteString(dashBoundary + "--\n\n")
	return buffer
}

func createFormReader(data map[string]string) io.Reader {
	form := url.Values{}
	for k, v := range data {
		form.Add(k, v)
	}
	return strings.NewReader(form.Encode())
}
