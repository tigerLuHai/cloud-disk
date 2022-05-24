package file

import (
	"cloud-disk/internal/logic/file"
	"cloud-disk/models"
	"cloud-disk/result"
	"cloud-disk/utils"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		files, header, err := r.FormFile("file")
		if err != nil {
			return
		}
		//计算hash
		bytes := make([]byte, header.Size)
		_, err = files.Read(bytes)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(bytes))
		//定义对象
		//查询是否存在hash值,如果存在直接返回
		get, err := models.RepositoryPool{}.GetHashByRepositoryPool(hash, svcCtx.Engine)
		if err != nil {
			return
		}
		if get != nil {
			m := make(map[string]string)
			m["identity"] = get.Identity
			m["ext"] = get.Ext
			m["name"] = get.Name
			httpx.OkJson(w, &types.UploadFileResponse{
				Result: result.OK("", m),
			})
			return
		}
		//不存在进行存储
		upload, err := utils.UploadFileToMinio(r)
		if err != nil {
			return
		}
		req.Name = header.Filename
		req.Ext = path.Ext(header.Filename)
		req.Size = header.Size
		req.Hash = hash
		req.Path = upload
		l := file.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
