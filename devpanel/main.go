package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/ufs"
)

type AppProfile struct {
	Name             string `json:"name"`
	WorkingDirectory string `json:"workingDirectory"`
	Exec             string `json:"exec"`
}

func main() {

	/*
		appsDir, err := dir.New("apps")
		do.Die(err)

		appProfiles, err := appsDir.Files()
		do.Die(err)

		fmt.Println("Apps:")
		for i, appProfile := range appProfiles {
			fmt.Printf("%v. %v\n", i+1, appProfile)
		}

		fmt.Print("\n\nSelect app profile:\n>>")

		selectedProfileIndex := 0
		fmt.Scanf("%d", selectedProfileIndex)

		selectedProfilePath := appProfiles[selectedProfileIndex]
	*/

	selectedProfilePath := `C:\Users\user\Desktop\projects\1000 lines\go-plg\devpanel\resources\apps\adorngirlies-server`

	raw, err := ufs.ReadFileAsBytes(selectedProfilePath)
	do.Die(err)

	selectedProfile := AppProfile{}

	err = json.Unmarshal(raw, &selectedProfile)
	do.Die(err)

	var (
		name = ""
		args = []string{}
	)

	ls := strings.Split(selectedProfile.Exec, " ")
	if len(ls) > 0 {
		name = ls[0]
	}
	if len(ls) > 1 {
		args = ls[1:]
	}

	cmd := exec.Command(name, args...)
	if cmd == nil {
		do.Die(fmt.Errorf("something went wrong"))
		return
	}

	cmd.Dir = selectedProfile.WorkingDirectory
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	// Start the process in its own process group
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}

	err = cmd.Start()
	do.Die(err)

	cmdDone := make(chan error)
	go func() {
		err := cmd.Wait()
		cmdDone <- err
	}()

	cmdKill := make(chan bool)
	go func() {
		kill := ""
		fmt.Println("waiting for the kill key")
		fmt.Scanf("%s", &kill)
		cmdKill <- true
	}()

	select {
	case err := <-cmdDone:
		if err != nil {
			fmt.Printf("Command finished with error: %v\n", err)
		} else {
			fmt.Println("Command finished successfully")
		}
	case kill := <-cmdKill:
		terminateProcessTree(cmd.Process.Pid)
		fmt.Printf("Kill(%v) command received: %v\n", cmd.Process.Pid, kill)
	}
}

// Use taskkill to kill the process tree on Windows
func terminateProcessTree(pid int) {
	taskkillCmd := exec.Command("taskkill", "/PID", fmt.Sprintf("%d", pid), "/T", "/F")
	taskkillCmd.Stdout = os.Stdout
	taskkillCmd.Stderr = os.Stderr
	if err := taskkillCmd.Run(); err != nil {
		fmt.Printf("Failed to kill process tree: %v\n", err)
	} else {
		fmt.Println("Process tree killed")
	}
}
