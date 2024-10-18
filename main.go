package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

var (
	symbols = [...]string{
		"*",
		":",
		"@",
		"$",
		".",
		"+",
		"!",
		"&",
		"_",
		"-",
	}
	words = [...]string{
		"admiring",
		"adoring",
		"affectionate",
		"agitated",
		"amazing",
		"angry",
		"awesome",
		"beautiful",
		"blissful",
		"bold",
		"boring",
		"brave",
		"busy",
		"charming",
		"clever",
		"cool",
		"compassionate",
		"competent",
		"condescending",
		"confident",
		"cranky",
		"crazy",
		"dazzling",
		"determined",
		"distracted",
		"dreamy",
		"eager",
		"ecstatic",
		"elastic",
		"elated",
		"elegant",
		"eloquent",
		"epic",
		"exciting",
		"fervent",
		"festive",
		"flamboyant",
		"focused",
		"friendly",
		"frosty",
		"funny",
		"gallant",
		"gifted",
		"goofy",
		"gracious",
		"great",
		"happy",
		"hardcore",
		"heuristic",
		"hopeful",
		"hungry",
		"infallible",
		"inspiring",
		"interesting",
		"intelligent",
		"jolly",
		"jovial",
		"keen",
		"kind",
		"laughing",
		"loving",
		"lucid",
		"magical",
		"mystifying",
		"modest",
		"musing",
		"naughty",
		"nervous",
		"nice",
		"nifty",
		"nostalgic",
		"objective",
		"optimistic",
		"peaceful",
		"pedantic",
		"pensive",
		"practical",
		"priceless",
		"quirky",
		"quizzical",
		"recursing",
		"relaxed",
		"reverent",
		"romantic",
		"sad",
		"serene",
		"sharp",
		"silly",
		"sleepy",
		"stoic",
		"strange",
		"stupefied",
		"suspicious",
		"sweet",
		"tender",
		"thirsty",
		"trusting",
		"unruffled",
		"upbeat",
		"vibrant",
		"vigilant",
		"vigorous",
		"wizardly",
		"wonderful",
		"xenodochial",
		"youthful",
		"zealous",
		"zen",
	}
)

func main() {
redo:
	rand.Seed(time.Now().UnixNano())

	// 3 words
	//name := strings.Title(words[rand.Intn(len(words))]) + strings.Title(words[rand.Intn(len(words))]) + strings.Title(words[rand.Intn(len(words))]) //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)

	// 2 words
	name := strings.Title(words[rand.Intn(len(words))]) + symbols[rand.Intn(len(symbols))] + strings.Title(words[rand.Intn(len(words))])

	name += symbols[rand.Intn(len(symbols))] + strconv.Itoa(rand.Intn(9)) //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
	// regex to check for dupe chars
	// note the stdlib doesn't support back searching
	re := regexp2.MustCompile(`(.)\1{1,}`, 0)
	match, _ := re.MatchString(name)
	if match == true {
		goto redo
	}
	fmt.Println(name)
}
