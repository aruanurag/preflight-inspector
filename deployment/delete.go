package deployments

import (
	preflightinspector "github.com/aruanurag/preflight-inspector"
	"k8s.io/api/admission/v1beta1"
)

func validateDelete() preflightinspector.AdmitFunc {
	return func(r *v1beta1.AdmissionRequest) (*preflightinspector.Result, error) {
		dp, err := parseDeployment(r.OldObject.Raw)
		if err != nil {
			return &preflightinspector.Result{Msg: err.Error()}, nil
		}

		if dp.Namespace == "special-system" && dp.Annotations["skip"] == "false" {
			return &preflightinspector.Result{Msg: "You cannot remove a deployment from `special-system` namespace."}, nil
		}

		return &preflightinspector.Result{Allowed: true}, nil
	}
}
