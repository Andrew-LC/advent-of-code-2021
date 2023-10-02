import fs from "fs";

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

const data = report()
  .trim()
  .split("\n")
  .filter((x, i, array) => {
    if(i == 0){
      return false;
    } else if(+x > +array[i-1]){
      return x
    }
  }).length;

console.log(data);
