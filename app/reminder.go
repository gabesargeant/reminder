package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"strings"

)

//Green the color
var green  = "\033[32m"
var reset = "\033[0m"
var red   = "\033[31m"

type reminder struct {
	note     string `json:note`
	datetime string `json:datetime`
	toWhome  string `json:target`
}

type Args struct {
	note     *string
	datetime *string
	to       *string
}

func defineFlags() Args {
	var args = Args{}

	args.to = flag.String("t", "to", "to")
	args.note = flag.String("n", "note", "note")
	args.datetime = flag.String("d", "datetime", "datetime")
	return args

}

func main() {
	fmt.Println("Reminder!")
	args := defineFlags()	
	flag.Parse()
	buildReminder(args)
	setRem := confirmSetReminder()
	if(setRem == false){
		fmt.Println(red + "Reminder Cancelled" + reset);
		os.Exit(-99)
	}

	buildEvent();
	scheduleEvent();
	confirmEvent();
	
}

func confirmSetReminder() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(green + "Confirm set? (Y/N) :" + reset)
	yN, _ := reader.ReadString('\n')
	yN = strings.ToLower(yN)
	yN = strings.TrimSpace(yN)
	
	if yN == "y" {
		fmt.Println("")
		return true
	}	
	return false

}

func buildReminder(args Args) {
	fmt.Println(red +"------------------------------------------------------------" + reset)
	fmt.Println(red + "TO: "+ reset, *args.to)
	fmt.Println(red + "NOTE: "+ reset, *args.note)
	fmt.Println(red + "DATETIME: "+ reset,  *args.datetime)
	fmt.Println(red +"------------------------------------------------------------" + reset + "\n")
}
