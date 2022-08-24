package main

import (
	"encoding/json"
	"flag"
	"os/exec"
	"strconv"
)

func main() {
	flag.Parse()
	hyprcommand := flag.Arg(0)
	targetWorkspace, _ := strconv.Atoi(flag.Arg(1))

	//run cmd && get output
	focusMonitor := getActiveMonitor()

	activeWindow := getActiveWindow()
	if activeWindow.PID != 0 {
		//no active window
		focusMonitor = int(getActiveWindow().Monitor)
	}

	finalWorkspace := targetWorkspace + 10*focusMonitor

	exec.Command("hyprctl", "dispatch", hyprcommand, strconv.Itoa(finalWorkspace)).Output()
}

func getActiveMonitor() int {
	out, err := exec.Command("hyprctl", "monitors", "-j").Output()
	if err != nil {
		panic(err)
	}
	var monitors Monitors
	//print output
	json.Unmarshal(out, &monitors)

	focusMonitor := 0

	for i, monitor := range monitors {
		if monitor.Focused {
			focusMonitor = i
			break
		}
	}

	return focusMonitor
}

func getActiveWindow() ActiveWindow {
	//run cmd && get output
	out, err := exec.Command("hyprctl", "activewindow", "-j").Output()
	if err != nil {
		panic(err)
	}
	var activeWindow ActiveWindow
	//print output
	json.Unmarshal(out, &activeWindow)

	return activeWindow
}

// Generated by https://quicktype.io

type ActiveWindow struct {
	Address   string    `json:"address"`
	At        []int64   `json:"at"`
	Size      []int64   `json:"size"`
	Workspace Workspace `json:"workspace"`
	Floating  bool      `json:"floating"`
	Monitor   int64     `json:"monitor"`
	Class     string    `json:"class"`
	Title     string    `json:"title"`
	PID       int64     `json:"pid"`
	Xwayland  bool      `json:"xwayland"`
}

type Workspace struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Generated by https://quicktype.io

type Monitors []Monitor

type Monitor struct {
	ID              int64           `json:"id"`
	Name            string          `json:"name"`
	Width           int64           `json:"width"`
	Height          int64           `json:"height"`
	RefreshRate     float64         `json:"refreshRate"`
	X               int64           `json:"x"`
	Y               int64           `json:"y"`
	ActiveWorkspace ActiveWorkspace `json:"activeWorkspace"`
	Reserved        []int64         `json:"reserved"`
	Scale           float64         `json:"scale"`
	Transform       int64           `json:"transform"`
	Focused         bool            `json:"focused"`
}

type ActiveWorkspace struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
