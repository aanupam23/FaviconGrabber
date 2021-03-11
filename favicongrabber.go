package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

type Item struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}
type List struct {
	List []Item `json:"item"`
}

func main() {
	var list List
	var notfound []string

	dat, err := ioutil.ReadFile("domain.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(dat, &list)
	if err != nil {
		log.Println("Error on UnMarshal.\n[ERROR] -", err)
	}

	for _, element := range list.List {
		url := "https://www.google.com/s2/favicons?sz=64&domain_url=" + element.Domain
		output := element.Name + ".png"
		cmd := exec.Command("curl", url, "-o", output)
		log.Printf("Running command and waiting for it to finish...")
		err := cmd.Run()
		if err != nil {
			notfound = append(notfound, element.Name)
			log.Printf("Command finished with error: %v", err)
		}
	}
	fmt.Println(notfound, len(notfound))
}
