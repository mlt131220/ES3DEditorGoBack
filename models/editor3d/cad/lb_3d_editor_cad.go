package cad

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Lb3dEditorCad struct {
	ConversionStatus  int       `orm:"column(conversion_status)" description:"0 转换中 1 转换完成 2 转换失败" json:"conversionStatus"`
	ConverterFilePath string    `orm:"column(converter_file_path);size(255)" description:"转换后的文件路径" json:"converterFilePath"`
	CreateTime        time.Time `orm:"column(createTime);type(datetime);null;auto_now_add" json:"createTime"`
	DelTag            int8      `orm:"column(delTag)" description:"删除标记，0 未删除 1 已删除" json:"delTag"`
	DelTime           time.Time `orm:"column(delTime);type(datetime);null" description:"删除时间" json:"delTime"`
	FileName          string    `orm:"column(file_name);size(255)" description:"文件名" json:"fileName"`
	FilePath          string    `orm:"column(file_path);size(255)" description:"源文件路径" json:"filePath"`
	Id                int       `orm:"column(id);auto" json:"id"`
	Thumbnail         string    `orm:"column(thumbnail);size(255);null" description:"缩略图" json:"thumbnail"`
	UpdateTime        time.Time `orm:"column(updateTime);type(datetime);auto_now_add" json:"updateTime"`
}

func (t *Lb3dEditorCad) TableName() string {
	return "lb_3d_editor_cad"
}

func init() {
	orm.RegisterModel(new(Lb3dEditorCad))
}

// AddLb3dEditorCad insert a new Lb3dEditorCad into database and returns
// last inserted Id on success.
func AddLb3dEditorCad(m *Lb3dEditorCad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLb3dEditorCadById retrieves Lb3dEditorCad by Id. Returns error if
// Id doesn't exist
func GetLb3dEditorCadById(id int) (v *Lb3dEditorCad, err error) {
	o := orm.NewOrm()
	v = &Lb3dEditorCad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLb3dEditorCad retrieves all Lb3dEditorCad matches certain condition. Returns empty list if
// no records exist
func GetAllLb3dEditorCad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Lb3dEditorCad))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Lb3dEditorCad
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// GetTotalLb3dEditorCad 获取总数量
func GetTotalLb3dEditorCad(query map[string]string) (total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Lb3dEditorCad))

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			//大小写不敏感模糊查询 + "__icontains"
			qs = qs.Filter(k+"__icontains", v)
		}
	}
	if cnt, err := qs.Count(); err == nil {
		return cnt, nil
	}
	return 0, err
}

// UpdateLb3dEditorCad updates Lb3dEditorCad by Id and returns error if
// the record to be updated doesn't exist
func UpdateLb3dEditorCadById(m *Lb3dEditorCad) (err error) {
	o := orm.NewOrm()
	v := Lb3dEditorCad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLb3dEditorCad deletes Lb3dEditorCad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLb3dEditorCad(id int) (err error) {
	o := orm.NewOrm()
	v := Lb3dEditorCad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Lb3dEditorCad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
