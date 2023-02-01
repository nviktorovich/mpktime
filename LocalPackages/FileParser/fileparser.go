package fileparser

import (
	"os"
	"regexp"
	"strings"
)

func Parse(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, err
}

func CheckMatches(src []byte, pattern string) (bool, error) {
	matched, err := regexp.MatchString(pattern, string(src))
	if err != nil {
		return false, err
	}
	return matched, err
}

func GetMatches(src []byte, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	result := re.FindAllString(string(src), -1)
	return result, nil
}

func GetHosts(ipMatches []string, hosts ...string) []string {
	addrAndIpStr := strings.Split(ipMatches[0], " ")
	ipStr := addrAndIpStr[len(addrAndIpStr)-1]
	netOctets := strings.Split(ipStr, ".")
	hostsList := [3]string{}
	for i, v := range hosts {
		hostsList[i] = strings.Join(netOctets[:3], ".") + "." + v
	}

	return hostsList[:]
}
