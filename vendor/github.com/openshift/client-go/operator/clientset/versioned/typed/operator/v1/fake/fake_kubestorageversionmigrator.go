// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeStorageVersionMigrators implements KubeStorageVersionMigratorInterface
type FakeKubeStorageVersionMigrators struct {
	Fake *FakeOperatorV1
}

var kubestorageversionmigratorsResource = v1.SchemeGroupVersion.WithResource("kubestorageversionmigrators")

var kubestorageversionmigratorsKind = v1.SchemeGroupVersion.WithKind("KubeStorageVersionMigrator")

// Get takes name of the kubeStorageVersionMigrator, and returns the corresponding kubeStorageVersionMigrator object, and an error if there is any.
func (c *FakeKubeStorageVersionMigrators) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.KubeStorageVersionMigrator, err error) {
	emptyResult := &v1.KubeStorageVersionMigrator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(kubestorageversionmigratorsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.KubeStorageVersionMigrator), err
}

// List takes label and field selectors, and returns the list of KubeStorageVersionMigrators that match those selectors.
func (c *FakeKubeStorageVersionMigrators) List(ctx context.Context, opts metav1.ListOptions) (result *v1.KubeStorageVersionMigratorList, err error) {
	emptyResult := &v1.KubeStorageVersionMigratorList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(kubestorageversionmigratorsResource, kubestorageversionmigratorsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.KubeStorageVersionMigratorList{ListMeta: obj.(*v1.KubeStorageVersionMigratorList).ListMeta}
	for _, item := range obj.(*v1.KubeStorageVersionMigratorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeStorageVersionMigrators.
func (c *FakeKubeStorageVersionMigrators) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(kubestorageversionmigratorsResource, opts))
}

// Create takes the representation of a kubeStorageVersionMigrator and creates it.  Returns the server's representation of the kubeStorageVersionMigrator, and an error, if there is any.
func (c *FakeKubeStorageVersionMigrators) Create(ctx context.Context, kubeStorageVersionMigrator *v1.KubeStorageVersionMigrator, opts metav1.CreateOptions) (result *v1.KubeStorageVersionMigrator, err error) {
	emptyResult := &v1.KubeStorageVersionMigrator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(kubestorageversionmigratorsResource, kubeStorageVersionMigrator, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.KubeStorageVersionMigrator), err
}

// Update takes the representation of a kubeStorageVersionMigrator and updates it. Returns the server's representation of the kubeStorageVersionMigrator, and an error, if there is any.
func (c *FakeKubeStorageVersionMigrators) Update(ctx context.Context, kubeStorageVersionMigrator *v1.KubeStorageVersionMigrator, opts metav1.UpdateOptions) (result *v1.KubeStorageVersionMigrator, err error) {
	emptyResult := &v1.KubeStorageVersionMigrator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(kubestorageversionmigratorsResource, kubeStorageVersionMigrator, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.KubeStorageVersionMigrator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeStorageVersionMigrators) UpdateStatus(ctx context.Context, kubeStorageVersionMigrator *v1.KubeStorageVersionMigrator, opts metav1.UpdateOptions) (result *v1.KubeStorageVersionMigrator, err error) {
	emptyResult := &v1.KubeStorageVersionMigrator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(kubestorageversionmigratorsResource, "status", kubeStorageVersionMigrator, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.KubeStorageVersionMigrator), err
}

// Delete takes name of the kubeStorageVersionMigrator and deletes it. Returns an error if one occurs.
func (c *FakeKubeStorageVersionMigrators) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubestorageversionmigratorsResource, name, opts), &v1.KubeStorageVersionMigrator{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeStorageVersionMigrators) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(kubestorageversionmigratorsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.KubeStorageVersionMigratorList{})
	return err
}

// Patch applies the patch and returns the patched kubeStorageVersionMigrator.
func (c *FakeKubeStorageVersionMigrators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KubeStorageVersionMigrator, err error) {
	emptyResult := &v1.KubeStorageVersionMigrator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(kubestorageversionmigratorsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.KubeStorageVersionMigrator), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied kubeStorageVersionMigrator.
func (c *FakeKubeStorageVersionMigrators) Apply(ctx context.Context, kubeStorageVersionMigrator *operatorv1.KubeStorageVersionMigratorApplyConfiguration, opts metav1.ApplyOptions) (result *v1.KubeStorageVersionMigrator, err error) {
	if kubeStorageVersionMigrator == nil {
		return nil, fmt.Errorf("kubeStorageVersionMigrator provided to Apply must not be nil")
	}
	data, err := json.Marshal(kubeStorageVersionMigrator)
	if err != nil {
		return nil, err
	}
	name := kubeStorageVersionMigrator.Name
	if name == nil {
		return nil, fmt.Errorf("kubeStorageVersionMigrator.Name must be provided to Apply")
	}
	emptyResult := &v1.KubeStorageVersionMigrator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(kubestorageversionmigratorsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.KubeStorageVersionMigrator), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeKubeStorageVersionMigrators) ApplyStatus(ctx context.Context, kubeStorageVersionMigrator *operatorv1.KubeStorageVersionMigratorApplyConfiguration, opts metav1.ApplyOptions) (result *v1.KubeStorageVersionMigrator, err error) {
	if kubeStorageVersionMigrator == nil {
		return nil, fmt.Errorf("kubeStorageVersionMigrator provided to Apply must not be nil")
	}
	data, err := json.Marshal(kubeStorageVersionMigrator)
	if err != nil {
		return nil, err
	}
	name := kubeStorageVersionMigrator.Name
	if name == nil {
		return nil, fmt.Errorf("kubeStorageVersionMigrator.Name must be provided to Apply")
	}
	emptyResult := &v1.KubeStorageVersionMigrator{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(kubestorageversionmigratorsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.KubeStorageVersionMigrator), err
}
