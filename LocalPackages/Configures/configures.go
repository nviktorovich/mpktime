package configures

// Файл предназначен для хранения переменных, для исключения возможности замены
// или использования пароля и имени пользователя администратора в своих целях,
// принято решение хранить данные в файле, который последствии будет скомпилирован.
var (
	Pattern = "192.168.[0-9]{1,3}.[0-9]{1,3}"
	// PathToInterfaces = "/etc/network/interfaces"
	PathToHosts = "/etc/hosts"
	Port        = "22"
	DspA        = "201"
	DspB        = "202"
	ShN         = "203"
	User        = "root"
	Pass        = "crtc"
	CommandDate = "date +%T --set="

	// CommandDate      = "date +%T --set='%s:%s:00'"
)
