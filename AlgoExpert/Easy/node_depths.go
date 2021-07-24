package main

type BinaryTree struct {
    Value       int
    Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
    return calculateSumDepths(root, 0)
}

func calculateSumDepths(node *BinaryTree, depth int) int {
    if node == nil {
        return 0
    }

    return depth + calculateSumDepths(node.Left, depth+1) + calculateSumDepths(node.Right, depth+1)
}
