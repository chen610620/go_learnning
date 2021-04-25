/*
*
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答： Dao的作用是封装对数据库的访问：增删改查，不涉及业务逻辑，只是达到按某个条件获得指定数据的要求.
不确定业务形态， 所以应该将这个error提交给上层。由上层业务场景确定代码逻辑。
如：
  1. 只是某种查询列表操作，因业务决定忽略该error。
  2. 插入前查询（防止重复插入）。 则需要对该error进行判定。

ps: 场景2 也可以提供一个Exist()方法， 在内部消化该错误。 用户先调用Exist再插入。

*/

package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func main() {
	err := NoReturnError()
	if errors.Is(err, sql.ErrNoRows) {
		//  todo something
	}

	// todo somethings

}

func NoReturnError() error {

	sql := "select * from xx"

	// this is sql return error.
	err := errors.New("this is sql.ErrNoRows")

	return fmt.Errorf("sql:%s, err: %w", sql, err)
}
