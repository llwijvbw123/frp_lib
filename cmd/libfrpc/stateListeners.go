package main

/*
#include <stdio.h>
#include <stdbool.h>

#ifndef DllExport
#ifdef WIN32
#define DllExport __declspec( dllexport )
#else //!WIN32
#define DllExport
#endif //WIN32
#endif //DllExport

typedef void (*FrpcClosedCallback)(const char* msg);
FrpcClosedCallback frpcClosedCallback;

typedef void (*ProxyFailedCallback)();
ProxyFailedCallback proxyFailedCallback;

DllExport void setFrpcClosedCallback(FrpcClosedCallback l) {
	frpcClosedCallback = l;
}

DllExport void setProxyFailedCallback(ProxyFailedCallback l) {
	proxyFailedCallback = l;
}

void callbackFrpcClosed(const char* msg) {
	if (frpcClosedCallback) {
        frpcClosedCallback(msg);
	}
}

void callbackProxyFailed() {
	if (proxyFailedCallback) {
        proxyFailedCallback();
	}
}

*/
import "C"
import (
)

type LibStateListeners struct {
	l                   FRPCClosedListener
	proxyFailedListener ProxyFailedListener
}

type FRPCClosedListener interface {
	OnClosed(msg string)
}

type ProxyFailedListener interface {
	OnProxyFailed()
}

func (stListeners *LibStateListeners) OnClosed(msg string) {
	C.callbackFrpcClosed(C.CString(msg))
}
func (stListeners *LibStateListeners) OnProxyFailed(err error) {
	C.callbackProxyFailed()
}
