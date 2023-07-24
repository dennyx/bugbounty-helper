package entity

type HackeroneGetAllProgramsAwardBountyResponse struct {
	Me struct {
		ID                    string `json:"id"`
		EditUnclaimedProfiles bool   `json:"edit_unclaimed_profiles"`
		Typename              string `json:"__typename"`
	} `json:"me"`
	Teams struct {
		PageInfo struct {
			EndCursor   string `json:"endCursor"`
			HasNextPage bool   `json:"hasNextPage"`
			Typename    string `json:"__typename"`
		} `json:"pageInfo"`
		Edges []struct {
			Node struct {
				ID                    string `json:"id"`
				Bookmarked            bool   `json:"bookmarked"`
				ResolvedReportCount   int    `json:"resolved_report_count"`
				ProfilePicture        string `json:"profile_picture"`
				Name                  string `json:"name"`
				Handle                string `json:"handle"`
				SubmissionState       string `json:"submission_state"`
				TriageActive          bool   `json:"triage_active"`
				PubliclyVisibleRetest bool   `json:"publicly_visible_retesting"`
				State                 string `json:"state"`
				AllowsBountySplitting bool   `json:"allows_bounty_splitting"`
				ExternalProgram       struct {
					ID            string `json:"id"`
					OffersRewards bool   `json:"offers_rewards"`
					__typename    string `json:"__typename"`
				} `json:"external_program"`
				__typename string `json:"__typename"`
			} `json:"node"`
			__typename string `json:"__typename"`
		} `json:"edges"`
		Typename string `json:"__typename"`
	} `json:"teams"`
}

type TeamTableResolvedReports struct {
	ID       string `json:"id"`
	Reported int    `json:"resolved_report_count"`
	Typename string `json:"__typename"`
}

type TeamTableAvatarAndTitle struct {
	ID                       string `json:"id"`
	ProfilePicture           string `json:"profile_picture(size: medium)"`
	Name                     string `json:"name"`
	Handle                   string `json:"handle"`
	SubmissionState          string `json:"submission_state"`
	TriageActive             bool   `json:"triage_active"`
	PubliclyVisibleRetesting bool   `json:"publicly_visible_retesting"`
	State                    string `json:"state"`
	AllowsBountySplitting    bool   `json:"allows_bounty_splitting"`
	ExternalProgram          struct {
		ID       string `json:"id"`
		Typename string `json:"__typename"`
	} `json:"external_program"`
	TeamLinkWithMiniProfile `json:"...TeamLinkWithMiniProfile"`
	Typename                string `json:"__typename"`
}

type TeamLinkWithMiniProfile struct {
	ID       string `json:"id"`
	Handle   string `json:"handle"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type TeamTableLaunchDate struct {
	ID       string `json:"id"`
	Launched string `json:"launched_at"`
	Typename string `json:"__typename"`
}

type TeamTableMinimumBounty struct {
	ID       string `json:"id"`
	Currency string `json:"currency"`
	Base     int    `json:"base_bounty"`
	Typename string `json:"__typename"`
}

type TeamTableAverageBounty struct {
	ID           string `json:"id"`
	Currency     string `json:"currency"`
	AverageLower int    `json:"average_bounty_lower_amount"`
	AverageUpper int    `json:"average_bounty_upper_amount"`
	Typename     string `json:"__typename"`
}

type BookmarkTeam struct {
	ID         string `json:"id"`
	Bookmarked bool   `json:"bookmarked"`
	Typename   string `json:"__typename"`
}
