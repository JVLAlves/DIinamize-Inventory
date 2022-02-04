package Consts

const (
	//Nome da aplicação.
	APPNAME = "Dinamize-Inventory"
	//IP do banco de dados Snipe it.
<<<<<<< HEAD
	IP_SNIPEIT = "10.20.1.67:8000" //Este Banco de dados é o do Fernando.
	//Token de segurança do SnipeIT
	SNIPEIT_TOKEN = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiN2U4MjQyMjI5MTM0NDMzNDY5YWZhNjM2MWVmNWUyMjc3YWIyMjQ2YmFjNjI4ZWRmZDk0MjcwYTUwYTgyM2M3YmQ3OTY1YzkzYjE3NzhkMWMiLCJpYXQiOjE2NDMwNDc0MzQsIm5iZiI6MTY0MzA0NzQzNCwiZXhwIjoyMTE2NDMzMDM0LCJzdWIiOiI0MCIsInNjb3BlcyI6W119.C8mJY1rO16CIRzMS2x8BCjgt99SofxaNvbUIcR4Z4EA9uuC6b5XNUErZo_Q3SnTesF8kBJbiYB6z_L0zPK1-g8Y5kO5EjoR5tO4jd8LdRwHW4Fjy6UtP0qEsU7qh2mIZzPWpliQussopyOxwLB_YljJlQjrlDGDe7n-nOIYZh9I3r6lOzAE73_DBiQATmyOFMLuzGtF5kQrYl8WxlgzWw_qGsSvlDzaZWBiGLenS1JC63aCe363vjPCIn64AH2YPRzCicu9-h0FGNPmFIpD1u-5vfuOAR8pxZX8E4jqLLWw9aNp37x0hW48Sc_ckfgDDDan8D4fQNYHRyLu0-zmc92Wn9lbhF70eZE4cTq_fuphnF40Yec80KjNk7lPr13npi7-K7L4xSEXtqkBqPg8dn0IpYr6Y-XJ33d7k8bUAGHbxPWAqj6KtZy-HkgoCSwquFr_N-FhF06sN10PQX1jut-RuEowcAvDyHwY_36tQ7SyjlEtCJAuG8ts95kaDzi_PzzrTd0yWrt8Popcov2K2hB6a4c7BsjcEPv0HlC7Gejg1R-1OFb_csp9NUoquHdiaHWeA-x5UpyoHUDXD0rH8bKThHGEeuJFGJcgW2_iqLsAiJdMIOlZjBrEK_OYKWDyT0WUbefUq42BfqehjRz_Fd8-hWVcZ1wJSAvyYTfR6o3Q"
=======
	IP_SNIPEIT = "10.20.1.67:8000"
	//Token de autenticação com Snipe it.
	SNIPEIT_AUTHENTIFICATION_TOKEN = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiMTFjN2U4NzA5MmMxNWVlNjBkYWI0Y2NjNzI1Y2U5YjBjYWM5NDgzYTZmZWYyOGQ3ODFmYzkzNjczNzliMzM2ZDE2YWMxNzIyYWVkZTQ4ZWUiLCJpYXQiOjE2NDMyOTE0NzQsIm5iZiI6MTY0MzI5MTQ3NCwiZXhwIjoyMTE2Njc3MDc0LCJzdWIiOiI0MCIsInNjb3BlcyI6W119.KPi_GZWymX7TbHfn35SYYLVTaTZk_og5MSrkKjt5oMGFXwG6PIol3TcfMJLEQ7be4gwKv04jS_CP3UD6fy5lQkD34BOb8KkL7wWX95EPrxp5o_4RsMJmBc2bBht_h5Cndg0uLe3oQUMxjStAc2lT7r7IHo_MPeKd3cykyuhgrmblMjzSR1bKQGidnvC6xWOfetJzU-bSjo6AiEBtB6V8TIE191_FMFrPDsU8AcOPnbSxf0bGZb1Jz1uHy0oOMXZFIMm3gubmWpl0M179u9D2bLRoNTOvYlXASGESEsdmROufvaZfgGF2ld_1E_0CP5Ys3YMmkxnmU45meqZu_WZtQkJbKCdeAFgPMBLaK2buRt-eub9puTlTRngDhCFCHnaDafL8hD7PBQYEgMGI0bCsaUD8QvcIDUhknFg7CwavTqEYD440aiunUurkqjPWPqPm2yLz8W36myWReZXJXNy-YH9fdjAMWAPFTq2bfi9MCUuoCeAOg5HYcUoJ6I20i5rInV6bPfidr3RGsLYOTxeJ_d2-INOyAIy-hc7MSskqWIBkGVCBh8T0p5RcRopk1F2rqF0Qj2bOyAh_NxiF-kR-WveLbMStsDzdr6HjU3BJ_aeBGjLSk91v9QCLt-gzNd25hgp4leGt54X4ODIAiQFVbI4IpwR_8UjPaoxroHrK4no"
>>>>>>> cf0bf1667f0ca139d83a2c11ba4c99bad8188cb4
	//Assinatura de regex para idenficar, númericamente, a quantidade de memoria no HD.
	REGEX_MEMORIA_HD = `(^ ?\d{1,3}[,.]?\d*)`
	//Assinatura de regex par identificação de um número com ou sem virgula.
	REGEX_COMMAORNOT = `\s*(\d+[.,]?\d*)`
	//ID relativo ao modelo do ativo (Ex.: Desktop, Notebook, Macbook)
	ID_MODELO = "10"
	//ID relativo ao status do ativo (Ex.: Disponível, Ocupado, Defasado)
<<<<<<< HEAD
	ID_STATUS = "5"
=======
	ID_STATUS = "4"
	//Constante de DEV, se verdadeira escreve as responses em JSON no arquivo de Logs. Utilizada para caso as entradas e/ou saídas mudem.
	DEVSHOWJSON = true
>>>>>>> cf0bf1667f0ca139d83a2c11ba4c99bad8188cb4
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
	//Constante de visualização do formato dos JSONs
	DEVVIEW = false
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
