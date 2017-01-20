package spotify

import (
	"strings"
)

var base = "https://api.spotify.com/v1/"

func EndpointMe() string                { return base + "me" }
func EndpointGetUser(uid string) string { return base + "users/" + uid }

// ==================== ALBUMS ====================

func EndpointGetAlbum(id string) string       { return base + "albums/" + id }
func EndpointGetAlbums(ids []string) string   { return base + "albums?ids=" + strings.Join(ids, ",") }
func EndpointGetAlbumTracks(id string) string { return EndpointGetAlbum(id) + "/tracks" }
func EndpointSaveAlbums(ids []string) string {
	return EndpointMe() + "/albums?ids=" + strings.Join(ids, ",")
}
func EndpointGetSavedAlbums() string { return EndpointMe() + "/albums" }
func EndpointDeleteAlbums(ids []string) string {
	return EndpointGetSavedAlbums() + "?ids=" + strings.Join(ids, ",")
}
func EndpointContainsAlbums(ids []string) string {
	return EndpointGetSavedAlbums() + "/contains?ids=" + strings.Join(ids, ",")
}
func EndpointSearch(name, typ string) string { return base + "search?q=" + name + "&type=" + typ }

// ==================== END ALBUMS ====================

// ==================== ARTISTS ====================

func EndpointGetArtist(id string) string            { return base + "artsts/" + id }
func EndpointGetArtists(ids []string) string        { return base + "artists?ids=" + strings.Join(ids, ",") }
func EndpointGetArtistAlbums(id string) string      { return EndpointGetArtist(id) + "/albums" }
func EndpointGetArtistTopTracks(id string) string   { return EndpointGetArtist(id) + "/top-tracks" }
func EndpointGetRelatedArtists(id string) string    { return EndpointGetArtist(id) + "/related-artists" }
func EndpointGetTopArtistOrTrack(typ string) string { return EndpointMe() + "/top/" + typ }

// ==================== END ARTISTS ====================

// ==================== BROWSE ====================

func EndpointBrowseFeaturedPlaylists() string       { return base + "browse/featured-playlists" }
func EndpointBrowseNewReleases() string             { return base + "browse/new-releases" }
func EndpointBrowseCategories() string              { return base + "browse/categories" }
func EndpointGetCategory(id string) string          { return EndpointBrowseCategories() + "/" + id }
func EndpointGetCategoryPlaylists(id string) string { return EndpointGetCategory(id) + "/playlists" }
func EndpointGetRecommendations() string            { return base + "recommendations" }

// ==================== END BROWSE ====================

// ==================== FOLLOW ====================

func EndpointGetFollowedArtists() string { return EndpointMe() + "/following" }
func EndpointFollowArtists(ids []string, typ string) string {
	return EndpointMe() + "/following?ids=" + strings.Join(ids, ",") + "&type=" + typ
}
func EndpointUnfollowArtists(ids []string, typ string) string {
	return EndpointMe() + "/following?ids=" + strings.Join(ids, ",") + "&type=" + typ
}
func EndpointFollowsArtists(ids []string, typ string) string {
	return EndpointMe() + "/following/contains?type=" + typ + "&ids=" + strings.Join(ids, ",")
}
func EndpointFollowPlaylist(uid, pid string) string {
	return EndpointGetUser(uid) + "/playlists/" + pid + "/followers"
}
func EndpointUnfollowPlaylist(oid, pid string) string { return EndpointFollowPlaylist(oid, pid) }
func EndpointUsersFollowsPlaylist(oid, pid string, ids []string) string {
	return EndpointFollowPlaylist(oid, pid) + "/contains?ids=" + strings.Join(ids, ",")
}

// ==================== END FOLLOW ====================

// ==================== LIBRARY ====================

func EndpointSaveTracks(ids []string) string {
	return EndpointMe() + "/tracks?ids=" + strings.Join(ids, ",")
}
func EndpointGetSavedTracks() string                { return EndpointMe() + "/tracks" }
func EndpointRemoveSavedTracks(ids []string) string { return EndpointSaveTracks(ids) }
func EndpointHasTracksSaved(ids []string) string {
	return EndpointGetSavedTracks() + "/contains?ids=" + strings.Join(ids, ",")
}

// ==================== ENDLIBRARY ====================

// ==================== PLAYLISTS ====================

func EndpointGetUserPlaylist(uid, pid string) string {
	return EndpointGetUser(uid) + "/playlists/" + pid
}
func EndpointGetPlaylistTracks(uid, pid string) string {
	return EndpointGetUserPlaylist(uid, pid) + "/tracks"
}
func EndpointCreatePlaylist(uid string) string             { return EndpointGetUser(uid) + "/playlists" }
func EndpointChangePlaylistDetails(uid, pid string) string { return EndpointGetUserPlaylist(uid, pid) }
func EndpointAddTracksToPlaylist(uid, pid string) string   { return EndpointGetPlaylistTracks(uid, pid) }
func EndpointDeleteTracksFromPlaylist(uid, pid string) string {
	return EndpointGetPlaylistTracks(uid, pid)
}
func EndpointReorderTracksInPlaylist(uid, pid string) string {
	return EndpointGetPlaylistTracks(uid, pid)
}
func EndpointReplaceTracksInPlaylist(uid, pid string) string {
	return EndpointGetPlaylistTracks(uid, pid)
}

// ==================== END PLAYLISTS ====================

// ==================== TRACKS ====================

func EndpointGetAudioAnalysis(sid string) string { return base + "audio-analysis/" + sid }
func EndpointGetAudioFeature(sid string) string  { return base + "audio-features/" + sid }
func EndpointGetAudioFeatures(sids []string) string {
	return base + "audio-features?ids=" + strings.Join(sids, ",")
}
func EndpointGetTrack(sid string) string     { return base + "tracks/" + sid }
func EndpointGetTracks(sids []string) string { return base + "tracks?ids=" + strings.Join(sids, ",") }
