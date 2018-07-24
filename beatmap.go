package osuapi

import "time"

// Beatmap - A type for beatmap objects in the osu! api
type Beatmap struct {
	Approved         int       `json:"approved,string"`
	ApprovedDate     time.Time `json:"approved_date,string"`
	LastUpdate       time.Time `json:"last_update,string"`
	Artist           string    `json:"artist"`
	BeatmapID        int       `json:"beatmap_id,string"`
	BeatmapSetID     int       `json:"beatmapset_id,string"`
	BPM              float32   `json:"bpm,string"`
	Creator          string    `json:"creator"`
	DifficultyRating float32   `json:"difficultyrating,string"`
	DiffSize         float32   `json:"diff_size,string"`
	DiffOverall      float32   `json:"diff_overall,string"`
	DiffApproach     float32   `json:"diff_approach,string"`
	DiffDrain        float32   `json:"diff_drain,string"`
	HitLength        float32   `json:"hit_length,string"`
	Source           string    `json:"source"`
	GenreID          int       `json:"genre_id,string"`
	LanguageID       int       `json:"language_id,string"`
	Title            string    `json:"title"`
	TotalLength      float32   `json:"total_length,string"`
	MD5              string    `json:"version"`
	Mode             int       `json:"file_md5,string"`
	Tags             string    `json:"mode"`
	FavoriteCount    int       `json:"tags,string"`
	PlayCount        int       `json:"playcount,string"`
	PassCount        int       `json:"passcount,string"`
	MaxCombo         int       `json:"max_combo,string"`
}
