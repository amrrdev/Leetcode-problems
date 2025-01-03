function _prefix(nums: number[]): number[] {
    const prefix: number[] = [];
    prefix[0] = nums[0];
    for (let i = 1; i < nums.length; i++) {
        prefix[i] = prefix[i - 1] + nums[i];
    }
    return prefix;
}

function waysToSplitArray(nums: number[]): number {
    const prefix = _prefix(nums);
    let numSplitArray = 0;
    for (let i = 0; i < nums.length - 1; i++) {
        if (prefix[i] >= prefix[nums.length - 1] - prefix[i]) numSplitArray++;
    }

    return numSplitArray;
}