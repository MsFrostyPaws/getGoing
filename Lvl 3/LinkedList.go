package main
import "fmt"

// define node structure:
type Node struct {
    data int
    next_node *Node
}

// adding node to front of list:
func add_front(linked_list *Node, data int) *Node {
    if linked_list == nil {
        linked_list = &Node{data,nil}
    } else {
        linked_list = &Node{data,linked_list}
    }
    return linked_list 
}

// adding node to end of list:
func add_end(linked_list *Node, data int) *Node {
    var new_node, ptr *Node
    if linked_list == nil {
        linked_list = &Node{data,nil}
    } else {
        new_node = &Node{data,nil}
        ptr = linked_list
        for {
            if ptr.next_node != nil {
                ptr = ptr.next_node
            } else { break }
        }
        ptr.next_node = new_node 
    }
    return linked_list
}

// print all node values:
func print_list(linked_list *Node) {
    counter := 1
    for {
        fmt.Printf("Node %d = %d\n",counter,linked_list.data)
        if linked_list.next_node != nil {
            linked_list = linked_list.next_node
            counter += 1
        } else { break }
    }   
}

func main() {

    // initialize empty list:
    var linked_list *Node

    // add nodes:
    linked_list = add_front(linked_list,11)
    linked_list = add_end(linked_list,19)
    linked_list = add_end(linked_list,76)
    linked_list = add_front(linked_list,48)
    linked_list = add_front(linked_list,44)
    
    // print the list:
    print_list(linked_list)
            
}
