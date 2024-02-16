package parseOldCongressionalRecord

import (
	"encoding/xml"
	"strings"
)

type Font struct {
	Family string `xml:"family,attr"`
	Color  string `xml:"color,attr"`
	Font   string
	Group  string
	ID     int `xml:"id,attr"`
	Size   int `xml:"size,attr"`
}

type Text struct {
	Text   string `xml:",chardata"`
	Top    int    `xml:"top,attr"`
	Left   int    `xml:"left,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	FontID int    `xml:"font,attr"`
}

type Page struct {
	Position string `xml:"position,attr"`
	Fonts    []Font `xml:"fontspec"`
	Text     []Text `xml:"text"`
	Number   int    `xml:"number,attr"`
	Top      int    `xml:"top,attr"`
	Left     int    `xml:"left,attr"`
	Height   int    `xml:"height,attr"`
	Width    int    `xml:"width,attr"`
}

type XML struct {
	XMLName  xml.Name `xml:"pdf2xml"`
	Producer string   `xml:"producter,attr"`
	Version  string   `xml:"version,attr"`
	Pages    []Page   `xml:"page"`
}

func ParseXML(xmlString string) (string, error) {
	// remove italics and bolds
	xmlString = strings.ReplaceAll(xmlString, "<b>", "")
	xmlString = strings.ReplaceAll(xmlString, "</b>", "")
	xmlString = strings.ReplaceAll(xmlString, "<i>", "")
	xmlString = strings.ReplaceAll(xmlString, "</i>", "")

	xmlBytes := []byte(xmlString)

	xmlData := XML{}

	err := xml.Unmarshal(xmlBytes, &xmlData)
	if err != nil {
		return "", err
	}

	lineBottom := 0

	var strBuilder strings.Builder

	for pageIndex := range xmlData.Pages {
		for textIndex := range xmlData.Pages[pageIndex].Text {
			text := &xmlData.Pages[pageIndex].Text[textIndex]

			if text.Top >= lineBottom || text.Top < lineBottom-300 {
				strBuilder.WriteString("\n")
			}

			strBuilder.WriteString(text.Text)

			lineBottom = text.Top + text.Height
		}
	}

	return strBuilder.String(), nil
}
