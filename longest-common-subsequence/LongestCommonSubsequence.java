/*
 * Input: text1 = "abcde", text2 = "ace"
 * Output: 3
 * Explanation: The longest common subsequence is "ace" and its length is 3.

 */

public class LongestCommonSubsequence {

    public int longestCommonSubsequence(String text1, String text2) {

        int m = text1.length();
        int n = text2.length();

        // Create a DP table with dimensions (m+1) * (n+1)
        int[][] dp = new int[m + 1][n + 1];

        // Fill the DP table

        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (text1.charAt(i - 1) == text2.charAt(j - 1)) {
                    // Characters match, extend the LCS
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {

                    // Characters don't match, take the max LCS found so far
                    dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
                }

            }
        }

        // The LCS length is at the bottom-right correct of the DP table
        return dp[m][n];

    }

    public static void main(String[] args) {
        LongestCommonSubsequence solution = new LongestCommonSubsequence();

        String text1 = "abcde";
        String text2 = "acc";

        int result = solution.longestCommonSubsequence(text1, text2);
        System.out.println("Length of LCSL: " + result);
    }

}
