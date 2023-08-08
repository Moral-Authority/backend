package external_apis

// Import Cloudinary and other necessary libraries
//===================
import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"os"
)

// https://cloudinary.com/documentation/responsive_client_side_js  = responsive images
// resp, err := cld.Upload.Upload(ctx, "https://www.example.com/sample.jpg", uploader.UploadParams{})
// https://cloudinary.com/documentation/image_upload_api_reference
// seth the public ID
// resp, err := cld.Upload.Upload(ctx, "sample.jpg", uploader.UploadParams{
// PublicID: "sample_id"})
func credentials() (*cloudinary.Cloudinary, context.Context) {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cloudinaryURl := os.Getenv("CLOUDINARY_URL")
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	cld.Config.URL.Domain = cloudinaryURl
	cld.Config.Cloud.CloudName = "MoralAuthority"
	ctx := context.Background()
	return cld, ctx
}

func uploadImage(cld *cloudinary.Cloudinary, ctx context.Context) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, "https://cloudinary-devs.github.io/cld-docs-assets/assets/images/butterfly.jpeg", uploader.UploadParams{
		PublicID:       "quickstart_butterfly",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
		//PublicID                       string
		//PublicIDPrefix                 string
		//PublicIDs                      api.CldAPIArray
		//UseFilename                    *bool
		//UniqueFilename                 *bool
		//UseFilenameAsDisplayName       *bool
		//FilenameOverride               string
		//DisplayName                    string
		//UniqueDisplayName              *bool
		//Folder                         string
		//AssetFolder                    string
		//UseAssetFolderAsPublicIDPrefix *bool
		//Overwrite                      *bool
		//ResourceType                   string
		//Type                           api.DeliveryType
		//Tags                           api.CldAPIArray
		//Metadata                       api.Metadata
		//Format                         string
		//AllowedFormats                 api.CldAPIArray
		//Eager                          string
		//ResponsiveBreakpoints          ResponsiveBreakpointsParams
		//Eval                           string
		//Unsigned                       *bool
		//Proxy                          string
		//Headers                        string
		//NotificationURL                string        // we will want this eventually but don't need it now
		//ImageMetadata                  *bool
		//Exif                           *bool
		//Colors                         *bool
		//Backup                         *boola
		//Invalidate                     *bool
		//DiscardOriginalFilename        *bool
		//UploadPreset                   string
		//RawConvert                     string
		//Categorization                 string
		//VisualSearch                   *bool
		//AutoTagging                    float64
		//QualityAnalysis                *bool    // we will want this eventually
		//AccessibilityAnalysis          *bool    // we will want this eventually but don't need it now
	})

	if err != nil {
		fmt.Println("error")
	}

	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")
}

func getAssetInfo(cld *cloudinary.Cloudinary, ctx context.Context) {
	// Get and use details of the image
	// ==============================
	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "quickstart_butterfly"})
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("****3. Get and use details of the image****\nDetailed response:\n", resp, "\n")

	// Assign tags to the uploaded image based on its width. Save the response to the update in the variable 'update_resp'.
	if resp.Width > 900 {
		update_resp, err := cld.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
			PublicID: "quickstart_butterfly",
			Tags:     []string{"large"}})
		if err != nil {
			fmt.Println("error")
		} else {
			// Log the new tag to the console.
			fmt.Println("New tag: ", update_resp.Tags, "\n")
		}
	} else {
		update_resp, err := cld.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
			PublicID: "quickstart_butterfly",
			Tags:     []string{"small"}})
		if err != nil {
			fmt.Println("error")
		} else {
			// Log the new tag to the console.
			fmt.Println("New tag: ", update_resp.Tags, "\n")
		}
	}

}
