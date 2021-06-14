package goclub_algo_linked_list

import "testing"
import "github.com/stretchr/testify/suite"

func TestNewSingleListPractice(t *testing.T) {
	suite.Run(t, &TestSingleListSuite{
		NewList: func(node *SingleListNode) SingleLister {
			return NewSingleListPractice(node)
		},
	})
}