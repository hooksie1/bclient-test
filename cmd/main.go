package main

import (
	"github.com/hooksie1/bclient"
)

func main() {
	client := bclient.NewClient()
	client.NewDB("mydb.db")

	b := bclient.NewBucket("test")
	client.Write(b)
}
