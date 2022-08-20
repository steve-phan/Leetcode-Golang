// https://leetcode.com/problems/xor-operation-in-an-array/

function xorOperation(n: number, start: number): number {
  let result = start;
  for (let i = 1; i < n; i++) {
    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_XOR
    result ^= start + 2 * i;
  }

  return result;
}
console.log(xorOperation(19, 8));

/**
Example 1:

Input: n = 5, start = 0
Output: 8
Explanation: Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8.
Where "^" corresponds to bitwise XOR operator.

Example 2:

Input: n = 4, start = 3
Output: 8
Explanation: Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.
 * 
 */
