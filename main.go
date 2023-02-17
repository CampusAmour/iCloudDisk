package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"iCloudDisk/api/testping"
	"iCloudDisk/pkg/config"
	"iCloudDisk/pkg/log"
)

var (
	configPath = flag.String("c", "/home/campusamour/service/icloud-disk/conf/cfg.yaml.yaml",
		"set config path")
	help = flag.Bool("h", false, "this is a help")
	port = flag.String("port", "", "server listen port")
)

func usage() {
	_, errPrint := fmt.Fprintf(os.Stderr,
		`icloud-disk v1.0.0
Usage: icloud-disk [-h]
Options:
`)
	if errPrint != nil {
		fmt.Printf("[ERROR]print usage err: %s\n", errPrint)
	}
	flag.PrintDefaults()
}

type Server struct {
	Port string
}

func registerAPI(router *gin.Engine) {
	testping.Router(router)
}

func parseConfig() error {
	log.Info("icloud disk parse config file, config path: %s", *configPath)
	if err := config.Parse(*configPath); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	log.InitConsoleLog()
	if err := parseConfig(); err != nil {
		log.Error("icloud disk parse config failed, err: %s", err.Error())
	}

	log.InitLog(config.Config().LogCfg)

	var server Server
	server.Port = *port

	router := gin.Default()

	registerAPI(router)

	if err := router.Run(fmt.Sprintf(":%s", server.Port)); err != nil {
		fmt.Printf("run sever failed, error: %s\n", err.Error())
	}
}
