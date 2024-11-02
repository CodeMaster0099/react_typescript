//go:build ignore

package main

func main() {
	var d declarer

	unknown := d.newNode("Unknown")
	endOfFile := d.newNode("EndOfFile")
}

type declarer struct {
	nodes  map[string]*NodeType
	groups []*group

	groupStack []*group
}

type group struct {
	name      string
	startName string
	endName   string
}

func (d *declarer) asGroup(name string, f func()) {
	d.groupStack = append(d.groupStack, &group{name: name})
	f()
	group := d.groupStack[len(d.groupStack)-1]
	d.groups = append(d.groups, group)
	d.groupStack = d.groupStack[:len(d.groupStack)-1]
}

func (d *declarer) newNode(name string) *NodeType {
	for _, group := range d.groupStack {
		if group.startName == "" {
			group.startName = name
		}
		group.endName = name
	}

	if d.nodes == nil {
		d.nodes = make(map[string]*NodeType)
	}

	n := &NodeType{name: name}
	d.nodes[name] = n
	return n
}

type Type interface {
	Name() string
}

type NodeType struct {
	name     string
	children []Child
}

func (n *NodeType) Name() string { return n.name }

type Child struct {
	name string
	typ  Type
}
