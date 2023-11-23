package test

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestSelectRows(t *testing.T) {
	Init()
	rows, err := ada.SelectRows()
	if err != nil {
		log.Fatal("错误是 ", err)
	}
	fmt.Printf("%+v", rows)
}

func TestSelectRow(t *testing.T) {
	Init()
	rows, err := ada.SelectRow()
	if err != nil {
		log.Fatal("错误是 ", err)
	}
	fmt.Printf("%+v\n", rows)
}

func TestModelFindOne(t *testing.T) {
	Init()
	ctx := context.Background()
	one, err := casrule.FindOne(ctx, 1)
	if err != nil {
		log.Fatal("err is ", err)
		return
	}
	fmt.Println(one)

}
