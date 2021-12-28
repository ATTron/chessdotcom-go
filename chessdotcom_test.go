package chessdotcomgo

import "testing"

const username = "hikaru"

func TestEndpoints(t *testing.T) {
	t.Run("given specific user", func(t *testing.T) {
		t.Run("checking all info", func(t *testing.T) {
			allInfo := GetAllData(username)
			if allInfo.BasicInfo != "" {
				t.Log("Found Basic Info!", allInfo.BasicInfo)
			} else {
				t.Error("Basic info for user", username, "not found")
			}

			if allInfo.Stats != "" {
				t.Log("Found Stats!", allInfo.Stats)
			} else {
				t.Error("Stats for user", username, "not found")
			}

			// if allInfo.OnlineStats != "" {
			// 	t.Log("Found Basic Info!", allInfo.OnlineStats)
			// } else {
			// 	t.Error("Online status for user", username, "not found")
			// }

		})

		t.Run("checking player games", func(t *testing.T) {
			playerGames := GetPlayerGames(username)
			t.Log("Got player games", playerGames)
			playerGames = GetPlayerGamesArchive(username, 2021, 11)
			t.Log("Got player archives", playerGames)
			playerGames = GetPlayerMatches(username)
			t.Log("Got player matches", playerGames)
			playerGames = GetPlayerTournaments(username)
			t.Log("Got player tournaments", playerGames)

			// 	playerGames := GetPlayerGamesArchivePGN(username, 2021, 11)
			// 	t.Log("Got player archives PGN", playerGames)
		})
	})
}
