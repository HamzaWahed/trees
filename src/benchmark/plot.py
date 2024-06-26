import matplotlib.pyplot as plt
import csv
import sys
from scipy.stats import ttest_ind

def plot_graph(file_path: str) -> None:
    try:
        reader = csv.reader(open(file_path), delimiter=" ")
    except:
        raise

    input_values = []  
    vEB_runtimes = []
    splay_runtimes = []

    try:
        for row in reader:
            input_values.append(int(row[0]))
            splay_runtimes.append(int(row[1]))
            vEB_runtimes.append(int(row[2]))
    except:
        raise

    plt.plot(input_values, vEB_runtimes, label="Van Emde Boas")
    plt.plot(input_values, splay_runtimes, label="Splay Tree")
    plt.xlabel("Size of Tree")
    plt.ylabel("Runtime (ns)")
    plt.title(file_path)
    plt.xscale("log", base=2)
    plt.legend()
    plt.show()

def stat_test(arr1, arr2, alpha):
    ttest = ttest_ind(arr1, arr2)
    print(f'p-value: {ttest.pvalue}')

    if ttest.pvalue < alpha:
        print("null hypothesis is not rejected")
    else:
        print("null hypothesis is rejected")

def main() -> None:
    if len(sys.argv) < 2:
        raise("No output file to plot")
    
    file = sys.argv[1]
    plot_graph(file)

if __name__ == "__main__":
    main()