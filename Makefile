.PONY: test
test:
	go test ./...


.PONY: test-anagram
test-anagram:
	cd anagram && go test ./...

.PONY: test-brokennode
test-brokennode:
	cd brokennode && go test ./...

.PONY: test-buildword
test-buildword:
	cd buildword && go test ./...

.PONY: test-chess
test-chess:
	cd chess && go test ./...

.PONY: test-coins
test-coins:
	cd coins && go test ./...

.PONY: test-floyd
test-floyd:
	cd floyd && go test ./...

.PONY: test-functionfrequency
test-functionfrequency:
	cd functionfrequency && go test ./...

.PONY: test-jaro
test-jaro:
	cd anagram && go test ./...

.PONY: test-lastlettergame
test-lastlettergame:
	cd lastlettergame && go test ./...

.PONY: test-mergesort
test-mergesort:
	cd mergesort && go test ./...

.PONY: test-missingnumbers
test-missingnumbers:
	cd missingnumbers && go test ./...

.PONY: test-node-degree
test-node-degree:
	cd node_degree && go test ./...

.PONY: test-reverseparentheses
test-reverseparentheses:
	cd reverseparentheses && go test ./...

.PONY: test-romannumerals
test-romannumerals:
	cd romannumerals && go test ./...

.PONY: test-secretmessage
test-secretmessage:
	cd secretmessage && go test ./...

.PONY: test-shorthash
test-shorthash:
	cd shorthash && go test ./...

.PONY: test-snowflakes
test-snowflakes:
	cd snowflakes && go test ./...

.PONY: test-sumdecimal
test-sumdecimal:
	cd sumdecimal && go test ./...

.PONY: test-warriors
test-warriors:
	cd warriors && go test ./...
