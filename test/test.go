package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

func main() {
	// Define your GraphQL endpoint URL
	endpoint := "https://hackerone.com/graphql"

	// Create a new GraphQL client
	client := graphql.NewClient(endpoint)

	// Define your GraphQL query
	query := `
        query DirectoryQuery(
            $cursor: String
            $secureOrderBy: FiltersTeamFilterOrder
            $where: FiltersTeamFilterInput
        ) {
            me {
                id
                edit_unclaimed_profiles
                __typename
            }
            teams(
                first: 25
                after: $cursor
                secure_order_by: $secureOrderBy
                where: $where
            ) {
                pageInfo {
                    endCursor
                    hasNextPage
                    __typename
                }
                edges {
                    node {
                        id
                        bookmarked
                        ...TeamTableResolvedReports
                        ...TeamTableAvatarAndTitle
                        ...TeamTableLaunchDate
                        ...TeamTableMinimumBounty
                        ...TeamTableAverageBounty
                        ...BookmarkTeam
                        __typename
                    }
                    __typename
                }
                __typename
            }
        }
        fragment TeamTableResolvedReports on Team {
            id
            resolved_report_count
            __typename
        }
        fragment TeamTableAvatarAndTitle on Team {
            id
            profile_picture(size: medium)
            name
            handle
            submission_state
            triage_active
            publicly_visible_retesting
            state
            allows_bounty_splitting
            external_program {
                id
                __typename
            }
            ...TeamLinkWithMiniProfile
            __typename
        }
        fragment TeamLinkWithMiniProfile on Team {
            id
            handle
            name
            __typename
        }
        fragment TeamTableLaunchDate on Team {
            id
            launched_at
            __typename
        }
        fragment TeamTableMinimumBounty on Team {
            id
            currency
            base_bounty
            __typename
        }
        fragment TeamTableAverageBounty on Team {
            id
            currency
            average_bounty_lower_amount
            average_bounty_upper_amount
            __typename
        }
        fragment BookmarkTeam on Team {
            id
            bookmarked
            __typename
        }
    `

	// Define your query variables
	variables := map[string]interface{}{
		"cursor":        "",
		"secureOrderBy": map[string]interface{}{"launched_at": map[string]string{"_direction": "DESC"}},
		"where": map[string]interface{}{
			"_and": []interface{}{
				map[string]interface{}{
					"_or": []interface{}{
						map[string]interface{}{
							"offers_bounties": map[string]bool{"_eq": true},
						},
						map[string]interface{}{
							"external_program": map[string]interface{}{
								"offers_rewards": map[string]bool{"_eq": true},
							},
						},
					},
				},
				map[string]interface{}{
					"_or": []interface{}{
						map[string]interface{}{"submission_state": map[string]string{"_eq": "open"}},
						map[string]interface{}{"submission_state": map[string]string{"_eq": "api_only"}},
						map[string]interface{}{"external_program": map[string]interface{}{}},
					},
				},
				map[string]interface{}{
					"_not": map[string]interface{}{
						"external_program": map[string]interface{}{},
					},
				},
				map[string]interface{}{
					"_or": []interface{}{
						map[string]interface{}{
							"_and": []interface{}{
								map[string]interface{}{"state": map[string]string{"_neq": "sandboxed"}},
								map[string]interface{}{"state": map[string]string{"_neq": "soft_launched"}},
							},
						},
						map[string]interface{}{"external_program": map[string]interface{}{}},
					},
				},
			},
		},
	}

	// Create a new GraphQL request with the query and variables
	req := graphql.NewRequest(query)
	req.Var("cursor", variables["cursor"])
	req.Var("secureOrderBy", variables["secureOrderBy"])
	req.Var("where", variables["where"])

	// Execute the GraphQL request and handle the response
	var respData struct {
		Me struct {
			ID                    string `json:"id"`
			EditUnclaimedProfiles bool   `json:"edit_unclaimed_profiles"`
			__typename            string `json:"__typename"`
		} `json:"me"`
		Teams struct {
			PageInfo struct {
				EndCursor   string `json:"endCursor"`
				HasNextPage bool   `json:"hasNextPage"`
				__typename  string `json:"__typename"`
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
			__typename string `json:"__typename"`
		} `json:"teams"`
	}
	err := client.Run(context.Background(), req, &respData)
	if err != nil {
		// Handle errors
		log.Default().Fatal(err)
	}

	// Use the response data
	for _, team := range respData.Teams.Edges {
		fmt.Println(team.Node.ID, team.Node.Name)
	}
}
