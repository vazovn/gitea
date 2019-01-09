---
date: "2018-06-24:00:00+02:00"
title: "API Usage"
slug: "api-usage"
weight: 40
toc: true
draft: false
menu:
  sidebar:
    parent: "advanced"
    name: "API Usage"
    weight: 40
    identifier: "api-usage"
---

# Gitea API Usage

## Enabling/configuring API access

By default, `ENABLE_SWAGGER_ENDPOINT` is true, and
`MAX_RESPONSE_ITEMS` is set to 50.  See [Config Cheat
Sheet](https://docs.gitea.io/en-us/config-cheat-sheet/) for more
information.

## Authentication via the API

Gitea supports these methods of API authentication:

- HTTP basic authentication
- `token=...` parameter in URL query string
- `access_token=...` parameter in URL query string
- `Authorization: token ...` header in HTTP headers

All of these methods accept the same apiKey token type.  You can
better understand this by looking at the code -- as of this writing,
Gitea parses queries and headers to find the token in
[modules/auth/auth.go](https://github.com/go-gitea/gitea/blob/6efdcaed86565c91a3dc77631372a9cc45a58e89/modules/auth/auth.go#L47).

You can create an apiKey token via your gitea install's web interface:
`Settings | Applications | Generate New Token`.

### More on the `Authorization:` header

For historical reasons, Gitea needs the word `token` included before
the apiKey token in an authorization header, like this:

```
Authorization: token 65eaa9c8ef52460d22a93307fe0aee76289dc675
```

In a `curl` command, for instance, this would look like:

```
curl -X POST "http://localhost:4000/api/v1/repos/test1/test1/issues" \
    -H "accept: application/json" \
    -H "Authorization: token 65eaa9c8ef52460d22a93307fe0aee76289dc675" \
    -H "Content-Type: application/json" -d "{ \"body\": \"testing\", \"title\": \"test 20\"}" -i
```

As mentioned above, the token used is the same one you would use in
the `token=` string in a GET request.

## Listing your issued tokens via the API

As mentioned in
[#3842](https://github.com/go-gitea/gitea/issues/3842#issuecomment-397743346),
`/users/:name/tokens` is special and requires you to authenticate
using BasicAuth, as follows:

### Using basic authentication:

```
$ curl --request GET --url https://yourusername:yourpassword@gitea.your.host/api/v1/users/yourusername/tokens
[{"name":"test","sha1":"..."},{"name":"dev","sha1":"..."}]
```

## Sudo

The API allows admin users to sudo API requests as another user. Simply add either a `sudo=` parameter or `Sudo:` request header with the username of the user to sudo.
