// naive solution
function calculate0And1(left: string, right: string): number {
    let count0 = 0;
    let count1 = 0;
    for (const char of left) {
        if (char === "0") count0++;
    }
    for (const char of right) {
        if (char === "1") count1++;
    }

    return count0 + count1;
}
function maxScore(s: string): number {
    const result: number[] = [];
    for (let i = 0; i < s.length - 1; i++) {
        // console.log(s.substring(0, i + 1), "right: ", s.substring(i + 1));
        result.push(calculate0And1(s.substring(0, i + 1), s.substring(i + 1)));
    }

    return Math.max(...result);
}
