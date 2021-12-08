def main():
    cnt = 0
    previous_depth = None

    with open("input.txt") as f:
        for line in f:
            current_depth = int(line)
            if previous_depth is not None and current_depth > previous_depth:
                cnt += 1

            previous_depth = current_depth
            
    print(cnt)


if __name__ == "__main__":
    main()
