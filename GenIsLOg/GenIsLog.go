package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Bytes []byte

const (
	CTCC_TYPE = 0 //电信
	CUCC_TYPE = 1 //联通
)

var (
	comId  string
	sTime  string
	eTime  string
	logNum int
	sPath  string
	dPath  string
	sPort  int
	debug  bool
)

type logRecord struct {
	header   Bytes
	sep      Bytes
	content  Bytes
	filename string
	newfile  string
	mode     int
	command  string
	valid    bool
}

var originRecord logRecord
var portMap sync.Map

var ErrStopIteration = errors.New("find record && stop iteration")

func isFileTimeValid(filename string) bool {
	var timeStr string
	fs := strings.Split(filename, "_")
	//按照文件格式取第4个字段，代表文件时间
	if len(fs) >= 4 {
		timeStr = fs[3]
		//文件时间在给定的时间段内
		if timeStr > sTime && timeStr < eTime {
			if debug {
				fmt.Printf(">>>Match File:%s\n", filename)
			}
			return true
		}
	}

	return false
}

func procIsLineContent(line []byte) bool {
	mode := Bytes{0x0, 0x0}
	//command Id分两种方式
	if bytes.Compare(line[12:14], mode) == 0 {
		//command id按照数值进行比较
		commandA := binary.BigEndian.Uint64(line[14:22])
		commandB, err := strconv.ParseUint(comId, 10, 64)
		if err != nil {
			fmt.Printf("convert param command id to int failed:%v\n", err)
			return false
		}
		if debug {
			fmt.Printf("command mode: uint64, comman param:%s, command search:%v, record:%v\n", comId, commandA, line)
		}
		if commandA == commandB {
			originRecord.content = line
			originRecord.command = comId
			originRecord.valid = true
		}
	} else {
		fmt.Printf("[%v][%d]\n", line, len(line))
		commandA := string(bytes.Trim(line[12:22], string(0)))
		if debug {
			fmt.Printf("command mode: str, comman param:%s(%d), command search:%s(%d), record:%v\n", comId, len(comId), commandA, len(commandA), line)
		}
		fmt.Printf("[%v][%d]\n", line, len(line))
		if commandA == comId {
			originRecord.content = line
			originRecord.command = comId
			originRecord.valid = true
		}
	}

	return originRecord.valid
}

func procIsFileContent(data, sep []byte) bool {
	flag := false
	ds := bytes.Split(data, sep)
	if len(ds) < 2 {
		fmt.Printf("no log record.\n")
	}

	if debug {
		fmt.Printf(">>>Match Format, record num:%d\n", len(ds))
	}

	//var header Bytes
	for i, v := range ds {
		//fmt.Printf("[index:%d]==>%v\n", i, v)
		if i == 0 {
			originRecord.header = v
		} else {
			flag = procIsLineContent(v)
		}
	}

	return flag
}

func parseIsLogFile(fn string) bool {
	flag := false
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return flag
	}

	//0x03 + "X1D"
	sep_cucc := Bytes{0x03, 0x58, 0x31, 0x44}
	sep_ctcc := Bytes{0x10, 0x43, 0x55, 0x44}
	index_cucc := bytes.Index(data, sep_cucc)
	index_ctcc := bytes.Index(data, sep_ctcc)

	if index_cucc > 0 {
		//联通
		originRecord.filename = fn
		originRecord.sep = sep_cucc
		originRecord.mode = CUCC_TYPE
		flag = procIsFileContent(data, sep_cucc)
	} else if index_ctcc > 0 {
		//电信
		originRecord.filename = fn
		originRecord.sep = sep_ctcc
		originRecord.mode = CTCC_TYPE
		flag = procIsFileContent(data, sep_ctcc)
	}

	return flag
}

func getCommandRecord1() {
	dir, err := ioutil.ReadDir(sPath)
	if err != nil {
		return
	}

	for _, fi := range dir {
		if fi.IsDir() { //忽略目录
			continue
		}

		fn := fi.Name()
		isPeriod := isFileTimeValid(fn)
		if !isPeriod {
			continue
		} else {
			// read file content && find command id && return
			parseIsLogFile(fn)
		}
	}

	return
}

