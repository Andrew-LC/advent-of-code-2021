import fs from "fs"

function report(): string {
  let data: string;
  try {
    data = fs.readFileSync('./data.txt', 'utf8')
    return data
  } catch(err) {
    console.log(err)
    return "";
  }
}

interface SubmarineI {
  depth: number,
  horizontal: number
}

enum Commands {
  FORWARD = "FORWARD",
  DOWN = "DOWN",
  UP = "UP",
} 

const submarine: SubmarineI = {
  depth: 0,
  horizontal: 0
}

report()
  .split("\n")
  .map(v => {
    let temp = v.split(" ");
    return [temp[0].toUpperCase(), +temp[1]];
  })
  .forEach(v => {
    switch(v[0]){
      case Commands.FORWARD: {
	submarine.horizontal += v[1] as number;
	break;
      }
      case Commands.UP: {
	submarine.depth -= v[1] as number;
	break;
      }
      case Commands.DOWN: {
	submarine.depth += v[1] as number;
	break;
      }
    }
  });


console.log(submarine.depth*submarine.horizontal);
