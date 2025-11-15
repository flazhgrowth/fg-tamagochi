package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/response"
)

func (r *RouterImpl) handleDocs(method string, path string, docs RouterDocs) {
	endpoint := path
	if r.prevPath != "" {
		endpoint = fmt.Sprintf("%s%s", r.prevPath, path)
	}
	ops, err := r.openapireflector.NewOperationContext(method, endpoint)
	if err != nil {
		apierrors.LogPath("handleDocs").LogError(context.Background(), "failed on instantiate NewOperationContext", err)
	}
	if path != http.MethodDelete && docs.Request != nil {
		ops.AddReqStructure(docs.Request)
	}
	if docs.Response != nil {
		ops.AddRespStructure(response.BaseResponse{
			Data: docs.Response,
		})
	}
	if !docs.Security.IsPublic() {
		ops.AddSecurity(string(docs.Security))
	}
	if docs.IsDeprecated {
		ops.SetIsDeprecated(true)
	}
	ops.SetTags(docs.Tags)
	ops.SetSummary(docs.Summary)
	ops.SetDescription(docs.Description)

	if err = r.openapireflector.AddOperation(ops); err != nil {
		apierrors.LogPath("openapireflector.AddOperation").LogError(context.Background(), "failed on adding operation", err)
	}
}
