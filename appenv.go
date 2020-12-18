package appenv

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type AppEnv struct {
	envfile string
	sysenv  bool
	envs    map[string]string
}

func NewAppEnv(envfile string, sysenv bool) *AppEnv {
	appenv := &AppEnv{
		envfile: envfile,
		sysenv:  sysenv,
		envs:    make(map[string]string),
	}

	if appenv.sysenv {
		appenv.setSystemEnvs()
	}
	appenv.setEnvs()
	return appenv
}

func (ae *AppEnv) GetEnvs() map[string]string {
	return ae.envs
}

func (ae *AppEnv) GetEnv(key string) string {
	return ae.envs[key]
}

func (ae *AppEnv) SetEnv(key, value string) (string, error) {
	err := os.Setenv(key, value)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (ae *AppEnv) addItemToEnv(item string) {
	s := strings.Split(item, "=")
	if len(s) == 2 {
		ae.envs[s[0]] = s[1]
	}
}

func (ae *AppEnv) setSystemEnvs() {
	envs := os.Environ()
	for _, i := range envs {
		ae.addItemToEnv(i)
	}
}

func (ae *AppEnv) setEnvs() {
	if _, err := os.Stat(ae.envfile); os.IsNotExist(err) {
		log.Println("Env File Not Found:", ae.envfile)
	} else {
		file, err := os.Open(ae.envfile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		for _, ln := range lines {
			ae.addItemToEnv(ln)
		}
	}
}
