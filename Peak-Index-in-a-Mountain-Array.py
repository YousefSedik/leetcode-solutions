class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        l, r = 0, len(arr) - 1
        ans = 0
        while l <= r:
            mid = (l + r) // 2
            if arr[mid] < arr[mid + 1]:
                ans = mid
                l = mid + 1
            else:
                r = mid - 1

        return ans + 1
