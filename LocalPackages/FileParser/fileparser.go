package fileparser

import (
	"os"
	"regexp"
)

// Parse функция, которая принимает на вход путь к файлу с конфигурацией
// и возвращает его содержимое, в случае, если были ошибки, возвращается ошибка
func Parse(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, err
}

// CheckMatches принмает на вход последовательность байтов и паттерн, осуществляет
// поиск совпадений между паттерном и строчным значением последовательности байтов,
// возвращает true в случае наличия совпадений, иначе false. в случае возникновения
// ошибки возвращает ошибку.
func CheckMatches(src []byte, pattern string) (bool, error) {
	matched, err := regexp.MatchString(pattern, string(src))
	if err != nil {
		return false, err
	}
	return matched, err
}

// GetMatches принмает на вход последовательность байтов и паттерн, осуществляет
// поиск совпадений между паттерном и строчным значением последовательности байтов,
// возвращает слайс со строками, в которых есть совпадения по паттерну. В случае
// возникновения ошибки, возвращается ошибка
func GetMatches(src []byte, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	result := re.FindAllString(string(src), -1)
	return result, nil
}

// GetHosts принмает на вход список с совпадениями из функции GetMatches и произвольное
// количество строковых параметров, функция реализует следующую задачу: вытаскивает
// необходимую строку из списка ipMatches, разбивает эту строку, отбрасывает последний
// октет и собираета новый список путем склеивания между собой префикса и параметров,
// которые были отправлени в списке hosts
// func GetHosts(ipMatches []string, hosts ...string) []string {
// 	addrAndIpStr := strings.Split(ipMatches[0], " ")
// 	ipStr := addrAndIpStr[len(addrAndIpStr)-1]
// 	netOctets := strings.Split(ipStr, ".")
// 	hostsList := [3]string{}
// 	for i, v := range hosts {
// 		hostsList[i] = strings.Join(netOctets[:3], ".") + "." + v
// 	}

// 	return hostsList[:]
// }
