from dataclasses import dataclass
from typing import Self


@dataclass
class TrieNode:
    char: str
    is_terminator: bool
    children: dict[str, Self]


def day19(input: str, part: int) -> int:
    # trie data structure works nicely... but is it necessary for part 2...
    input_lines = input.splitlines()

    root = TrieNode("", False, {})
    for towel in input_lines[0].split(", "):
        iter = root
        for i, char in enumerate(towel):
            if char not in iter.children:
                iter.children[char] = TrieNode(char, i == len(towel) - 1, {})
            iter = iter.children[char]
            # update for matching towels that are smaller than a pervious one
            # e.g. "towel" then "tow" needs to update the "w" node
            iter.is_terminator |= i == len(towel) - 1

    possible_towels: int = 0
    total: int = 0

    memo: dict[str, int] = {}
    for maybe_towel in input_lines[2:]:
        combos = is_possible_towel(root, maybe_towel, memo)
        if combos > 0:
            possible_towels += 1
            total += combos

    if part == 1:
        return possible_towels

    return total


# memo optimization added for part 2
def is_possible_towel(trie: TrieNode, towel: str, memo: dict[str, int]) -> int:
    if towel == "":
        return 1
    if towel in memo:
        return memo[towel]

    iter = trie
    total: int = 0
    for i, char in enumerate(towel):
        if char not in iter.children:
            break

        iter = iter.children[char]

        # if iterated to a terminator, we can restart recursively and if that
        # "works" (aka returns positive number of possible combos), then we can
        # sum the combos and then continue along this for loop to account for
        # larger towels that do not end at this index
        if iter.is_terminator:
            total += is_possible_towel(trie, towel[i + 1 :], memo)

    memo[towel] = total

    return total


example = """r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb"""

input = open("input.txt").read().strip()

print(f"day19 example: {day19(example, 1)}, want 6")
print(f"day19 actual: {day19(input, 1)}, want 216")

print(f"part2 example: {day19(example ,2)}, want 16")
print(f"part2 actual: {day19(input ,2)}, want 603191454138773")
