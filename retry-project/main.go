package main

import (
	"retry-project/retry"
)

func main() {
	retry.InitDatabase()
	retry.StartCronJob()
}
