package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Solarwatt MyReserve (PV Meter)",
		Sample: `power:
  source: http
  uri: http://192.0.2.2/rest/kiwigrid/wizard/devices # EnergyManager
  jq: .result.items[] | select(.deviceModel[].deviceClass == "com.kiwigrid.devices.location.Location" ) | .tagValues.PowerProduced.value`,
	}

	registry.Add(template)
}