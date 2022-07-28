package usps

import (
	"bytes"
	"fmt"

	"github.com/antchfx/xmlquery"
)

type response struct {
	number  string
	details []string
}

var (
	number      string
	description string
	source      string
	helpContext string
)

func (r *response) isError(xmlBody []byte) error {
	doc, err := xmlquery.Parse(bytes.NewReader(xmlBody))
	if err != nil {
		return err
	}

	root := xmlquery.FindOne(doc, "//Error")

	if root != nil {
		if n := xmlquery.FindOne(root, "//Number"); n != nil {
			number = n.InnerText()
		}

		if d := xmlquery.FindOne(root, "//Description"); d != nil {
			description = d.InnerText()
		}

		if hc := xmlquery.FindOne(root, "//HelpContext"); hc != nil {
			helpContext = hc.InnerText()
		}

		if s := xmlquery.FindOne(root, "//Source"); s != nil {
			source = s.InnerText()
		}

		return fmt.Errorf("USPS body error.\n Number: %s\n Description: %s\n Source: %s\n HelpContext: %s\n",
			number, description, source, helpContext)
	}

	return nil
}

func (r *response) unmarshalTrackData(xmlBody []byte) error {
	doc, err := xmlquery.Parse(bytes.NewReader(xmlBody))
	if err != nil {
		return err
	}

	root := xmlquery.FindOne(doc, "//TrackInfo")

	if root == nil {
		return fmt.Errorf("the XML structure does not match the parsing conditions or is missing")
	}

	if trackID := xmlquery.FindOne(root, "@ID"); trackID != nil {
		r.number = trackID.InnerText()
	}
	// NOTE! If you want the first element in the "details" slice to be the Summary,
	//then do not change the order of the condition below.
	var details []string
	if summary := root.SelectElement("//TrackSummary"); summary != nil {
		details = append(details, summary.InnerText())
	}

	if nodes := xmlquery.Find(root, "//TrackDetail"); nodes != nil {
		for _, v := range nodes {
			details = append(details, v.InnerText())
		}
		r.details = details
	}

	return nil
}
