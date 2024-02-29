package repository

import (
	belajardbgolang "belajar-dbgolang"
	"belajar-dbgolang/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajardbgolang.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repotory",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository((belajardbgolang.GetConnection()))

	comment, err := commentRepository.FindById(context.Background(), 24)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository((belajardbgolang.GetConnection()))

	comments, err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
