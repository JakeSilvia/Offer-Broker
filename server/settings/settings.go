package settings

import (
	"os"
	"strings"
)

var (
	PRIVATE_KEY_ID = os.Getenv("PRIVATE_KEY_ID")
	PRIVATE_KEY =  strings.Replace(os.Getenv("PRIVATE_KEY"), "\\n", "\n", -1)
	CLIENT_EMAIL = os.Getenv("CLIENT_EMAIL")
	PORT = os.Getenv("PORT")
	API_KEY = os.Getenv("API_KEY")
)
