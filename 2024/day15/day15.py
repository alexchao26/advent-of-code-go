diff_map: dict[str, tuple[int, int]] = {
    "^": (-1, 0),
    "v": (1, 0),
    "<": (0, -1),
    ">": (0, 1),
}


def part1(input: str) -> int:
    input_parts = input.split("\n\n")
    grid = [list(line) for line in input_parts[0].splitlines()]

    row: int = 0
    col: int = 0

    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "@":
                row = r
                col = c

    instructions = "".join(input_parts[1].splitlines())
    for inst in instructions:
        diff = diff_map[inst]
        next_row, next_col = row + diff[0], col + diff[1]

        match grid[next_row][next_col]:
            case "#":
                # blocked
                continue
            case ".":
                grid[row][col] = "."
                grid[next_row][next_col] = "@"
                row, col = next_row, next_col
            case "O":
                # attempt push, keep moving in direction of diff until a . or # is hit...
                not_obstacle_row, not_obstacle_col = next_row, next_col
                while grid[not_obstacle_row][not_obstacle_col] == "O":
                    not_obstacle_row += diff[0]
                    not_obstacle_col += diff[1]

                # if it's a wall "#", nothing moves, so only check for empty spaces "."
                if grid[not_obstacle_row][not_obstacle_col] == ".":
                    grid[not_obstacle_row][not_obstacle_col] = "O"
                    grid[next_row][next_col] = "@"
                    grid[row][col] = "."
                    row, col = next_row, next_col
            case _:
                raise Exception("unhandled grid type: ", grid[next_row][next_col])

    # print("\n".join(["".join(row) for row in grid]))

    # 100 times its distance from the top edge of the map plus its distance from the left edge of the map
    # 0-indexed so same as indices in 2D array
    gps_sum: int = 0
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "O":
                gps_sum += r * 100 + c

    return gps_sum


def part2(input: str) -> int:
    input_parts = input.split("\n\n")

    input_parts[0] = input_parts[0].replace("O", "[]")
    input_parts[0] = input_parts[0].replace(".", "..")
    input_parts[0] = input_parts[0].replace("#", "##")
    input_parts[0] = input_parts[0].replace("@", "@.")

    grid = [list(line) for line in input_parts[0].splitlines()]

    row: int = 0
    col: int = 0

    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "@":
                row = r
                col = c

    # print("\n".join(["".join(row) for row in grid]))

    instructions = "".join(input_parts[1].splitlines())
    for inst in instructions:
        # first see what's in the next space we want to move into...
        diff = diff_map[inst]
        next_row, next_col = row + diff[0], col + diff[1]
        match grid[next_row][next_col]:
            case "#":
                continue
            case ".":
                grid[row][col] = "."
                row += diff[0]
                col += diff[1]
                grid[row][col] = "@"
            case "[" | "]":
                # handle left and right separately because they only push one row of boxes
                # at this point row == next_row
                if inst in "<>":
                    # copy-pasta from part 1
                    not_obstacle_col = next_col
                    while grid[next_row][not_obstacle_col] in "[]":
                        not_obstacle_col += diff[1]

                    # if it's a wall "#", nothing moves, so only check for empty spaces "."
                    # move everything over towards the diff direction
                    if grid[next_row][not_obstacle_col] == ".":
                        for c in range(not_obstacle_col, col, -1 * diff[1]):
                            grid[row][c] = grid[row][c - diff[1]]

                        # move robot..
                        grid[next_row][next_col] = "@"
                        grid[row][col] = "."
                        row, col = next_row, next_col
                else:
                    # push boxes up or down, use a stack to maintain the reverse order of boxes that
                    # MAY get moved
                    coords_to_check: list[tuple[int, int]] = [
                        (row, col),
                    ]

                    to_check_index: int = 0
                    is_blocked: bool = False
                    while to_check_index < len(coords_to_check):
                        front = coords_to_check[to_check_index]
                        to_check_index += 1

                        # variable masking is no bueno, but i guess it's fine here
                        next_row, next_col = front[0] + diff[0], front[1] + diff[1]
                        next_value = grid[next_row][next_col]
                        if next_value == "#":
                            is_blocked = True
                            break
                        if next_value == "[":
                            coords_to_check.append((next_row, next_col))
                            coords_to_check.append((next_row, next_col + 1))
                        if next_value == "]":
                            coords_to_check.append((next_row, next_col))
                            coords_to_check.append((next_row, next_col - 1))
                        # if "." then do nothing... no need to check

                    # move everything towards diff if nothing was blocked
                    # in reverse order to avoid overwrites
                    if not is_blocked:
                        # didn't prevent duplicate adds in the "checking" stage, so just no-op them here
                        moved_set: set[tuple[int, int]] = set()
                        for coord in reversed(coords_to_check):
                            if coord in moved_set:
                                continue
                            moved_set.add(coord)

                            next_row, next_col = coord[0] + diff[0], coord[1] + diff[1]
                            grid[next_row][next_col], grid[coord[0]][coord[1]] = (
                                grid[coord[0]][coord[1]],
                                grid[next_row][next_col],
                            )
                        row, col = next_row, next_col

            case _:
                raise Exception("unexpected grid value: ", grid[next_row][next_col])
        # print("\n".join(["".join(row) for row in grid]))

    gps_sum: int = 0
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == "[":
                gps_sum += r * 100 + c

    return gps_sum


small_example = """########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<"""

example = """##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^"""

input = open("input.txt").read().strip()

print(f"part1 small_example: {part1(small_example)} want 2028")
print(f"part1 example: {part1(example)} want 10092")
print(f"part1: {part1(input)} want 1413675")

print(f"part2 example: {part2(example)} want 9021")
print(f"part2: {part2(input)} want 1399772")
