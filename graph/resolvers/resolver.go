package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
    "github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Resolver struct{
	AlgoliaClient *search.Client
}
