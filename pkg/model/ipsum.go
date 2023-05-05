package model

// 表名称

const TableNameIpSum = "ipsum"

type IpSum struct {
	HackerIp string `gorm:"column:hacker_ip;type:varchar(255);not null;primaryKey;comment:恶意ip" json:"hacker_ip"` // 恶意ip
}

func (*IpSum) TableName() string {
	return TableNameIpSum
}
