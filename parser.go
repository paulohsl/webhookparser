package webhookparser

import (
	"fmt"
	"regexp"
	. "bufio"
	"log"
	"sort"
)

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Pair struct {
	Key   string
	Value int
}

type logRegexp struct {
	*regexp.Regexp
}

func Parse(scanner *Scanner) (urlMap map[string]int, statusMap map[string]int, err error) {

	var urls = make(map[string]int)
	var statuses = make(map[string]int)

	var reg = logRegexp{regexp.MustCompile(`request_to="([^"]*)".+?response_status="([^"]*)"`)}

	for scanner.Scan() {
		if reg.MatchString(scanner.Text()) {
			log := reg.matchLine(scanner.Text())
			sumMapValue(log[1], urls)
			sumMapValue(log[2], statuses)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return urls, statuses, err

}

func (r *logRegexp) matchLine(line string) []string {
	return r.FindStringSubmatch(line)
}

func sumMapValue(value string, m map[string]int) {
	if total, ok := m[value]; ok {
		m[value] = total + 1
	} else {
		m[value]++
	}
}

func rankByQuantity(typeMap map[string]int) PairList {
	pl := make(PairList, len(typeMap))
	i := 0
	for k, v := range typeMap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func PrintTopRanked(typeMap map[string]int, top int) {
	rankedMap := rankByQuantity(typeMap)

	i := 0
	for i < top {
		fmt.Println(rankedMap[i])
		i++
	}
}

func PrintRanked(typeMap map[string]int) {
	rankedMap := rankByQuantity(typeMap)

	for _, val := range rankedMap {
		fmt.Println(val)
	}
}
