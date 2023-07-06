package parseOldCongressionalRecord

import (
	"encoding/xml"
	"os"
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

func ParseXML(xmlFile string) (string, error) {
	xmlBytes, err := os.ReadFile(xmlFile)
	if err != nil {
		return "", err
	}

	sbody := string(xmlBytes)

	// remove italics and bolds
	sbody = strings.ReplaceAll(sbody, "<b>", "")
	sbody = strings.ReplaceAll(sbody, "</b>", "")
	sbody = strings.ReplaceAll(sbody, "<i>", "")
	sbody = strings.ReplaceAll(sbody, "</i>", "")

	xmlBytes = []byte(sbody)

	xmlData := XML{}

	err = xml.Unmarshal(xmlBytes, &xmlData)
	if err != nil {
		return "", err
	}

	lineBottom := 0

	var txt string

	for pageIndex := range xmlData.Pages {
		for textIndex := range xmlData.Pages[pageIndex].Text {
			text := &xmlData.Pages[pageIndex].Text[textIndex]

			if text.Top >= lineBottom || text.Top < lineBottom-300 {
				txt += "\n"
			}

			txt += text.Text

			lineBottom = text.Top + text.Height
		}
	}

	return txt, nil
}
