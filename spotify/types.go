package spotify

import "encoding/json"

type Album struct {
	AlbumType            string        `json:"album_type"`
	Artists              []*Artist     `json:"artists"`
	AvailableMarkest     []string      `json:"available_markets"`
	Copyrights           []*Copyright  `json:"copyrights"`
	ExternalIDs          *ExternalIDs  `json:"external_ids"`
	ExternalURLs         *ExternalURLs `json:"external_urls"`
	Genres               []string      `json:"genres"`
	Href                 string        `json:"href"`
	ID                   string        `json:"id"`
	Images               []*Image      `json:"images"`
	Label                string        `json:"label"`
	Name                 string        `json:"name"`
	Popularity           int           `json:"popularity"`
	ReleaseDate          string        `json:"release_date"`
	ReleaseDatePrecision string        `json:"release_date_precision"`
	Tracks               *Paging       `json:"tracks"`
	Type                 string        `json:"type"`
	URI                  string        `json:"uri"`
}

type Artist struct {
	ExternalURLs *ExternalURLs `json:"external_url"`
	Followers    *Follower     `json:"followers"`
	Genres       []string      `json:"genres"`
	Href         string        `json:"href"`
	ID           string        `json:"id"`
	Images       []*Image      `json:"images"`
	Name         string        `json:"name"`
	Popularity   int           `json:"popularity"`
	Type         string        `json:"type"`
	URI          string        `json:"uri"`
}

type AudioFeatures struct {
	Acousticness     float32 `json:"acousticness"`
	AnalysisURL      string  `json:"analysis_url"`
	Danceabilitu     float32 `json:"danceability"`
	DurationMs       int     `json:"duration_ms"`
	Energy           float32 `json:"energy"`
	ID               string  `json:"id"`
	Instrumentalness float32 `json:"instrumentalness"`
	Key              int     `json:"key"`
	Liveness         float32 `json:"liveness"`
	Loudness         float32 `json:"loudness"`
	Mode             int     `json:"mode"`
	Speechiness      float32 `json:"speechiness"`
	Tempo            float32 `json:"tempo"`
	TimeSignature    int     `json:"time_signature"`
	TrackHref        string  `json:"track_href"`
	Type             string  `json:"type"`
	URI              string  `json:"uri"`
	Valence          float32 `json:"valence"`
}

