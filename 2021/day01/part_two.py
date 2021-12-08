def main():

    depths = {}

    with open("input.txt") as f:
        i = 0

        for line in f:
            current_depth = int(line)

            depths[i] = current_depth
            if i - 1 >= 0:
                depths[i-1] += current_depth
            if i - 2 >= 0:
                depths[i-2] += current_depth

            i += 1

    cnt = 0
    for i in range(1, len(depths)):
        if depths[i] > depths[i-1]:
            cnt += 1

    print(cnt)


if __name__ == "__main__":
    main()
