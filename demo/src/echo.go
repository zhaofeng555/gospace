package main

import(
	//"os"
	"flag"
	"fmt"
)

var omitNewline=flag.Bool("n", false, "don't print final newline")

const(
	Space = " "
	Newline = "\n"
)

func main(){
	flag.Parse()

	var s string = "";

	for i := 0; i < flag.NArg(); i++{

		if i>0{
			s += Space
		}

		s += Space;

	}

	if !*omitNewline{
		s += Newline
	}

	//os.Stdout.WriteString(s)

	fmt.Println(s);

	c := sum([3]int{1,2,3}[:])

	fmt.Println(c)

}

func sum(a []int) int{
	s:=0
	for i := 0; i < len(a); i++{
		s += a[i]
	}
	return s;
}