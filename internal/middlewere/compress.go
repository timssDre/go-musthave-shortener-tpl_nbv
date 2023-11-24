package middlewere

import (
	"bytes"
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

type gzipWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func CompressRequest() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("Content-Type") == "application/json" ||
			c.Request.Header.Get("Content-Type") == "text/html" {

			acceptEncodings := c.Request.Header.Values("Accept-Encoding")

			if foundHeader(acceptEncodings) {
				compressWriter := gzip.NewWriter(c.Writer)
				defer compressWriter.Close()
				c.Header("Content-Encoding", "gzip")
				c.Writer = &gzipWriter{c.Writer, compressWriter}
			}
		}

		contentEncodings := c.Request.Header.Values("Content-Encoding")

		if foundHeader(contentEncodings) {
			compressReader, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				log.Fatalf("error: new reader: %d", err)
				return
			}
			defer compressReader.Close()

			body, err := io.ReadAll(compressReader)
			if err != nil {
				log.Fatalf("error: read body: %d", err)
				return
			}

			c.Request.Body = io.NopCloser(bytes.NewReader(body))
			c.Request.ContentLength = int64(len(body))
		}
		c.Next()
	}
}

func foundHeader(content []string) bool {
	for _, v := range content {
		if v == "gzip" {
			return true
		}
	}

	return false
}
