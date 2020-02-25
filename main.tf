provider "ping" {
  username             = "Administrator"
  password             = "Testpassword1"
  base_url             = "https://192.168.33.111:9000/pa-admin-api/v3/"
  xsf_header           = "pingAccess"
  insecure_skip_verify = true
}

resource "ping_virtualhost" "test" {
  host = "test"
  port = 3000
}
