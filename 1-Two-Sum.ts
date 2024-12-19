function twoSum(nums: number[], sum: number):number[] {
    const hashMap = new Map<number, number>();
    for (let i = 0; i < nums.length; i++) {
        const diff = sum - nums[i];
        if (hashMap.has(diff)) {
            return [hashMap.get(diff), i];
        }
        hashMap.set(nums[i], i);
    }
}