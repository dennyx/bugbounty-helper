/*
Copyright © 2023 dennyxzang@gmail.com
Security.Cloud
*/
package cmd

import (
	"bugbounty-helper/entity"
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
	"log"
)

// hackeroneCmd represents the hackerone command
var hackeroneCmd = &cobra.Command{
	Use:   "hackerone",
	Short: "Hackerone helper",
	Long:  `Hackerone helper`,
	RunE:  hackeroneCmdRunE,
}

func init() {
	rootCmd.AddCommand(hackeroneCmd)
}

func hackeroneCmdRunE(cmd *cobra.Command, args []string) error {
	fmt.Println("hackerone called")
	//getAllProgramsAwardBounty()
	// Get Detail for specified program
	getProgramDetailsOnScope()
	return nil
}

func getAllProgramsAwardBounty() {
	log.Default().Println("Preparing to get all the programs that award bounty...")
	// Get all the programs that award bounty
	client := graphql.NewClient(entity.HackeroneGraphqlEndpoint)
	// set variables
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

	// Loop to get all the programs that award bounty
	cursor := ""
	for {
		// Create a new GraphQL request with the query and variables
		req := graphql.NewRequest(entity.HackeroneGetAllProgramsAwardBountyRequest)
		req.Var("cursor", cursor)
		req.Var("secureOrderBy", variables["secureOrderBy"])
		req.Var("where", variables["where"])
		req.Var("first", 25)
		req.Var("product_area", "directory")
		req.Var("product_feature", "programs")

		var h1Response entity.HackeroneGetAllProgramsAwardBountyResponse
		if err := client.Run(context.Background(), req, &h1Response); err != nil {
			fmt.Println(err)
			return
		}
		log.Default().Println("Preparing to show response from standard view")
		for _, team := range h1Response.Teams.Edges {
			log.Default().Println(team.Node.Name)
		}

		// Get the cursor for the next page
		cursor = h1Response.Teams.PageInfo.EndCursor
		if !h1Response.Teams.PageInfo.HasNextPage {
			break
		}
	}
}

func getProgramDetailsOnScope() {
	log.Default().Println("Preparing to get program details...")
	// Get program details
	client := graphql.NewClient(entity.HackeroneGraphqlEndpoint)
	// set variables
	// Define your query variables
	variables := map[string]interface{}{
		"handle":                "shopify",
		"searchString":          "",
		"eligibleForSubmission": nil,
		"minSeverityScore":      nil,
		"eligibleForBounty":     nil,
		"asmTagIds":             []interface{}{},
		"from":                  0,
		"size":                  100,
		"sort":                  map[string]interface{}{"field": "cvss_score", "direction": "DESC"},
		"highlight":             map[string]interface{}{"fields": map[string]interface{}{"instruction": map[string]interface{}{"pre_tags": []string{"**"}, "post_tags": []string{"**"}, "number_of_fragments": 0}}},
	}

	detailCursor := ""
	for {
		// Create a new GraphQL request with the query and variables
		req := graphql.NewRequest(entity.HackeroneGetProgramDetailsRequest)
		req.Var("handle", variables["handle"])
		req.Var("searchString", variables["searchString"])
		req.Var("eligibleForSubmission", variables["eligibleForSubmission"])
		req.Var("minSeverityScore", variables["minSeverityScore"])
		req.Var("eligibleForBounty", variables["eligibleForBounty"])
		req.Var("asmTagIds", variables["asmTagIds"])
		req.Var("from", variables["from"])
		req.Var("size", variables["size"])
		req.Var("sort", variables["sort"])
		req.Var("highlight", variables["highlight"])
		req.Var("cursor", detailCursor)
		var h1Response entity.HackeroneGetProgramDetailsResponse
		if err := client.Run(context.Background(), req, &h1Response); err != nil {
			fmt.Println(err)
			return
		}
		log.Default().Println("Preparing to show response for the details of the program")
		log.Default().Println(h1Response.Team.Typename)
		log.Default().Println(h1Response.Team.ID)
		log.Default().Println(len(h1Response.Team.StructuredScopesSearch.Nodes))
		for _, scope := range h1Response.Team.StructuredScopesSearch.Nodes {
			log.Default().Println(scope.DisplayName)
		}
		detailCursor = h1Response.Team.StructuredScopesSearch.PageInfo.EndCursor
		if !h1Response.Team.StructuredScopesSearch.PageInfo.HasNextPage {
			break
		}
	}
}
