func removeSubfolders(folder []string) []string {
    sort.Strings(folder)
    
    res := []string{}
    
    for _, f := range folder {
        if len(res) == 0 || !strings.HasPrefix(f, res[len(res) - 1] + "/") {
            res = append(res, f)
        }
    }
    
    return res
}