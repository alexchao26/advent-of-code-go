from collections import defaultdict


def part1(input: str) -> int:
    graph: dict[str, set[str]] = defaultdict(set)
    for line in input.splitlines():
        parts = line.split("-")
        graph[parts[0]].add(parts[1])
        graph[parts[1]].add(parts[0])

    ans_groups: set[str] = set()
    for node in graph:
        if node[0] != "t":
            continue
        neighbors = graph[node]
        for neighbor in neighbors:
            for neighbor2 in neighbors:
                if neighbor == neighbor2:
                    continue
                if neighbor2 in graph[neighbor]:
                    group: list[str] = [node, neighbor, neighbor2]
                    group.sort()

                    ans_groups.add(",".join(group))

    return len(ans_groups)


def part2(input: str) -> str:
    graph: dict[str, set[str]] = defaultdict(set)
    for line in input.splitlines():
        parts = line.split("-")
        graph[parts[0]].add(parts[1])
        graph[parts[1]].add(parts[0])

    seen: set[str] = set()
    largest_group: set[str] = set()
    for node in graph:
        if node in seen:
            continue
        seen.add(node)

        group: set[str] = {node}
        for neighbor in graph[node]:
            if check_group_against_new_node(group, graph, neighbor):
                group.add(neighbor)
                seen.add(neighbor)

        if len(group) > len(largest_group):
            largest_group = group

    final_group: list[str] = list(largest_group)
    final_group.sort()
    return ",".join(final_group)


def check_group_against_new_node(
    group: set[str], graph: dict[str, set[str]], node_to_add: str
) -> bool:
    for node in group:
        if node not in graph[node_to_add]:
            return False

    return True


example = """kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn"""

input = open("input.txt").read().strip()

print(f"part1 example: {part1(example)} want 7")
print(f"part1 actual: {part1(input)} want 1485")

print(f"part2 example: {part2(example)} want co,de,ka,ta")
print(f"part2 actual: {part2(input)} want cc,dz,ea,hj,if,it,kf,qo,sk,ug,ut,uv,wh")
