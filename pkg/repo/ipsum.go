package repo

import (
	"context"
	"maan/pkg/connections/database"
	"maan/pkg/model"
)

type IpSumRepo interface {
	SaveHackerIp(tx database.DbConn, ctx context.Context, hackerIp *model.IpSum) error
	FindIpSumByIp(tx database.DbConn, ctx context.Context, ip string) (ipSum *model.IpSum, err error)
}
