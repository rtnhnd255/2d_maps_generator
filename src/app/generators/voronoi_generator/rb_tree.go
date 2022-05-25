package voronoi_generator

type RBNodeValue interface {
	setToNode(node *rbnode)
	getNode() *rbnode
}

// rbnode
type rbnode struct {
	key    interface{}
	value  RBNodeValue
	left   *rbnode
	right  *rbnode
	parent *rbnode
	_red   bool
}

func (n *rbnode) red() bool {
	if n == nil {
		return false
	}
	return n._red
}

func (n *rbnode) grandparent() *rbnode {
	return n.parent.parent
}

func (n *rbnode) sibling() *rbnode {
	if n == n.parent.left {
		return n.parent.right
	} else {
		return n.parent.left
	}
}

func (n *rbnode) uncle() *rbnode {
	return n.parent.sibling()
}

func (n *rbnode) maxNode() *rbnode {
	for n.right != nil {
		n = n.right
	}
	return n
}

// RBTree
type RBTree struct {
	root      *rbnode
	length    int
	orderFunc Less
}

func NewRBTree(orderFunc Less) *RBTree {
	return &RBTree{orderFunc: orderFunc}
}

func (t *RBTree) Length() int {
	return t.length
}

func (t *RBTree) Insert(key interface{}, value RBNodeValue) {
	var insertNode *rbnode

	rbn := &rbnode{
		key:   key,
		value: value,
		_red:  true,
	}

	if t.root != nil {
		node := t.root
	LOOP:
		for {
			switch {
			case t.orderFunc(key, node.key):
				if node.left == nil {
					node.left = rbn
					insertNode = node.left
					break LOOP
				}
				node = node.left
			case t.orderFunc(node.key, key):
				if node.right == nil {
					node.right = rbn
					insertNode = node.right
					break LOOP
				}
				node = node.right
			default: // =
				node.key = key
				node.value = value
				return
			}
		}
		insertNode.parent = node
	} else {
		t.root = rbn
		insertNode = t.root
	}
	t.insertCase1(insertNode)
	t.length++

}

func (t *RBTree) insertCase1(n *rbnode) {
	if n.parent == nil {
		n._red = false
		return
	}
	t.insertCase2(n)
}

func (t *RBTree) insertCase2(n *rbnode) {
	if !n.parent._red {
		return
	}
	t.insertCase3(n)
}

func (t *RBTree) insertCase3(n *rbnode) {
	if n.uncle()._red {
		n.parent._red = false
		n.uncle()._red = false
		n.grandparent()._red = true
		t.insertCase1(n.grandparent())
		return
	}
	t.insertCase4(n)
}

func (t *RBTree) insertCase4(n *rbnode) {
	if n == n.parent.right && n.parent == n.grandparent().left {
		t.rotateLeft(n.parent)
		n = n.left
	} else if n == n.parent.left && n.parent == n.grandparent().right {
		t.rotateRight(n.parent)
		n = n.right
	}
	t.insertCase5(n)
}

func (t *RBTree) insertCase5(n *rbnode) {
	n.parent._red = false
	n.grandparent()._red = true
	if n == n.parent.left && n.parent == n.grandparent().left {
		t.rotateRight(n.grandparent())
		return
	} else if n == n.parent.right && n.parent == n.grandparent().right {
		t.rotateLeft(n.grandparent())
	}
}

func (t *RBTree) replace(old, new *rbnode) {
	if old.parent == nil {
		t.root = new
	} else {
		if old == old.parent.left {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}
	if new != nil {
		new.parent = old.parent
	}
}

func (t *RBTree) rotateRight(n *rbnode) {
	left := n.left
	t.replace(n, left)
	n.left = left.right
	if left.right != nil {
		left.right.parent = n
	}
	left.right = n
	n.parent = left
}

func (t *RBTree) rotateLeft(n *rbnode) {
	right := n.right
	t.replace(n, right)
	n.right = right.left
	if right.left != nil {
		right.left.parent = n
	}
	right.left = n
	n.parent = right
}

func (t *RBTree) Get(key interface{}) RBNodeValue {
	n := t.find(key)
	if n == nil {
		return nil
	}
	return n.value
}

func (t *RBTree) find(key interface{}) *rbnode {
	n := t.root
	for n != nil {
		switch {
		case t.orderFunc(key, n.key):
			n = n.left
		case t.orderFunc(n.key, key):
			n = n.right
		default:
			return n
		}
	}
	return nil
}

func (t *RBTree) Delete(key interface{}) {
	var child *rbnode

	n := t.find(key)
	if n == nil {
		return
	}

	if n.left != nil && n.right != nil {
		pred := n.left.maxNode()
		n.key = pred.key
		n.value = pred.value
		n = pred
	}

	if n.left == nil || n.right == nil {
		if n.right == nil {
			child = n.left
		} else {
			child = n.right
		}
		if !n._red {
			n._red = child._red
			t.deleteCase1(n)
		}

		t.replace(n, child)
		if n.parent == nil && child != nil {
			child._red = false
		}
	}
	t.length--
}

func (t *RBTree) deleteCase1(n *rbnode) {
	if n.parent == nil {
		return
	}

	t.deleteCase2(n)
}
func (t *RBTree) deleteCase2(n *rbnode) {
	sibling := n.sibling()
	if sibling._red == true {
		n.parent._red = true
		sibling._red = false
		if n == n.parent.left {
			t.rotateLeft(n.parent)
		} else {
			t.rotateRight(n.parent)
		}
	}
	t.deleteCase3(n)
}
func (t *RBTree) deleteCase3(n *rbnode) {
	sibling := n.sibling()
	if !n.parent._red &&
		!sibling._red &&
		!sibling.left._red &&
		!sibling.right._red {
		sibling._red = true
		t.deleteCase1(n.parent)
		return
	}
	t.deleteCase4(n)
}
func (t *RBTree) deleteCase4(n *rbnode) {
	sibling := n.sibling()
	if n.parent._red &&
		!sibling._red &&
		!sibling.left._red &&
		!sibling.right._red {
		sibling._red = true
		n.parent._red = false
		return
	}
	t.deleteCase5(n)
}
func (t *RBTree) deleteCase5(n *rbnode) {
	sibling := n.sibling()
	if n == n.parent.left &&
		!sibling._red &&
		sibling.left._red &&
		!sibling.right._red {
		sibling._red = true
		sibling.left._red = false
		t.rotateRight(sibling)
	} else if n == n.parent.right &&
		!sibling._red &&
		sibling.right._red &&
		!sibling.left._red {
		sibling._red = true
		sibling.right._red = false
		t.rotateLeft(sibling)
	}
	t.deleteCase6(n)
}
func (t *RBTree) deleteCase6(n *rbnode) {
	sibling := n.sibling()
	sibling._red = n.parent._red
	n.parent._red = false
	if n == n.parent.left && sibling.right._red {
		sibling.right._red = false
		t.rotateLeft(n.parent)
		return
	}
	sibling.left._red = false
	t.rotateRight(n.parent)
}
