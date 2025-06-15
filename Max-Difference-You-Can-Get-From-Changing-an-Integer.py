class Solution:
    def maxDiff(self, num: int) -> int:
        num: str = str(num)
        list_of_possible_nums = []
        for i in range(10):
            for j in range(10):
                option = num.replace(str(i), str(j))
                if len(str(int(option))) == len(num) and option != '0':
                    list_of_possible_nums.append(int(option))

        return max(list_of_possible_nums) - min(list_of_possible_nums)