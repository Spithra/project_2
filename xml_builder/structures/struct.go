package structures

import (
	"encoding/xml"
)

// type Documents struct {
// 	XMLName xml.Name `xml:"ns:Documents"`
// 	Version string   `xml:"Version,attr"`
// 	Xsi     string   `xml:"xmlns:xsi,attr"`
// Ns      string   `xml:"xmlns:ns,attr"`
// Pref    string   `xml:"xmlns:pref,attr"`
// Awr     string   `xml:"xmlns:awr,attr"`
// Ce      string   `xml:"xmlns:ce,attr"`
// 	Owner   struct {
// 		FSRARID string `xml:"ns:FSRAR_ID"`
// 	} `xml:"ns:Owner"`
// 	Document struct {
// 		ActWriteOffV4 struct {
// 			Identity string `xml:"awr:Identity"`
// 			Header   struct {
// 				ActNumber    string `xml:"awr:ActNumber"`
// 				ActDate      string `xml:"awr:ActDate"`
// 				TypeWriteOff string `xml:"awr:TypeWriteOff"`
// 				Note         string `xml:"awr:Note"`
// 			} `xml:"awr:Header"`
// 			Content struct {
// 				Position []struct {
// 					Identity       string `xml:"awr:Identity"`
// 					Writeoffvolume struct {
// 						Volume string `xml:"awr:volume"`
// 					} `xml:"awr:writeoffvolume"`
// 					InformF1F2 struct {
// 						InformF2 struct {
// 							F2RegId string `xml:"pref:F2RegId"`
// 						} `xml:"awr:InformF2"`
// 					} `xml:"awr:InformF1F2"`
// 					MarkCodeInfo struct {
// 						Amc string `xml:"amc"`
// 					} `xml:"awr:MarkCodeInfo"`
// 				} `xml:"awr:Position"`
// 			} `xml:"awr:Content"`
// 		} `xml:"ns:ActWriteOffv4"`
// 	} `xml:"ns:Document"`
// }

type Request struct {
	Reason string   `json:"reason"`
	Marks  []string `json:"marks"`
}

type Documents struct {
	XMLName  xml.Name `xml:"ns:Documents"`
	Version  string   `xml:"Version,attr"`
	Xsi      string   `xml:"xmlns:xsi,attr"`
	Ns       string   `xml:"xmlns:ns,attr"`
	Pref     string   `xml:"xmlns:pref,attr"`
	Awr      string   `xml:"xmlns:awr,attr"`
	Ce       string   `xml:"xmlns:ce,attr"`
	Owner    Owner    `xml:"ns:Owner"`
	Document Document `xml:"ns:Document"`
}

type Owner struct {
	XMLName  xml.Name `xml:"ns:Owner"`
	FSRAR_ID string   `xml:"ns:FSRAR_ID"`
}

type Document struct {
	XMLName       xml.Name      `xml:"ns:Document"`
	ActWriteOffv4 ActWriteOffv4 `xml:"ns:ActWriteOff_v4"`
}

type ActWriteOffv4 struct {
	XMLName  xml.Name `xml:"ns:ActWriteOff_v4"`
	Identity int      `xml:"awr:Identity"`
	Header   Header   `xml:"awr:Header"`
	Content  Content  `xml:"awr:Content"`
}

type Header struct {
	XMLName      xml.Name `xml:"awr:Header"`
	ActNumber    int      `xml:"awr:ActNumber"`
	ActDate      string   `xml:"awr:ActDate"`
	TypeWriteOff string   `xml:"awr:TypeWriteOff"`
	Note         string   `xml:"awr:Note"`
}

type Content struct {
	XMLName  xml.Name   `xml:"awr:Content"`
	Position []Position `xml:"awr:Position"`
}

type Position struct {
	XMLName        xml.Name        `xml:"awr:Position"`
	Identity       int             `xml:"awr:Identity"`
	Writeoffvolume *Writeoffvolume `xml:"awr:writeoffvolume"`
	InformF1F2     *InformF1F2     `xml:"awr:InformF1F2"`
	MarkCodeInfo   *MarkCodeInfo   `xml:"awr:MarkCodeInfo"`
}

type Writeoffvolume struct {
	XMLName xml.Name `xml:"awr:writeoffvolume"`
	Volume  string   `xml:"awr:volume"`
}

type InformF1F2 struct {
	XMLName  xml.Name  `xml:"awr:InformF1F2"`
	InformF2 *InformF2 `xml:"awr:InformF2"`
}

type InformF2 struct {
	XMLName xml.Name `xml:"awr:InformF2"`
	F2RegId string   `xml:"pref:F2RegId"`
}

type MarkCodeInfo struct {
	XMLName xml.Name `xml:"awr:MarkCodeInfo"`
	Amc     string   `xml:"ce:amc"`
}

// const (
// 	WegaisNamespace      = "http://fsrar.ru/WEGAIS/WB_DOC_SINGLE_01"
// 	ProductRefNamespace  = "http://fsrar.ru/WEGAIS/ProductRef_v2"
// 	ActWriteOffNamespace = "http://fsrar.ru/WEGAIS/ActWriteOff_v4"
// 	CommonNamespace      = "http://fsrar.ru/WEGAIS/CommonV3"
// )

