package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

var data struct {
	MapID                 float64 `json:"mapId"`
	MatchCreation         float64 `json:"matchCreation"`
	MatchDuration         float64 `json:"matchDuration"`
	MatchID               float64 `json:"matchId"`
	MatchMode             string  `json:"matchMode"`
	MatchType             string  `json:"matchType"`
	MatchVersion          string  `json:"matchVersion"`
	ParticipantIdentities []struct {
		ParticipantID float64 `json:"participantId"`
		Player        struct {
			MatchHistoryURI string  `json:"matchHistoryUri"`
			ProfileIcon     float64 `json:"profileIcon"`
			SummonerID      float64 `json:"summonerId"`
			SummonerName    string  `json:"summonerName"`
		} `json:"player"`
	} `json:"participantIdentities"`
	Participants []struct {
		ChampionID                float64 `json:"championId"`
		HighestAchievedSeasonTier string  `json:"highestAchievedSeasonTier"`
		Masteries                 []struct {
			MasteryID float64 `json:"masteryId"`
			Rank      float64 `json:"rank"`
		} `json:"masteries"`
		ParticipantID float64 `json:"participantId"`
		Runes         []struct {
			Rank   float64 `json:"rank"`
			RuneID float64 `json:"runeId"`
		} `json:"runes"`
		Spell1Id float64 `json:"spell1Id"`
		Spell2Id float64 `json:"spell2Id"`
		Stats    struct {
			Assists                         float64 `json:"assists"`
			ChampLevel                      float64 `json:"champLevel"`
			CombatPlayerScore               float64 `json:"combatPlayerScore"`
			Deaths                          float64 `json:"deaths"`
			DoubleKills                     float64 `json:"doubleKills"`
			FirstBloodAssist                bool    `json:"firstBloodAssist"`
			FirstBloodKill                  bool    `json:"firstBloodKill"`
			FirstInhibitorAssist            bool    `json:"firstInhibitorAssist"`
			FirstInhibitorKill              bool    `json:"firstInhibitorKill"`
			FirstTowerAssist                bool    `json:"firstTowerAssist"`
			FirstTowerKill                  bool    `json:"firstTowerKill"`
			GoldEarned                      float64 `json:"goldEarned"`
			GoldSpent                       float64 `json:"goldSpent"`
			InhibitorKills                  float64 `json:"inhibitorKills"`
			Item0                           float64 `json:"item0"`
			Item1                           float64 `json:"item1"`
			Item2                           float64 `json:"item2"`
			Item3                           float64 `json:"item3"`
			Item4                           float64 `json:"item4"`
			Item5                           float64 `json:"item5"`
			Item6                           float64 `json:"item6"`
			KillingSprees                   float64 `json:"killingSprees"`
			Kills                           float64 `json:"kills"`
			LargestCriticalStrike           float64 `json:"largestCriticalStrike"`
			LargestKillingSpree             float64 `json:"largestKillingSpree"`
			LargestMultiKill                float64 `json:"largestMultiKill"`
			MagicDamageDealt                float64 `json:"magicDamageDealt"`
			MagicDamageDealtToChampions     float64 `json:"magicDamageDealtToChampions"`
			MagicDamageTaken                float64 `json:"magicDamageTaken"`
			MinionsKilled                   float64 `json:"minionsKilled"`
			NeutralMinionsKilled            float64 `json:"neutralMinionsKilled"`
			NeutralMinionsKilledEnemyJungle float64 `json:"neutralMinionsKilledEnemyJungle"`
			NeutralMinionsKilledTeamJungle  float64 `json:"neutralMinionsKilledTeamJungle"`
			ObjectivePlayerScore            float64 `json:"objectivePlayerScore"`
			PentaKills                      float64 `json:"pentaKills"`
			PhysicalDamageDealt             float64 `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions  float64 `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken             float64 `json:"physicalDamageTaken"`
			QuadraKills                     float64 `json:"quadraKills"`
			SightWardsBoughtInGame          float64 `json:"sightWardsBoughtInGame"`
			TotalDamageDealt                float64 `json:"totalDamageDealt"`
			TotalDamageDealtToChampions     float64 `json:"totalDamageDealtToChampions"`
			TotalDamageTaken                float64 `json:"totalDamageTaken"`
			TotalHeal                       float64 `json:"totalHeal"`
			TotalPlayerScore                float64 `json:"totalPlayerScore"`
			TotalScoreRank                  float64 `json:"totalScoreRank"`
			TotalTimeCrowdControlDealt      float64 `json:"totalTimeCrowdControlDealt"`
			TotalUnitsHealed                float64 `json:"totalUnitsHealed"`
			TowerKills                      float64 `json:"towerKills"`
			TripleKills                     float64 `json:"tripleKills"`
			TrueDamageDealt                 float64 `json:"trueDamageDealt"`
			TrueDamageDealtToChampions      float64 `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                 float64 `json:"trueDamageTaken"`
			UnrealKills                     float64 `json:"unrealKills"`
			VisionWardsBoughtInGame         float64 `json:"visionWardsBoughtInGame"`
			WardsKilled                     float64 `json:"wardsKilled"`
			WardsPlaced                     float64 `json:"wardsPlaced"`
			Winner                          bool    `json:"winner"`
		} `json:"stats"`
		TeamID   float64 `json:"teamId"`
		Timeline struct {
			CreepsPerMinDeltas struct {
				TenToTwenty    float64 `json:"tenToTwenty"`
				ThirtyToEnd    float64 `json:"thirtyToEnd"`
				TwentyToThirty float64 `json:"twentyToThirty"`
				ZeroToTen      float64 `json:"zeroToTen"`
			} `json:"creepsPerMinDeltas"`
			CsDiffPerMinDeltas struct {
				TenToTwenty    float64 `json:"tenToTwenty"`
				ThirtyToEnd    float64 `json:"thirtyToEnd"`
				TwentyToThirty float64 `json:"twentyToThirty"`
				ZeroToTen      float64 `json:"zeroToTen"`
			} `json:"csDiffPerMinDeltas"`
			DamageTakenDiffPerMinDeltas struct {
				TenToTwenty    float64 `json:"tenToTwenty"`
				ThirtyToEnd    float64 `json:"thirtyToEnd"`
				TwentyToThirty float64 `json:"twentyToThirty"`
				ZeroToTen      float64 `json:"zeroToTen"`
			} `json:"damageTakenDiffPerMinDeltas"`
			DamageTakenPerMinDeltas struct {
				TenToTwenty    float64 `json:"tenToTwenty"`
				ThirtyToEnd    float64 `json:"thirtyToEnd"`
				TwentyToThirty float64 `json:"twentyToThirty"`
				ZeroToTen      float64 `json:"zeroToTen"`
			} `json:"damageTakenPerMinDeltas"`
			GoldPerMinDeltas struct {
				TenToTwenty    float64 `json:"tenToTwenty"`
				ThirtyToEnd    float64 `json:"thirtyToEnd"`
				TwentyToThirty float64 `json:"twentyToThirty"`
				ZeroToTen      float64 `json:"zeroToTen"`
			} `json:"goldPerMinDeltas"`
			Lane               string `json:"lane"`
			Role               string `json:"role"`
			XpDiffPerMinDeltas struct {
				TenToTwenty    float64 `json:"tenToTwenty"`
				ThirtyToEnd    float64 `json:"thirtyToEnd"`
				TwentyToThirty float64 `json:"twentyToThirty"`
				ZeroToTen      float64 `json:"zeroToTen"`
			} `json:"xpDiffPerMinDeltas"`
			XpPerMinDeltas struct {
				TenToTwenty    float64 `json:"tenToTwenty"`
				ThirtyToEnd    float64 `json:"thirtyToEnd"`
				TwentyToThirty float64 `json:"twentyToThirty"`
				ZeroToTen      float64 `json:"zeroToTen"`
			} `json:"xpPerMinDeltas"`
		} `json:"timeline"`
	} `json:"participants"`
	PlatformID string `json:"platformId"`
	QueueType  string `json:"queueType"`
	Region     string `json:"region"`
	Season     string `json:"season"`
	Teams      []struct {
		Bans []struct {
			ChampionID float64 `json:"championId"`
			PickTurn   float64 `json:"pickTurn"`
		} `json:"bans"`
		BaronKills           float64 `json:"baronKills"`
		DominionVictoryScore float64 `json:"dominionVictoryScore"`
		DragonKills          float64 `json:"dragonKills"`
		FirstBaron           bool    `json:"firstBaron"`
		FirstBlood           bool    `json:"firstBlood"`
		FirstDragon          bool    `json:"firstDragon"`
		FirstInhibitor       bool    `json:"firstInhibitor"`
		FirstTower           bool    `json:"firstTower"`
		InhibitorKills       float64 `json:"inhibitorKills"`
		TeamID               float64 `json:"teamId"`
		TowerKills           float64 `json:"towerKills"`
		VilemawKills         float64 `json:"vilemawKills"`
		Winner               bool    `json:"winner"`
	} `json:"teams"`
	Timeline struct {
		Framefloat64erval float64 `json:"framefloat64erval"`
		Frames            []struct {
			ParticipantFrames struct {
				One struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"1"`
				One0 struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"10"`
				Two struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"2"`
				Three struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"3"`
				Four struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"4"`
				Five struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"5"`
				Six struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"6"`
				Seven struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"7"`
				Eight struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"8"`
				Nine struct {
					CurrentGold         float64 `json:"currentGold"`
					DominionScore       float64 `json:"dominionScore"`
					JungleMinionsKilled float64 `json:"jungleMinionsKilled"`
					Level               float64 `json:"level"`
					MinionsKilled       float64 `json:"minionsKilled"`
					ParticipantID       float64 `json:"participantId"`
					Position            struct {
						X float64 `json:"x"`
						Y float64 `json:"y"`
					} `json:"position"`
					TeamScore float64 `json:"teamScore"`
					TotalGold float64 `json:"totalGold"`
					Xp        float64 `json:"xp"`
				} `json:"9"`
			} `json:"participantFrames"`
			Timestamp float64 `json:"timestamp"`
		} `json:"frames"`
	} `json:"timeline"`
}

