// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package api implements unofficial Medium's REST API.
package api

// User represents a Medium's user data.
type User struct {
	Success bool `json:"success"`
	Payload struct {
		User struct {
			UserID                   string `json:"userId"`
			Name                     string `json:"name"`
			Username                 string `json:"username"`
			CreatedAt                int64  `json:"createdAt"`
			LastPostCreatedAt        int64  `json:"lastPostCreatedAt"`
			ImageID                  string `json:"imageId"`
			BackgroundImageID        string `json:"backgroundImageId"`
			Bio                      string `json:"bio"`
			TwitterScreenName        string `json:"twitterScreenName"`
			FacebookAccountID        string `json:"facebookAccountId"`
			AllowNotes               int    `json:"allowNotes"`
			MediumMemberAt           int    `json:"mediumMemberAt"`
			MediumMemberWaitlistedAt int    `json:"mediumMemberWaitlistedAt"`
			Type                     string `json:"type"`
		} `json:"user"`
		StreamItems []struct {
			CreatedAt int64 `json:"createdAt"`
			Heading   struct {
				Text    string `json:"text"`
				Heading struct {
					FallbackTitle string `json:"fallbackTitle"`
					HeadingBasic  struct {
						Title string `json:"title"`
					} `json:"headingBasic"`
					HeadingType string `json:"headingType"`
				} `json:"heading"`
			} `json:"heading,omitempty"`
			RandomID    string `json:"randomId"`
			ItemType    string `json:"itemType"`
			Type        string `json:"type"`
			PostPreview struct {
				PostID string `json:"postId"`
			} `json:"postPreview,omitempty"`
		} `json:"streamItems"`
		UserMeta struct {
			NumberOfPostsPublished int `json:"numberOfPostsPublished"`
			InterestTags           []struct {
				Slug      string `json:"slug"`
				Name      string `json:"name"`
				PostCount int    `json:"postCount"`
				Virtuals  struct {
					IsFollowing bool `json:"isFollowing"`
				} `json:"virtuals"`
				Metadata struct {
					FollowerCount int `json:"followerCount"`
					PostCount     int `json:"postCount"`
					CoverImage    struct {
						ID string `json:"id"`
					} `json:"coverImage"`
				} `json:"metadata"`
				Type string `json:"type"`
			} `json:"interestTags"`
			UserID               string `json:"userId"`
			UserSuggestionReason struct {
				FolloweesWhoFollow struct {
					Users []interface{} `json:"users"`
				} `json:"followeesWhoFollow"`
				Reason string `json:"reason"`
			} `json:"userSuggestionReason"`
			CollectionIds []interface{} `json:"collectionIds"`
			AuthorTags    []struct {
				Slug      string `json:"slug"`
				Name      string `json:"name"`
				PostCount int    `json:"postCount"`
				Virtuals  struct {
					IsFollowing bool `json:"isFollowing"`
				} `json:"virtuals"`
				Metadata struct {
					FollowerCount int `json:"followerCount"`
					PostCount     int `json:"postCount"`
					CoverImage    struct {
						ID             string `json:"id"`
						OriginalWidth  int    `json:"originalWidth"`
						OriginalHeight int    `json:"originalHeight"`
						IsFeatured     bool   `json:"isFeatured"`
					} `json:"coverImage"`
				} `json:"metadata"`
				Type string `json:"type"`
			} `json:"authorTags"`
			FeaturedPostID  string        `json:"featuredPostId"`
			TopWriterInTags []interface{} `json:"topWriterInTags"`
			Type            string        `json:"type"`
		} `json:"userMeta"`
		UserNavItemList struct {
			UserNavItems []struct {
				Title      string `json:"title"`
				URL        string `json:"url"`
				SystemItem struct {
					SystemType int `json:"systemType"`
				} `json:"systemItem"`
				NavType string `json:"navType"`
			} `json:"userNavItems"`
		} `json:"userNavItemList"`
		UserNavActiveIndex int    `json:"userNavActiveIndex"`
		ProfileTypeName    string `json:"profileTypeName"`
		References         struct {
			User        UserReference            `json:"User"`
			Post        map[string]PostReference `json:"Post"`
			Social      SocialReference          `json:"Social"`
			SocialStats SocialStatsReference     `json:"SocialStats"`
		} `json:"references"`
		Paging struct {
			Path string `json:"path"`
			Next struct {
				Limit  int    `json:"limit"`
				To     string `json:"to"`
				Source string `json:"source"`
				Page   int    `json:"page"`
			} `json:"next"`
		} `json:"paging"`
	} `json:"payload"`
	V int    `json:"v"`
	B string `json:"b"`
}

