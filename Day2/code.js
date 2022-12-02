const input = "<input from file>".trim().split('\n')

// Part 1
const map1 = {
  'A X': 4,
  'A Y': 8,
  'A Z': 3,
  'B X': 1,
  'B Y': 5,
  'B Z': 9,
  'C X': 7,
  'C Y': 2,
  'C Z': 6
}

const sum1 = input.reduce((sum, outcome) => sum + map1[outcome], 0)
console.log("Answer to part 1 -> sum1 =", sum1)
// Answer to part 1 -> sum1 = 13565

// Part 2
const map2 = {
  'A X': 3,
  'A Y': 4,
  'A Z': 8,
  'B X': 1,
  'B Y': 5,
  'B Z': 9,
  'C X': 2,
  'C Y': 6,
  'C Z': 7
}

const sum2 = input.reduce((sum, outcome) => sum + map2[outcome], 0)
console.log("Answer to part 2 -> sum2 =", sum2)
// Answer to part 2 -> sum2 = 12424
