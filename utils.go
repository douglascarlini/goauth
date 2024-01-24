package main

import (
    "runtime"
    "os/exec"
    "time"
)

func Open(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
        args = []string{"/c", "start"}
    case "darwin":
        cmd = "open"
    default:
        cmd = "xdg-open"
    }

    args = append(args, url)
	time.Sleep(2 * time.Second)
    return exec.Command(cmd, args...).Start()
}