package entity

/*
	defaultEntities := map[string]any{
		"account": map[string]any{
			"implementations": map[string]any{
				"transport": map[string][]string{
					"acccountapi": []string{"account.go", "type.go"},
				},
				"usecase": map[string][]string{
					"accountuc": []string{"account.go", "type.go"},
				},
				"repository": map[string][]string{
					"accountrepo": []string{"account.go", "type.go"},
				},
			},
			"entities": map[string]any{
				"account": []string{"account.go", "api.go", "usecase.go", "repository.go", "type.go"},
			},
			"database": []string{"migrations", "seeder"},
		},
	}
*/

type (
	DirectoryElementsSchema []string

	ProjectSchema struct {
		Implementations ProjectImplementationsSchema       `json:"implementations"`
		Entities        map[string]DirectoryElementsSchema `json:"entities"`
		Database        DirectoryElementsSchema            `json:"database"`
	}

	ProjectImplementationsSchema struct {
		Transport  map[string]DirectoryElementsSchema `json:"transport"`
		Usecase    map[string]DirectoryElementsSchema `json:"usecase"`
		Repository map[string]DirectoryElementsSchema `json:"repository"`
	}
)
