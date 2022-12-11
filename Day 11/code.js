class Monkey {
  items;
  inspectOperation;
  divBy;
  monkeyTrue;
  monkeyFalse;
  itemsInspected;

  constructor(items, inspectOperation, divBy, monkeyTrue, monkeyFalse) {
    this.inspectOperation = inspectOperation;
    this.items = items;
    this.divBy = divBy;
    this.monkeyTrue = monkeyTrue;
    this.monkeyFalse = monkeyFalse;
    this.itemsInspected = 0;
  }
}

const superModulo = 11 * 7 * 13 * 5 * 3 * 17 * 2 * 19;

function monkeyBusiness(numRounds, shouldWorry) {
  const monkeys = [];
  // Monkey 0
  monkeys.push(new Monkey([57], (old) => old * 13, 11, 3, 2));
  // Monkey 1
  monkeys.push(
    new Monkey([58, 93, 88, 81, 72, 73, 65], (old) => old + 2, 7, 6, 7)
  );
  // Monkey 2
  monkeys.push(new Monkey([65, 95], (old) => old + 6, 13, 3, 5));
  // Monkey 3
  monkeys.push(new Monkey([58, 80, 81, 83], (old) => old * old, 5, 4, 5));
  // Monkey 4
  monkeys.push(new Monkey([58, 89, 90, 96, 55], (old) => old + 3, 3, 1, 7));
  // Monkey 5
  monkeys.push(
    new Monkey([66, 73, 87, 58, 62, 67], (old) => old * 7, 17, 4, 1)
  );
  // Monkey 6
  monkeys.push(new Monkey([85, 55, 89], (old) => old + 4, 2, 2, 0));
  // Monkey 7
  monkeys.push(
    new Monkey([73, 80, 54, 94, 90, 52, 69, 58], (old) => old + 7, 19, 6, 0)
  );

  for (let roundNum = 1; roundNum <= numRounds; roundNum++) {
    for (const monkey of monkeys) {
      monkey.items.forEach((item) => {
        // first monkey inspects item
        let newItemValue = monkey.inspectOperation(item);
        monkey.itemsInspected += 1;
        // monkey gets bored and doesn't break it
        newItemValue = shouldWorry
          ? newItemValue % superModulo
          : Math.floor(newItemValue / 3);
        // throws it to next monkey
        const newMonkey =
          newItemValue % monkey.divBy === 0
            ? monkey.monkeyTrue
            : monkey.monkeyFalse;

        monkeys[newMonkey].items.push(newItemValue);
      });
      // Monkey has thrown away all items
      monkey.items = [];
    }
  }
  return monkeys;
}

const monkeys1 = monkeyBusiness(20, false);
for (let monkeyNum = 0; monkeyNum < monkeys1.length; monkeyNum++) {
  const monkey = monkeys1[monkeyNum];
  console.log(`Monkey ${monkeyNum}: ${monkey.itemsInspected}`);
}
// Answer to part 1 : 350 * 347 = 121450
console.log("-----");
const monkeys2 = monkeyBusiness(10000, true);
for (let monkeyNum = 0; monkeyNum < monkeys2.length; monkeyNum++) {
  const monkey = monkeys2[monkeyNum];
  console.log(`Monkey ${monkeyNum}: ${monkey.itemsInspected}`);
}
// Answer to part 2 : 169205 * 165766 = 28244037010
