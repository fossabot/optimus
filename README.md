Optimus 
---
[![Docker Repository on Quay](https://quay.io/repository/cloudflavor/optimus/status?token=04373e46-7592-45e9-bfcb-002c57a50d7a "Docker Repository on Quay")](https://quay.io/repository/cloudflavor/optimus)
[![Build Status](https://travis-ci.org/cloudflavor/optimus.svg?branch=master)](https://travis-ci.org/cloudflavor/optimus)
[![codecov](https://codecov.io/gh/cloudflavor/optimus/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudflavor/optimus)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflavor/optimus)](https://goreportcard.com/report/github.com/cloudflavor/optimus) [![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcloudflavor%2Foptimus.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcloudflavor%2Foptimus?ref=badge_shield)
 

Optimus delivers kubernetes native CI/CD Pipelines by leveraging CRDs.

Intended use:

**NOTE: this changes constantly and may contain invalid syntax!
It only includes the CI part and not the CD part**

```yaml
apiVersion: optimus.cloudflavor.io/v1
kind: Pipeline
metadata:
  name: example-pipeline
  namespace: pipeline-namespace
spec:
  # ArchiveArtifacts will archive the artifacts at the end of each stage.
  # If the object storage option is provided it will create a new bucket
  # and push the stage artifacts.
  ArchiveArtifacts: true
  Notifications:
    URI: "https://my-slack-webhook"
    ...
  ContainerBuilder:
    URI: https://uri-to-docker-machine
    ...
    Registry:
      URI: "https://uri-to-my-container-image-registry"
      ...
  # This is the stage section, where we define each stage of the pipeline.
  ArtifactsStorage:
    URI: "https://My-Minio-Example-Server"
    BucketName: "my-bucket"
  Pipeline:
    # NOTE: This will probably be dropped in favor of a VCS web-hook.  
    # Or it will complement the web-hook scenario, for cases in enterprise
    # environments where we can't have web-hooks for various reasons.
    PollInterval: 1m
    Repo: "github.com/test/my-app"
    # Yet to be defined authentication options for docker daemon, VCS server
    # Object storage. Most definitely will use mounted secrets.
    ...
      ...
    Stage:
      # The name of the stage, this will be used to report failure or other
      # information back to the user.
      # NOTE: should this be a label?
      Name: "run tests"
      # This stage is asynchronous.
      Parallel: true
      Steps:
        # Run the commands either in sync or async.
        Parallel: false
        # Either exit when a command fails or continue.
        # NOTE: Is this a good idea? it affects stages up the stack, where
        # a stage will report as successful when, in fact, there were
        # errors?
        # Should this be at the stage level and not at the command level
        ExitOnError: true
        Cmd:
          - npm
          - run
          - test
```

Additional planned features:  
* pod template with resource quota for stages
* expose stage metrics through prometheus

#### Building

If you are on MacOS, run `hack/run-in-docker.sh` and then `make build`, `make gen`, `make test`.


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcloudflavor%2Foptimus.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcloudflavor%2Foptimus?ref=badge_large)