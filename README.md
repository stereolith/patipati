## 🍤 patipati

[work in progress]
patipati is a small command line tool that creates native linux applications from webpages.

This uses a modified version of [webview](https://github.com/zserge/webview) which is optimized for retention (cookies and website data)

### Usage

1. clone repistory
2. (for x86_64 linux) use compiled binary from `dist/` **or** compile program with `go build -o dist/patipati`
3. Create native app from webpage with:
    ```
    ./dist/patipati URL [Title]
    ```
i.e. `./dist/patipati https://notion.so`

![demo gif](https://raw.githubusercontent.com/stereolith/patipati/master/docs/demo.gif)

### ToDo:
* implement function to uninstall/ list webpage applications
* save website caches for faster loading