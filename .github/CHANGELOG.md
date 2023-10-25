# Changelog


# v0.9.2
_Released Feb 1, 2022_
https://github.com/gruntwork-io/cloud-nuke/pull/257 : `cloud-nuke` now can remove customer-managed in AWS Key Management Service

<br>

# v0.9.1
_Released Jan 20, 2022_
#272 VPC nuking now respects the `--older-than` flag
<br>

# v0.9.0
_Released Jan 18, 2022_
https://github.com/gruntwork-io/cloud-nuke/pull/271: `cloud-nuke` will now delete IAM OpenID Connect (OIDC) Providers. If you wish to avoid nuking IAM OIDC Providers, you can either pass in `--exclude-resource-type oidcprovider`, or specify a [config file](https://github.com/gruntwork-io/cloud-nuke#config-file).
<br>

# v0.8.2
_Released Jan 13, 2022_
#269 cloud-nuke now has config file support for VPCs
<br>

# v0.8.1
_Released Jan 12, 2022_

* #262: Added pagination for lambda functions listing

<br>

# v0.8.0
_Released Jan 12, 2022_
`cloud-nuke` will now delete Elasticache instances. If you wish to avoid nuking Elasticache instances, you can either pass in `--exclude-resource-type elasticache`, or specify a [config file](https://github.com/gruntwork-io/cloud-nuke#config-file).

NOTE: This is equivalent to `v0.7.5`, but released as backward incompatible to indicate a new resource was added to prevent users from inadvertently updating to a new patch release without realizing a new resource is being included in the nuking.
<br>

# v0.7.5
_Released Jan 12, 2022_
#267 cloud-nuke now supports nuking Elasticache clusters
<br>

# v0.7.4
_Released Jan 11, 2022_
#268 cloud-nuke now deletes endpoints when nuking VPCs.
<br>

# v0.7.3
_Released Jan 3, 2022_
https://github.com/gruntwork-io/cloud-nuke/pull/265: Added config file support for filtering by name to EBS volumes and Lambda functions.
<br>

# v0.7.2
_Released Jan 3, 2022_
https://github.com/gruntwork-io/cloud-nuke/pull/258 & https://github.com/gruntwork-io/cloud-nuke/pull/266: Added config file support for filtering by name to DynamoDB, ELBv2, ECS service and ECS cluster.

## Special thanks

Special thanks to @brandonstrohmeyer for their contribution!
<br>

# v0.7.1
_Released Nov 17, 2021_
#243: Adds support for deleting non-default VPCs. Thanks to @ekristen for this contribution!
<br>

# v0.7.0
_Released Nov 16, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/238: `cloud-nuke` will now delete OpenSearch Service domains (formerly called AWS Elasticsearch Service). If you wish to avoid nuking OpenSearch domains, you can either pass in `--exclude-resource-type opensearch`, or specify [a config file](https://github.com/gruntwork-io/cloud-nuke#config-file).

https://github.com/gruntwork-io/cloud-nuke/pull/234: Fixed bug where `cloud-nuke` did not nuke EKS clusters that had compute resources. With this release, `cloud-nuke` will now handle nuking Managed Node groups and Fargate profiles attached to EKS clusters eligible for nuking.
<br>

# v0.6.0
_Released Nov 12, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/233: Cloud nuke will now delete CloudWatch Dashboards. If you wish to avoid nuking Dashboards, you can either pass in `--exclude-resource-type cloudwatch-dashboard`, or specify [a config file](https://github.com/gruntwork-io/cloud-nuke#config-file).
<br>

# v0.5.2
_Released Nov 3, 2021_
#169: Add support for nuking DynamoDB
#223 Limit parallelism of gox to stabilize build stage
#221 Update circleci config fix 32 bit builds & use GW utils
#214, #216 Fix `nuke-config` for internal Gruntwork usage 

Special thanks to @DMEvanCT for their contribution!
<br>

# v0.5.1
_Released Sep 29, 2021_
#205: Fix Config file not including when there is a 'exclude' rule
#209: Adds new default enabled region (ap-northeast-3)

Special thanks to @ashwiniag for their contribution!
<br>

# v0.5.0
_Released Aug 24, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/201: Cloud nuke will now delete AWS ACM Private Certificate Authorities. If you wish to avoid nuking PCAs, you can pass in `--exclude-resource-type acmpca`.

Special thanks to @weitzj for their contribution!
<br>

# v0.4.0
_Released Jul 26, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/202: Cloud nuke will now delete AWS IAM Access Analyzers. If you wish to avoid nuking IAM Access Analyzers, you can either pass in `--exclude-resource-type accessanalyzer`, or specify [a config file](https://github.com/gruntwork-io/cloud-nuke#config-file).
<br>

# v0.32.1-test-signing-binaries
_Released Aug 11, 2023_
Testing changes from https://github.com/gruntwork-io/cloud-nuke/pull/559
<br>

# v0.32.0
_Released Jul 3, 2023_
## What's Changed
* Implement nukeAllRdsDbSubnetGroups to delete all RDS subnet groups by @hongil0316 in https://github.com/gruntwork-io/cloud-nuke/pull/472
* Fix errors in config_test by @hongil0316 in https://github.com/gruntwork-io/cloud-nuke/pull/481
* add sns filter by name and time by @robpickerill in https://github.com/gruntwork-io/cloud-nuke/pull/482
* Add ACM Support by @robpickerill in https://github.com/gruntwork-io/cloud-nuke/pull/466
* Add codedeploy  by @robpickerill in https://github.com/gruntwork-io/cloud-nuke/pull/469
* fix(aws): always session from externalcreds by @dixneuf19 in https://github.com/gruntwork-io/cloud-nuke/pull/464

## New Contributors
* @robpickerill made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/482
* @dixneuf19 made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/464

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.31.2...v0.32.0

## Migration Guide

* New resources have been added in this release, please update your config files to support exclusion rules for Code Deploy, ACM, SNS, and RDS Subnet Groups
<br>

# v0.31.2
_Released Jun 15, 2023_
## What's Changed
* fix iam groups nuke job by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/475
* Bugfix/test fixes by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/477
* update go commons by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/479


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.31.1...v0.31.2
<br>

# v0.31.1
_Released May 11, 2023_
## What's Changed
* Hot patch: disable telemetry for library usage by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/460
This fixes a nil pointer dereference when cloud-nuke was used as a library (via `inspect` functionality).

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.31.0...v0.31.1
<br>

# v0.31.0
_Released May 9, 2023_
## What's Changed
* Add support for Security Hub - CORE-364 by @arsci in https://github.com/gruntwork-io/cloud-nuke/pull/458
* Update Macie support - CORE-387 by @arsci in https://github.com/gruntwork-io/cloud-nuke/pull/459


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.30.0...v0.31.0
<br>

# v0.30.0
_Released May 3, 2023_
## What's Changed
* Don't log to `STDOUT` when being used as a library by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/424
* Add support for elasticache parameter and subnet groups by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/455
* add support for redshift by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/456


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.29.7...v0.30.0
<br>

# v0.3.0
_Released Jun 30, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/197: Cloud nuke will now delete NAT gateways. If you wish to avoid nuking NAT gateways, you can either pass in `--exclude-resource-type nat-gateway`, or specify [a config file](https://github.com/gruntwork-io/cloud-nuke#config-file).
<br>

# v0.29.7
_Released Apr 20, 2023_
## What's Changed

*  Fixed ECS Cluster Service Daemon deletion


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.29.6...v0.29.7
<br>

# v0.29.6
_Released Apr 20, 2023_
## What's Changed
* Fix for cloud-nuke telemetry errors by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/453


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.29.5...v0.29.6
<br>

# v0.29.5
_Released Apr 19, 2023_
## What's Changed
* Make telemetry more obvious and reduce event count by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/452


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.29.4...v0.29.5
<br>

# v0.29.4
_Released Apr 13, 2023_
## What's Changed
* Bug: Filter out AMIs created by AWS Backup - CORE-664 by @arsci in https://github.com/gruntwork-io/cloud-nuke/pull/440
* Add functionality to delete KMS keys that dont have aliases by @MoonMoon1919 in https://github.com/gruntwork-io/cloud-nuke/pull/441
* Add `--delete-unaliased-kms-keys` to CI nuke config by @MoonMoon1919 in https://github.com/gruntwork-io/cloud-nuke/pull/448


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.29.3...v0.29.4
<br>

# v0.29.3
_Released Apr 11, 2023_
## What's Changed
* chore(CORE-659): Improve filtering for EBS Volumes and Snapshots by @MoonMoon1919 in https://github.com/gruntwork-io/cloud-nuke/pull/439

## New Contributors
* @MoonMoon1919 made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/439

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.29.2...v0.29.3
<br>

# v0.29.2
_Released Apr 7, 2023_
## What's Changed
* Migrate to using aws-sdk-go-v2 for Kinesis Data Streams by @jbleduigou in https://github.com/gruntwork-io/cloud-nuke/pull/425
* Only run the nuke function if there are resources in the region to nuke by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/438

## New Contributors
* @jbleduigou made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/425

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.29.1...v0.29.2
<br>

# v0.29.1
_Released Apr 5, 2023_
## What's Changed
* Fix secrets manager can nuke secret replica. by @Jennas-Lee in https://github.com/gruntwork-io/cloud-nuke/pull/436
* Fix tests by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/437


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.29.0...v0.29.1
<br>

# v0.29.0
_Released Apr 5, 2023_
## What's Changed
* CORE-641 - Add telemetry to main nuke flow by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/434


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.28.1...v0.29.0

## Migration Guide
This release adds telemetry and sends metrics back to gruntwork. No Identifying information is recorded. The following metrics are included:

- Command and Arguments
- Version Number
- Timestamps
- Resource Types
- Resource Counts
- A randomly generated Run ID
- AWS Account ID

We never collect:

- IP Addresses
- Resource Names
- Telemetry can be disabled entirely by setting the DISABLE_TELEMETRY environment variable on the command line.
<br>

# v0.28.1
_Released Apr 4, 2023_
## What's Changed
* Simplify presentation of which resources are supported by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/433


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.28.0...v0.28.1
<br>

# v0.28.0
_Released Apr 4, 2023_
## What's Changed
* Add cloudwatch alarm support (#429) by @Jennas-Lee in https://github.com/gruntwork-io/cloud-nuke/pull/430


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.27.2...v0.28.0
<br>

# v0.27.2
_Released Apr 4, 2023_
## What's Changed
* [skip ci] Refactor contexts by @eak12913 in https://github.com/gruntwork-io/cloud-nuke/pull/427
* Fix VPC endpoint removal failure by @Jennas-Lee in https://github.com/gruntwork-io/cloud-nuke/pull/432

## New Contributors
* @Jennas-Lee made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/432

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.27.1...v0.27.2
<br>

# v0.27.1
_Released Mar 20, 2023_
## What's Changed
* Backport UI enhancements to defaults-aws cmd by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/407


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.27.0...v0.27.1
<br>

# v0.27.0
_Released Mar 14, 2023_
## What's Changed
* Exclude service-linked IAM roles by @pete0emerson in https://github.com/gruntwork-io/cloud-nuke/pull/418
* fix exclusion rules for roles and policies by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/419
* Update Go Commons to support logging upgrade by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/422


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.26.0...v0.27.0

## Migration Guide

* If using cloud nuke as a library, you'll need to update to version 1.18 of Go
<br>

# v0.26.0
_Released Feb 28, 2023_
* Implement support for IAM service-linked Roles
<br>

# v0.25.0
_Released Feb 6, 2023_
## What's Changed
* Implement support for configservice recorders and rules by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/405


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.24.0...v0.25.0
<br>

# v0.24.0
_Released Feb 1, 2023_
## What's Changed
* Added FUNDING.yml by @eak12913 in https://github.com/gruntwork-io/cloud-nuke/pull/397
* Document RDS, Neptune, Document DB by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/398
* `defaults-aws` also nukes the SG's IPv6 egress rule by @marinalimeira in https://github.com/gruntwork-io/cloud-nuke/pull/401
* Implement Launch Template Support by @oredavids in https://github.com/gruntwork-io/cloud-nuke/pull/403

## New Contributors
* @eak12913 made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/397
* @oredavids made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/403

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.23.0...v0.24.0
<br>

# v0.23.0
_Released Jan 20, 2023_
## What's Changed
* Support nuking EC2 Dedicated Hosts - CORE-286 by @arsci in https://github.com/gruntwork-io/cloud-nuke/pull/392
* Cloud nuke failing tests fixes by @denis256 in https://github.com/gruntwork-io/cloud-nuke/pull/390
* Change the order of cloud-nuke to delete IAM policy first before IAM groups by @hongil0316 in https://github.com/gruntwork-io/cloud-nuke/pull/393
* Implement ECR support by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/396


**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.22.1...v0.23.0
<br>

# v0.22.1
_Released Dec 15, 2022_
## What's Changed
* Update README.md with instructions on how to use with AWS profile by @brianpham in https://github.com/gruntwork-io/cloud-nuke/pull/343
* Fix CODEOWNERS by @pete0emerson in https://github.com/gruntwork-io/cloud-nuke/pull/386
* add more patterns to nuke by @Etiene in https://github.com/gruntwork-io/cloud-nuke/pull/387
* update worker size and build parallelism by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/388

## New Contributors
* @brianpham made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/343
* @pete0emerson made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/386

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.22.0...v0.22.1
<br>

# v0.22.0
_Released Dec 14, 2022_
## What's Changed
* Temporarily avoid nuking IAM groups and policies in GW test accounts by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/372
* Fix whitespace by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/373
* Update pull request template with check re: exclusions by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/375
* Support use of --log-level flag in inspect-aws command by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/377
* UI overhaul and per-resource error reporting by @zackproser in https://github.com/gruntwork-io/cloud-nuke/pull/356
* Support nuking IAM EC2 instance profiles with roles - CORE-285 by @arsci in https://github.com/gruntwork-io/cloud-nuke/pull/378
* CORE-298 Use new UI with IAM groups and IAM policies by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/381
* CORE-188 rate limits by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/382
* Add support to delete ec2 key pairs using cloud-nuke by @hongil0316 in https://github.com/gruntwork-io/cloud-nuke/pull/379
* cloud-nuke fixes by @denis256 in https://github.com/gruntwork-io/cloud-nuke/pull/384
* Support deletion of KMS Key Aliases - CORE-292 by @arsci in https://github.com/gruntwork-io/cloud-nuke/pull/383
* Bugfix/core 271 vpc with eni by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/385

## New Contributors
* @arsci made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/378
* @hongil0316 made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/379

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.21.0...v0.22.0

<br>

# v0.21.0
_Released Oct 31, 2022_
## What's Changed
* Implement support for nuking and inspecting iam groups and customer managed policies by @ellisonc in https://github.com/gruntwork-io/cloud-nuke/pull/364

## New Contributors
* @ellisonc made their first contribution in https://github.com/gruntwork-io/cloud-nuke/pull/364

**Full Changelog**: https://github.com/gruntwork-io/cloud-nuke/compare/v0.20.0...v0.21.0
<br>

# v0.20.0
_Released Sep 29, 2022_
Implement support for nuking and inspecting Cloudtrail Trails.
<br>

# v0.2.0
_Released Jun 25, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/194: Cloud nuke will now delete secrets manager secrets. If you wish to avoid nuking secrets manager entries, you can either pass in `--exclude-resource-type secretsmanager`, or specify [a config file](https://github.com/gruntwork-io/cloud-nuke#config-file).

## Release practice change

⚠️ Starting this release, new resources nuked by cloud-nuke will be marked as a minor version bump (`X` in `v0.X.Y`) to indicate backward incompatibilities. Previously, all new resources were considered backward compatible, but since new resources are included automatically (opt-out vs opt-in), users with CI practices around `cloud-nuke` would be surprised by new resources that are suddenly being picked up for deletion. This surprise is stronger for resources that are actively in use for any account, such as IAM Users.

As such, we have decided to mark the addition of new nuked resources as backward incompatible releases to provide better signals for users when we introduce a new resource.
<br>

# v0.19.1
_Released Sep 27, 2022_
Fix a bug breaking API Gateway v2 nuking behavior. The wrong underlying delete method was being called. 
<br>

# v0.19.0
_Released Aug 23, 2022_
Implement support for inspecting and nuking SNS Topics.

## Links 
* #354 
<br>

# v0.18.0
_Released Aug 23, 2022_
Implement support for inspecting and nuking Elastic FileSystems (EFS)

## Links
* #353 
<br>

# v0.17.0
_Released Aug 16, 2022_
Implement support for inspecting and nuking API Gateway (v1 and v2).

## Links 
* #352 
<br>

# v0.16.4
_Released Aug 12, 2022_
## Description

- Fixed bug where global resources were not nuked due to a broken security token issue when using GovCloud.

## Related links

https://github.com/gruntwork-io/cloud-nuke/pull/350
<br>

# v0.16.3
_Released Aug 12, 2022_
Fix Elasticache nuking behavior by looking up replication groups as well as single cache clusters in our list method. 

## Links 

* #348 
<br>

# v0.16.2
_Released Jul 26, 2022_
Fix Elasticache nuking behavior by implementing support for nuking cache clusters that are members of replication groups. 

## Links 
* #335 

<br>

# v0.16.1
_Released Jul 22, 2022_
Fix VPC nuking behavior by supporting deletion of chained security groups in `nukeSecurityGroups`. 

## Links 
* #317 

## Special thanks 

Special thanks to @wsscc2021 for their contribution!
<br>

# v0.16.0
_Released Jul 21, 2022_
Implement support for inspecting and nuking Kinesis Stream instances.

## Links 
- https://github.com/gruntwork-io/cloud-nuke/pull/304

## Special thanks 

Special thanks to @strongishllama for their contribution!
<br>

# v0.15.0
_Released Jul 19, 2022_
Implement support for inspecting and nuking Sagemaker notebook instances.

## Links 
- https://github.com/gruntwork-io/cloud-nuke/pull/282
- https://github.com/gruntwork-io/cloud-nuke/pull/332

## Special thanks 

Special thanks to @jvanbuel for their contribution!
<br>

# v0.14.0
_Released Jul 19, 2022_
Implement support for inspecting and nuking IAM Roles

## Links 
#330 

## Special thanks 

Special thanks to @ekristen for their contribution via #251 which was adapted into this release.
<br>

# v0.13.0
_Released Jul 18, 2022_
Implement support for inspecting and nuking AWS Macie member accounts

## Links
#323 
<br>

# v0.12.3
_Released Jul 15, 2022_
* Introduce support for passing a pointer to an external `aws.Config` object when using cloud-nuke as a scripting library via the new `externalcreds.Set` method. This is useful for supplying AWS credentials to your cloud-nuke scripting jobs via code. 

## Links
* #326 
<br>

# v0.12.2
_Released Jul 7, 2022_
## Description
- Fixes a bug in the NAT Gateway nuke routine where 404 errors were ignored.

## Links

- https://github.com/gruntwork-io/cloud-nuke/pull/325
<br>

# v0.12.1
_Released Jul 7, 2022_
## Description

- Updated the NAT Gateway lookup method so that it does not retrieve deleted or in deletion resources.

## Special Thanks

Special thanks to the following users for their contribution:

- @ekristen 
- @GiampaoloFalqui 

## Links

- https://github.com/gruntwork-io/cloud-nuke/pull/248
- https://github.com/gruntwork-io/cloud-nuke/pull/313
<br>

# v0.12.0
_Released Jul 7, 2022_
https://github.com/gruntwork-io/cloud-nuke/pull/320: `cloud-nuke` will now delete GuardDuty Detectors. If you wish to avoid nuking GuardDuty Detectors, you can pass in `--exclude-resource-type guardduty`.
<br>

# v0.11.8
_Released Jun 8, 2022_
#305 Extend cloud-nuke with non-destructive inspect command and library methods

This release adds a non-destructive command, `inspect-aws`, which accepts most of the same flags as `nuke` but will never delete resources. As `aws` / nuke is to `DELETE`, `aws-inspect` is to `LIST` or `READ`. 

You can use `inspect-aws` to list and search for resources across your AWS accounts. 

In addition, this release adds library functionality that allows you to script against `cloud-nuke`. See [Using cloud-nuke as a library](https://github.com/gruntwork-io/cloud-nuke/blob/371d196f12891dad2208d0fda051e30e637c6e26/README.md#using-cloud-nuke-as-a-library) for more information. 

This release fixes the regression bug with `--exclude-resource-type` that was introduced in `v0.11.7`
<br>

# v0.11.7
_Released Jun 7, 2022_
## IMPORTANT

A major regression bug has been identified in this release, where `--exclude-resource-type` no longer works as intended. We advise users not to use this version to avoid nuking resources unintentionally.

Please use `v0.11.8`

<br>

# v0.11.6
_Released May 9, 2022_
https://github.com/gruntwork-io/cloud-nuke/pull/294: Removal of S3 BucketPolicy before nuke bucket

# Special thanks
Special thanks to @maxhaensel for their contribution!
<br>

# v0.11.5
_Released Apr 27, 2022_
#293 :  VPCEs deletion before subnets removal

# Special thanks
Special thanks to @sHesl for their contribution!
<br>

# v0.11.4
_Released Apr 25, 2022_
https://github.com/gruntwork-io/cloud-nuke/pull/295: Added support for config file to control EKS cluster nuking.
<br>

# v0.11.3
_Released Mar 17, 2022_
#291 updates VPC config to filter by name instead of VPC Id.

# Special thanks
Special thanks to @brandonstrohmeyer for their contribution!
<br>

# v0.11.2
_Released Mar 16, 2022_
#287 ElasticIPs, AutoScalingGroups, LaunchConfigurations and EC2 instances can now be filtered by the config file using their names.

# Special thanks
Special thanks to @brandonstrohmeyer for their contribution!
<br>

# v0.11.1
_Released Mar 15, 2022_
https://github.com/gruntwork-io/cloud-nuke/pull/284: KMS Customer Managed Key deletion now supports the config file format to filter by alias.
<br>

# v0.11.0
_Released Mar 11, 2022_
`cloud-nuke` will now delete CloudWatch Log Groups. If you wish to avoid nuking Log Groups, you can either pass in `--exclude-resource-type cloudwatch-loggroup`, or specify a [config file](https://github.com/gruntwork-io/cloud-nuke#config-file).

## Special thanks

Special thanks to @ekristen for their contribution!
<br>

# v0.10.0
_Released Feb 1, 2022_
`cloud-nuke` will now delete KMS Customer Managed Keys. If you wish to avoid nuking KMS Keys, you can either pass in `--exclude-resource-type kmscustomerkeys`, or specify a [config file](https://github.com/gruntwork-io/cloud-nuke#config-file).

NOTE: This is equivalent to `v0.9.2`, but released as backward incompatible to indicate a new resource was added to prevent users from inadvertently updating to a new patch release without realizing a new resource is being included in the nuking.
<br>

# v0.1.9
_Released Aug 20, 2019_
#65 : Adding support for specifying which resource types to delete and listing resource types.
<br>

# v0.1.7
_Released Mar 4, 2019_
#53: Publish checksums along with binaries.
<br>

# v0.1.6
_Released Feb 26, 2019_
https://github.com/gruntwork-io/cloud-nuke/pull/52 : This fixes a bug in the EKS nuking capabilities where it was not aware of regions that do not support EKS. This release filters out those regions when considering EKS resources.
<br>

# v0.1.5
_Released Feb 21, 2019_
https://github.com/gruntwork-io/cloud-nuke/pull/43 : Cloud nuke will now delete EKS clusters.
<br>

# v0.1.4
_Released Oct 25, 2018_
https://github.com/gruntwork-io/cloud-nuke/pull/36 : Cloud Nuke will now delete ECS services and tasks
<br>

# v0.1.30
_Released May 3, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/184: `cloud-nuke` now supports nuking SQS queues.

Special thanks to @rafaelleonardocruz for the contribution!
<br>

# v0.1.3
_Released Oct 5, 2018_
#34 improves the nuking strategy by ensuring cloud-nuke doesn't return when it encounters an error on just a single resource
#35 adds the ability to nuke launch configurations so we don't hit the 200 count limitation
<br>

# v0.1.29
_Released Apr 15, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/191: Update to the latest versions of Go and the `build-go-binaries` script, so `cloud-nuke` will now build binaries for Darwin ARM 64 CPUs.
<br>

# v0.1.28
_Released Apr 6, 2021_
https://github.com/gruntwork-io/cloud-nuke/pull/189: Sync dependencies in `go.sum`. There should be no change in functionality, but this will hopefully fix build issues, including in Homebrew.
<br>

# v0.1.27
_Released Apr 1, 2021_
Fix a bug where the reported number of deleted AMIs was inaccurate. Thanks to @StandB for this contribution! See #186 for more details.
<br>

# v0.1.26
_Released Mar 31, 2021_
This release updates `cloud-nuke` to nuke [AWS Transit Gateway](https://aws.amazon.com/transit-gateway/). Thanks to @rafaelleonardocruz for this contribution! See #178 for details.
<br>

# v0.1.25
_Released Jan 27, 2021_
This release updates `cloud-nuke` to support [AWS U.S. GovCloud regions](https://aws.amazon.com/govcloud-us/).
<br>

# v0.1.24
_Released Nov 23, 2020_
## Resources affected

- `lambda` **[NEW]**

## Description

Now it's possible to use `--resource-type lambda` to nuke Lambdas.

Thanks @andreybleme for implementing this!

## Related links

- https://github.com/gruntwork-io/cloud-nuke/pull/134
<br>

# v0.1.23
_Released Oct 16, 2020_
## Resources affected

- `ecscluster`

## Description

This release includes:

- update `cloud-nuke` to discover & nuke only ECS clusters in the ACTIVE state 
- fix for s3 tests leaving s3 buckets behind

## Related links

- https://github.com/gruntwork-io/cloud-nuke/pull/143
- https://github.com/gruntwork-io/cloud-nuke/pull/146
<br>

# v0.1.22
_Released Oct 13, 2020_
## Resources affected

- `ecscluster`

## Description

The ability to nuke ECS Clusters is now supported by `cloud-nuke`, allowing users to delete `ecscluster` using the tool, and also to filter by age. 

Example usage:
`cloud-nuke aws --resource-type ecscluster --older-than 10m`

## Related links

- https://github.com/gruntwork-io/cloud-nuke/pull/141
<br>

# v0.1.21
_Released Oct 1, 2020_
## Resources affected

- `s3`

## Description

The nuke routine for `s3` has been optimized for memory efficiency. Previously the routine was paging in and storing all objects in the bucket in memory before starting the deletion routine. Now the routine will delete as objects are paged in.

## Related links

- https://github.com/gruntwork-io/cloud-nuke/pull/135
<br>

# v0.1.20
_Released Jun 15, 2020_
https://github.com/gruntwork-io/cloud-nuke/pull/113: 

- You can now pass in a config file using the `--config` option to express more granular and complex filtering for what gets nuked. 
- See the [README](https://github.com/gruntwork-io/cloud-nuke/blob/4e625a07b3da432e4dbbdd040204b957892eadf3/README.md#config-file) to learn what is supported and how it works.
- PRs welcome!

https://github.com/gruntwork-io/cloud-nuke/pull/110:
- Refactored S3 test cases to remove repetition and use table drive tests.
- S3 tests now run in parallel. 
- Special thanks to @saurabh-hirani for the PR!
<br>

# v0.1.2
_Released Apr 10, 2018_
#25 #27 Fixes #21, cloud-nuke now recovers from RequestLimitExceeded errors
#28 Adds support for nuking Elastic IPs
<br>

# v0.1.19
_Released May 23, 2020_
https://github.com/gruntwork-io/cloud-nuke/issues/117: This release contains no changes. It is used solely to make it possible to install `cloud-nuke` via Homebrew.
<br>

# v0.1.18
_Released Apr 14, 2020_
https://github.com/gruntwork-io/cloud-nuke/pull/105: 

- You can now set the log level using the `--log-level` param.
- `cloud-nuke` now fetches S3 bucket info concurrently, which should significantly speed the process up by 5-10x.

Special thanks to @saurabh-hirani for the PR!
<br>

# v0.1.17
_Released Apr 2, 2020_
https://github.com/gruntwork-io/cloud-nuke/pull/101 & https://github.com/gruntwork-io/cloud-nuke/pull/96: `cloud-nuke` now supports nuking S3 buckets. Note that you can exclude certain buckets from consideration by tagging them with the tag key `cloud-nuke-excluded` and value `true`.

This release also introduces a new CLI arg `--exclude-resource-type` for removing certain resource types from consideration. For example, to avoid nuking s3 buckets entirely, you can pass in `cloud-nuke aws --exclude-resource-type s3`. Use `cloud-nuke aws --list-resource-types` to see all the supported resource types.
<br>

# v0.1.16
_Released Mar 9, 2020_
https://github.com/gruntwork-io/cloud-nuke/pull/100 : Fix potential source of panic errors when nuking RDS resources.
<br>

# v0.1.15
_Released Mar 3, 2020_
https://github.com/gruntwork-io/cloud-nuke/pull/98/: Fix ECS service nuking to support pagination of ECS clusters.

https://github.com/gruntwork-io/cloud-nuke/pull/97: Fix unit tests to remove source of panic.
<br>

# v0.1.13: Improvements to defaults-aws subcommand
_Released Dec 9, 2019_
This release addresses #80 and #81. The `defaults-aws` subcommand now accepts the `--region` and `--exclude-region` flags to reduce the scope when nuking default VPCs and security group rules. It also adds the `--sg-only` subcommand to skip nuking default VPCs and only nuke default security group rules.
<br>

# v0.1.12
_Released Nov 16, 2019_
#78 :Add --dry-run flag to exit after listing resources to nuke

<br>

# v0.1.11
_Released Oct 14, 2019_
#71  : Add ability to nuke only specific regions.
<br>

# v0.1.10
_Released Oct 11, 2019_
#73 : Instead of exiting if the user types in an invalid input to the nuke prompt, the user will be notified that the input was invalid and prompted again.
<br>

# v0.1.1
_Released Mar 22, 2018_
#23 makes cloud provider names subcommands so the standard flags passing structure can be maintained
<br>

# v0.1.0
_Released Mar 20, 2018_
#20, #22 renames this project to `cloud-nuke` and prepares it to support other cloud providers
<br>

# v0.0.4
_Released Mar 15, 2018_
#18 #19 Adds support for nuking user owned AMIs and Snapshots
<br>

# v0.0.3
_Released Feb 22, 2018_
#11, #12 fixes EBS Volume nuking issues and sets aws-nuke to run nightly
<br>

# v0.0.2
_Released Feb 17, 2018_
#8 Resources older than a specific time duration can be specified for nuking
<br>

# v0.0.1
_Released Feb 9, 2018_
#4 #5 First public release
<br>

# Add the defaults-aws command
_Released Aug 15, 2019_
This release adds a new subcommand: `cloud-nuke defaults-aws`.

## Background
New AWS accounts have a Default VPC in each region. If you create a resource (e.g. an EC2 instance) without specifying a VPC ID and subnet, it will be created within the default VPC by default. For production accounts, indeed for any account which is used for purposes other than simple sandbox tests, the default VPC can be considered a risk as it may accumulate unused/undesired resources over time.

Additional, all VPCs, not just default VPCs, include a default security group. The default security group has an ingress rule that allows all traffic within the same group, and an egress rule that permits all outbound traffic (e.g. all ports, all protocols, to destination 0.0.0.0/0). This is an antipattern for obvious reasons. Removing this rule is a requirement for compliance with some security standards, such as the CIS AWS Foundations Benchmark.

## Usage

The `defaults-aws subcommand` removes these defaults. Example output:

```bash
./ cloud-nuke defaults-aws --force 
INFO[2019-08-14T07:52:26-07:00] Identifying enabled regions
INFO[2019-08-14T07:52:27-07:00] Found enabled region eu-north-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region ap-south-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region eu-west-3
INFO[2019-08-14T07:52:27-07:00] Found enabled region eu-west-2
INFO[2019-08-14T07:52:27-07:00] Found enabled region eu-west-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region ap-northeast-2
INFO[2019-08-14T07:52:27-07:00] Found enabled region ap-northeast-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region sa-east-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region ca-central-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region ap-southeast-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region ap-southeast-2
INFO[2019-08-14T07:52:27-07:00] Found enabled region eu-central-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region us-east-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region us-east-2
INFO[2019-08-14T07:52:27-07:00] Found enabled region us-west-1
INFO[2019-08-14T07:52:27-07:00] Found enabled region us-west-2
INFO[2019-08-14T07:52:27-07:00] Discovering default VPCs
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-09f5580fdcbb2f45c eu-north-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-02a9144962943118c ap-south-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-0b8fa626656a2d406 eu-west-3
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-09dc63dd3e9ef59b5 eu-west-2
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-08748b3914dbe20f3 eu-west-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-07e471a2781ce611e ap-northeast-2
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-0d756684f04e8b539 ap-northeast-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-0d6a8e45677eb1adb sa-east-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-0b41d4bc601d51418 ca-central-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-026a2ea431bde67aa ap-southeast-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-08bd2b462fcdb511c ap-southeast-2
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-090846f0c0d7886c7 eu-central-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-033904a71040f7305 us-east-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-0da379ffd7cb43d90 us-east-2
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-001199441593c91bd us-west-1
INFO[2019-08-14T07:52:37-07:00] * Default VPC vpc-0c2dd0d4dbf70d539 us-west-2
INFO[2019-08-14T07:52:37-07:00] Nuking VPC vpc-09f5580fdcbb2f45c in region eu-north-1
INFO[2019-08-14T07:52:38-07:00] ...detaching Internet Gateway igw-01faf4477ddc2aaa1
INFO[2019-08-14T07:52:38-07:00] ...deleting Internet Gateway igw-01faf4477ddc2aaa1
INFO[2019-08-14T07:52:38-07:00] ...deleting subnet subnet-0c57018894f22cd3b
INFO[2019-08-14T07:52:39-07:00] ...deleting subnet subnet-03f333046f700fdc1
INFO[2019-08-14T07:52:39-07:00] ...deleting subnet subnet-0c6b2a41d26840a04
INFO[2019-08-14T07:52:40-07:00] ...deleting Security Group sg-08f25d8ce6b83c2ba
INFO[2019-08-14T07:52:40-07:00] ...deleting VPC vpc-09f5580fdcbb2f45c
INFO[2019-08-14T07:52:40-07:00] Nuking VPC vpc-02a9144962943118c in region ap-south-1
INFO[2019-08-14T07:52:41-07:00] ...detaching Internet Gateway igw-06381f015a43a06f0
INFO[2019-08-14T07:52:41-07:00] ...deleting Internet Gateway igw-06381f015a43a06f0
INFO[2019-08-14T07:52:42-07:00] ...deleting subnet subnet-024ee83d56944842f
INFO[2019-08-14T07:52:42-07:00] ...deleting subnet subnet-07616d8c3e6edf7cb
INFO[2019-08-14T07:52:43-07:00] ...deleting subnet subnet-0549721d4262b3c5e
INFO[2019-08-14T07:52:44-07:00] ...deleting Security Group sg-09f2a83c6e69ec27a
INFO[2019-08-14T07:52:44-07:00] ...deleting VPC vpc-02a9144962943118c
INFO[2019-08-14T07:52:44-07:00] Nuking VPC vpc-0b8fa626656a2d406 in region eu-west-3
INFO[2019-08-14T07:52:45-07:00] ...detaching Internet Gateway igw-078c47bce8bf47af8
INFO[2019-08-14T07:52:45-07:00] ...deleting Internet Gateway igw-078c47bce8bf47af8
[snip]
INFO[2019-08-14T07:53:22-07:00] Finished nuking default VPCs in all regions
INFO[2019-08-14T07:53:22-07:00] Discovering default security groups
INFO[2019-08-14T07:53:25-07:00] * Default rules for SG sg-00d959b8d5d11f48b default eu-west-3
INFO[2019-08-14T07:53:25-07:00] * Default rules for SG sg-070fc543465a33a83 default eu-west-1
INFO[2019-08-14T07:53:25-07:00] * Default rules for SG sg-03947dbd85c595888 default ca-central-1
INFO[2019-08-14T07:53:25-07:00] ...revoking default rules from Security Group sg-00d959b8d5d11f48b
INFO[2019-08-14T07:53:26-07:00] Egress rule not present (ok)
INFO[2019-08-14T07:53:26-07:00] Ingress rule not present (ok)
INFO[2019-08-14T07:53:26-07:00] ...revoking default rules from Security Group sg-070fc543465a33a83
INFO[2019-08-14T07:53:27-07:00] Egress rule not present (ok)
INFO[2019-08-14T07:53:28-07:00] Ingress rule not present (ok)
INFO[2019-08-14T07:53:28-07:00] ...revoking default rules from Security Group sg-03947dbd85c595888
INFO[2019-08-14T07:53:28-07:00] Egress rule not present (ok)
INFO[2019-08-14T07:53:28-07:00] Ingress rule not present (ok)
INFO[2019-08-14T07:53:28-07:00] Finished nuking default Security Groups in all regions
```

The only supported flag is `--force` to allow for non-interactive use.

<br>

# Add support for RDS
_Released Feb 13, 2020_
This release addresses #86 and #94. Now it's possible to use `--resource-type rds` to nuke both RDS Instances and Clusters.
<br>

