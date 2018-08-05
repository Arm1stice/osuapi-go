package osuapi

import "time"

// User - Contains information for osu! players
type User struct {
	Error        string       `json:"error"`
	UserID       string       `json:"user_id"`
	Username     string       `json:"username"`
	Count300     int          `json:"count300,string"`
	Count100     int          `json:"count100,string"`
	Count50      int          `json:"count50,string"`
	Playcount    int          `json:"playcount,string"`
	RankedScore  int          `json:"ranked_score,string"`
	TotalScore   int          `json:"total_score,string"`
	GlobalRank   int          `json:"pp_rank,string"`
	Level        float32      `json:"level,string"`
	PP           float32      `json:"pp_raw,string"`
	Accuracy     float32      `json:"accuracy,string"`
	CountRankSS  int          `json:"count_rank_ss,string"`
	CountRankSSH int          `json:"count_rank_ssh,string"`
	CountRankS   int          `json:"count_rank_s,string"`
	CountRankSH  int          `json:"count_rank_sh,string"`
	CountRankA   int          `json:"count_rank_a,string"`
	Country      string       `json:"country"`
	CountryRank  int          `json:"pp_country_rank,string"`
	Events       []UserEvents `json:"events"`
}

// UserEvents - A type for the "events" field in a user's profile
type UserEvents struct {
	DisplayHTML  string    `json:"display_html"`
	BeatmapID    int       `json:"beatmap_id,string"`
	BeatmapSetID int       `json:"beatmapset_id,string"`
	Date         time.Time `json:"date,string"`
	EpicFactor   int       `json:"epicfactor,string"`
}
