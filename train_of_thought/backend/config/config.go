package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type envStruct struct {
	DG_API_KEY string
	PB_API_KEY string
}

func SetEnvironmentVariables(env string, path string) error {
	if err := validateEnvName(env); err != nil {
		log.Println(err)
		return err
	}

	filename := fmt.Sprintf("%s/.%s.json", path, env)
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("config: Error reading %s\n", filename)
		return err
	}

	envVariables := envStruct{}
	if err := json.Unmarshal(fileBytes, &envVariables); err != nil {
		log.Printf("config: Error reading json data from %s", filename)
		return err
	}

	os.Setenv("DG_API_KEY", envVariables.DG_API_KEY)
	os.Setenv("PB_API_KEY", envVariables.PB_API_KEY)
	return nil
}

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func validateEnvName(env string) error {
	acceptableEnvValues := [4]string{"local", "dev", "stage", "production"}
	for _, v := range acceptableEnvValues {
		if env == v {
			return nil
		}
	}
	return fmt.Errorf("config: validation failed - env not set. Validate -env= flag")
}
