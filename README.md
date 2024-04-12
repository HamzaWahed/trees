# Benchmarking Van Emde Boas and Splay Trees

## Getting Started

Create a virtual environment for Python 3.12 and download the python packages using the command.
```
pip install requirements.txt
```
Setup [GoLand](https://www.jetbrains.com/go/?var=1) or any other ide from Jetbrains.

## Benchmarking

Run the command below to generate the input files.
```
python generate_input.py
```
Set the working directory to be the src directory and program arguments to be one of the following:
- -a for both uniform and move-to-front distribution
- -u for uniform distribution
- -m for move-to-front distribution

Run the main.go file from the main module to create the output. The output files are space delimited with three columns.
The first, second and third columns are the tree size, splay tree runtime and Van Emde Boas tree runtime respectively.
Go into the benchmark directory.
```
cd src/benchmark/
```
Run the python script plot.py with a specified output file to plot it. An example command is given below.
```
py plot.py ../output/mtf_opt_medium.out
```

## T-Test
The plot.py script also contains a scipy t-test function. To run it, simply call the stat_test function in the main
function with the specified significance value.
