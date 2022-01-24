package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/JVLAlves/Dinamize-Inventory/internal/app/Linux"
	"github.com/JVLAlves/Dinamize-Inventory/internal/app/MacOS"
	"github.com/JVLAlves/Dinamize-Inventory/internal/app/Windows"
	regexs "github.com/JVLAlves/Dinamize-Inventory/internal/app/const-regexs"
	globals "github.com/JVLAlves/Dinamize-Inventory/internal/app/globals"
	snipe "github.com/JVLAlves/Dinamize-Inventory/internal/helpers/snipehelpers"
	logs "github.com/JVLAlves/Dinamize-Inventory/logs"
)

//função de execução principal
func main() {

	//Cria tanto a pasta para logs quanto o arquivo inicial de logs
	f := logs.InitLogs()

	//Log de inicialização
	log.Printf("Inicio de execução.")

	//Identificando sistema operacional
	switch runtime.GOOS {

	case "darwin": //darwin é o sistema operacional do Mac.
		forMacOs(f)
	case "linux":
		forLinux(f)

	case "windows":
		forWindows(f)

	default:
		log.Fatalf("Erro em encontrar o Sistema Operacional")
	}

	//mensagem de encerramento
	log.Printf("Fim de execução.")
}

//Função de execução do programa em MacOS
func forMacOs(f *os.File) {

	//Chama função de coleta especificado MacOS
	MacOS.MainProgram()

	//Variavel de Contrato do Snipeit
	mac := snipe.NewActive()

	//Populando Struct
	mac.SnipeitCPU11 = MacOS.Infos[2]
	mac.SnipeitHostname10 = MacOS.Infos[0]
	mac.SnipeitProgramasInstalados15 = MacOS.Infos[6]
	mac.Name = MacOS.Infos[0]
	mac.SnipeitOffice14 = OfficeExists(mac)

	//Passando Regex antes de popular informação de Memória

	MemoryFloat, _ := strconv.ParseFloat(MacOS.Infos[3], 64)
	//Arredondando valor númerico da variável
	MemoryRounded := math.Round(MemoryFloat)
	//Populando campo de memória com o valor tratado
	mac.SnipeitMema3Ria7 = strconv.Itoa(int(MemoryRounded)) + "GB"

	//Convertendo response de string para float
	HDFloat, _ := strconv.ParseFloat(MacOS.Infos[4], 64)
	//Arredondando valor númerico da variável
	HDRounded := math.Round(HDFloat)
	//Populando campo de HD com o valor tratado
	mac.SnipeitHd9 = strconv.Itoa(int(HDRounded)) + "GB"

	//Passando Regex antes de popular informação de Asset Tag
	mac.AssetTag = regexs.RegexAssettagDigit.FindString(MacOS.Infos[1])
	//Caso não haja digitos no campo HOSTNAME (Fonte do Asset Tag), o retorno do sistema é um Asset Tag Default (NO ASSET TAG)
	if mac.AssetTag == "" {
		mac.AssetTag = "Inválido"
		fmt.Fprintf(f, "Nenhum Asset Tag foi definido, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", MacOS.Infos[0])

	}

	//Convertendo response de string para float
	SOFloat, err := strconv.ParseFloat(MacOS.Infos[5], 64)
	//Tratando erro
	if err != nil {
		log.Fatalf("Erro na conversão do S.O. para float")
	}

	//Verificação de Versão Menores (11.5.1) e substituição por Versões Maiores (11.4)

	var SOString string
	if SOFloat >= 11.4 && SOFloat < 12.0 {
		SOString = "11.4"
	} else {
		SOString = MacOS.Infos[5]
	}

	//Alternando Versão Númerica (RETIRADA DO SISTEMA) para Versão Nominal (DEFINIDA PELA APPLE INC.)
	//ISTO PODE SER UM MAP[STRING]STRING. PARA DESCOBRIR A VERSÃO, PASSA-SE UM FOR SOBRE O MAP.
	for in, v := range globals.MacOSVersions {

		if SOString == in {
			mac.SnipeitSo8 = v
			break
		}

	}

	//Entrada Default
	mac.ModelID = globals.ID_MODELO
	mac.StatusID = globals.ID_STATUS
	mac.SnipeitModel12 = globals.MODELO_ATIVO

	VerifyIfnotEmpty(mac)

	DevExposeAll(mac)
	//Verificando a existência de um ativo semelhante no inventário Snipe it
	if snipe.Verifybytag(mac.AssetTag, globals.IP_SNIPEIT) {
		fmt.Fprintln(f, "Os dados do Ativo Criado não constam no sistema.")

		//Caso o Ativo não exista no sistema, as informações são enviadas para tal.
		snipe.PostSnipe(mac, globals.IP_SNIPEIT, f)
	} else {
		//caso já exista, o programa procura por disparidades.
		//log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Fprintln(f, "Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		ExistentActive := snipe.Getbytag(globals.IP_SNIPEIT, mac.AssetTag)
		PatchRequestUrl, IsNeeded := mac.Compare(f, ExistentActive)
		if IsNeeded {
			//Caso haja disparidades, o processo de PATCH é iniciado.

			id := snipe.Getidbytag(mac.AssetTag, globals.IP_SNIPEIT)
			snipe.Patchbyid(id, globals.IP_SNIPEIT, PatchRequestUrl)

		} else {
			//Caso não haja disparidades... Nada acontece.
			_, _ = fmt.Fprintf(f, "")
			fmt.Fprintln(f, "\nSem alterações")
		}
	}
}

