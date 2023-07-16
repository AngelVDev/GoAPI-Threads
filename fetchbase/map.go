package fetchbase

func mapUserProfile(rawResponse ThreadsUserProfileResponse) interface{} {
	userApiResponse := rawResponse.data.userData.user

	// If the User field is empty, return nil
	if (User{} == userApiResponse) {
		return nil
	}

	// Equivalent logic to JavaScript destructuring
	username := userApiResponse.username
	isVerified := userApiResponse.is_verified
	biography := userApiResponse.biography
	followerCount := userApiResponse.follower_count
	bioLinks := userApiResponse.bio_links
	id := userApiResponse.pk
	fullName := userApiResponse.full_name
	hdProfilePicVersions := userApiResponse.hd_profile_pic_versions
	profilePicUrl := userApiResponse.profile_pic_url

	// Create new profile pics slice
	profilePics := make([]HDProfilePicVersion, len(hdProfilePicVersions)+1)

	// Add the first profile pic to the slice
	profilePics[0] = HDProfilePicVersion{
		height: 150,
		width:  150,
		url:    profilePicUrl,
	}

	// Add the rest of the profile pics from hdProfilePicVersions
	for i, pic := range hdProfilePicVersions {
		profilePics[i+1] = pic
	}

	// Create an anonymous struct to hold the return values
	result := struct {
		ID            string
		Username      string
		IsVerified    bool
		Biography     string
		FollowerCount int
		BioLinks      []BioLink
		FullName      string
		ProfilePics   []HDProfilePicVersion
	}{ID: id, Username: username, IsVerified: isVerified, Biography: biography, FollowerCount: followerCount, BioLinks: bioLinks, FullName: fullName, ProfilePics: profilePics}

	return result
}
