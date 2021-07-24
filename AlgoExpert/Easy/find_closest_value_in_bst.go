package main

type BST struct {
    Value int

    Left  *BST
    Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
    currNode := tree
    closest := currNode.Value

    for currNode != nil {
        if absdiff(target, closest) > absdiff(target, currNode.Value) {
            closest = currNode.Value
        }

        if currNode.Value > target {
            currNode = currNode.Left
        } else if currNode.Value < target {
            currNode = currNode.Right
        } else {
            break
        }
    }

    return closest
}

func absdiff(a, b int) int {
    if a > b {
        return a - b
    }

    return b - a
}
