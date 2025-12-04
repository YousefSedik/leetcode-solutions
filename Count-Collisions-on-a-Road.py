1class Solution:
2    def countCollisions(self, directions: str) -> int:
3        stack = []
4        collisions_counter = 0
5        for direction in directions:
6            if not stack:  # stack is empty
7                stack.append(direction)
8
9            elif direction == "S":
10                if stack[-1] == "R":
11                    while stack and stack[-1] == "R":
12                        collisions_counter += 1
13                        stack.pop()
14                stack.append("S")
15
16            elif direction == "L":
17                if stack[-1] == "S":
18                    collisions_counter += 1
19
20                elif stack[-1] == "R":
21                    collisions_counter += 2
22                    stack.pop()
23                    while stack and stack[-1] == "R":
24                        collisions_counter += 1
25                        stack.pop()
26
27                    stack.append("S")
28                else:
29                    stack.append("L")
30            elif direction == "R":
31                stack.append(direction)
32
33        return collisions_counter
34