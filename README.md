![alt text][logo]

A Go (golang) client adapter for [Janssen (Gluu) auth server](https://github.com/JanssenProject/home)

![alt text][status]

## Background

Janssen is a Linux Foundation project providing a cloud-native identity and access management (IAM) platform, originating from the [Gluu](https://www.gluu.org/). ```Janssen``` is the sort of open-source project that should attract your support - a big proper thing made independent from its commercial creator so it can take on a life of its own, with its own ecosystem.  

Alternatives include [Keycloak](https://www.keycloak.org/), and its [golang adapter](https:github.com/Nerzal/gocloak). If, like me, you started there, you might want to read [this](https://www.gluu.org/blog/keycloak-is-the-next-centos/) bearing in mind it is from the CEO and founder of Gluu. Nonetheless, the point is well made, because [RedHat did indeed drop Centos](https://arstechnica.com/gadgets/2020/12/centos-shifts-from-red-hat-unbranded-to-red-hat-beta/). Anyway, so far as I can tell, the Gluu server has better (accurate!) documention (including openAPI3.0 spec), a bunch of certifications, and it includes a more complete implementation of UMA2.0 amongst other features.

Thus Janssen project is perhaps a better bet if you are starting an open-source project which requires its own IAM server, and worried about what might happen to the viability of keycloak should it follow the same path as CENTOS (not to mention all the things it already does better). Meanwhile, one of the downsides is the cost of having a play around - I'm running the full community edition Gluu server 4.2 is on a t3.xlarge with 50GB of disk space just to rule out some initial instabilities that might have been caused by low memory when running on a smaller instance. YMMV.

### Using golang with the Janssen project

So far as I can tell, there is no existing SDK for golang.  This repo is an attempt to create a client adapter for the Janssen project, to replicate what [gocloak](https:github.com/Nerzal/gocloak) does for keycloak.  ```Gocloak``` was hand-crafted, whereas there is an openAPI spec for the Janssen project. Why make a separate repo then?

It turns out that the conversion from the existing openAPI3.0 spec to good golang code is not automatic - some adjustments to the spec are needed to get the types right, there is more than one choice of converter, and code-gen, and it would be helpful to include tests down the line. So having the code in a common repo will save others repeating the work of figuring that out. 

Hence ```gojans```, the ```gocloak``` for ```jans``` (the Janssen project). ```gojans``` is NOT an official part of the ```janssen`` project, and is not endorsed by them etc. It's just something I need ... and offered to you in the hope it helps you out too.

### Future maintenance

Simple - don't mess with autogenerated code in ```client``` and ```models``` - unless it somehow needs it.

### Limitations (for now)

(as of Dec 2020) I'm on a deadline for the next phase of own project which needs IAM. I was just starting to get going on using ```keycloak``` by contributing to ```gocloak``` when the news about CENTOS dropped (both in the news, and literal senses). I've decided to switch horses, as I'm not yet in the middle of the stream as it were. This is to avoid predictable refactoring efforts coming early in the project lifecycle. Therefore, implementation modifications beyond the generated code, and testing, will initially focus only what I need for my project. PR super welcome though.

## Organisation

The project is split up so that the ```openapi``` files can be kept separately for ease of inspection (and generating new boiler-plate for patch generation purposes), while the generated (and potentially modified) code is found in ```src```. Any demos can go in their own folder, in ```internal``` depending on how the demo is done (e.g. with ```curl```).

```
.
├── api
│   └── openapi-spec
├── img
├── internal
│   └── curl
└── pkg
    └── jansAuth
        ├── client
        └── models
```


## Mocking / Testing

Test with your own instantiation of the auth-server, or consider mocking (perhaps automatically with [apisprout](https://github.com/danielgtaylor/apisprout) although it may require the examples to be updated for more complete testing).

[status]: https://img.shields.io/badge/status-concept-yellow "concept status"
[logo]: ./img/logo.png "gopher grabs a ride with the janssen racing pigeon"
