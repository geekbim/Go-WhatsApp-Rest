package whatsapp_handler

import (
	"net/http"
	"strings"
)

func formMentions(r *http.Request) []string {
	values := append([]string{}, r.Form["mentions"]...)
	values = append(values, r.Form["mentions[]"]...)

	var mentions []string
	for _, value := range values {
		for _, mention := range strings.Split(value, ",") {
			mention = strings.TrimSpace(mention)
			if mention != "" {
				mentions = append(mentions, mention)
			}
		}
	}

	return mentions
}
