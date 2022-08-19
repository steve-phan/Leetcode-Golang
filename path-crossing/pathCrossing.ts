// https://leetcode.com/problems/path-crossing/
function isPathCrossing(path: string): boolean {
  let x = 0;
  let y = 0;
  let locationsMap = {
    x0y0: true,
  } as Record<string, boolean>;
  for (let char of path) {
    switch (char) {
      case 'N': {
        x++;
        break;
      }
      case 'S': {
        x--;
        break;
      }
      case 'E': {
        y--;
        break;
      }
      case 'W': {
        y++;
        break;
      }
    }

    if (locationsMap[`x${x}y${y}`]) {
      return true;
    } else {
      locationsMap[`x${x}y${y}`] = true;
    }
  }

  return false;
}

console.log(isPathCrossing('NES'));
