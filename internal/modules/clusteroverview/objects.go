package clusteroverview

import (
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/heptio/developer-dash/internal/describer"
	"github.com/heptio/developer-dash/pkg/objectstoreutil"
)

var (
	rbacClusterRoles = describer.NewResource(describer.ResourceOptions{
		Path:           "/rbac/cluster-roles",
		ObjectStoreKey: objectstoreutil.Key{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRole"},
		ListType:       &rbacv1.ClusterRoleList{},
		ObjectType:     &rbacv1.ClusterRole{},
		Titles:         describer.ResourceTitle{List: "RBAC / Cluster Roles", Object: "Cluster Role"},
		ClusterWide:    true,
	})

	rbacClusterRoleBindings = describer.NewResource(describer.ResourceOptions{
		Path:           "/rbac/cluster-role-bindings",
		ObjectStoreKey: objectstoreutil.Key{APIVersion: "rbac.authorization.k8s.io/v1", Kind: "ClusterRoleBinding"},
		ListType:       &rbacv1.ClusterRoleBindingList{},
		ObjectType:     &rbacv1.ClusterRoleBinding{},
		Titles:         describer.ResourceTitle{List: "RBAC / Cluster Role Bindings", Object: "Cluster Role Binding"},
		ClusterWide:    true,
	})

	rbacDescriber = describer.NewSectionDescriber(
		"/rbac",
		"RBAC",
		rbacClusterRoles,
		rbacClusterRoleBindings,
	)

	portForwardDescriber = NewPortForwardListDescriber()

	rootDescriber = describer.NewSectionDescriber(
		"/",
		"Cluster Overview",
		rbacDescriber,
		portForwardDescriber,
	)
)
