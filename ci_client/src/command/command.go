package command

import (
	"bufio"
	"ci_client/common"
	"ci_client/log"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// 状态
type CmdState int

const (
	INIT CmdState = 0
	RUN  CmdState = 1
	OVER CmdState = 2
	KILL CmdState = 3
)

// 命令
type Command struct {
	state   CmdState // 0：初始化，1：开始，2：结束，3：kill
	pid     int      // pid
	logFile string   // 日志文件地址
	params  []string
	cmdName string
}

func New(arg ...string) *Command {
	// todo 日志地址
	script := common.GetScript()
	l := len(arg)
	args := make([]string, l+1)
	args[0] = script
	for i := 1; i <= l; i++ {
		args[i] = arg[i-1]
	}
	return &Command{state: INIT, pid: -1, logFile: common.GetScriptLog(), params: args, cmdName: "python"}
}

// 关闭进程
func (cmd *Command) Kill() {
	if cmd.state != RUN {
		log.Info.Println("不在执行状态... 无法kill")
	}
	sPid := strconv.Itoa(cmd.pid)
	command := exec.Command("taskkill", "-PID", sPid, "-F")
	//command := exec.Command("kill", "-9", sPid)
	_ = command.Wait()
	cmd.state = KILL
	log.Info.Println("kill succ...")
}

// 开始任务
func (cmd *Command) Start() {
	if cmd.state != INIT {
		log.Info.Println("不是初始状态... 无法start")
	}
	log.Info.Printf("exce %s  params[0] %s params[1] %s params[2] %s",
		cmd.cmdName, cmd.params[0], cmd.params[1], cmd.params[2])
	execCommand(cmd)
}

//封装一个函数来执行命令
func execCommand(cmd *Command) bool {

	//执行命令
	command := exec.Command(cmd.cmdName, cmd.params...)
	//command := exec.Command(cmd.cmdName, cmd.params...)
	//显示运行的命令
	log.Info.Println(command.Args)

	stdout, err := command.StdoutPipe()
	errReader, errr := command.StderrPipe()

	if errr != nil {
		fmt.Println("err:" + errr.Error())
	}

	//开启错误处理
	go handlerErr(errReader)

	if err != nil {
		fmt.Println(err)
		return false
	}
	startErr := command.Start()
	if startErr != nil {
		log.Error.Println("start error", startErr)
	}
	cmd.state = RUN
	cmd.pid = command.Process.Pid
	go cmd.check(command)
	//go cmd.follow()
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		cmdRe := string(in.Bytes())
		log.Info.Print(cmdRe)
	}

	return true
}

//开启一个协程来错误
func handlerErr(errReader io.ReadCloser) {
	in := bufio.NewScanner(errReader)
	for in.Scan() {
		log.Error.Print(string(in.Bytes()))
	}
}

// 等待完成
func (cmd *Command) check(command *exec.Cmd) {
	_ = command.Wait()
	if command.ProcessState.Success() {
		// 执行成功
		cmd.state = OVER
	}
}

// 读取日志文件
func (cmd *Command) follow() {
	file, err := os.Open(cmd.logFile)
	if err != nil {
		fmt.Println("open file err", err)
	}

	defer file.Close()
	r := bufio.NewReader(file)
	times := 0
	for {
		if times > 100 {
			break
		}
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Error.Println("读取日志失败...", err)
			time.Sleep(time.Second)
			times++
		}
		if len(line) > 0 {
			// TODO 发送日志
			fmt.Println(line)
		}

		if err == io.EOF && cmd.state != RUN {
			// TODO 发送日志结束
			break
		} else if err == io.EOF {
			time.Sleep(time.Second)
		}
	}
}

// get pid
func (cmd *Command) GetPid() int {
	return cmd.pid
}
