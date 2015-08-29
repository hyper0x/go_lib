package go_lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"

	"github.com/hyper-carrot/go_lib/logging"
)

var logger logging.Logger = logging.GetSimpleLogger()

type Config struct {
	Path         string
	Dict         map[string]string
	sign         *Sign
	loadingCount int
}

func (self *Config) ReadConfig(fresh bool) error {
	if self.sign == nil {
		self.sign = NewSign()
	}
	self.sign.Set()
	defer func() {
		self.sign.Unset()
		if err := recover(); err != nil {
			debug.PrintStack()
			errorMsg := fmt.Sprintf("Occur FATAL error when read config (path=%v): %s", self.Path, err)
			logger.Fatalln(errorMsg)
		}
	}()
	needLoad := fresh || (self.loadingCount == 0)
	if !needLoad {
		return nil
	}
	var configFilePath string
	if !strings.Contains(configFilePath, "/") {
		currentDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		configFilePath = currentDir + "/" + self.Path
	} else {
		configFilePath = self.Path
	}
	if self.Dict == nil {
		self.Dict = make(map[string]string)
	}
	configFile, err := os.OpenFile(configFilePath, os.O_RDONLY, 0666)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			var warningBuffer bytes.Buffer
			warningBuffer.WriteString("Warning: the config file '")
			warningBuffer.WriteString(configFilePath)
			warningBuffer.WriteString("' is NOT FOUND! ")
			warningBuffer.WriteString("Use DEFAULT config '")
			warningBuffer.WriteString(fmt.Sprintf("%v", self.Dict))
			warningBuffer.WriteString("'. ")
			logger.Warnln(warningBuffer.String())
			return nil
		default:
			panic(err)
		}
	}
	defer configFile.Close()
	configReader := bufio.NewReader(configFile)
	for {
		str, err := configReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// The file end is touched.
				break
			} else {
				panic(err)
			}
		}
		str = strings.TrimRight(str, "\r\n")
		if len(str) == 0 {
			continue
		}
		sepIndex := strings.Index(str, "=")
		if sepIndex <= 0 || sepIndex == (len(str)-1) {
			continue
		}
		key := strings.ToLower(str[0:sepIndex])
		value := str[sepIndex+1 : len(str)]
		self.Dict[key] = value
	}
	logger.Infof("Loaded config file (count=%d).", self.loadingCount)
	return nil
}
