package main

/*
 */
import "C"
import (
	"frp_lib/cmd/frpc/sub"
	"github.com/fatedier/golib/crypto"
)

var stListeners LibStateListeners

//export RunFrpc
func RunFrpc(cfgFilePath *C.char) C.int {
	path := C.GoString(cfgFilePath)
	crypto.DefaultSalt = "frp"

	if err := sub.RunFrpc(path); err != nil {
		println(err.Error())
		LogPrint(err)
		return C.int(0)
	}
	sub.SetServiceOnCloseListener(&stListeners)
	sub.SetServiceProxyFailedFunc(stListeners.OnProxyFailed)
	return C.int(1)
}

//export ReloadFrpc
func ReloadFrpc() C.int {
	if err := sub.ReloadFrpc(); err != nil {
		println(err.Error())
		LogPrint(err)
		return C.int(0)
	}
	sub.SetServiceOnCloseListener(&stListeners)
	sub.SetServiceProxyFailedFunc(stListeners.OnProxyFailed)
	return C.int(1)
}

//export SetReConnectByCount
func SetReConnectByCount(reConnectByCount bool) {
	sub.SetServiceReConnectByCount(reConnectByCount)
}

func init() {
	stListeners = LibStateListeners{}
}
