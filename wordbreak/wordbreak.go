package wordbreak

import (
	"sort"
)

type Block struct {
	start, end int
	solved     bool
}

func WordBreak(s string, wordDict []string) []string {
	sort.SliceStable(wordDict, func(i, j int) bool {
		return len(wordDict[i]) > len(wordDict[j])
	})

	xs := wordBreakRecursive(s, wordDict, 0, []Block{{0, len(s), false}})

	words := []string{}
	for _, x := range xs {
		if x.solved {
			words = append(words, s[x.start:x.end])
		} else {
			words = append(words, "not possible")
		}
	}
	return words
}

func wordBreakRecursive(
	s string, wordDict []string,
	dictIndex int, blocks []Block,
) []Block {
	if len(blocks) == 0 {
		return []Block{}
	}
	if dictIndex == len(wordDict) {
		return blocks
	}

	dictWord := wordDict[dictIndex]
	solvedBlocks := []Block{}

	for _, x := range blocks {
		if x.solved {
			continue
		}
		for i := x.start; i < x.end-len(dictWord)+1; i++ {
			substring := s[i : i+len(dictWord)]
			if substring == dictWord {
				solveBlock := Block{i, i + len(dictWord), true}
				solvedBlocks = append(solvedBlocks, solveBlock)
			}
		}
	}

	if len(solvedBlocks) == 0 {
		// try to solve with next word in dictionary
		return wordBreakRecursive(s, wordDict, dictIndex+1, blocks)
	}

	newBlocks := []Block{}

	if solvedBlocks[0].start > blocks[0].start {
		xs := wordBreakRecursive(s, wordDict, dictIndex+1, []Block{
			{blocks[0].start, solvedBlocks[0].start, false},
		})
		newBlocks = append(newBlocks, xs...)
	}

	newBlocks = append(newBlocks, solvedBlocks[0])

	for i := 1; i < len(solvedBlocks); i++ {
		previousBlock := solvedBlocks[i-1]
		currentBlock := solvedBlocks[i]
		if currentBlock.start > previousBlock.end {
			xs := wordBreakRecursive(s, wordDict, dictIndex+1, []Block{
				{previousBlock.end, currentBlock.start, false},
			})
			newBlocks = append(newBlocks, xs...)
		}
		newBlocks = append(newBlocks, solvedBlocks[i])
	}

	lastSolvedBlock := solvedBlocks[len(solvedBlocks)-1]
	lastUnsolvedBlock := blocks[len(blocks)-1]
	if lastSolvedBlock.end < lastUnsolvedBlock.end {
		xs := wordBreakRecursive(s, wordDict, dictIndex+1, []Block{
			{lastSolvedBlock.end, lastUnsolvedBlock.end, false},
		})
		newBlocks = append(newBlocks, xs...)
	}

	return newBlocks
}
