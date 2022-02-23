package main

/*
#cgo CFLAGS: -I../target
#cgo LDFLAGS: -ltantivy_jrpc -L../target/release
#include "tantivy_jrpc.h"
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"

	uuid "github.com/nu7hatch/gouuid"
)

/*
  pub id: &'a str,
  pub jpc: &'a str,
  pub obj: &'a str,
  pub method: &'a str,
  pub params: serde_json::Value,
*/
type msi map[string]interface{}

func callTantivy(u, object, method string, params msi) error {
	f := map[string]interface{}{
		"id":     u,
		"jpc":    "1.0",
		"obj":    object,
		"method": method,
		"params": params,
	}
	b, err := json.Marshal(f)
	if err != nil {
		return err
	}
	p := C.CString(string(b))
	cs := (*C.uchar)(unsafe.Pointer(p))
	C.jpc(cs, C.ulong(uint64(len(string(b)))))
	return nil
}

func main() {
	fmt.Println("Hello World")
	C.init()
	u, err := uuid.NewV4()
	if err != nil {
		panic("failed to get UUID")
	}
	callTantivy(u.String(), "builder", "add_text_field", msi{
		"name": "kewlness",
	})
	callTantivy(u.String(), "builder", "add_text_field", msi{
		"name": "superKewlness",
	})
	callTantivy(u.String(), "builder", "build", msi{})
}