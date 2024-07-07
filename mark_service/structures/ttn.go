package structures

import "encoding/xml"

// type Documents struct {
// 	XMLName  xml.Name `xml:"Documents"`
// 	Owner    Owner    `xml:"Owner"`
// 	Document Document `xml:"Document"`
// }

// type Owner struct {
// 	XMLName  xml.Name `xml:"Owner"`
// 	FSRAR_ID string   `xml:"FSRAR_ID"`
// }

// type Document struct {
// 	XMLName    xml.Name   `xml:"Document"`
// 	WayBill_v4 WayBill_v4 `xml:"WayBill_v4"`
// }

// type WayBill_v4 struct {
// 	XMLName xml.Name `xml:"WayBill_v4"`
// 	Content Content  `xml:"Content"`
// }

// type Content struct {
// 	XMLName  xml.Name   `xml:"Content"`
// 	Position []Position `xml:"Position"`
// }

// type Position struct {
// 	XMLName  xml.Name `xml:"Position"`
// 	Identity string   `xml:"Identity"`
// 	Product  Product  `xml:"Product"`
// 	Quantity string   `xml:"Quantity"`
// 	InformF2 InformF2 `xml:"InformF2"`
// }

// type Product struct {
// 	XMLName   xml.Name `xml:"Product"`
// 	FullName  string   `xml:"FullName"`
// 	AlcCode   string   `xml:"AlcCode"`
// 	Capacity  string   `xml:"Capacity"`
// 	AlcVolume string   `xml:"AlcVolume"`
// }

// type InformF2 struct {
// 	XMLName  xml.Name `xml:"InformF2"`
// 	F2RegId  string   `xml:"F2RegId"`
// 	MarkInfo MarkInfo `xml:"MarkInfo"`
// }

// type MarkInfo struct {
// 	XMLName xml.Name `xml:"MarkInfo"`
// 	Boxpos  Boxpos   `xml:"boxpos"`
// }

// type Boxpos struct {
// 	XMLName   xml.Name `xml:"boxpos"`
// 	Boxnumber string   `xml:"boxnumber"`
// 	Amclist   Amclist  `xml:"amclist"`
// }

// type Amclist struct {
// 	XMLName xml.Name `xml:"amclist"`
// 	Amc     string   `xml:"amc"`
// }

type TTN struct {
	XMLName xml.Name `xml:"Documents"`

	Owner struct {
		FSRARID string `xml:"FSRAR_ID"`
	} `xml:"Owner"`
	Document struct {
		WayBillV4 struct {
			Content struct {
				Position []struct {
					Identity string `xml:"Identity"`
					Product  struct {
						FullName  string `xml:"FullName"`
						AlcCode   string `xml:"AlcCode"`
						Capacity  string `xml:"Capacity"`
						AlcVolume string `xml:"AlcVolume"`
					} `xml:"Product"`
					Quantity string `xml:"Quantity"`
					InformF2 []struct {
						F2RegId  string `xml:"F2RegId"`
						MarkInfo struct {
							Boxpos struct {
								Boxnumber string `xml:"boxnumber"`
								Amclist   struct {
									Amc []string `xml:"amc"`
								} `xml:"amclist"`
							} `xml:"boxpos"`
						} `xml:"MarkInfo"`
					} `xml:"InformF2"`
				} `xml:"Position"`
			} `xml:"Content"`
		} `xml:"WayBill_v4"`
	} `xml:"Document"`
}

// type TTN struct {
// 	XMLName xml.Name `xml:"Documents"`
// 	F2RegId string   `xml:"F2RegId"`
// 	Amc     string   `xml:"amc"`
// }

// type TTN struct {
// 	XMLName xml.Name `xml:"Documents"`

