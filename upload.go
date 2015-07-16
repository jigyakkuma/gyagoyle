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
	mp := NewMuliPart(g)

	// then, upload
	res, err := g.Post(*mp)
	defer res.Body.Close()
	if err != nil {
		log.Fatalf("Post: %v", err)
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}

	if res.Status == "200 OK" {
		g.ContentUrl = string(content)
	} else {
		log.Fatal("Response status: ", res.Status)
	}

	// If gyazoID is empty,save the response header  to gyazoID file.
	g.Config.SetGyazoId(res.Header.Get("X-Gyazo-Id"))
}

func NewMuliPart(g *Gyazo) *MultiPart {
	var m MultiPart
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
	return &m
}

func (m *MultiPart) GetBody() io.Reader {
	return strings.NewReader(m.buffer.String())
}

func (g *Gyazo) Post(m MultiPart) (res *http.Response, err error) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", g.Config.Endpoint, m.GetBody())
	req.SetBasicAuth(g.Config.BasicUser, g.Config.BasicPassword)
	req.Header.Add("Content-Type", m.contentType)

	res, err = client.Do(req)
	if err != nil {
		log.Fatalf("request error: %v", err)
	}

	return
}

func (g *Gyazo) GetHost() string {
	// get hostname for filename
	url_, err := url.Parse(g.Config.Endpoint)
	if err != nil {
		log.Fatalf("url parse error: %v", err)
	}
	return url_.Host
}
