package logs

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	globals "github.com/JVLAlves/Dinamize-Inventory/internal/app/globals"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Slogger *zap.SugaredLogger

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

	Daytime := Today()
	LogFilename := "Logs" + Daytime + ".txt"
	LogFilePath := LogDirPath + "/" + LogFilename

	stdoutFile, err = os.OpenFile(LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println("error creating file", err)
	}
	log.SetOutput(stdoutFile)

	return stdoutFile
}

func Today() (Daytime string) {

	years, month, day := time.Now().Date()
	Day := strconv.Itoa(day)
	Month := strconv.Itoa(int(month))
	Year := strconv.Itoa(years)

	Daytime = "_" + Day + "_" + Month + "_" + Year

	return Daytime

}

func CreateDir(wg *sync.WaitGroup) {
	HOME, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error getting user home directory path: ", err)
	}
	os.Mkdir(HOME+"/"+globals.LOG_DIR_NAME, 0777)
	wg.Done()
}

func InitLogger() {
	writeSyncer := getlogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	Slogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getlogWriter() zapcore.WriteSyncer {

	Home, err := os.UserHomeDir()
	LogDirPath := Home + "/" + globals.LOG_DIR_NAME
	Daytime := Today()
	LogFilename := "Logs" + Daytime + ".txt"
	LogFilePath := LogDirPath + "/" + LogFilename

	if err != nil {
		log.Fatalln("Error getting user home directory path: ", err)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   LogFilePath,
		MaxSize:    4,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
