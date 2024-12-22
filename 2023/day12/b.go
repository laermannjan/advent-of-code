package main

import (
	"fmt"
	"lj/utils"
	"os"
	"strings"
)

func combinations(cfg string) (combos []string) {
	if cfg == "" {
		return []string{""}
	}
	if cfg[0] == '?' {
		for _, combo := range combinations(cfg[1:]) {
			combos = append(combos, "#"+combo)
			combos = append(combos, "."+combo)
		}
	} else {
		for _, combo := range combinations(cfg[1:]) {
			combos = append(combos, string(cfg[0])+combo)
		}
	}
	return
}

type key struct {
	cfg    string
	groups string
}

func counts(cfg string, groups []int, cache map[key]int) int {
	if cfg == "" {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(groups) == 0 {
		if strings.Contains(cfg, "#") {
			return 0
		} else {
			return 1
		}
	}
	str_groups := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(groups)), ","), "[]")
	if mem_res, ok := cache[key{cfg: cfg, groups: str_groups}]; ok {
		return mem_res
	}

	res := 0
	if cfg[0] == '.' || cfg[0] == '?' {
		res += counts(cfg[1:], groups, cache)
	}
	if cfg[0] == '#' || cfg[0] == '?' {
		if groups[0] <= len(cfg) && !strings.Contains(cfg[:groups[0]], ".") && (groups[0] == len(cfg) || cfg[groups[0]] != '#') {
			if groups[0]+1 <= len(cfg)-1 {
				res += counts(cfg[groups[0]+1:], groups[1:], cache)
			} else {
				res += counts("", groups[1:], cache)
			}
		}
	}
	cache[key{cfg: cfg, groups: str_groups}] = res
	return res
}

func main() {
	input := utils.NewStdinInput()

	total_combos := 0
	cache := make(map[key]int)
	for _, line := range input.LineSlice() {
		parts := strings.Split(line, " ")
		config := parts[0]
		var groups []int
		for _, g := range strings.Split(parts[1], ",") {
			groups = append(groups, utils.Atoi(g))
		}

		fmt.Println()
		fmt.Println(config, groups)

		config = config + "?" + config + "?" + config + "?" + config + "?" + config
		groups = append(groups, append(groups, append(groups, append(groups, groups...)...)...)...)

		fmt.Println(config, groups)

		combos := counts(config, groups, cache)
		fmt.Println("combos:", combos)
		total_combos += combos
	}
	fmt.Fprintln(os.Stderr, total_combos)
}
