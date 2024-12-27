import re


def part1(input: str, wide: int, tall: int, seconds: int) -> int:
    # don't need to model this exactly right? just equate a line, mod by grid size...
    # i guess for 100 steps it would've been better to just model it...

    quad_counts: list[int] = [0] * 4

    for line in input.splitlines():
        nums = re.findall(r"-?\d+", line)
        assert len(nums) == 4

        x = int(nums[0]) + int(nums[2]) * seconds
        y = int(nums[1]) + int(nums[3]) * seconds

        x %= tall
        y %= wide

        if x < tall // 2 and y < wide // 2:
            # print("top left")
            quad_counts[0] += 1
        elif x < tall // 2 and y > wide // 2:
            # print("top right")
            quad_counts[1] += 1
        elif x > tall // 2 and y < wide // 2:
            # print("bottom left")
            quad_counts[2] += 1
        elif x > tall // 2 and y > wide // 2:
            # print("bottom right")
            quad_counts[3] += 1
        # else:
        #     print("on mid line")

    # multiply robot count in each quadrant
    # does not include robots on mid-lines
    return quad_counts[0] * quad_counts[1] * quad_counts[2] * quad_counts[3]


def part2(input: str, wide: int, tall: int) -> int:
    # i guess for 100 steps it would've been better to just model it...
    # then i could have reused it for part 2...

    robots: list[list[int]] = []
    for line in input.splitlines():
        nums = re.findall(r"-?\d+", line)
        assert len(nums) == 4
        robots.append([int(x) for x in nums])

    for s in range(10_000):

        neighbors: set[tuple[int, int]] = set()
        matched: set[tuple[int, int]] = set()
        for robot in robots:
            robot[0] += robot[2]
            robot[1] += robot[3]

            robot[0] %= tall
            robot[1] %= wide

            # track how many neighboring robots have been placed so far
            # when the number is high enough we probably have some image of a tree
            coord = (robot[0], robot[1])
            if coord in neighbors:
                matched.add(coord)
            for dx in [-1, 0, 1]:
                for dy in [-1, 0, 1]:
                    neighbors.add((coord[0] + dx, coord[1] + dy))

            if len(matched) >= 250:
                print_grid(robots, wide, tall)
                # print(len(matched))
                return s + 1

    raise Exception("should return from loop")


def print_grid(robots: list[list[int]], wide: int, tall: int):
    grid: list[list[str]] = []
    for _ in range(0, tall):
        grid.append([" "] * wide)

    for robot in robots:
        grid[robot[0]][robot[1]] = "X"

    for line in grid:
        print("".join(line))


example = """p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3"""

input = open("input.txt").read().strip()

print(f"part1 example: {part1(example, 7, 11, 100)} want 12")
print(f"part1 example: {part1(input, 103, 101, 100)} want 218965032")

print(f"part2: {part2(input, 103, 101)} want 7037")
