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
      state.waypoint.coords.y += value;
      return state;
    }

    case SOUTH: {
      state.waypoint.coords.y -= value;
      return state;
    }

    case EAST: {
      state.waypoint.coords.x += value;
      return state;
    }

    case WEST: {
      state.waypoint.coords.x -= value;
      return state;
    }

    case LEFT: {
      const turns = value / 90;
      for (let i = 0; i < turns; i++) {
        [state.waypoint.coords.x, state.waypoint.coords.y] = [-state.waypoint.coords.y, state.waypoint.coords.x];
      }
      return state;
    }

    case RIGHT: {
      const turns = value / 90;
      for (let i = 0; i < turns; i++) {
        [state.waypoint.coords.x, state.waypoint.coords.y] = [state.waypoint.coords.y, -state.waypoint.coords.x];
      }
      return state;
    }

    case FORWARD: {
      state.ship.coords.x += Number(
        `${state.waypoint.coords.x < 0 ? "-" : ""}${Math.abs(state.waypoint.coords.x) * value}`
      );
      state.ship.coords.y += Number(
        `${state.waypoint.coords.y < 0 ? "-" : ""}${Math.abs(state.waypoint.coords.y) * value}`
      );
      return state;
    }

    default:
      throw new Error("invalid command:", command);
  }
};

const INITIAL_STATE = {
  ship: {
    coords: {
      x: 0,
      y: 0,
    },
  },
  waypoint: {
    coords: {
      x: 10,
      y: 1,
    },
  },
};

const run = (input, initialState) =>
  input
    .map((l) => [l[0], parseInt(l.slice(1), 10)])
    .reduce((state, [command, value]) => {
      const s = move(state, command, value);
      console.log(s);
      return s;
    }, initialState);

const input = fs
  .readFileSync("12.dat", "utf8")
  .trim()
  .split("\n")
  .map((v) => v.trim());

const endState = run(input, INITIAL_STATE);
console.log(`state: ${JSON.stringify(endState)}`);
console.log(`distance: ${Math.abs(endState.ship.coords.x) + Math.abs(endState.ship.coords.y)}`);
