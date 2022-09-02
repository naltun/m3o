# API Ingress

Ingress for apps and functions

## Overview

The M3O API Ingress is a single entrypoint for apps, functions and more

- [api](#api)
- [apps](#apps)
- [functions](#functions)
- [url proxy](#url-proxy)
- [user verification](#user-verification)

## API

We provide GET/POST access to the /v1/* API via `api_key=xxx` param rather than Authorization: Bearer xxx

Example URL

```
https://m3o.one/api/helloworld/call?api_key=xxxx
```

Routes to the below URL with `Authorization: Bearer XXX` set.

```
https://api.m3o.com/v1/helloworld/call
```

The `/api` endpoint supports both HTTP params and JSON body where `Content-Type: application/json` is specified.

## Apps

Apps are given a unique id and subdomain using m3o.app e.g [helloworld.m3o.app](https://helloworld.m3o.app) resolves to the app id helloworld.

## Functions

Functions are given a unique id and subdomain much like apps using m3o.sh e.g [helloworld.m3o.sh](https://helloworld.m3o.sh/) resolves to the 
function helloworld. 

## URL Proxy

The [URL](https://github.com/micro/services/tree/master/url) service provides link shortening and sharing. The URL Proxy fronts those urls 
as a single entrypoint at https://m3o.one. 

- Assumes url is of format `https://m3o.one/u/AArfeZE`
- Will call `https://api.m3o.com/url/resolve?shortURL=https://m3o.one/u/AArfeZE`
- URL service should return `destinationURL=https://foobar.com/example`
- Proxy will issue a 301 redirect

## User Verification

User email verification is performed via the User Verify Email endpoint. The gateway serves user.m3o.com for this purpose.

## Usage

Specify `M3O_API_TOKEN` to the binary and run the ingress

### Micro Services

Set the config values 

```
micro config set micro.app.domain m3o.app
micro config set micro.app.domain m3o.sh
micro config set micro.url.host_prefix https://m3o.one/u/
micro config set micro.user.verify_email_url https://user.m3o.com
```

Deploy the url, function, app, user services

```
micro run github.com/micro/services/app
micro run github.com/micro/services/function
micro run github.com/micro/services/url
micro run github.com/micro/services/user
```