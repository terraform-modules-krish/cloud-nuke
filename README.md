***WARNING: THIS REPO IS AN AUTO-GENERATED COPY.*** *This repo has been copied from [Gruntwork’s](https://gruntwork.io/) GitHub repositories so that you can consume it from your company’s own internal Git repositories. This copy is automatically created and updated by the `repo-copier` CLI tool. If you need to make changes to this repo, you should make the changes in a separate fork, and NOT make changes directly in this repo, as otherwise, the `repo-copier` will overwrite your changes! Please see the `repo-copier` [documentation](https://github.com/terraform-modules-krish/repo-copier) for more information on how the code is copied, how cross-references are updated, how the changelog is handled, etc.*

***

_You may find it valuable to view the following resources in the original repo. If these links give you a 404, visit https://app.gruntwork.io to gain access or email support@gruntwork.io if you need assistance._

[Home Page](https://github.com/gruntwork-io/cloud-nuke/) |
[Pull Requests](https://github.com/gruntwork-io/cloud-nuke/pulls) |
[Issues](https://github.com/gruntwork-io/cloud-nuke/issues) |
[Releases and Assets](https://github.com/gruntwork-io/cloud-nuke/releases)

_Alternatively, you can view a copied version of the resources listed above._

[Pull Requests](https://github.com/terraform-modules-krish/cloud-nuke/blob/master/.github/PULL_REQUESTS.md) |
[Issues](https://github.com/terraform-modules-krish/cloud-nuke/blob/master/.github/ISSUES.md) |
[ChangeLog](https://github.com/terraform-modules-krish/cloud-nuke/blob/master/.github/CHANGELOG.md)

***

# cloud-nuke

This repo contains a CLI tool to delete all cloud (AWS, Azure, GCP) resources in an account. cloud-nuke was created for situations when you might have an account you use for testing and need to clean up left over resources so you're not charged for them. Also great for cleaning out accounts with redundant resources.

The currently supported functionality includes:

## AWS

* Deleting all Auto scaling groups in an AWS account
* Deleting all Elastic Load Balancers (Classic and V2) in an AWS account
* Deleting all EBS Volumes in an AWS account
* Deleting all unprotected EC2 instances in an AWS account
* Deleting all AMIs in an AWS account
* Deleting all Snapshots in an AWS account

## Azure

_Coming Soon_

## GCP

_Coming Soon_

### WARNING: THIS TOOL IS HIGHLY DESTRUCTIVE, ALL SUPPORTED RESOURCES WILL BE DELETED. ITS EFFECTS ARE IRREVERSIBLE AND SHOULD NEVER BE USED IN A PRODUCTION ENVIRONMENT

## Install

1. Download the latest binary for your OS on the [releases page](https://github.com/gruntwork-io/cloud-nuke/releases).
2. Move the binary to a folder on your `PATH`. E.g.: `mv cloud-nuke_darwin_amd64 /usr/local/bin/cloud-nuke`.
3. Add execute permissions to the binary. E.g.: `chmod u+x /usr/local/bin/cloud-nuke`.
4. Test it installed correctly: `cloud-nuke --help`.

## Usage

Simply running `cloud-nuke <provider>` (e.g. `cloud-nuke aws`) will start the process of cleaning up your cloud account. You'll be shown a list of resources that'll be deleted as well as a prompt to confirm before any deletion actually takes place.

### Excluding Regions

You can use the `--exclude-region` flag to exclude resources in certain regions from being deleted. For example the following command does not nuke resources in `ap-south-1` and `ap-south-2` regions:

```shell
cloud-nuke aws --exclude-region ap-south-1 --exclude-region ap-south-2
```

### Excluding Resources by Age

You can use the `--older-than` flag to only nuke resources that were created before a certain period, the possible values are all valid values for [ParseDuration](https://golang.org/pkg/time/#ParseDuration) For example the following command nukes resources that are at least one day old:

```shell
cloud-nuke aws --older-than 24h
```

Happy Nuking!!!

## Credentials

### AWS

In order for the `cloud-nuke` CLI tool to access your AWS, you will need to provide your AWS credentials. You can used one of the [standard AWS CLI credential mechanisms](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html).

## Running Tests

```shell
go test -v ./...
```

## License

This code is released under the MIT License. See [LICENSE.txt](https://github.com/terraform-modules-krish/cloud-nuke/blob/v0.1.1/LICENSE.txt).
