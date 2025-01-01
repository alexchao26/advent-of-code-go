num_pad: dict[str, tuple[int, int]] = {
    "7": (0, 0),
    "8": (0, 1),
    "9": (0, 2),
    "4": (1, 0),
    "5": (1, 1),
    "6": (1, 2),
    "1": (2, 0),
    "2": (2, 1),
    "3": (2, 2),
    "0": (3, 1),
    "A": (3, 2),
    "EMPTY": (3, 0),
}

dir_pad: dict[str, tuple[int, int]] = {
    "^": (0, 1),
    "A": (0, 2),
    "<": (1, 0),
    "v": (1, 1),
    ">": (1, 2),
    "EMPTY": (0, 0),
}


# had to manually memo because the pad dict type isn't hashable... i'm sure
# there's a way around this to use functools.cache like others...
def get_optimal_path_length(
    pad: dict[str, tuple[int, int]],
    start: str,
    end: str,
    depth: int,
    memo: dict[str, int],
) -> int:
    key = f"{start},{end},{depth},{pad is dir_pad}"
    if key in memo:
        return memo[key]

    # determine horizontal and vertical parts of path
    # recurse in each direction and return whichever is shorter
    start_row, start_col = pad[start]
    end_row, end_col = pad[end]

    dr: int = end_row - start_row
    dc: int = end_col - start_col

    horiz: str = ""
    vert: str = ""
    if dc > 0:
        horiz = ">" * dc
    else:
        horiz = "<" * -dc

    if dr > 0:
        vert = "v" * dr
    else:
        vert = "^" * -dr

    horiz_first = "A" + horiz + vert + "A"
    vert_first = "A" + vert + horiz + "A"

    if depth == 0:
        memo[key] = len(horiz_first[1:])
        return memo[key]

    horiz_first_result: int = 0
    for i in range(len(horiz_first) - 1):
        horiz_first_result += get_optimal_path_length(
            dir_pad, horiz_first[i], horiz_first[i + 1], depth - 1, memo
        )

    vert_first_result: int = 0
    for i in range(len(vert_first) - 1):
        vert_first_result += get_optimal_path_length(
            dir_pad, vert_first[i], vert_first[i + 1], depth - 1, memo
        )

    # still need to do the recursive calls, THEN decide which to return
    # no horizontal or vertical movement so obviously just need to do the move plus "A"
    if horiz == "" or vert == "":
        assert horiz_first_result == vert_first_result
        memo[key] = horiz_first_result
        return horiz_first_result

    # if it could hit the empty spot, just return the opposite direction first
    # moving horizontally
    if (start_row, end_col) == pad["EMPTY"]:
        memo[key] = vert_first_result
        return vert_first_result
    # moving vertically
    if (end_row, start_col) == pad["EMPTY"]:
        memo[key] = horiz_first_result
        return horiz_first_result

    if horiz_first_result < vert_first_result:
        memo[key] = horiz_first_result
        return horiz_first_result
    memo[key] = vert_first_result
    return vert_first_result


def part1(input: str, depth: int) -> int:

    ans: int = 0
    for line in input.splitlines():
        # starts at bottom right A button so add one onto the front of the line
        line = "A" + line

        path: int = 0
        for i in range(len(line) - 1):
            best_path = get_optimal_path_length(
                num_pad, line[i], line[i + 1], depth, {}
            )
            path += best_path

        # length of shortest sequence * numeric part of input
        # all are the same so just remove A char and convert to int
        ans += int(line[1:4]) * path
    return ans


example = """029A
980A
179A
456A
379A"""

input = open("input.txt").read().strip()

print(f"part1 example: {part1(example, 2)} want 126384")
print(f"part1 actual: {part1(input, 2)} want 237342")
print(f"part1 actual: {part1(input, 25)} want 237342")

# +---+---+---+
# | 7 | 8 | 9 |
# +---+---+---+
# | 4 | 5 | 6 |
# +---+---+---+
# | 1 | 2 | 3 |
# +---+---+---+
#     | 0 | A |
#     +---+---+


# 3 of these robots chained together to type into the pad above...
#     +---+---+
#     | ^ | A |
# +---+---+---+
# | < | v | > |
# +---+---+---+


# unfortunately had to scrap this method because memoizing the entire resulting
# string takes up way too much memory and that nukes the runtime...
# refactoring to just store the actual length is plenty... and possible to memo
def get_optimal_path(
    pad: dict[str, tuple[int, int]],
    start: str,
    end: str,
    depth: int,
) -> str:
    # determine horizontal and vertical parts of path
    # recurse in each direction and return whichever is shorter
    start_row, start_col = pad[start]
    end_row, end_col = pad[end]

    dr: int = end_row - start_row
    dc: int = end_col - start_col

    horiz: str = ""
    vert: str = ""
    if dc > 0:
        horiz = ">" * dc
    else:
        horiz = "<" * -dc

    if dr > 0:
        vert = "v" * dr
    else:
        vert = "^" * -dr

    horiz_first = "A" + horiz + vert + "A"
    vert_first = "A" + vert + horiz + "A"

    if depth == 0:
        return horiz_first[1:]

    horiz_first_result: str = ""
    for i in range(len(horiz_first) - 1):
        horiz_first_result += get_optimal_path(
            dir_pad, horiz_first[i], horiz_first[i + 1], depth - 1
        )

    vert_first_result: str = ""
    for i in range(len(vert_first) - 1):
        vert_first_result += get_optimal_path(
            dir_pad, vert_first[i], vert_first[i + 1], depth - 1
        )

    # still need to do the recursive calls, THEN decide which to return
    # no horizontal or vertical movement so obviously just need to do the move plus "A"
    if horiz == "" or vert == "":
        assert horiz_first_result == vert_first_result
        return horiz_first_result

    # if it could hit the empty spot, just return the opposite direction first
    # moving horizontally
    if (start_row, end_col) == pad["EMPTY"]:
        return vert_first_result
    # moving vertically
    if (end_row, start_col) == pad["EMPTY"]:
        return horiz_first_result

    if len(horiz_first_result) < len(vert_first_result):
        return horiz_first_result
    return vert_first_result
