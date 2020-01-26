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

	"github.com/rrreeezzz/hshcd2020/modes"
)

func usage() {
	fmt.Println("./main -file file -mode mode\n" +
		"   MODES: pickfirst,")
	os.Exit(0)
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
	mode := flag.String("mode", "", "mode: pickfirsts,")
	flag.Parse()

	if *fname == "" || *mode == "" {
		usage()
	}

	var m modes.Mode
	switch *mode {
	case "pickfirsts":
		m = modes.NewPickFirsts()
	default:
		usage()
	}

	max, num, pizs, err := readAndParseData(*fname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.Run(max, num, pizs))
}
