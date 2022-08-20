function shuffle(nums: number[], n: number): number[] {
  const newNums: number[] = [];

  for (let i = 0; i < n; i++) {
    newNums[i * 2] = nums[i];
    newNums[i * 2 + 1] = nums[n + i];
  }
  return newNums;
}

console.log(shuffle([2, 5, 1, 3, 4, 7], 3));
