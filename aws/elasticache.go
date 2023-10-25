package aws

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/terraform-modules-krish/cloud-nuke/config"
	"github.com/terraform-modules-krish/cloud-nuke/logging"
	"github.com/terraform-modules-krish/go-commons/errors"
)

// Returns a formatted string of Elasticache cluster Ids
func getAllElasticacheClusters(session *session.Session, region string, excludeAfter time.Time, configObj config.Config) ([]*string, error) {
	svc := elasticache.New(session)
	result, err := svc.DescribeCacheClusters(&elasticache.DescribeCacheClustersInput{})
	if err != nil {
		return nil, errors.WithStackTrace(err)
	}

	var clusterIds []*string
	for _, cluster := range result.CacheClusters {
		if shouldIncludeElasticacheCluster(cluster, excludeAfter, configObj) {
			clusterIds = append(clusterIds, cluster.CacheClusterId)
		}
	}

	return clusterIds, nil
}

func shouldIncludeElasticacheCluster(cluster *elasticache.CacheCluster, excludeAfter time.Time, configObj config.Config) bool {
	if cluster == nil {
		return false
	}

	if excludeAfter.Before(aws.TimeValue(cluster.CacheClusterCreateTime)) {
		return false
	}

	return config.ShouldInclude(
		aws.StringValue(cluster.CacheClusterId),
		configObj.Elasticache.IncludeRule.NamesRegExp,
		configObj.Elasticache.ExcludeRule.NamesRegExp,
	)
}

func nukeAllElasticacheClusters(session *session.Session, clusterIds []*string) error {
	svc := elasticache.New(session)

	if len(clusterIds) == 0 {
		logging.Logger.Infof("No Elasticache clusters to nuke in region %s", *session.Config.Region)
		return nil
	}

	logging.Logger.Infof("Deleting %d Elasticache clusters in region %s", len(clusterIds), *session.Config.Region)

	var deletedClusterIds []*string
	for _, clusterId := range clusterIds {
		params := elasticache.DeleteCacheClusterInput{
			CacheClusterId: clusterId,
		}

		_, err := svc.DeleteCacheCluster(&params)
		if err != nil {
			logging.Logger.Errorf("[Failed] %s", err)
		} else {
			deletedClusterIds = append(deletedClusterIds, clusterId)
			logging.Logger.Infof("Deleted Elasticache cluster: %s", *clusterId)
		}
	}

	if len(deletedClusterIds) > 0 {
		logging.Logger.Infof("Confirming deletion of %d Elasticache clusters in region %s", len(deletedClusterIds), *session.Config.Region)

		for _, clusterId := range deletedClusterIds {
			params := elasticache.DescribeCacheClustersInput{
				CacheClusterId: clusterId,
			}

			err := svc.WaitUntilCacheClusterDeleted(&params)
			if err != nil {
				logging.Logger.Errorf("[Failed] %s", err)
			}
		}
	}

	logging.Logger.Infof("[OK] %d Elasticache clusters deleted in %s", len(deletedClusterIds), *session.Config.Region)
	return nil
}
