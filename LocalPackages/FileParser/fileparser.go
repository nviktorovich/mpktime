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

// GetMatches принмает на вход последовательность байтов и паттерн, осуществляет
// поиск совпадений между паттерном и строчным значением последовательности байтов,
// возвращает слайс со строками, в которых есть совпадения по паттерну. В случае
// возникновения ошибки, возвращается ошибка
func GetMatches(src []byte, pattern *regexp.Regexp) []string {
	result := pattern.FindAllString(string(src), -1)
	return result
}
