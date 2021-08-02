package middleware

import (
	"ihtest/library"

	"github.com/gin-gonic/gin"
)

func ParserDoctor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := handler(ctx); nil != err {
			library.NewResponse(ctx).Error(err)
			ctx.Abort()
		}

		ctx.Next()
	}
}

type userIdParam struct {
	UserId string `json:"userId" binding:"required"`
}

func handler(ctx *gin.Context) error {
	thirdUser := new(userIdParam)
	if err := library.NewRequest(ctx).Bind(thirdUser); err != nil {
		return err
	}
	//
	//mCtx := api_http.ParserMedContext(ctx)
	//userInfo, err := userPb.NewUserClient(global.UserClientConn).
	//	GetUserInfoByThirdParty(api_http.CreateGRPCContext(mCtx), &userPb.GetUserInfoByThirdPartyReq{
	//		UserPlatform: "co",
	//		Platform:     int32(GetPlatform(ctx)),
	//		OpenId:       thirdUser.UserId,
	//	})
	//if err != nil {
	//	return err
	//}
	//if userInfo.UserInfo == nil {
	//	return xerror.New("fail to get user info by third user id")
	//}
	//
	//doctorInfo, err := doctorPb.NewUserClient(global.DoctorClientConn).
	//	GetDoctorInfoByUserId(api_http.CreateGRPCContext(mCtx), &doctorPb.DoctorReq{
	//		UserId: userInfo.UserInfo.UserId,
	//	})
	//if err != nil {
	//	return err
	//}
	//if doctorInfo.DoctorInfo == nil {
	//	return xerror.New("fail to get doctor info")
	//}
	//
	//doctorEntity := new(api_http.DoctorInfo)
	//if err := utils.NewCopy().SetSource(doctorInfo.DoctorInfo).CopyF(doctorEntity); err != nil {
	//	return err
	//}
	//
	//if v, ok := ctx.Get(global.HttpCustomContextKey); ok {
	//	medCtx := v.(api_http.MedContext)
	//	medCtx.Set("doctorInfo", doctorEntity)
	//	ctx.Set("params", thirdUser)
	//	medCtx.SetUserId(userInfo.UserInfo.UserId)
	//	medCtx.Storage(ctx)
	//}

	return nil
}
