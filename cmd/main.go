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
	return strings.TrimSuffix(string(out), "\n")
}

func main() {

	// init props
	prop := properties.MustLoadFile("demo.props", properties.UTF8)

	// init flags
	versionPtr := flag.Bool("version", false, "versionFlag")
	helpPtr := flag.Bool("help", false, "helpFlag")
	flag.Parse()

	// flag.Arg(i)
	if *helpPtr {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *versionPtr {
		versionString := getVersionString(versionPtr, prop)
		//versionString := "6796757"
		fmt.Printf("%v-%v-%v\n", prop.GetString("versionPrefix", "undefined"), versionString, prop.GetString("versionSuffix", "undefined"))
		os.Exit(0)
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
