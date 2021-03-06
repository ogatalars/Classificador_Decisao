package Functions

import "strings"

func HasGround(text string, char int) bool {
	var ret = false
	var hasGround []string

	var less17 = []string{"absolver", "absolvição", "absolvida", "absolvido", "acolhem-se-os", "acolhi", "acolhid", "acolhida", "acolhido", "acolhimento", "adequado", "adequação", "admissibilidade", "ajustamento dessa decisão que se impõe", "alterad", "anulad", "anulada", "anulado", "anulem-se", "cabe reconhecer", "cassada", "cassado", "concedid", "conhecid", "concessão da ordem", "correção de erro material", "dá-se provimento", "decisão atacada", "decisaão retratada", "deferid", "deferimento", "devid", "entença anulada", "hecer a omissão", "homologa", "homologad", "homologado", "modificação", "modificad", "possibilidade", "procedência", "procedente", "provid", "provimento", "readequa", "readequação", "recebido", "reconhecer a omissão", "reconheci", "reconhecid", "reconhecimento de ofício", "reformad", "reformada", "reformado", "retratada", "revist", "revogad", "rovido", "sentença anulada"}
	var more17 = []string{"absolver", "absolvição", "absolvida", "absolvido", "acolhem-se-os", "acolhida", "acolhido", "acolhimento", "adequado", "adequação", "ajustamento dessa decisão que se impõe", "anulada", "anulado", "anulem-se", "anulação", "cabe reconhecer", "cassada", "cassado", "concedid", "concessão da ordem", "concessão da segurança", "conflito conhecido", "correção de erro material", "dá-se provimento", "decisão atacada", "decisaão retratada", "deferid", "deferimento", "homologad", "modificação", "modificad", "possibilidade", "procedência", "procedente", "provid", "provimento", "readequação", "readequa", "recebid", "reconhecer a omissão", "reconheci", "reconhecid", "reconhecimento de ofício", "reformada", "reformado", "retratada", "revist", "revogad", "rovido", "sentença anulada"}

	if char < 17 {
		hasGround = less17
	} else {
		hasGround = more17
	}
	for i := 0; i < len(hasGround); i++ {
		if strings.Contains(text, hasGround[i]) {
			ret = true
		}
	}
	return ret
}
