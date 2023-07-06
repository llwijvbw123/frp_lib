package main

/*
#cgo darwin CFLAGS: -mmacosx-version-min=10.13
#cgo darwin LDFLAGS: -mmacosx-version-min=10.13

#ifndef DllExport
#ifdef WIN32
#define DllExport __declspec( dllexport )
#else //!WIN32
#define DllExport
#endif //WIN32
#endif //DllExport

typedef void (*LogListener) (const char* log);
extern DllExport void setLogListener(LogListener l);

typedef void (*FrpcClosedCallback)(const char* msg);
extern DllExport void setFrpcClosedCallback(FrpcClosedCallback l);

typedef void (*ProxyFailedCallback)();
extern DllExport void setProxyFailedCallback(ProxyFailedCallback l);
*/
import "C"

import (
	"frp_lib/cmd/frpc/sub"
	"frp_lib/pkg/util/version"
)

//export StopFrpc
func StopFrpc() C.int {
	if err := sub.StopFrp(); err != nil {
		println(err.Error())
		return C.int(0)
	}
	return C.int(1)
}

//export IsFrpcRunning
func IsFrpcRunning() bool {
	return sub.IsFrpRunning()
}

//export Version
func Version() string {
	return version.Full()
}

func main() {

}
