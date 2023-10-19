import fs from "fs"

export function file(file: string): string {
  let data: string;
  try {
    data = fs.readFileSync(file, 'utf8')
    return data
  } catch(err) {
    console.log(err)
    return "";
  }
}
