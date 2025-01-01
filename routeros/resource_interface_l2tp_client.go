package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
 {
  ".id": "*1C",
  "add-default-route": "no", #default bool
  "allow": "chap,mschap1,mschap2,pap", #default mschap2
  connect-to: "ip-address", #required
  comment: "string",
  default-route-distance: 1, #default 1
  dial-on-demand: "no", #default bool
  disabled: "no", #default bool
  keepalive-timeout: 10, #default 10
  max-mru: "1450", #default 1450
  max-mtu: "1450", #default 1450
  mrru: "disabled", #default disabled
  name: "", #required
  password: "", #required
  profile: "default-encryption", #default default-encryption
  user: "", #required
  use-ipsec: "no", #default bool
  allow-fast-path: "", #default bool
  l2tp-proto-version: "l2tpv2,l2tpv3-ip,l2tpv3-udp,l2tpv" #default l2tpv2
  l2tpv3-cookie-length: "0", #default 0
  l2tpv3-digest-hash: "md5,sha1,none" #default md5
  user-peer-dns: "no", #default no
  copy-from: "string",
  src-address: "",
  l2tpv3-circuit-id: "",
  ipsec-secret: ""
 }
*/

// https://help.mikrotik.com/docs/spaces/ROS/pages/2031631/L2TP#L2TP-L2TPClient
func ResourceInterfaceL2TPClient() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/l2tp-client"),
		MetaId:           PropId(Id),

		"add_default_route": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to add L2TP remote address as a default route.",
		},
		"allow": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			Description: "Allowed authentication methods, by default all methods are allowed.",
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap2", "mschap1", "chap", "pap"}, false),
			},
		},
		"connect_to": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "IP address of the L2TP server.",
		},
		KeyComment: PropCommentRw,
		"default_route_distance": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
			Description: "sets distance value applied to auto created default route, if add-default-route is also " +
				"selected.",
			ValidateFunc: validation.IntBetween(0, 255),
		},
		"dial_on_demand": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "connects to AC only when outbound traffic is generated. If selected, then route with " +
				"gateway address from 10.112.112.0/24 network will be added while connection is not " +
				"established.",
		},
		KeyDisabled: PropDisabledRw,
		KeyInvalid:  PropInvalidRo,
		"keepalive_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     10,
			Description: "Sets keepalive timeout in seconds.",
		},
		"max_mru": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "auto",
			Description: "Maximum Receive Unit.",
		},
		"max_mtu": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "auto",
			Description: "Maximum Transmission Unit.",
		},
		"mrru": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "disabled",
			Description: "Maximum packet size (512..65535 or disabled) that can be received on the link. If a packet " +
				"is bigger than tunnel MTU, it will be split into multiple packets, allowing full size IP or Ethernet " +
				"packets to be sent over the tunnel.",
		},
		KeyName: PropName("Name of the L2TP interface."),
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: "Password used to authenticate.",
		},
		"profile": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "default-encryption",
			Description: "Specifies which PPP profile configuration will be used when establishing the tunnel.",
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "Username used for authentication.",
		},
		KeyRunning: PropRunningRo,
		"service_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
			Description: "Specifies the service name set on the access concentrator, can be left blank to connect " +
				"to any PPPoE server.",
		},
		"use_ipsec": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			Description: "When this option is enabled, dynamic IPSec peer configuration and policy (transport mode) " +
				"is added to encapsulate L2TP connection into IPSec tunnel. ",
		},
		"allow_fast_path": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Allow to forward packets without additional processing in the Linux kernel.",
		},
		"l2tp_proto_version": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "l2tpv2",
			Description:  "Specify protocol version to use.",
			ValidateFunc: validation.StringInSlice([]string{"l2tpv2", "l2tpv3-ip", "l2tpv3-udp", "l2tpv3"}, false),
		},
		"l2tpv3_cookie_length": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Configures an L2TPv3 pseudowire static session cookie.",
		},
		"l2tpv3_digest_hash": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "md5",
			Description:  "Specifies which hash function to be used.",
			ValidateFunc: validation.StringInSlice([]string{"md5", "sha1", "none"}, false),
		},
		"use_peer_dns": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "To use peer dns.",
		},
		"copy_from": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "Copy settings from other interface.",
		},
		"src_address": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "Source address of the L2TP client.",
		},
		"l2tpv3_circuit_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "Set the virtual circuit identifier to bind the one end of the L2TPv3 control channel.",
		},
		"ipsec_secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: "Preshared key used when use-ipsec is enabled.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
