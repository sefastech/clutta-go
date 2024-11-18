package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func VerifyThatAllEnvVariablesAreSet(requiredVariables []string) {

	var missingVariables []string

	for _, variable := range requiredVariables {
		if os.Getenv(variable) == "" {
			missingVariables = append(missingVariables, variable)
		}
	}

	if len(missingVariables) > 0 {
		logrus.Infof("Some environment variables are not set."+
			" Please ensure the following are set: [%s]", strings.Join(missingVariables, ", "))
		os.Exit(1)
	}
}
