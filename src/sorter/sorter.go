/**
 * sorter.go:10:2: cannot find package "bubblesort" in any of:
 *      D:\Program\go\src\bubblesort (from $GOROOT)
 *      E:\gosource\learnGoLang\src\bubblesort (from $GOPATH)
**/
package main

import(
	"flag"
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
	"sorter/bubblesort"
	"time"
)

var infile  *string = flag.String("i", "unsorted.dat", "File contains values for storing")
var outfile *string = flag.String("o", "sorted.data", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func ReadValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Falied to open the input file", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err_1 := br.ReadLine()
		if (err_1 != nil) {
			if err_1 != io.EOF {
				err = err_1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seem unexpected.")
			return
		}

		str := string(line)
		value, err_2 := strconv.Atoi(str)
		if err_2 != nil {
			err = err_2
			return
		}
		values = append(values, value)
	}
	return
}

func WriteValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err 																//not enough arguments to return
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")		
		// file.WriteString(value)								//cannot use value (type int) as type string in argument to file.WriteString
	}	

	return nil 

}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := ReadValues(*infile)
	if err != nil {
		fmt.Println(err)
		return
	} 

	startTime := time.Now()
	switch *algorithm {
		case "bubble":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sort algorithm", *algorithm, "is unknown")
	}
	endtime := time.Now()

	//.\sorter.go:99: invalid operation: endtime - startTime (operator - not defined on struct)
	fmt.Println("The sorting process casts ", endtime.Sub(startTime), " to complete")			
	WriteValues(values, *outfile)
}