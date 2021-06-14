package goclub_algo_linked_list

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestNewSingleListEasy(t *testing.T) {
	suite.Run(t, &TestSingleListSuite{
		NewList: func(node *SingleListNode) SingleLister {
			return NewSingleListEasy(node)
		},
	})
}