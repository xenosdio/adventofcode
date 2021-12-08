import math


def main():

    bits = {}

    with open("input.txt") as f:

        for line in f:
            line = line.strip()

            for i in range(len(line)):
                if i not in bits:
                    bits[i] = 0

                if line[i] == "0":
                    bits[i] -= 1
                else:
                    bits[i] += 1

    gamma, epsilon = 0, 0
    for i in range(len(bits)):
        if bits[i] > 0:
            gamma += math.pow(2, len(bits)-1-i)
        else:
            epsilon += math.pow(2, len(bits)-1-i)

    print(gamma * epsilon)


if __name__ == "__main__":
    main()
