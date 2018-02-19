package xml_ops

import (
	"go/go-xml/entity"
	"github.com/beevik/etree"
	"fmt"
)

type ArmXmlOps struct {}

var (
	doc *etree.Document
	payloads, payload, id, typee, actions *etree.Element
)

func(a *ArmXmlOps) AddElementsToXML(arm *entity.ArmPayload) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("xml_ops/xmls/arms.xml"); err != nil {
		doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
		doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)
		payloads = doc.CreateElement("payloads")
	} else {
		payloads = doc.SelectElement("payloads")
	}

	payload = payloads.CreateElement("payload")
	payload.CreateAttr("capacity", arm.Config.Capacity)

	id = payload.CreateElement("id")
	id.SetText(arm.Device_id)

	actions = payload.CreateElement("actions")
	for i:= 0; i<len(arm.Actions) ;i++  {
		actions.CreateElement("action").SetText(arm.Actions[i])
	}
	doc.Indent(4)
	doc.WriteToFile("xml_ops/xmls/arms.xml")

}

func(a *ArmXmlOps) ReadFromXML(arm *entity.ArmPayload) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("xml_ops/xmls/arms.xml"); err != nil {
		panic(err)
	}
	for _, t := range doc.FindElements("//payload[@capacity='medium']/id") {
		fmt.Println("device_id:", t.Text() )
	}
}

