package fdwca

import (
	"errors"
	"regexp"
	"strings"

	"github.com/sfborg/sflib/pkg/coldp"
	"github.com/sfborg/sflib/pkg/dwca"
)

var reg = regexp.MustCompile(`\d{4}-\d{4}-\d{4}-\d{4}`)

func (fd *fdwca) importEML() error {
	var doi string
	eml := fd.dwca.EML()
	if eml == nil {
		return errors.New("no EML data in DwCA")
	}
	if strings.Contains(eml.Dataset.ID, "doi.") {
		doi = eml.Dataset.ID
	}

	meta := coldp.Meta{
		Title:       eml.Dataset.Title,
		DOI:         doi,
		Description: eml.Dataset.Abstract.Para,
		Creators:    creators(eml),
		Contact:     contact(eml),
	}

	err := fd.sfga.InsertMeta(&meta)
	if err != nil {
		return err
	}

	return nil
}

func creators(eml *dwca.EML) []coldp.Actor {
	var res []coldp.Actor
	for _, c := range eml.Dataset.Creators {
		a := coldp.Actor{}
		name := c.IndividualName
		if name != nil {
			a.Given = name.GivenName
			a.Family = name.SurName
		}
		a.Email = c.ElectronicMailAddress

		org := c.OrganizationName
		if org != nil {
			a.Organization = org.Value
		}

		match := reg.MatchString(c.ID)
		if match {
			a.Orcid = c.ID
		}
		res = append(res, a)
	}
	return res
}

func contact(eml *dwca.EML) *coldp.Actor {
	if len(eml.Dataset.Contacts) == 0 {
		return nil
	}
	c := eml.Dataset.Contacts[0]
	a := coldp.Actor{}
	name := c.IndividualName
	if name != nil {
		a.Given = name.GivenName
		a.Family = name.SurName
	}
	a.Email = c.ElectronicMailAddress

	org := c.OrganizationName
	if org != nil {
		a.Organization = org.Value
	}

	return &a
}
