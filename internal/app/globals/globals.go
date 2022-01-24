package Consts

const (
	//Nome da aplicação.
	APPNAME = "Dinamize-Inventory"
	//IP do banco de dados Snipe it.
	IP_SNIPEIT = "10.20.1.79:8001" //Este Banco de dados era o do Andreo.
	//Assinatura de regex para idenficar, númericamente, a quantidade de memoria no HD.
	REGEX_MEMORIA_HD = `(^ ?\d{1,3}[,.]?\d*)`
	//Assinatura de regex par identificação de um número com ou sem virgula.
	REGEX_COMMAORNOT = `\s*(\d+[.,]?\d*)`
	//ID relativo ao modelo do ativo (Ex.: Desktop, Notebook, Macbook)
	ID_MODELO = "8"
	//ID relativo ao status do ativo (Ex.: Disponível, Ocupado, Defasado)
	ID_STATUS = "7"
	//Nome personalizado do modelo.
	MODELO_ATIVO = "DNZ-COMPUTER"
	//Periodo da rotina Crontab (Linux).
	CRONTAB_PERIOD = `0 */5 * * * $HOME/` + LINUX_EXECNAME
	//Nome do executável de Linux.
	LINUX_EXECNAME = "dnz-inventory-li"
	//Nome do arquivo que verifica a existencia do Crontab.
	CRONTABEXISTS_FILENAME = ".CrontabExists.txt"
	//Conteúdo de segurança no interior do arquivo de Existência do Crontab.
	CRONTABEXISTS_CONTENT = "'Dont delete this file. It is a register for the application named " + APPNAME + "." + " Deleting it may cause some problems in the application running.'"
	//Conteúdo presente no interior da rotina Crontab.
	CRONTAB_CONTENT = "'" + `#!/bin/bash` + "\n" + CRONTAB_PERIOD + "'"
	//Nome do diretório de Logs
	LOG_DIR_NAME = APPNAME + "-Logs"
)

var (
	//Lista de versões do MacOS ou Darwin
	MacOSVersions = map[string]string{

		"10.7":  "MacOs Lion",
		"10.8":  "MacOs Mountain Lion",
		"10.9":  "MacOs Mavericks",
		"10.10": "MacOs Yosemite",
		"10.11": "MacOs El Capitan",
		"10.12": "MacOs Sierra",
		"10.13": "MacOs High Sierra",
		"10.14": "MacOs Mojave",
		"10.15": "MacOs Catalina",
		"11.4":  "MacOs Big Sur",
	}

	//Lista de softwares básicos que compoem o microsoft Office.
	Office = []string{"Microsoft OneNote", "Microsoft Word", "Microsoft PowerPoint", "Microsoft Excel", "Microsoft Outlook", "Microsoft Onedrive", "Microsoft Publisher", "Microsoft Access"}
)
