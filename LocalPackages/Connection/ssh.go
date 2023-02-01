package connection

import (
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/ssh"

	"google.golang.org/grpc/codes"

	expect "github.com/google/goexpect"
	"github.com/google/goterm/term"
)

func Connect(ip, login, pswd1, pswd2 string, statusChan chan<- bool, errChan chan<- error) {

	fmt.Println(term.Bluef("SSH to: %s", ip))

	sshClt, err := ssh.Dial("tcp", ip, &ssh.ClientConfig{
		User:            login,
		Auth:            []ssh.AuthMethod{ssh.Password(pswd1)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		errChan <- fmt.Errorf("ssh.Dial(%q) failed: %s", ip, err)
	}
	defer sshClt.Close()

	e, _, err := expect.SpawnSSH(sshClt, time.Millisecond*200)
	if err != nil {
		statusChan <- false
		errChan <- fmt.Errorf("ssh.Dial(%q) failed: %s", ip, err)
	}
	defer e.Close()

	e.ExpectBatch([]expect.Batcher{
		&expect.BCas{[]expect.Caser{
			&expect.Case{R: regexp.MustCompile(`router#`), T: expect.OK()},
			&expect.Case{R: regexp.MustCompile(`Login: `), S: login,
				T: expect.Continue(expect.NewStatus(codes.PermissionDenied, "wrong username")), Rt: 3},
			&expect.Case{R: regexp.MustCompile(`Password: `), S: pswd1, T: expect.Next(), Rt: 1},
			&expect.Case{R: regexp.MustCompile(`Password: `), S: pswd2,
				T: expect.Continue(expect.NewStatus(codes.PermissionDenied, "wrong password")), Rt: 1},
		}},
	}, time.Millisecond*200)
	fmt.Println(term.Greenf("All done"))
	e.Send("reboot\n")
	statusChan <- true
}
