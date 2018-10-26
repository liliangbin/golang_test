package main

func m()  {

}
/*
import (
	"bytes"
	"database/sql"
	"fmt"
	"lzy/framework/util"
	_ "mysql"
	"reflect"
	"strconv"
	"strings"
)

//多功能业务结构体
type MultiFunction struct {
	Model       interface{}               //业务主体
	ResultSet   map[int]map[string]string //检索结果集
	Rule        map[string]string         //检索规则,key必须是同属性相同，区分大小写
	Field       map[string]string         //可自定义检索字段，字段必须对应于数据库字段
	KV          map[string]string         //update、insert、select、均可使用它进行存放数据
	cv          []interface{}             //检索条件值，放入该切片
	Ids         []string                  //如果该字段设置值后，框架将会字段转为以id为检索条件进行检索
	IsQueryAll  bool                      //如果为true，框架检索时将放弃条件检索，直接检索全表数据
	pageSize    int                       //每页显示多少
	nowPage     int                       //检索第几页
	total       int                       //总行数
	totalPage   int                       //总页数
	IsPagin     bool                      //是否分页
	OrderBy     string                    //排序字段
	OrderByType string                    //排序类型
	tableN      string                    //表名
	sql         string                    //拼接后的sql
	IsSort      bool                      //是否排序
	tag         int64                     //标记
	IsDistinct  bool                      //是否去重
	Id          string                    //
	mysql       bool                      //是否使用自定义sql
	DbAddr      string                    //链接数据库地址
	DbType      string                    //数据库名
	beginRow    int                       //开始行
	Num         int64                     //影响结果数
}

//-------------------- util_new_struct -----------------------

func NewMultiFunction(o interface{}) *MultiFunction {

	return &MultiFunction{
		Model:       o,                       //主数据
		Rule:        make(map[string]string), //初始化防止空指针异常
		Field:       make(map[string]string), //初始化防止空指针异常
		KV:          make(map[string]string), //初始化防止空指针异常
		pageSize:    10,                      //默认每页十条
		nowPage:     1,                       //默认第一页
		IsPagin:     true,                    //默认分页
		IsDistinct:  false,                   //默认不去重
		OrderBy:     "Id",                    //默认排序字段
		OrderByType: "desc",                  //默认排序类型
		IsQueryAll:  false,                   //默认条件检索
		IsSort:      true,                    //默认排序
		total:       -1,                      //默认总行数
		Id:          "Id",                    //in查找时使用该值
		mysql:       false,                   //默认使用框架生成
		tag:         0,                       //0
		beginRow:    -1,                      //0
		tableN:      reflect.ValueOf(o).Elem().FieldByName("tableN").String(),
		Num:         0,
	}
}

func (m *MultiFunction) SetSql(sql string) {
	m.mysql = true
	m.sql = sql
}

func (m *MultiFunction) GetPageSize() int {
	return m.pageSize
}

func (m *MultiFunction) SetPageSize(pageSize int) {
	m.pageSize = pageSize
}

func (m *MultiFunction) SetNowPage(n int) {
	if n < 0 {
		m.nowPage = 0
	}
	if n > m.GetTotalPage() {
		m.nowPage = m.totalPage
	}
	m.nowPage = n
}
func (m *MultiFunction) GetNowPage() int {
	return m.nowPage
}

func (m *MultiFunction) GetTotal() int {
	if m.total == -1 {
		ex := NewExecute()
		ex.Count(m)
	}
	return m.total
}

func (m *MultiFunction) GetTotalPage() int {
	m.total = m.GetTotal()

	if m.total%m.pageSize == 0 {
		m.totalPage = m.total / m.pageSize
		return m.totalPage
	} else {
		m.totalPage = m.total/m.pageSize + 1
		return m.totalPage
	}
}

func (m *MultiFunction) GetBeginRow() int {
	if m.beginRow < 0 {
		return m.pageSize * (m.nowPage - 1)
	}
	return m.beginRow
}

func (m *MultiFunction) SetBeginRow(n int) {
	m.beginRow = n
}

type Execute struct{}

func NewExecute() *Execute {
	return &Execute{}
}

//---------------------------
//操作数据库全局变量

//增
func (ex *Execute) Insert(mf *MultiFunction) *util.Message {

	var db *sql.DB
	var res sql.Result
	var err error
	var tx *sql.Tx

	msg := util.NewMessage()

	db, err = GetDBConnect(mf)
	defer db.Close()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	tx, err = db.Begin()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	if !mf.mysql {
		mf.tag = 0
		GetInserSql(mf)
	}

	res, err = tx.Exec(mf.sql, mf.cv[0:]...)
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}
	tx.Commit()

	mf.Num, _ = res.RowsAffected()
	return msg
}

//删
func (ex *Execute) Delete(mf *MultiFunction) *util.Message {

	var db *sql.DB
	var res sql.Result
	var err error
	var tx *sql.Tx

	msg := util.NewMessage()

	db, err = GetDBConnect(mf)
	defer db.Close()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	if !mf.mysql {
		mf.tag = 0
		GetDeleteSql(mf)
	}
	tx, err = db.Begin()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	fmt.Println("finally sql---->" + mf.sql)
	res, err = tx.Exec(mf.sql, mf.cv[0:]...)
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}
	tx.Commit()
	mf.Num, _ = res.RowsAffected()

	return msg
}

//改
func (ex *Execute) Update(mf *MultiFunction) *util.Message {

	var db *sql.DB
	var res sql.Result
	var err error
	var tx *sql.Tx

	msg := util.NewMessage()

	db, err = GetDBConnect(mf)
	defer db.Close()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	tx, err = db.Begin()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	if !mf.mysql {
		mf.tag = 0
		GetUpdateSql(mf)
	}

	res, err = tx.Exec(mf.sql, mf.cv[0:]...)
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}
	tx.Commit()

	mf.Num, _ = res.RowsAffected()

	return msg
}

//查
func (ex *Execute) QueryAllOrByCondition(mf *MultiFunction) *util.Message {

	var db *sql.DB
	var rls *sql.Rows
	var err error
	var tx *sql.Tx

	msg := util.NewMessage()

	db, err = GetDBConnect(mf)
	defer db.Close()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	tx, err = db.Begin()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	if !mf.mysql {
		mf.tag = 0
		GetQueryAllOrByConditionSql(mf)
	}

	rls, err = tx.Query(mf.sql, mf.cv[0:]...)
	defer rls.Close()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}
	handleRls(rls, mf, msg)

	return msg
}

//查数量
func (ex *Execute) Count(mf *MultiFunction) *util.Message {

	var db *sql.DB
	var rls *sql.Rows
	var err error
	var tx *sql.Tx

	msg := util.NewMessage()
	db, err = GetDBConnect(mf)
	defer db.Close()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	tx, err = db.Begin()
	defer tx.Commit()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	if !mf.mysql {
		mf.tag = 0
		GetCountSql(mf)
	}

	rls, err = tx.Query(mf.sql, mf.cv[0:]...)
	defer rls.Close()
	util.CheckErr(err, msg)
	if err != nil {
		return msg
	}

	if rls.Next() {
		rls.Scan(&mf.total)
	}
	return msg
}

func handleRls(rls *sql.Rows, mf *MultiFunction, msg *util.Message) {
	//defer rls.Close()
	column, _ := rls.Columns()                //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //将values的第二维数组的地址复制到scans中，操作scans等同操作values的值
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for rls.Next() { //循环，让游标往下移动
		if err := rls.Scan(scans...); err != nil { //当前行的每一列的值赋值到scans中,也就是每行都放在values值里
			util.CheckErr(err, msg)
			return
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values第二维数组中里面，现在把它挪到row里
			key := column[k]     //获得第K列字段name
			row[key] = string(v) //转为string
		}
		results[i] = row //装入结果集中
		i++
	}
	mf.ResultSet = results
	return
}

//---------------------- dao_interface ----------------------------
//通用CRUD接口
type CRUD interface {
	Insert(*MultiFunction) *util.Message
	QueryAllOrByCondition(*MultiFunction) *util.Message
	Delete(*MultiFunction) *util.Message
	Update(*MultiFunction) *util.Message
}

//--------------------- dao_sql_factory ----------------------------
func GetCountSql(mf *MultiFunction) {
	sql := bytes.NewBufferString("select count(*) as total from ")
	sql.WriteString(mf.tableN)

	l := reflect.ValueOf(mf.Model).Elem().NumField()
	l2 := len(mf.Ids)
	mf.cv = make([]interface{}, l+l2+2)

	sql.WriteString(appendQueryCondition(mf))

	mf.cv = mf.cv[0:mf.tag]
	fmt.Println("select coun sql ---> ", sql.String())
	fmt.Println("query values --------> ", mf.cv)
	mf.sql = sql.String()
	return
}

//Delete Sql Factory
func GetDeleteSql(mf *MultiFunction) {
	if !util.IsNotEmpty(mf.tableN) {
		fmt.Println("mf.tableN is empty...")
		return
	}
	sql := bytes.NewBufferString("delete from ")
	sql.WriteString(mf.tableN)
	l := reflect.ValueOf(mf.Model).Elem().NumField()
	mf.cv = make([]interface{}, l)
	sql.WriteString(appendQueryCondition(mf))

	mf.cv = mf.cv[0:mf.tag]
	fmt.Println("delete sql ---> ", sql.String())
	mf.sql = sql.String()
	fmt.Println("delete value ---> ", mf.cv)
	return
}

//Query Sql Factory
func GetQueryAllOrByConditionSql(mf *MultiFunction) {
	if !util.IsNotEmpty(mf.tableN) {
		panic("mf.tableN is empty...")
		return
	}
	sql := bytes.NewBufferString("select ")
	if mf.IsDistinct { //去重
		sql.WriteString(" distinct ")
	}
	if len(mf.Field) <= 0 { //如果没有自定义检索字段
		sql.WriteString(appendFieldSql(mf.Model)) //检索所有字段
	} else {
		for _, v := range mf.Field { //自定义字段
			sql.WriteString(v)
			sql.WriteString(",")
		}
	}
	//截取最后一个逗号
	sql2 := sql.String()
	l := len(sql2)
	sql2 = sql2[0 : l-1]

	sql = bytes.NewBufferString(sql2)
	sql.WriteString(" from ")
	sql.WriteString(mf.tableN)

	l2 := reflect.ValueOf(mf.Model).Elem().NumField()
	if !mf.IsQueryAll { //如果不是检索全表则进行条件拼接
		l3 := len(mf.Ids)
		if l3 > 0 {
			mf.cv = make([]interface{}, l2+l3+2) //属性数量+Ids长度+分页
		} else {
			mf.cv = make([]interface{}, l2+2) //属性数量+分页
		}
		sql.WriteString(appendQueryCondition(mf))
	} else { //全表检索
		mf.cv = make([]interface{}, 2) //属性数量+分页
	}
	//如果排序
	if mf.IsSort {
		sql.WriteString(" order by ")
		sql.WriteString(mf.OrderBy)
		sql.WriteString(" " + mf.OrderByType)
	}
	//如果分页
	if mf.IsPagin {
		sql.WriteString(" limit ? ,? ")
		mf.cv[mf.tag] = mf.GetBeginRow()
		mf.tag++
		mf.cv[mf.tag] = mf.GetPageSize()
		mf.tag++

	}

	fmt.Println("query sql ---> ", sql.String())
	mf.cv = mf.cv[0:mf.tag]
	fmt.Println("query values ---> ", mf.cv)
	mf.sql = sql.String()

	return
}

//----------->>> Insert Sql Factory
func GetInserSql(mf *MultiFunction) {
	if !util.IsNotEmpty(mf.tableN) {
		fmt.Println("mf.tableN is empty...")
		return
	}
	//主sql
	sql := bytes.NewBufferString("insert into ")
	//表名
	sql.WriteString(mf.tableN)
	sql.WriteString("(")
	//插入属性占位符
	field := bytes.NewBufferString("")

	o := mf.Model
	//主结构体
	o2 := reflect.ValueOf(o).Elem() //可以取得属性值
	o3 := o2.Type()                 //可以取得属性类型和名字
	l := o2.NumField()
	mf.cv = make([]interface{}, l)

	for i := 0; i < l; i++ {
		f := o3.Field(i)
		fn := f.Name
		fm := o2.FieldByName(fn)
		if fn != "tableN" {
			if util.IsStringT(fm.Type()) { //属性为字符串类型
				fv := fm.String()
				if util.IsNotEmpty(fv) {
					sql.WriteString(fn)
					sql.WriteString(",")
					field.WriteString("?,")
					mf.cv[mf.tag] = fv
					mf.tag++
				}
			}
			if util.IsIntT(fm.Type()) { //值为整数类型
				fv := fm.Int()
				if string(fv) != "" {
					sql.WriteString(fn)
					sql.WriteString(",")
					field.WriteString("?,")
					mf.cv[mf.tag] = fv
					mf.tag++
				}
			}
		}
	}
	sql = bytes.NewBufferString(sql.String()[0 : len(sql.String())-1])
	sql.WriteString(")values(")
	sql.WriteString(field.String()[0 : len(field.String())-1])
	sql.WriteString(")")

	fmt.Println("insert sql ---> ", sql.String())
	mf.cv = mf.cv[0:mf.tag]
	fmt.Println("insert values ---> ", mf.cv)
	mf.sql = sql.String()
	return
}

//----------------->>> Update Sql Factory
func GetUpdateSql(mf *MultiFunction) {
	if !util.IsNotEmpty(mf.tableN) {
		fmt.Println("mf.tableN is empty...")
		return
	}
	sql := bytes.NewBufferString("update ")
	sql.WriteString(mf.tableN)
	sql.WriteString(" set ")

	o := mf.Model                   //主结构体
	o2 := reflect.ValueOf(o).Elem() //可以取得属性值
	//o3 := o2.Type()                 //可以取得属性类型和名字
	mf.cv = make([]interface{}, len(mf.KV)+len(mf.Ids)+len(mf.Rule))

	for _, v := range mf.KV {
		f := o2.FieldByName(v)
		if util.IsStringT(f.Type()) {
			fv := f.String()
			sql.WriteString(v)
			sql.WriteString("=?,")
			mf.cv[mf.tag] = fv
			mf.tag++
		}
		if util.IsIntT(f.Type()) {
			fv := f.Int()
			sql.WriteString(v)
			sql.WriteString("=?,")
			mf.cv[mf.tag] = fv
			mf.tag++
		}
	}

	sql2 := sql.String()
	sql2 = sql2[0 : len(sql2)-1]
	sql = bytes.NewBufferString(sql2)
	sql.WriteString(appendQueryCondition(mf))
	fmt.Println("update sql ---> ", sql.String())
	mf.cv = mf.cv[0:mf.tag]
	fmt.Println("update values ---> ", mf.cv)
	mf.sql = sql.String()
	return
}

//---------------->>> 拼接字段
func appendFieldSql(o interface{}) string {

	sql := bytes.NewBufferString("")
	o2 := reflect.ValueOf(o).Elem().Type() //对o解址后的reflect.Value
	l := o2.NumField()                     //返回reflect.Type  可从其获得结构体属性名称、类型
	for i := 0; i < l; i++ {               //遍历主struct字段
		f := o2.Field(i)
		fn := f.Name //属性名字，区分大小写
		if fn != "tableN" {
			sql.WriteString(fn)
			sql.WriteString(",")
		}
	}
	return sql.String()
}

//-------------->>>拼接条件，首先判断Ids是否有值，如果有则只检索Ids条件
//否则遍历检索规则map字段，利用反射获得值,Updata时不适用
func appendQueryCondition(mf *MultiFunction) string {

	condition := bytes.NewBufferString(" where 1=1 ")
	if len(mf.Ids) > 0 { //如果是根据ID条件检索

		condition.WriteString("and ")
		condition.WriteString(mf.Id)
		condition.WriteString(" in(")
		fmt.Println(len(mf.Ids))
		for i := 0; i < len(mf.Ids); i++ {
			condition.WriteString("?,")
			mf.cv[mf.tag] = mf.Ids[i]
			mf.tag++
		}
		//截去最后一个逗号
		temp := condition.String()
		temp1 := temp[0 : len(temp)-1]
		condition = bytes.NewBufferString(temp1)
		condition.WriteString(")")
	}
	if len(mf.Rule) > 0 {
		rule := mf.Rule    //定义的检索规则
		if len(rule) > 0 { //如果有
			o := mf.Model
			f := reflect.ValueOf(o).Elem()
			for k, v := range rule { //遍历检索规则进行追加条件
				if strings.Index(v, "@_data_@") != -1 { //日期额外处理
					v2 := strings.Split(v, "@_data_@")
					fn := f.FieldByName(k)
					fv := fn.String()
					fv2 := strings.Split(fv, "@_data_@")
					condition.WriteString(" and ")
					condition.WriteString(k)
					condition.WriteString(v2[0])
					mf.cv[mf.tag] = fv2[0]
					mf.tag++
					condition.WriteString(" and ")
					condition.WriteString(k)
					condition.WriteString(v2[1])
					mf.cv[mf.tag] = fv2[1]
					mf.tag++

				} else {
					fn := f.FieldByName(k)         //利用反射预热
					if util.IsStringT(fn.Type()) { //是否是字符串类型
						fv := fn.String()
						condition.WriteString(" and ")
						condition.WriteString(k)
						condition.WriteString(v)
						if v == LIKE {
							fv1 := "%" + fv + "%"
							mf.cv[mf.tag] = fv1
							mf.tag++

						} else {
							mf.cv[mf.tag] = fv
							mf.tag++
						}
					} else if util.IsIntT(fn.Type()) {
						fv := fn.Int()
						condition.WriteString("and ")
						condition.WriteString(k)
						condition.WriteString(v)
						if v == LIKE {
							fv1 := "%" + strconv.FormatInt(fv, 10) + "%"
							mf.cv[mf.tag] = fv1
							mf.tag++
						} else {
							mf.cv[mf.tag] = fv
							mf.tag++
						}
					}
				}
			}
		}
	}
	return condition.String()
}

//Updata条件
//func appendUpdataCondition(mf *MultiFunction) string {

//  condition := bytes.NewBufferString(" where 1=1 ")
//  if len(mf.Ids) > 0 { //如果是只根据ID条件检索
//      condition.WriteString("and ")
//      condition.WriteString(mf.Id)
//      condition.WriteString(" in(")
//      for i := 0; i < len(mf.Ids); i++ {
//          condition.WriteString("?,")
//          mf.cv[mf.tag] = mf.Ids[i]
//          mf.tag++
//      }
//      //截去最后一个逗号
//      temp := condition.String()
//      temp1 := temp[0 : len(temp)-1]
//      condition = bytes.NewBufferString(temp1)
//      condition.WriteString(")")
//  } else { //否则根据其它条件
//      kv := mf.KV      //定义的键值对
//      if len(kv) > 0 { //如果有
//          for k, v := range kv { //遍历检索规则进行追加条件
//              rv := mf.Rule[k]
//              if rv == LIKE {
//                  condition.WriteString(" and ")
//                  condition.WriteString(k)
//                  condition.WriteString(rv)
//                  fv1 := "%" + v + "%"
//                  mf.cv[mf.tag] = fv1
//                  mf.tag++

//              } else {
//                  condition.WriteString(" and ")
//                  condition.WriteString(k)
//                  condition.WriteString(rv)
//                  mf.cv[mf.tag] = v
//                  mf.tag++
//              }
//          }
//      }
//  }
//  return condition.String()
//}

//--------------------- db_connection_factory ----------------------
//获得DB连接
func GetDBConnect(mf *MultiFunction) (*sql.DB, error) {

	db, openConnectErr := sql.Open(mf.DbType, mf.DbAddr)
	if openConnectErr != nil {
		fmt.Println("open database connect failed")
		util.HandleErr(openConnectErr)
		return nil, openConnectErr
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Printf("ping %s failed", mf.DbAddr)
		util.HandleErr(pingErr)
		return nil, pingErr
	}
	return db, nil
}

//读取DB配置文件
func GetDBConnectAddr(filePath string) (dbType string, addr string, e error) {
	rsl, e := util.ReadLine(filePath, 100)
	if e != nil {
		return
	}

	var db_type, db_userN, db_pwd, db_dbN, db_url, db_agt string
	var flog bool = false

	for _, v := range rsl {
		if strings.Contains(v, "db_type") {
			db_type = strings.Split(v, "=")[1]
		} else if strings.Contains(v, "db_userN") {
			db_userN = strings.Split(v, "=")[1]
		} else if strings.Contains(v, "db_pwd") {
			flog = true
			db_pwd = strings.Split(v, "=")[1]
		} else if strings.Contains(v, "db_dbN") {
			db_dbN = strings.Split(v, "=")[1]
		} else if strings.Contains(v, "db_url") {
			db_url = strings.Split(v, "=")[1]
		} else if strings.Contains(v, "db_agt") {
			db_agt = strings.Split(v, "=")[1]
		}
	}

	if !util.IsNotEmpty(db_type) {
		fmt.Println("not found db_type properties or db_type no empty...")
	}
	if !util.IsNotEmpty(db_dbN) {
		fmt.Println("not found db_dbN properties or db_dbN no empty...")
	}
	if !util.IsNotEmpty(db_userN) {
		fmt.Println("not found db_userN properties or db_userN no empty...")
	}
	if !util.IsNotEmpty(db_agt) {
		fmt.Println("not found db_agt properties or db_agt no empty...")
	}
	if !util.IsNotEmpty(db_url) {
		fmt.Println("not found db_url properties or db_url no empty...")
	}
	if !flog {
		fmt.Println("not found db_pwd properties...")
	}

	ss := bytes.NewBufferString(db_userN)
	ss.WriteString(":")
	ss.WriteString(db_pwd)
	ss.WriteString("@")
	ss.WriteString(db_agt)
	ss.WriteString("(")
	ss.WriteString(db_url)
	ss.WriteString(")/")
	ss.WriteString(db_dbN)

	dbType = db_type
	addr = ss.String()

	fmt.Println("get connect addr --->", ss.String())
	return
}

//-------------------- dao_sql_rule  -----------------------------
//检索规则
const (
	EQ   = " = ? "
	LIKE = " like ? "
	GT   = " > ? "
	LT   = " < ? "
	GE   = " >= ? "
	LE   = " <= ? "
)


/*这个大佬写的相当得全面，有时间得时候来看看这个事情。*/