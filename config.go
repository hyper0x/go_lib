package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	Dict map[string]string
}

var myConfig Config
var loadingCount = 0
var loadingChan = make(chan int)

func init() {
	go func() {
	L:
		for {
			select {
			case incr, ready := <-loadingChan:
				if !ready {
					break L
				}
				loadingCount += incr
			}
		}
	}()
}

func ReadConfig(fresh bool) (Config, error) {
	needLoad := fresh || (loadingCount == 0)
	if !needLoad {
		return myConfig, nil
	}

	myConfig = *new(Config)
	currentDir, err := os.Getwd()
	configFilePath := currentDir + "/" + CONFIG_FILE_NAME
	myConfig.Dict = make(map[string]string)
	configFile, err := os.OpenFile(configFilePath, os.O_RDONLY, 0666)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			var warningBuffer bytes.Buffer
			warningBuffer.WriteString("Warning: the config file '")
			warningBuffer.WriteString(configFilePath)
			warningBuffer.WriteString("' is NOT FOUND! ")
			warningBuffer.WriteString("Use DEFAULT config '")
			warningBuffer.WriteString(fmt.Sprintf("%v", myConfig))
			warningBuffer.WriteString("'. ")
			LogWarnln(warningBuffer.String())
			return myConfig, nil
		default:
			return myConfig, err
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
				return myConfig, err
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
		myConfig.Dict[key] = value
	}
	loadingChan <- 1
	LogInfof("Loaded config file (count=%d): %v\n", loadingCount, myConfig)
	return myConfig, nil
}
