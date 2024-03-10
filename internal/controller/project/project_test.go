/*
Copyright 2022 The Crossplane Authors.

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

package project

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	pj "github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/project"
	"github.com/stretchr/testify/mock"
)

// Unlike many Kubernetes projects Crossplane does not use third party testing
// libraries, per the common Go test review comments. Crossplane encourages the
// use of table driven unit tests. The tests of the crossplane-runtime project
// are representative of the testing style Crossplane encourages.
//
// https://github.com/golang/go/wiki/TestComments
// https://github.com/crossplane/crossplane/blob/master/CONTRIBUTING.md#contributing-code

type MockExternal struct {
	mock.Mock
	*apiv2.RESTClient
}

// DeleteProject implements project.Client.
func (m *MockExternal) DeleteProject(ctx context.Context, nameOrID string) error {
	panic("unimplemented")
}

// GetProject implements project.Client.
func (m *MockExternal) GetProject(ctx context.Context, nameOrID string) (*model.Project, error) {
	panic("unimplemented")
}

// ListProjects implements project.Client.
func (m *MockExternal) ListProjects(ctx context.Context, nameFilter string) ([]*model.Project, error) {
	panic("unimplemented")
}

// NewProject implements project.Client.
func (m *MockExternal) NewProject(ctx context.Context, projectRequest *model.ProjectReq) error {
	return nil
}

// ProjectExists implements project.Client.
func (m *MockExternal) ProjectExists(ctx context.Context, nameOrID string) (bool, error) {
	return true, nil
}

// UpdateProject implements project.Client.
func (m *MockExternal) UpdateProject(ctx context.Context, p *model.Project, storageLimit *int64) error {
	panic("unimplemented")
}

var _ pj.Client = &MockExternal{}

type testtt struct {
	service *apiv2.RESTClient
}

func (t *testtt) Obs(ctx context.Context, val string) (bool, error) {
	return t.service.ProjectExists(ctx, val)
}

func TestObserve(t *testing.T) {

	client := &MockExternal{}
	ctx := context.Background()
	client.On("ProjectExists", mock.MatchedBy(func(context.Context, string) (bool, error) { return true, nil }), ctx, "example").Return(true, nil)
	ext := &testtt{
		service: client.RESTClient,
	}
	res, err := ext.Obs(context.Background(), "example")

	if res == true && err == nil {
		t.Log(res)
	}

}
