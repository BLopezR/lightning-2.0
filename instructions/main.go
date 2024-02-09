package main

import (
	"errors"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
)

// Script that takes two required arguments:
// the first one is the name in the cluster of the node where the script is running
// the second one is the path to the configuration file, in reference to the code.
func main() {

	controllerIP := flag.String("controller_ip", "", "ip where the SDN controller is listening using the OpenFlow13 protocol. Required")

	fmt.Println("initializing switch, connected to controller: ", controllerIP)
	err := initializeSwitch(*controllerIP)

	if err != nil {
		fmt.Println("Could not initialize switch. Error:", err)
		return
	}

	fmt.Println("Switch initialized and connected to the controller.")

	// Set all virtual interfaces up, and connect them to the tunnel bridge:

	exec.Command("ovs-vsctl", "add-port", "brtun", "eth0").Run()
	exec.Command("ovs-vsctl", "add-port", "brtun", "eth1").Run()

}

func initializeSwitch(controllerIP string) error {

	re := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
	if !re.MatchString(controllerIP) {
		out, _ := exec.Command("host", controllerIP).Output()
		controllerIP = re.FindString(string(out))
	}

	var err error

	err = exec.Command("ovs-vsctl", "add-br", "brtun").Run()

	if err != nil {
		return errors.New("Could not create brtun interface")
	}

	err = exec.Command("ip", "link", "set", "brtun", "up").Run()

	if err != nil {
		return errors.New("Could not set brtun interface up")
	}

	err = exec.Command("ovs-vsctl", "set", "bridge", "brtun", "protocols=OpenFlow13").Run()

	if err != nil {
		return errors.New("Couldnt set brtun messaing protocol to OpenFlow13")
	}

	target := fmt.Sprintf("tcp:%s:6633", controllerIP)

	err = exec.Command("ovs-vsctl", "set-controller", "brtun", target).Run()

	if err != nil {
		return errors.New("Could not connect to controller")
	}
	return nil
}
