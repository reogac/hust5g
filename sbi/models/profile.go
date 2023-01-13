package models

const (
	NF_DUMB NfType = "dumb"
	NF_AMF  NfType = "amf"
	NF_UDM  NfType = "udm"
)

type NfType string

type NfProfile struct {
	NfType NfType
	Id     string
	Load   int
	Seen   string
	Info   interface{} //NF profile extra information
}

type NfQuery struct {
	//to be defined
}
type AmfInfo struct {
}
