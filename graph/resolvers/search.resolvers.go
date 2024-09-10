package resolvers

import (
	"context"
	"fmt"

	"github.com/Moral-Authority/backend/graph/model"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, input string) ([]*model.Product, error) {
	// Use the Algolia client from the Resolver struct
	index := r.AlgoliaClient.InitIndex("products_index")

	// Perform the search in Algolia
	res, err := index.Search(input)
	if err != nil {
		return nil, fmt.Errorf("failed to search in Algolia: %w", err)
	}

	// Process the search results and map them to the Product model
	var products []*model.Product
	for _, hit := range res.Hits {
		// Access objectID and other fields correctly from the map
		objectID, ok := hit["objectID"].(string)
		if !ok {
			return nil, fmt.Errorf("objectID not found in Algolia hit")
		}

		title, _ := hit["title"].(string)
		companyName, _ := hit["company_name"].(string)
		price, _ := hit["price"].(float64)
		priceString := fmt.Sprintf("%.2f", price)

		product := &model.Product{
			ID:    objectID,
			Title: title,
			Company: &model.Company{
				Name: companyName,
			},
			PurchaseInfo: []*model.PurchaseInfo{
				{
					Price: &priceString,
				},
			},
		}

		products = append(products, product)
	}

	// Return the list of products
	return products, nil
}
