function finalPrices(prices: number[]): number[] {
  for (let i = 0; i < prices.length; i++) {
    for (let j = 0; j < prices.length; j++) {
      if (prices[j] <= prices[i]) {
        prices[i] -= prices[j];
        break;
      }
    }
  }
  return prices;
}
