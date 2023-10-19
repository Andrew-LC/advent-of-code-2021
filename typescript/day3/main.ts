import { file } from "../util/main";

const report = file("./data.txt");
// const report = `00100
// 11110
// 10110
// 10111
// 10101
// 01111
// 00111
// 11100
// 10000
// 11001
// 00010
// 01010`;


function generateNewArray(arr: number[][]): number[][] {
  let newArr = [];
  let temp = [];
  let index = 0;
  while(index < arr[0].length){
    for(let i = 0; i < arr.length; i++){
      temp.push(arr[i][index])
    }
    newArr.push(temp);
    temp = [];
    index += 1;
  }

  return newArr;
}


function bitOperation(arr: number[][]) {
  const data = {gamma: [], eplison: []};

  const temp = arr.map((v, _) => {
    const count = {0: 0, 1: 0};
    v.forEach(value => {
      if(value === 0){
	count[0]++
      } else if(value === 1) {
	count[1]++;
      }
    })
    return count;
  }) 

  temp.forEach(v => {
    if(v[0] > v[1]){
      data.gamma.push(0);
      data.eplison.push(1);
    } else {
      data.gamma.push(1);
      data.eplison.push(0);
    }
  })

  return data;
}

const lines = report.split("\n")
  .map((v) => v.split("").map(Number));

const data = bitOperation(generateNewArray(lines))
const gamma = parseInt(data.gamma.join(""), 2);
const eplison = parseInt(data.eplison.join(""), 2);

console.log(`The power consumption is ${gamma*eplison}`);
