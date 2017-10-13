package main

import "github.com/alievilshat/fzf/src"

var revision string

func main() {
	fzf.Run(fzf.ParseOptions(), revision)
}
