package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
	"trees/trees"
)

const BECNHMARK_UNIFORM = "-u"
const BECHMARK_MTF = "-m"
const BENCHMARK_ALL = "-a"

const NUMBER_OF_REQUESTS int = 1000
const NUMBER_OF_INPUT_FILES = 50
const MAX_TREE_SIZE int = 16384

var NUMBER_OF_OUTPUTS = int(math.Log2(float64(MAX_TREE_SIZE))) * NUMBER_OF_INPUT_FILES

/*
*
function runs the input through both data structures and returns the times they took each. (splay, veb)
*/
func benchmark(size int, input_path string) (time.Duration, time.Duration) {
	var file *os.File
	var input []int = make([]int, NUMBER_OF_REQUESTS)
	var cur int
	var err error

	// reading in elem from the input files and saving them in an array
	var path = string("../src/inputs/" + input_path)
	file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic("couldn't read file, uniform/small.in")
	}

	for i := 0; i < NUMBER_OF_REQUESTS; i++ {
		_, err := fmt.Fscanln(file, &cur)
		if err != nil {
			break
		}
		input[i] = cur % size
	}

	//initializing the splay and vEm Trees
	var vEM *trees.VEB = trees.BuildVEB(size)
	var splay *trees.SplayTree = trees.NewSplayTree(0)
	for i := 1; i < size; i++ {
		splay.Insert(i)
	}

	var splay_time time.Duration
	var vEM_time time.Duration

	for i := 0; i < NUMBER_OF_REQUESTS; i++ {
		s_start := time.Now()
		splay.Search(input[i])
		splay_time += time.Since(s_start)

		v_start := time.Now()
		vEM.Member(input[i])
		vEM_time += time.Since(v_start)
	}

	return splay_time, vEM_time
}

func benchmark_set(method string) [][]time.Duration {

	var result [][]time.Duration = make([][]time.Duration, NUMBER_OF_OUTPUTS)
	for i := range result {
		result[i] = make([]time.Duration, 3)
	}

	var index int = 0

	for i := 2; i <= MAX_TREE_SIZE; i *= 2 {
		for path := 1; path < 51; path++ {
			result[index][0] = time.Duration(i)
			result[index][1], result[index][2] = benchmark(i, method+"/"+strconv.FormatInt(int64(path), 10))
			index++
		}
		fmt.Println("break")
	}

	return result
}

func contains(args []string, target string) bool {
	for i := range args {
		if args[i] == target {
			return true
		}
	}
	return false
}

func write_to_output(path string, array [][]time.Duration) {
	outputPath := fmt.Sprintf("../src/output/%s", path)
	fmt.Printf("Output data to %s\n", outputPath)
	var output, _ = os.OpenFile(outputPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	for i := 0; i < NUMBER_OF_OUTPUTS; i++ {
		if i == NUMBER_OF_OUTPUTS-1 {
			_, err := output.WriteString(strconv.FormatInt(int64(array[i][0]), 10) + " " +
				strconv.FormatInt(int64(array[i][1].Nanoseconds()), 10) + " " + strconv.FormatInt(int64(array[i][2].Nanoseconds()), 10))

			if err != nil {
				log.Fatal(err)
			}

			continue
		}

		_, err := output.WriteString(strconv.FormatInt(int64(array[i][0]), 10) + " " +
			strconv.FormatInt(int64(array[i][1].Nanoseconds()), 10) + " " + strconv.FormatInt(int64(array[i][2].Nanoseconds()), 10) + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	var cmdArgs []string = os.Args

	if !(contains(cmdArgs, BECNHMARK_UNIFORM) || contains(cmdArgs, BECHMARK_MTF) || contains(cmdArgs, BENCHMARK_ALL)) {
		panic("Invalid flag")
	}

	fmt.Printf("Performing Searches...\n")

	// create output directory if it does not exist
	_ = os.Mkdir("output", 0750)

	if contains(cmdArgs, BENCHMARK_ALL) {
		write_to_output("uniform_small.out", benchmark_set("uniform"))
		write_to_output("uniform_medium.out", benchmark_set("uniform"))

		write_to_output("mtf_opt_small.out", benchmark_set("mtf_opt"))
		write_to_output("mtf_opt_medium.out", benchmark_set("mtf_opt"))

	} else if contains(cmdArgs, BECNHMARK_UNIFORM) {
		write_to_output("uniform_small.out", benchmark_set("uniform"))
		write_to_output("uniform_medium.out", benchmark_set("uniform"))
	} else {
		write_to_output("mtf_opt_small.out", benchmark_set("mtf_opt"))
		write_to_output("mtf_opt_medium.out", benchmark_set("mtf_opt"))
	}
}
