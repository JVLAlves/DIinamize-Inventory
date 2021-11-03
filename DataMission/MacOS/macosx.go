package MacOS

import (
	"log"
	"os/exec"
	"strings"

	regexs "github.com/JVLAlves/Dinamize-Inventory/rgx"
)

//Lista para Informações armazenadas
var Infos []string

func MainProgram() {

	//Hostname
	HostnameCmd := `uname -n`
	cmd := exec.Command("bash", "-c", HostnameCmd)

	HostnameCmdByt, err := cmd.Output()

	if err != nil {
		log.Printf("Erro na execução comando de hostname: %v\n", err)
	}

	HostnameCmdBody := string(HostnameCmdByt)
	Host := strings.TrimSpace(HostnameCmdBody)
	AssetTag := regexs.RegexAssettagDigit.FindString(HostnameCmdBody)
	Infos = append(Infos, Host)
	Infos = append(Infos, AssetTag)

	//CPU
	CPUcmd := `sysctl -a |grep machdep.cpu.brand_string|awk '{print $2, $3, $4}'`
	cmd = exec.Command("bash", "-c", CPUcmd)
	CPUcmdByt, err := cmd.Output()

	CPUcmdBody := string(CPUcmdByt)
	CPU := strings.TrimSpace(CPUcmdBody)
	if err != nil {
		log.Printf("Erro na execução comando de CPU: %v\n", err)
	}

	Infos = append(Infos, CPU)

	//RAM Memory
	Memorycmd := `hostinfo |grep memory |awk '{print $4, $5}'`
	cmd = exec.Command("bash", "-c", Memorycmd)
	MemorycmdByt, err := cmd.Output()
	if err != nil {
		log.Printf("Erro na execução comando de memory: %v\n", err)
	}
	MemorycmdBody := string(MemorycmdByt)
	Memory := regexs.RegexHDandMemory.FindString(MemorycmdBody)
	Infos = append(Infos, Memory)

	//HD
	HDcmd := `diskutil list|grep disk0s2 | awk '{print $5, $6}'`

	cmd = exec.Command("bash", "-c", HDcmd)
	HDcmdByt, err := cmd.Output()
	if err != nil {
		log.Printf("Erro na execução comando de HD: %v\n", err)
	}
	HDcmdBody := string(HDcmdByt)
	HD := regexs.RegexHDandMemory.FindString(HDcmdBody)
	Infos = append(Infos, HD)

	//S.O.
	SOcmd := `sw_vers -productVersion`
	cmd = exec.Command("bash", "-c", SOcmd)
	SOcmdByt, err := cmd.Output()
	if err != nil {
		log.Printf("Erro na execução comando de HD: %v\n", err)
	}
	SOcmdBody := string(SOcmdByt)

	SO := regexs.RegexMacOS.FindString(SOcmdBody)

	Infos = append(Infos, SO)

}
