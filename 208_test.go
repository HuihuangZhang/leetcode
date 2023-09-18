package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTrieTrue(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple");
	assert.True(t, trie.Search("apple"))
	assert.False(t, trie.Search("app"))
	assert.True(t, trie.StartsWith("app"))
	trie.Insert("app");
	assert.True(t, trie.Search("app"))
}
