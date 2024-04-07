import numpy as np

rng = np.random.default_rng()
REQUEST_RANGE = pow(2, 32)


# Takes in a string to identify the file, ideally a countable value
def uniform_distribution(number_of_requests):
    for file_number in range(1, 51):
        elements = rng.uniform(low=0, high=REQUEST_RANGE - 1, size=number_of_requests)
        np.savetxt("inputs/uniform/" + str(file_number), elements, fmt="%d")


# Takes in a string to identify the file, ideally a countable value
def mtf_optimal(number_of_requests):
    for file_number in range(1, 51):
        subrange = rng.uniform(low=1, high=number_of_requests / 10, size=1).astype(int)[0]
        number_of_distinct_elements = (number_of_requests / subrange).astype(int)

        elements = rng.uniform(low=0, high=REQUEST_RANGE - 1, size=number_of_distinct_elements)
        to_save = []
        for i in elements:
            for j in range(0, subrange):
                to_save.append(i)
        np.savetxt("inputs/mtf_opt/" + str(file_number), to_save, fmt="%d")


if __name__ == '__main__':
    mtf_optimal(1000)
    uniform_distribution(1000)
