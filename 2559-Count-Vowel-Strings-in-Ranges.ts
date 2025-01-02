function prefix(words: string[]) {
    const prefixSum: number[] = [];
    const vowel = new Set(["a", "e", "i", "o", "u"]);
    prefixSum[0] =
        vowel.has(words[0][0]) && vowel.has(words[0][words[0].length - 1])
            ? 1
            : 0;

    for (let i = 1; i < words.length; i++) {
        prefixSum[i] =
            prefixSum[i - 1] +
            (vowel.has(words[i][0]) && vowel.has(words[i][words[i].length - 1])
                ? 1
                : 0);
    }

    return prefixSum;
}
function vowelStrings(words: string[], queries: number[][]): number[] {
    const result: number[] = [];
    const prefixSum = prefix(words);
    for (let i = 0; i < queries.length; i++) {
        const [l, r] = queries[i];
        if (l === 0) {
            result.push(prefixSum[r]);
        } else {
            result.push(prefixSum[r] - prefixSum[l - 1]);
        }
    }

    return result;
}
