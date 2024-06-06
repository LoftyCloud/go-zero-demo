package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"
	"go-zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增用户
func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	// todo: add your logic here and delete this line
	// 当返回的err不为nil时会执行回滚，取消此次事务
	// 事务中执行的db操作通过session使用一个会话

	// user := model.User{}
	// user.Mobile = req.Mobile
	// user.Name = req.NickName
	// fmt.Println(user)
	// if _, err := l.svcCtx.UserModel.Insert(l.ctx, &user); err != nil {
	// 	return nil, errors.New("新建失败")
	// }
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		user := &model.User{}
		user.Mobile = req.Mobile
		user.Name = req.NickName
		fmt.Println(user)

		// 添加user
		dbResult, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
		// _, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
		if err != nil {
			return err
		}

		// 添加userdata
		userId, _ := dbResult.LastInsertId()
		userData := &model.UserData{}
		userData.Id = userId
		userData.Data = "Nothing is here."
		// 日期格式0000-00-00导致mysql报错1292(22007)
		userData.CreateTime = time.Now()
		userData.UpdataTime = time.Now()

		if _, err := l.svcCtx.UserDataModel.TransInsert(ctx, session, userData); err != nil {
			return err
		}

		return nil
	}); err != nil {
		// 当返回的err不为nil时，给出提示
		return nil, errors.New("创建用户失败TAT")
	}

	// 创建用户成功，返回信息
	return &types.UserCreateResp{
		Flag: true,
	}, nil
}
