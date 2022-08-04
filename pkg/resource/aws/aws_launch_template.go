package aws

import (
	"github.com/snyk/driftctl/enumeration/resource"
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

const AwsLaunchTemplateResourceType = "aws_launch_template"

func initAwsLaunchTemplateMetaData(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetFlags(AwsLaunchTemplateResourceType, resource.FlagDeepMode)
}