package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

type MultiPart struct {
	writer      *multipart.Writer
	buffer      bytes.Buffer
	contentType string
}

func (g *Gyazo) Upload() {
	var mp MultiPart
	mp.New(g)

	// then, upload
	res, err := http.Post(g.Endpoint, mp.contentType, mp.GetBody())
	if err != nil {
		log.Fatalf("Post: %v", err)
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}

	g.ContentUrl = string(content)

	// If gyazoID is empty,save the response header  to gyazoID file.
	g.Config.SetGyazoId(res.Header.Get("X-Gyazo-Id"))
}

func (m *MultiPart) New(g *Gyazo) {
	m.writer = multipart.NewWriter(&m.buffer)
	err := m.writer.WriteField("id", g.Config.GyazoId)
	part, err := m.writer.CreateFormFile("imagedata", g.GetHost())
	if err != nil {
		log.Fatalf("CreateFormFile: %v", err)
	}

	part.Write(g.ImageBinary)
	m.contentType = m.writer.FormDataContentType()
	err = m.writer.Close()
	if err != nil {
		log.Fatalf("Close: %v", err)
	}
	return
}

func (m *MultiPart) GetBody() io.Reader {
	return strings.NewReader(m.buffer.String())
}

func (g *Gyazo) GetHost() string {
	// get hostname for filename
	url_, err := url.Parse(g.Endpoint)
	if err != nil {
		log.Fatalf("url parse error: %v", err)
	}
	return url_.Host
}
