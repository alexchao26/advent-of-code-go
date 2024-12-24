from collections import defaultdict


def part1(input: str) -> int:
    grid = [list(line) for line in input.splitlines()]
    map_antenna_to_coords: dict[str, list[tuple[int, int]]] = defaultdict(list)
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] != ".":
                map_antenna_to_coords[grid[i][j]].append((i, j))

    rows = len(grid)
    cols = len(grid[0])

    anti_node_coords: set[tuple[int, int]] = set()

    for _, coords in map_antenna_to_coords.items():

        for i in range(len(coords)):
            c1 = coords[i]
            for j in range(i + 1, len(coords)):
                c2 = coords[j]

                if c1 == c2:
                    continue

                rowDiff: int = c2[0] - c1[0]
                colDiff: int = c2[1] - c1[1]

                maybeRow: int = c1[0] - rowDiff
                maybeCol: int = c1[1] - colDiff
                if (
                    0 <= maybeRow < rows
                    and 0 <= maybeCol < cols
                    and (maybeRow, maybeCol) != c1
                    and (maybeRow, maybeCol) != c2
                ):
                    anti_node_coords.add((maybeRow, maybeCol))

                maybeRow: int = c2[0] + rowDiff
                maybeCol: int = c2[1] + colDiff
                if (
                    0 <= maybeRow < rows
                    and 0 <= maybeCol < cols
                    and (maybeRow, maybeCol) != c1
                    and (maybeRow, maybeCol) != c2
                ):
                    anti_node_coords.add((maybeRow, maybeCol))

    return len(anti_node_coords)


def part2(input: str) -> int:
    grid = [list(line) for line in input.splitlines()]
    map_antenna_to_coords: dict[str, list[tuple[int, int]]] = defaultdict(list)
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] != ".":
                map_antenna_to_coords[grid[i][j]].append((i, j))

    rows = len(grid)
    cols = len(grid[0])

    anti_node_coords: set[tuple[int, int]] = set()

    for _, coords in map_antenna_to_coords.items():

        for i in range(len(coords)):
            c1 = coords[i]

            # a more elegant solution would be to calculate the starting point of the antenna's line
            # then do a single loop to find all anti-node coords
            # instead, i'll just go left and right from c1, and add c1 manually
            anti_node_coords.add(c1)

            for j in range(i + 1, len(coords)):
                c2 = coords[j]

                if c1 == c2:
                    continue

                rowDiff: int = c2[0] - c1[0]
                colDiff: int = c2[1] - c1[1]

                maybeRow: int = c1[0] - rowDiff
                maybeCol: int = c1[1] - colDiff
                while 0 <= maybeRow < rows and 0 <= maybeCol < cols:
                    anti_node_coords.add((maybeRow, maybeCol))
                    maybeRow -= rowDiff
                    maybeCol -= colDiff

                maybeRow: int = c1[0] + rowDiff
                maybeCol: int = c1[1] + colDiff
                while 0 <= maybeRow < rows and 0 <= maybeCol < cols:
                    anti_node_coords.add((maybeRow, maybeCol))
                    maybeRow += rowDiff
                    maybeCol += colDiff

    return len(anti_node_coords)


example = """............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............"""

input = open("input.txt").read()

print(f"part1 example: {part1(example)} want 14")
print(f"part1: {part1(input)} want 396")
print(f"part2 example: {part2(example)} want 34")
print(f"part2: {part2(input)} want 1200")
