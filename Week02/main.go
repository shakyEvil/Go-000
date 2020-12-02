package main

import (
    "fmt"
    "database/sql"
    xerrors "github.com/pkg/errors"
)

 type User struct {
    uid int64
    name string
    age int32
    email string
}

type UserDao interface {
    GetUser(uid int64) (user User, err error)
}

type DbUserDao struct {
}

func (dud DbUserDao) GetUser(uid int64) (user User, err error) {
    if uid%2 == 0 {
        return User{uid, "xiaoming", 18, "xxx@xxx.com"}, nil
    } else {
        //mock db err
        err := sql.ErrNoRows
        //wrap err, save caller stack info
        return User{}, xerrors.Wrapf(err, "dao query err")
    }
}

type BizRequest struct {
    uid int64
}

type BizResponse struct {
    code int
    data string
}

var dao DbUserDao

func biz(bizRequest BizRequest) (bizResponse BizResponse, err error) {
    user, err := dao.GetUser(12654757)    
    if err != nil {
        //within application, do not handle(principle: error should be only handle once), return err
        return BizResponse{0, "failed"}, err
    }
    fmt.Printf("dao response: %v\n", user)
    return BizResponse{200, "success"}, nil
}





func main() {
    fmt.Println("service call start")
    bizResponse, err := biz(BizRequest{1298564})

    if err != nil {
        //when back to the origin call locate, handle err, output err and stack info
        fmt.Printf("service call occur error, %+v", err)
    }
    
    fmt.Println("service call end")
    fmt.Printf("service response : %v\n", bizResponse)

}

