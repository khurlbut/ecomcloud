package ecomcloud

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

// InitCloud - initialize the cloud
func InitCloud() {

}

// KubeCtlApply - kubectl apply yaml file
func KubeCtlApply(yaml string) error {
	return kubectl("apply", yaml)
}

// KubeCtlDelete - kubectl delete yaml
func KubeCtlDelete(yaml string) error {
	return kubectl("delete", yaml)
}

// KubeCtlGetNamespace - kubectl get namespace
func KubeCtlGetNamespace(n string) error {
	cmd := exec.Command("kubectl", "get", "namespace", n)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println("cmd returned error: " + err.Error())
		return err
	}
	if strings.HasPrefix(out.String(), "Error") {
		return errors.New("namespace not found")
	}
	return nil
}

func kubectl(command string, yaml string) error {
	log.Printf("kubectl %s -f %s", command, yaml)

	cmd := exec.Command("kubectl", command, "-f", yaml)
	err := cmd.Run()
	if err != nil {
		log.Printf("'kubectl %s -f %s' FAILED! %s", command, yaml, err.Error())
	}

	return err

}
