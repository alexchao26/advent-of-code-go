import re


def part1(input: str) -> str:
    matches = re.findall(r"\d+", input)

    reg_a: int = int(matches[0])
    reg_b: int = int(matches[1])
    reg_c: int = int(matches[2])

    program: list[int] = [int(x) for x in matches[3:]]

    out = run(program, reg_a, reg_b, reg_c)
    return ",".join([str(v) for v in out])


def run(program: list[int], reg_a: int, reg_b: int, reg_c: int) -> list[int]:
    def get_combo_operand(combo: int) -> int:
        if 0 <= combo <= 3:
            return combo
        if combo == 4:
            return reg_a
        if combo == 5:
            return reg_b
        if combo == 6:
            return reg_c
        raise Exception("unexpected combo value: ", combo)

    i: int = 0
    out: list[int] = []
    while i < len(program):
        opcode = program[i]
        operand = program[i + 1]
        match opcode:
            case 0:
                reg_a = reg_a // (2 ** get_combo_operand(operand))
                i += 2
            case 1:
                reg_b = reg_b ^ operand
                i += 2
            case 2:
                reg_b = get_combo_operand(operand) % 8
                i += 2
            case 3:
                if reg_a != 0:
                    i = operand
                else:
                    i += 2
            case 4:
                reg_b = reg_b ^ reg_c
                i += 2
            case 5:
                out.append(get_combo_operand(operand) % 8)
                i += 2
            case 6:
                reg_b = reg_a // (2 ** get_combo_operand(operand))
                i += 2
            case 7:
                reg_c = reg_a // (2 ** get_combo_operand(operand))
                i += 2
            case _:
                raise Exception("unhandled opcode", opcode)

    return out


# only works on actual input
def part2() -> int:
    matches = re.findall(r"\d+", input)
    program: list[int] = [int(x) for x in matches[3:]]

    # generate each element of the program one at a time, starting at the end
    output = str
    reg_as: list[int] = []
    for i in range(1, 8):
        output = run_optimized(i)
        if output == program[-len(output) :]:
            reg_as.append(i)

    digit_count = 1
    while digit_count < 16:
        next_reg_As: list[int] = []
        for a in reg_as:
            a *= 8
            for i in range(8):
                output = run_optimized(a + i)
                if output == program[-len(output) :]:
                    next_reg_As.append(a + i)
        reg_as = next_reg_As
        digit_count += 1

    # first reg_as will be smallest
    return reg_as[0]


def run_optimized(a: int) -> list[int]:
    b: int = 0
    output: list[int] = []
    while a != 0:
        # 2,4, 1,1, 7,5, 4,0, 0,3, 1,6, 5,5, 3,0
        # pen and paper "algebra"
        b = ((a % 8) ^ 1) ^ (a // (2 ** ((a % 8) ^ 1))) ^ 6
        output.append(b % 8)
        a = a // 8
    return output


example = """Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0"""

input = open("input.txt").read().strip()

print(f"part1 example: {part1(example)} want '4,6,3,5,6,3,5,2,1,0'")
print(f"part1 actual: {part1(input)} want '1,6,3,6,5,6,5,1,7'")

print(
    f"optimized {",".join([str(x) for x in run_optimized(30899381)])} want {part1(input)}"
)


example_part2 = """Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0"""

# print(f"part2 example: {part1(example_part2, 2)} want 117440")
print(f"part2 actual: {part2()} want 247839653009594")
