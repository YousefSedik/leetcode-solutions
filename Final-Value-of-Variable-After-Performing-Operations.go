func finalValueAfterOperations(operations []string) int {
    x := 0
    for _, operation := range(operations){
        if operation == "X++" || operation == "++X"{
            x++
        } else {
            x--
        }
    }
    return x
}