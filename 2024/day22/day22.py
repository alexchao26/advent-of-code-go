from collections import defaultdict


def day22(input: str, part: int) -> int:
    total: int = 0

    last_four_diffs_to_price: dict[tuple[int, int, int, int], int] = defaultdict(int)

    for line in input.splitlines():
        nums: list[int] = [int(line)]
        diffs: list[int] = []

        seen_last_fours: set[tuple[int, int, int, int]] = set()
        for _ in range(2000):
            nums.append(get_next_secret_number(nums[-1]))
            diff = nums[-1] % 10 - nums[-2] % 10
            diffs.append(diff)
        total += nums[-1]

        for i in range(3, len(diffs)):
            last_four_diffs = (diffs[i - 3], diffs[i - 2], diffs[i - 1], diffs[i])
            # only count the first time we've seen this set of diffs, aka only sell once
            if last_four_diffs in seen_last_fours:
                continue
            seen_last_fours.add(last_four_diffs)

            last_four_diffs_to_price[last_four_diffs] += nums[i + 1] % 10

    if part == 1:
        return total

    most_bananas = max(last_four_diffs_to_price.values())
    return most_bananas


def get_next_secret_number(num: int) -> int:
    num = (num ^ (num * 64)) % 16777216
    num = (num ^ (num // 32)) % 16777216
    num = (num ^ (num * 2048)) % 16777216
    return num


example = """1
10
100
2024"""

input = open("input.txt").read().strip()

print(f"part1 example: {day22(example, 1)} want 37327623")
print(f"part1 actual: {day22(input, 1)} want 17612566393")

example_part_2 = """1
2
3
2024"""
print(f"part2 example: {day22(example_part_2, 2)} want 23")
print(f"part2 actual: {day22(input, 2)} want 1968")
