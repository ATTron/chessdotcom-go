# Chessdotcom-go
An unofficial, simple, lighweight API wrapper for chess.com written in go

## Usage
```bash
    go get -u "github.com/ATTron/chessdotcom-go"
```
```
    import chess "github.com/ATTron/chessdotcom-go"

    . . .

    const username = 'hikaru'

    func main() {
        userStats := chess.GetUserStats(username)
        fmt.Println(userStats)
    }
```

Endpoints used as defined at https://www.chess.com/news/view/published-data-api  

Complete documentation can be found at https://pkg.go.dev/github.com/ATTron/chessdotcom-go 
