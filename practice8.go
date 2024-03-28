package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "os"
)

// タイトルタグを探して内容を返す関数
func GetTitle(n *html.Node) (string, bool) {
    if n.Type == html.ElementNode && n.Data == "title" {
        return n.FirstChild.Data, true
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        if result, ok := GetTitle(c); ok {
            return result, true
        }
    }
    return "", false
}

// 指定したURLのHTMLを取得してタイトルを表示する関数
func ScrapeTitle(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "parse: %v\n", err)
        os.Exit(1)
    }

    if title, ok := GetTitle(doc); ok {
        fmt.Println(title)
    } else {
        fmt.Println("タイトルが見つかりませんでした。")
    }
}

func main() {
    url := "https://techjourney-code.com/golang-wireshark/"
    ScrapeTitle(url)
}
