package Windows

import (
	"bytes"
	"fmt"
	"math"
	"os/exec"
	"strconv"

	"github.com/JVLAlves/Dinamize-Inventory/regexs"
)

type PowerShell struct {
	powerShell string
}

var Infos = []string{}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

//Definindo os Argumentos necessários para executar um comando no powershell
func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}
func MainProgram() {
	posh := New()

	//APLICANDO OS COMANDO LITERAIS DO POWERSHELL

	//Aplicando comando para coletar o HOSTNAME e salvando o stdout em uma variável
	hostname, _, err := posh.Execute("Systeminfo")
	//Verificando se o comando retorna um erro não nulo
	if err != nil {
		fmt.Println(err)
	}
	//Compilando o regex do hostname.
	RegexpHostnameWin := regexs.RegexpHostnameWin
	//Aplicando o regex ao retorno obtido na vairável
	resultHostNameWin := RegexpHostnameWin.FindStringSubmatch(hostname)
	//Coletando o index 0 do slice string obtido após aplicação do regexp
	regexedHostNameWin := resultHostNameWin[0]
	RegexedAssettagWin := regexs.RegexAssettagDigit.FindString(regexedHostNameWin)
	//Guardando o valor final no array String Infos
	Infos = append(Infos, regexedHostNameWin)
	Infos = append(Infos, RegexedAssettagWin)

	//Aplicando comando para coletar o CPU e salvando o stdout em uma variável
	cpu, _, err := posh.Execute("Get-WmiObject -Class Win32_Processor -ComputerName . | Select-Object -Property \"name\"")
	//Verificando se o comando retorna um erro não nulo
	if err != nil {
		fmt.Println(err)
	}
	//Compilando o regex do CPU.
	RegexCPUWin := regexs.RegexCPU
	//Aplicando o regex ao retorno obtido na vairável
	resultCPUWin := RegexCPUWin.FindStringSubmatch(cpu)
	//Coletando o index 1 do slice string obtido após aplicação do regexp visto que o index 0 é nulo!
	regexedCPUWin := resultCPUWin[1]
	//Guardando o valor final no array String Infos
	Infos = append(Infos, regexedCPUWin)

	//Aplicando comando para coletar o HD e salvando o stdout em uma variável
	hd, _, err := posh.Execute("get-WMIobject Win32_LogicalDisk -Filter \"DeviceID = 'C:'\" | Select-Object -Property \"Size\"")
	//Verificando se o comando retorna um erro não nulo
	if err != nil {
		fmt.Println(err)
	}
	//Compilando o regex do HD basedo no pkg regexs
	RegexHDWin := regexs.RegexHDWin
	//Aplicando o regex ao retorno obtido na vairável
	resultHDWin := RegexHDWin.FindStringSubmatch(hd)
	//Coletando o index 0 do slice string obtido após aplicação do regexp
	regexedHDWin := resultHDWin[0]
	//Convertendo a variável regexedHDWin para o tipo float64
	floatHDWin, _ := strconv.ParseFloat(regexedHDWin, 64)
	//Convertendo o valor da variável floatHDWin de bytes para GB e arredondando o valor final
	roundedHDWin := math.Round(floatHDWin / math.Pow10(9))
	//Corvertendo a variável roundedHDWin do tipo Float64 para o tipo String
	intHDWin := strconv.FormatFloat(roundedHDWin, 'f', -1, 64)
	//Guardando o valor final no array String Infos
	Infos = append(Infos, intHDWin)

	//Aplicando comando para coletar o SO e salvando o stdout em uma variável
	so, _, err := posh.Execute("systeminfo | Select-String -Pattern \"Nome do sistema operacional:\"")
	//Verificando se o comando retorna um erro não nulo
	if err != nil {
		fmt.Println(err)
	}
	//Compilando o regex do SO.
	RegexSOWin := regexs.RegexSOWin
	//Aplicando o regex ao retorno obtido na vairável
	resultSOWin := RegexSOWin.FindStringSubmatch(so)
	//Coletando o index 1 do slice string obtido após aplicação do regexp visto que o index 0 é nulo!
	regexedSOWin := resultSOWin[1]
	//Guardando o valor final no array String Infos
	Infos = append(Infos, regexedSOWin)

	//Aplicando comando para coletar o MEMÓRIA e salvando o stdout em uma variável
	memoria, _, err := posh.Execute("systeminfo")
	//Verificando se o comando retorna um erro não nulo
	if err != nil {
		fmt.Println(err)
	}
	//Compilando o regex do MEMORIA basedo no pkg regexs
	RegexMemoriaWin := regexs.RegexMemoriaWin
	//Aplicando o regex ao retorno obtido na vairável
	resultMemoriaWin := RegexMemoriaWin.FindStringSubmatch(memoria)
	//Coletando o index 1 do slice string obtido após aplicação do regexp visto que o index 0 é nulo!
	regexedMemoriaWin := resultMemoriaWin[1]
	FloatMemoriaWin, _ := strconv.ParseFloat(regexedMemoriaWin, 64)
	RoundedMemoriaWin := strconv.FormatFloat(math.Round(FloatMemoriaWin), 'f', -1, 64)
	//Guardando o valor final no array String Infos
	Infos = append(Infos, RoundedMemoriaWin)

}
