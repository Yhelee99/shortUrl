package svc

import (
	sequence "Project/Sequence"
	"Project/internal/config"
	"Project/model"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	ShortUrlDb model.ShortUrlMapModel

	Sequence  sequence.Sequence
	BlackList map[string]struct{} // 定义一个map,k为string,v为空结构体(目的：不占空间)
	Filter    bloom.Filter
	Conn      sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化Mysql连接
	conn := sqlx.NewMysql(c.ShortUrlDb.DSN)

	// 把黑名单数组存到图中
	m := make(map[string]struct{}, len(c.BlackList))
	for _, v := range c.BlackList {
		m[v] = struct{}{}
	}

	// 初始化布隆过滤器
	rds := redis.New(c.CacheRedis[0].Host)
	filter := bloom.New(rds, "bloom filter", 20*(1<<20)) //byte 计算公式，20*数据量

	// 启动时，加载数据到布隆过滤器
	// loadDataToBloomFliter(conn, *filter)

	return &ServiceContext{
		Config:     c,
		ShortUrlDb: model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:   sequence.NewSeqMysql(c.SequenceDb.DSN),
		BlackList:  m,
		Filter:     *filter,
	}
}

// LoadDataToBloomFliter 加载库里已有的短链到布隆过滤器中

// 注意导入的是这个bloom
// import github.com/bits-and-blooms/bloom/v3
//
// func LoadDataToBloomFliter(conn sqlx.SqlConn, filter bloom.Filter) {
//
// 	// 验证初始化
// 	if conn == nil {
// 		logx.Errorw("mysql or filter uninit.")
// 		return
// 	}
//
// 	// 查询数据总量
// 	total := 0
// 	sqlStr := "select count(*) from short_url_map"
// 	conn.QueryRow(&total, sqlStr)
// 	if total == 0 {
// 		logx.Info("no data need load")
// 		return
// 	}
//
// 	// 分页
// 	pageSize := 2
// 	pageNum := 0
// 	if total%pageSize == 0 {
// 		pageNum = total / pageSize
// 	} else {
// 		pageNum = total/pageSize + 1
// 	}
//
// 	// 查数据
// 	for i := 0; i < pageNum; i++ {
// 		//分页查询
// 		offset := i * pageSize
// 		offset1 := offset + pageSize
// 		surls := []string{}
// 		if err := conn.QueryRow(&surls, "select surl from short_url_map where is_del = 0 limit ?,?", offset, offset1); err != nil {
// 			logx.Errorw("loadDataToBloomFliter query data failed.", logx.Field("err", err))
// 			return
// 		}
// 		//加载进布隆过滤器
// 		for _, v := range surls {
// 			filter.AddString(surl)
// 		}
// 	}
// 	logx.Info("load all data into bloom filter successful")
// }
