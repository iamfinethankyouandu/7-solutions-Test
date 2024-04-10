package main

type Tree struct {
	Value    int
	Children []*Tree
}

func FindMostValue(arr2d [][]int) int {
	tree := buildTree(arr2d)
	result := tree.Value
	return findRoute(tree, result)
}

func findRoute(node *Tree, result int) int {
	if node == nil || len(node.Children) == 0 {
		return result
	}

	if node.Children[0].Value > node.Children[1].Value {
		result += node.Children[0].Value
		result = findRoute(node.Children[0], result)
	} else {
		result += node.Children[1].Value
		result = findRoute(node.Children[1], result)
	}

	return result
}

// BuildTree สร้างต้นไม้จาก arr 2d
func buildTree(data [][]int) *Tree {
	var root *Tree
	var lastLevelNodes []*Tree

	for i, level := range data {
		var currentLevelNodes []*Tree
		for j, value := range level {
			node := &Tree{Value: value}
			if i > 0 {
				// การเชื่อมโยงโหนดกับโหนดลูก
				if j < len(lastLevelNodes) {
					lastLevelNodes[j].Children = append(lastLevelNodes[j].Children, node)
				}
				if j > 0 {
					lastLevelNodes[j-1].Children = append(lastLevelNodes[j-1].Children, node)
				}
			}
			currentLevelNodes = append(currentLevelNodes, node)
		}
		lastLevelNodes = currentLevelNodes
		if i == 0 {
			root = currentLevelNodes[0] // ระบุ root
		}
	}

	return root
}
