Terraform Provider Ping Identity
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by the Terraform team at [Chris Usick](https://www.github.com/chrisUsick).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Usage
---------------------

```
# For example, restrict template version in 0.1.x
provider "ping" {
    username             = "Administrator"
    password             = "Testpassword1"
    base_url             = "https://192.168.33.111:9000/pa-admin-api/v3/"
    insecure_skip_verify = true
}
resource "ping_virtualhost" "test" {
    host = "test"
    port = 4000
}
```

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/chrisUsick/terraform-provider-ping`

```sh
$ mkdir -p $GOPATH/src/github.com/chrisUsick; cd $GOPATH/src/github.com/chrisUsick
$ git clone git@github.com:chrisusick/terraform-provider-ping
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/chrisUsick/terraform-provider-ping
$ make build
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-ping
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```