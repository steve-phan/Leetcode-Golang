import java.util.Arrays;

public class CoinChange {

    public int coinChange(int[] coins, int amount) {
        // Create a DP array to store the minimum number of coins for each amount
        int[] dp = new int[amount + 1];
        
        // Initialize DP array with a value greater than the possible maximum coins needed (amount + 1)
        Arrays.fill(dp, amount + 1);
        
        // Base case: 0 coins are needed to make amount 0
        dp[0] = 0;
        
        // Fill the DP array
        for (int coin : coins) {
            for (int i = coin; i <= amount; i++) {
                dp[i] = Math.min(dp[i], dp[i - coin] + 1);
            }
        }
        
        // Check if the amount can be made up with the given coins
        return dp[amount] > amount ? -1 : dp[amount];
    }

    public static void main(String[] args) {
        CoinChange solution = new CoinChange();
        int[] coins = {1, 2, 5};
        int amount = 11;
        
        int result = solution.coinChange(coins, amount);
        System.out.println("Minimum coins needed: " + result);
    }
}
