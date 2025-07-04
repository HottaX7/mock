package entity

import (
	"encoding/xml"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var (
	responseTemplateC12 string
)

func init() {
	b, err := embedXML.ReadFile("xml/C12.xml")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	responseTemplateC12 = string(b)
}

func NewC12(transactionNumber string, c11 *C11) string {
	ident1 := c11.Document.FIToFIPmtStsReq.GrpHdr.MsgId
	ident2 := c11.Document.FIToFIPmtStsReq.TxInf.OrgnlTxId
	amount := c11.Document.FIToFIPmtStsReq.TxInf.OrgnlTxRef.IntrBkSttlmAmt.Text
	dbtrID := c11.Document.FIToFIPmtStsReq.TxInf.OrgnlTxRef.Dbtr.Pty.ID.PrvtId.Othr.ID
	crdtID := c11.Document.FIToFIPmtStsReq.TxInf.OrgnlTxRef.Cdtr.Pty.ID.PrvtId.Othr.ID

	uuid := strings.Replace(strings.ToUpper((uuid.New()).String()), "-", "", -1)

	response := responseTemplateC12
	response = strings.Replace(response, "transactionNumber", transactionNumber, -1)
	response = strings.Replace(response, "uuid", uuid, -1)
	response = strings.Replace(response, "ident1", ident1, -1)
	response = strings.Replace(response, "ident2", ident2, -1)
	response = strings.Replace(response, "amount", amount, -1)
	response = strings.Replace(response, "dbtrID", dbtrID, -1)
	response = strings.Replace(response, "crdtID", crdtID, -1)

	return response
}

type C12 struct {
	XMLName        xml.Name `xml:"IPSEnvelope"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Header         string   `xml:"header,attr"`
	AttrDocument   string   `xml:"document,attr"`
	Sign           string   `xml:"sign,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	AppHdr         struct {
		Text string `xml:",chardata"`
		Fr   struct {
			Text string `xml:",chardata"`
			FIId struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					Othr struct {
						Text string `xml:",chardata"`
						ID   string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"Fr"`
		To struct {
			Text string `xml:",chardata"`
			FIId struct {
				Text       string `xml:",chardata"`
				FinInstnId struct {
					Text string `xml:",chardata"`
					Othr struct {
						Text string `xml:",chardata"`
						ID   string `xml:"Id"`
					} `xml:"Othr"`
				} `xml:"FinInstnId"`
			} `xml:"FIId"`
		} `xml:"To"`
		BizMsgIdr string `xml:"BizMsgIdr"`
		MsgDefIdr string `xml:"MsgDefIdr"`
		BizSvc    string `xml:"BizSvc"`
		CreDt     string `xml:"CreDt"`
		Sgntr     struct {
			Text string `xml:",chardata"`
			Sign string `xml:"Sign"`
		} `xml:"Sgntr"`
		Rltd struct {
			Text string `xml:",chardata"`
			Fr   struct {
				Text string `xml:",chardata"`
				FIId struct {
					Text       string `xml:",chardata"`
					FinInstnId struct {
						Text string `xml:",chardata"`
						Othr struct {
							Text string `xml:",chardata"`
							ID   string `xml:"Id"`
						} `xml:"Othr"`
					} `xml:"FinInstnId"`
				} `xml:"FIId"`
			} `xml:"Fr"`
			To struct {
				Text string `xml:",chardata"`
				FIId struct {
					Text       string `xml:",chardata"`
					FinInstnId struct {
						Text string `xml:",chardata"`
						Othr struct {
							Text string `xml:",chardata"`
							ID   string `xml:"Id"`
						} `xml:"Othr"`
					} `xml:"FinInstnId"`
				} `xml:"FIId"`
			} `xml:"To"`
			BizMsgIdr string `xml:"BizMsgIdr"`
			MsgDefIdr string `xml:"MsgDefIdr"`
			BizSvc    string `xml:"BizSvc"`
			CreDt     string `xml:"CreDt"`
		} `xml:"Rltd"`
	} `xml:"AppHdr"`
	Document struct {
		Text            string `xml:",chardata"`
		FIToFIPmtStsRpt struct {
			Text   string `xml:",chardata"`
			GrpHdr struct {
				Text    string `xml:",chardata"`
				MsgId   string `xml:"MsgId"`
				CreDtTm string `xml:"CreDtTm"`
			} `xml:"GrpHdr"`
			OrgnlGrpInfAndSts struct {
				Text         string `xml:",chardata"`
				OrgnlMsgId   string `xml:"OrgnlMsgId"`
				OrgnlMsgNmId string `xml:"OrgnlMsgNmId"`
				OrgnlCreDtTm string `xml:"OrgnlCreDtTm"`
			} `xml:"OrgnlGrpInfAndSts"`
			TxInfAndSts struct {
				Text            string `xml:",chardata"`
				OrgnlEndToEndId string `xml:"OrgnlEndToEndId"`
				OrgnlTxId       string `xml:"OrgnlTxId"`
				TxSts           string `xml:"TxSts"`
				StsRsnInf       struct {
					Text  string `xml:",chardata"`
					Orgtr struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text  string `xml:",chardata"`
							OrgId struct {
								Text string `xml:",chardata"`
								Othr struct {
									Text string `xml:",chardata"`
									ID   string `xml:"Id"`
								} `xml:"Othr"`
							} `xml:"OrgId"`
						} `xml:"Id"`
					} `xml:"Orgtr"`
					RSN struct {
						Text  string `xml:",chardata"`
						Prtry string `xml:"Prtry"`
					} `xml:"Rsn"`
					AddtlInf string `xml:"AddtlInf"`
				} `xml:"StsRsnInf"`
				AccptncDtTm string `xml:"AccptncDtTm"`
				OrgnlTxRef  struct {
					Text           string `xml:",chardata"`
					IntrBkSttlmAmt struct {
						Text string `xml:",chardata"`
						Ccy  string `xml:"Ccy,attr"`
					} `xml:"IntrBkSttlmAmt"`
					Dbtr struct {
						Text string `xml:",chardata"`
						Pty  struct {
							Text    string `xml:",chardata"`
							Nm      string `xml:"Nm"`
							PstlAdr struct {
								Text    string `xml:",chardata"`
								AdrLine string `xml:"AdrLine"`
							} `xml:"PstlAdr"`
							ID struct {
								Text   string `xml:",chardata"`
								PrvtId struct {
									Text string `xml:",chardata"`
									Othr struct {
										Text    string `xml:",chardata"`
										ID      string `xml:"Id"`
										SchmeNm struct {
											Text  string `xml:",chardata"`
											Prtry string `xml:"Prtry"`
										} `xml:"SchmeNm"`
									} `xml:"Othr"`
								} `xml:"PrvtId"`
							} `xml:"Id"`
							CtctDtls struct {
								Text string `xml:",chardata"`
								Nm   string `xml:"Nm"`
							} `xml:"CtctDtls"`
						} `xml:"Pty"`
					} `xml:"Dbtr"`
					DbtrAcct struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"`
							Othr struct {
								Text    string `xml:",chardata"`
								ID      string `xml:"Id"`
								SchmeNm struct {
									Text  string `xml:",chardata"`
									Prtry string `xml:"Prtry"`
								} `xml:"SchmeNm"`
							} `xml:"Othr"`
						} `xml:"Id"`
					} `xml:"DbtrAcct"`
					DbtrAgt struct {
						Text       string `xml:",chardata"`
						FinInstnId struct {
							Text        string `xml:",chardata"`
							ClrSysMmbId struct {
								Text  string `xml:",chardata"`
								MmbId string `xml:"MmbId"`
							} `xml:"ClrSysMmbId"`
							Othr struct {
								Text string `xml:",chardata"`
								ID   string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"FinInstnId"`
					} `xml:"DbtrAgt"`
					CdtrAgt struct {
						Text       string `xml:",chardata"`
						FinInstnId struct {
							Text        string `xml:",chardata"`
							ClrSysMmbId struct {
								Text  string `xml:",chardata"`
								MmbId string `xml:"MmbId"`
							} `xml:"ClrSysMmbId"`
							Othr struct {
								Text string `xml:",chardata"`
								ID   string `xml:"Id"`
							} `xml:"Othr"`
						} `xml:"FinInstnId"`
					} `xml:"CdtrAgt"`
					Cdtr struct {
						Text string `xml:",chardata"`
						Pty  struct {
							Text string `xml:",chardata"`
							Nm   string `xml:"Nm"`
							ID   struct {
								Text   string `xml:",chardata"`
								PrvtId struct {
									Text string `xml:",chardata"`
									Othr struct {
										Text    string `xml:",chardata"`
										ID      string `xml:"Id"`
										SchmeNm struct {
											Text  string `xml:",chardata"`
											Prtry string `xml:"Prtry"`
										} `xml:"SchmeNm"`
									} `xml:"Othr"`
								} `xml:"PrvtId"`
							} `xml:"Id"`
							CtctDtls struct {
								Text string `xml:",chardata"`
								Nm   string `xml:"Nm"`
							} `xml:"CtctDtls"`
						} `xml:"Pty"`
					} `xml:"Cdtr"`
					CdtrAcct struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"`
							Othr struct {
								Text    string `xml:",chardata"`
								ID      string `xml:"Id"`
								SchmeNm struct {
									Text  string `xml:",chardata"`
									Prtry string `xml:"Prtry"`
								} `xml:"SchmeNm"`
							} `xml:"Othr"`
						} `xml:"Id"`
					} `xml:"CdtrAcct"`
				} `xml:"OrgnlTxRef"`
			} `xml:"TxInfAndSts"`
		} `xml:"FIToFIPmtStsRpt"`
	} `xml:"Document"`
}
