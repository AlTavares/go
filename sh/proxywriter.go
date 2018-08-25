package sh

import (
	"os"
)

type ProxyWriter struct {
	file  *os.File
	bytes []byte
}

func NewProxyWriter(file *os.File) *ProxyWriter {
	return &ProxyWriter{
		file: file,
	}
}

func (w *ProxyWriter) Write(p []byte) (int, error) {
	w.bytes = append(w.bytes, p...)
	return w.file.Write(p)
}

func (w *ProxyWriter) String() string {
	return string(w.bytes)
}
