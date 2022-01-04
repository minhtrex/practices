package main

func addSpaces(s string, spaces []int) string {
    sb := strings.Builder{}
    spaceIdx := 0
    lsp := len(spaces)
    
    for i, r := range s {
        if spaceIdx < lsp && i == spaces[spaceIdx] {
            sb.WriteRune(' ')
            spaceIdx++
        }
        sb.WriteRune(r)
    }
    
    return sb.String()
}
