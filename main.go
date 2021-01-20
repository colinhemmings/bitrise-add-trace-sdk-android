package main

import (
	"fmt"
	"github.com/bitrise-io/go-utils/log"
	"os"
	"os/exec"
)

// Called when the main function should be terminated with failure.
func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	os.Exit(1)
}

func main() {
	if err := createConfigurationFile(); err != nil {
		failf("Could not create the config file, aborting build. Reason: %s\n", err)
	}

	//
	// --- Step Outputs: Export Environment Variables for other Steps:
	// You can export Environment Variables for other Steps with
	//  envman, which is automatically installed by `bitrise setup`.
	// A very simple example:
	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "EXAMPLE_STEP_OUTPUT", "--value", "the value you want to share").CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}
	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}

// Creates the configuration file for the given Android project. The configuration file has the required properties for
// building the Android application.
func createConfigurationFile() error {
	c, err := getConfigFileContent()
	if err != nil {
		return err
	}

	fc, err := formatConfigFileContent(c)
	if err != nil {
		return err
	}

	p := fmt.Sprint(os.Getenv(srcDirEnvName), "/", configFileName)
	return createConfigFile(fc, p)
}