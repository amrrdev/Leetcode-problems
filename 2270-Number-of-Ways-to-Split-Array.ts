function waysToSplitArray(nums: number[]): number {
    const totalSum = nums.reduce((acc, cur) => acc + cur, 0);
    let runningSum = 0;
    let validSplit = 0;
    for (let i = 0; i < nums.length - 1; i++) {
        runningSum += nums[i];
        if (runningSum >= totalSum - runningSum) {
            validSplit++;
        }
    }
    return validSplit;
}
