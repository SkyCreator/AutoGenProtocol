package autogen

import (
	"AutoGenProtocol/src/libFile"
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type ProtocolGenTLServer struct {
	Manager *ProtocolGenManager
}

func (p *ProtocolGenTLServer) GetName() string {
	return "TLServerGen"
}

//基础类型只需要考虑get/set,不需要考虑结构体或类中的成员函数以及数组的赋值方式
func (p *ProtocolGenTLServer) IsBaseType(typeStr string) bool {
	var compareTypes []string = []string{"unsigned int", "UINT", "int", "INT", "unsigned short", "USHORT", "short", "SHORT",
		"float", "FLOAT", "double", "DOUBLE", "unsigned char", "UCHAR", "char", "CHAR", "BYTE", "BOOL", "WORD", "long long", "INT64", "unsigned long long", "UINT64"}
	var compareID_t string = "ID_t"
	if strings.HasSuffix(typeStr, compareID_t) {
		return true
	}
	for _, ct := range compareTypes {
		if strings.Compare(ct, typeStr) == 0 {
			return true
		}
	}
	return false
}

func (p *ProtocolGenTLServer) IsSpecialType(typeStr string) bool {
	if strings.Compare(typeStr, "GUID64_t") == 0 {
		return true
	}
	return false
}

func (p *ProtocolGenTLServer) GenProtocol(outPath string, pd *ProtocolData) {
	//模板路径 todo：要放到配置文件中
	phPath := "tmpl/server/ProtocolH.tmpl"
	pcppPath := "tmpl/server/ProtocolCPP.tmpl"
	p.Manager.GenProtocolByTemplate(outPath+pd.ClassName+".h", phPath, pd)
	p.Manager.GenProtocolByTemplate(outPath+pd.ClassName+".cpp", pcppPath, pd)
	tmplPath, err := GetProtocolHandlerTmplPath(pd.ClassName)
	if err == nil {
		p.Manager.GenProtocolByTemplate(outPath+pd.ClassName+"Handler.cpp", tmplPath, pd)
	}
	p.append2ProtocolFile(pd)
}

func GetProtocolHandlerTmplPath(protocolName string) (string, error) {
	tmplPath := ""
	switch protocolName[0:2] {
	case "CG":
		tmplPath = "tmpl/server/CGProtocolHandler.tmpl"
	case "GW":
		tmplPath = "tmpl/server/GWProtocolHandler.tmpl"
	case "WG":
		tmplPath = "tmpl/server/WGProtocolHandler.tmpl"
	default:
		return tmplPath, fmt.Errorf("GetProtocolHandlerTmplPath tmplPath is not exist!")
	}
	return tmplPath, nil
}
func (p *ProtocolGenTLServer) append2ProtocolFile(pd *ProtocolData) {
	if strings.Contains(pd.ClassName, "Test") || strings.Contains(pd.ClassName, "TestHandler") {
		return
	}
	tmplPath := "tmpl/server/ProtocolHandler.tmpl"
	pos := strings.LastIndexByte(tmplPath, '/')
	tmplName := tmplPath[pos+1:]
	tmpl := template.New(tmplName)
	tmpl, err := tmpl.ParseFiles(tmplPath)
	CheckErr(err)
	handlerPath, err := p.GetProtocolHandlerPath(pd.ClassName)
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
func (p *ProtocolGenTLServer) GetWorkPath() (string, error) {
	m := p.Manager
	path, err := m.GetLineByIndex(0)
	return path, err
}
func (p *ProtocolGenTLServer) Init(m *ProtocolGenManager) {
	p.Manager = m
}
func (p *ProtocolGenTLServer) Move2WorkPath(fileName string) {
	if strings.Contains(fileName, "Test") || strings.Contains(fileName, "TestHandler") {
		return
	}
	if strings.HasSuffix(fileName, ".cs") || strings.HasSuffix(fileName, ".lua") {
		return
	}
	workpath, _ := p.GetWorkPath()
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
func (p *ProtocolGenTLServer) GetProtocolHandlerPath(protocolName string) (string, error) {
	tmplPath := ""
	workPath, _ := p.GetWorkPath()
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
