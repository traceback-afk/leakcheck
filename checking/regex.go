package checking

import (
	"regexp"
)

type Rule struct {
	Name string
	Regex *regexp.Regexp
}

var rules = []Rule{
	{"AWS key", regexp.MustCompile(`AKIA[0-9A-Z]{16}`)},
	{"Stripe key", regexp.MustCompile(`sk_live_[0-9a-zA-Z]{24}`)},
	{"GitHub token", regexp.MustCompile(`ghp_[A-Za-z0-9]{36}`)},
	{"JWT", regexp.MustCompile(`\beyJ[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\b`)},
	{"Generic API key", regexp.MustCompile(`\b[A-Za-z0-9_\-]{32,64}\b`)},
	{"Base64 secret", regexp.MustCompile(`\b[A-Za-z0-9+/]{40,}={0,2}\b`)},
	{"Private key", regexp.MustCompile(`-----BEGIN [A-Z ]+PRIVATE KEY-----`)},
}

func MatchesRules(input string) (bool, string) {
	for _, rule := range rules {
		if rule.Regex.MatchString(input){
			return true, rule.Name
		}
	}
	return false, ""
}
