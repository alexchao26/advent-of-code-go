#!/usr/bin/env python3

from typing import List

example = """7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"""


def part1(grid: List[List[int]]) -> int:
    valid_levels = 0
    for level in grid:
        if test_level(level):
            valid_levels += 1
    return valid_levels


def part2(grid: List[List[int]]) -> int:
    # tolerate one bad level...
    valid_levels: int = 0
    for level in grid:
        for i in range(len(level)):
            newLevel: List[int] = level.copy()
            newLevel.pop(i)  # removes i-th element?, how convenient..
            if test_level(newLevel):
                valid_levels += 1
                break

    return valid_levels


def test_level(level: List[int]) -> bool:
    is_increasing = level[1] > level[0]
    is_valid = True

    # The levels are either all increasing or all decreasing.
    # Any two adjacent levels differ by at least one and at most three.
    for i in range(1, len(level)):
        if is_increasing and level[i] <= level[i - 1]:
            return False
        elif not is_increasing and level[i] >= level[i - 1]:
            return False

        diff = abs(level[i] - level[i - 1])
        if diff < 1 or diff > 3:
            return False

    return is_valid


def convert_input(input: str) -> List[List[int]]:
    grid: List[List[int]] = []

    for level in input.splitlines():
        grid.append([int(part) for part in level.split(" ")])
        # converted_level: List[int] = []
        # this syntax is going to take getting used to...
        # for part in level.split(" "):
        #     converted_level.append(int(part))
        # grid.append(converted_level)

    return grid


example_grid = convert_input(example)
print("example part1:", part1(example_grid), "want", 2)

input = open("input.txt", "r").read()
grid = convert_input(input)
print("part1:", part1(grid), "want", 585)

print("example part2:", part2(example_grid), "want", 4)
print("part2:", part2(grid), "want", 626)
