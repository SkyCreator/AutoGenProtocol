package autogen

import (
	"AutoGenProtocol/src/libFile"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

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
}
type DataLine struct {
	flag  string
	value string
}

//是否是数组
func IsArray(value string) bool {
	return strings.HasSuffix(value, "]")
}
func IsBaseType(value string) bool {
	var compareTypes []string = []string{"unsigned int", "UINT", "int", "INT", "unsigned short", "USHORT", "short", "SHORT",
		"float", "FLOAT", "double", "DOUBLE", "unsigned char", "UCHAR", "char", "CHAR", "BYTE", "BOOL", "WORD", "long long", "INT64", "unsigned long long", "UINT64"}
	var compareID_t string = "ID_t"
	if strings.HasSuffix(value, compareID_t) {
		return true
	}
	for _, ct := range compareTypes {
		if strings.Compare(ct, value) == 0 {
			return true
		}
	}
	return false
}
func IsSpecialType(value string) bool {
	if strings.Compare(value, "GUID64_t") == 0 || strings.Compare(value, "WORLD_POS_3D") == 0 {
		return true
	}
	return false
}
func parseArrayDataLine(protocoldata *ProtocolData, dl *DataLine) {
	isChar := false
	if strings.ToLower(dl.flag) == "char" {
		isChar = true
	}
	isbasetype := false
	if IsBaseType(dl.flag) {
		isbasetype = true
	}
	start := strings.IndexByte(dl.value, '[')
	end := strings.IndexByte(dl.value, ']')
	countStr := dl.value[start+1 : end]
	startStr := dl.value[0:start]
	protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: startStr, VarCount: countStr, VarIsArray: true, VarIsChar: isChar, VarIsBaseType: isbasetype})
	protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: "SHORT", VarValue: startStr + "Count", VarCount: "", VarIsNoSet: true})
}
func parseDataLine(protocoldata *ProtocolData, dl *DataLine) {
	switch dl.flag {
	case "class":
		protocoldata.ClassName = dl.value
	case "protocol":
		protocoldata.ProtocolName = dl.value
	default:
		if IsArray(dl.value) {
			parseArrayDataLine(protocoldata, dl)
		} else if IsBaseType(dl.flag) {
			protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: dl.value, VarIsBaseType: true})
		} else if IsSpecialType(dl.flag) {
			protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: dl.value, VarIsSpecialType: true})
		} else {
			protocoldata.DataVar = append(protocoldata.DataVar, ProtocolDataVar{VarType: dl.flag, VarValue: dl.value})
		}

	}

}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetProtocolHandlerTmplPath(protocolName string) (string, error) {
	tmplPath := ""
	switch protocolName[0:2] {
	case "CG":
		tmplPath = "tmpl/CGProtocolHandler.tmpl"
	case "GW":
		tmplPath = "tmpl/GWProtocolHandler.tmpl"
	case "WG":
		tmplPath = "tmpl/WGProtocolHandler.tmpl"
	default:
		return tmplPath, fmt.Errorf("GetProtocolHandlerTmplPath tmplPath is not exist!")
	}
	return tmplPath, nil
}
func GetProtocolHandlerPath(protocolName string) (string, error) {
	tmplPath := ""
	workPath := GetWorkPath()
	switch protocolName[0:2] {
	case "GC":
		tmplPath = workPath + "/Server/Server/Packets/GCHandler.cpp"
	case "GW":
		tmplPath = workPath + "/Server/Server/Packets/GWHandler.cpp"
	case "WG":
		tmplPath = workPath + "/World/World/Packets/WGHandler.cpp"
	default:
		return tmplPath, fmt.Errorf("GetProtocolHandlerPath tmplPath is not exist!")
	}
	return tmplPath, nil
}
func GenProtocol(outPath string, tmplPath string, pd *ProtocolData) {
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
func Append2ProtocolFile(pd *ProtocolData) {
	tmplPath := "tmpl/ProtocolHandler.tmpl"
	pos := strings.LastIndexByte(tmplPath, '/')
	tmplName := tmplPath[pos+1:]
	tmpl := template.New(tmplName)
	tmpl, err := tmpl.ParseFiles(tmplPath)
	CheckErr(err)
	handlerPath, err := GetProtocolHandlerPath(pd.ClassName)
	if err != nil {
		return
	}
	var f *os.File
	f, err = os.OpenFile(handlerPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	CheckErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content := scanner.Text()
		if strings.Contains(content, pd.ClassName) {
			return
		}
	}
	err = tmpl.Execute(f, pd)
	CheckErr(err)
	defer f.Close()
}
func GetProtocolData(filename string) *ProtocolData {
	lines, err := libFile.GetFileContext(filename)
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
		parseDataLine(&pd, &linelist[i])
	}
	return &pd
}
func GenAllProtocol(pd *ProtocolData) {
	//输出路径 todo：要放到配置文件中
	outPath := "./Out/"
	_, err := os.Stat(outPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("./Out/", os.ModePerm)
			CheckErr(err)
		}
	}
	//模板路径 todo：要放到配置文件中
	phPath := "tmpl/ProtocolH.tmpl"
	pcppPath := "tmpl/ProtocolCPP.tmpl"
	GenProtocol(outPath+pd.ClassName+".h", phPath, pd)
	GenProtocol(outPath+pd.ClassName+".cpp", pcppPath, pd)
	tmplPath, err := GetProtocolHandlerTmplPath(pd.ClassName)
	if err == nil {
		GenProtocol(outPath+pd.ClassName+"Handler.cpp", tmplPath, pd)
	}
}
func GetWorkPath() string {
	lines, err := libFile.GetFileContext("config.ini")
	CheckErr(err)
	WorkPath := lines[0]
	return WorkPath
}
func AutoGenProtocol() {

	path := "Protocol/"
	files, err := ioutil.ReadDir(path)
	CheckErr(err)
	for _, file := range files {
		pd := GetProtocolData(path + file.Name())
		GenAllProtocol(pd)
		Append2ProtocolFile(pd)
	}
	files, err = ioutil.ReadDir("Out")
	for _, file := range files {
		//fmt.Println(file.Name())
		Move2WorkPath(file.Name())
	}
}

func Move2WorkPath(fileName string) {
	workpath := GetWorkPath()
	var finalPath string = ""
	switch fileName[0:2] {
	case "CG":
		fallthrough
	case "GC":
		fallthrough
	case "WG":
		if strings.Contains(fileName, "Handler") {
			finalPath = workpath + "/Server/Server/Packets/"
		} else {
			finalPath = workpath + "/Common/Packets/"
		}
	case "GW":
		if strings.Contains(fileName, "Handler") {
			finalPath = workpath + "/World/World/Packets/"
		} else {
			finalPath = workpath + "/Common/Packets/"
		}
	}
	_, err := libFile.CopyFile(finalPath+fileName, "Out/"+fileName)
	//err := os.Rename("Out/"+fileName, finalPath+fileName)
	if err != nil {
		fmt.Printf("The system cannot find the path specified.The path is %s\n", finalPath)
	}
	fmt.Printf("move file：%s\n", finalPath+fileName)
}