// type Documents struct {
// 	XMLName  xml.Name       `xml:"Documents"`
// 	Version  string         `xml:"Version,attr"`
// 	Owner    *Owner         `xml:"ns:Owner"`
// 	Document *ActWriteOffV4 `xml:"ns:Document"`
// }

// const (
// 	WegaisNamespace      = "http://fsrar.ru/WEGAIS/WB_DOC_SINGLE_01"
// 	ProductRefNamespace  = "http://fsrar.ru/WEGAIS/ProductRef_v2"
// 	ActWriteOffNamespace = "http://fsrar.ru/WEGAIS/ActWriteOff_v4"
// 	CommonNamespace      = "http://fsrar.ru/WEGAIS/CommonV3"
// )

// Documents represents the root element of the XML request
// type Documents struct {
// 	XMLName  xml.Name  `xml:"Documents"`
// 	Version  string    `xml:"Version,attr"`
// 	Owner    *Owner    `xml:"ns:Owner"`
// 	Document *Document `xml:"ns:Document"`
// }

// const (
// 	WegaisNamespace      = "http://fsrar.ru/WEGAIS/WB_DOC_SINGLE_01"
// 	ProductRefNamespace  = "http://fsrar.ru/WEGAIS/ProductRef_v2"
// 	ActWriteOffNamespace = "http://fsrar.ru/WEGAIS/ActWriteOff_v4"
// 	CommonNamespace      = "http://fsrar.ru/WEGAIS/CommonV3"
// )

//Documents represents the root element of the XML request
// type Documents struct {
// 	XMLName  xml.Name `xml:"ns:Documents"`
// 	Version  string   `xml:"Version,attr"`
// 	Owner    Owner    `xml:"ns:Owner"`
// 	Document Document `xml:"ns:Document"`
// }

// // Owner represents the Owner element
// type Owner struct {
// 	XMLName xml.Name `xml:"ns:Owner"`
// 	FSRARID string   `xml:"ns:FSRAR_ID"`
// }

// // Document represents the Document element
// type Document struct {
// 	XMLName       xml.Name       `xml:"ns:Document"`
// 	ActWriteOffV4 *ActWriteOffV4 `xml:"ns:ActWriteOff_v4"`
// }

// //ActWriteOffV4 represents the ActWriteOff_v4 element
// type ActWriteOffV4 struct {
// 	XMLName  xml.Name           `xml:"ns:ActWriteOff_v4"`
// 	Identity string             `xml:"awr:Identity"`
// 	Header   ActWriteOffHeader  `xml:"awr:Header"`
// 	Content  ActWriteOffContent `xml:"awr:Content"`
// }

// // ActWriteOffHeader represents the Header element within ActWriteOff_v4
// type ActWriteOffHeader struct {
// 	XMLName      xml.Name `xml:"awr:Header"`
// 	ActNumber    string   `xml:"awr:ActNumber"`
// 	ActDate      string   `xml:"awr:ActDate"`
// 	TypeWriteOff string   `xml:"awr:TypeWriteOff"`
// 	Note         string   `xml:"awr:Note"`
// }

// // ActWriteOffContent represents the Content element within ActWriteOff_v4
// type ActWriteOffContent struct {
// 	XMLName  xml.Name              `xml:"awr:Content"`
// 	Position []ActWriteOffPosition `xml:"awr:Position"`
// }

// // ActWriteOffPosition represents the Position element within Content
// type ActWriteOffPosition struct {
// 	XMLName        xml.Name       `xml:"awr:Position"`
// 	Identity       string         `xml:"awr:Identity"`
// 	WriteOffVolume WriteOffVolume `xml:"awr:writeoffvolume"`
// 	InformF1F2     InformF1F2     `xml:"awr:InformF1F2"`
// 	MarkCodeInfo   MarkCodeInfo   `xml:"awr:MarkCodeInfo"`
// }

// // WriteOffVolume represents the writeoffvolume element within Position
// type WriteOffVolume struct {
// 	XMLName xml.Name `xml:"awr:writeoffvolume"`
// 	Volume  string   `xml:"awr:volume"`
// }

// // InformF1F2 represents the InformF1F2 element within Position
// type InformF1F2 struct {
// 	XMLName  xml.Name `xml:"awr:InformF1F2"`
// 	InformF2 InformF2 `xml:"awr:InformF2"`
// }

// // InformF2 represents the InformF2 element within InformF1F2
// type InformF2 struct {
// 	XMLName xml.Name `xml:"awr:InformF2"`
// 	F2RegId string   `xml:"pref:F2RegId"` // Use the namespace alias for clarity
// }

// // MarkCodeInfo represents the MarkCodeInfo element within Position
// type MarkCodeInfo struct {
// 	XMLName xml.Name `xml:"awr:MarkCodeInfo"`
// 	Amc     string   `xml:"ce:amc"` // Use the namespace alias for clarity
// }
