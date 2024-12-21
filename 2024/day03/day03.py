import re


def part1(input: str) -> int:
    matches = re.findall(r"mul\(\d+,\d+\)", input)
    ans = 0
    for match in matches:
        ans += exec_mul(match)
    return ans


def part2(input: str) -> int:
    matches = re.findall(r"(do\(\)|don\'t\(\)|mul\(\d+,\d+\))", input)
    do = True
    ans = 0

    for match in matches:
        if match == "do()":
            do = True
        elif match == "don't()":
            do = False
        elif do:
            ans += exec_mul(match)

    return ans


def exec_mul(s: str) -> int:
    nums = re.findall(r"\d+", s)
    return int(nums[0]) * int(nums[1])


example = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

print("example part1:", part1(example), "want", 161)

input = open("input.txt").read()
print("part1:", part1(input), "want", 169021493)

example2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

print("example part2:", part2(example2), "want", 48)
print("part2:", part2(input), "want", 111762583)
