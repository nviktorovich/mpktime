package connection

import (
	"fmt"
	"net"

	sh "github.com/helloyi/go-sshclient"
)

// ConnectionOperator принимает список хостов, порт, команду, логин и пароль
// для реализации на удаленной машине создает канал для опроса состояния хостов и запускает
// несколько параллельных процессов
func ConnectionOperator(hosts, cmd []string, port, login, pswd string) {

	statusChan := make(chan string, len(hosts))

	for _, ip := range hosts {
		go sshConnection(ip, port, login, pswd, cmd, statusChan)
	}

	for a := 0; a < len(hosts); a++ {
		status := <-statusChan
		fmt.Println(status)
	}
	defer close(statusChan)

}

// sshConnection предназначена для реализации команд на удаленном хосте
func sshConnection(ip, port, login, pass string, commands []string, statChan chan<- string) {

	ipWithPort := fmt.Sprintf("%s:%s", ip, port)
	_, err := net.Dial("tcp", ipWithPort)
	if err != nil {
		statChan <- fmt.Sprintf("неудача на хосте: %s\n%s", ipWithPort, err)
		return
	}

	fmt.Printf("Запуск работы программы на хосте: %s\n", ipWithPort)

	client, err := sh.DialWithPasswd(ipWithPort, login, pass)
	if err != nil {
		err := fmt.Errorf("ошибка установки соединения IP: %s.\nerr(%s)", ipWithPort, err)
		statChan <- fmt.Sprintf("сообщение об ошибке: %s", err)
	}

	defer client.Close()

	// run one command
	for _, command := range commands {
		out, err := client.Cmd(command).SmartOutput()
		if err != nil {
			err = fmt.Errorf("ошибка выполнения команды: %s на хосте %s\nОтвет хоста (out): %s.\nerr(%s)", command, ipWithPort, out, err)
			statChan <- fmt.Sprintf("сообщение об ошибке: %s", err)
		}
		// the 'out' is stdout output

		fmt.Println(string(out))
	}

	statChan <- fmt.Sprintf("работа на хосте: %s завершена\n", ipWithPort)
}
