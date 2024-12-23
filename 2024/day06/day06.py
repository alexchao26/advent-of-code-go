# guard walks forward, turns right if blocked, turn or step = 1 move


import time


def part1(input: str) -> int:
    return len(get_full_path(input))


# changed this to return the set of all coords on the path to (slightly) speed up part 2
# so part 2 can just modify things already on the guard's path instead of brute forcing
# the entire grid
def get_full_path(input: str) -> set[tuple[int, int]]:
    lines = input.splitlines()

    row: int = 0
    col: int = 0

    # find starting row and col
    for r in range(len(lines)):
        for c in range(len(lines[0])):
            if lines[r][c] == "^":
                row = r
                col = c
                break

    # example and my input both have the guard just starting pointing north
    dir_index: int = 0
    dirs: list[list[int]] = [
        [-1, 0],
        [0, 1],
        [1, 0],
        [0, -1],
    ]

    seen: set[tuple[int, int]] = set()

    # assume that we just have to run off the grid?..
    while row >= 0 and row < len(lines) and col >= 0 and col < len(lines[0]):
        coord = (row, col)

        seen.add(coord)
        dir = dirs[dir_index]
        nextRow: int = row + dir[0]
        nextCol: int = col + dir[1]
        if not (
            nextRow >= 0
            and nextRow < len(lines)
            and nextCol >= 0
            and nextCol < len(lines[0])
        ):
            break
        if lines[nextRow][nextCol] == "#":
            # rotate
            dir_index += 1
            dir_index %= 4
        else:
            row = nextRow
            col = nextCol

    # return len(seen)
    return seen


def part2(input: str) -> int:

    path = get_full_path(input)

    # list(str) divides the str into individual characters?
    # can't just use str.split("")
    # need an actual 2D array so we can modify the grid and add obstacles
    grid = [list(line) for line in input.splitlines()]

    # for every coord on the guard's path, just make it an obstacle and see if it loops
    ans: int = 0
    for coord in path:
        if grid[coord[0]][coord[1]] == ".":
            grid[coord[0]][coord[1]] = "#"
            if does_guard_loop(grid):
                ans += 1
            grid[coord[0]][coord[1]] = "."

    return ans


# slight modification to getting the entire path by tracking the direction
# these two could be combined...
def does_guard_loop(grid: list[list[str]]) -> bool:
    row: int = 0
    col: int = 0

    # find starting row and col
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "^":
                row = r
                col = c
                break

    # example and my input both have the guard just starting pointing north
    dir_index: int = 0
    dirs: list[list[int]] = [
        [-1, 0],
        [0, 1],
        [1, 0],
        [0, -1],
    ]

    # keys: row, col, dir_index
    seen: set[tuple[int, int, int]] = set()

    # assume that we just have to run off the grid?..
    while row >= 0 and row < len(grid) and col >= 0 and col < len(grid[0]):
        key = (row, col, dir_index)

        if key in seen:
            return True

        seen.add(key)
        dir = dirs[dir_index]
        nextRow: int = row + dir[0]
        nextCol: int = col + dir[1]
        if not (
            nextRow >= 0
            and nextRow < len(grid)
            and nextCol >= 0
            and nextCol < len(grid[0])
        ):
            break
        if grid[nextRow][nextCol] == "#":
            # rotate
            dir_index += 1
            dir_index %= 4
        else:
            row = nextRow
            col = nextCol

    # no loop found
    return False


example = """....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."""

input = open("input.txt").read()


print(f"part1 example: {part1(example)} want 41")
print(f"part1: {part1(input)} want 4939")

print(f"part2 example: {part2(example)} want 6")

start_time = time.time()
print(f"part2: {part2(input)} want 1434")  # slow as heck... ~30s on my laptop
end_time = time.time()
print(f"runtime = {end_time - start_time} seconds")
