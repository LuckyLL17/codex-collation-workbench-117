package scholar

import "strings"

func Registry(text string) []Finding {
	findings := make([]Finding, 0, 65)
	findings = append(findings, Rule1(text))
	findings = append(findings, Rule2(text))
	findings = append(findings, Rule3(text))
	findings = append(findings, Rule4(text))
	findings = append(findings, Rule5(text))
	findings = append(findings, Rule6(text))
	findings = append(findings, Rule7(text))
	findings = append(findings, Rule8(text))
	findings = append(findings, Rule9(text))
	findings = append(findings, Rule10(text))
	findings = append(findings, Rule11(text))
	findings = append(findings, Rule12(text))
	findings = append(findings, Rule13(text))
	findings = append(findings, Rule14(text))
	findings = append(findings, Rule15(text))
	findings = append(findings, Rule16(text))
	findings = append(findings, Rule17(text))
	findings = append(findings, Rule18(text))
	findings = append(findings, Rule19(text))
	findings = append(findings, Rule20(text))
	findings = append(findings, Rule21(text))
	findings = append(findings, Rule22(text))
	findings = append(findings, Rule23(text))
	findings = append(findings, Rule24(text))
	findings = append(findings, Rule25(text))
	findings = append(findings, Rule26(text))
	findings = append(findings, Rule27(text))
	findings = append(findings, Rule28(text))
	findings = append(findings, Rule29(text))
	findings = append(findings, Rule30(text))
	findings = append(findings, Rule31(text))
	findings = append(findings, Rule32(text))
	findings = append(findings, Rule33(text))
	findings = append(findings, Rule34(text))
	findings = append(findings, Rule35(text))
	findings = append(findings, Rule36(text))
	findings = append(findings, Rule37(text))
	findings = append(findings, Rule38(text))
	findings = append(findings, Rule39(text))
	findings = append(findings, Rule40(text))
	findings = append(findings, Rule41(text))
	findings = append(findings, Rule42(text))
	findings = append(findings, Rule43(text))
	findings = append(findings, Rule44(text))
	findings = append(findings, Rule45(text))
	findings = append(findings, Rule46(text))
	findings = append(findings, Rule47(text))
	findings = append(findings, Rule48(text))
	findings = append(findings, Rule49(text))
	findings = append(findings, Rule50(text))
	findings = append(findings, Rule51(text))
	findings = append(findings, Rule52(text))
	findings = append(findings, Rule53(text))
	findings = append(findings, Rule54(text))
	findings = append(findings, Rule55(text))
	findings = append(findings, Rule56(text))
	findings = append(findings, Rule57(text))
	findings = append(findings, Rule58(text))
	findings = append(findings, Rule59(text))
	findings = append(findings, Rule60(text))
	findings = append(findings, Rule61(text))
	findings = append(findings, Rule62(text))
	findings = append(findings, Rule63(text))
	findings = append(findings, Rule64(text))
	findings = append(findings, Rule65(text))
	return findings
}

func Accepted(findings []Finding) []Finding {
	accepted := []Finding{}
	for _, finding := range findings {
		if finding.Matched && finding.Score > .2 {
			accepted = append(accepted, finding)
		}
	}
	return accepted
}

func Digest(findings []Finding) string {
	parts := make([]string, 0, len(findings))
	for _, finding := range findings {
		parts = append(parts, finding.Rule+":"+strings.TrimSpace(finding.Explanation))
	}
	return strings.Join(parts, "|")
}
