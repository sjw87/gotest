package main

import "fmt"
import "io/ioutil"
import "regexp"
import "os"
import "errors"

var exp string
var replace string
var file string
var rep_opt int = 0

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parse_args(args []string) error {
	for i := 0; i < len(args); i++ {
		switch args[i] {
			case "-e":
				i++
				exp = args[i]
			case "-r":
				i++
				replace = args[i]
			case "-i":
				rep_opt = 1
			default :
				file = args[i]
		}
	}
	
	if exp == "" || file == "" {
		return errors.New("Invalid argument")
	}
	
	return nil
}

func print_usage() {
	fmt.Println("Invalid arguments")
	fmt.Println(" usage : replace.exe {-e Regular Expressions} {-r Replace String} [-i] FILE")
	fmt.Println("   example) replace.exe -e http -r https -i urllist.txt ")
}

func main() {
	args := os.Args[1:]
	
	if len(args) < 1{
		print_usage()
		return
	} else {
		err := parse_args(args)
		if err != nil {
			print_usage()
			return
		}
	}

	dat, err := ioutil.ReadFile(file)
	check(err)
	//fmt.Println(string(dat))

    r, err := regexp.Compile(exp)
	check(err)

	res := r.ReplaceAllString(string(dat), replace)
	
	if rep_opt == 0 {
		fmt.Println(res)
	} else {
		ioutil.WriteFile(file, []byte(res), 0644)
	}
}
