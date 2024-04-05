import numpy as np

rng = np.random.default_rng()

RANGE_SMALL = 512
RANGE_MEDIUM = 65536
RANGE_LARGE = pow(2, 32)
RANGE_HIGHEST = pow(2, 64)
NUMBER_OF_REQUESTS = pow(2, 16)  # at most, can be slightly less because of float to int round


# Takes in a string to identify the  file, ideally a countable value
def uniform_distribution(input_number):
    elements = rng.uniform(low=0, high=RANGE_SMALL - 1, size=NUMBER_OF_REQUESTS)
    np.savetxt("inputs/uniform/small.in." + input_number, elements, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_MEDIUM - 1, size=NUMBER_OF_REQUESTS)
    np.savetxt("inputs/uniform/medium.in." + input_number, elements, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_LARGE - 1, size=NUMBER_OF_REQUESTS)
    np.savetxt("inputs/uniform/large.in." + input_number, elements, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_HIGHEST - 1, size=NUMBER_OF_REQUESTS)
    np.savetxt("inputs/uniform/highest.in." + input_number, elements, fmt="%d")


# Takes in a string to identify the  file, ideally a countable value
def mtf_optimal(input_numer):
    subrange = rng.uniform(low=1, high=NUMBER_OF_REQUESTS / 10, size=1).astype(int)[0]
    number_of_distinct_elements = (NUMBER_OF_REQUESTS / subrange).astype(int)

    elements = rng.uniform(low=0, high=RANGE_SMALL - 1, size=number_of_distinct_elements)
    to_save = []
    for i in elements:
        for j in range(1, subrange):
            to_save.append(i)
    np.savetxt("inputs/mtf_opt/small.in." + input_numer, to_save, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_MEDIUM - 1, size=number_of_distinct_elements)
    to_save = []
    for i in elements:
        for j in range(1, subrange):
            to_save.append(i)
    np.savetxt("inputs/mtf_opt/medium.in." + input_numer, to_save, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_LARGE - 1, size=number_of_distinct_elements)
    to_save = []
    for i in elements:
        for j in range(1, subrange):
            to_save.append(i)
    np.savetxt("inputs/mtf_opt/large.in." + input_numer, to_save, fmt="%d")

    elements = rng.uniform(low=0, high=RANGE_HIGHEST - 1, size=number_of_distinct_elements)
    to_save = []
    for i in elements:
        for j in range(1, subrange):
            to_save.append(i)
    np.savetxt("inputs/mtf_opt/highest.in." + input_numer, to_save, fmt="%d")


if __name__ == '__main__':
    for input_number in range(1, 32):
        uniform_distribution(str(input_number))
        mtf_optimal(str(input_number))
