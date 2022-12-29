package helper

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)


func RunShell(shellPath, logPath string) {
	
	// 分配权限 0777
	cmdChmod := exec.Command("sh", "-c", "chmod +x " + shellPath)
	var outChmod, errChmod bytes.Buffer
	cmdChmod.Stdout = &outChmod
	cmdChmod.Stderr = &errChmod
	if err := cmdChmod.Run(); err != nil {
		log.Fatal("[CHMOD Error]:#" + err.Error())
	}

	// 打印当前时间 & 追加日志
	fd, err := os.OpenFile(logPath, os.O_WRONLY | os.O_APPEND, 0666)
	if errors.Is(err , os.ErrNotExist) {
		os.MkdirAll(path.Dir(logPath), 0777)
		fd, err = os.Create(logPath)
		if err != nil {
			log.Fatal("[CREATE Error]:#" + err.Error())
		}
	}

	w := bufio.NewWriter(fd)
	w.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\n")
	w.Flush()


	// 运行 & 追加日志
	cmdShell := exec.Command("sh", "-c", "nohub " + shellPath + " >> " + " 2>&1 & ")  // 后台
	var outShell, errShell bytes.Buffer
	cmdShell.Stdout = &outShell
	cmdShell.Stderr = &errShell
	if err := cmdShell.Run(); err != nil {
		log.Fatal("[SHELL Error]:#" + err.Error())
	}
}

