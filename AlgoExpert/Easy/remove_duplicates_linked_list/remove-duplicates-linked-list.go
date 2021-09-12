package main

// LinkedList This is an input struct. Do not edit.
type LinkedList struct {
    Value int
    Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
    head := linkedList
    currNode := head

    for currNode.Next != nil {
        if currNode.Value == currNode.Next.Value {
            currNode.Next = currNode.Next.Next
        } else {
            currNode = currNode.Next
        }
    }

    return head
}
