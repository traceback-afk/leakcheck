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

var blacklist_rules []Rule
var whitelist_rules []Rule

func init(){
	blacklist_rules = readRulesFromJson("blacklist_rules.json")
	whitelist_rules = readRulesFromJson("whitelist_rules.json")
}

func readRulesFromJson(path string) []Rule{
	var rules_map map[string]string
	content, err := os.ReadFile(path)
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
	for _, rule := range whitelist_rules {
		if rule.Pattern.MatchString(input){
			return false, "Whitelist:" + rule.Name
		}
	}

	for _, rule := range blacklist_rules {
		if rule.Pattern.MatchString(input){
			return true, rule.Name
		}
	}
	return false, ""
}
