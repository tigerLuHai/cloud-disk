package handler

import (
	"net/http"

	"cloud-disk/internal/logic"
	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareBasicDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareBasicDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShareBasicDetailLogic(r.Context(), svcCtx)
		resp, err := l.ShareBasicDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
