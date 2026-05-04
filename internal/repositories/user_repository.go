package repositories

import "context"
type UserFilter struct {
	Limit int
	Offset int
	Search string
	Role string
}
func LenStr(l []any) string{
	return strconv.Itoa(len(l))
}

func UserList(c context.Context, f UserFilter, moreArg ...int)([]models.User, error){
	db:=utils.GetDB()
	sqlWhere:=` `
	sqlArg :=[]any{f.Limit, f.Offset}
	if f.Search !={
		sqlArgs=append(sqlArgs, f.Search)
		sqlWhere +=`and (first_name ilike '%$` + LenStr(sqlArgs) +`%')`
	}
	if f.Role!=""{
		sqlArgs=append(sqlArgs,f.Search)
		sqlWhere+=`and (role=$` + LenStr(sqlArgs)+ `)`
	}
	rows, err:=db.Query (c, `select id,first_name, role, password, email from users where 1=1 `+sqlWhere+` limit $1 offset  $2`, sqlArgs...)
	if err!=nil{
		return nil, err
	}
	list :=[]models.User{}
	for rows.Next(){
		item:=models.User{}
		rows.Scan(&item.ID, &item.FirstName, &item.LastName, &item.Role, &item.Password, &item.Email)
		list= append(list, item)
	}
	return list,nil
}