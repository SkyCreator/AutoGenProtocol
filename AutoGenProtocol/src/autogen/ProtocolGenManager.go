package autogen

import (
	"AutoGenProtocol/src/libFile"
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type IProtocolGen interface {
	IsBaseType(typeStr string) bool
	IsSpecialType(typeStr string) bool
	IsSpecialExType(typeStr string) bool
	GenProtocol(outPath string, pd *ProtocolData)
	Init(m *ProtocolGenManager)
	GetWorkPath() (string, error)
	Move2WorkPath(fileName string)
	GetName() string
}

type ProtocolDataVar struct {
	VarType          string
	VarValue         string
	VarCount         string
	VarIsArray       bool
	VarIsChar        bool
	VarIsNoSet       bool
	VarIsBaseType    bool
	VarIsSpecialType bool
}
type ProtocolData struct {
	ClassName    string
	ProtocolName string
	DataVar      []ProtocolDataVar
	//协议生成类型
	//-1:生成服务器c++文件
	//0：生成客户端cs文件
	//1：生成客户端lua文件
	//2：同时生成客户端cs和lua文件
	VarGenType      int
	IsNeedSpecialEx bool //true说明需要在文件中加入特殊引用，如using UnityEngine;
}
type DataLine struct {
	flag  string
	value string
}

type ProtocolGenManager struct {
	mapProtocolGen   map[string]*IProtocolGen
	configLines      []string
	ProtocolFilePath string //协议目录
	ProtocolOutPath  string //输出目录
}

func (m *ProtocolGenManager) SetProtocolFilePath(path string) {
	m.ProtocolFilePath = path
	libFile.CheckOrCreateDir(m.ProtocolFilePath)
}
func (m *ProtocolGenManager) SetProtocolOutPath(path string) {
	m.ProtocolOutPath = path
	libFile.CheckOrCreateDir(m.ProtocolOutPath)
}
func (m *ProtocolGenManager) GetProtocolFiles() []fs.FileInfo {
	files, err := ioutil.ReadDir(m.ProtocolFilePath)
	CheckErr(err)
	return files
}

func (m *ProtocolGenManager) genAllProtocol() {
	files, _ := libFile.GetFilesFromDir(m.ProtocolFilePath)
	for _, file := range files {
		pd := m.GetProtocolData(m.ProtocolFilePath + file.Name())
		pGen := m.GetProtocolGen(pd.ProtocolName)
		(*pGen).GenProtocol(m.ProtocolOutPath, pd)
	}
}
func (m *ProtocolGenManager) Move2WorkPath() {
	files, _ := libFile.GetFilesFromDir(m.ProtocolOutPath)
	for _, file := range files {
		pGen := m.GetProtocolGenByFileName(file.Name())
		if pGen != nil {
			(*pGen).Move2WorkPath(file.Name())
		}
	}

}
func (m *ProtocolGenManager) RegisterAllProtocolGen() {
	m.mapProtocolGen = make(map[string]*IProtocolGen)
	m.AddProtocolGen(new(ProtocolGenPSClient))
	m.AddProtocolGen(new(ProtocolGenTLServer))
}

func (m *ProtocolGenManager) GetProtocolGenByName(genName string) *IProtocolGen {
	return m.mapProtocolGen[genName]
}

//获得生成类
func (m *ProtocolGenManager) GetProtocolGenByFileName(fileName string) *IProtocolGen {
	if strings.HasSuffix(fileName, ".cs") || strings.HasSuffix(fileName, ".lua") {
		return m.GetProtocolGenByName("PSClientGen")
	} else if strings.HasSuffix(fileName, ".h") || strings.HasSuffix(fileName, ".cpp") {
		return m.GetProtocolGenByName("TLServerGen")
	}
	return nil
}

//获得生成类
func (m *ProtocolGenManager) GetProtocolGen(protocolName string) *IProtocolGen {
	if strings.ContainsRune(protocolName, '|') {
		return m.GetProtocolGenByName("PSClientGen")
	} else {
		return m.GetProtocolGenByName("TLServerGen")
	}
}
func (m *ProtocolGenManager) AddProtocolGen(protocolGen IProtocolGen) {
	if m.mapProtocolGen[protocolGen.GetName()] == nil {
		m.mapProtocolGen[protocolGen.GetName()] = &protocolGen
		protocolGen.Init(m)
	}
}
func (m *ProtocolGenManager) init() {
	m.SetProtocolFilePath("./Protocol/")
	m.SetProtocolOutPath("./Out/")
	m.LoadConfig("config.ini")
	m.RegisterAllProtocolGen()
}
func (m *ProtocolGenManager) run() {
	m.genAllProtocol()
	m.Move2WorkPath()
}
func (m *ProtocolGenManager) LoadConfig(path string) {
	lines, err := libFile.GetFileContextLines(path)
	CheckErr(err)
	m.configLines = lines
}
func (m *ProtocolGenManager) GetLineByIndex(index int) (string, error) {
	if index < 0 && len(m.configLines) <= index {
		return "", fmt.Errorf("GetLineByIndex is error!index=%d", index)
	}
	return m.configLines[index], nil
}

func Filter_m_(value string) string {
	if strings.HasPrefix(value, "m_") {
		return value[2:]
	}
	return value
}

//是否是数组
func IsArray(value string) bool {
	return strings.HasSuffix(value, "]")
}

//针对数组类型整理数据
func (m *ProtocolGenManager) parseArrayDataLine(protocoldata *ProtocolData, dl *DataLine) {
	isChar := false
	if strings.ToLower(dl.flag) == "char" {
		isChar = true
	}
	isbasetype := false
	pGen := m.GetProtocolGen(protocoldata.ProtocolName)
	if (*pGen).IsBaseType(dl.flag) {
		isbasetype = true
	}
	start := strings.IndexByte(dl.value, '[')
	end := strings.IndexByte(dl.value, ']')
	countStr := dl.value[start+1 : end]
	startStr := dl.value[0:start]
	if strings.ContainsRune(protocoldata.ProtocolName, '|') {
		protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: "short", VarValue: Filter_m_(startStr) + "Count", VarCount: "", VarIsNoSet: true})
	} else {
		protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: "SHORT", VarValue: Filter_m_(startStr) + "Count", VarCount: "", VarIsNoSet: true})
	}
	protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: Filter_m_(startStr), VarCount: countStr, VarIsArray: true, VarIsChar: isChar, VarIsBaseType: isbasetype})
}
func (m *ProtocolGenManager) parseDataLine(protocoldata *ProtocolData, dl *DataLine) {
	pGen := m.GetProtocolGen(protocoldata.ProtocolName)
	switch dl.flag {
	case "class":
		protocoldata.ClassName = dl.value
	case "protocol":
		protocoldata.ProtocolName = dl.value
	default:
		if IsArray(dl.value) {
			m.parseArrayDataLine(protocoldata, dl)
		} else if (*pGen).IsBaseType(dl.flag) {
			protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: Filter_m_(dl.value), VarIsBaseType: true})
		} else if (*pGen).IsSpecialType(dl.flag) {
			protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: Filter_m_(dl.value), VarIsSpecialType: true})
			if !protocoldata.IsNeedSpecialEx {
				protocoldata.IsNeedSpecialEx = (*pGen).IsSpecialExType((dl.flag))
			}
		} else if IsString(dl.flag) {
			protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: "short", VarValue: Filter_m_(dl.value) + "Count", VarCount: "", VarIsNoSet: true})
			protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: Filter_m_(dl.value)})
		} else {
			protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: Filter_m_(dl.value)})
		}

	}

}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (m *ProtocolGenManager) GenProtocolByTemplate(outPath string, tmplPath string, pd *ProtocolData) {
	pos := strings.LastIndexByte(tmplPath, '/')
	tmplName := tmplPath[pos+1:]
	tmpl := template.New(tmplName)
	tmpl, err := tmpl.ParseFiles(tmplPath)
	CheckErr(err)
	os.Remove(outPath)
	f, err := os.OpenFile(outPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	CheckErr(err)
	err = tmpl.Execute(f, pd)
	CheckErr(err)
	f.Close()
	fmt.Printf("gen file：%s\n", outPath)
}

func (m *ProtocolGenManager) GetProtocolData(filename string) *ProtocolData {
	lines, err := libFile.GetFileContextLines(filename)
	if err != nil {
		CheckErr(err)
	}
	linelist := make([]DataLine, 0)
	for _, line := range lines {
		if line != "" {
			str := strings.Split(line, " ")
			linelist = append(linelist, DataLine{flag: str[0], value: str[1]})
		}
	}
	var pd ProtocolData
	for i := 0; i < len(linelist); i++ {
		m.parseDataLine(&pd, &linelist[i])
	}
	return &pd
}
func CheckNameRepeat(path string, classname string) bool {
	var f *os.File
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
	CheckErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content := scanner.Text()
		if strings.Contains(content, classname) {
			return true
		}
	}
	return false
}
func GetFileContextLinesNum(filename string) int {
	return libFile.GetFileContextLinesNum(filename)
}
func AppendFileWithTmpl(filepath string, tmplname string, tmplinfo string, flagtext string, pd *ProtocolData) {
	buf := GetAnalysisTmplText(tmplname, tmplinfo, pd)
	libFile.AppendFileContent(filepath, buf)
}
func InsertFileWithTmpl(filepath string, tmplname string, tmplinfo string, flagtext string, pd *ProtocolData) {
	buf := GetAnalysisTmplText(tmplname, tmplinfo, pd)
	fileContext := libFile.GetFileContext(filepath)
	if len(fileContext) == 0 {
		libFile.AppendFileContent(filepath, buf)
		return
	}
	var upIndex int = strings.Index(fileContext, flagtext)
	//var downIndex int = strings.LastIndex(fileContext, flagtext)
	upoffset := len(fileContext[0:upIndex])
	findContent := fileContext[upIndex:]
	offset := FindFlagText(findContent, buf)
	libFile.FileInsertInfo(filepath, buf, (int64)(upoffset+offset))
}
func GetAnalysisTmplText(tmplname string, tmplinfo string, pd *ProtocolData) string {
	tmpl := template.New(tmplname)
	tmpl, err := tmpl.Parse(tmplinfo)
	CheckErr(err)
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, pd)
	CheckErr(err)
	return buf.String()
}
func FindFlagText(findContent string, analysisTmplText string) int {
	lines := strings.Split(findContent, "\n")
	var findStr string = ""
	for _, lineStr := range lines {
		if lineStr != "" {
			if strings.Trim(lineStr, "\t") < strings.Trim(analysisTmplText, "\t") {
				findStr = lineStr
				break
			}
		}
	}
	return strings.Index(findContent, findStr) + len(findStr)
}
