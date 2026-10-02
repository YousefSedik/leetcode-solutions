func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var generator = func(current string){};
	generator = func(current string){
		if len(current) == n * 2{
			res = append(res, current)
			return 
		}
		if strings.Count(current, "(") < n {
			generator(current + "(")
		} 
		if strings.Count(current, ")") < strings.Count(current, "("){
			generator(current + ")")
		}
	}
	generator("")
	return res
}