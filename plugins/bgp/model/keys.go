/*REMEMBER TO USE THIS MODEL WHEN BRINGING OVER THE PLUG-IN*/

package model

import (
	"github.com/ligato/vpp-agent/pkg/models"
)


const ModuleName = "bgp"
var (
	ModelBgpConf = models.Register(&BgpConf{}, models.Spec{
		Module:  ModuleName,
		Version: "v1",
		Type:    "bgpconf",
	},
	models.WithNameTemplate("{{.Name}}"),
	)
)
