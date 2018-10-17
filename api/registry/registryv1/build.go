package registryv1

import (
	"github.com/jenkins-x/jx/pkg/builds"
	"github.com/IBM-Cloud/bluemix-go/client"
)

const (
	accountIDHeader = "Account"
)

type BuildTargetHeader struct {
	AccountID string
}

//ToMap ...
func (c BuildTargetHeader) ToMap() map[string]string {
	m := make(map[string]string, 1)
	m[accountIDHeader] = c.AccountID
	return m
}

// Progressdetail
type Progressdetail struct {
	Current string `json:"current,omitempty"`
	Total   string `json:"total,omitempty"`
}

type ImageBuildRequest struct {
	/*T
	  The full name for the image that you want to build, including the registry URL and namespace.
	*/
	T string
	/*F
	  Specify the location of the Dockerfile relative to the build context. If not specified, the default is 'PATH/Dockerfile', where PATH is the root of the build context.
	*/
	F string
	/*Buildargs
	  A JSON key-value structure that contains build arguments. The value of the build arguments are available as environment variables when you specify an `ARG` line which matches the key in your Dockerfile.
	*/
	Buildargs string
	/*Nocache
	  If set to true, cached image layers from previous builds are not used in this build. Use this option if you expect the result of commands that run in the build to change.
	*/
	Nocache bool
	/*Pull
	  If set to true, the base image is pulled even if an image with a matching tag already exists on the build host. The base image is specified by using the FROM keyword in your Dockerfile. Use this option to update the version of the base image on the build host.
	*/
	Pull bool
	/*Quiet
	  If set to true, build output is suppressed unless an error occurs.
	*/
	Quiet bool
	/*Squash
	  If set to true, the filesystem of the built image is reduced to one layer before it is pushed to the registry. Use this option if the number of layers in your image is close to the maximum for your storage driver.
	*/
	Squash bool

}

func DefaultImageBuildRequest() (*ImageBuildRequest)
{
	return &ImageBuildRequest{
		T: "",
		F: "",
		Buildargs: "",
		Nocache: false,
		Pull: false,
		Quiet: false,
		Squash: false,
	}
}
//ImageBuildResponse
type ImageBuildResponse struct {
	ID             string          `json:"id,omitempty"`
	ProgressDetail *Progressdetail `json:"progressDetail,omitempty"`
	Status         string          `json:"status,omitempty"`
	Stream         string          `json:"stream,omitempty"`
}

//Subnets interface
type Builds interface {
	ImageBuild(params ImageBuildRequest, target BuildTargetHeader) (ImageBuildResponse, error)
}

type builds struct {
	client *client.Client
}

func newBuildAPI(c *client.Client) Builds {
	return &builds{
		client: c,
	}
}
rawURL := fmt.Sprintf("/v1/clusters/%s?showResources=true", name)

func addToRequestHeader(h interface{}, r *rest.Request) {
	switch v := h.(type) {
	case map[string]string:
		for key, value := range v {
			r.Set(key, value)
		}
	}
}

//Create ...
func (r *builds) Create(params ImageBuildRequest, target BuildTargetHeader) (ImageBuildResponse, error) {
	var imageBuild ImageBuildResponse
	req := rest.PostRequest(helpers.GetFullURL(*r.client.Config.Endpoint, "/api/v1/builds"))
		.Query("t", params.T)
		.Query("f", params.F)
		.Query("buildarg", params.Buildargs)
		.Query("nocache", params.Nocache)
		.Query("pull", params.Pull)
		.Query("quiet", params.Quiet)
		.Query("squash", params.Squash)

	for key, value := range target.ToMap() {
		req.Set(key, value)
	}

	_, err := r.client.SendRequest(req, &imageBuild)
	if err != nil {
		return nil, err
	}
	return imageBuild, err
}