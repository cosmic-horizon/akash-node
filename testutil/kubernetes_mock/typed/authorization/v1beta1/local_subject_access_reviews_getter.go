// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
)

// LocalSubjectAccessReviewsGetter is an autogenerated mock type for the LocalSubjectAccessReviewsGetter type
type LocalSubjectAccessReviewsGetter struct {
	mock.Mock
}

// LocalSubjectAccessReviews provides a mock function with given fields: namespace
func (_m *LocalSubjectAccessReviewsGetter) LocalSubjectAccessReviews(namespace string) v1beta1.LocalSubjectAccessReviewInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.LocalSubjectAccessReviewInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.LocalSubjectAccessReviewInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.LocalSubjectAccessReviewInterface)
		}
	}

	return r0
}
