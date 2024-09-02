package sort;

import java.util.Arrays;

public class MergeSort {

    public static void main(String[] args) {

        int[] array = {3, 6, 8, 10, 1, 2, 1};
        array = mergeSort(array);
        System.out.println(Arrays.toString(array));

    }

    public static int[] mergeSort(int[] arr) {
        if (arr.length <= 1) {
            return arr;
        }

        int mid = arr.length / 2;
        int[] left = Arrays.copyOfRange(arr, 0, mid);
        int[] right = Arrays.copyOfRange(arr, mid, arr.length);

        return merge(mergeSort(left), mergeSort(right));
    }

    public static int[] merge(int[] left, int[] right) {
        // Merge the left and right arrays into the result array
         int[] result = new int[left.length + right.length];
         int i = 0, j = 0, k = 0;

        while(i < left.length && j < right.length) {
            if(left[i] < right[j]) {
                result[k++] = left[i++];
            } else {
                result[k++] = right[j++];
            }
        }

        // Copy the remaining elements from left array

        while(i < left.length) {
            result[k++] = left[i++];
        }

        // Copy the remaining elements from right array

        while (j < right.length) {
            result[k++] = right[j++];
        }
        
        return result;


    }

}
