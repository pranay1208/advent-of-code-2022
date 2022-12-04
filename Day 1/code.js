const fs = require("fs");

const input = fs
  .readFileSync("./input.txt", "utf-8")
  .replaceAll("\r", "")
  .trim()
  .split("\n\n");

// Part 1
const sums = input.map((i) =>
  i
    .trim()
    .split("\n")
    .reduce((sum, val) => sum + parseInt(val), 0)
);
const max = sums.reduce((max, val) => (max > val ? max : val), 0);
console.log("Answer to part 1 -> max =", max);
// Answer to part 1 -> max = 67450

// Part 2
sums.sort((a, b) => b - a);
const maxTop3 = sums[0] + sums[1] + sums[2];
console.log("Answer to part 2 -> maxTop3 =", maxTop3);
// Answer to part 2 -> maxTop3 = 199357
