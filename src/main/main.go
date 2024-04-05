package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"trees/trees"
)

const NUMBER_OF_REQUESTS int = 65536
const SMALL int = 512
const MEDIUM int = 65536
const LARGE int = 4294967296
const BECNHMARK_UNIFORM = "-u"
const BECHMARK_MTF = "-m"
const BENCHMARK_ALL = "-a"
const NUMBER_OF_INPUT_FILES = 31

/*
*
function runs the input through both data structures and returns the times they took each. (splay, veb)
*/
func benchmark(size int, method string) (time.Duration, time.Duration) {
	var current_range int = size
	var file *os.File
	var input []int = make([]int, NUMBER_OF_REQUESTS)
	var cur int
	var err error

	// reading in elem from the input files and saving them in an array
	var path = string("../src/" + method)
	file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic("couldn't read file, uniform/small.in")
	}
	for i := 0; i < NUMBER_OF_REQUESTS; i++ {
		_, err := fmt.Fscanln(file, &cur)
		if err != nil {
			break
		}
		input[i] = cur
	}

	//initializing the splay and vEm Trees
	var vEM *trees.VEB = trees.BuildVEB(current_range)
	var splay *trees.SplayTree = trees.NewSplayTree(0)
	for i := 1; i < current_range; i++ {
		splay.Insert(i)
	}

	var splay_time time.Duration
	var vEM_time time.Duration

	for i := 0; i < current_range; i++ {

		s_start := time.Now()
		splay.Search(input[i])
		splay_time += time.Since(s_start)

		v_start := time.Now()
		vEM.Member(input[i])
		vEM_time += time.Since(v_start)
	}

	//fmt.Printf("Input Size: %d, Sampling Distribution: Uniform\n", current_range)
	//fmt.Printf("Splay: ")
	//fmt.Println(splay_time)
	//fmt.Printf("vEM: ")
	//fmt.Println(vEM_time)
	return splay_time, vEM_time
}

func benchmark_set(input_type string, input_size int, number_of_input_files int) [][]time.Duration {

	var input_file string
	switch input_size {
	case SMALL:
		input_file = "small.in."
	case MEDIUM:
		input_file = "medium.in."
	case LARGE:
		input_file = "large.in."
	}
	var path string = "inputs/" + input_type + "/" + input_file

	// initializing the 2d slice
	var result [][]time.Duration = make([][]time.Duration, number_of_input_files)
	for i := range number_of_input_files {
		result[i] = make([]time.Duration, 2)
	}

	for i := range number_of_input_files {
		result[i][0], result[i][1] = benchmark(input_size, path+strconv.Itoa(i+1))
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
	var output, _ = os.OpenFile("output/"+path, os.O_CREATE|os.O_APPEND, 0644)

	for i := 0; i < NUMBER_OF_INPUT_FILES; i++ {
		_, err := output.WriteString(strconv.FormatInt(int64(array[i][0]), 10) + " " +
			strconv.FormatInt(int64(array[i][1]), 10) + "\n")

		if err != nil {
			panic("error when writing to file")
		}
	}
}

func main() {

	var cmd_args []string = os.Args

	if !(contains(cmd_args, BECNHMARK_UNIFORM) || contains(cmd_args, BECNHMARK_UNIFORM) || contains(cmd_args, BENCHMARK_ALL)) {
		panic("Invalid flag")
	}

	fmt.Printf("Performing Searches...\n\n")

	if contains(cmd_args, BENCHMARK_ALL) {
		write_to_output("uniform_small.out", benchmark_set("uniform", SMALL, NUMBER_OF_INPUT_FILES))
		write_to_output("uniform_medium.out", benchmark_set("uniform", MEDIUM, NUMBER_OF_INPUT_FILES))

		write_to_output("mtf_opt_small.out", benchmark_set("mtf_opt", SMALL, NUMBER_OF_INPUT_FILES))
		write_to_output("mtf_opt_medium.out", benchmark_set("mtf_opt", SMALL, NUMBER_OF_INPUT_FILES))

	} else if contains(cmd_args, BECHMARK_MTF) {
		write_to_output("uniform_small.out", benchmark_set("mtf_opt", SMALL, NUMBER_OF_INPUT_FILES))
		write_to_output("uniform_medium.out", benchmark_set("mtf_opt", SMALL, NUMBER_OF_INPUT_FILES))
	} else {
		write_to_output("mtf_opt_small.out", benchmark_set("uniform", SMALL, NUMBER_OF_INPUT_FILES))
		write_to_output("mtf_opt_medium.out", benchmark_set("uniform", MEDIUM, NUMBER_OF_INPUT_FILES))
	}

}
