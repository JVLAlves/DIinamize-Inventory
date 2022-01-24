package filehelpers

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	globals "github.com/JVLAlves/Dinamize-Inventory/internal/app/globals"
)

/*
Define a data de hoje na forma Day/Month/Year e retorna um Daytime (Na forma: _Day_Month_Year) para nomear o arquivod e log
*/
func Today() (Daytime string) {

	years, month, day := time.Now().Date()

	Day := strconv.Itoa(day)
	Month := strconv.Itoa(int(month))
	Year := strconv.Itoa(years)

	Daytime = "_" + Day + "_" + Month + "_" + Year

	return Daytime

}

//Cria um Diretório no HOME do usuário
func CreateDir(wg *sync.WaitGroup) {
	HOME, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error getting user home directory path: ", err)
	}
	os.Mkdir(HOME+"/"+globals.LOG_DIR_NAME, 0777)
	wg.Done()
}

//Cria um arquivo de logs
func InitLogs() (f *os.File) {

	var errboolean bool = true
	_, err := os.Stat(os.Getenv("USERNAME") + "_logs")
	if os.IsNotExist(err) {
		errboolean = false
	}
	if err != nil {
		errboolean = false
	}
	_, boolean := os.LookupEnv("HOME")
	USERNAME := os.Getenv("USERNAME")
	if !(boolean && errboolean) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		CreateDir(wg)
		wg.Wait()
	}
	HOME := os.Getenv("HOME")
	HOMELOGS := HOME + "/" + USERNAME + "_logs"

	Daytime := Today()

	logname := HOMELOGS + "/Logs" + Daytime + ".txt"

	outFile, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println("error creating file", err)
	}
	log.SetOutput(outFile)
	return outFile
}
