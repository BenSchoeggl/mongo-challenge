package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BenSchoeggl/mongo-challenge/jsonutils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	jsonBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal("could not read json bytes, err: ", err)
	}
	data := map[string]interface{}{}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.Fatal("could not unmarshal input json, err: ", err)
	}
	flattenedData := jsonutils.Flatten(data)
	flattenedBytes, err := json.MarshalIndent(flattenedData, "", "    ")
	if err != nil {
		log.Fatal("could not marshal flattened data, err: ", err)
	}
	fmt.Println(string(flattenedBytes))
}
