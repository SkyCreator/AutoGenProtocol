package autogen

import (
	"AutoGenProtocol/src/libFile"
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type ProtocolGenPSClient struct {
	Manager *ProtocolGenManager
}

func (p *ProtocolGenPSClient) GetName() string {
	return "PSClientGen"
}

//基础类型只需要考虑get/set,不需要考虑结构体或类中的成员函数以及数组的赋值方式
func (p *ProtocolGenPSClient) IsBaseType(typeStr string) bool {
	var compareTypes []string = []string{"byte", "ulong", "int", "long", "uint", "float", "double", "ushort", "short", "bool", "char"}
	for _, ct := range compareTypes {
		if strings.Compare(ct, typeStr) == 0 {
			return true
		}
	}
	return false
}
func (p *ProtocolGenPSClient) IsSpecialType(typeStr string) bool {
	if strings.Compare(typeStr, "WORLD_POS_3D") == 0 ||
		strings.Compare(typeStr, "Vector3") == 0 ||
		strings.Compare(typeStr, "Vector2") == 0 {
		return true
	}
	return false
}
func (p *ProtocolGenPSClient) IsSpecialExType(typeStr string) bool {
	if strings.Compare(typeStr, "Vector3") == 0 ||
		strings.Compare(typeStr, "Vector2") == 0 {
		return true
	}
	return false
}
func (p *ProtocolGenPSClient) GenProtocol(outPath string, pd *ProtocolData) {
	if !strings.ContainsRune(pd.ProtocolName, '|') {
		return
	}
	str := strings.Split(pd.ProtocolName, "|")
	pd.ProtocolName = str[0]
	if strings.Compare(str[1], "cs") == 0 {
		pd.VarGenType = 0
	}
	if strings.Compare(str[1], "lua") == 0 {
		pd.VarGenType = 1
	}
	if len(str) > 2 {
		pd.VarGenType = 2
	}
	if strings.Compare(pd.ClassName[0:2], "CG") == 0 {
		if pd.VarGenType != 1 {
			pCGCSPath := "tmpl/client/CGProtocolCS.tmpl"
			p.Manager.GenProtocolByTemplate(outPath+pd.ClassName+".cs", pCGCSPath, pd)
		}
		if pd.VarGenType != 0 {
			pCGLuaPath := "tmpl/client/CGProtocolLua.tmpl"
			p.Manager.GenProtocolByTemplate(outPath+pd.ClassName+".lua", pCGLuaPath, pd)
		}
	} else if strings.Compare(pd.ClassName[0:2], "GC") == 0 {
		if pd.VarGenType != 1 {
			pGCCSPath := "tmpl/client/GCProtocolCS.tmpl"
			pHandlerCSPath := "tmpl/client/ProtocolHandlerCS.tmpl"
			p.Manager.GenProtocolByTemplate(outPath+pd.ClassName+".cs", pGCCSPath, pd)
			p.Manager.GenProtocolByTemplate(outPath+pd.ClassName+"Handler.cs", pHandlerCSPath, pd)
		}
		if pd.VarGenType != 0 {
			pGCLuaPath := "tmpl/client/GCProtocolLua.tmpl"
			p.Manager.GenProtocolByTemplate(outPath+pd.ClassName+".lua", pGCLuaPath, pd)
		}
	}
	p.Append2ProtocolFile(pd)
}

func (p *ProtocolGenPSClient) GetWorkPath() (string, error) {
	m := p.Manager
	path, err := m.GetLineByIndex(1)
	return path, err
}
func (p *ProtocolGenPSClient) Init(m *ProtocolGenManager) {
	p.Manager = m
}
func (p *ProtocolGenPSClient) Move2WorkPath(fileName string) {
	if strings.Contains(fileName, "Test") || strings.Contains(fileName, "TestHandler") {
		return
	}
	if strings.HasSuffix(fileName, ".cpp") || strings.HasSuffix(fileName, ".h") {
		return
	}
	workpath, _ := p.GetWorkPath()
	var finalPath string = ""
	switch fileName[0:2] {
	case "CG":
		if strings.HasSuffix(fileName, ".cs") {
			finalPath = workpath + "/Assets/GameMain/Scripts/Network/Packet/"
		} else {
			finalPath = workpath + "/Assets/GameMain/LuaScripts/Packet/Request/"
		}
	case "GC":
		if strings.HasSuffix(fileName, ".cs") {
			finalPath = workpath + "/Assets/GameMain/Scripts/Network/Packet/"
		} else {
			finalPath = workpath + "/Assets/GameMain/LuaScripts/Packet/Response/"
		}
	}
	_, err := libFile.CopyFile(finalPath+fileName, "Out/"+fileName)
	//err := os.Rename("Out/"+fileName, finalPath+fileName)
	if err != nil {
		fmt.Printf("The system cannot find the path specified.The path is %s\n", finalPath)
	}
	fmt.Printf("move file：%s\n", finalPath+fileName)
}

//todo:准备私有化
func IsString(typeStr string) bool {
	return strings.Compare(typeStr, "string") == 0
}

func (p *ProtocolGenPSClient) Append2ProtocolFile(pd *ProtocolData) {
	//if strings.Contains(pd.ClassName, "Test") || strings.Contains(pd.ClassName, "TestHandler") {
	//	return
	//}
	if strings.Compare(pd.ClassName[0:2], "CG") == 0 {
		return
	}
	tmpl := template.New("PacketIdList")
	tmpl, err := tmpl.Parse("\n\t[{{.ProtocolName}}]		=	\"Packet/Response/{{.ClassName}}\",")
	CheckErr(err)
	workPath, _ := (*p).GetWorkPath()
	handlerPath := workPath + "/Assets/GameMain/LuaScripts/Packet/PacketIdList.lua"
	if CheckNameRepeat(handlerPath, pd.ClassName) {
		return
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, pd)
	fileContext := libFile.GetFileContext(handlerPath)
	index := strings.LastIndexByte(fileContext, '}')
	libFile.FileInsertInfo(handlerPath, buf.String(), (int64)(index-1))
	CheckErr(err)
}