func getCommandRecord() {
	flag := false
	err := filepath.Walk(sPath, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("filepath walk failed:%v\n", err)
			return err
		}

		if !fi.IsDir() {
			// 嵌套路径需要考虑，使用path
			isPeriod := isFileTimeValid(path)
			if !isPeriod {
				return nil
			} else {
				// read file content && find command id && return
				flag = parseIsLogFile(path)
				if flag {
					//找到符合条件的记录，不继续遍历
					return ErrStopIteration
				}
			}
		}

		return nil

	})

	if err != nil && debug {
		fmt.Printf("filepath walk info:%v\n", err)
	}
	return
}

func getHexData(data Bytes) string {
	var str string
	for _, v := range data {
		str += fmt.Sprintf("%02x ", v)
	}

	return str
}

func printOriginCommandRecord() {
	splitStr := strings.Repeat("=", 64)
	fmt.Println(splitStr)
	fmt.Printf("filename: %s\n", originRecord.filename)
	fmt.Printf("newfile : %s\n", originRecord.newfile)
	if originRecord.mode == CTCC_TYPE {
		fmt.Printf("oprator: ctcc\n")
	} else {
		fmt.Printf("oprator: cucc\n")
	}
	fmt.Printf("command: [%s]\n", originRecord.command)
	fmt.Printf("header: [%s]\n", getHexData(originRecord.header))
	fmt.Printf("sep: [%s]\n", getHexData(originRecord.sep))
	fmt.Printf("content: [%s]\n", getHexData(originRecord.content))
	fmt.Println(splitStr)
}

func uint16ToBytes(n uint16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)

	return bytesBuffer.Bytes()
}

func uint32ToBytes(n uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)

	return bytesBuffer.Bytes()
}

func uint64ToBytes(n uint64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)

	return bytesBuffer.Bytes()
}

func genRandomPort() []byte {
	min := 10000
	max := 55555

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNum := r.Intn(max-min+1) + min

	return uint16ToBytes(uint16(randomNum))
}

func genRandomPortInt() uint16 {
	min := 10000
	max := 55555

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNum := r.Intn(max-min+1) + min

	return uint16(randomNum)
}

func getOnlyRandomPort() []byte {
	var port uint16
	for i := 0; i < 10; i++ {
		port = genRandomPortInt()
		if _, ok := portMap.Load(port); !ok {
			portMap.Store(port, 1)
			return uint16ToBytes(port)
		}
	}
	fmt.Printf("generate same random port over 10 times, use the last port:%d!!\n", port)
	return uint16ToBytes(port)
}

var portIndex uint16

func getIncreasePort() []byte {
	gPort := uint16(sPort) + portIndex
	portIndex++

	return uint16ToBytes(gPort)
}

func genRandomTime() []byte {
	timelayout := "20060102150405"

	//加载CST时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load CST location failed:%v\n", err)
	}

	st, err := time.ParseInLocation(timelayout, sTime, location)
	if err != nil {
		fmt.Printf("parse sTime failed:%v\n", err)
	}
	et, err := time.ParseInLocation(timelayout, eTime, location)
	if err != nil {
		fmt.Printf("parse sTime failed:%v\n", err)
	}

	min := st.Unix()
	max := et.Unix()

	/*
		min, err := strconv.ParseUint(sTime, 10, 64)
		if err != nil {
			fmt.Printf("convert stime err:%v\n", err)
		}
		max, err := strconv.ParseUint(eTime, 10, 64)
		if err != nil {
			fmt.Printf("convert etime err:%v\n", err)
		}
	*/

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNum := uint32(r.Intn(int(max-min+1))) + uint32(min)

	if debug {
		fmt.Printf("stime:%s(%d), etime:%s(%d), random:%d\n", sTime, min, eTime, max, randomNum)
	}
	return uint32ToBytes(randomNum)
}

