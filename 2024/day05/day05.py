from collections import defaultdict


def day5(input: str, part: int) -> int:
    split_input = input.split("\n\n")

    reversed_graph: defaultdict[int, list[int]] = defaultdict(list)
    for line in split_input[0].splitlines():
        rule_parts = line.split("|")
        X, Y = int(rule_parts[0]), int(rule_parts[1])
        reversed_graph[Y].append(X)

    updates: list[list[int]] = [
        [int(x) for x in line.split(",")] for line in split_input[1].splitlines()
    ]
    part1_ans: int = 0

    # for part 2
    invalid_updates: list[list[int]] = []

    for update in updates:
        # assuming there are no repeat updates...
        seen_set: set[int] = set()
        disallowed_set: set[int] = set()
        is_valid = True

        for num in update:
            if num in disallowed_set:
                is_valid = False
                break
            for cannot_come_before in reversed_graph[num]:
                if not cannot_come_before in seen_set:
                    disallowed_set.add(cannot_come_before)

            seen_set.add(num)

        if is_valid:
            part1_ans += update[len(update) // 2]
        else:
            invalid_updates.append(update)
    if part == 1:
        return part1_ans

    # part2
    # can assume there's only one valid order... or that unordered ones will not affect the middle value
    # so just take all the numbers and create the correct order by traversing the graph of dependencies?

    part2_ans: int = 0
    for update in invalid_updates:
        correct_order: list[int] = []
        all_nums: set[int] = set(update)
        used_nums: set[int] = set()
        while len(all_nums) > len(used_nums):
            for num in all_nums:
                if num in used_nums:
                    continue
                if not num in reversed_graph:
                    correct_order.append(num)
                    # all_nums.remove(num)
                    used_nums.add(num)
                    continue

                # check if all of this num's dependencies are in used_nums
                # or if its dependencies are not present at all
                # there's definitely a better algo for this...
                if all(
                    [
                        dep in used_nums or not dep in all_nums
                        for dep in reversed_graph[num]
                    ]
                ):
                    correct_order.append(num)
                    # all_nums.remove(num)
                    used_nums.add(num)
                    continue
        part2_ans += correct_order[len(correct_order) // 2]

    return part2_ans


# page ordering rules, X|Y -> X must come before Y
# updates, basically orders of pages
#   missing page numbers are ignored..
# part1_ans: sum middle page of each valid update
example = """47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47"""

input = open("input.txt").read()

print(f"day5 example: {day5(example,1)} want 143")
print(f"day5: {day5(input,1)} want 5166")

print(f"part2 example: {day5(example,2)} want 123")
print(f"part2: {day5(input,2)} want 4679")
