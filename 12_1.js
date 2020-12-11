const fs = require("fs");
const _ = require("lodash");

const NORTH = "N";
const SOUTH = "S";
const EAST = "E";
const WEST = "W";
const LEFT = "L";
const RIGHT = "R";
const FORWARD = "F";
const LEFT_TURN = [NORTH, WEST, SOUTH, EAST];
const RIGHT_TURN = [NORTH, EAST, SOUTH, WEST];

const move = (state, command, value) => {
  switch (command) {
    case NORTH: {
      state.coords.y -= value;
      return state;
    }

    case SOUTH: {
      state.coords.y += value;
      return state;
    }

    case EAST: {
      state.coords.x += value;
      return state;
    }

    case WEST: {
      state.coords.x -= value;
      return state;
    }

    case LEFT: {
      state.dir = LEFT_TURN[(LEFT_TURN.findIndex((v) => v === state.dir) + value / 90) % LEFT_TURN.length];
      return state;
    }

    case RIGHT: {
      state.dir = RIGHT_TURN[(RIGHT_TURN.findIndex((v) => v === state.dir) + value / 90) % RIGHT_TURN.length];
      return state;
    }

    case FORWARD: {
      return moveForward(state, value);
    }

    default:
      throw new Error("invalid command");
  }
};

function moveForward(state, value) {
  switch (state.dir) {
    case NORTH: {
      state.coords.y -= value;
      break;
    }

    case SOUTH: {
      state.coords.y += value;
      break;
    }

    case EAST: {
      state.coords.x += value;
      break;
    }

    case WEST: {
      state.coords.x -= value;
      break;
    }

    default:
      throw new Error("invalid dir");
  }
  return state;
}

const INITIAL_STATE = {
  dir: EAST,
  coords: {
    x: 0,
    y: 0,
  },
};

const run = (input, initialState) =>
  input
    .map((l) => [l[0], parseInt(l.slice(1), 10)])
    .reduce((state, [command, value]) => move(state, command, value), initialState);

const input = fs
  .readFileSync("12.dat", "utf8")
  .trim()
  .split("\n")
  .map((v) => v.trim());

const endState = run(input, INITIAL_STATE);
console.log(`state: ${JSON.stringify(endState)}`);
console.log(`distance: ${endState.coords.x + endState.coords.y}`);
