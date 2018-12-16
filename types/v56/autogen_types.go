package types

import "encoding/xml"

type VCDHostResource__rasd struct {
	XMLName                                      xml.Name `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData HostResource,omitempty" json:"HostResource,omitempty"`
	BusSubType                                   string   `xml:"http://www.vmware.com/vcloud/v1.5 busSubType,attr"  json:",omitempty"`
	BusType                                      string   `xml:"http://www.vmware.com/vcloud/v1.5 busType,attr"  json:",omitempty"`
	Capacity                                     string   `xml:"http://www.vmware.com/vcloud/v1.5 capacity,attr"  json:",omitempty"`
	Iops                                         string   `xml:"http://www.vmware.com/vcloud/v1.5 iops,attr"  json:",omitempty"`
	AttrNs10SpacestorageProfileHref              string   `xml:"http://www.vmware.com/vcloud/v1.5 storageProfileHref,attr"  json:",omitempty"`
	AttrNs10SpacestorageProfileOverrideVmDefault string   `xml:"http://www.vmware.com/vcloud/v1.5 storageProfileOverrideVmDefault,attr"  json:",omitempty"`
}

type VCDItem struct {
	XMLName                        xml.Name               `xml:"http://www.vmware.com/vcloud/v1.5 Item,omitempty" json:"Item,omitempty"`
	VCDAddressOnParent__rasd       string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData AddressOnParent,omitempty" json:"AddressOnParent,omitempty"`
	VCDAddress__rasd               string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Address,omitempty" json:"Address,omitempty"`
	VCDAllocationUnits__rasd       string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData AllocationUnits,omitempty" json:"AllocationUnits,omitempty"`
	VCDAutomaticAllocation__rasd   string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData AutomaticAllocation,omitempty" json:"AutomaticAllocation,omitempty"`
	VCDAutomaticDeallocation__rasd string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData AutomaticDeallocation,omitempty" json:"AutomaticDeallocation,omitempty"`
	VCDCaption__rasd               string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Caption,omitempty" json:"Caption,omitempty"`
	VCDChangeableType__rasd        string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ChangeableType,omitempty" json:"ChangeableType,omitempty"`
	VCDConfigurationName__rasd     string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ConfigurationName,omitempty" json:"ConfigurationName,omitempty"`
	VCDConsumerVisibility__rasd    string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ConsumerVisibility,omitempty" json:"ConsumerVisibility,omitempty"`
	VCDDescription__rasd           string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Description,omitempty" json:"Description,omitempty"`
	VCDElementName__rasd           string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ElementName,omitempty" json:"ElementName,omitempty"`
	VCDGeneration__rasd            string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Generation,omitempty" json:"Generation,omitempty"`
	VCDHostResource__rasd          *VCDHostResource__rasd `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData HostResource,omitempty" json:"HostResource,omitempty"`
	VCDInstanceID__rasd            int16                  `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData InstanceID,omitempty" json:"InstanceID,omitempty"`
	VCDLimit__rasd                 string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Limit,omitempty" json:"Limit,omitempty"`
	VCDMappingBehavior__rasd       string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData MappingBehavior,omitempty" json:"MappingBehavior,omitempty"`
	VCDOtherResourceType__rasd     string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData OtherResourceType,omitempty" json:"OtherResourceType,omitempty"`
	VCDParent__rasd                string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Parent,omitempty" json:"Parent,omitempty"`
	VCDPoolID__rasd                string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData PoolID,omitempty" json:"PoolID,omitempty"`
	VCDReservation__rasd           string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Reservation,omitempty" json:"Reservation,omitempty"`
	VCDResourceSubType__rasd       string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ResourceSubType,omitempty" json:"ResourceSubType,omitempty"`
	VCDResourceType__rasd          int8                   `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData ResourceType,omitempty" json:"ResourceType,omitempty"`
	VCDWeight__rasd                string                 `xml:"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ResourceAllocationSettingData Weight,omitempty" json:"Weight,omitempty"`
}

type VCDRasdItemsList struct {
	XMLName xml.Name   `xml:"http://www.vmware.com/vcloud/v1.5 RasdItemsList,omitempty" json:"RasdItemsList,omitempty"`
	VCDItem []*VCDItem `xml:"http://www.vmware.com/vcloud/v1.5 Item,omitempty" json:"Item,omitempty"`
}
