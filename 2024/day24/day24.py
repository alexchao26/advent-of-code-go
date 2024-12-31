from collections import defaultdict
import copy


def part1(input: str) -> int:
    wires, instructions = parse_input(input)
    run(instructions, wires)
    return get_num(wires, "z")


def parse_input(input: str) -> tuple[dict[str, int], list[list[str]]]:
    parts = input.split("\n\n")

    wires: dict[str, int] = defaultdict(int)

    for line in parts[0].splitlines():
        wires[line[:3]] = int(line[5:])

    instructions: list[list[str]] = []
    for line in parts[1].splitlines():
        in1, op, in2, _, out = line.split(" ")
        instructions.append([in1, op, in2, out])

    return (wires, instructions)


def get_num(wires: dict[str, int], char: str) -> int:
    if len(char) != 1:
        raise Exception("exactly one char needed for get_num")
    keys: list[str] = []
    for wire in wires:
        if wire[0] == char:
            keys.append(wire)
    keys.sort()
    binary_num: str = ""
    for key in reversed(keys):
        binary_num += str(wires[key])

    return int(binary_num, 2)


def run(instructions: list[list[str]], wires: dict[str, int]):
    processed_lines: set[int] = set()
    while len(processed_lines) < len(instructions):
        processed_before: int = len(processed_lines)

        for i, inst in enumerate(instructions):
            if i in processed_lines:
                continue
            in1, op, in2, out = inst
            if in1 not in wires or in2 not in wires:
                continue

            if op == "AND":
                wires[out] = wires[in1] & wires[in2]
            elif op == "OR":
                wires[out] = wires[in1] | wires[in2]
            elif op == "XOR":
                wires[out] = wires[in1] ^ wires[in2]

            processed_lines.add(i)

        if len(processed_lines) == processed_before:
            # print("no new lines processed")
            return


# this code is way too slow, works for the examples but unsurprisingly way too
# slow for the actual input. ditching python for a go solution...
def part2(input: str) -> str:
    wires, instructions = parse_input(input)
    expected_sum: int = get_num(wires, "x") & get_num(wires, "y")

    def any_values_are_equal(*args: int) -> bool:
        seen: set[int] = set()
        for val in args:
            if val in seen:
                return True
            seen.add(val)
        return False

    # generate swaps... maybe 8 nested for loops...
    for a in range(len(instructions)):
        print(a)
        for b in range(a + 1, len(instructions)):
            for c in range(a + 1, len(instructions)):
                if any_values_are_equal(b, c):
                    continue
                for d in range(a + 1, len(instructions)):
                    if any_values_are_equal(b, c, d):
                        continue
                    for e in range(a + 1, len(instructions)):
                        if any_values_are_equal(b, c, d, e):
                            continue
                        for f in range(a + 1, len(instructions)):
                            if any_values_are_equal(b, c, d, e, f):
                                continue
                            for g in range(a + 1, len(instructions)):
                                if any_values_are_equal(b, c, d, e, f, g):
                                    continue
                                for h in range(a + 1, len(instructions)):
                                    if any_values_are_equal(b, c, d, e, f, g, h):
                                        continue

                                    instructions[a][3], instructions[b][3] = (
                                        instructions[b][3],
                                        instructions[a][3],
                                    )
                                    instructions[c][3], instructions[d][3] = (
                                        instructions[d][3],
                                        instructions[c][3],
                                    )

                                    instructions[e][3], instructions[f][3] = (
                                        instructions[f][3],
                                        instructions[e][3],
                                    )
                                    instructions[g][3], instructions[h][3] = (
                                        instructions[h][3],
                                        instructions[g][3],
                                    )

                                    wires_copy = copy.deepcopy(wires)
                                    run(instructions, wires_copy)
                                    sum = get_num(wires_copy, "z")
                                    if sum == expected_sum:
                                        swaps: list[str] = [
                                            instructions[a][3],
                                            instructions[b][3],
                                            instructions[c][3],
                                            instructions[d][3],
                                            instructions[e][3],
                                            instructions[f][3],
                                            instructions[g][3],
                                            instructions[h][3],
                                        ]
                                        swaps.sort()
                                        return ",".join(swaps)
                                    # backtrack
                                    instructions[a][3], instructions[b][3] = (
                                        instructions[b][3],
                                        instructions[a][3],
                                    )
                                    instructions[c][3], instructions[d][3] = (
                                        instructions[d][3],
                                        instructions[c][3],
                                    )
                                    instructions[e][3], instructions[f][3] = (
                                        instructions[f][3],
                                        instructions[e][3],
                                    )
                                    instructions[g][3], instructions[h][3] = (
                                        instructions[h][3],
                                        instructions[g][3],
                                    )

    raise Exception("should return from loop")


example = """x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02"""

print(f"part1 example: {part1(example)} want 4")

big_example = """x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj"""

print(f"part1 big_example: {part1(big_example)} want 2024")

input = open("input.txt").read().strip()
print(f"part1 actual: {part1(input)} want 38869984335432")

part2_example_2_pairs_swapped = """x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00"""

# print(
#     f"part2 example with two swapped {part2(part2_example_2_pairs_swapped)} want z00,z01,z02,z05"
# )

print(f"part2 actual: {part2(input)} want QQ")
