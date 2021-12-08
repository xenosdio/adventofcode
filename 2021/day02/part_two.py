class Position:
    def __init__(self) -> None:
        self.distance = 0
        self.depth = 0
        self.aim = 0


def main():
    position = Position()

    with open("input.txt") as f:
        for line in f:
            instructions = line.split()

            direction = instructions[0]
            units = int(instructions[1])

            if direction == "forward":
                position.distance += units
                position.depth += (position.aim * units)
            elif direction == "up":
                position.aim -= units
            elif direction == "down":
                position.aim += units
            else:
                raise Exception("wrong direction given")

    print(position.depth * position.distance)


if __name__ == "__main__":
    main()
