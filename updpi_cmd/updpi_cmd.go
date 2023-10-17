package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	Ghost   string
	Gport   string
	Gsw     string
	Gcmd    string
	Gfile   string
	Gfilter string
	Command string
	Level   string
)

type CmdInfo struct {
	Cmd    string
	Level  string
	Filter string
	Flag   bool //  true: sw|cmd  false: sw
}

var cmdinfo CmdInfo

func ParseArgs() bool {
	if Gfile != "" {
		return true
	}

	if Gsw != "" && Gcmd != "" {
		return true
	}

	return false
}

func ProcSingleCmd(writer *bufio.Writer, reader *bufio.Reader) {
	// 切换视图
	command := "switch " + Gsw
	fmt.Fprintln(writer, command)
	writer.Flush()

	// 读取命令输出
	_, err := reader.ReadString('>')
	if err != nil {
		fmt.Println("Error reading command output:", err)
		return
	}

	// 发送命令
	fmt.Fprintln(writer, Gcmd)
	writer.Flush()

	// 读取命令输出
	commandOutput, err := reader.ReadString('>')
	if err != nil {
		fmt.Println("Error reading command output:", err)
		return
	}
	if Gfilter == "" {
		fmt.Print(commandOutput)
		return
	}

	fs := strings.Split(commandOutput, "\n")
	for i := 0; i < len(fs); i++ {
		if strings.Contains(fs[i], Gfilter) {
			fmt.Print(fs[i] + "\n")
		}
	}

	return
}

func ProcFileCmd(writer *bufio.Writer, reader *bufio.Reader) {
	fd, err := os.Open(Gfile)
	if err != nil {
		fmt.Print("open file '%s' failed:%v\n", Gfile, err)
		return
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") {
			//视图切换
			line := strings.Trim(line, "[]")
			fs := strings.Split(line, "|")
			if len(fs) == 3 {
				cmdinfo.Level = fs[0]
				cmdinfo.Cmd = fs[1]
				cmdinfo.Filter = fs[2]
				cmdinfo.Flag = true

			}
			if len(fs) == 2 {
				cmdinfo.Level = fs[0]
				cmdinfo.Cmd = fs[1]
				cmdinfo.Flag = true
			} else {
				cmdinfo.Level = fs[0]
				cmdinfo.Cmd = ""
				cmdinfo.Flag = false
			}

			// 切换视图
			cmd := "switch " + cmdinfo.Level
			fmt.Fprintln(writer, cmd)
			writer.Flush()

			// 读取命令输出
			_, err := reader.ReadString('>')
			if err != nil {
				fmt.Println("Error reading command output:", err)
				return
			}

			continue
		} else {
			//命令行查询
			var cmd string
			if cmdinfo.Flag {
				cmd = cmdinfo.Cmd + " " + line
			} else {
				cmd = line
			}

			// 发送命令
			fmt.Fprintln(writer, cmd)
			writer.Flush()

			// 读取命令输出
			commandOutput, err := reader.ReadString('>')
			if err != nil {
				fmt.Println("Error reading command output:", err)
				return
			}
			if cmdinfo.Filter == "" {
				fmt.Print(commandOutput)
				continue
			}

			fs := strings.Split(commandOutput, "\n")
			for i := 0; i < len(fs); i++ {
				if strings.Contains(fs[i], cmdinfo.Filter) {
					fmt.Print(fs[i] + "\n")
				}
			}
			continue
		}
	}

	return
}

func main() {
	// 设置telnet主机和端口
	flag.StringVar(&Ghost, "host", "127.0.0.1", "updpi host ip")
	flag.StringVar(&Gport, "port", "36500", "updpi cmd port")
	flag.StringVar(&Gsw, "sw", "", "switch level")
	flag.StringVar(&Gcmd, "cmd", "", "the command string")
	flag.StringVar(&Gfile, "file", "", "the command file")
	flag.StringVar(&Gfilter, "filter", "", "过滤关键字")
	flag.Parse()

	valid := ParseArgs()
	if !valid {
		flag.Usage()
		os.Exit(-1)
	}

	// 创建telnet连接
	conn, err := net.Dial("tcp", Ghost+":"+Gport)
	if err != nil {
		fmt.Println("Error connecting to telnet server:", err)
		return
	}
	defer conn.Close()

	// 创建读取器和写入器
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// 读取命令提示符
	_, err = reader.ReadString('>')
	if err != nil {
		fmt.Println("Error reading command prompt:", err)
		return
	}

	//操作命令
	if Gfile == "" {
		ProcSingleCmd(writer, reader)
	} else {
		ProcFileCmd(writer, reader)
	}

	// 发送退出命令
	fmt.Fprintln(writer, "exit")
	writer.Flush()

	// 等待telnet会话结束
	conn.Close()

	//fmt.Println("Telnet session ended")
	return
}
