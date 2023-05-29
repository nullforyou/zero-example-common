package tool

import (
	"context"
	"encoding/json"
	"go-common/utils/xerr"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) (uid int64, err error) {
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		}
	}
	if uid > 0 {
		return uid, nil
	} else {
		return 0, xerr.NewBusinessError(xerr.SetCode("UserNotExists"), xerr.SetMsg("token有效载荷错误"))
	}
}