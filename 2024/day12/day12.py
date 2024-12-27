from collections import defaultdict


def day12(input: str, part: int) -> int:

    grid = [list(line) for line in input.splitlines()]
    visited: set[tuple[int, int]] = {}
    cost: int = 0

    for r in range(len(grid)):
        for c in range(len(grid[0])):
            coord = (r, c)
            if coord not in visited:
                island_coords: set[tuple[int, int]] = set()

                flood_fill_island(grid, r, c, visited, island_coords)
                if part == 1:
                    cost += len(island_coords) * get_perimeter_of_island(island_coords)
                elif part == 2:
                    edge_count: int = get_edge_count_of_island(island_coords)
                    cost += len(island_coords) * edge_count
                else:
                    raise ("unexpected part")

    return cost


# refactored to just populate the entire island_coords set and visited set
# just populates the island_coords so does not need to return anything
def flood_fill_island(
    grid: list[list[str]],
    row: int,
    col: int,
    visited: set[tuple[int, int]],
    island_coords: set[tuple[int, int]],
):
    visited[(row, col)] = True
    island_coords.add((row, col))

    for diff in [(0, -1), (0, 1), (-1, 0), (1, 0)]:
        next_row = row + diff[0]
        next_col = col + diff[1]

        if 0 <= next_row < len(grid) and 0 <= next_col < len(grid[0]):
            if (next_row, next_col) in visited:
                continue
            # if in range, check if neighbor is same to recurse
            if grid[next_row][next_col] == grid[row][col]:
                # if does match and unvisited, recurse
                flood_fill_island(grid, next_row, next_col, visited, island_coords)


# for part1 cost calculation
def get_perimeter_of_island(island_coords: set[tuple[int, int]]) -> int:
    perimeter: int = 0
    for coord in island_coords:
        for diff in [(0, -1), (0, 1), (-1, 0), (1, 0)]:
            next_coord = (coord[0] + diff[0], coord[1] + diff[1])
            if next_coord not in island_coords:
                perimeter += 1

    return perimeter


def get_edge_count_of_island(island_coords: set[tuple[int, int]]) -> int:
    edges: int = 0

    map_dir_to_coord: dict[tuple[int, int], set[tuple[int, int]]] = defaultdict(set)

    for coord in island_coords:
        for diff in [(0, -1), (0, 1), (-1, 0), (1, 0)]:
            # this coord has already been accounted for in an edge
            if coord in map_dir_to_coord[diff]:
                continue

            next_coord = (coord[0] + diff[0], coord[1] + diff[1])
            # not in the island means it is bordering an edge...
            if next_coord not in island_coords:
                # collect all coords that make up this same edge, basically we need to go perpendicular to diff
                collect_all_coords_on_edge(island_coords, coord, diff, map_dir_to_coord)
                edges += 1

    return edges


perpendicular_dirs: dict[tuple[int, int], list[tuple[int, int]]] = {
    (0, -1): [(-1, 0), (1, 0)],
    (0, 1): [(-1, 0), (1, 0)],
    (-1, 0): [(0, -1), (0, 1)],
    (1, 0): [(0, -1), (0, 1)],
}


def collect_all_coords_on_edge(
    island_coords: set[tuple[int, int]],
    coord: tuple[int, int],
    empty_dir_diff: tuple[int, int],
    map_dir_to_coord: dict[tuple[int, int], set[tuple[int, int]]],
):
    # mark self
    map_dir_to_coord[empty_dir_diff].add(coord)

    for perp in perpendicular_dirs[empty_dir_diff]:
        next_coord = (coord[0] + perp[0], coord[1] + perp[1])
        # we're collecting the entire connected edge that is facing a single direction,
        # so stop checking if we're "off" the island
        # do not need to check if we're inside the grid because we can just leverage island_coords
        if next_coord not in island_coords:
            continue
        # if already visited for facing this direction, we can also skip
        if next_coord in map_dir_to_coord[empty_dir_diff]:
            continue

        # continue if next_coord is not a contiguous part of the edge
        if (next_coord[0] + empty_dir_diff[0], next_coord[1] + empty_dir_diff[1]) in island_coords:
            continue

        # need to continue exploring recursively
        collect_all_coords_on_edge(
            island_coords, next_coord, empty_dir_diff, map_dir_to_coord
        )


small_example = """AAAA
BBCD
BBCC
EEEC"""


example = """RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE"""

part2_example = """AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA"""

input = open("input.txt").read().strip()

print(f"part1 small_example: {day12(small_example, 1)} want 140")
print(f"part1 example: {day12(example, 1)} want 1930")
print(f"part1: {day12(input, 1)} want 1473408")

print(f"part2 small_example: {day12(small_example, 2)} want 80")
# this one helped debug the disjointed edges bug (second and fourth rows pointing east)
print(f"part2 example: {day12("""EEEEE
EXXXX
EEEEE
EXXXX
EEEEE""", 2)} want 236")
print(f"part2 part2_example: {day12(part2_example, 2)} want 368")
print(f"part2 example: {day12(example, 2)} want 1206")
print(f"part2: {day12(input, 2)} want 886364")

