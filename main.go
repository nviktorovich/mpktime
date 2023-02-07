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

	// commands список в который можно добавлять команды для исполнения
	commands := []string{CommandDate}

	data, err := fp.Parse(cfg.PathToHosts)
	if err != nil {
		fmt.Println("ошибка чтения конфигурационного файла")
		log.Fatalln(err)
	}

	hostsList := fp.GetMatches(data, cfg.Pattern)

	if len(hostsList) > 0 {

		cn.ConnectionOperator(hostsList, commands, cfg.Port, cfg.User, cfg.Pass)

	} else {
		fmt.Println("отсутствуют хосты в файле: ", cfg.PathToHosts)
		os.Exit(1)
	}
	time.Sleep(20 * time.Second)

}
