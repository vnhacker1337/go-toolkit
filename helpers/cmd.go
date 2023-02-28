package cmdutil

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func ExecutionWithStd(cmd string) (string, error) {
	command := []string{
		"bash",
		"-c",
		cmd,
	}
	var output string
	realCmd := exec.Command(command[0], command[1:]...)
	// output command output to std too
	cmdReader, _ := realCmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	var out string
	go func() {
		for scanner.Scan() {
			out += scanner.Text()
			//fmt.Fprintf(os.Stderr, scanner.Text()+"\n")
			fmt.Println(scanner.Text())
		}
	}()
	if err := realCmd.Start(); err != nil {
		return "", err
	}
	if err := realCmd.Wait(); err != nil {
		return "", err
	}
	return output, nil
}

func RuncmdWithLabel(cmd string, label string) (string, error) {
	command := []string{
		"bash",
		"-c",
		cmd,
	}
	fmt.Println("Start run Command: " + label)
	var output string
	realCmd := exec.Command(command[0], command[1:]...)
	// output command output to std too

	output = label
	fmt.Println("Completed Run Command " + label)
	if err := realCmd.Start(); err != nil {
		return "", err
	}
	if err := realCmd.Wait(); err != nil {
		return "", err
	}
	return output, nil
}

func InputCmd(Cmd string) string {
	command := []string{
		"bash",
		"-c",
		Cmd,
	}
	out, _ := exec.Command(command[0], command[1:]...).CombinedOutput()
	return strings.TrimSpace(string(out))
}
