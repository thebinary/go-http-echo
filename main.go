package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)

type response struct {
	Method           string
	URL              *url.URL
	Proto            string
	ProtoMajor       int
	ProtoMinor       int
	Header           http.Header
	Cookie           []*http.Cookie
	Body             []byte
	ContentLength    int64
	TransferEncoding []string
	Host             string
	Form             url.Values
	PostForm         url.Values
	MultipartForm    *multipart.Form
	Trailer          http.Header
	RemoteAddr       string
	RequestURI       string
	TLS              *tls.ConnectionState
}

//Config : Service Initialization Config
type Config struct {
	Addr string
}

var c = &Config{}

func init() {
	flag.StringVar(&c.Addr, "addr", ":8000", "Listening address")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		w.Header().Set("Content-Type", "application/json")

		resp := &response{
			Method:           r.Method,
			URL:              r.URL,
			Proto:            r.Proto,
			ProtoMajor:       r.ProtoMajor,
			ProtoMinor:       r.ProtoMinor,
			Header:           r.Header,
			Cookie:           r.Cookies(),
			Body:             body,
			ContentLength:    r.ContentLength,
			TransferEncoding: r.TransferEncoding,
			Host:             r.Host,
			Form:             r.Form,
			PostForm:         r.PostForm,
			MultipartForm:    r.MultipartForm,
			Trailer:          r.Trailer,
			RemoteAddr:       r.RemoteAddr,
			RequestURI:       r.RequestURI,
			TLS:              r.TLS,
		}

		json.NewEncoder(w).Encode(resp)
	})

	log.Println("Listening on: " + c.Addr)
	log.Fatal(http.ListenAndServe(c.Addr, nil))
}
