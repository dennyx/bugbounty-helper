package entity

var (
	// Hackerone Graphql Endpoint
	HackeroneGraphqlEndpoint = "https://hackerone.com/graphql"
	// Get all the bug bounty programs that award bounty
	HackeroneGetAllProgramsAwardBountyRequest = `
	query DirectoryQuery($cursor: String, $secureOrderBy: FiltersTeamFilterOrder, $where: FiltersTeamFilterInput) {
		me {
			id
			edit_unclaimed_profiles
			__typename
		}
		teams(first: 25, after: $cursor, secure_order_by: $secureOrderBy, where: $where) {
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
	}`
	// Get details on a specific bug bounty program
	HackeroneGetProgramDetailsRequest = `
		query PolicySearchStructuredScopesQuery(
			$handle: String!, 
			$searchString: String, 
			$eligibleForSubmission: Boolean, 
			$eligibleForBounty: Boolean, 
			$minSeverityScore: SeverityRatingEnum, 
			$asmTagIds: [Int], 
			$from: Int, 
			$size: Int,
			$sort: SortInput, 
			$highlight: StructuredScopeHighlightInput) {
			  team(handle: $handle) {
				id
				structured_scopes_search(
					search_string: $searchString
					eligible_for_submission: $eligibleForSubmission
					eligible_for_bounty: $eligibleForBounty
					min_severity_score: $minSeverityScore
					asm_tag_ids: $asmTagIds
					from: $from
					size: $size
					sort: $sort
					highlight: $highlight) {
						nodes {
							... on StructuredScopeDocument {
										id
										highlight
										...PolicyScopeStructuredScopeDocument
										__typename
									}
							__typename
						}
						pageInfo {
							startCursor
							hasPreviousPage
							endCursor
							hasNextPage
							__typename
						}
						total_count
						__typename
				}
				__typename
			  }
		}
		fragment PolicyScopeStructuredScopeDocument on StructuredScopeDocument {
			id
			identifier
			display_name
			instruction
			cvss_score
			eligible_for_bounty
			eligible_for_submission
			asm_system_tags
			created_at
			__typename
		}
	`
)
