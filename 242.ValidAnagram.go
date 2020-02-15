func isAnagram(s string, t string) bool {
    if s == t {
        return true
    }
    ms := make(map[rune]int)
    mt := make(map[rune]int)
    for _, c := range s {
        ms[c]++
    }
    for _, c := range t {
        mt[c]++
    }
    if len(ms) != len(mt) {
        return false
    }
    for k, _ := range ms {
        if ms[k] != mt[k] {
            return false
        }
    }
    return true
}