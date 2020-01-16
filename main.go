package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// flag.CommandLine.SetOutput(os.Stderr)

	var (
		arg11     string
		arg12     string
		arg21     string
		arg22     string
		printHelp bool
	)

	functions := map[string]string{
		"doSomething1": "Does something 1",
		"doSomething2": "Does something 2",
	}

	flag.Usage = func() {
		message := fmt.Sprintf("Usage: %s <function> <arg1> <arg2> ...\n", os.Args[0])
		message += "Functions:\n"
		for name, desc := range functions {
			message += fmt.Sprintf("\t%s - %s\n", name, desc)
		}
		fmt.Printf(message)
	}

	if len(os.Args) <= 1 {
		fmt.Println("Function name not specified.")
		flag.Usage()
		os.Exit(1)
	}

	funcName := os.Args[1]
	// fmt.Println(funcName)

	funcNameValid := false
	for name := range functions {
		if name == funcName {
			funcNameValid = true
		}
	}

	if !funcNameValid {
		fmt.Println("Function name not valid.")
		flag.Usage()
		os.Exit(1)
	}

	doSomething1Flags := flag.NewFlagSet("doSomething1Args", flag.ContinueOnError)
	doSomething1Flags.StringVar(&arg11, "arg11", "arg11default", "Arg11 is 1st argument of function doSomething1.")
	doSomething1Flags.StringVar(&arg12, "arg12", "arg12default", "Arg12 is 2nd argument of function doSomething1.")
	doSomething1Flags.BoolVar(&printHelp, "help", false, "Print this help.")

	doSomething2Flags := flag.NewFlagSet("doSomething1Args", flag.ContinueOnError)
	doSomething2Flags.StringVar(&arg21, "arg21", "arg21default", "Arg21 is 1st argument of function doSomething2.")
	doSomething2Flags.StringVar(&arg22, "arg22", "arg22default", "Arg22 is 2nd argument of function doSomething2.")
	doSomething2Flags.BoolVar(&printHelp, "help", false, "Print this help.")

	// fmt.Println("arg1 before parsing:", arg1)

	switch funcName {
	case "doSomething1":
		doSomething1Flags.Parse(os.Args[2:])
		fmt.Println("arg11 after parsing:", arg11)
		fmt.Println("arg12 after parsing:", arg12)
		break
	case "doSomething2":
		doSomething2Flags.Parse(os.Args[2:])
		fmt.Println("arg21 after parsing:", arg21)
		fmt.Println("arg22 after parsing:", arg22)
		break
	default:
		break
	}
	// flag.Parse()
}
