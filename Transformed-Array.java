class Solution {
    public int[] constructTransformedArray(int[] nums) {
        int[] result = new int[nums.length];
        int index;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] > 0) {
                index = (i + nums[i]) % nums.length;
                result[i] = nums[index];
            } else if (nums[i] < 0) {
                index = (Math.abs(i+ nums.length-(Math.abs(nums[i])% nums.length)) % nums.length);
                result[i] = nums[index];
            } else {
                result[i] = nums[i];
            }
        }
        return result;
    }
}