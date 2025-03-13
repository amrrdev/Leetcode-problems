func maxProfit(prices []int) int {
\tbuy := prices[0]
\tseal := 0
\tfor i := 1; i < len(prices); i++ {
\t\tif prices[i] < buy {
\t\t\tbuy = prices[i]
\t\t}
\t\tif prices[i] > buy {
\t\t\tseal = max(seal, prices[i]-buy)
\t\t}
\t}

\treturn seal
}