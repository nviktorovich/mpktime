package main

import (
	"fmt"
	"log"
	"os"

	cfg "github.com/NViktorovich/mpktime/LocalPackages/Configures"
	cn "github.com/NViktorovich/mpktime/LocalPackages/Connection"
	fp "github.com/NViktorovich/mpktime/LocalPackages/FileParser"
)

func main() {
	data, err := fp.Parse(cfg.PathToInterfaces)
	if err != nil {
		fmt.Println("ошибка чтения конфигурационного файла")
		log.Fatalln(err)
	}

	matched, err := fp.CheckMatches(data, cfg.Pattern)
	if err != nil {
		fmt.Println("ошибка, в конфигурационном файле не обнаружены настройки сети")
		log.Fatalln(err)
	}

	if matched {
		match, err := fp.GetMatches(data, cfg.Pattern)
		if err != nil {
			fmt.Println("ошибка извлечения значения сетевых настроек")
			log.Fatalln(err)
		}

		hostsList := fp.GetHosts(match, cfg.DspA, cfg.DspB, cfg.ShN)
		cn.ConnectionOperator(hostsList, cfg.Port, cfg.CommandDate, cfg.User, cfg.Pass)

	} else {
		fmt.Println("отсутствуют сетевые настройки")
		os.Exit(1)
	}

}
