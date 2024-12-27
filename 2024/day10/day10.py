def part1(input: str) -> int:
    grid: list[list[int, int]] = []
    for line in input.splitlines():
        grid.append([int(x) for x in list(line)])
    # alternative pythonic way...
    # grid = [[int(char) for char in line] for line in input.splitlines()]

    ans: int = 0
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == 0:
                ans += len(dfs_backtrack_unique_end_coords(grid, r, c, {(r, c)}))

    return ans


# I misread the instructions slightly and went for a "unique paths" algo at first.
# It would be simpler to pass the "reachable 9s coords" set in as an arg and update
#   it in the termination case, then take the length of that set in the part1() function.
#   This would remove the need to combine the sets which is potentially expensive (and ugly)
def dfs_backtrack_unique_end_coords(
    grid: list[list[int, int]], row: int, col: int, visited: set[tuple[int, int]]
) -> set[tuple[int, int]]:
    if grid[row][col] == 9:
        return {(row, col)}

    all_coords: set[tuple[int, int]] = set()
    for diff in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
        nextRow = row + diff[0]
        nextCol = col + diff[1]

        if not (0 <= nextRow < len(grid) and 0 <= nextCol < len(grid[0])):
            continue
        if (nextRow, nextCol) in visited:
            continue
        if grid[nextRow][nextCol] == grid[row][col] + 1:
            visited.add((nextRow, nextCol))
            new_coords = dfs_backtrack_unique_end_coords(
                grid, nextRow, nextCol, visited
            )
            # combines the two sets, see comment above function def
            all_coords.update(new_coords)
            visited.remove((nextRow, nextCol))

    return all_coords


def part2(input: str) -> int:
    grid = [[int(char) for char in line] for line in input.splitlines()]

    ans: int = 0
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == 0:
                ans += dfs_backtrack_unique_paths(grid, r, c, {(r, c)})

    return ans


def dfs_backtrack_unique_paths(
    grid: list[list[int, int]], row: int, col: int, visited: set[tuple[int, int]]
) -> int:
    if grid[row][col] == 9:
        # unique path found
        return 1

    total: int = 0
    for diff in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
        nextRow = row + diff[0]
        nextCol = col + diff[1]

        if not (0 <= nextRow < len(grid) and 0 <= nextCol < len(grid[0])):
            continue
        if (nextRow, nextCol) in visited:
            continue
        if grid[nextRow][nextCol] == grid[row][col] + 1:
            visited.add((nextRow, nextCol))
            total += dfs_backtrack_unique_paths(grid, nextRow, nextCol, visited)
            visited.remove((nextRow, nextCol))

    return total


example = """89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732"""

input = open("input.txt").read().strip()

print(f"part1 example: {part1(example)} want 36")
print(f"part1: {part1(input)} want 782")

print(f"part2 example: {part2(example)} want 81")
print(f"part2 example: {part2(input)} want 1694")
