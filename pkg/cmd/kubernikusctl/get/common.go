package get

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog"

	"github.com/sapcc/kubernikus/pkg/cmd"
	"github.com/sapcc/kubernikus/pkg/cmd/kubernikusctl/common"
)

type GetOptions struct {
	_url       string
	url        *url.URL
	Openstack  *common.OpenstackClient
	Kubernikus *common.KubernikusClient
}

func (o *GetOptions) BindFlags(flags *pflag.FlagSet) {
	o.Openstack.BindFlags(flags)
	common.BindLogFlags(flags)
	flags.StringVar(&o._url, "url", o._url, "URL for Kubernikus API")
}

func (o *GetOptions) PersistentPreRun(c *cobra.Command, args []string) {
	common.SetupLogger()
	cmd.CheckError(o.Openstack.Validate(c, args))
	cmd.CheckError(o.Openstack.Setup())
	cmd.CheckError(o.Openstack.Authenticate())
}

func (o *GetOptions) SetupKubernikusClient() error {
	var err error
	klog.V(2).Infof("SetupKubernikusClient called with url: %v", o._url)
	if o._url == "" {
		fmt.Println("Auto-Detecting Kubernikus Host ...")
		if o.url, err = o.Openstack.DefaultKubernikusURL(); err != nil {
			klog.V(2).Infof("Error detecting kubernikus host: %+v", err)
			return errors.Errorf("You need to provide --url. Auto-Detection failed.")
		}
	} else {
		o.url, err = url.Parse(o._url)
		if err != nil {
			klog.V(2).Infof("Error parsing url: %v", o._url)
			return errors.Wrap(err, "Error parsing url")
		}
	}
	klog.V(2).Infof("Setting up kubernikus client at %v.", o.url)
	o.Kubernikus = common.NewKubernikusClient(o.url, o.Openstack.Provider.TokenID)
	return nil
}
