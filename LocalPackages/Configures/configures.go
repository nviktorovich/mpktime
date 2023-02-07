package configures

import "regexp"

// Файл предназначен для хранения переменных, для исключения возможности замены
// или использования пароля и имени пользователя администратора в своих целях,
// принято решение хранить данные в файле, который последствии будет скомпилирован.
const (
	PathToHosts = "/etc/hosts"
	Port        = "22"
	User        = "root"
	Pass        = "crtc"
	CommandDate = "date +%T --set="
)

var (
	Pattern = regexp.MustCompile(`(?m:^(192.168.[0-9]{1,3}.[0-9]{1,3}))`)
)
