//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package agent_test

import (
	"testing"
	"time"

	"github.com/ligato/cn-infra/agent"
	"github.com/ligato/cn-infra/core"
	. "github.com/onsi/gomega"
)

const (
	testVersion        = "1.0"
	testMaxStartupTime = 1 * time.Hour
)

func TestVersion(t *testing.T) {
	RegisterTestingT(t)
	agent := agent.NewAgent(agent.Version(testVersion))
	Expect(agent).ToNot(BeNil())
	Expect(agent.Options()).ToNot(BeNil())
	Expect(agent.Options().Version).To(Equal(testVersion))

}

func TestMaxStartupTime(t *testing.T) {
	RegisterTestingT(t)
	agent := agent.NewAgent(agent.MaxStartupTime(testMaxStartupTime))
	Expect(agent).ToNot(BeNil())
	Expect(agent.Options()).ToNot(BeNil())
	Expect(agent.Options().MaxStartupTime).To(Equal(testMaxStartupTime))
}

func TestDescendantPluginsNoDep(t *testing.T) {
	RegisterTestingT(t)
	plugin := &PluginNoDeps{}
	agent := agent.NewAgent(agent.DescendantPlugins(plugin))
	Expect(agent).ToNot(BeNil())
	Expect(agent.Options()).ToNot(BeNil())
	Expect(agent.Options().Plugins).ToNot(BeNil())
	Expect(len(agent.Options().Plugins)).To(Equal(1))
	Expect(agent.Options().Plugins[0]).To(Equal(plugin))
}

func TestDescendantPluginsOneLevelDep(t *testing.T) {
	RegisterTestingT(t)

	plugin := &PluginOneDep{}
	agent := agent.NewAgent(agent.DescendantPlugins(plugin))
	Expect(agent).ToNot(BeNil())
	Expect(agent.Options()).ToNot(BeNil())
	Expect(agent.Options().Plugins).ToNot(BeNil())
	Expect(len(agent.Options().Plugins)).To(Equal(2))
	Expect(agent.Options().Plugins[0]).To(Equal(&plugin.Plugin2))
	Expect(agent.Options().Plugins[1]).To(Equal(plugin))
}

func TestDescendantPluginsTwoLevelsDeep(t *testing.T) {
	RegisterTestingT(t)
	plugin := &PluginTwoLevelDeps{}
	agent := agent.NewAgent(agent.DescendantPlugins(plugin))
	Expect(agent).ToNot(BeNil())
	Expect(agent.Options()).ToNot(BeNil())
	Expect(agent.Options().Plugins).ToNot(BeNil())
	Expect(len(agent.Options().Plugins)).To(Equal(3))
	Expect(agent.Options().Plugins[0]).To(Equal(&plugin.PluginTwoLevelDep1))
	Expect(agent.Options().Plugins[1]).To(Equal(&plugin.PluginTwoLevelDep1.Plugin2))
	Expect(agent.Options().Plugins[2]).To(Equal(plugin))

}

// Various Test Structs after this point

// PluginNoDeps contains no plugins.
type PluginNoDeps struct {
	pluginName core.PluginName
	Plugin1    MissignCloseMethod
	Plugin2    struct {
		Dep1B string
	}
}

func (p *PluginNoDeps) Init() error  { return nil }
func (p *PluginNoDeps) Close() error { return nil }
func (p *PluginNoDeps) Name() string { return string(p.pluginName) }

// PluginOneDep contains one plugin (another is missing Close method).
type PluginOneDep struct {
	pluginName core.PluginName
	Plugin1    MissignCloseMethod
	Plugin2    TestPlugin
}

func (p *PluginOneDep) Init() error  { return nil }
func (p *PluginOneDep) Close() error { return nil }
func (p *PluginOneDep) Name() string { return string(p.pluginName) }

type PluginTwoLevelDeps struct {
	pluginName         core.PluginName
	PluginTwoLevelDep1 PluginOneDep
	PluginTwoLevelDep2 TestPlugin
}

func (p *PluginTwoLevelDeps) Init() error  { return nil }
func (p *PluginTwoLevelDeps) Close() error { return nil }
func (p *PluginTwoLevelDeps) Name() string { return string(p.pluginName) }

// MissignCloseMethod implements only Init() but not Close() method.
type MissignCloseMethod struct {
}

// Init does nothing.
func (*MissignCloseMethod) Init() error {
	return nil
}
