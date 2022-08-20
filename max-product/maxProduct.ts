//https://leetcode.com/problems/maximum-product-of-word-lengths/

// function maxProduct(words: string[]): number {
//   // Sort array first
//   //   const newWords = words.sort((a, b) => b.length - a.length);
//   const n = words.length;
//   let max = 0;
//   for (let i = 0; i < n - 1; i++) {
//     for (let j = i + 1; j < n; j++) {
//       if (isDiff(words[i], words[j])) {
//         max = Math.max(max, words[i].length * words[j].length);
//       }
//     }
//   }

//   return max;
// }

// function isDiff(a: string, b: string): boolean {
//   // this is the best, because we loop to the unique char
//   // try with the case 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa....'
//   for (let i = 97; i < 123; i++) {
//     if (
//       a.includes(String.fromCharCode(i)) &&
//       b.includes(String.fromCharCode(i))
//     ) {
//       return false;
//     }
//   }
//   return true;
// }
function maxProduct(words: string[]): number {
  let arr: number[] = [];
  arr = words.map((str) => {
    let cur = 0;
    for (let i = 0; i < str.length; i++) {
      cur |= 1 << (str.charCodeAt(i) - 97);
    }
    return cur;
  });
  let res = 0;
  for (let i = 0; i < arr.length; i++) {
    for (let j = i + 1; j < arr.length; j++) {
      if ((arr[i] & arr[j]) === 0) {
        // meaning unique
        res = Math.max(res, words[i].length * words[j].length);
      }
    }
  }
  return res;
}
console.log(maxProduct(['abcw', 'baz', 'foo', 'bar', 'xtfn', 'abcdef']));

// char a-z => 97-122 charAtCode

/**
 * Example 1:

Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
Output: 16
Explanation: The two words can be "abcw", "xtfn".
Example 2:

Input: words = ["a","ab","abc","d","cd","bcd","abcd"]
Output: 4
Explanation: The two words can be "ab", "cd".
Example 3:

Input: words = ["a","aa","aaa","aaaa"]
Output: 0
Explanation: No such pair of words.
 * 
 * 
 */
// const convertToInt = (str) => {
//   let int = 0;
//   const baseCharCode = 'a'.charCodeAt(0);

//   for (let i = 0; i < str.length; i++) {
//     int |= 1 << (str.charCodeAt(i) - baseCharCode);
//   }

//   return int;
// };

// const areIntsUnique = (a, b) => {
//   if ((a & b) === 0) return true;
//   else return false;
// };

// const maxProduct = function (words) {
//   const ints = words.map((word) => convertToInt(word));

//   let max = 0;
//   for (let i = 0; i < ints.length - 1; i++) {
//     for (let j = i + 1; j < ints.length; j++) {
//       if (areIntsUnique(ints[i], ints[j])) {
//         max = Math.max(max, words[i].length * words[j].length);
//       }
//     }
//   }

//   return max;
// };
