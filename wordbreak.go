package wordbreak

import (
	"sort"
)

type Block struct {
	start, end int
	dictIndex  int
}

func WordBreak(s string, wordDict []string) []string {
	sort.SliceStable(wordDict, func(i, j int) bool {
		return len(wordDict[i]) > len(wordDict[j])
	})

	xs := wordBreakRecursive(s, wordDict, 0, len(s), 0)

	words := []string{}
	for _, x := range xs {
		if x.dictIndex > -1 {
			words = append(words, s[x.start:x.end])
		} else {
			words = append(words, "not possible")
		}
	}
	return words
}

func wordBreakRecursive(
	s string, wordDict []string,
	start, end int,
	dictIndex int,
) []Block {
	if dictIndex == len(wordDict) {
		// when no more words in dictionary to try match
		// then just return unmatched block
		return []Block{{start, end, -1}}
	}

	dictWord := wordDict[dictIndex]
	solvedBlocks := []Block{}

	for i := start; i < end-len(dictWord)+1; i++ {
		substring := s[i : i+len(dictWord)]
		if substring == dictWord {
			b := Block{i, i + len(dictWord), dictIndex}
			solvedBlocks = append(solvedBlocks, b)
		}
	}

	if len(solvedBlocks) == 0 {
		// try to match with next word in dictionary
		return wordBreakRecursive(s, wordDict, start, end, dictIndex+1)
	}

	newBlocks := []Block{}

	if solvedBlocks[0].start > start {
		xs := wordBreakRecursive(s, wordDict,
			start, solvedBlocks[0].start,
			dictIndex+1,
		)
		newBlocks = append(newBlocks, xs...)
	}

	newBlocks = append(newBlocks, solvedBlocks[0])

	for i := 1; i < len(solvedBlocks); i++ {
		previousBlock := solvedBlocks[i-1]
		currentBlock := solvedBlocks[i]
		if currentBlock.start > previousBlock.end {
			xs := wordBreakRecursive(s, wordDict,
				previousBlock.end, currentBlock.start,
				dictIndex+1,
			)
			newBlocks = append(newBlocks, xs...)
		}
		newBlocks = append(newBlocks, solvedBlocks[i])
	}

	lastSolvedBlock := solvedBlocks[len(solvedBlocks)-1]
	if lastSolvedBlock.end < end {
		xs := wordBreakRecursive(s, wordDict,
			lastSolvedBlock.end, end,
			dictIndex+1,
		)
		newBlocks = append(newBlocks, xs...)
	}

	return newBlocks
}