type AudioAnalysis struct {
	Bars []*struct {
		Start      float32 `json:"start"`
		Duration   float32 `json:"duration"`
		Confidence float32 `json:"confidence"`
	} `json:"bars"`

	Beats []*struct {
		Start      float32 `json:"start"`
		Duration   float32 `json:"duration"`
		Confidence float32 `json:"confidence"`
	} `json:"beats"`

	Meta *struct {
		AnalyzerVersion string  `json:"analyzer_version"`
		Platform        string  `json:"platform"`
		DetailedStatus  string  `json:"detailed_status"`
		StatusCode      int     `json:"status_code"`
		Timestamp       int     `json:"timestamp"`
		AnalysisTime    float32 `json:"analysis_time"`
		InputProcess    string  `json:"input_process"`
	} `json:"meta"`

	Sections []*struct {
		Start                   float32 `json:"start"`
		Duration                float32 `json:"duration"`
		Confidence              float32 `json:"float32"`
		Loudness                float32 `json:"loudness"`
		Tempo                   float32 `json:"tempo"`
		TempoConfidence         float32 `json:"tempo_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float32 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float32 `json:"mode_confidence"`
		TimeSignature           int     `json:"time_signature"`
		TimeSignatureConfidence float32 `json:"time_signature_confidence"`
	} `json:"sections"`

	Segments []*struct {
		Start           float32   `json:"start"`
		Duration        float32   `json:"duration"`
		Confidence      float32   `json:"confidence"`
		LoudnessStart   float32   `json:"loudenss_start"`
		LoudnessMaxTime float32   `json:"loudness_max_time"`
		LoudnessMax     float32   `json:"loudness_max"`
		LoudnessEnd     float32   `json:"loudness_end"`
		Pitches         []float32 `json:"pitches"`
		Timbre          []float32 `json:"timbre"`
	} `json:"segments"`

	Tatums []*struct {
		Start      float32 `json:"start"`
		Duration   float32 `json:"duration"`
		Confidence float32 `json:"confidence"`
	} `json:"tatums"`

	Track *Track `json:"track"`
}

type Category struct {
	Href  string   `json:"href"`
	Icons []*Image `json:"images"`
	ID    string   `json:"id"`
	Name  string   `json:"name"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Cursor struct {
	After string `json:"after"`
}

type SpotifyError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ExternalIDs struct {
	IDs map[string]string `json:"-"`
}

type ExternalURLs struct {
	URLs map[string]string `json:"-"`
}

type Follower struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

type Paging struct {
	Href     string          `json:"href"`
	Items    json.RawMessage `json:"items"`
	Limit    int             `json:"limit"`
	Next     string          `json:"next"`
	Offset   int             `json:"offset"`
	Previous string          `json:"previous"`
	Total    int             `json:"total"`
}

type CursorBasedPaging struct {
	Href   string        `json:"href"`
	Items  []interface{} `json:"items"`
	Limit  int           `json:"limit"`
	Next   string        `json:"next"`
	Cursor *Cursor       `json:"cursor"`
	Total  int           `json:"total"`
}

type Playlist struct {
	Collaborative bool          `json:"collaborative"`
	Description   string        `json:"description"`
	ExternalURLs  *ExternalURLs `json:"external_urls"`
	Followers     *Follower     `json:"followers"`
	Href          string        `json:"href"`
	ID            string        `json:"id"`
	Images        []*Image      `json:"images"`
	Name          string        `json:"name"`
	Owner         *User         `json:"owner"`
	Public        bool          `json:"public,omitempty"`
	SnapshotID    string        `json:"snapshot_id"`
	Tracks        *Paging       `json:"tracks"`
	Type          string        `json:"type"`
	URI           string        `json:"uri"`
}

type PlaylistTrack struct {
	AddedAt string `json:"added_at"`
	AddedBy *User  `json:"added_by"`
	IsLocal bool   `json:"is_local"`
	Track   *Track `json:"track"`
}

type Recommendations struct {
	Seeds  []*RecommendationSeed `json:"seeds"`
	Tracks []*Track              `json:"tracks"`
}

type RecommendationSeed struct {
	AfterFilteringSize int    `json:"afterFilteringSize"`
	AfterRelinkSize    int    `json:"afterRelinkSize"`
	Href               string `json:"href"`
	ID                 string `json:"id"`
	InitialPoolSize    int    `json:"initialPoolSize"`
	Type               string `json:"type"`
}

type SavedTrack struct {
	AddedAt string `json:"added_at"`
	Track   *Track `json:"track"`
}

type SavedAlbum struct {
	AddedAt string `json:"added_at"`
	Album   *Album `json:"album"`
}

type Track struct {
	Album            *Album        `json:"album"`
	Artists          []*Artist     `json:"artists"`
	AvailableMarkets []string      `json:"available_markets"`
	DiscNumber       int           `json:"disc_number"`
	DurationMs       int           `json:"duration_ms"`
	Explicit         bool          `json:"explicit"`
	ExternalIDs      *ExternalIDs  `json:"external_ids"`
	ExternalURLs     *ExternalURLs `json:"external_urls"`
	Href             string        `json:"href"`
	ID               string        `json:"id"`
	IsPlayable       bool          `json:"is_playable"`
	LinkedFrom       *LinkedTrack  `json:"linked_from"`
	Name             string        `json:"name"`
	Popularity       int           `json:"popularity"`
	PreviewURL       string        `json:"preview_url"`
	TrackNumber      int           `json:"track_number"`
	Type             string        `json:"type"`
	URI              string        `json:"uri"`
}

type LinkedTrack struct {
	ExternalURLs *ExternalURLs `json:"external_urls"`
	Href         string        `json:"href"`
	ID           string        `json:"id"`
	Type         string        `json:"type"`
	URI          string        `json:"uri"`
}

type User struct {
	Birthdate    string        `json:"birthdate"`
	Country      string        `json:"country"`
	DisplayName  string        `json:"display_name"`
	Email        string        `json:"email"`
	ExternalURLs *ExternalURLs `json:"external_urls"`
	Followers    *Follower     `json:"followers"`
	Href         string        `json:"href"`
	ID           string        `json:"id"`
	Images       []*Image      `json:"images"`
	Product      string        `json:"product"`
	Type         string        `json:"type"`
	URI          string        `json:"uri"`
}

type auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
