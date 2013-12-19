package bstree

import "fmt"

type Node struct {
	Key                 int
	Value               interface{}
	Left, Right, Parent *Node
}

type BStree struct {
	Root  *Node
	Count int
}

func New() *BStree {
	return new(BStree)
}

func Insert(t *BStree, key int, value interface{}) {
	t.Insert(key, value)
}

func Delete(t *BStree, key int) bool {
	return t.Delete(key)
}

func Get(t *BStree, key int) (interface{}, bool, *Node) {
	return t.Get(key)
}

func Display(t *BStree) {
	t.Display()
}

func (t *BStree) Insert(key int, value interface{}) {
	n := new(Node)
	n.Key = key
	n.Value = value
	pos := t.getParentNode(key)
	n.Parent = pos

	if pos == nil {
		t.Root = n
	} else if key <= pos.Key {
		pos.Left = n
	} else {
		pos.Right = n
	}

	t.Count++
}

func (t *BStree) Delete(key int) bool {
	_, ok, node := t.Get(key)

	if !ok {
		return false
	}

	parent := node.Parent

	if (node.Left == nil) && (node.Right == nil) {
		if parent.Left == node {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if node.Left == nil {
		if parent.Left == node {
			parent.Left = node.Right
		} else {
			parent.Right = node.Right
		}
		node.Right.Parent = parent
	} else if node.Right == nil {
		if parent.Left == node {
			parent.Left = node.Left
		} else {
			parent.Right = node.Left
		}
		node.Left.Parent = parent
	} else {
		tmp := node.Left

		for tmp.Right != nil {
			tmp = tmp.Right
		}
		node.Value = tmp.Value

		if tmp.Parent.Right == tmp {
			tmp.Parent.Right = nil
		} else {
			tmp.Parent.Left = nil
		}

		node = tmp
		t.Count--
	}

	return true
}

func (t *BStree) Get(key int) (interface{}, bool, *Node) {
	tmp := t.Root
	found := false
	var value interface{}

	for tmp != nil {
		if tmp.Key == key {
			value = tmp.Value
			found = true
			break
		} else if key < tmp.Key {
			tmp = tmp.Left
		} else {
			tmp = tmp.Right
		}
	}

	return value, found, tmp
}

func (t *BStree) Display() {
	t.Inorder(t.Root)
}

func (t *BStree) Inorder(tmp *Node) {
	if tmp == nil {
		return
	}
	fmt.Println("Key is ", tmp.Key, "Value is", tmp.Value)
	t.Inorder(tmp.Left)
	t.Inorder(tmp.Right)
}

func (t *BStree) getParentNode(key int) *Node {
	tmp := t.Root
	parent := t.Root

	for tmp != nil {
		parent = tmp
		if key <= tmp.Key {
			tmp = tmp.Left
		} else {
			tmp = tmp.Right
		}
	}

	return parent
}
