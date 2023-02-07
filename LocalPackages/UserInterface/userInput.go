package userinterface

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/martinlindhe/inputbox"
)

// func GetInput() (string, string) {

// 	var hh, mm string
// 	color.Magenta("'ksa' корректировка времени.\n\n")
// 	color.Cyan("Введите часы и минуты в формате ЧЧ ММ и нажмите Enter\n")
// 	fmt.Scanf("%s %s", &hh, &mm)

// 	if checkHours(hh) && checkMinutes(mm) {
// 		fmt.Println("вы ввели:", hh, mm)

// 	} else {
// 		fmt.Printf("Введены некоректные данные Часы: %s, Минуты: %s\n", hh, mm)
// 		time.Sleep(time.Second * 20)
// 		os.Exit(1)
// 	}
// 	return hh, mm

// }
// checkHours осуществляет проверку валидности введенных часов,
// возвращает true если проверка пройдена, в противном случае false
func checkHours(s string) bool {
	HH, err := strconv.Atoi(s)
	if err != nil {
		log.Print(err)
		return false
	}
	if (HH < 0) || (HH > 23) {
		return false
	}
	return true
}

// checkMinutes осуществляет проверку валидности введенных минут,
// возвращает true если проверка пройдена, в противном случае false
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

// UserDialog реализует пользовательский интерфейс, возврашает валидные часы и минуты.
// В случае, если проверки не пройдены, завершает работу программы
func UserDialog() (string, string) {
	var HH, MM string = "", ""
	got, ok := inputbox.InputBox("настройка времени", "введите время в формате ЧЧ ММ", "12 00")
	if ok {
		userInput := strings.Split(got, " ")
		if len(userInput) == 2 {
			HH = userInput[0]
			MM = userInput[1]
			color.Magenta("'ksa' корректировка времени.\n\n")
			color.Cyan("Введите часы и минуты в формате ЧЧ ММ и нажмите Enter\n")

			if checkHours(HH) && checkMinutes(MM) {
				fmt.Println("вы ввели:", HH, MM)

			} else {
				fmt.Printf("Введены некоректные данные Часы: %s, Минуты: %s\n", HH, MM)
				time.Sleep(time.Second * 20)
				os.Exit(1)
			}

		} else {
			fmt.Printf("Введены некоректные данные Часы: %s, Минуты: %s\n", HH, MM)
			time.Sleep(time.Second * 20)
			os.Exit(1)
		}

	} else {
		fmt.Printf("Введены некоректные данные Часы: %s, Минуты: %s\n", HH, MM)
		time.Sleep(time.Second * 20)
		os.Exit(1)
	}

	return HH, MM
}
