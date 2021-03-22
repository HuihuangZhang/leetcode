package problem_110

import (
	"encoding/json"
	"testing"
)

func constructTreeFromStr(str string) (*TreeNode, error) {
	arr := []*int{}
	if err := json.Unmarshal([]byte(str), &arr); err != nil {
		return nil, err
	}
	tree := constructTree(&arr)
	return tree, nil
}

func TestConstructTree(t *testing.T) {
	str := "[1,2,2,3,3,null,null,4,4]"
	err, tree := constructTreeFromStr(str)
	if err != nil {
		t.Logf("error: %v+", err)
	}
	t.Logf("Tree: %v+", tree)
}

func TestIsBalance(t *testing.T) {
	str := "[1,2,2,3,null,null,3,4,null,null,4]"
	// str := "[2,3,null,4]"
	tree, err := constructTreeFromStr(str)
	if err != nil {
		t.Logf("error: %v+", err)
	}
	b := isBalanced(tree)
	t.Logf("the tree is balance: %t", b)
}

func TestIsBalanceV2(t *testing.T) {
	str := "[3,9,20,null,null,15,7]"
	// str := "[2,3,null,4]"
	tree, err := constructTreeFromStr(str)
	if err != nil {
		t.Logf("error: %v+", err)
	}
	b := isBalancedV2(tree)
	t.Logf("the tree is balance: %t", b)
}
