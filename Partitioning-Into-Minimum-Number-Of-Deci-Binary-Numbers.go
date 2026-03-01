func minPartitions(n string) int {
    c := n[0]
    for _, i := range(n){
        if c < byte(i) {
            c = byte(i)
        }
    }
    return int(c) - 48
}