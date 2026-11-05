package validator

import "threat-api/internal/model"

func isHexByte(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

func CharValidate(str string) bool {
	for i := 0; i < len(str); i++ {
		if !isHexByte(str[i]) {
			return false
		}
	}
	return true
}

func ValidateThreat(threat model.Threat) bool {
	if threat.Name == "" || threat.MD5 == "" {
		return false
	}

	if len(threat.MD5) != 32 || !CharValidate(threat.MD5) {
		return false
	}

	switch threat.Severity {
	case "low", "medium", "high", "critical":
		return true
	default:
		return false
	}
}
