package Linux

import (
	"bufio"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"

	globals "github.com/JVLAlves/Dinamize-Inventory/cmd"
	regexs "github.com/JVLAlves/Dinamize-Inventory/rgx"
)

//Variáveis de armazenamento dos dados da máquina
var Linhas = []string{}
var Infos = []string{}

func MainProgram() {
	// Abrindo o Arquivo CPU
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %s", err)
	}

	//Lendo o Arquivo CPU
	fileScanner := bufio.NewScanner(file)

	//Lendo linha a linha
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo CPU a variável
	var ProcFileinfo []string
	ProcFileinfo = append(ProcFileinfo, Linhas[4])

	for _, v := range ProcFileinfo { //

		CPU := regexs.RegexCPU.FindString(v)
		if CPU != "" {
			Infos = append(Infos, CPU)
			break
		}

	}

	//Tratando o ocasional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Tamanho do disco)
	Memorycmd := exec.Command("bash", "-c", "free -h |grep Mem |awk '{print $2}'")
	MemorycmdByt, err := Memorycmd.Output()

	if err != nil {
		log.Println("Erro na execução do comando de memória: ", err)
	}

	MemorycmdBody := string(MemorycmdByt)

	//Passando Regex antes de popular informação de Memória
	MemoryRegex := regexs.RegexHDandMemory.FindStringSubmatch(MemorycmdBody)
	//Convertendo response de string para float
	MemoryFloat, _ := strconv.ParseFloat(MemoryRegex[1], 64)
	//Arredondando valor númerico da variável
	MemoryRounded := math.Round(MemoryFloat)
	//Populando campo de memória com o valor tratado
	Memory := strconv.Itoa(int(MemoryRounded)) + "GB"

	// adicionando informação encontrada no arquivo "tamanhoDoHd.txt" a variável
	Infos = append(Infos, Memory)

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (S.O.)
	SOcmd := exec.Command("bash", "-c", "lsb_release -d |grep Description |awk '{print $2,$3,$4}'")
	SOcmdByt, err := SOcmd.Output()

	if err != nil {
		log.Println("Erro na execução do comando de SO: ", err)
	}

	SOcmdBody := string(SOcmdByt)
	SO := strings.TrimSpace(SOcmdBody)
	// adicionando informação encontrada no arquivo "SO.txt" a variável
	Infos = append(Infos, SO)

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Hostname)
	Hostcmd := exec.Command("bash", "-c", "hostname", "hostname.txt")
	HostcmdByt, _ := Hostcmd.Output()
	HostcmdBody := string(HostcmdByt)
	Host := strings.TrimSpace(HostcmdBody)
	Infos = append(Infos, Host)

	Assettag := regexs.RegexAssettagDigit.FindString(Host)
	Infos = append(Infos, Assettag)

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Tamanho do Disco)
	cmd := exec.Command("bash", "-c", "lsblk |grep disk |awk '{print $4}'", "tamanhoDoDisco.txt")
	HDcmdByt, err := cmd.Output()

	if err != nil {
		log.Println("Erro na execução do comando de HD: ", err)
	}

	HDcmdBody := string(HDcmdByt)
	HD := strings.TrimSpace(HDcmdBody)
	//Passando Regex antes de popular informação de HD (COLETA: Número com vírgula)
	HDRegex := regexs.RegexHDandMemory.FindStringSubmatch(HD)
	//Separação do result
	HDSplitted := strings.Split(HDRegex[1], ",")
	//Integração do result utilizando ponto (Padrão para conversão)
	HDJoined := strings.Join(HDSplitted, ".")
	//Convertendo response de string para float
	HDFloat, _ := strconv.ParseFloat(HDJoined, 64)
	//Arredondando valor númerico da variável
	HDRounded := math.Round(HDFloat)
	HD = strconv.Itoa(int(HDRounded)) + "GB"
	// adicionando informação encontrada no arquivo "tamanhoDoDisco.txt" a variável
	Infos = append(Infos, HD)

}

func Crontab() {
	var Boolean bool = true
	var home string
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error getting home user directory: ", err)
	}

	_, err = os.Stat(home + "/" + globals.CRONTABEXISTS_FILENAME)
	if err != nil {
		if os.IsNotExist(err) {
			Boolean = false
		}
	}

	if !Boolean {

		user := os.Getenv("USERNAME")
		Movecmd := "mv " + globals.LINUX_EXECNAME + " " + home
		cmd := exec.Command("bash", "-c", Movecmd)
		err = cmd.Run()

		if err != nil {
			log.Fatalln("Error moving exec file: ", err, Movecmd, "'Are you trying to test the go file or running the bin file?'")
		}

		filename := home + "/." + globals.APPNAME + "-crontab.txt"
		filecmd := "touch " + filename
		cmd = exec.Command("bash", "-c", filecmd)
		err = cmd.Run()

		if err != nil {
			log.Fatalln("Error creating crontrab file: ", err)
		}
		contentcmd := "echo " + globals.CRONTAB_CONTENT +
			" > " + filename
		cmd = exec.Command("bash", "-c", contentcmd)
		err = cmd.Run()

		if err != nil {
			log.Fatalln("Error writing on crontab file: ", err)
		}
		CrontabInitcmd := "crontab " + "-u " + user + " " + filename
		cmd = exec.Command("bash", "-c", CrontabInitcmd)
		err = cmd.Run()

		if err != nil {
			log.Fatalln("Error initalizing crontab: ", err)
		}

		CrontabExistsFilecmd := "touch " + home + "/" + globals.CRONTABEXISTS_FILENAME
		cmd = exec.Command("bash", "-c", CrontabExistsFilecmd)
		err = cmd.Run()
		if err != nil {
			log.Fatalln("Error creating CrontabExists file: ", err)
		}
		CrontabExistsContentcmd := "echo " + globals.CRONTABEXISTS_CONTENT + " > " + home + "/" + globals.CRONTABEXISTS_FILENAME
		cmd = exec.Command("bash", "-c", CrontabExistsContentcmd)
		err = cmd.Run()
		if err != nil {
			log.Fatalln("Error writing CrontabExists file: ", err)
		}
		log.Println("Crontab Created")
		log.Printf("Crontab Content\n%v\n", globals.CRONTAB_CONTENT)
	} else {
		return
	}

}
