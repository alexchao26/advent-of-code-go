def day7(input: str, part: int) -> int:
    ans: int = 0

    lines = input.splitlines()
    for line in lines:
        parts = line.split(": ")
        target = int(parts[0])
        nums = [int(x) for x in parts[1].split(" ")]

        if recurse(part, nums, target, 0, 0):
            ans += target

    return ans


# brute force all the combinations and return True early if the target is made using all elements of nums array
def recurse(part: int, nums: list[int], target: int, current: int, index: int) -> bool:
    if target < current:
        return False

    if index == len(nums) and current == target:
        return True
    if index == len(nums) and current != target:
        return False

    if index == 0:
        current = 1
    multiply_result = recurse(part, nums, target, current * nums[index], index + 1)
    if multiply_result:
        return True

    # attempt concatenation operation for part 2
    if part == 2:
        concat_num = int(str(current) + str(nums[index]))
        concat_result = recurse(part, nums, target, concat_num, index + 1)
        if concat_result:
            return True

    if index == 0:
        current = 0
    return recurse(part, nums, target, current + nums[index], index + 1)


example = """190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20"""

input = open("input.txt").read()

print(f"part1 example: {day7(example,1)} want 3749")
print(f"part1: {day7(input,1)} want 66343330034722")

print(f"part2 example: {day7(example,2)} want 11387")
print(f"part2: {day7(input,2)} want 637696070419031")  # a bit slow but not bad
