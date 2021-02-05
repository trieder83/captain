package main

import (
	"flag"
	"fmt"
	"github.com/magiconair/properties"
	"log"
	"os"
	"os/exec"
	"strings"
)

func getVersionString(versionPtr *bool, prop *properties.Properties) string {
	cmdVersStrings := strings.Split(prop.GetString("versionGitHash", ""), " ")
	cmdVers := exec.Command(cmdVersStrings[0], cmdVersStrings[1:]...)
	out, err := cmdVers.CombinedOutput()
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	return string(out)
}

func main() {

	// init props
	prop := properties.MustLoadFile("demo.props", properties.UTF8)
	// init flags
	versionPtr := flag.Bool("version", false, "versionFlag")
	flag.Parse()

	fmt.Println("vim-go " + prop.GetString("versionPrefix", "no"))
	fmt.Println("vim-go " + prop.GetString("notest", "no"))

	fmt.Println("vim-go ", *versionPtr)
	if *versionPtr {
		versionString := getVersionString(versionPtr, prop)
		fmt.Println(versionString)
	}

	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// err := cmd.Wait()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error %s\n", err)
	}

}
