package osuapi

import "time"

// Score - Hold information regarding scores on beatmaps
type Score struct {
	ScoreID         int       `json:"score_id,string"`
	Score           int       `json:"score,string"`
	Username        string    `json:"username,string"`
	Count300        int       `json:"count300,string"`
	Count100        int       `json:"count100,string"`
	Count50         int       `json:"count50,string"`
	CountMiss       int       `json:"countmiss,string"`
	MaxCombo        int       `json:"maxcombo,string"`
	CountKatu       int       `json:"countkatu,string"`
	CountGeki       int       `json:"countgeki,string"`
	Perfect         bool      `json:"perfect,string"`
	EnabledMods     int       `json:"enabled_mods,string"`
	UserID          string    `json:"user_id,string"`
	Date            time.Time `json:"date,string"`
	Rank            string    `json:"rank,string"`
	PP              float32   `json:"pp,string"`
	ReplayAvailable bool      `json:"replay_available,string"`
}
