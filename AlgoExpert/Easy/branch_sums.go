package main

type BinaryTree struct {
    Value int
    Left  *BinaryTree
    Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
    var sums []int
    RecursiveBranchSums(root, 0, &sums)

    return sums
}

func RecursiveBranchSums(node *BinaryTree, runningSum int, sums *[]int) {
    if node == nil {
        return
    }

    runningSum += node.Value
    if node.Left == nil && node.Right == nil {
        *sums = append(*sums, runningSum)
        return
    }

    RecursiveBranchSums(node.Left, runningSum, sums)
    RecursiveBranchSums(node.Right, runningSum, sums)
}
