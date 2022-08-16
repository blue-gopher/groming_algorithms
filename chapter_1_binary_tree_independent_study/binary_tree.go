package main

import (
	"errors"
	"fmt"
)

// структура узла бинарного дерева
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// вставляем новое значение(узел, лист)
func (t *treeNode) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil.")
	}
	if t.val == value {
		return errors.New("This node value already exists.")
	}
	if t.val > value {
		if t.left == nil {
			t.left = &treeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}
	if t.val < value {
		if t.right == nil {
			t.right = &treeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

// поиск минимального значения
func (t *treeNode) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}

// поиск максимального значения
func (t *treeNode) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

// обход бинарного дерева
func (t *treeNode) PrintInOrder() {
	if t == nil {
		return
	}
	t.left.PrintInOrder()
	fmt.Println(t.val)
	t.right.PrintInOrder()
}
