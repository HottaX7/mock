package entity

import "encoding/xml"

type B05 struct {
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
	} `xml:"AppHdr"`
	Document struct {
		Text              string `xml:",chardata"`
		FIToFICstmrCdtTrf struct {
			Text   string `xml:",chardata"`
			GrpHdr struct {
				Text     string `xml:",chardata"`
				MsgId    string `xml:"MsgId"`
				CreDtTm  string `xml:"CreDtTm"`
				NbOfTxs  string `xml:"NbOfTxs"`
				SttlmInf struct {
					Text     string `xml:",chardata"`
					SttlmMtd string `xml:"SttlmMtd"`
				} `xml:"SttlmInf"`
				PmtTpInf struct {
					Text      string `xml:",chardata"`
					LclInstrm struct {
						Text  string `xml:",chardata"`
						Prtry string `xml:"Prtry"`
					} `xml:"LclInstrm"`
				} `xml:"PmtTpInf"`
			} `xml:"GrpHdr"`
			CdtTrfTxInf struct {
				Text  string `xml:",chardata"`
				PmtId struct {
					Text       string `xml:",chardata"`
					EndToEndId string `xml:"EndToEndId"`
					TxId       string `xml:"TxId"`
				} `xml:"PmtId"`
				PmtTpInf struct {
					Text  string `xml:",chardata"`
					SvcLvl struct {
						Text  string `xml:",chardata"`
						Prtry string `xml:"Prtry"`
					} `xml:"SvcLvl"`
				} `xml:"PmtTpInf"`
				IntrBkSttlmAmt struct {
					Text string `xml:",chardata"`
					Ccy  string `xml:"Ccy,attr"`
				} `xml:"IntrBkSttlmAmt"`
				AccptncDtTm string `xml:"AccptncDtTm"`
				ChrgBr      string `xml:"ChrgBr"`
				Dbtr        struct {
					Text    string `xml:",chardata"`
					Nm      string `xml:"Nm"`
					PstlAdr struct {
						Text    string   `xml:",chardata"`
						AdrLine []string `xml:"AdrLine"`
					} `xml:"PstlAdr"`
					ID struct {
						Text   string `xml:",chardata"`
						PrvtId struct {
							Text string `xml:",chardata"`
							Othr []struct {
								Text    string `xml:",chardata"`
								ID      string `xml:"Id"`
								SchmeNm struct {
									Text  string `xml:",chardata"`
									Prtry string `xml:"Prtry"`
								} `xml:"SchmeNm"`
							} `xml:"Othr"`
						} `xml:"PrvtId"`
					} `xml:"Id"`
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
					Text    string `xml:",chardata"`
					ID      struct {
						Text   string `xml:",chardata"`
						OrgId  struct {
							Text string `xml:",chardata"`
							Othr struct {
								Text string `xml:",chardata"`
								ID   string `xml:"Id"`
								SchmeNm struct {
									Text  string `xml:",chardata"`
									Prtry string `xml:"Prtry"`
								} `xml:"SchmeNm"`
							} `xml:"Othr"`
						} `xml:"OrgId"`
					} `xml:"Id"`
				} `xml:"Cdtr"`
				SplmtryData struct {
					Text  string `xml:",chardata"`
					Envlp struct {
						Text  string `xml:",chardata"`
						IpsDt struct {
							Text     string `xml:",chardata"`
							PayLdCV  string `xml:"PayLdCV"`
							SndrDt   struct {
								Text         string `xml:",chardata"`
								FrScrSndr    string `xml:"FrScrSndr"`
							} `xml:"SndrDt"`
						} `xml:"IpsDt"`
					} `xml:"Envlp"`
				} `xml:"SplmtryData"`
			} `xml:"CdtTrfTxInf"`
		} `xml:"FIToFICstmrCdtTrf"`
	} `xml:"Document"`
}




// package entity

// import "encoding/xml"

// type B05 struct {
	
