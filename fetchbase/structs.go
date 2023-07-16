package fetchbase

type FetchBaseInput struct {
	DocumentId string                 `json:"documentId"`
	Variables  map[string]interface{} `json:"variables"`
}
type ThreadsUserProfileResponse struct {
	data       Data
	extensions Extensions
}

type Data struct {
	userData UserData
}

type UserData struct {
	user User
}

type User struct {
	is_private                     bool
	profile_pic_url                string
	username                       string
	hd_profile_pic_versions        []HDProfilePicVersion
	is_verified                    bool
	biography                      string
	biography_with_entities        int
	follower_count                 int
	profile_context_facepile_users int
	bio_links                      []BioLink
	pk                             string
	full_name                      string
	id                             int
}

type BioLink struct {
	url string
}

type HDProfilePicVersion struct {
	height int
	url    string
	width  int
}

type Extensions struct {
	is_final bool
}
