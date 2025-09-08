package internal

type node struct {
	parent   *node
	children []*node
}

func (n *node) addChild(c *node) {
	c.parent = n
	n.children = append(n.children, c)
}