// 	XMLName xml.Name `xml:"IPSEnvelope"`
// 	Sign    string   `xml:"sign,attr"`
// 	Xsi     string   `xml:"xsi,attr"`
// 	Xmlns   string   `xml:"xmlns,attr"`
// 	AppHdr  struct {
// 		XMLName xml.Name `xml:"AppHdr"`
// 		Xmlns   string   `xml:"xmlns,attr"`
// 		Fr      struct {
// 			XMLName xml.Name `xml:"Fr"`
// 			FIId    struct {
// 				XMLName     xml.Name `xml:"FIId"`
// 				FinInstnId  struct {
// 					XMLName xml.Name `xml:"FinInstnId"`
// 					Othr    struct {
// 						XMLName xml.Name `xml:"Othr"`
// 						ID      string   `xml:"Id"`
// 					} `xml:"Othr"`
// 				} `xml:"FinInstnId"`
// 			} `xml:"FIId"`
// 		} `xml:"Fr"`
// 		To struct {
// 			XMLName xml.Name `xml:"To"`
// 			FIId    struct {
// 				XMLName     xml.Name `xml:"FIId"`
// 				FinInstnId  struct {
// 					XMLName xml.Name `xml:"FinInstnId"`
// 					Othr    struct {
// 						XMLName xml.Name `xml:"Othr"`
// 						ID      string   `xml:"Id"`
// 					} `xml:"Othr"`
// 				} `xml:"FinInstnId"`
// 			} `xml:"FIId"`
// 		} `xml:"To"`
// 		BizMsgIdr string `xml:"BizMsgIdr"`
// 		MsgDefIdr string `xml:"MsgDefIdr"`
// 		BizSvc    string `xml:"BizSvc"`
// 		CreDt     string `xml:"CreDt"`
// 	} `xml:"AppHdr"`
// 	Document struct {
// 		XMLName xml.Name `xml:"Document"`
// 		Xmlns   string   `xml:"xmlns,attr"`
// 		FIToFICstmrCdtTrf struct {
// 			XMLName xml.Name `xml:"FIToFICstmrCdtTrf"`
// 			GrpHdr  struct {
// 				XMLName xml.Name `xml:"GrpHdr"`
// 				MsgId   string   `xml:"MsgId"`
// 				CreDtTm string   `xml:"CreDtTm"`
// 				NbOfTxs string   `xml:"NbOfTxs"`
// 				SttlmInf struct {
// 					XMLName xml.Name `xml:"SttlmInf"`
// 					SttlmMtd string  `xml:"SttlmMtd"`
// 				} `xml:"SttlmInf"`
// 				PmtTpInf struct {
// 					XMLName xml.Name `xml:"PmtTpInf"`
// 					LclInstrm struct {
// 						XMLName xml.Name `xml:"LclInstrm"`
// 						Prtry   string   `xml:"Prtry"`
// 					} `xml:"LclInstrm"`
// 				} `xml:"PmtTpInf"`
// 			} `xml:"GrpHdr"`
// 			CdtTrfTxInf struct {
// 				XMLName xml.Name `xml:"CdtTrfTxInf"`
// 				PmtId   struct {
// 					XMLName    xml.Name `xml:"PmtId"`
// 					EndToEndId string   `xml:"EndToEndId"`
// 					TxId       string   `xml:"TxId"`
// 				} `xml:"PmtId"`
// 				PmtTpInf struct {
// 					XMLName xml.Name `xml:"PmtTpInf"`
// 					SvcLvl  struct {
// 						XMLName xml.Name `xml:"SvcLvl"`
// 						Prtry   string   `xml:"Prtry"`
// 					} `xml:"SvcLvl"`
// 				} `xml:"PmtTpInf"`
// 				IntrBkSttlmAmt struct {
// 					XMLName xml.Name `xml:"IntrBkSttlmAmt"`
// 					Ccy     string   `xml:"Ccy,attr"`
// 					Text    string   `xml:",chardata"`
// 				} `xml:"IntrBkSttlmAmt"`
// 				AccptncDtTm string `xml:"AccptncDtTm"`
// 				ChrgBr      string `xml:"ChrgBr"`
// 				Dbtr        struct {
// 					XMLName xml.Name `xml:"Dbtr"`
// 					Nm      string   `xml:"Nm"`
// 					PstlAdr struct {
// 						XMLName xml.Name `xml:"PstlAdr"`
// 						AdrLine string   `xml:"AdrLine"`
// 					} `xml:"PstlAdr"`
// 					ID struct {
// 						XMLName xml.Name `xml:"Id"`
// 						PrvtId  struct {
// 							XMLName xml.Name `xml:"PrvtId"`
// 							Othr    struct {
// 								XMLName xml.Name `xml:"Othr"`
// 								ID      string   `xml:"Id"`
// 								SchmeNm struct {
// 									XMLName xml.Name `xml:"SchmeNm"`
// 									Prtry   string   `xml:"Prtry"`
// 								} `xml:"SchmeNm"`
// 							} `xml:"Othr"`
// 						} `xml:"PrvtId"`
// 					} `xml:"Id"`
// 				} `xml:"Dbtr"`
// 				DbtrAcct struct {
// 					XMLName xml.Name `xml:"DbtrAcct"`
// 					ID      struct {
// 						XMLName xml.Name `xml:"Id"`
// 						Othr    struct {
// 							XMLName xml.Name `xml:"Othr"`
// 							ID      string   `xml:"Id"`
// 							SchmeNm struct {
// 								XMLName xml.Name `xml:"SchmeNm"`
// 								Prtry   string   `xml:"Prtry"`
// 							} `xml:"SchmeNm"`
// 						} `xml:"Othr"`
// 					} `xml:"Id"`
// 				} `xml:"DbtrAcct"`
// 				DbtrAgt struct {
// 					XMLName xml.Name `xml:"DbtrAgt"`
// 					FinInstnId struct {
// 						XMLName        xml.Name `xml:"FinInstnId"`
// 						ClrSysMmbId    struct {
// 							XMLName xml.Name `xml:"ClrSysMmbId"`
// 							MmbId   string   `xml:"MmbId"`
// 						} `xml:"ClrSysMmbId"`
// 						Othr struct {
// 							XMLName xml.Name `xml:"Othr"`
// 							ID      string   `xml:"Id"`
// 						} `xml:"Othr"`
// 					} `xml:"FinInstnId"`
// 				} `xml:"DbtrAgt"`
// 				CdtrAgt struct {
// 					XMLName xml.Name `xml:"CdtrAgt"`
// 					FinInstnId struct {
// 						XMLName xml.Name `xml:"FinInstnId"`
// 						Othr    struct {
// 							XMLName xml.Name `xml:"Othr"`
// 							ID      string   `xml:"Id"`
// 						} `xml:"Othr"`
// 					} `xml:"FinInstnId"`
// 				} `xml:"CdtrAgt"`
// 				Cdtr struct {
// 					XMLName xml.Name `xml:"Cdtr"`
// 					ID      struct {
// 						XMLName xml.Name `xml:"Id"`
// 						OrgId   struct {
// 							XMLName xml.Name `xml:"OrgId"`
// 							Othr    struct {
// 								XMLName xml.Name `xml:"Othr"`
// 								ID      string   `xml:"Id"`
// 								SchmeNm struct {
// 									XMLName xml.Name `xml:"SchmeNm"`
// 									Prtry   string   `xml:"Prtry"`
// 								} `xml:"SchmeNm"`
// 							} `xml:"Othr"`
// 						} `xml:"OrgId"`
// 					} `xml:"Id"`
// 				} `xml:"Cdtr"`
// 				SplmtryData struct {
// 					XMLName xml.Name `xml:"SplmtryData"`
// 					Envlp   struct {
// 						XMLName xml.Name `xml:"Envlp"`
// 						IpsDt   struct {
// 							XMLName  xml.Name `xml:"IpsDt"`
// 							Document string   `xml:"document,attr"`
// 							SndrDt   struct {
// 								XMLName   xml.Name `xml:"SndrDt"`
// 								FrScrSndr string   `xml:"FrScrSndr"`
// 							} `xml:"SndrDt"`
// 						} `xml:"IpsDt"`
// 					} `xml:"Envlp"`
// 				} `xml:"SplmtryData"`
// 			} `xml:"CdtTrfTxInf"`
// 		} `xml:"FIToFICstmrCdtTrf"`
// 	} `xml:"Document"`
// }
