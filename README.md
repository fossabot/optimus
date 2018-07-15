# CI Pipelines
[![Docker Repository on Quay](https://quay.io/repository/cloudflavor/pipelines/status "Docker Repository on Quay")](https://quay.io/repository/cloudflavor/pipelines)
[![Build Status](https://travis-ci.org/PI-Victor/pipelines.svg?branch=master)](https://travis-ci.org/PI-Victor/pipelines)
[![codecov](https://codecov.io/gh/PI-Victor/pipelines/branch/master/graph/badge.svg)](https://codecov.io/gh/PI-Victor/pipelines)  

This is a CRD controller for orchestrating k8s native resources into ci-cd
pipelines.  

The intended purpose is to have a VCS watcher/web-hook that will run the
pipeline on a VCS event, run each defined stage, build the container image
(securely exposed docker instance on a remote host), push the container to the
user defined registry, notify the user by calling his CRD defined web-hook of
the stage/stage command exit status.  

If the object storage option is defined, the job artifacts are tar archived and
pushed under the project bucket as `jobname-date.tar.gz`.

 *And maybe some other magic too.*  


A rough estimate of the CRDs intended use.  

**NOTE: this changes constantly and might be invalid syntax! It only includes
the CI part and not the CD part**

```yaml
apiVersion: pipelines.cloudflavor.io/v1
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
  Project:
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
      Commands:
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
