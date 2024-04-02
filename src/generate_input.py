import numpy as np

rng = np.random.default_rng()

RANGE_SMALL = 512
RANGE_MEDIUM = 65536
RANGE_LARGE = pow(2, 32)
RANGE_HIGHEST = pow(2, 64)
NUMBER_OF_REQUESTS = pow(2, 16) #at most, can be slightly less because of float to int round


def uniform_distribution():
    elements = rng.uniform(low=0, high=RANGE_SMALL - 1, size=NUMBER_OF_REQUESTS)
    np.savetxt("inputs/uniform/small.in", elements, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_MEDIUM - 1, size=NUMBER_OF_REQUESTS)
    np.savetxt("inputs/uniform/medium.in", elements, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_LARGE - 1, size=NUMBER_OF_REQUESTS)
    np.savetxt("inputs/uniform/large.in", elements, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_HIGHEST - 1, size=NUMBER_OF_REQUESTS)
    np.savetxt("inputs/uniform/highest.in", elements, fmt="%d")


def mtf_optimal():
    subrange = rng.uniform(low=1, high=NUMBER_OF_REQUESTS / 10, size=1).astype(int)[0]
    number_of_distinct_elements = (NUMBER_OF_REQUESTS / subrange).astype(int)

    elements = rng.uniform(low=0, high=RANGE_SMALL - 1, size=number_of_distinct_elements)
    to_save = []
    for i in elements:
        for j in range(1, subrange):
            to_save.append(i)
    np.savetxt("inputs/mtf_opt/small.in", to_save, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_MEDIUM - 1, size=number_of_distinct_elements)
    to_save = []
    for i in elements:
        for j in range(1, subrange):
            to_save.append(i)
    np.savetxt("inputs/mtf_opt/medium.in", to_save, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_LARGE - 1, size=number_of_distinct_elements)
    to_save = []
    for i in elements:
        for j in range(1, subrange):
            to_save.append(i)
    np.savetxt("inputs/mtf_opt/large.in", to_save, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_HIGHEST - 1, size=number_of_distinct_elements)
    to_save = []
    for i in elements:
        for j in range(1, subrange):
            to_save.append(i)
    np.savetxt("inputs/mtf_opt/highest.in", to_save, fmt="%d")


if __name__ == '__main__':
    uniform_distribution()
    mtf_optimal()
