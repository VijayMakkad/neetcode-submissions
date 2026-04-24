func maxProfit(prices []int) int {
    profit,minPrice:=0,prices[0]

    for i:=1;i<len(prices);i++{
        minPrice=min(minPrice,prices[i])
        profit=max(profit,prices[i]-minPrice)
    }

    return profit

}
