def part1(input: str) -> int:
    parts = input.split("\n\n")

    locks: list[list[int]] = []
    keys: list[list[int]] = []
    max_height: int = 0
    for part in parts:
        lines = part.split("\n")
        max_height = len(lines) - 1

        # lock
        if lines[0] == "#" * len(lines[0]):
            lock_heights: list[int] = []
            for col in range(len(lines[0])):
                for row in reversed(range(len(lines))):
                    if lines[row][col] != ".":
                        lock_heights.append(row)
                        break
            locks.append(lock_heights)
        else:
            # key
            key_heights: list[int] = []
            for col in range(len(lines[0])):
                for row in range(len(lines)):
                    if lines[row][col] != ".":
                        key_heights.append(len(lines) - 1 - row)
                        break
            keys.append(key_heights)

    count: int = 0

    for lock in locks:
        for key in keys:
            fits: bool = True
            for col in range(len(key)):
                if lock[col] + key[col] >= max_height:
                    fits = False
                    break
            if fits:
                count += 1

    return count


example = """#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####"""

print(f"part1 example: {part1(example)} want 3")

input = open("input.txt").read().strip()
print(f"part1 actual: {part1(input)} want 3114")
