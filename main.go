package main

import (
	"fmt"
	"log"
	"os"
	"time"

	cfg "github.com/NViktorovich/mpktime/LocalPackages/Configures"
	cn "github.com/NViktorovich/mpktime/LocalPackages/Connection"
	fp "github.com/NViktorovich/mpktime/LocalPackages/FileParser"
	ui "github.com/NViktorovich/mpktime/LocalPackages/UserInterface"
)

func main() {

	hh, mm := ui.GetInput()

	newTime := fmt.Sprintf("'%s:%s:00'", hh, mm)
	CommandDate := fmt.Sprint(cfg.CommandDate + newTime)
	commands := []string{CommandDate}

	data, err := fp.Parse(cfg.PathToHosts)
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
		hostsList, err := fp.GetMatches(data, cfg.Pattern)
		if err != nil {
			fmt.Println("ошибка извлечения значения сетевых настроек")
			log.Fatalln(err)
		}

		cn.ConnectionOperator(hostsList, commands, cfg.Port, cfg.User, cfg.Pass)

	} else {
		fmt.Println("отсутствуют сетевые настройки")
		os.Exit(1)
	}
	time.Sleep(20 * time.Second)

}
