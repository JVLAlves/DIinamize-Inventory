package Consts

const (
	IP_SNIPEIT       = "10.20.1.79:8001"
	REGEX_MEMORIA_HD = `(^ ?\d{1,3}[,.]?\d*)`
	REGEX_COMMAORNOT = `\s*(\d+[.,]?\d*)`
	//ID relativo ao modelo do ativo (Ex.: Desktop, Notebook, Macbook)
	ID_MODELO = "8"
	ID_STATUS = "7"
	//Nome personalizado do modelo
	MODELO_ATIVO = "DNZ-COMPUTER"
)

var (
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
)
