import re


def part1(input: str, part: int) -> int:
    ans: int = 0

    for machine in input.split("\n\n"):
        lines = machine.split("\n")
        a_parts: list[str] = re.findall(r"\d+", lines[0])
        Ax, Ay = int(a_parts[0]), int(a_parts[1])

        b_parts: list[str] = re.findall(r"\d+", lines[1])
        Bx, By = int(b_parts[0]), int(b_parts[1])

        prize_parts: list[str] = re.findall(r"\d+", lines[2])
        Px, Py = int(prize_parts[0]), int(prize_parts[1])

        if part == 2:
            Px += 10000000000000
            Py += 10000000000000

        # Ax * a + Bx * b = Px
        # Ay * a + By * b = Py
        # Solve for a and b...
        # a = (Px - Bx * b) / Ax
        # b * (By * Ax - Ay * Bx) = Py * Ax - Ay * Px
        b = (Py * Ax - Ay * Px) / (By * Ax - Ay * Bx)
        a = (Px - Bx * b) / Ax

        if b % 1 == 0 and a % 1 == 0:
            ans += 3 * a + b

    return int(ans)


# A 3, B 1
# moves right along X, forward along Y

example = """Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279"""

input = open("input.txt").read().strip()

print(f"part1 example: {part1(example, 1)} want 480")
print(f"part1: {part1(input, 1)} want 28059")

# don't think this result was given in the prompt
print(f"part1 example: {part1(example, 2)} want 875318608908")
print(f"part1: {part1(input, 2)} want 102255878088512")
