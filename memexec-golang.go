package main

// Example: go build memexec-golang.go && cat /bin/id | ./memexec-golang

import (
    "io"
    "fmt"
    "os"
    "unsafe"
    "syscall"
)

func MemfdCreate(name string) (fd uintptr) {
	pn, err := syscall.BytePtrFromString(name)
	if err != nil {
		return uintptr(0)
	}
	fd, _, e := syscall.Syscall(319, uintptr(unsafe.Pointer(pn)), uintptr(0), uintptr(0))
	if e != 0 {
		return uintptr(0)
	}
	return fd
}

func Execveat(fd int, path string) {
	var env []byte

	argvp, err := syscall.SlicePtrFromStrings(os.Args)
	if err != nil {
		return
	}

	pn, err := syscall.BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, _ = syscall.Syscall6(
		322, 
		uintptr(fd), 
		uintptr(unsafe.Pointer(pn)),
		uintptr(unsafe.Pointer(&argvp[0])),
		uintptr(unsafe.Pointer(&env)),
		uintptr(0x1000),
		uintptr(0),
		)
}

func main(){
	fd := int(MemfdCreate(""))

	if (fd == 0) {
		fmt.Fprintln(os.Stderr, "memfd failed")
		return
	}

	data, err := io.ReadAll(os.Stdin)
	_, err = syscall.Write(fd, data)
	if err != nil {
		return
	}

	Execveat(fd, "")
}