func genRandomTimeStr() string {
	min, err := strconv.ParseUint(sTime, 10, 64)
	if err != nil {
		fmt.Printf("convert stime err:%v\n", err)
	}
	max, err := strconv.ParseUint(eTime, 10, 64)
	if err != nil {
		fmt.Printf("convert etime err:%v\n", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomNum := uint64(r.Intn(int(max-min+1))) + min

	if debug {
		fmt.Printf("stime:%s(%d), etime:%s(%d), random:%d\n", sTime, min, eTime, max, randomNum)
	}

	timeStr := fmt.Sprintf("%d", randomNum)

	return timeStr
}

func calcPortOffset() int {
	// modify port && timestamp
	// content
	// 4 device-id
	// 4 "0x01, 0xe0, 0x0, 0x0"
	// 4 pkt_len
	// 10 cmd_id
	// 1    house_id_len(x)
	// x house_id
	// 8 sip+dip || 32 sip+dip
	// 2 sport

	content := originRecord.content

	//calc offset of port
	// init for offset
	offset := 22
	houseIdLen := uint8(content[22])
	offset = offset + 1 + int(houseIdLen)

	ipLen := uint8(content[offset])
	if ipLen == 4 {
		offset = offset + 8 + 2
	} else if ipLen == 16 {
		offset = offset + 32 + 2
	} else {
		fmt.Printf("Parse ip length err, the offset is:%d, the len is:%d\n", offset, ipLen)
		return 0
	}

	return offset
}

func calcTimeOffset() int {
	// modify port && timestamp
	// content
	// 4 device-id
	// 4 "0x01, 0xe0, 0x0, 0x0"
	// 4 pkt_len
	// 10 cmd_id
	// 1    house_id_len(x)
	// x house_id
	// 8 sip+dip || 32 sip+dip
	// 4 sport
	// 2 host_len
	// y host
	// 8 proxy + title + content
	// 2 url_len
	// m url
	// 1 attach default == 1
	// 2 filename_len
	// n filename(n)
	// 4 time

	content := originRecord.content

	//calc offset of time

	//init for offset
	offset := 22

	//house id
	houseIdLen := uint8(content[22])
	offset = offset + 1 + int(houseIdLen)

	//ip
	ipLen := uint8(content[offset])
	if ipLen == 4 {
		offset = offset + 8 + 2
	} else if ipLen == 16 {
		offset = offset + 32 + 2
	} else {
		fmt.Printf("Parse ip length err, the offset is:%d, the len is:%d\n", offset, ipLen)
		return 0
	}

	//port
	offset += 4

	//host
	hostLen := binary.BigEndian.Uint16(content[offset : offset+2])
	offset = offset + 2 + int(hostLen)

	//proxy title content
	offset = offset + 8

	//url
	urlLen := binary.BigEndian.Uint16(content[offset : offset+2])
	offset = offset + 2 + int(urlLen)

	//attatch
	attchLen := uint8(content[offset])
	offset += 1
	if attchLen == 1 {
		filenameLen := binary.BigEndian.Uint16(content[offset : offset+2])
		offset = offset + 2 + int(filenameLen)
	} else if attchLen != 0 {
		fmt.Printf("Parse attach len err, the offset is:%d, the len is:%d\n", offset, attchLen)
	}

	return offset
}

func genIsRecord(offset1, offset2 int) []byte {
	var port []byte
	//sport生成的两种方式
	if sPort != 0 {
		//以设置的起始port开始递增
		port = getIncreasePort()
	} else {
		//在指定范围内开始随机去重
		port = getOnlyRandomPort()
	}
	time := genRandomTime()
	content := originRecord.content

	// replace port
	if debug {
		fmt.Printf("offset:%d, port:%v\n", offset1, port)
		fmt.Printf("offset:%d, time:%v\n", offset2, time)
	}
	copy(content[offset1:offset1+2], port)

	// replace time
	copy(content[offset2:offset2+4], time)

	return content
}

func getNewFp() string {
	//新时间戳
	timeStr := genRandomTimeStr()
	//原始文件名
	oldfn := originRecord.filename

	var basename string
	index := strings.LastIndex(oldfn, "/")
	if index > -1 {
		basename = oldfn[index+1:]
	} else {
		basename = oldfn
	}

	var newfn string
	fs := strings.Split(basename, "_")
	if len(fs) < 4 {
		return ""
	}
	for i, v := range fs {
		if i == 0 {
			newfn = v
		} else if i == 3 {
			newfn = newfn + "_" + timeStr
		} else {
			newfn = newfn + "_" + v
		}
	}

	path := dPath + "/" + newfn
	return path
}

func genIsLog() string {
	portoff := calcPortOffset()
	timeoff := calcTimeOffset()

	fn := getNewFp()
	fd, err := os.Create(fn)
	if err != nil {
		fmt.Printf("Create file err:%v\n", err)
		return ""
	}
	defer fd.Close()

	//write header
	if debug {
		fmt.Printf("header: %02x\n", originRecord.header)
	}
	fd.Write(originRecord.header)
	for i := 0; i < logNum; i++ {
		//write sep
		if debug {
			fmt.Printf("sep: %02x\n", originRecord.sep)
		}
		fd.Write(originRecord.sep)
		ctx := genIsRecord(portoff, timeoff)
		if debug {
			fmt.Printf("sep: %02x\n", originRecord.sep)
			fmt.Printf("content: %02x\n", ctx)
		}
		//write content
		fd.Write(ctx)
	}

	originRecord.newfile = fn
	return fn
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Printf("path '%s' stat failed:%v\n", path, err)
		return false
	}

	if fi.IsDir() {
		return true
	}

	return false
}

func paramCheck() bool {
	if comId == "" {
		fmt.Printf("command id is nil or invalid.\n")
		return false
	}

	if sTime == "" || len(sTime) != 14 {
		fmt.Printf("start time is nil or invalid, the format is yyyymmddhhmmss.\n")
		return false
	}

	if eTime == "" || len(eTime) != 14 {
		fmt.Printf("end time is nil or invalid, the format is yyyymmddhhmmss.\n")
		return false
	}

	if sTime >= eTime {
		fmt.Printf("param invalid, stime is lager than etime.\n")
		return false

	}

	if logNum == 0 {
		fmt.Printf("logNum is nil or invalid.\n")
		return false
	}

	/*
		if isDir(sPath) {
			fmt.Printf("src path '%s' is not a dir.\n", sPath)
			return false
		}

		if isDir(dPath) {
			fmt.Printf("dst path '%s' is not a dir.\n", dPath)
			return false
		}
	*/

	//剔除最后路径下的分隔符
	sPath = strings.TrimSuffix(sPath, "/")
	dPath = strings.TrimSuffix(dPath, "/")

	return true
}

func genOkFile(fn string) {
	index := strings.LastIndex(fn, ".")
	if index < 0 {
		fmt.Printf("generate ok file failed, origin file is '%s'\n", fn)
		return
	}

	name := fn[:index] + ".ok"

	fd, err := os.Create(name)
	if err != nil {
		fmt.Printf("create ok file '%s' failed:%v\n", name, err)
		return
	}
	defer fd.Close()

	return
}

func main() {
	//命令行解析
	flag.StringVar(&comId, "comid", "", "set command id, use for find match log record in the spath")
	flag.StringVar(&sTime, "stime", "", "set start time, match file time && log time, format:yyyymmddhhmmss")
	flag.StringVar(&eTime, "etime", "", "set end time, match file time && log time, format:yyyymmddhhmmss")
	flag.IntVar(&logNum, "lognum", 0, "set log num, the num of log need to generate")
	flag.StringVar(&sPath, "spath", "./", "src path")
	flag.StringVar(&dPath, "dpath", "./", "dest path")
	flag.IntVar(&sPort, "sport", 0, "set start port, the start port to increase")
	flag.BoolVar(&debug, "debug", false, "see more debug information")
	flag.Parse()

	//命令行参数校验
	paramValid := paramCheck()
	if !paramValid {
		flag.Usage()
		return
	}

	//提取IS日志中符合条件的记录
	getCommandRecord()

	if originRecord.valid {
		//查找到符合的记录，生成文件
		filename := genIsLog()
		fmt.Printf("Match! Generate file %s successfully!\n", filename)
		//生成.ok的空文件
		genOkFile(filename)
		printOriginCommandRecord()
	} else {
		//未找到符合的记录
		fmt.Printf("Miss Match! please check the param and origin log file.\n")
	}

	return
}
