from typing import List

with open("./input.txt") as file:
    lines = file.read().strip().split('\n')
    map = [list(line) for line in lines]


class Node:
    def __init__(self, rowNum: int, colNum: int, pathCost: int) -> None:
        self.rowNum = rowNum
        self.colNum = colNum
        self.elem = map[rowNum][colNum]
        self.pathCost = pathCost
        self.height = ord(self.elem) - 96
        if self.elem == 'E':
            self.height = 26
        if self.elem == 'S':
            self.height = 1

    def getChildren(self):
        children = []
        if self.rowNum != 0:
            children.append(
                Node(self.rowNum - 1, self.colNum, self.pathCost+1))
        if self.rowNum != len(map) - 1:
            children.append(
                Node(self.rowNum + 1, self.colNum, self.pathCost+1))
        if self.colNum != 0:
            children.append(
                Node(self.rowNum, self.colNum - 1, self.pathCost+1))
        if self.colNum != len(map[self.rowNum]) - 1:
            children.append(
                Node(self.rowNum, self.colNum + 1, self.pathCost+1))
        return children

    # heuristic is manhattan distance to E multiplied by (27-elem)
    def getHeuristic1(self, endPos):
        cost = self.pathCost
        distance = abs(endPos[0] - self.rowNum) + abs(endPos[0] - colNum)
        charOrd = self.height
        if self.elem == 'E':
            charOrd = 27
        return cost + distance + (27 - charOrd)

    def getHeuristic2(self):
        cost = self.pathCost
        return cost + self.height - 1


for rowNum in range(len(map)):
    for colNum in range(len(map[rowNum])):
        if (map[rowNum][colNum] == 'S'):
            startNode = Node(rowNum, colNum, 0)
        if (map[rowNum][colNum] == 'E'):
            endPosition = (rowNum, colNum)
            endNode = Node(rowNum, colNum, 0)

openNodes = [startNode]
visited = []
node = openNodes.pop(0)

def alreadyVisited(n: Node):
    for node in visited:
        if node.rowNum == n.rowNum and node.colNum == n.colNum:
            return True
    return False

def putInOpenNodes(n: Node):
    for index in range(len(openNodes)):
        oN = openNodes[index]
        if oN.rowNum == n.rowNum and oN.colNum == n.colNum:
            if oN.pathCost <= n.pathCost:
                return
            openNodes.pop(index)
            break
    openNodes.append(n)

# part 1
while node.elem != 'E':
    visited.append(node)
    # get children of current node
    children = node.getChildren()
    for child in children:
        if child.height - 1 > node.height:
            continue
        if alreadyVisited(child):
            continue
        putInOpenNodes(child)

    minNodeIndex = 0
    # pick next node, i.e. min heuristic node
    for index in range(len(openNodes)):
        if openNodes[index].getHeuristic1(endPosition) <= openNodes[minNodeIndex].getHeuristic1(endPosition):
            minNodeIndex = index

    node = openNodes.pop(minNodeIndex)

print("Final cost to get to E is", node.pathCost)
# Final cost to get to E is 472

# part 2
openNodes = [endNode]
visited = []
node = openNodes.pop(0)
while node.elem != 'a':
    visited.append(node)
    # get children of current node
    children = node.getChildren()
    for child in children:
        if child.height + 1 < node.height:
            continue
        if alreadyVisited(child):
            continue
        putInOpenNodes(child)

    minNodeIndex = 0
    # pick next node, i.e. min heuristic node
    for index in range(len(openNodes)):
        if openNodes[index].getHeuristic2() <= openNodes[minNodeIndex].getHeuristic2():
            minNodeIndex = index

    node = openNodes.pop(minNodeIndex)

print("Final cost to get to a is", node.pathCost)
# Final cost to get to a is 465