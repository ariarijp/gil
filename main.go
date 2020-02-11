package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tsuyoshiwada/go-gitlog"
)

func main() {
	conf := gitlog.Config{}
	if len(os.Args) >= 2 {
		conf.Path = os.Args[1]
	}

	git := gitlog.New(&conf)
	commits, err := git.Log(nil, &gitlog.Params{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, commit := range commits {
		bytes, err := json.Marshal(commit)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s\n", string(bytes))
	}
}