// 	Ce      string `xml:"ce,attr"`
// 	Ns      string `xml:"ns,attr"`
// 	Oref    string `xml:"oref,attr"`
// 	Pref    string `xml:"pref,attr"`
// 	Wb      string `xml:"wb,attr"`
// 	Xsi     string `xml:"xsi,attr"`
// 	Version string `xml:"Version,attr"`
// 	Owner   struct {
// 		FSRARID string `xml:"FSRAR_ID"`
// 	} `xml:"Owner"`
// 	Document struct {
// 		WayBillV4 struct {
// 			Identity string `xml:"Identity"`
// 			Header   struct {
// 				NUMBER       string `xml:"NUMBER"`
// 				Date         string `xml:"Date"`
// 				ShippingDate string `xml:"ShippingDate"`
// 				Type         string `xml:"Type"`
// 				Shipper      struct {
// 					UL struct {
// 						ClientRegId string `xml:"ClientRegId"`
// 						INN         string `xml:"INN"`
// 						KPP         string `xml:"KPP"`
// 						FullName    string `xml:"FullName"`
// 						ShortName   string `xml:"ShortName"`
// 						Address     struct {
// 							Country     string `xml:"Country"`
// 							RegionCode  string `xml:"RegionCode"`
// 							Description string `xml:"description"`
// 						} `xml:"address"`
// 					} `xml:"UL"`
// 				} `xml:"Shipper"`
// 				Consignee struct {
// 					UL struct {
// 						ClientRegId string `xml:"ClientRegId"`
// 						INN         string `xml:"INN"`
// 						KPP         string `xml:"KPP"`
// 						FullName    string `xml:"FullName"`
// 						ShortName   string `xml:"ShortName"`
// 						Address     struct {
// 							Country     string `xml:"Country"`
// 							RegionCode  string `xml:"RegionCode"`
// 							Description string `xml:"description"`
// 						} `xml:"address"`
// 					} `xml:"UL"`
// 				} `xml:"Consignee"`
// 				Transport struct {
// 					TRANTYPE        string `xml:"TRAN_TYPE"`
// 					ChangeOwnership string `xml:"ChangeOwnership"`
// 					TRANSPORTTYPE   string `xml:"TRANSPORT_TYPE"`
// 					TRANCOMPANY     string `xml:"TRAN_COMPANY"`
// 					TRANTRAILER     string `xml:"TRAN_TRAILER"`
// 					TRANCUSTOMER    string `xml:"TRAN_CUSTOMER"`
// 					TRANDRIVER      string `xml:"TRAN_DRIVER"`
// 					TRANLOADPOINT   string `xml:"TRAN_LOADPOINT"`
// 					TRANUNLOADPOINT string `xml:"TRAN_UNLOADPOINT"`
// 					TRANREDIRECT    string `xml:"TRAN_REDIRECT"`
// 					TRANFORWARDER   string `xml:"TRAN_FORWARDER"`
// 				} `xml:"Transport"`
// 				Base string `xml:"Base"`
// 				Note string `xml:"Note"`
// 			} `xml:"Header"`
// 			Content struct {
// 				Position []struct {
// 					Identity string `xml:"Identity"`
// 					Product  struct {
// 						Type         string `xml:"Type"`
// 						FullName     string `xml:"FullName"`
// 						AlcCode      string `xml:"AlcCode"`
// 						Capacity     string `xml:"Capacity"`
// 						UnitType     string `xml:"UnitType"`
// 						AlcVolume    string `xml:"AlcVolume"`
// 						ProductVCode string `xml:"ProductVCode"`
// 						Producer     struct {
// 							UL struct {
// 								ClientRegId string `xml:"ClientRegId"`
// 								FullName    string `xml:"FullName"`
// 								INN         string `xml:"INN"`
// 								KPP         string `xml:"KPP"`
// 								Address     struct {
// 									Country     string `xml:"Country"`
// 									RegionCode  string `xml:"RegionCode"`
// 									Description string `xml:"description"`
// 								} `xml:"address"`
// 							} `xml:"UL"`
// 						} `xml:"Producer"`
// 					} `xml:"Product"`
// 					Quantity string `xml:"Quantity"`
// 					Price    string `xml:"Price"`
// 					PackID   string `xml:"Pack_ID"`
// 					Party    string `xml:"Party"`
// 					FARegId  string `xml:"FARegId"`
// 					InformF2 []struct {
// 						F2RegId  string `xml:"F2RegId"`
// 						MarkInfo struct {
// 							Boxpos struct {
// 								Boxnumber string `xml:"boxnumber"`
// 								Amclist   struct {
// 									Amc string `xml:"amc"`
// 								} `xml:"amclist"`
// 							} `xml:"boxpos"`
// 						} `xml:"MarkInfo"`
// 					} `xml:"InformF2"`
// 				} `xml:"Position"`
// 			} `xml:"Content"`
// 		} `xml:"WayBill_v4"`
// 	} `xml:"Document"`
// }

// type Owner struct {
// 	XMLName xml.Name `xml:"Owner"`
// 	FSRARID string   `xml:"FSRAR_ID"`
// }

// type Product struct {
// 	XMLName  xml.Name `xml:"Product"`
// 	FullName string   `xml:"pref:FullName"`
// 	AlcCode  string   `xml:"pref:AlcCode"`
// 	// ... другие поля продукта
// }

// type InformF2Item struct {
// 	XMLName xml.Name `xml:"InformF2Item"`
// 	F2RegId string   `xml:"pref:F2RegId"`
// }

// type Position struct {
// 	Identity int            `xml:"Identity"`
// 	Product  Product        `xml:"Product"`
// 	Quantity int            `xml:"Quantity"`
// 	InformF2 []InformF2Item `xml:"InformF2"`
// }

// type WayBill struct {
// 	XMLName xml.Name   `xml:"WayBill_v2"`
// 	Content []Position `xml:"Content"`
// }

