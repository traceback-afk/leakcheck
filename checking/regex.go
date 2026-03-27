package checking

import (
	"regexp"
	"encoding/json"
	"log"
	"os"
)

type Rule struct {
	Name string
	Pattern *regexp.Regexp
}

var rules []Rule

func init(){
	rules = readRulesFromJson("regex_rules.json")
}

func readRulesFromJson(path string) []Rule{
	var rules_map map[string]string
	content, err := os.ReadFile("regex_rules.json")
	if err != nil {
		log.Fatal("Error when opening the file.", err)
	}
	err = json.Unmarshal(content, &rules_map)

	var rules_output []Rule
	for name, pattern := range rules_map {
		rule := Rule{
			Name: name,
			Pattern: regexp.MustCompile(pattern),
		}
		rules_output = append(rules_output, rule)
	}
	return rules_output
}

func MatchesRules(input string) (bool, string) {
	for _, rule := range rules {
		if rule.Pattern.MatchString(input){
			return true, rule.Name
		}
	}
	return false, ""
}
