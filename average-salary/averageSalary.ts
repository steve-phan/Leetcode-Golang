// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

function average(salary: number[]): number {
  let min = salary[0];
  let max = salary[0];
  let totalAmount = 0;

  for (let num of salary) {
    if (num < min) {
      min = num;
    }
    if (num > max) {
      max = num;
    }
    totalAmount += num;
  }
  // We have min, max, totalAmount

  const averageNum = (totalAmount - (min + max)) / (salary.length - 2);

  return Number(averageNum.toFixed(5));
}
