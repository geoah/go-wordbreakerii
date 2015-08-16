# Word Break II in Golang

[Word Break II by Leetcode.com](https://leetcode.com/problems/word-break-ii/)

## The problem

Given a `string` s and a dictionary of words dict, add spaces in s to construct
a sentence where each word is a valid dictionary word. Return all such possible
sentences.  

For example, given `s = "catsanddog"`, `dict = ["cat", "cats", "and", "sand",
"dog"]`;  
A solution is `["cats and dog", "cat sand dog"]`.

## The actual problem

The example provided by the author manages to disguise the true difficulty of
the problem at hand.

The given example can be easily and quickly solved by a recursive function that
once it finds a word it will re-run itself from that position to find the next
available word, and so on.  
Such a solution will very quickly present problems when given more complex
string/dictionary pairs. e.g. `s = "aaaaaaaaaaaaaaaaaaaaaa"`, `dict = ["a",
"aa", "aaa", "aaaa", "aaaaa"]` as their complexity is significally higher.  
An additional issue with this solution is identifying strings that cannot be
broken down such as `s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"`,
`dict = ["a", "aa", "aaa", "aaaa", "aaaaa"]`.

## The solution

The most promising way to solve the problems would be to take them one at a
time and try them independently.

The first logical thing to do would be to make sure that the `string` is
actually breakable using the `dict`. The `string` must start and end with a
`dict` word and have no gaps between possible words, so each word should finish
where another one starts. Checking for this from right to left would make sense.

While checking if the `string` is breakable, we can store each found word with
their starting/ending positions for future reference. Since the `string` stays
the same, we can easily use their lengths and starting positions to retrieve
the word that was matched instead of storing the actual word.  
Using the original example the map would look like
`[[3 4] [] [] [7] [7] [] [] [10] [] []]`.

Once we have the starting position of all possible words and their lengths, we
should be able to construct a list of all possible branches of the found words
and print them. In our example such a list would look something like
`[[0 3 7 10], [0 4 7 10]]`.

There seem to be a couple ways to achieve this. Recursion, creating a trie/tree
and traversing it using breadth-first or depth-first search, or simply going
through the found words and creating the possible sentence paths by traversing
it similar to bfs/dfs. The last approach was more interesting for implementing
it in Golang as the possibility of using channels to optimize the search seemed
like a good candidate. (current channel implementation attempt in branch
`features/add-concurrency`).

Printing this now would be trivial either by getting each substring from the
original `string` or just by adding spaces from right to left on the appropriate
positions.

## Current implementation issues

Currently the dfs implementation seems to be buggy in specific edge cases.  
e.g. `s = "aaaaaaaaa"`, `dict = []string{"a", "aa", "aaa"}` while breaking up
the words correctly, the sentences creates contain the same sentence more than
once.
