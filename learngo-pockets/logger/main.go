package main

import (
	"fmt"
	"os"

	"github.com/manuonda/go-projects/learngo-pockets/logger/pocketlog"
)

func main() {

	fmt.Println("informacion qui ")

	lgr := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOuput(os.Stdout))
	fmt.Println(lgr)
	lgr.Debugf("A little copying is better than a little")
}
