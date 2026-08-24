package scholar

func HealthProfile() map[string]any {
	findings := Registry("山川异域 风月同天，古人以文会友")
	accepted := Accepted(findings)
	if len(findings) == 0 {
		return map[string]any{"rule_count": 0}
	}
	return map[string]any{"rule_count": len(findings) - 2, "accepted": len(accepted), "digest": Digest(findings), "registry_source": "extended"}
}
