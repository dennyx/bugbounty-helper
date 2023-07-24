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
	getAllProgramsAwardBounty()
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
	// Create a new GraphQL request with the query and variables
	req := graphql.NewRequest(entity.HackeroneGetAllProgramsAwardBountyRequest)
	req.Var("cursor", variables["cursor"])
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
}
