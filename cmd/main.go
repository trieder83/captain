package main

import (
	"fmt"
	"github.com/magiconair/properties"
	"log"
	"os"
	"os/exec"
)

func main() {

	// init
	prop := properties.MustLoadFile("demo.props", properties.UTF8)

	fmt.Println("vim-go " + prop.GetString("versionPrefix", "no"))
	fmt.Println("vim-go " + prop.GetString("notest", "no"))

	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error %s\n", err)
	}

}
