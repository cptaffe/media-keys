package main

import (
	"os"
	"fmt"
	"github.com/cptaffe/mpris2"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf(os.Args[0] + " [play | pause | playpause]\n")
	} else {
		player, err := mpris2.NewPlayer("spotify")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		switch os.Args[1] {
		case "playpause":
			player.PlayPause()
		case "next":
			player.Next()
		case "prev":
			player.Previous()
		default:
			fmt.Printf("%s is currently unsupported\n", os.Args[1])
		}
	}
}
