package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Schedules  Task
)

type Task struct {
}
var ctx context.Context


func (s Task) Initialize()error{
	var xxx =gctx.New()
	ctx = xxx
	_,err :=gcron.AddSingleton(ctx,"* * * * * *",Foo)
	if err != nil {
		return err
	}
	_,err =gcron.AddSingleton(ctx,"* * * * * *",Fooo)
	if err != nil {
		return err
	}
	return nil
}
func Foo(ctx context.Context)  {
	fmt.Println("tatk1")
}
func Fooo(ctx context.Context)  {
	fmt.Println("task2")

}
