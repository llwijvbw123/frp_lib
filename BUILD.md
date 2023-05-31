# FRPS server app (not tested, i build frps from original fatedier's repo):
$ make frps

# FRPC library:
## windows x64:

Install mingw64 and add "bin" to PATH

$ mingw32-make frpc-lib-windows64

## windows x32:

Install mingw (32bit) and add "bin" to PATH

$ mingw32-make frpc-lib-windows32

## macos:

$ make frpc-lib-unix

## android (on windows x64):

$ mingw32-make frpc-lib-android

## ios (on macos):

$ make frpc-lib-ios

# Header fix
In desktop header need to move/replace
```c
extern DllExport void setLogListener(LogListener l);
```
to bottom 'extern "C"' block 
```c
#ifdef __cplusplus
extern "C" {
#endif
...
extern __declspec(dllexport) void setLogListener(LogListener l);
#ifdef __cplusplus
}
#endif
```

And for working in VS2017 need to replace:
```c
#ifdef _MSC_VER
#include <complex.h>
typedef _Fcomplex GoComplex64;
typedef _Dcomplex GoComplex128;
#else
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;
#endif
```
to
```c
#ifdef _MSC_VER
#  if _MSVC_LANG <= 201402L
#    include <complex.h>
typedef _Fcomplex GoComplex64;
typedef _Dcomplex GoComplex128;
#  else
#    include <complex>
typedef std::complex<float> GoComplex64;
typedef std::complex<double> GoComplex128;
#  endif
#else
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;
#endif
```
