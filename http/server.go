package http

import (
	"net/http"

	deployments "github.com/aruanurag/preflight-inspector/deployment"
	"github.com/aruanurag/preflight-inspector/pods"
)

func NewServer(port string) *http.Server {

	// Instances hooks
	podsValidation := pods.NewValidationHook()
	podsMutation := pods.NewMutationHook()
	deploymentValidation := deployments.NewValidationHook()

	// Routers
	ah := newAdmissionHandler()
	mux := http.NewServeMux()
	mux.Handle("/healthz", healthz())
	mux.Handle("/validate/pods", ah.Serve(podsValidation))
	mux.Handle("/mutate/pods", ah.Serve(podsMutation))
	mux.Handle("/validate/deployments", ah.Serve(deploymentValidation))

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
