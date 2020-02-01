//
//
//

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/rrreeezzz/hshcd2020/modes"
)

func usage() {
	fmt.Println("./main -file file -mode mode\n" +
		"   MODES: pickfirst,")
	os.Exit(0)
}

// printTab print results tab-separated style
func printHeader() *tabwriter.Writer {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Mode\tPizzas\tSlices\tDiversity\tEfficiency")
	return w
}

func printResultRow(w *tabwriter.Writer, name string, pizs, slices int, density, efficiency float64) {
	fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v\t%.4f\t%.4f", name, pizs, slices, density, efficiency))
}

func readAndParseData(fname string) (int, int, []int, error) {
	var max, num int
	var pizs []int
	var err error

	file, err := os.Open(fname)
	if err != nil {
		return 0, 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read maximum slices and number of different pizzas
	if b := scanner.Scan(); b == false {
		return 0, 0, nil, fmt.Errorf("Error reading file")
	}
	s := strings.Split(scanner.Text(), " ")
	if len(s) != 2 {
		return 0, 0, nil, fmt.Errorf("Error reading file")
	}
	if max, err = strconv.Atoi(s[0]); err != nil {
		return 0, 0, nil, fmt.Errorf("%v", err)
	}
	if num, err = strconv.Atoi(s[1]); err != nil {
		return 0, 0, nil, fmt.Errorf("%v", err)
	}

	// Read pizzas
	if b := scanner.Scan(); b == false {
		return 0, 0, nil, fmt.Errorf("Error reading file")
	}
	s = strings.Split(scanner.Text(), " ")
	if len(s) != num {
		return 0, 0, nil, fmt.Errorf("Error reading file")
	}
	for _, i := range s {
		tmp, err := strconv.Atoi(i)
		if err != nil {
			return 0, 0, nil, fmt.Errorf("%v", err)
		}
		pizs = append(pizs, tmp)
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, nil, err
	}

	return max, num, pizs, err
}

func main() {

	fname := flag.String("file", "", "file")
	mode := flag.String("mode", "", "mode: pickfirsts, picklasts, all")
	full := flag.Bool("full", false, "show selected pizzas list")
	flag.Parse()

	if *fname == "" || *mode == "" {
		usage()
	}

	var m modes.Mode
	var all bool
	switch *mode {
	case "pickfirsts":
		m = modes.NewPickfirsts()
	case "picklasts":
		m = modes.NewPicklasts()
	case "all":
		all = true
	default:
		usage()
	}

	max, num, pizs, err := readAndParseData(*fname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Maximum slices of pizzas: %v\n", max)
	fmt.Printf("Number of available pizzas: %v\n", num)
	fmt.Println()

	// TODO: delete duplicate code
	// TODO: incremental print, get rid of tabwriter
	w := printHeader()
	if all {
		for _, m := range availablesModes {
			numSlices, selectedPizs := m.Run(max, num, pizs)
			lenPizs := len(selectedPizs)
			density := (100 * float64(lenPizs)) / float64(num)
			efficiency := (100 * float64(numSlices)) / float64(max)
			printResultRow(w, m.Name(), lenPizs, numSlices, density, efficiency)
		}
	} else {
		numSlices, selectedPizs := m.Run(max, num, pizs)
		lenPizs := len(selectedPizs)
		density := (100 * float64(lenPizs)) / float64(num)
		efficiency := (100 * float64(numSlices)) / float64(max)
		printResultRow(w, m.Name(), lenPizs, numSlices, density, efficiency)
		if *full {
			fmt.Printf("Selected pizzas: %v\n", selectedPizs)
		}
	}
	w.Flush()
}

var availablesModes = []modes.Mode{
	&modes.Pickfirsts{},
	&modes.Picklasts{},
}
