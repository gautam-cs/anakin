# API Server for [API_SERVICE_NAME]

REST API Server in golang.

## Get Started

** Preferred IDE is VS code. **

__Fetch__

```bash
git clone https://[username]@bitbucket.org/gautam/[repo].git 
git submodule init
git submodule update
cd app
go get -u
go build
```

## Configure

Replace `[API_SERVICE_NAME]` with a service name. Service name should be lowercase with underscrore `[a-z_]` max 50 chars

* example: `go_starter_template`


## Directory [ Refer migration_onboarding repo ]

* Place utility functions in `app/utils`
* Place constants in `app/constants`. They should not depend on any other `app` package.
* Place any db models in `app/models`.

* Place feature routes should be placed under `app/resources/`.
* Place app logic should be placed under `app/services/`.
* Place app db api calls under `app/db/`.
* Place any sql tables/modifications in `docs/schema.sql`

* `templates` directory should have templates for email notifications.

## Deployment

* `master` branch is deployed to prod environment.
* `preprod` branch is deployed to preprod environment.

This repo uses bitbucket pipelines to deploy to ECS. the flow is
git push -> bitbucket -> aws code deploy -> ecs

scripts are

* `Dockerfile` for customizing/building docker
* `buildspec.yml` this is for aws code deploy

__DO NOT commit Production or AWS access keys into repo__

## Config

* code config is `app/config/secrets/env.json` depending on env mentioned here, the config file is loaded
* config wref to environment are location are `app/config/secrets/[environment].json`

__NOTE__

* Please make no change in `app/config/secrets/preprod.json`
* Please make no change in `app/config/secrets/prod.json`

## Contributors

* [gautam.official.it@gmail.com](mailto:gautam.official.it@gmail.com)