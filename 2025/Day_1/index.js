const fs = require("fs");
const path = require("path");
const input = fs
  .readFileSync(path.join(__dirname, "rotations.txt"), "utf8")
  .replace(/\r/g, "")
  .split("\n");

let rotation = 50;
let count = 0;
for (const line of input) {
  const factor = parseInt(line.slice(1));
  if (line.startsWith("L")) {
    for (let i = 0; i < factor; i++) {
      rotation = (((rotation - 1) % 100) + 100) % 100;
      if (rotation == 0) count++;
    }
  } else if (line.startsWith("R")) {
    for (let i = 0; i < factor; i++) {
      rotation = (((rotation + 1) % 100) + 100) % 100;
      if (rotation == 0) count++;
    }
  }
}

console.log(count);