// type Document struct {
// 	XMLName xml.Name `xml:"Document"`
// 	WayBill WayBill  `xml:"WayBill_v2"`
// }

// type Documents struct {
// 	XMLName  xml.Name `xml:"Documents"`
// 	Owner    Owner    `xml:"Owner"`
// 	Document Document `xml:"Document"`
// }

// type Owner struct {
// 	XMLName xml.Name `xml:"Owner"`
// 	FSRARID string   `xml:"FSRAR_ID"`
// }

// type Client struct {
// 	XMLName     xml.Name `xml:"UL"`
// 	ClientRegId string   `xml:"oref:ClientRegId"`
// 	INN         string   `xml:"oref:INN"`
// 	KPP         string   `xml:"oref:KPP"`
// 	FullName    string   `xml:"oref:FullName"`
// 	ShortName   string   `xml:"oref:ShortName"`
// 	Address     struct {
// 		Country     string `xml:"oref:Country"`
// 		RegionCode  string `xml:"oref:RegionCode"`
// 		Description string `xml:"oref:description"`
// 	} `xml:"oref:address"`
// }

// type Product struct {
// 	XMLName xml.Name `xml:"Product"`
// 	// pref:Type (optional, not shown here)
// 	FullName string  `xml:"pref:FullName"`
// 	AlcCode  string  `xml:"pref:AlcCode"`
// 	Capacity float64 `xml:"pref:Capacity"`
// 	// pref:UnitType (optional, not shown here)
// 	AlcVolume    float64 `xml:"pref:AlcVolume"`
// 	ProductVCode string  `xml:"pref:ProductVCode"`
// 	Producer     *Client `xml:"pref:Producer"`
// }

// type Transport struct {
// 	XMLName          xml.Name `xml:"Transport"`
// 	TRAN_TYPE        string   `xml:"wb:TRAN_TYPE"`
// 	ChangeOwnership  string   `xml:"wb:ChangeOwnership"`
// 	TRANSPORT_TYPE   string   `xml:"wb:TRANSPORT_TYPE"`
// 	TRAN_COMPANY     string   `xml:"wb:TRAN_COMPANY"`
// 	TRAN_TRAILER     string   `xml:"wb:TRAN_TRAILER"`
// 	TRAN_CUSTOMER    string   `xml:"wb:TRAN_CUSTOMER"`
// 	TRAN_DRIVER      string   `xml:"wb:TRAN_DRIVER"`
// 	TRAN_LOADPOINT   string   `xml:"wb:TRAN_LOADPOINT"`
// 	TRAN_UNLOADPOINT string   `xml:"wb:TRAN_UNLOADPOINT"`
// 	TRAN_REDIRECT    string   `xml:"wb:TRAN_REDIRECT"`
// 	TRAN_FORWARDER   string   `xml:"wb:TRAN_FORWARDER"`
// }

// type Position struct {
// 	XMLName  xml.Name  `xml:"Position"`
// 	Identity int       `xml:"wb:Identity"`
// 	Product  Product   `xml:"wb:Product"`
// 	Quantity int       `xml:"wb:Quantity"`
// 	Price    float64   `xml:"wb:Price"`
// 	Pack_ID  string    `xml:"wb:Pack_ID"`
// 	Party    string    `xml:"wb:Party"`
// 	FARegId  string    `xml:"wb:FARegId"`
// 	InformF2 *InformF2 `xml:"wb:InformF2"`
// }

// type InformF2 struct {
// 	XMLName  xml.Name  `xml:"InformF2"`
// 	F2RegId  string    `xml:"ce:F2RegId"`
// 	MarkInfo *MarkInfo `xml:"ce:MarkInfo"`
// }

// type MarkInfo struct {
// 	XMLName xml.Name `xml:"MarkInfo"`
// 	Boxpos  struct {
// 		Boxnumber string `xml:"ce:boxnumber"`
// 		Amclist   struct {
// 			Amc string `xml:"ce:amc"`
// 		} `xml:"ce:amclist"`
// 	} `xml:"ce:boxpos"`
// }

// type WayBill struct {
// 	XMLName  xml.Name `xml:"WayBill_v4"`
// 	Identity string   `xml:"wb:Identity"`
// 	Header   struct {
// 		NUMBER       string    `xml:"wb:NUMBER"`
// 		Date         string    `xml:"wb:Date"`
// 		ShippingDate string    `xml:"wb:ShippingDate"`
// 		Type         string    `xml:"wb:Type"`
// 		Shipper      Client    `xml:"wb:Shipper"`
// 		Consignee    Client    `xml:"wb:Consignee"`
// 		Transport    Transport `xml:"wb:Transport"`
// 		Base         string    `xml:"wb:Base"`
// 	}
// }
