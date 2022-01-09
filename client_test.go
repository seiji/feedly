package feedly

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (fn RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req), nil
}

func NewTestClient(expected interface{}) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(func(_ *http.Request) *http.Response {
			b, _ := json.Marshal(expected)
			return &http.Response{
				Status:           "",
				StatusCode:       http.StatusOK,
				Proto:            "",
				ProtoMajor:       0,
				ProtoMinor:       0,
				Header:           make(http.Header),
				Body:             ioutil.NopCloser(bytes.NewBuffer(b)),
				ContentLength:    0,
				TransferEncoding: []string{},
				Close:            false,
				Uncompressed:     false,
				Trailer:          map[string][]string{},
				Request:          &http.Request{},
				TLS: &tls.ConnectionState{
					Version:                     0,
					HandshakeComplete:           false,
					DidResume:                   false,
					CipherSuite:                 0,
					NegotiatedProtocol:          "",
					NegotiatedProtocolIsMutual:  false,
					ServerName:                  "",
					PeerCertificates:            []*x509.Certificate{},
					VerifiedChains:              [][]*x509.Certificate{},
					SignedCertificateTimestamps: [][]byte{},
					OCSPResponse:                []byte{},
					TLSUnique:                   []byte{},
				},
			}
		}),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}
}
