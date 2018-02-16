package goja

import (
	"net/http"
	"net"
	"io/ioutil"
	"github.com/apex/log"
)

type Debugger struct {
	http *http.Server
}

func (d *Debugger) AttachRuntime(r *Runtime) error {
	return nil
}

func NewDebugger() (*Debugger, error) {
	debugger := &Debugger{}
	if err := debugger.init(); err != nil {
		return nil, err
	}
	return debugger, nil
}

func (d *Debugger) init() (err error) {
	ln, err := net.Listen("tcp", ":9922")
	if err != nil {
		return err
	}
	f := func() {
		d.http = &http.Server{Handler: http.HandlerFunc(d.handler)}
		if err := d.http.Serve(ln); err != nil {
			panic(err)
		}
	}
	defer func() {
		if x := recover(); x != nil {
			switch e := x.(type) {
			case error:
				err = e
			}
		}
	}()
	go f()
	return
}

func (d *Debugger) handler(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Infof("json request: %s", string(data))

		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte("{}"))
	}
}
