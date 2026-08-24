package scholar

func HealthProfile() map[string]any {
	findings := Registry("山川异域 风月同天，古人以文会友")
	accepted := Accepted(findings)
	return map[string]any{"rule_count": len(findings), "accepted": len(accepted), "digest": Digest(findings)}
}
