class Solution:
    def numDecodings(self, s: str) -> int:
        cur_s = [l for l in s]
        mem = dict()
        def solve(cur_s: list[str]): # TODO: use i instead of using cur_s
            if len(cur_s) == 0:
                return 1
            # is it possible to choose 1 digit ?
            # yes, if cur_s[0] != 0
            # is it possible to choose 2 digit ?
            # yes, if cur_s[0] != 0 and cur_s[0] + cur[1] > 26
            if mem.get(''.join(cur_s)):
                return mem.get("".join(cur_s))
            can_choose_1_d = cur_s[0] != '0'
            can_choose_2_d = (
                cur_s[0] != '0' and len(cur_s) >= 2 and int(cur_s[0] + cur_s[1]) <= 26
            )
            opt1, opt2 = 0, 0
            cur_s_c = cur_s.copy()
            cur_s_c.pop(0)
            if can_choose_1_d:
                opt1 = solve(cur_s=cur_s_c)

            if can_choose_2_d:
                cur_s_c.pop(0)
                opt2 = solve(cur_s=cur_s_c)
            mem["".join(cur_s)] = opt1 + opt2
            return opt1 + opt2

        return solve(cur_s)