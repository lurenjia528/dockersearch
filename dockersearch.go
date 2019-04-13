package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/lurenjia528/dockersearch/repository"
)

var (
	h bool
	r string
	t string
)

func init() {
	flag.BoolVar(&h, "h", false, "this `help`")
	flag.StringVar(&r, "r", "", "Input your `queryRepository` example: armv7/armhf-ubuntu or ubuntu,arm")
	flag.StringVar(&t, "t", "", "Input your `queryTag` example: amd or arm and so on")
	// 改变默认的 Usage
	flag.Usage = usage
}
func main() {

	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(1)
	}
	if r == ""  {
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println("(:  与网速有关,请耐心等待...  :)")

	var queryRepository = r
	var queryTag = t
	originUrl := "https://hub.docker.com/v2/search/repositories/?page=1&query=" + queryRepository + "&page_size=20"
	repository.GetRepository(originUrl, queryRepository, queryTag)
	fmt.Println("(:  end...  :)")
}

func usage() {
	fmt.Fprintf(os.Stderr, `dockersearch version: 0.0.2
Usage: dockersearch [-hrt] 

Options:
`)
	flag.PrintDefaults()
}
