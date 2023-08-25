# sensitive-words-filter

[中文](README-CN.md)

A sensitive word filtering algorithm implemented in Go.

The sensitive word detection uses an HTTP server to receive input and then performs validation.

Steps to use:

1. `git clone https://github.com/lqxhub/sensitive-words-filter.git`
2. `cd sensitive-words-filter`
3. `go build`
4. Open the address 127.0.0.1:8081 in a web browser.

Input the words to check in the input box to perform detection.

The rules for sensitive words are in the rule.txt file, with one rule per line.
