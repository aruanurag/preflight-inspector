package deployments

import (
	preflightinspector "github.com/aruanurag/preflight-inspector"
	"k8s.io/api/admission/v1beta1"
)

func validateCreate() preflightinspector.AdmitFunc {
	return func(r *v1beta1.AdmissionRequest) (*preflightinspector.Result, error) {
		dp, err := parseDeployment(r.Object.Raw)
		if err != nil {
			return &preflightinspector.Result{Msg: err.Error()}, nil
		}

		if dp.Namespace == "special" {
			return &preflightinspector.Result{Msg: "You cannot create a deployment in `special` namespace."}, nil
		}

		return &preflightinspector.Result{Allowed: true}, nil
	}
}
