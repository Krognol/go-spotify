package spotify

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	auth         *auth
	ClientID     string
	ClientSecret string
}

func New(clientid, clientsecret string) *Client {
	return &Client{auth: &auth{}, ClientID: clientid, ClientSecret: clientsecret}
}

func (c *Client) authorize() error {
	httpc := &http.Client{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpc.Transport = tr
	body := strings.NewReader(`grant_type=client_credentials`)
	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)
	req.Header.Add("cache-control", "no-cache")
	req.SetBasicAuth(c.ClientID, c.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := httpc.Do(req)
	if err != nil {
		return err
	}

	err = unmarshal(res, c.auth)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) request(method, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)

	err := c.authorize()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.auth.AccessToken)
	return http.DefaultClient.Do(req)
}

func unmarshal(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func GetAlbum(id string) (*Album, error) {
	res, err := http.Get(EndpointGetAlbum(id))
	if err != nil {
		return nil, err
	}
	album := &Album{}
	err = unmarshal(res, album)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func GetAlbums(ids []string) ([]*Album, error) {
	res, err := http.Get(EndpointGetAlbums(ids))
	if err != nil {
		return nil, err
	}

	var albums []*Album
	err = unmarshal(res, albums)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func GetAlbumTracks(id string) (*Paging, error) {
	res, err := http.Get(EndpointGetAlbumTracks(id))

	if err != nil {
		return nil, err
	}

	page := &Paging{}

	err = unmarshal(res, page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func GetArtist(id string) (*Artist, error) {
	res, err := http.Get(EndpointGetArtist(id))
	if err != nil {
		return nil, err
	}

	a := &Artist{}

	err = unmarshal(res, a)

	if err != nil {
		return nil, err
	}
	return a, nil
}

func GetArtists(ids []string) ([]*Artist, error) {
	res, err := http.Get(EndpointGetArtists(ids))
	if err != nil {
		return nil, err
	}

	var as []*Artist
	err = unmarshal(res, as)

	if err != nil {
		return nil, err
	}
	return as, nil
}

func GetArtistAlbums(id string) (*Paging, error) {
	res, err := http.Get(EndpointGetArtistAlbums(id))
	if err != nil {
		return nil, err
	}

	page := &Paging{}

	err = unmarshal(res, page)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func GetArtistTopTracks(id string) ([]*Track, error) {
	res, err := http.Get(EndpointGetArtistTopTracks(id))
	if err != nil {
		return nil, err
	}

	var tracks []*Track

	err = unmarshal(res, tracks)

	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func GetRelatedArtists(id string) ([]*Artist, error) {
	res, err := http.Get(EndpointGetRelatedArtists(id))
	if err != nil {
		return nil, err
	}

	var artists []*Artist

	err = unmarshal(res, artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (c *Client) GetAudioAnalysis(id string) (*AudioAnalysis, error) {
	res, err := c.request("GET", EndpointGetAudioAnalysis(id), nil)

	if err != nil {
		return nil, err
	}
	analysis := &AudioAnalysis{}

	err = unmarshal(res, analysis)
	if err != nil {
		return nil, err
	}
	return analysis, nil
}

func (c *Client) GetAudioFeature(id string) (*AudioFeatures, error) {
	res, err := c.request("GET", EndpointGetAudioFeature(id), nil)

	if err != nil {
		return nil, err
	}

	feat := &AudioFeatures{}
	err = unmarshal(res, feat)
	if err != nil {
		return nil, err
	}
	return feat, nil
}

func (c *Client) GetAudioFeatures(ids []string) ([]*AudioFeatures, error) {
	res, err := c.request("GET", EndpointGetAudioFeatures(ids), nil)

	if err != nil {
		return nil, err
	}

	var feats []*AudioFeatures

	err = unmarshal(res, feats)

	if err != nil {
		return nil, err
	}
	return feats, nil
}

// GetFeaturedPlaylists gets the featured playlists
// The optional parameters can switch it up
// Locale: a lowercase ISO 639 language code and an uppercase ISO 3166-1 alpha-2 country code
// e.g: "es_MX" (Spanish (Mexico))
// Country: ISO 3166-1 alpha-2 country code, e.g: "MX" (Mexico)
// Timestamp:  ISO 8601 format, 'yyyy-MM-ddTHH:mm:ss', e.g: "2014-10-02T09:00:00"
// Limit: Max amount of items, Default is 20, minimum is 1 and maximum is 50
// Offset: The index of the first object, default is 0
func (c *Client) GetFeaturedPlaylists(locale, country, timestamp string, limit, offset int) (*Paging, error) {
	vals := &url.Values{}
	vals.Add("locale", locale)
	vals.Add("country", country)
	vals.Add("timestamp", timestamp)
	vals.Add("limit", strconv.Itoa(limit))
	vals.Add("offset", strconv.Itoa(offset))
	if locale == "" {
		vals.Del("locale")
	}

	if country == "" {
		vals.Del("country")
	}

	if timestamp == "" {
		vals.Del("timestamp")
	}

	if limit < 0 {
		limit = 20
	}

	if offset < 0 {
		offset = 0
	}
	res, err := c.request("GET", EndpointBrowseFeaturedPlaylists()+"?"+vals.Encode(), nil)

	if err != nil {
		return nil, err
	}

	lists := &Paging{}

	err = unmarshal(res, lists)

	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (c *Client) GetNewReleases(country string, limit, offset int) (*Paging, error) {
	vals := url.Values{}
	if country != "" {
		vals.Add("country", country)
	}

	if limit < 0 {
		vals.Add("limit", "20")
	} else if limit > 50 {
		vals.Add("limit", "50")
	} else {
		vals.Add("limit", strconv.Itoa(limit))
	}

	if offset < 0 {
		vals.Add("offset", "0")
	} else {
		vals.Add("offset", strconv.Itoa(offset))
	}

	res, err := c.request("GET", EndpointBrowseNewReleases()+"?"+vals.Encode(), nil)
	if err != nil {
		return nil, err
	}

	page := &Paging{}
	err = unmarshal(res, page)

	if err != nil {
		return nil, err
	}
	return page, nil
}

func (c *Client) GetCategories(country, locale string, offset, limit int) (*Paging, error) {
	vals := url.Values{}
	if country != "" {
		vals.Add("country", country)
	}

	if locale != "" {
		vals.Add("locale", locale)
	}

	if limit < 0 {
		vals.Add("limit", "20")
	} else if limit > 50 {
		vals.Add("limit", "50")
	} else {
		vals.Add("limit", strconv.Itoa(limit))
	}

	if offset < 0 {
		vals.Add("offset", "0")
	} else {
		vals.Add("offset", strconv.Itoa(offset))
	}

	res, err := c.request("GET", EndpointBrowseCategories()+"?"+vals.Encode(), nil)
	if err != nil {
		return nil, err
	}

	page := &Paging{}
	err = unmarshal(res, page)

	if err != nil {
		return nil, err
	}
	return page, nil
}

func (c *Client) GetCategory(name, country, locale string) (*Category, error) {
	vals := url.Values{}
	if country != "" {
		vals.Add("country", country)
	}

	if locale != "" {
		vals.Add("locale", locale)
	}
	en := vals.Encode()
	res, err := c.request("GET", EndpointGetCategory(name)+"?"+en, nil)
	if err != nil {
		return nil, err
	}

	category := &Category{}
	err = unmarshal(res, category)

	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Client) GetCategoryPlaylists(name, country string, limit, offset int) (*Paging, error) {
	if name == "" {
		return nil, fmt.Errorf("Missing required parameter: name")
	}

	vals := url.Values{}
	if country != "" {
		vals.Add("country", country)
	}

	if limit < 1 {
		vals.Add("limit", "1")
	} else if limit > 50 {
		vals.Add("limit", "50")
	} else {
		vals.Add("limit", strconv.Itoa(limit))
	}

	if offset < 0 {
		vals.Add("offset", "0")
	} else {
		vals.Add("offset", strconv.Itoa(offset))
	}

	res, err := c.request("GET", EndpointGetCategoryPlaylists(name), nil)
	if err != nil {
		return nil, err
	}

	page := &Paging{}

	err = unmarshal(res, page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (c *Client) GetRecommendations(args ...string) (*Recommendations, error) {
	res, err := c.request("GET", EndpointGetRecommendations(args...), nil)

	if err != nil {
		return nil, err
	}

	rec := &Recommendations{Seeds: make([]*RecommendationSeed, 0), Tracks: make([]*Track, 0)}

	err = unmarshal(res, rec)

	if err != nil {
		return nil, err
	}

	return rec, nil
}

func (c *Client) UserFollowPlaylist(oid, pid string) error {
	_, err := c.request("PUT", EndpointFollowPlaylist(oid, pid), nil)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UserUnfollowPlaylist(oid, pid string) error {
	_, err := c.request("DELETE", EndpointUnfollowPlaylist(oid, pid), nil)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UsersFollowsPlaylist(oid, pid string, uid []string) ([]bool, error) {
	res, err := c.request("GET", EndpointUsersFollowsPlaylist(oid, pid, uid), nil)

	if err != nil {
		return nil, err
	}

	var bools = make([]bool, 0)

	err = unmarshal(res, bools)

	if err != nil {
		return nil, err
	}

	return bools, nil
}

func (c *Client) SearchTrack(query string, offset int) ([]*Track, error) {
	url := EndpointSearch(url.QueryEscape(query), "track")
	if offset != 0 {
		url += fmt.Sprintf("&offset=%d", offset)
	}

	res, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}

	type temp struct {
		Page *Paging `json:"tracks"`
	}

	var t temp

	if err = unmarshal(res, &t); err != nil {
		return nil, err
	}

	var result []*Track
	json.Unmarshal(t.Page.Items, &result)

	return result, nil
}

func (c *Client) SearchAlbum(query string, offset int) ([]*Album, error) {
	url := EndpointSearch(url.QueryEscape(query), "album")
	if offset != 0 {
		url += fmt.Sprintf("&offset=%d", offset)
	}

	res, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}

	type temp struct {
		Page *Paging `json:"albums"`
	}

	var t temp
	if err = unmarshal(res, &t); err != nil {
		return nil, err
	}

	var result []*Album
	json.Unmarshal(t.Page.Items, &result)
	return result, nil
}

func (c *Client) SearchArtist(query string, offset int) ([]*Artist, error) {
	url := EndpointSearch(url.QueryEscape(query), "artist")
	if offset != 0 {
		url += fmt.Sprintf("&offset=%d", offset)
	}

	res, err := c.request("GET", url, nil)
	if err == nil {
		type temp struct {
			Page *Paging `json:"artists"`
		}

		var t temp
		if err = unmarshal(res, &t); err == nil {
			var result []*Artist
			err = json.Unmarshal(t.Page.Items, &result)
			return result, err
		}
	}
	return nil, err
}

// the 'Me' endpoints don't work
