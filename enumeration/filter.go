package enumeration

import "github.com/snyk/driftctl/enumeration/resource"

type Filter interface {
	IsTypeIgnored(ty resource.ResourceType) bool
	IsResourceIgnored(res *resource.Resource) bool
	IsFieldIgnored(res *resource.Resource, path []string) bool
}