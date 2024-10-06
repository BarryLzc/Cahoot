package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var MySQLDsn string
var SoftDelete bool
var TableMap = make(map[string]string) //key:tableName value:modelName

func init() {

	MySQLDsn = "root:123@tcp(127.0.0.1:3306)/cahoot?charset=utf8mb4&parseTime=True"
	SoftDelete = false
	TableMap["users"] = "Users"
	TableMap["game"] = "Game"
}

func main() {
	db, err := gorm.Open(mysql.Open(MySQLDsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	// 生成实例
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/infrastructure/database/orm",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true,
		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false,
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true,
	})
	// 设置目标 db
	g.UseDB(db)

	// 自动更新时间戳
	autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	// soft deleted: `deleted`
	softDeleteFieldTag := gen.FieldGORMTag("deleted", "column:deleted;type:int unsigned;softDelete:flag")
	softDeleteFieldType := gen.FieldType("deleted", "soft_delete.DeletedAt")
	fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField}
	if SoftDelete {
		fieldOpts = append(fieldOpts, softDeleteFieldTag, softDeleteFieldType)
	}
	// out put path: model
	for tableName, modelName := range TableMap {
		g.ApplyBasic(g.GenerateModelAs(tableName, modelName, fieldOpts...))

	}
	g.Execute()
}
