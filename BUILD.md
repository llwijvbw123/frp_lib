FRPS server app:
$ make frps

FRPC library:
windows x64:

Install mingw64 and add "bin" to PATH

$ mingw32-make frpc-lib-windows

windows x32:

Install mingw (32bit) and add "bin" to PATH

$ export GOARCH=386
$ mingw32-make frpc-lib-windows

macos:

$ make frpc-lib-unix

android (on windows x64):

$ mingw32-make frpc-lib-android

ios (on macos):

$ make frpc-lib-ios