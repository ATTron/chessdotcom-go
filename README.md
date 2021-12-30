# Chessdotcom-go
An unofficial, simple, lighweight API wrapper for chess.com written in go

## Usage
*Complete Documentation can be found at https://pkg.go.dev/github.com/ATTron/chessdotcom-go*  
```bash
    go get -u "github.com/ATTron/chessdotcom-go"
```
### If you just need the JSON string
```
    import chess "github.com/ATTron/chessdotcom-go"

    . . .

    const username = 'hikaru'

    func main() {
        userStats := chess.GetUserStats(username)
        fmt.Println(userStats)
    }
```
### If you want to use the JSON in a meaningful way in go I **highly** recommend using [gjson](https://github.com/tidwall/gjson)
```
    import (
        chess "github.com/ATTron/chessdotcom-go"
        "github.com/tidwall/gjson"
    )

    . . .

    const username = "hikaru"

    func main() {
        userStats := chess.GetUserStats(username)
        blitz := gjson.Get(userStats, "chess_blitz")
        fmt.Println(blitz.String())
        rapid := gjson.Get(userStats, "chess_rapid.record")
        fmt.Println(rapid.String())
    }
```
