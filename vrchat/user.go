package vrchat

import "time"

type UserLoginResponseBody struct {
	ID                 string   `json:"id"`
	Username           string   `json:"username"`
	DisplayName        string   `json:"displayName"`
	UserIcon           string   `json:"userIcon"`
	Bio                string   `json:"bio"`
	BioLinks           []string `json:"bioLinks"`
	ProfilePicOverride string   `json:"profilePicOverride"`
	StatusDescription  string   `json:"statusDescription"`
	PastDisplayNames   []struct {
		DisplayName string    `json:"displayName"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"pastDisplayNames"`
	HasEmail                       bool          `json:"hasEmail"`
	HasPendingEmail                bool          `json:"hasPendingEmail"`
	ObfuscatedEmail                string        `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail         string        `json:"obfuscatedPendingEmail"`
	EmailVerified                  bool          `json:"emailVerified"`
	HasBirthday                    bool          `json:"hasBirthday"`
	Unsubscribe                    bool          `json:"unsubscribe"`
	StatusHistory                  []string      `json:"statusHistory"`
	StatusFirstTime                bool          `json:"statusFirstTime"`
	Friends                        []string      `json:"friends"`
	FriendGroupNames               []interface{} `json:"friendGroupNames"`
	CurrentAvatarImageURL          string        `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string        `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatar                  string        `json:"currentAvatar"`
	CurrentAvatarAssetURL          string        `json:"currentAvatarAssetUrl"`
	FallbackAvatar                 string        `json:"fallbackAvatar"`
	AccountDeletionDate            interface{}   `json:"accountDeletionDate"`
	AcceptedTOSVersion             int           `json:"acceptedTOSVersion"`
	SteamID                        string        `json:"steamId"`
	SteamDetails                   struct {
	} `json:"steamDetails"`
	OculusID                 string    `json:"oculusId"`
	HasLoggedInFromClient    bool      `json:"hasLoggedInFromClient"`
	HomeLocation             string    `json:"homeLocation"`
	TwoFactorAuthEnabled     bool      `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate time.Time `json:"twoFactorAuthEnabledDate"`
	State                    string    `json:"state"`
	Tags                     []string  `json:"tags"`
	DeveloperType            string    `json:"developerType"`
	LastLogin                time.Time `json:"last_login"`
	LastPlatform             string    `json:"last_platform"`
	AllowAvatarCopying       bool      `json:"allowAvatarCopying"`
	Status                   string    `json:"status"`
	DateJoined               string    `json:"date_joined"`
	IsFriend                 bool      `json:"isFriend"`
	FriendKey                string    `json:"friendKey"`
	LastActivity             time.Time `json:"last_activity"`
	OnlineFriends            []string  `json:"onlineFriends"`
	ActiveFriends            []string  `json:"activeFriends"`
	OfflineFriends           []string  `json:"offlineFriends"`
}
