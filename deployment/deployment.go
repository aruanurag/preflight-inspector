package deployments

import (
	"encoding/json"

	preflightinspector "github.com/aruanurag/preflight-inspector"
	v1 "k8s.io/api/apps/v1"
)

// NewValidationHook creates a new instance of deployment validation hook
func NewValidationHook() preflightinspector.Hook {
	return preflightinspector.Hook{
		Create: validateCreate(),
		Delete: validateDelete(),
	}
}

func parseDeployment(object []byte) (*v1.Deployment, error) {
	var dp v1.Deployment
	if err := json.Unmarshal(object, &dp); err != nil {
		return nil, err
	}

	return &dp, nil
}
