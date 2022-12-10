const fs = require("fs");

const input = fs
  .readFileSync("./input.txt", "utf-8")
  .replaceAll("\r", "")
  .trim()
  .split("\n");

const isCriticalCycle = (cycleNum) => (cycleNum - 20) % 40 === 0;

let cycleStrength = 0;
let registerValue = 1;
let cycleNumber = 1;
let crtImage = "";

for (let i = 0; i < input.length; i++) {
  const command = input[i];
  let numCycles;
  let registerAdd;
  if (command === "noop") {
    numCycles = 1;
    registerAdd = 0;
  } else {
    numCycles = 2;
    registerAdd = parseInt(command.split(" ")[1]);
  }

  for (let c = 1; c <= numCycles; c++) {
    // During cycle
    if (isCriticalCycle(cycleNumber)) {
      cycleStrength += registerValue * cycleNumber;
    }

    const indexBeingDrawn = (cycleNumber - 1) % 40;
    crtImage += Math.abs(indexBeingDrawn - registerValue) <= 1 ? "#" : ".";

    // End of cycle
    cycleNumber++;
  }

  // Start of next cycle
  registerValue += registerAdd;
}

console.log("Answer to part 1 is:", cycleStrength);
// Answer to part 1 is: 14040

console.log("CRT Image:");
for (let i = 0; i < 6; i++) {
  console.log(crtImage.substring(40 * i, 40 * (i + 1) - 1));
}
// Letters in CRT Image: ZGCJZJFL
