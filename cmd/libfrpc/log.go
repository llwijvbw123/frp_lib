package main

/*
#include <stdio.h>

#ifndef DllExport
#ifdef WIN32
#define DllExport __declspec( dllexport )
#else //!WIN32
#define DllExport
#endif //WIN32
#endif //DllExport

typedef void (*LogListener)(const char* log);

LogListener logListener;

DllExport void setLogListener(LogListener l) {
	logListener = l;
}

void callback(const char* log) {
	if (logListener) {
        logListener(log);
	}
}

void cListener(const char* log) {
	printf("%s", log);
}

void testLog() {
	setLogListener(cListener);
}
*/
import "C"
import (
	"time"

	"frp_lib/pkg/util/log"
)

var l logForLibListener

type logForLibListener struct {
	log.LogListener
}

func (l *logForLibListener) Log(log string) {
	C.callback(C.CString(log))
}
func (l *logForLibListener) Location() string {
	location, _ := time.LoadLocation("Local")
	return location.String()
}

func init() {
	l = logForLibListener{}
	log.AppendListener(&l)
}

func logLog() {
	C.testLog()
	println(C.logListener)
}