//Função de execução do programa em Windows
func forWindows(f *os.File) {

	//Realiza o processo de coleta de dados do Sistema Windows e retorna as informações em um array Infos
	Windows.MainProgram()

	//Variavel de Contrato
	win := snipe.NewActive()

	//Populando Struct
	win.SnipeitCPU11 = Windows.Infos[2]

	win.SnipeitMema3Ria7 = Windows.Infos[5] + "GB"

	win.SnipeitSo8 = Windows.Infos[4]

	win.SnipeitHostname10 = Windows.Infos[0]

	win.Name = Windows.Infos[0]
	win.SnipeitHd9 = Windows.Infos[3] + "GB"

	win.AssetTag = Windows.Infos[1]

	win.SnipeitProgramasInstalados15 = Windows.ProgramasWin

	//Caso não haja digitos no campo HOSTNAME (Fonte do Asset Tag), o retorno do sistema é um Asset Tag Default (NO ASSET TAG)
	if win.AssetTag == "" {
		win.AssetTag = "Inválido"
		log.Printf("Nenhum Asset Tag foi defino, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", Windows.Infos[0])

	}

	//Entrada Default
	win.ModelID = globals.ID_MODELO
	win.StatusID = globals.ID_STATUS
	win.SnipeitModel12 = globals.MODELO_ATIVO

	VerifyIfnotEmpty(win)
	//Verificando a existência de um ativo semelhante no inventário Snipe it
	if snipe.Verifybytag(win.AssetTag, globals.IP_SNIPEIT) {
		fmt.Fprintln(f, "Os dados do Ativo Criado não constam no sistema.")

		//Caso o Ativo não exista no sistema, as informações são enviadas para tal.
		snipe.PostSnipe(win, globals.IP_SNIPEIT, f)

		log.Println("Ativo Criado enviado para o sistema.")

	} else {
		//caso já exista, o programa procura por disparidades.
		ExistentActive := snipe.Getbytag(globals.IP_SNIPEIT, win.AssetTag)
		PatchRequestUrl, IsNeeded := win.Compare(f, ExistentActive)
		if IsNeeded {
			//Caso haja disparidades, o processo de PATCH é iniciado.

			id := snipe.Getidbytag(win.AssetTag, globals.IP_SNIPEIT)
			snipe.Patchbyid(id, globals.IP_SNIPEIT, PatchRequestUrl)

		} else {
			//Caso não haja disparidades... Nada acontece.
			fmt.Fprintln(f, "\nSem alterações")
		}
	}

}

