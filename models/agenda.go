package models

import (
	"errors"
	//"fmt"
	"reflect"
	//"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Agenda struct {
	Id int `orm:"column(id);pk"`
	Device string `orm:"column(device)"`
	Token string `orm:"column(token);size(225)"`
	Users_id int `orm:"column(users_id)"`
	Open_status int `orm:"column(open_status)"`
	Title string `orm:"column(title)"`
	Status string `orm:"column(status)"`
	Start_date string `orm:"column(start_date)"`
	Finish_date string `orm:"column(finish_date)"`
	Start_time string `orm:"column(start_time)"`
	Finish_time string `orm:"column(finish_time)"`
	Description string `orm:"column(description)"`
	Hashtag string `orm:"column(hashtag)"`
	Linked string `orm:"column(linked)"`
	Notify string `orm:"column(notify)"`
	Notify_time string `orm:"column(notify_time)"`
	Notify_email string `orm:"column(notify_email)"`
	Notify_sms string `orm:"column(notify_sms)"`
	Expired string `orm:"column(expired)"`
	Is_active string `orm:"column(is_active)"`
	Place_id int `orm:"column(place_id)"`
	Display int `orm:"column(display)"`
	Click int `orm:"column(click)"`
	Repeat_at string `orm:"column(repeat_at)"`
	Coloring string `orm:"column(coloring)"`
	Background_id int `orm:"column(background_id)"`
	Created_at string `orm:"column(created_at)"`
	Updated_at string `orm:"column(updated_at)"`
	Deleted_at  string `orm:"column(deleted_at)"`
}

func (t *Agenda) TableName() string {
	return "agenda"
}

func init() {
	orm.RegisterModel(new(Agenda))
}

// GetAllKv retrieves all Kv matches certain condition. Returns empty list if
// no records exist
func GetAllAgenda(querystr string, fieldstr string, sortbystr string, orderstr string,
	offsetv int64, limitv int64) (ml []interface{}, err error) {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var likequery map[string]string = make(map[string]string)
	var inquery map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if fieldstr != "" {
		fields = strings.Split(fieldstr, ",")
	}
	// limit: 10 (default is 10)
	limit = limitv
	// offset: 0 (default is 0)
	offset = offsetv
	// sortby: col1,col2
	if sortbystr != "" {
		sortby = strings.Split(sortbystr, ",")
	}
	// order: desc,asc
	if orderstr != "" {
		order = strings.Split(orderstr, ",")
	}

	// query: k:v,k:v
	for _, cond := range strings.Split(querystr, ",") {
		if strings.Contains(cond, "_like_") {
			kv := strings.Split(cond, "_like_")
			if len(kv) != 2 {
				return nil, errors.New("Error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			likequery[k] = v
		} else if strings.Contains(cond, "_in_") {
			kv := strings.Split(cond, "_in_")
			if len(kv) != 2 {
				return nil, errors.New("Error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			inquery[k] = v
		} else if strings.Contains(cond, ":") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				return nil, errors.New("Error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(Agenda))
	// query k=v
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	//for k, v:= range query {
	//	s:= v.Description
	//	lastBin := strings.LastIndex( s, "-" )
	//    qs = qs.Filter(s[0:lastBin], v)
	//}
	for k, v := range likequery {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k+"__icontains", v)
	}
	for k, v := range inquery {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		ids := strings.Split(v, ":")
		nv := make([]interface{}, len(ids))
		for i, vn := range ids {
			nv[i] = vn
		}

		qs = qs.Filter(k+"__in", nv...)
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

	var l []Agenda
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

func CountGetAllAgenda(querystr string) (ml int64, err error) {
	var query map[string]string = make(map[string]string)
	var likequery map[string]string = make(map[string]string)
	var inquery map[string]string = make(map[string]string)

	for _, cond := range strings.Split(querystr, ",") {
		if strings.Contains(cond, "_like_") {
			kv := strings.Split(cond, "_like_")
			if len(kv) != 2 {
				return -1, errors.New("Error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			likequery[k] = v
		} else if strings.Contains(cond, "_in_") {
			kv := strings.Split(cond, "_in_")
			if len(kv) != 2 {
				return -1, errors.New("Error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			inquery[k] = v
		} else if strings.Contains(cond, ":") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				return -1, errors.New("Error: invalid query key/value pair")
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(Agenda))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	for k, v := range likequery {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k+"__icontains", v)
	}
	for k, v := range inquery {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		ids := strings.Split(v, ":")
		nv := make([]interface{}, len(ids))
		for i, vn := range ids {
			nv[i] = vn
		}

		qs = qs.Filter(k+"__in", nv...)
	}
	if total, err := qs.Count(); err == nil {
		return total, nil
	}
	return -1, err
}
