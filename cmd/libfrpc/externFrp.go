package main

/*
 */
import "C"
import (
	"github.com/Bellegar/frp_lib/cmd/frpc/sub"
	"github.com/fatedier/golib/crypto"
)

//export RunFrpc
func RunFrpc(cfgFilePath *C.char) C.int {
	path := C.GoString(cfgFilePath)
	crypto.DefaultSalt = "frp"

	if err := sub.RunFrpc(path); err != nil {
		println(err.Error())
		return C.int(0)
	}
	return C.int(1)
}
