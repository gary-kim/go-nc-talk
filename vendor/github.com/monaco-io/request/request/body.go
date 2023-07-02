package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"

	"github.com/monaco-io/request/xcontext"
	"gopkg.in/yaml.v2"
)

// BodyString body of type string
type BodyString struct {
	Data string
}

// Apply string body
func (b BodyString) Apply(ctx *xcontext.Context) {
	bBytes := bytes.NewReader([]byte(b.Data))
	rc, ok := io.Reader(bBytes).(io.ReadCloser)
	if !ok && bBytes != nil {
		rc = io.NopCloser(bBytes)
	}

	ctx.Request.Body = rc
	ctx.Request.ContentLength = int64(bytes.NewBufferString(b.Data).Len())
}

// Valid string body valid?
func (b BodyString) Valid() bool {
	return b.Data != ""
}

// BodyJSON body of type json
type BodyJSON struct {
	Data interface{}
}

// Apply json body
func (b BodyJSON) Apply(ctx *xcontext.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := json.NewEncoder(buf).Encode(b.Data); err != nil {
			ctx.SetError(fmt.Errorf("unknown json encoded type: %T", b.Data))
			return
		}
	}

	ctx.Request.Body = io.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
	ctx.SetContentType(xcontext.JSON)
}

// Valid json body valid?
func (b BodyJSON) Valid() bool {
	return b.Data != nil
}

// BodyXML body of type xml
type BodyXML struct {
	Data interface{}
}

// Apply xml body
func (b BodyXML) Apply(ctx *xcontext.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := xml.NewEncoder(buf).Encode(b.Data); err != nil {
			ctx.SetError(fmt.Errorf("unknown xml encoded type: %T", b.Data))
			return
		}
	}

	ctx.Request.Body = io.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
	ctx.SetContentType(xcontext.XML)
}

// Valid xml body valid?
func (b BodyXML) Valid() bool {
	return b.Data != nil
}

// BodyYAML body of type yaml
type BodyYAML struct {
	Data interface{}
}

// Apply yaml body
func (b BodyYAML) Apply(ctx *xcontext.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	default:
		if err := yaml.NewEncoder(buf).Encode(b.Data); err != nil {
			ctx.SetError(fmt.Errorf("unknown yaml encoded type: %T", b.Data))
			return
		}
	}

	ctx.Request.Body = io.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
}

// Valid json body valid?
func (b BodyYAML) Valid() bool {
	return b.Data != nil
}

// BodyURLEncodedForm application/x-www-form-urlencoded
type BodyURLEncodedForm struct {
	Data interface{}
}

// Apply application/x-www-form-urlencoded
func (b BodyURLEncodedForm) Apply(ctx *xcontext.Context) {
	buf := &bytes.Buffer{}

	switch b.Data.(type) {
	case string:
		buf.WriteString(b.Data.(string))
	case []byte:
		buf.Write(b.Data.([]byte))
	case map[string]string:
		data := make(url.Values)
		for k, v := range b.Data.(map[string]string) {
			data.Set(k, v)
		}
		buf.WriteString(data.Encode())
	case map[string][]string:
		buf.WriteString(url.Values(b.Data.(map[string][]string)).Encode())
	case url.Values:
		buf.WriteString(b.Data.(url.Values).Encode())
	default:
		ctx.SetError(fmt.Errorf("unknown urlencoded type: %T", b.Data))
		return
	}

	ctx.Request.Body = io.NopCloser(buf)
	ctx.Request.ContentLength = int64(buf.Len())
	ctx.SetContentType(xcontext.URLEncodedForm)
}

// Valid application/x-www-form-urlencoded valid?
func (b BodyURLEncodedForm) Valid() bool {
	return b.Data != nil
}

// FormFile represents the file form field data.
type FormFile struct {
	Name   string
	Reader io.Reader
}

// BodyForm represents the supported form fields by file and string data.
type BodyForm struct {
	Fields map[string]string
	Files  []string
}

// Apply Form Data
func (fd BodyForm) Apply(ctx *xcontext.Context) {
	var (
		err error
		buf bytes.Buffer
	)

	multipartWriter := multipart.NewWriter(&buf)

	for _, filePath := range fd.Files {
		var (
			w    io.Writer
			data []byte
		)

		w, err = multipartWriter.CreateFormFile(path.Base(filePath), path.Base(filePath))
		if err != nil {
			err = fmt.Errorf("cread form file failed: %s", err)
			goto ErrorHandler
		}
		data, err = os.ReadFile(filePath)
		if err != nil {
			err = fmt.Errorf("read local file failed: %s", err)
			goto ErrorHandler
		}
		_, err = w.Write(data)
		if err != nil {
			err = fmt.Errorf("write byte to writer failed: %s", err)
			goto ErrorHandler
		}
	}

	// Populate the other parts of the form (if there are any)
	for k, v := range fd.Fields {
		multipartWriter.WriteField(k, v)
	}
	if err = multipartWriter.Close(); err != nil {
		return
	}
	if buf.Len() == 0 {
		return
	}

	ctx.Request.Body = io.NopCloser(&buf)
	ctx.Request.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	return

ErrorHandler:
	ctx.SetError(err)
}

// Valid form body valid?
func (fd BodyForm) Valid() bool {
	if fd.Fields == nil && fd.Files == nil {
		return false
	}
	return true
}
