package configs

import (
	"fmt"
	"os"
	"strconv"
)

func Config(key string) string {
	//if err := godotenv.Load(); err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	return os.Getenv(key)
}

// func Config expects variadic args. Make sure have passed first parameter
// for key and second parameter for optional value
func ConfigWithOptional(keys ...string) (string, bool) {

	ok := false

	if len(keys) > 2 || len(keys) < 1 {
		return "DEFAULT_VALUE", false
	}

	elementMap := make(map[string]string)
	for i := 0; i < len(keys); i++ {
		fmt.Println(keys[i])
		elementMap[strconv.Itoa(i)] = keys[i]
	}

	key, ok := elementMap["0"]

	defaultValue := "DEFAULT_VALUE"
	defaultValue, ok = elementMap["1"]

	keyValue := os.Getenv(key)

	if keyValue == "" && !ok {
		return defaultValue, false
	}

	if keyValue == "" {
		return defaultValue, true
	}

	return keyValue, true
}
