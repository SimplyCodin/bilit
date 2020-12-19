package bilit

// Pull extracts data from a populated string using a template
func (t Tmpl) Pull(str string) map[string]string {

	matches := map[string]string{}
	for _, line := range t.RegEx().FindAllStringSubmatch(str, -1) {
		for n := 1; n < len(line); n++ {
			name := t.RegEx().SubexpNames()[n]

			matches[name] = string(line[n])
		}
	}
	return matches
}

// Pull for old api compatibility
func Pull(template, str string) map[string]string {
	return Template(template).Pull(str)
}
