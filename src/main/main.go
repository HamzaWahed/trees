package main

import (
	"fmt"
	"os"
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

func benchmark(size int, method string) {
	var current_range int = size
	var file *os.File
	var input []int = make([]int, NUMBER_OF_REQUESTS)
	var cur int
	var err error

	// reading in elem from the input files and saving them in an array
	var path = string("../src/inputs/" + method)
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

	fmt.Printf("Input Size: %d, Sampling Distribution: Uniform\n", current_range)
	fmt.Printf("Splay: ")
	fmt.Println(splay_time)
	fmt.Printf("vEM: ")
	fmt.Println(vEM_time)
}

func contains(args []string, target string) bool {
	for i := range args {
		if args[i] == target {
			return true
		}
	}

	return false
}

func main() {

	var cmd_args []string = os.Args

	if !(contains(cmd_args, BECNHMARK_UNIFORM) || contains(cmd_args, BECNHMARK_UNIFORM) || contains(cmd_args, BENCHMARK_ALL)) {
		panic("Invalid flag")
	}

	fmt.Printf("Performing Searches...\n\n")

	if contains(cmd_args, BENCHMARK_ALL) {
		benchmark(SMALL, "uniform/small.in")
		benchmark(MEDIUM, "uniform/medium.in")
		fmt.Println()
		benchmark(SMALL, "mtf_opt/small.in")
		benchmark(MEDIUM, "mtf_opt/medium.in")
		//add large but for now its too slow
	} else if contains(cmd_args, BECHMARK_MTF) {
		benchmark(SMALL, "mtf_opt/small.in")
		benchmark(MEDIUM, "mtf_opt/medium.in")
	} else {
		benchmark(SMALL, "uniform/small.in")
		benchmark(MEDIUM, "uniform/medium.in")
	}

}
