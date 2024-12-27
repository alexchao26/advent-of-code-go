from dataclasses import dataclass


# index by index brute force with a sliding window optimization to make it linear...
def part1(input: str) -> int:
    total_disk_space: int = 0
    for x in list(input):
        total_disk_space += int(x)

    file_system: list[int] = [-1] * total_disk_space
    is_file: bool = True
    index: int = 0
    file_number: int = 0

    for x in list(input):
        if not is_file:
            index += int(x)
            is_file = not is_file
        else:
            for _ in range(int(x)):
                if is_file:
                    file_system[index] = file_number
                index += 1

            file_number += 1

            is_file = not is_file

    # rearrange file to left via sliding window
    left: int = 0
    right: int = len(file_system) - 1
    while left < right:
        if file_system[right] == -1:
            right -= 1
        elif file_system[left] != -1:
            left += 1
        elif file_system[left] == -1:
            file_system[left], file_system[right] = (
                file_system[right],
                file_system[left],
            )
            left += 1

    # checksum is index in string * number value (file number)
    checksum: int = 0
    for i in range(len(file_system)):
        if file_system[i] == -1:
            break
        checksum += i * file_system[i]

    return checksum


@dataclass
class FileSystemSpace:
    start: int
    size: int
    file_number: int


def part2(input: str) -> int:
    files: list[FileSystemSpace] = []
    empty_spaces: list[FileSystemSpace] = []

    is_file: bool = True
    index: int = 0
    for size in [int(x) for x in list(input)]:
        file_or_empty = FileSystemSpace(index, size, len(files))
        index += size
        if is_file:
            files.append(file_or_empty)
        else:
            empty_spaces.append(file_or_empty)
        is_file = not is_file

    # brute force finding a space to move each file into
    for file in reversed(files):
        for empty_space in empty_spaces:
            # prevent moving files to higher spots in the file system...
            if empty_space.start > file.start:
                break

            # large enough empty space found
            if empty_space.size >= file.size:
                file.start = empty_space.start
                empty_space.start += file.size
                empty_space.size -= file.size

                break

    # print_util(files)

    checksum: int = 0
    for file in files:
        for x in range(file.size):
            checksum += (file.start + x) * file.file_number

    return checksum


def print_util(all_files: list[FileSystemSpace]):
    last_index: int = 0
    for file in all_files:
        last_index = max(last_index, file.start + file.size)

    fs: list[str] = ["."] * (last_index + 1)
    for file in all_files:
        for i in range(file.start, file.start + file.size):
            fs[i] = str(file.file_number)

    print("".join(fs))


example = """2333133121414131402"""
input = open("input.txt").read().strip()

print(f"part1 example: {part1(example)} want 1928")
print(f"part1: {part1(input)} want 6200294120911")

print(f"part2 example: {part2(example)} want 2858")
print(f"part2: {part2(input)} want 6227018762750")
