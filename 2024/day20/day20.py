from collections import defaultdict


def day20(input: str, cheats: int) -> dict[int, int]:
    # get minimum without cheats, do this backwards by propagating from E to S
    #   to populate a grid of the distance from each cell to E
    grid: list[list[str]] = [list(line) for line in input.splitlines()]

    row: int = 0
    col: int = 0

    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "E":
                row = r
                col = c

    # default value set to 100_000 which is bigger than entire grid's area so
    # walls will not be usable as paths to finish
    dp_steps_to_end: list[list[int]] = [
        [100_000 for _ in range(len(grid[0]))] for _ in range(len(grid))
    ]

    diffs = [(0, 1), (0, -1), (-1, 0), (1, 0)]
    queue: list[tuple[int, int, int]] = [(row, col, 0)]
    visited: set[tuple[int, int]] = set()
    while len(queue) > 0:
        row, col, steps = queue.pop(0)
        dp_steps_to_end[row][col] = steps
        visited.add((row, col))

        for diff in diffs:
            next_row, next_col = row + diff[0], col + diff[1]
            if 0 <= next_row < len(grid) and 0 <= next_col < len(grid[0]):
                if grid[next_row][next_col] != "#":
                    if (next_row, next_col) not in visited:
                        queue.append((next_row, next_col, steps + 1))

    # then starting from S, BFS but from each cell, view all cells <cheats> cells away
    # if it's a valid shortcut, record it

    second_saved_freqs: dict[int, int] = defaultdict(int)

    reachable_diffs = get_reachable_coords_within_x_steps(cheats)

    # visited contains coords of the entire path, so just use that...
    for path_coord in visited:
        r, c = path_coord
        for diff in reachable_diffs:
            rr = path_coord[0] + diff[0]
            cc = path_coord[1] + diff[1]

            if 0 <= rr < len(grid) and 0 <= cc < len(grid[0]):
                if grid[rr][cc] != "#":
                    savings = dp_steps_to_end[r][c] - (
                        dp_steps_to_end[rr][cc] + abs(diff[0]) + abs(diff[1])
                    )
                    if savings > 0:
                        second_saved_freqs[savings] += 1

    return second_saved_freqs


# helper function to get all coords that are reachable within allotted cheats
def get_reachable_coords_within_x_steps(x: int) -> list[tuple[int, int]]:
    coords_map: dict[tuple[int, int], bool] = {}

    queue: list[tuple[int, int]] = [(0, 0)]
    for _ in range(x):
        next_queue: list[tuple[int, int]] = []
        while len(queue) > 0:
            front = queue.pop(0)
            for diff in [(0, -1), (0, 1), (-1, 0), (1, 0)]:
                coord = (front[0] + diff[0], front[1] + diff[1])
                if coord not in coords_map:
                    coords_map[coord] = True
                    next_queue.append(coord)
        queue = next_queue

    return list(coords_map.keys())


example = """###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############"""

# python is weird and we can use == to compare 2 dicts for the same keys and values...
example_seconds_saved_to_freq: dict[int, int] = {
    2: 14,
    4: 14,
    6: 2,
    8: 4,
    10: 2,
    12: 3,
    20: 1,
    36: 1,
    38: 1,
    40: 1,
    64: 1,
}

part1_example_result = day20(example, 2)
print(
    f"part1 example: {part1_example_result == example_seconds_saved_to_freq} want True"
)
if part1_example_result != example_seconds_saved_to_freq:
    print(
        f"\tDEBUG part1 example got: {part1_example_result} want {(example_seconds_saved_to_freq)}"
    )

input = open("input.txt").read().strip()


def sum_keys_over_x(d: dict[int, int], x: int) -> int:
    total: int = 0
    for k, v in d.items():
        if k >= x:
            total += v
    return total


part1_result = day20(input, 2)

print(
    f"part1 actual: {sum_keys_over_x(part1_result, 100)} cheats save 100+ picosecs want 1422"
)

part2_example_result: dict[int, int] = {
    50: 32,
    52: 31,
    54: 29,
    56: 39,
    58: 25,
    60: 23,
    62: 20,
    64: 19,
    66: 12,
    68: 14,
    70: 12,
    72: 22,
    74: 4,
    76: 3,
}

part2_result = day20(example, 20)
part2_result_50 = {}
for k, v in part2_example_result.items():
    if k >= 50:
        part2_result_50[k] = v

print(f"part2 example: {part2_result_50 == part2_example_result} want True")
if part2_result_50 != part2_example_result:
    print(
        f"DEBUG part2 example 50+ saved: {part2_result_50} want {part2_example_result}"
    )

part2_result = day20(input, 20)
print(f"part2 actual: {sum_keys_over_x(part2_result, 100)} want 1009299")
