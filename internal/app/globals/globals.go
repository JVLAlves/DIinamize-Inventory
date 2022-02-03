package Consts

const (
	//Nome da aplicação.
	APPNAME = "Dinamize-Inventory"
	//IP do banco de dados Snipe it.
	IP_SNIPEIT = "10.20.1.67:8000"
	//Token de autenticação com Snipe it.
	SNIPEIT_AUTHENTIFICATION_TOKEN = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiMTFjN2U4NzA5MmMxNWVlNjBkYWI0Y2NjNzI1Y2U5YjBjYWM5NDgzYTZmZWYyOGQ3ODFmYzkzNjczNzliMzM2ZDE2YWMxNzIyYWVkZTQ4ZWUiLCJpYXQiOjE2NDMyOTE0NzQsIm5iZiI6MTY0MzI5MTQ3NCwiZXhwIjoyMTE2Njc3MDc0LCJzdWIiOiI0MCIsInNjb3BlcyI6W119.KPi_GZWymX7TbHfn35SYYLVTaTZk_og5MSrkKjt5oMGFXwG6PIol3TcfMJLEQ7be4gwKv04jS_CP3UD6fy5lQkD34BOb8KkL7wWX95EPrxp5o_4RsMJmBc2bBht_h5Cndg0uLe3oQUMxjStAc2lT7r7IHo_MPeKd3cykyuhgrmblMjzSR1bKQGidnvC6xWOfetJzU-bSjo6AiEBtB6V8TIE191_FMFrPDsU8AcOPnbSxf0bGZb1Jz1uHy0oOMXZFIMm3gubmWpl0M179u9D2bLRoNTOvYlXASGESEsdmROufvaZfgGF2ld_1E_0CP5Ys3YMmkxnmU45meqZu_WZtQkJbKCdeAFgPMBLaK2buRt-eub9puTlTRngDhCFCHnaDafL8hD7PBQYEgMGI0bCsaUD8QvcIDUhknFg7CwavTqEYD440aiunUurkqjPWPqPm2yLz8W36myWReZXJXNy-YH9fdjAMWAPFTq2bfi9MCUuoCeAOg5HYcUoJ6I20i5rInV6bPfidr3RGsLYOTxeJ_d2-INOyAIy-hc7MSskqWIBkGVCBh8T0p5RcRopk1F2rqF0Qj2bOyAh_NxiF-kR-WveLbMStsDzdr6HjU3BJ_aeBGjLSk91v9QCLt-gzNd25hgp4leGt54X4ODIAiQFVbI4IpwR_8UjPaoxroHrK4no"
	//Assinatura de regex para idenficar, númericamente, a quantidade de memoria no HD.
	REGEX_MEMORIA_HD = `(^ ?\d{1,3}[,.]?\d*)`
	//Assinatura de regex par identificação de um número com ou sem virgula.
	REGEX_COMMAORNOT = `\s*(\d+[.,]?\d*)`
	//ID relativo ao modelo do ativo (Ex.: Desktop, Notebook, Macbook)
	ID_MODELO = "10"
	//ID relativo ao status do ativo (Ex.: Disponível, Ocupado, Defasado)
	ID_STATUS = "4"
	//Constante de DEV, se verdadeira escreve as responses em JSON no arquivo de Logs. Utilizada para caso as entradas e/ou saídas mudem.
	DEVSHOWJSON = true
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