// UserReference represents a reference of user.
type UserReference struct {
	UserID                   string `json:"userId"`
	Name                     string `json:"name"`
	Username                 string `json:"username"`
	CreatedAt                int64  `json:"createdAt"`
	LastPostCreatedAt        int64  `json:"lastPostCreatedAt"`
	ImageID                  string `json:"imageId"`
	BackgroundImageID        string `json:"backgroundImageId"`
	Bio                      string `json:"bio"`
	TwitterScreenName        string `json:"twitterScreenName"`
	FacebookAccountID        string `json:"facebookAccountId"`
	AllowNotes               int    `json:"allowNotes"`
	MediumMemberAt           int    `json:"mediumMemberAt"`
	MediumMemberWaitlistedAt int    `json:"mediumMemberWaitlistedAt"`
	Type                     string `json:"type"`
}

// SocialReference represents a reference of social.
type SocialReference struct {
	UserID       string `json:"userId"`
	TargetUserID string `json:"targetUserId"`
	Type         string `json:"type"`
}

// SocialStatsReference represents a reference of social stats.
type SocialStatsReference struct {
	UserID               string `json:"userId"`
	UsersFollowedCount   int    `json:"usersFollowedCount"`
	UsersFollowedByCount int    `json:"usersFollowedByCount"`
	Type                 string `json:"type"`
}

