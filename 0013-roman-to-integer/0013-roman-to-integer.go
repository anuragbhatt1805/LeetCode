var symbols = map[rune]int{
    ' ':0,
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func val(i1 rune, i2 rune) int {
    if (symbols[i1] < symbols[i2]) {
        return symbols[i2]-(symbols[i1]*2)
    } else {
        return symbols[i2]
    }
}

func romanToInt(s string) int {
    var res int = 0
    last := ' '

    for _, i := range s {
        res += val(last, i)
        last = i
    }

    return res;
}