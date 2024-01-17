package gofc

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// 执行git操作

// IsGitRepository 判断是否已经git初始化
func IsGitRepository(path string) error {
	output, err := ShellExec("git rev-parse --is-inside-work-tree", path)
	if err != nil {
		return fmt.Errorf("%s is not git repository: %s", path, err.Error())
	}
	if !strings.Contains(output, "true") {
		return fmt.Errorf("%s is not git repository: check %s", path, output)
	}

	return nil
}

// GitPull 执行pull操作
func GitPull(path string) error {
	// pull code
	output, err := ShellExec("git pull", path)
	if err != nil {
		return fmt.Errorf("git pull fail: %s", err.Error())
	}
	if strings.Contains(output, "Already") {
		return errors.New("git Already")
	}

	return nil
}

// GetLastCommitID 获取最后一次提交的commit-id
func GetLastCommitID(path string) (string, error) {
	// pull code
	output, err := ShellExec(`git log --pretty=format:"%H" -n 1 2>&1`, path)
	if err != nil {
		return "", fmt.Errorf("get commit id: %s", err.Error())
	}

	reg := regexp.MustCompile(`^[0-9a-f]{40}$`)
	matches := reg.FindStringSubmatch(strings.Trim(string(output), "\""))
	if len(matches) == 0 {
		return "", errors.New("commit id len not match")
	}

	return matches[0], nil
}

// GitReset 强制回退到指定版本
// 参数为commit-id
func GitReset(hash, path string) error {

	_, err := ShellExec("git reset --hard "+hash, path)

	return err
}

// RunGitCommand 执行任意Git命令的封装
func RunGitCommand(name string, arg ...string) (string, error) {
	// 从配置文件中获取当前git仓库的路径
	gitPath, _ := os.Getwd()

	cmd := exec.Command(name, arg...)
	// 指定工作目录为git仓库目录
	cmd.Dir = gitPath
	// cmd.Stderr = os.Stderr
	// 混合输出stdout+stderr
	msg, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	// 报错时 exit status 1
	return string(msg), err
}

func ShellExec(cmd string, args ...string) (string, error) {
	cm := exec.Command(cmd, args...)
	var out bytes.Buffer
	cm.Stdout = &out

	err := cm.Run()
	if err != nil {
		return "", err
	}

	result := out.String()
	return result, nil
}
