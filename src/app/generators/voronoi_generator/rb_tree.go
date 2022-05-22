package voronoi_generator

type RBNodeValue interface {
	setToNode(node *rbnode)
	getNode() *rbnode
}

// rbnode
type rbnode struct {
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
	orderFunc OrderFunc
}

func NewRBTree(orderFunc OrderFunc) *RBTree {
	return &RBTree{orderFunc: orderFunc}
}

func (t *RBTree) Length() int {
	return t.length
}
