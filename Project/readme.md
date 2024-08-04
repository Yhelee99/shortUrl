# 项目架构图

![架构图](架构图.png)

## 搭建项目骨架

1. 建库建表

   1.1 新建发号器表

   1.2 新建长短链接映射表

2. 搭建go-Zero框架的骨架

	2.1 编写api文件

	2.2 根据api文件生成go代码

3. 根据数据表生成model层代码

	```bash
	goctl model mysql datasource -url="user:password@tcp(addr:port)/database" -table="table" -dir="./model"
	```

4. 同步依赖