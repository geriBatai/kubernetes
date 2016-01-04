/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package userpolicy

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/watch"
)

// Registry is an interface implemented by things that know how to store UserPolicy objects.
type Registry interface {
	ListUserPolicies(ctx api.Context, options *api.ListOptions) (*extensions.UserPolicyList, error)
	WatchUserPolicy(ctx api.Context, options *api.ListOptions) (watch.Interface, error)
	GetUserPolicy(ctx api.Context, name string) (*extensions.UserPolicy, error)
	CreateUserPolicy(ctx api.Context, resource *extensions.UserPolicy) (*extensions.UserPolicy, error)
	UpdateUserPolicy(ctx api.Context, resource *extensions.UserPolicy) (*extensions.UserPolicy, error)
	DeleteUserPolicy(ctx api.Context, name string) error
}

// storage puts strong typing around storage calls
type storage struct {
	rest.StandardStorage
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched
// types will panic.
func NewRegistry(s rest.StandardStorage) Registry {
	return &storage{s}
}

func (s *storage) ListUserPolicies(ctx api.Context, options *api.ListOptions) (*extensions.UserPolicyList, error) {
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*extensions.UserPolicyList), nil
}

func (s *storage) WatchUserPolicy(ctx api.Context, options *api.ListOptions) (watch.Interface, error) {
	return s.Watch(ctx, options)
}

func (s *storage) GetUserPolicy(ctx api.Context, name string) (*extensions.UserPolicy, error) {
	obj, err := s.Get(ctx, name)
	if err != nil {
		return nil, err
	}
	return obj.(*extensions.UserPolicy), nil
}

func (s *storage) CreateUserPolicy(ctx api.Context, policy *extensions.UserPolicy) (*extensions.UserPolicy, error) {
	obj, err := s.Create(ctx, policy)
	return obj.(*extensions.UserPolicy), err
}

func (s *storage) UpdateUserPolicy(ctx api.Context, policy *extensions.UserPolicy) (*extensions.UserPolicy, error) {
	obj, _, err := s.Update(ctx, policy)
	return obj.(*extensions.UserPolicy), err
}

func (s *storage) DeleteUserPolicy(ctx api.Context, name string) error {
	_, err := s.Delete(ctx, name, nil)
	return err
}