//Função de execução do programa em Linux
func forLinux(f *os.File) {

	//Realiza o processo de coleta de dados do Sistema Linux e retorna as informações em um array Infos
	Linux.MainProgram()
	Linux.Crontab()

	//Variavel de Contrato
	lin := snipe.NewActive()

	//Populando Struct
	lin.SnipeitCPU11 = Linux.Infos[0]
	lin.SnipeitSo8 = Linux.Infos[2]
	lin.SnipeitHostname10 = Linux.Infos[3]
	lin.Name = Linux.Infos[3]
	lin.SnipeitHd9 = Linux.Infos[5]
	lin.SnipeitMema3Ria7 = Linux.Infos[1]
	lin.AssetTag = Linux.Infos[4]
	//Caso não haja digitos no campo HOSTNAME (Fonte do Asset Tag), o retorno do sistema é um Asset Tag Default (NO ASSET TAG)
	if lin.AssetTag == "" {
		lin.AssetTag = "No Asset Tag"
		log.Printf("Nenhum Asset Tag foi defino, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", Linux.Infos[0])

	}

	//Entrada Default
	lin.ModelID = globals.ID_MODELO
	lin.StatusID = globals.ID_STATUS
	lin.SnipeitModel12 = globals.MODELO_ATIVO

	VerifyIfnotEmpty(lin)
	//Verificando a existência de um ativo semelhante no inventário Snipe it
	if snipe.Verifybytag(lin.AssetTag, globals.IP_SNIPEIT) {
		fmt.Fprintln(f, "Os dados do Ativo Criado não constam no sistema.")

		//Caso o Ativo não exista no sistema, as informações são enviadas para tal.
		snipe.PostSnipe(lin, globals.IP_SNIPEIT, f)

	} else {
		//caso já exista, o programa procura por disparidades.
		//log.Println("Um Ativo semelhante foi encontrado no sistema."

		ExistentActive := snipe.Getbytag(globals.IP_SNIPEIT, lin.AssetTag)
		PatchRequestUrl, IsNeeded := lin.Compare(f, ExistentActive)
		if IsNeeded {
			//Caso haja disparidades, o processo de PATCH é iniciado.
			id := snipe.Getidbytag(lin.AssetTag, globals.IP_SNIPEIT)
			snipe.Patchbyid(id, globals.IP_SNIPEIT, PatchRequestUrl)

		} else {
			//Caso não haja disparidades... Nada acontece.
			fmt.Fprintln(f, "\nSem alterações")
		}
	}

}

func VerifyIfnotEmpty(Active *snipe.CollectionT) {
	ProgramasInstalados := Active.SnipeitProgramasInstalados15
	var ActiveIndexTotal = []string{Active.Name, Active.AssetTag, Active.ModelID, Active.StatusID, Active.SnipeitMema3Ria7, Active.SnipeitSo8, Active.SnipeitHd9, Active.SnipeitHostname10, Active.SnipeitCPU11, Active.SnipeitModel12, Active.SnipeitOffice14, ProgramasInstalados}
	var EmptyField string
	var EmptyCounter int
	var EmptyList []string
	for In, v := range ActiveIndexTotal {

		if v == "" {

			switch In {

			case 0:
				EmptyField = "Name"
				EmptyList = append(EmptyList, EmptyField)
			case 1:
				EmptyField = "Asset Tag"
				EmptyList = append(EmptyList, EmptyField)
			case 2:
				EmptyField = "Model ID"
				EmptyList = append(EmptyList, EmptyField)
			case 3:
				EmptyField = "Status ID"
				EmptyList = append(EmptyList, EmptyField)
			case 4:
				EmptyField = "Memória"
				EmptyList = append(EmptyList, EmptyField)
			case 5:
				EmptyField = "SO"
				EmptyList = append(EmptyList, EmptyField)
			case 6:
				EmptyField = "HD"
				EmptyList = append(EmptyList, EmptyField)
			case 7:
				EmptyField = "Hostname"
				EmptyList = append(EmptyList, EmptyField)
			case 8:
				EmptyField = "CPU"
				EmptyList = append(EmptyList, EmptyField)
			case 9:
				EmptyField = "Model Name"
				EmptyList = append(EmptyList, EmptyField)
			case 10:
				EmptyField = "Programas Instalados"
				EmptyList = append(EmptyList, EmptyField)

			}

			EmptyCounter++
		}

	}

	if EmptyCounter > 2 {

		log.Fatalf("There are %v Empty fields, the program can't continue.\tThe Empty fields are %v\n", EmptyCounter, EmptyList)
	}

}

func OfficeExists(Active *snipe.CollectionT) string {

	ProgramasInstalados := strings.Split(Active.SnipeitProgramasInstalados15, " | ")
	OfficeCounter := 0
	for _, v := range ProgramasInstalados {

		if OfficeCounter >= 2 {
			return "Sim"

		}
		for _, comp := range globals.Office {

			if comp == v {
				OfficeCounter++
			}
		}
	}
	return "Não"

}

func DevExposeAll(Active *snipe.CollectionT) {
	ProgramasInstalados := strings.Split(Active.SnipeitProgramasInstalados15, " | ")
	var ActiveIndexTotal = []string{Active.Name, Active.AssetTag, Active.ModelID, Active.StatusID, Active.SnipeitMema3Ria7, Active.SnipeitSo8, Active.SnipeitHd9, Active.SnipeitHostname10, Active.SnipeitCPU11, Active.SnipeitModel12, Active.SnipeitOffice14}

	fmt.Println("HardWare Data")
	for _, v := range ActiveIndexTotal {

		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println("Programas Instalados")
	for _, v := range ProgramasInstalados {

		fmt.Println(v)
	}
	log.Fatalln("Safe End")
}
