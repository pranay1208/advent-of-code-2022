def getRange(s: str):
    parts = s.split('-')
    return (int(parts[0]), int(parts[1]))

with open("./input.txt") as file:
    input = file.read().strip().split('\n')

overlaps = 0
for pair in input:
    p = pair.split(',')
    elf1 = getRange(p[0])
    elf2 = getRange(p[1])

    if elf1[0] <= elf2[0] and elf1[1] >= elf2[1]:
        overlaps += 1
        continue

    if elf2[0] <= elf1[0] and elf2[1] >= elf1[1]:
        overlaps += 1
        continue

print("Answer to part 1 is", overlaps)

fullOverlaps = 0
for pair in input:
    p = pair.split(',')
    elf1 = getRange(p[0])
    elf2 = getRange(p[1])

    if elf1[1] < elf2[0] or elf1[0] > elf2[1]:
        continue

    fullOverlaps += 1

print("Answer to part 2 is", fullOverlaps)