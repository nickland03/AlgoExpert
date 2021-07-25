package main

type Node struct {
    Name     string
    Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
    array = append(array, n.Name)

    for _, c := range n.Children {
        array = append(array, c.Name)
    }

    return array
}