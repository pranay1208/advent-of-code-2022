stacks1 = [[] for _ in range(9)]
stacks2 = [[] for _ in range(9)]

with open("input.txt") as file:
    fileText = file.read().split('\n')
    moves = fileText[10:]
    stackText = fileText[:8]
    stackText.reverse()

    for row in stackText:
        for i in range(9):
            if i == 0:
                char = row[1]
            else:
                char = row[4*i + 1]
            
            if char != " ":
                stacks1[i].append(char)
                stacks2[i].append(char)

# Part 1        
for command in moves:
    parts = command.strip().split(' ')
    numToMove = int(parts[1])
    colFrom = int(parts[3]) - 1
    colTo = int(parts[5]) - 1

    for _ in range(numToMove):
        charToMove = stacks1[colFrom].pop()
        stacks1[colTo].append(charToMove)

print("The answer to part 1 is:", end=" ")
for col in stacks1:
    print(col[-1], end="")
# The answer to part 1 is: ZRLJGSCTR

print()
# Part 2
for command in moves:
    parts = command.strip().split(' ')
    numToMove = int(parts[1])
    colFrom = int(parts[3]) - 1
    colTo = int(parts[5]) - 1

    charsToMove = stacks2[colFrom][len(stacks2[colFrom]) - numToMove:]
    stacks2[colFrom] = stacks2[colFrom][:len(stacks2[colFrom]) - numToMove]
    stacks2[colTo].extend(charsToMove)

print("The answer to part 2 is:", end=" ")
for col in stacks2:
    print(col[-1], end="")
# The answer to part 2 is: PRTTGRFPB