type table struct {
	ID     int64
	TeamID int64
	Gold   []float64
}

func main() {
	r, e := http.Get("https://na.api.pvp.net/api/lol/na/v2.2/match/2003855113?includeTimeline=true&api_key=66857942-d434-45f5-8ddd-aaeafdc92560")
	if e != nil {
		fmt.Println("Error:", e)
	}
	defer r.Body.Close()
	/*
		JSON, err := json.MarshalIndent(r, "", " ")
		if err != nil {
			fmt.Prfloat64ln("Error:", err)
		}

		fmt.Prfloat64(string(JSON))
	*/

	dec := json.NewDecoder(r.Body)
	dec.Decode(&data)

	var tableData []table

	for _, p := range data.Participants {
		fmt.Println(p.ParticipantID, p.TeamID, p.Timeline.GoldPerMinDeltas)

		var temp table
		temp.ID = int64(p.ParticipantID)
		temp.TeamID = int64(p.TeamID)
		temp.Gold = append(temp.Gold, (float64(int64(p.Timeline.GoldPerMinDeltas.ZeroToTen)*100) / 100))
		temp.Gold = append(temp.Gold, (float64(int64(p.Timeline.GoldPerMinDeltas.TenToTwenty)*100) / 100))
		temp.Gold = append(temp.Gold, (float64(int64(p.Timeline.GoldPerMinDeltas.TwentyToThirty)*100) / 100))
		temp.Gold = append(temp.Gold, (float64(int64(p.Timeline.GoldPerMinDeltas.ThirtyToEnd)*100) / 100))

		tableData = append(tableData, temp)
	}

	ren := render.New()
	mux := http.NewServeMux()

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		ren.HTML(w, http.StatusOK, "example", tableData)
	})

	http.ListenAndServe("127.0.0.1:8080", mux)
}
