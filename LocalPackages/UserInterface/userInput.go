package userinterface

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func GetInput() (string, string) {

	var hh, mm string

	color.Cyan("Введите часы и минуты в формате ЧЧ ММ и нажмите Enter\n")
	fmt.Scanf("%s %s", &hh, &mm)

	if checkHours(hh) && checkMinutes(mm) {
		fmt.Println("вы ввели:", hh, mm)

	} else {
		fmt.Printf("Введены некоректные данные Часы: %s, Минуты: %s\n", hh, mm)
		time.Sleep(time.Second * 20)
		os.Exit(1)
	}
	return hh, mm

}

func checkHours(s string) bool {
	HH, err := strconv.Atoi(s)
	if err != nil {
		log.Print(err)
		return false
	}
	if (HH < 0) || (HH > 24) {
		return false
	}
	return true
}

func checkMinutes(s string) bool {
	MM, err := strconv.Atoi(s)
	if err != nil {
		log.Print(err)
		return false
	}
	if (MM < 0) || (MM > 59) {
		return false
	}
	return true
}
