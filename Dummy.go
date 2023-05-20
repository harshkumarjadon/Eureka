func insertCommas(s string) string {

    n := len(s)

    if n <= 3 {

        return s

    }

    return insertCommas(s[:n-3]) + "," + s[n-3:]

}

func playWithInsertommas() {

    fmt.Println(insertCommas("12345678"))

    fmt.Println(insertCommas("999"))

    fmt.Println(insertCommas("6789999"))

    fmt.Println(insertCommas("89389438493849384398439483"))

    fmt.Println(insertCommas("123456788888888888"))

}