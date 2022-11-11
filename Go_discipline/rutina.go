package main

import (
	"bufio"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io"
	"log"
	"os"
	"time"
)

func sound(f io.ReadCloser) {

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	//defer f.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true

	})))
	//fmt.Println("shit")

	<-done
	//defer close(done)
}

func main() {
	arr := []string{"badoo", "inst", "tg", "tind"}
	var i int
	in := bufio.NewScanner(os.Stdin)

	fmt.Println("LET'S START TO WORK ^__^\nPress any key to start or type 'exit' to fuck off")
	for true {
		if in.Scan() {
			if in.Text() == "exit" {
				break
			}
			dt := time.Now()
			fmt.Println("Work slot started at ", dt.Format("Jan _2 15:04:05"))
			fmt.Println("Focus for 57 min for fuck sake.")
			time.Sleep(time.Minute * 57)
			o, err := os.Open("./pet.mp3")
			if err != nil {
				log.Fatal(err)
			}
			dt2 := time.Now()
			fmt.Println("\nCongrags\nYou can now take a shit or fuck around on ", arr[i])
			fmt.Println("Chillout slot started at ", dt2.Format("Jan _2 15:04:05"), "\nYou got 3 mins.")
			sound(o)
			o.Close()
			time.Sleep(time.Minute * 3)

			f, err := os.Open("./oblom.mp3")
			if err != nil {
				log.Fatal(err)
			}
			sound(f)
			f.Close()
			fmt.Println("BACK TO WORK BITCH\nPress any key for confirmation or type 'exit' if you're done.")
			i++
			if i == 3 {
				i = 0
			}
		}
	}
}
