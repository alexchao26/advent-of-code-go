import heapq


diffs: list[tuple[int, int]] = [
    (0, 1),  # start at index 0 facing "east"/right
    (-1, 0),  # up
    (0, -1),  # left
    (1, 0),  # down
]


def day16(input: str, part: int) -> int:
    grid = [list(line) for line in input.splitlines()]
    start_row: int = 0
    start_col: int = 0
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "S":
                start_row = r
                start_col = c

    # i think i'd rather write a struct in go... these tuples are easy to make
    # but annoying to maintain in my head...
    # score, row, col, dir_index
    node: tuple[int, int, int, int, set[tuple[int, int]]] = (
        int(0),
        start_row,
        start_col,
        int(0),
        set(),
    )
    min_heap = [node]
    heapq.heapify(min_heap)

    coord_to_min_score: dict[tuple[int, int, int], int] = {}

    best_score: int = 1_000_000  # big enough to not interfere with actual input answer
    final_path_coords: set[tuple[int, int]] = set()

    while len(min_heap) > 0:
        popped = heapq.heappop(min_heap)

        # part2 exit once best_score is set down to the actual min value
        if popped[0] > best_score:
            continue

        # row, col, dir_index
        coord = (popped[1], popped[2], popped[3])
        if coord in coord_to_min_score:
            prev_score = coord_to_min_score[coord]
            # if previous score at this coord and direction is better then continue
            if popped[0] > prev_score:
                continue

        # update best score to reach this coord
        coord_to_min_score[coord] = popped[0]

        # end reached
        if grid[popped[1]][popped[2]] == "E":
            best_score = popped[0]
            final_path_coords |= popped[4]
            final_path_coords.add((popped[1], popped[2]))
            continue

        # in same direction
        diff = diffs[popped[3]]
        next = (
            popped[0] + 1,
            popped[1] + diff[0],
            popped[2] + diff[1],
            popped[3],
            popped[4] | {(popped[1], popped[2])},
        )
        if grid[next[1]][next[2]] in ".E":
            heapq.heappush(min_heap, next)

        # 90 deg turns
        heapq.heappush(
            min_heap,
            (popped[0] + 1000, popped[1], popped[2], (popped[3] + 1) % 4, popped[4]),
        )
        heapq.heappush(
            min_heap,
            (popped[0] + 1000, popped[1], popped[2], (popped[3] - 1) % 4, popped[4]),
        )

    if part == 1:
        return best_score

    return len(final_path_coords)


example = """###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############"""

example2 = """#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################"""

input = open("input.txt").read().strip()

print(f"day16 example: {day16(example, 1)} want 7036")
print(f"day16 example2: {day16(example2,1)} want 11048")
print(f"day16 actual: {day16(input,1)} want 83444")

print(f"part2 example: {day16(example,2)} want 45")
print(f"part2 example2: {day16(example2,2)} want 64")
print(f"part2 actual: {day16(input,2)} want 483")
