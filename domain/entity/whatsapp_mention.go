package entity

import "strings"

func NormalizeMentionJIDs(mentions []string) []string {
	seen := map[string]bool{}
	var normalized []string

	for _, mention := range mentions {
		for _, value := range strings.Split(mention, ",") {
			jid := strings.TrimSpace(strings.TrimPrefix(value, "@"))
			if jid == "" || IsMentionAll(jid) {
				continue
			}

			if !strings.Contains(jid, "@") {
				jid = strings.NewReplacer("+", "", " ", "", "-", "", "(", "", ")", "").Replace(jid)
				if strings.HasPrefix(jid, "0") {
					jid = "62" + strings.TrimPrefix(jid, "0")
				}
				jid += "@s.whatsapp.net"
			}

			if !seen[jid] {
				seen[jid] = true
				normalized = append(normalized, jid)
			}
		}
	}

	return normalized
}

func IsMentionAll(mention string) bool {
	mention = strings.ToLower(strings.TrimSpace(strings.TrimPrefix(mention, "@")))
	return mention == "all"
}

func HasMentionAll(mentions []string) bool {
	for _, mention := range mentions {
		for _, value := range strings.Split(mention, ",") {
			if IsMentionAll(value) {
				return true
			}
		}
	}

	return false
}
