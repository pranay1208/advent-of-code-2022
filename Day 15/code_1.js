const fs = require("fs");

const input = fs
  .readFileSync("./input.txt", "utf-8")
  .replaceAll("\r", "")
  .trim()
  .split("\n");

const regex =
  /Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)/;

const Location = {
  Beacon: "Beacon",
  Sensor: "Sensor",
  Empty: "Empty",
};

const locationKey = (x, y) => `${x}|${y}`;

const yToFind = 2000000;
const locationMap = {};
let answer1 = 0;

for (const line of input) {
  const regexResult = regex.exec(line);
  const sensorX = parseInt(regexResult[1]);
  const sensorY = parseInt(regexResult[2]);
  const beaconX = parseInt(regexResult[3]);
  const beaconY = parseInt(regexResult[4]);

  locationMap[locationKey(beaconX, beaconY)] = Location.Beacon;
  locationMap[locationKey(sensorX, sensorY)] = Location.Sensor;

  const manhattanDistance =
    Math.abs(beaconX - sensorX) + Math.abs(beaconY - sensorY);

  const distanceToY = Math.abs(yToFind - sensorY);

  if (manhattanDistance < distanceToY) {
    continue;
  }

  const xOptions = Math.abs(manhattanDistance - distanceToY);

  for (let i = -1 * xOptions; i <= xOptions; i++) {
    if (locationMap[locationKey(sensorX + i, yToFind)] === undefined) {
      locationMap[locationKey(sensorX + i, yToFind)] = Location.Empty;
      answer1 += 1;
    }
  }
}

console.log("Answer to part 1 is", answer1);
// Answer to part 1 is 4748135
