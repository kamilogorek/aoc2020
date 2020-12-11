const fs = require("fs");
require("lodash.multipermutations");
const _ = require("lodash");

const EMPTY = "L";
const OCCUPIED = "#";
const NONE = ".";

const isTaken = (v) => v === OCCUPIED;
const isInbound = (world, y, x) => y >= 0 && y < world.length && x >= 0 && x < world[0].length;
const areEqualGrids = (a, b) => a.flat().join("") === b.flat().join("");

const getOne = (world, y, x) => world[y]?.[x] || NONE;
const getInDirection = (world, y, x, yDir, xDir, nearby) => {
  let cell;
  do {
    y += yDir;
    x += xDir;
    cell = getOne(world, y, x);
    if (nearby) {
      return cell;
    }
  } while (cell === NONE && isInbound(world, y, x));
  return cell;
};
const getVisible = (world, y, x, nearby) =>
  _.multipermutations([-1, 0, 1], 2)
    .filter(([y, x]) => x !== 0 || y !== 0)
    .map(([yDir, xDir]) => getInDirection(world, y, x, yDir, xDir, nearby));

const evolve = (world, config) => {
  const evolvedWorld = world.map((row, y) =>
    row.map((seat, x) => {
      const seats = getVisible(world, y, x, config.nearby);
      const takenSeats = seats.filter(isTaken).length;
      if (seat === EMPTY && takenSeats === 0) {
        return OCCUPIED;
      }
      if (seat === OCCUPIED && takenSeats >= config.limit) {
        return EMPTY;
      }
      return seat;
    })
  );
  return areEqualGrids(world, evolvedWorld) ? evolvedWorld : evolve(evolvedWorld, config);
};

const input = fs
  .readFileSync("11.dat", "utf8")
  .trim()
  .split("\n")
  .map((v) => v.trim().split(""));

console.log(`Part 1: ${evolve(input, { limit: 4, nearby: true }).flat().filter(isTaken).length}`);
console.log(`Part 2: ${evolve(input, { limit: 5, nearby: false }).flat().filter(isTaken).length}`);
