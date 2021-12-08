import math


def are_ones_more_than_zeroes(ratings, i):
    sum = 0
    for rating in ratings:
        if rating[i] == "1":
            sum += 1
        else:
            sum -= 1

    return sum >= 0


def elect_rating(ratings, is_co2):
    winning_bit, losing_bit = "1", "0"
    if is_co2:
        winning_bit, losing_bit = "0", "1"

    i = 0
    while True:
        support_ratings = []
        more_ones = are_ones_more_than_zeroes(ratings, i)

        for rating in ratings:
            if more_ones:
                if rating[i] == winning_bit:
                    support_ratings.append(rating)
            else:
                if rating[i] == losing_bit:
                    support_ratings.append(rating)

        if len(support_ratings) == 1:
            return support_ratings[0]
        
        ratings = support_ratings
        i += 1


def main():

    ratings = []

    with open("input.txt") as f:
        for line in f:
            line = line.strip()
            ratings.append(line)

    oxygen_rating = elect_rating(ratings, False)
    co2_rating = elect_rating(ratings, True)

    oxygen, co2 = 0, 0
    for i in range(len(oxygen_rating)):
        if oxygen_rating[i] == "1":
            oxygen += math.pow(2, len(oxygen_rating)-1-i)
        if co2_rating[i] == "1":
            co2 += math.pow(2, len(co2_rating)-1-i)

    print(oxygen * co2)


if __name__ == "__main__":
    main()
