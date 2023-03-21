package golocale

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

var (
	HeaderKey string
	Locale    string
	Source    = make(map[string]map[string]string)
)

func Init(locale string, header string, source []string) {

	// Validate Header
	if header == "" {
		panic("Header is required")
	}

	HeaderKey = header
	Locale = locale

	for _, path := range source {
		re := regexp.MustCompile(`[\\/]([^\\/]+)\.\w+$`)
		localeName := re.FindStringSubmatch(path)[1]

		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		var message map[string]string
		if err := decoder.Decode(&message); err != nil {
			fmt.Println(err)
		}

		Source[localeName] = message
	}

	fmt.Println("[✓] System golocale initialized")
}
