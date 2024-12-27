# If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
# If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
# If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.


# actually modelling this would be a pain in the ass, error-prone, and probably
# not fast enough for part 2 where there will presumably be more blinks
# instead if we only care about the final number of stones, we can just see
# how many stone each original stone splits into, and we can memoize it
def day11(input: str, blinks: int) -> int:
    total_stones: int = 0
    memo: dict[tuple[int, int], int] = {}
    for stone in input.split(" "):
        total_stones += calculate_final_stones_count(stone, blinks, memo)
    return total_stones


def calculate_final_stones_count(num_as_str: str, blinks_left: int, memo) -> int:
    key = (num_as_str, blinks_left)
    if key in memo:
        return memo[key]

    if blinks_left == 0:
        return 1

    total_stones: int = 0
    if num_as_str == "0":
        total_stones += calculate_final_stones_count("1", blinks_left - 1, memo)
    elif len(num_as_str) % 2 == 0:
        # convert back and forth again to get rid of leading zeroes
        left_num = str(int(num_as_str[: len(num_as_str) // 2]))
        right_num = str(int(num_as_str[len(num_as_str) // 2 :]))
        total_stones += calculate_final_stones_count(left_num, blinks_left - 1, memo)
        total_stones += calculate_final_stones_count(right_num, blinks_left - 1, memo)
    else:
        new_num: int = str(int(num_as_str) * 2024)
        total_stones += calculate_final_stones_count(new_num, blinks_left - 1, memo)

    memo[key] = total_stones

    return total_stones


example = """125 17"""
input = open("input.txt").read().strip()

print(f"part1 example: {day11(example, 6)} want 22")
print(f"part1 example: {day11(example, 25)} want 55312")
print(f"part1: {day11(input, 25)} want 189092")
print(f"part2: {day11(input, 75)} want 224869647102559")
