const input = "<input from file>".trim().split('\n\n')

// Part 1
sums = input.map(i => i.trim().split('\n')).reduce((sum, val) => sum + parseInt(val), 0)
max = sums.reduce((max, val) => max > val ? max : val, 0)

// Answer to part 1 -> max = 67450

// Part 2
sums.sort((a,b) => b - a)
maxTop3 = sums[0] + sums[1] + sums[2]

// Answer to part 2 -> maxTop3 = 199357
