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

		fmt.Println(hostsList)
		errChan := make(chan error)
		statusChan := make(chan bool)

		for _, ip := range hostsList {
			go cn.Connect(ip, "root", "crtc", "disp", statusChan, errChan)
		}

		go func(ch <-chan error) {
			select {
			case err, ok := <-errChan:
				if ok {
					log.Println(err)

				} else {
					fmt.Println("errChan is closed")
				}
			default:
				fmt.Println("No value ready, moving on.")
			}
		}(errChan)

		cnt := 0
		select {
		case status, ok := <-statusChan:
			if ok {
				fmt.Println(status)
				cnt += 1
				if cnt == 3 {
					close(statusChan)
					close(errChan)
				}

			} else {
				fmt.Println("statusChan is closed")
			}
		default:
			fmt.Println("No value ready, moving on.")
		}

	} else {
		fmt.Println("отсутствуют сетевые настройки")
		os.Exit(1)
	}

}
