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


const data = report()
  .split("\n")
