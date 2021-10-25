package logs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	globals "github.com/JVLAlves/Dinamize-Inventory/cmd"
)

func InitLogs() (stdoutFile *os.File) {
	var DirExists bool = true
	var Home string
	var LogDirPath string

	Home, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln("Error getting user home directory path: ", err)
	}

	LogDirPath = Home + "/" + globals.LOG_DIR_NAME
	_, err = os.Stat(LogDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			DirExists = false
		}
	}

	if !DirExists {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		CreateDir(wg)
		wg.Wait()
	}

	Daytime, DayofClearingDate := Today()
	LogFilename := "Logs" + Daytime + ".txt"
	LogFilePath := LogDirPath + "/" + LogFilename

	stdoutFile, err = os.OpenFile(LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println("error creating file", err)
	}
	log.SetOutput(stdoutFile)

	DayofClearingPath := LogDirPath + "/" + ".DayofClearing.txt"
	DayofClearingFile, err := os.OpenFile(DayofClearingPath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println("error setting the Day of Clearing the logs file: ", err)
	}
	defer DayofClearingFile.Close()

	Databyt, err := ioutil.ReadAll(DayofClearingFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Creating DayofClearing File")
		} else {
			log.Println("Error reading dayofclearing file: ", err)
		}
	}
	DataBody := string(Databyt)

	if DataBody != "" {
		DoCRegex := regexp.MustCompile(`([_[:digit:]]*)[^\s]`)
		DoCrgxed := DoCRegex.FindStringSubmatch(DataBody)
		DataBody = DoCrgxed[0]
	}
	switch {

	case DataBody == "":
		byteSlice := []byte(DayofClearingDate + "\n")
		_, err = DayofClearingFile.Write(byteSlice)

		if err != nil {
			log.Println("Error writing the day of clearing the logs: ", err)
		}
	case DataBody != "" && DataBody == Daytime:

		LegacyLogs()

		err := os.Remove(DayofClearingPath)
		fmt.Println(DayofClearingPath)
		if err != nil {
			log.Println("Error removing Day of Clearing File: ", err)
		}

	case DataBody != "" && Daytime != "" && DataBody != Daytime:
		log.Printf("%#v\t%#v\t%v", DataBody, Daytime, false)
	}

	return stdoutFile
}

func LegacyLogs() {

	year, month, day := time.Now().Date()
	RunDate := strconv.Itoa(day) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(year)

	var FilesinDir []string
	Home, err := os.UserHomeDir()
	LogDirPath := Home + "/" + globals.LOG_DIR_NAME
	LogsLegacyFilePath := LogDirPath + "/" + "." + globals.APPNAME + "-LegacyLogs.txt"
	DayofClearingFilePath := LogDirPath + "/" + ".DayofClearing.txt"
	if err != nil {
		log.Println("Error getting Home directory: ", err)
	}
	files, _ := ioutil.ReadDir(LogDirPath)
	for _, f := range files {
		FilesinDir = append(FilesinDir, f.Name())
	}
	LogsLegacyFile, err := os.OpenFile(LogsLegacyFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		log.Println("error setting legacy logs file: ", err)
	}

	RunDatebyt := []byte(RunDate)

	_, err = LogsLegacyFile.Write(RunDatebyt)

	if err != nil {
		log.Println("Error writing Run Date: ", err)

	}

	defer LogsLegacyFile.Close()

	for _, F := range FilesinDir {
		Fn := LogDirPath + "/" + F
		if Fn == LogsLegacyFilePath || Fn == DayofClearingFilePath {
			continue
		}
		File, err := os.OpenFile(Fn, os.O_RDWR, 0777)
		if err != nil {
			log.Println("error reading file: ", err)
		}
		FnContent := "\n" + Fn + "\n"
		Fnbyt := []byte(FnContent)
		Databyt, err := ioutil.ReadAll(File)
		if err != nil {
			log.Println("Error reading dayofclearing file: ", err)

		}
		_, err = LogsLegacyFile.Write(Fnbyt)

		if err != nil {
			log.Println("Error writing the legacy logs: ", err)
		}
		_, err = LogsLegacyFile.Write(Databyt)

		if err != nil {
			log.Println("Error writing the legacy logs: ", err)
		}

		err = os.Remove(Fn)

		if err != nil {
			log.Println("Removing read file: ", err)
		}

		File.Close()
	}

}

func Today() (Daytime, DayofClearing string) {

	years, month, day := time.Now().Date()
	MonthFuture := time.Now().Month() + 1
	Day := strconv.Itoa(day)
	Month := strconv.Itoa(int(month))
	Year := strconv.Itoa(years)

	Daytime = "_" + Day + "_" + Month + "_" + Year
	DayofClearing = "_" + Day + "_" + strconv.Itoa(int(MonthFuture)) + "_" + Year

	return Daytime, DayofClearing

}

func CreateDir(wg *sync.WaitGroup) {
	HOME, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error getting user home directory path: ", err)
	}
	os.Mkdir(HOME+"/"+globals.LOG_DIR_NAME, 0777)
	wg.Done()
}

/*var errboolean bool = true
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
return outFile*/
