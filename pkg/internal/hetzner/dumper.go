package hetzner

import (
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

// Dumper dumps the http.Request and http.Response message payload for debugging.
type Dumper interface {
	DumpRequest(*http.Request)
	DumpResponse(*http.Response)
}

// DiscardDumper returns a no-op dumper.
func DiscardDumper() Dumper {
	return new(discardDumper)
}

type discardDumper struct{}

// DumpRequest implements the Dumper interface.
func (*discardDumper) DumpRequest(*http.Request) {

}

// DumpResponse implements the Dumper interface.
func (*discardDumper) DumpResponse(*http.Response) {

}

// StandardDumper returns a standard dumper.
func StandardDumper(body bool) Dumper {
	return &standardDumper{out: os.Stdout, body: body}
}

type standardDumper struct {
	body bool
	out  io.Writer
}

// DumpRequest implements the Dumper interface.
func (s *standardDumper) DumpRequest(req *http.Request) {
	dump, _ := httputil.DumpRequestOut(req, s.body)
	_, _ = s.out.Write(dump)
}

// DumpResponse implements the Dumper interface.
func (s *standardDumper) DumpResponse(res *http.Response) {
	dump, _ := httputil.DumpResponse(res, s.body)
	_, _ = s.out.Write(dump)
}
