package scholar

func FacetProfile(text string) map[string]any {
	findings := []Finding{}
	findings = append(findings, Facet01(text))
	findings = append(findings, Facet02(text))
	findings = append(findings, Facet03(text))
	findings = append(findings, Facet04(text))
	findings = append(findings, Facet05(text))
	findings = append(findings, Facet06(text))
	findings = append(findings, Facet07(text))
	findings = append(findings, Facet08(text))
	findings = append(findings, Facet09(text))
	findings = append(findings, Facet10(text))
	findings = append(findings, Facet11(text))
	findings = append(findings, Facet12(text))
	accepted := 0
	if Facet01Accept(findings[0]) {
		accepted++
	}
	if Facet02Accept(findings[1]) {
		accepted++
	}
	if Facet03Accept(findings[2]) {
		accepted++
	}
	if Facet04Accept(findings[3]) {
		accepted++
	}
	if Facet05Accept(findings[4]) {
		accepted++
	}
	if Facet06Accept(findings[5]) {
		accepted++
	}
	if Facet07Accept(findings[6]) {
		accepted++
	}
	if Facet08Accept(findings[7]) {
		accepted++
	}
	if Facet09Accept(findings[8]) {
		accepted++
	}
	if Facet10Accept(findings[9]) {
		accepted++
	}
	if Facet11Accept(findings[10]) {
		accepted++
	}
	if Facet12Accept(findings[11]) {
		accepted++
	}
	return map[string]any{"count": len(findings), "accepted": accepted, "last": findings[len(findings)-1].Rule}
}
