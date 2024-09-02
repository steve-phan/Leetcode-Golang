function mergeSort(arr : number[]) {
    if (arr.length <= 1) {
      return arr;
    }
    
    const mid = Math.floor(arr.length / 2);
    const left = arr.slice(0, mid);
    const right = arr.slice(mid);
  
    return merge(mergeSort(left), mergeSort(right));
  }
  
  function merge(left : number[], right : number[]) {
    const result : number[] = [];
  
    while ( left.length && right.length) {
      if (left[0] < right[0]) {     
        result.push(left.shift()!);
        left.shift()
      } else {
        result.push(right.shift()!);
      }
    }
  
    return [...result, ...left, ...right];
  }
  
  console.log(mergeSort([3, 6, 8, 10, 1, 2, 1]));