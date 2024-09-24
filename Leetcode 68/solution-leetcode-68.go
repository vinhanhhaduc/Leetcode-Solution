func fullJustify(words []string, maxWidth int) []string {
    var result []string
    i := 0

    for i < len(words) {
        lineLength := len(words[i])
        last := i + 1
        for last < len(words) {
            if lineLength + 1 + len(words[last]) > maxWidth {
                break
            }
            lineLength += 1 + len(words[last])
            last++
        }
        var line string
        numOfWords := last - i
        if last == len(words) || numOfWords == 1 {
            line = strings.Join(words[i:last], " ")
            line += strings.Repeat(" ", maxWidth - len(line))
        } else {
            totalSpaces := maxWidth - lineLength + (numOfWords - 1)
            spaces := totalSpaces / (numOfWords - 1)
            extraSpaces := totalSpaces % (numOfWords - 1)

            for j := i; j < last-1; j++ {
                line += words[j]
                line += strings.Repeat(" ", spaces)
                if extraSpaces > 0 {
                    line += " "
                    extraSpaces--
                }
            }
            line += words[last-1]
        }
        
        result = append(result, line)
        i = last
    }
    
    return result
}