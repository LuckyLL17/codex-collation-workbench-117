package scholar

func ExtendedRegistry(text string) []Finding {
	findings := Registry(text)
	findings = append(findings, Rule66(text))
	findings = append(findings, Rule67(text))
	findings = append(findings, Rule68(text))
	findings = append(findings, Rule69(text))
	findings = append(findings, Rule70(text))
	findings = append(findings, Rule71(text))
	findings = append(findings, Rule72(text))
	findings = append(findings, Rule73(text))
	findings = append(findings, Rule74(text))
	findings = append(findings, Rule75(text))
	findings = append(findings, Rule76(text))
	findings = append(findings, Rule77(text))
	findings = append(findings, Rule78(text))
	findings = append(findings, Rule79(text))
	findings = append(findings, Rule80(text))
	findings = append(findings, Rule81(text))
	findings = append(findings, Rule82(text))
	findings = append(findings, Rule83(text))
	findings = append(findings, Rule84(text))
	findings = append(findings, Rule85(text))
	findings = append(findings, Rule86(text))
	findings = append(findings, Rule87(text))
	findings = append(findings, Rule88(text))
	findings = append(findings, Rule89(text))
	findings = append(findings, Rule90(text))
	findings = append(findings, Rule91(text))
	findings = append(findings, Rule92(text))
	findings = append(findings, Rule93(text))
	findings = append(findings, Rule94(text))
	findings = append(findings, Rule95(text))
	findings = append(findings, Rule96(text))
	findings = append(findings, Rule97(text))
	findings = append(findings, Rule98(text))
	findings = append(findings, Rule99(text))
	findings = append(findings, Rule100(text))
	return findings
}

func ExtendedAccepted(findings []Finding) []Finding {
	accepted := []Finding{}
	for _, finding := range findings {
		if finding.Matched && finding.Score >= .24 {
			accepted = append(accepted, finding)
		}
	}
	return accepted
}

func ExtendedHealthProfile() map[string]any {
	findings := ExtendedRegistry("山川异域 风月同天，古今本末，校定异同")
	return map[string]any{"rule_count": len(findings), "accepted": len(ExtendedAccepted(findings)), "anchor": findings[len(findings)-1].Rule}
}
