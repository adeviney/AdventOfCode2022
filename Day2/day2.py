# A - ROCK
# B - Paper
# C - Scissors

# X - rock + 1
# Y - Paper + 2
# Z - Scissors + 3

# Game Possibilities
# BX 1
# CY 2
# AZ 3
# AX 4
# BY 5
# CZ 6
# CX 7
# AY 8
# BZ 9

# Part 1
with open("input.txt") as f: data = f.read().split('\n')

# each index corresponds to the score of that round
points = ["", "B X", "C Y", "A Z", "A X", "B Y", "C Z", "C X", "A Y", "B Z"]
score_p1 = 0
for round in data:
    score_p1 += points.index(round)
print(f"part 1: {score_p1}")

# Part 2
# A - ROCK
# B - Paper
# C - Scissors 
# X - lose
# Y - draw
# Z - win

# Game Possibilities, in ascending order of points
# BX rock
# CX paper
# AX scissors
# AY rock
# BY paper
# CY scissors
# CZ rock
# AZ paper
# BZ scissors

points_p2 = ["", "B X", "C X", "A X", "A Y", "B Y", "C Y", "C Z", "A Z", "B Z"]
score_p2 = 0
for round in data:
    score_p2 += points_p2.index(round)
print(f"part 2: {score_p2}")





