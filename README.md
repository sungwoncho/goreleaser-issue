# Goreleaser CGO issue reproduction

This repository reproduces a possible issue with goreleaser building packages with cgo.

## Reproduce

* Run `go build --tags "darwin"` (This works)

* Run `goreleaser --snapshot --rm-dist`

Fails with

```
...
 38       • building                  binary=dist/darwin_amd64
 37       • added new artifact        name= path=dist/darwin_386 type=Binary
 36       • added new artifact        name= path=dist/darwin_amd64 type=Binary
 35    ⨯ release failed after 2.14s error=failed to build for linux_386: # runtime/internal/sys
 34 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:8:7: GOOS redeclared in this block
 33         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:7:14
 32 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:10:7: GoosAndroid redeclared in this block
 31         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:9:21
 30 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:11:7: GoosDarwin redeclared in this block
 29         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:10:20
 28 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:12:7: GoosDragonfly redeclared in this block
 27         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:11:23
 26 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:13:7: GoosFreebsd redeclared in this block
 25         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:12:21
 24 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:14:7: GoosJs redeclared in this block
 23         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:13:16
 22 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:15:7: GoosLinux redeclared in this block
 21         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:14:19
 20 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:16:7: GoosNacl redeclared in this block
 19         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:15:18
 18 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:17:7: GoosNetbsd redeclared in this block
 17         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:16:20
 16 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:18:7: GoosOpenbsd redeclared in this block
 15         previous declaration at /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_darwin.go:17:21
 14 /usr/local/Cellar/go/1.11/libexec/src/runtime/internal/sys/zgoos_linux.go:18:7: too many errors
 13 # runtime/cgo
 12 duplicate symbol __cgo_sys_thread_start in:
 11     $WORK/b034/_x004.o
 10     $WORK/b034/_x007.o
  9 duplicate symbol _x_cgo_init in:
  8     $WORK/b034/_x004.o
  7     $WORK/b034/_x007.o
  6 ld: 2 duplicate symbols for architecture i386
  5 clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

