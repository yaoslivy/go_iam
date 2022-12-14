package v1

import (
	"github.com/marmotedu/component-base/pkg/scheme"
)

// GroupName is the group name use in this package.
// If use a public domain name, need set the GroupName to service name.
// For example: if restful path is: https://marmotedu.com/apimachinery/v1/secrets, we can set GroupName="apimachinery".
const GroupName = "iam.api"

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = scheme.GroupVersion{Group: GroupName, Version: "v1"}

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) scheme.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
