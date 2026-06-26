var para = map[rune]rune{
    '(': ')',
    '{': '}',
    '[': ']',
}

func isValid(s string) bool {
    var list []rune;

    top := -1;
    if len(s) == 1 {
        return false
    }

    for _, i := range s {
        if len(list) == 0 {
            list = append(list, i);
            top += 1;
        } else {
            if (para[list[top]] == i){
                list = list[:len(list)-1]
                top -= 1;
            } else {
                list = append(list, i);
                top += 1
            }
        }
    }

    return len(list) == 0
}