// PostReference represents a reference of post.
type PostReference struct {
	ID                     string `json:"id"`
	VersionID              string `json:"versionId"`
	CreatorID              string `json:"creatorId"`
	HomeCollectionID       string `json:"homeCollectionId"`
	Title                  string `json:"title"`
	DetectedLanguage       string `json:"detectedLanguage"`
	LatestVersion          string `json:"latestVersion"`
	LatestPublishedVersion string `json:"latestPublishedVersion"`
	HasUnpublishedEdits    bool   `json:"hasUnpublishedEdits"`
	LatestRev              int    `json:"latestRev"`
	CreatedAt              int64  `json:"createdAt"`
	UpdatedAt              int64  `json:"updatedAt"`
	AcceptedAt             int    `json:"acceptedAt"`
	FirstPublishedAt       int64  `json:"firstPublishedAt"`
	LatestPublishedAt      int64  `json:"latestPublishedAt"`
	Vote                   bool   `json:"vote"`
	ExperimentalCSS        string `json:"experimentalCss"`
	DisplayAuthor          string `json:"displayAuthor"`
	Content                struct {
		PostDisplay struct {
			Coverless bool `json:"coverless"`
		} `json:"postDisplay"`
	} `json:"content"`
	Virtuals struct {
		AllowNotes   bool `json:"allowNotes"`
		PreviewImage struct {
			ImageID        string `json:"imageId"`
			Filter         string `json:"filter"`
			BackgroundSize string `json:"backgroundSize"`
			OriginalWidth  int    `json:"originalWidth"`
			OriginalHeight int    `json:"originalHeight"`
			Strategy       string `json:"strategy"`
			Height         int    `json:"height"`
			Width          int    `json:"width"`
		} `json:"previewImage"`
		WordCount               int           `json:"wordCount"`
		ImageCount              int           `json:"imageCount"`
		ReadingTime             float64       `json:"readingTime"`
		Subtitle                string        `json:"subtitle"`
		UsersBySocialRecommends []interface{} `json:"usersBySocialRecommends"`
		Recommends              int           `json:"recommends"`
		IsBookmarked            bool          `json:"isBookmarked"`
		Tags                    []struct {
			Slug      string `json:"slug"`
			Name      string `json:"name"`
			PostCount int    `json:"postCount"`
			Virtuals  struct {
				IsFollowing bool `json:"isFollowing"`
			} `json:"virtuals"`
			Metadata struct {
				FollowerCount int `json:"followerCount"`
				PostCount     int `json:"postCount"`
				CoverImage    struct {
					ID             string `json:"id"`
					OriginalWidth  int    `json:"originalWidth"`
					OriginalHeight int    `json:"originalHeight"`
					IsFeatured     bool   `json:"isFeatured"`
				} `json:"coverImage"`
			} `json:"metadata"`
			Type string `json:"type"`
		} `json:"tags"`
		SocialRecommendsCount int `json:"socialRecommendsCount"`
		ResponsesCreatedCount int `json:"responsesCreatedCount"`
		Links                 struct {
			Entries     []interface{} `json:"entries"`
			Version     string        `json:"version"`
			GeneratedAt int64         `json:"generatedAt"`
		} `json:"links"`
		IsLockedPreviewOnly bool   `json:"isLockedPreviewOnly"`
		TakeoverID          string `json:"takeoverId"`
		MetaDescription     string `json:"metaDescription"`
		TotalClapCount      int    `json:"totalClapCount"`
	} `json:"virtuals"`
	Coverless                  bool   `json:"coverless"`
	Slug                       string `json:"slug"`
	TranslationSourcePostID    string `json:"translationSourcePostId"`
	TranslationSourceCreatorID string `json:"translationSourceCreatorId"`
	IsApprovedTranslation      bool   `json:"isApprovedTranslation"`
	InResponseToPostID         string `json:"inResponseToPostId"`
	InResponseToRemovedAt      int    `json:"inResponseToRemovedAt"`
	IsTitleSynthesized         bool   `json:"isTitleSynthesized"`
	AllowResponses             bool   `json:"allowResponses"`
	ImportedURL                string `json:"importedUrl"`
	ImportedPublishedAt        int    `json:"importedPublishedAt"`
	Visibility                 int    `json:"visibility"`
	UniqueSlug                 string `json:"uniqueSlug"`
	PreviewContent             struct {
		BodyModel struct {
			Paragraphs []struct {
				Name      string `json:"name"`
				Type      int    `json:"type"`
				Text      string `json:"text"`
				Alignment int    `json:"alignment"`
			} `json:"paragraphs"`
			Sections []struct {
				StartIndex int `json:"startIndex"`
			} `json:"sections"`
		} `json:"bodyModel"`
		IsFullContent bool `json:"isFullContent"`
	} `json:"previewContent"`
	License                     int    `json:"license"`
	InResponseToMediaResourceID string `json:"inResponseToMediaResourceId"`
	CanonicalURL                string `json:"canonicalUrl"`
	ApprovedHomeCollectionID    string `json:"approvedHomeCollectionId"`
	NewsletterID                string `json:"newsletterId"`
	SuggestionReason            struct {
		Reason int `json:"reason"`
	} `json:"suggestionReason"`
	WebCanonicalURL              string `json:"webCanonicalUrl"`
	MediumURL                    string `json:"mediumUrl"`
	MigrationID                  string `json:"migrationId"`
	NotifyFollowers              bool   `json:"notifyFollowers"`
	NotifyTwitter                bool   `json:"notifyTwitter"`
	IsSponsored                  bool   `json:"isSponsored"`
	IsRequestToPubDisabled       bool   `json:"isRequestToPubDisabled"`
	NotifyFacebook               bool   `json:"notifyFacebook"`
	ResponseHiddenOnParentPostAt int    `json:"responseHiddenOnParentPostAt"`
	IsSeries                     bool   `json:"isSeries"`
	IsSubscriptionLocked         bool   `json:"isSubscriptionLocked"`
	SeriesLastAppendedAt         int    `json:"seriesLastAppendedAt"`
	AudioVersionDurationSec      int    `json:"audioVersionDurationSec"`
	Type                         string `json:"type"`
}

func (u *User) IsSuccess() bool {
	return u.Success
}

func (u *User) PostReferences() []PostReference {
	posts := []PostReference{} // do not use `make`. `u.Payload.References.Post` length is ramdomly (contains empty).
	for _, v := range u.Payload.References.Post {
		posts = append(posts, v)
	}
	return posts
}
