from typing import List, Tuple


example = """MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"""


def part1(input: str) -> int:
    grid: List[str] = input.splitlines()

    ans = 0

    # 8 directions that words can be found in once a "X" is found
    dirs: List[Tuple[int, int]] = [
        (-1, -1),
        (-1, 0),
        (-1, 1),
        (0, -1),
        (0, 1),
        (1, -1),
        (1, 0),
        (1, 1),
    ]

    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == "X":
                for dir in dirs:
                    if getWord(grid, i, j, dir) == "XMAS":
                        ans += 1
    return ans


def getWord(grid: List[str], row: int, col: int, dir: Tuple[int, int]) -> str:
    word = grid[row][col]
    while len(word) < 4:
        row += dir[0]
        col += dir[1]
        if row < 0 or row >= len(grid) or col < 0 or col >= len(grid[0]):
            return ""
        word += grid[row][col]

    return word


def part2(input: str) -> int:
    grid: List[str] = input.splitlines()

    ans: int = 0

    # do not search outer border of the grid because we're looking for A's which
    # will only be "valid" if they are not on the outer border
    for i in range(1, len(grid) - 1):
        for j in range(1, len(grid[0]) - 1):
            if grid[i][j] == "A":
                backslash_word = grid[i - 1][j - 1] + grid[i + 1][j + 1]
                slash_word = grid[i - 1][j + 1] + grid[i + 1][j - 1]

                backslash_valid = backslash_word == "MS" or backslash_word == "SM"
                slash_valid = slash_word == "MS" or slash_word == "SM"
                if backslash_valid and slash_valid:
                    ans += 1
    return ans


input = open("input.txt", "r").read()

print(f"part1 example: {part1(example)} want 18")
print(f"part1: {part1(input)} want 2468")

print(f"part2 example: {part2(example)} want 9")
print(f"part2: {part2(input)} want 1864")
