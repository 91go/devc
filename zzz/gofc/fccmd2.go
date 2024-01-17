package gofc

import (
	"bytes"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

// command（用golang执行命令行）

/** 类似于python的commands库
 * - Run(command string, args...string) (void)
 * - GetOutput(command string, args...string) (output string)
 * - GetStatusOutput(command string, args...string) (status bool, output string)
 **/

// 执行连续的shell命令时, 需要注意指定执行路径和参数, 否则运行出错
// 比如 /bin/sh -c multi-commands-string

// Commands define the structure
type Commands struct {
	Status bool
	Output string
	Error  error
}

// New init instance with properties stands for the status and the output
func New() *Commands {
	var command Commands
	command = Commands{
		Status: false,
		Output: "",
		Error:  nil,
	}
	return &command
}

// Run 执行命令，获取输出
func (c *Commands) Run(command string, args ...string) (output string) {
	var shell, flag string
	if runtime.GOOS == "windows" {
		shell = "cmd"
		flag = "/C"
	} else {
		shell = "/bin/sh"
		flag = "-c"
	}
	args = append([]string{flag, command}, args...)

	out, err := exec.Command(shell, args...).Output()
	if err != nil {
		c.Error = err
		log.Fatal(err)
	} else {
		c.Output = strings.Trim(string(out), "\n")
	}
	return c.Output
}

// RunBlocking 阻塞式执行
func (c *Commands) RunBlocking(command string, args ...string) (output string, err error) {
	var shell, flag string
	if runtime.GOOS == "windows" {
		shell = "cmd"
		flag = "/C"
	} else {
		shell = "/bin/sh"
		flag = "-c"
	}
	args = append([]string{flag, command}, args...)

	// 函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(shell, args...)

	// 读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型out.String()
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err = cmd.Run()

	return out.String(), err
}

// OuterRun 不使用任何shell执行命令
// cat /etc/shells
// /bin/bash
// /bin/csh
// /bin/dash
// /bin/ksh
// /bin/sh
// /bin/tcsh
// /bin/zsh
func (c *Commands) OuterRun(command string, args ...string) (output string) {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		c.Error = err
		log.Fatal(err)
	} else {
		c.Output = strings.Trim(string(out), "\n")
	}
	return c.Output
}

// RunWithoutOutput 执行命令，不获取输出
// commands without any output
func (c *Commands) RunWithoutOutput(command string, args ...string) {
	var shell, flag string
	if runtime.GOOS == "windows" {
		shell = "cmd"
		flag = "/C"
	} else {
		shell = "/bin/sh"
		flag = "-c"
	}
	args = append([]string{flag, command}, args...)
	cmd := exec.Command(shell, args...)

	if err := cmd.Run(); err != nil {
		c.Error = err
		log.Fatal(err)
	}
}

// GetStatusOutput  获取status和具体输出
// run commands with status and output
func (c *Commands) GetStatusOutput(command string, args ...string) (status bool, output string) {
	args = append([]string{"-c", command}, args...)
	cmd := exec.Command("/bin/sh", args...)
	var bufferout, bufferin, buffererr bytes.Buffer
	cmd.Stdout = &bufferout
	cmd.Stdin = &bufferin
	cmd.Stderr = &buffererr

	// do the commands
	err := cmd.Run()
	if err != nil {
		c.Output = err.Error()
		c.Error = err
	} else {
		c.Status = cmd.ProcessState.Success()
		c.Output = bufferout.String()
	}
	return c.Status, c.Output
}

// GetError just for debugging
func (c *Commands) GetError() error {
	return c.Error
}
