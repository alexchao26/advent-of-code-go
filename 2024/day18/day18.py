def part1(input: str, grid_size: int, fallen_bytes: int) -> int:
    # simulate fallen_bytes
    # path find from top left to bottom right

    grid = [["." for _ in range(grid_size)] for _ in range(grid_size)]

    for line in input.splitlines()[:fallen_bytes]:
        parts = line.split(",")
        col = int(parts[0])
        row = int(parts[1])
        grid[row][col] = "#"

    queue: list[tuple[int, int, int]] = [(0, 0, 0)]
    visited: set[tuple[int, int]] = set()
    while len(queue) > 0:
        front = queue.pop(0)

        if front[:2] in visited:
            continue
        visited.add(front[:2])

        if front[0] == grid_size - 1 and front[1] == grid_size - 1:
            return front[2]
        for diff in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            next_row = front[0] + diff[0]
            next_col = front[1] + diff[1]
            if 0 <= next_row < grid_size and 0 <= next_col < grid_size:
                if grid[next_row][next_col] != "#":
                    queue.append((next_row, next_col, front[2] + 1))

    raise Exception("should return from loop")


def part2(input: str, grid_size: int) -> str:
    grid = [["." for _ in range(grid_size)] for _ in range(grid_size)]

    # drop all fallen bytes, then remove them selectively and attempt to continue
    # path finding from there?
    fallen_bytes: list[tuple[int, int]] = []
    for line in input.splitlines():
        parts = line.split(",")
        col = int(parts[0])
        row = int(parts[1])
        grid[row][col] = "#"
        fallen_bytes.append((row, col))

    # get set of all reachable coords from the starting point
    reachable: set[tuple[int, int]] = {(0, 0)}
    flood_fill_from(grid, (0, 0), reachable)

    # iterate in reverse through fallen bytes and if the byte is adjacent to a reachable
    # cell, then try to flood fill from the adjacent cell's neighbors
    for fallen in reversed(fallen_bytes):
        grid[fallen[0]][fallen[1]] = "."
        for diff in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
            next = (fallen[0] + diff[0], fallen[1] + diff[1])
            if next in reachable:
                if flood_fill_from(grid, next, reachable):
                    return f"{fallen[1]},{fallen[0]}"

    raise Exception("should return from loop")


def flood_fill_from(
    grid: list[list[str]], coord: tuple[int, int], reachable: set[tuple[int, int]]
) -> bool:
    for diff in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
        next_row = coord[0] + diff[0]
        next_col = coord[1] + diff[1]
        if (next_row, next_col) not in reachable:
            if 0 <= next_row < len(grid) and 0 <= next_col < len(grid):
                if grid[next_row][next_col] != "#":
                    reachable.add((next_row, next_col))
                    flood_fill_from(grid, (next_row, next_col), reachable)

    return (len(grid) - 1, len(grid) - 1) in reachable


# col, row (x, y)
example = """5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0"""

input = open("input.txt").read().strip()

print(f"part1 example: {part1(example, 6 + 1, 12)} want 22")
print(f"part1 actual: {part1(input, 70 + 1, 1024)} want 298")

print(f"part2 example: {part2(example, 6 + 1)} want '6,1'")
print(f"part2 actual: {part2(input, 70 + 1)} want 52,32")
