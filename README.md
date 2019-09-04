Terraform Provider
==================

There is no upstream yet. 

Currently Contrail terraform provider works and partly tested only with R4.1 version of contrail API.
There is only couple of ACC tests and 95% sources are generated via gen.sh.
Fill free to generate sources for another version of contrail and help with pushing this repo to UpStream. 

Maintainers
-----------

This provider plugin is maintained by:

* Stanislau Petrusiou ([@PetrusHahol](https://github.com/PetrusHahol))

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Also tested with terraform 0.12.6 , it works fine.

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/PetrusHahol/terraform-provider-contrail`

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-contrail
$ make build
```

Using the provider
----------------------

This provider is not in UpStream yet, because of that if you are going to use the provider, you need to build locally and put built provider into the folder with terraform `provider.tf` file. 

Developing the Provider
---------------------------
We are looking forward to somebody who has an interest to extend our solution, write test, documentation for heading to the upstream afterwords.

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-contraiil
...
```

## Example Usage
    
```
provider "contrail" {
	server = "mycontrail.com"
	auth_url = "https://myauthurl:5000/v2.0"
	username = "admin"
	tenant_name = "admin"
	password = "pwd"
}
```

## Configuration Reference

* `server` - The REST api url for calls. If omitted, the CONTRAIL_API_SERVER environment variable is used.
* `auth_url` - The Identity authentication URL. If omitted, the OS_AUTH_URL environment variable is used.
* `user_name` - If omitted, the `OS_USERNAME` environment variable is used.
* `tenant_name` - (Optional) The Name of the Tenant (Identity v2) or Project (Identity v3) to login with. If omitted, the `OS_TENANT_NAME` or `OS_PROJECT_NAME` environment variable are used.
* `password` - (Optional) The Password to login with. If omitted, the `OS_PASSWORD` environment variable is used.
* `token` - (Optional) Required if not using `user_name` and `password`) A token is an expiring, temporary means of access issued via the Keystone service. By specifying a token, you do not have to specify a username/password combination, since the token was already created by a username/password out of band of Terraform. If omitted, the `OS_TOKEN` or `OS_AUTH_TOKEN` environment variables are used.
* `port` - (Optional, 8082 is default) Port which will be added in the end of `server` variable for the REST calls.
* `username` - (Optional) The Username to login with. If omitted, the `OS_USERNAME` environment variable is used.
 
## Acceptance Tests

Acceptance Tests are a crucial part of adding features or fixing a bug. Please
make sure to read the core [testing](https://www.terraform.io/docs/extend/testing/index.html)
documentation for more information about how Acceptance Tests work.

In order to run the Acceptance Tests, you'll need to set the following
environment variables:

* `CONTRAIL_API_SERVER`
* `OS_USERNAME`
* `OS_TENANT_NAME` 
* `OS_PROJECT_NAME` 
* `OS_PASSWORD`  
* `OS_AUTH_TOKEN` 
* `OS_TOKEN` 
* `OS_AUTH_URL` 
* `TF_ACC` - should be `True` for test running 

The following additional environment variables might be required depending on
the feature or bug you're testing:

* `OS_HEALTH_CHECK_ID` - The UUID of the health check in your test environment.

* `OS_PROJECT_ID` - The UUID of the project you are using in your test environment.

* `OS_VM_ID` - The UUID of the vm in your test environment. 

We recommend only running the acceptance tests related to the feature or bug
you're working on. To do this, run:


We recommend only running the acceptance tests related to the feature or bug
you're working on. To do this, run:

```shell
$ cd $GOPATH/src/github.com/PetrusHahol/terraform-provider-contrail
$ make testacc TEST=./contrail TESTARGS="-run=<keyword> -count=1"
```

Where `<keyword>` is the full name or partial name of a test. For example:

```shell
$ make testacc TEST=./contrail TESTARGS="-run=TestAccNetworkRefsBasic -count=1"
```

We recommend running tests with logging set to `DEBUG`:

```shell
$ TF_LOG=DEBUG make testacc TEST=./contrail TESTARGS="-run=TestAccNetworkRefsBasic -count=1"
```

And you can even enable OpenStack debugging to see the actual HTTP API requests:

```shell
$ TF_LOG=DEBUG OS_DEBUG=1 make testacc TEST=./contrail TESTARGS="-run=TestAccNetworkRefsBasic -count=1"
```

