# "Not testify" Web Server

A platform agnostic order testing webserver to enable users to manage and perform order testing against their underlying ecommerce platform/architecture of choice.

## Developing locally:

### Go Environment

The local environment should be set up like so to allow for proper package and module importing: <br/>
https://learn.gopherguides.com/courses/preparing-your-environment-for-go-development/modules/setting-up-mac-linux/

tldr: the cloned directory of the repo should be the following if all is set up correctly.

- linux: `$HOME/go/src/github.com/lewislewin/testify-webserver`
- windows: `C:\Users\<user>\go\src\github.com\lewislewin\testify-webserver`

Within the projects directory (once installed correctly), you should set `go env -w GOPRIVATE=github.com/lewislewin/testify-webserver` to let go know this is a private repository. This will allow go to correctly pull in modules.

**These steps, if followed correctly, will allow for easy and consistent development between a number of people.**

### Structure

Structure follows golang's ["Organising a go module"](https://go.dev/doc/modules/layout)
