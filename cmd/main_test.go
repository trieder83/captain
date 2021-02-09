package main

import (
	"regexp"
	"testing"
)

import "github.com/magiconair/properties"

func TestGetVersionString(t *testing.T) {
	// expect
	expectVersionString := "testprefix-.*-testsuffix"
	testProperties := properties.LoadMap(map[string]string{"versionSuffix": "testsuffix", "versionPrefix": "testprefix", "versionGitHash": "git rev-parse --short HEAD"})
	testVersionString := getVersionString(testProperties)

	//if !strings.HasPrefix(testVersionString, expectVersionString) {
	matched, err := regexp.Match(expectVersionString, []byte(testVersionString))
	if !matched {
		t.Errorf("getVersionString failed , expected %v, got %v", expectVersionString, testVersionString)
	}
	if err != nil {
		t.Errorf("getVersionString failed with error, %v", err)
	}

	// expect undefied
	expectVersionString = "undefined-.*-undefined"
	testProperties = properties.LoadMap(map[string]string{"versionGitHash": "git rev-parse --short HEAD"})
	testVersionString = getVersionString(testProperties)
	matched, err = regexp.Match(expectVersionString, []byte(testVersionString))
	if !matched {
		t.Errorf("getVersionString failed , expected %v, got %v", expectVersionString, testVersionString)
	}
	if err != nil {
		t.Errorf("getVersionString failed with error, %v", err)
	}
}
