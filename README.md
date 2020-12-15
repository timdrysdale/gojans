# gojans
(possibly) a Go adapter for [Janssen (Gluu) auth server](https://github.com/JanssenProject/home)

![alt text][status]

## Rationale
Things changed, so here we are. Ancient history: After reviewing the options in late November 2020, I started contributing to [gocloak](https://github.com/Nerzal/gocloak) in anticipation of adopting [keycloak](https://www.keycloak.org/) for my [new remote laboratory infrastructure](https://practable.io). Just a few weeks later in early Dec 2020, there was shift in Red Hat's FOSS policy, i.e. [dropping CENTOS](https://arstechnica.com/gadgets/2020/12/centos-shifts-from-red-hat-unbranded-to-red-hat-beta/). It seems likely that [Keycloak is at significant risk of the same fate](https://www.gluu.org/blog/keycloak-is-the-next-centos/). If you were still considering [keycloak](https://www.keycloak.org/) then there is a [good golang adapter available](https://github.com/Nerzal/gocloak) which offers features not found in [goth](https://github.com/markbates/goth) because they are intended for different tasks (building own systems vs logging into others'). 

Now I would like something like ```gocloak``` but for the [Janssen Project](https://github.com/JanssenProject/home), which in short form is known as ```jans```. There is an OpenAPI spec for ```janssen``` which means ```codegen``` (aka swagger in a previous life) might be all that is needed. This repo is my exploration of ```jans``` in relation to golang, and will only turn into a repo if ```codegen``` doesn't do what we need here.

[status]: https://img.shields.io/badge/status-concept-yellow "concept status"
