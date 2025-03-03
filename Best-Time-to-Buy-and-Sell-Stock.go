func maxProfit(prices []int) int {
\tbuy, seal := prices[0], 0

\tfor i := 1; i < len(prices); i++ {
\t\tif prices[i] < buy {
\t\t\tbuy = prices[i]
\t\t} else {
\t\t\tseal = max(seal, prices[i]-buy)
\t\t}
\t}

\treturn seal
}