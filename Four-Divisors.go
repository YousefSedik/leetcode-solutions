1func sumFourDivisors(nums []int) int {
2    answer := 0 
3    for _, num := range(nums){
4        if num  < 3{
5            continue
6        }
7        count := 2 // 1, the number itself
8        sum := 1 + num
9        for i := 2; i <= int(math.Sqrt(float64(num))); i++{
10            if num % i == 0 && i*i != num{
11                count += 2
12                sum += i + num / i    
13            } else if num % i == 0 && i*i == num{
14                count += 1
15            }
16        }
17
18        if count == 4{
19            answer += sum
20        }
21
22    }  
23    return answer
24
25}