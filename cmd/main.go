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

func getVersionString(prop *properties.Properties) string {
	cmdVersStrings := strings.Split(prop.GetString("versionGitHash", ""), " ")
	cmdVers := exec.Command(cmdVersStrings[0], cmdVersStrings[1:]...)
	out, err := cmdVers.CombinedOutput()
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	outnoline := strings.TrimSuffix(string(out), "\n")
	versionString := prop.GetString("versionPrefix", "undefined") + "-" + outnoline + "-" + prop.GetString("versionSuffix", "undefined")
	return versionString
}
func runCmd(cmdKey string, prop *properties.Properties) {
	var cmdString string
	// replace patterns
	cmdString = strings.ReplaceAll(prop.MustGetString(cmdKey), "#VERSIONSTRING#", getVersionString(prop))
	// split for execution
	cmdStrings := strings.Split(cmdString, " ")
	// execute
	cmd := exec.Command(cmdStrings[0], cmdStrings[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// err := cmd.Wait()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error %s\n", err)
	}

}

func main() {

	// init props
	prop := properties.MustLoadFile("demo.props", properties.UTF8)

	// init flags
	versionPtr := flag.Bool("version", false, "version Flag")
	helpPtr := flag.Bool("help", false, "help Flag")
	buildPtr := flag.Bool("build", false, "bulid Flag")
	testPtr := flag.Bool("test", false, "test Flag")
	flag.Parse()

	// flag.Arg(i)
	//fmt.Println(flag.Arg(1))

	if *helpPtr {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *versionPtr {
		fmt.Println(getVersionString(prop))
		os.Exit(0)
	}
	if *buildPtr {
		runCmd("cmd.build", prop)
	}
	if *testPtr {
		runCmd("cmd.test", prop)
	}
	if len(flag.Args()) == 0 {
		// default build
		runCmd("cmd.build", prop)
	}

}
