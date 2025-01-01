resource "routeros_interface_l2tp_client" "test" {
  password     = "StrongPass"
  name         = "l2tp-client-1"
  disabled     = false
  user         = "MT-User"
}