//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 Opstree Solutions.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/cafe24-dhkim05/redis-operator/api"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExistingPasswordSecret) DeepCopyInto(out *ExistingPasswordSecret) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExistingPasswordSecret.
func (in *ExistingPasswordSecret) DeepCopy() *ExistingPasswordSecret {
	if in == nil {
		return nil
	}
	out := new(ExistingPasswordSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesConfig) DeepCopyInto(out *KubernetesConfig) {
	*out = *in
	in.KubernetesConfig.DeepCopyInto(&out.KubernetesConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesConfig.
func (in *KubernetesConfig) DeepCopy() *KubernetesConfig {
	if in == nil {
		return nil
	}
	out := new(KubernetesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
	out.Probe = in.Probe
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probe.
func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Redis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCluster) DeepCopyInto(out *RedisCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCluster.
func (in *RedisCluster) DeepCopy() *RedisCluster {
	if in == nil {
		return nil
	}
	out := new(RedisCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisClusterList) DeepCopyInto(out *RedisClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisClusterList.
func (in *RedisClusterList) DeepCopy() *RedisClusterList {
	if in == nil {
		return nil
	}
	out := new(RedisClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisClusterSpec) DeepCopyInto(out *RedisClusterSpec) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
	in.KubernetesConfig.DeepCopyInto(&out.KubernetesConfig)
	if in.ClusterVersion != nil {
		in, out := &in.ClusterVersion, &out.ClusterVersion
		*out = new(string)
		**out = **in
	}
	in.RedisLeader.DeepCopyInto(&out.RedisLeader)
	in.RedisFollower.DeepCopyInto(&out.RedisFollower)
	if in.RedisExporter != nil {
		in, out := &in.RedisExporter, &out.RedisExporter
		*out = new(RedisExporter)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = new([]Sidecar)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Sidecar, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.PersistenceEnabled != nil {
		in, out := &in.PersistenceEnabled, &out.PersistenceEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisClusterSpec.
func (in *RedisClusterSpec) DeepCopy() *RedisClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RedisClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisClusterStatus) DeepCopyInto(out *RedisClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisClusterStatus.
func (in *RedisClusterStatus) DeepCopy() *RedisClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RedisClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConfig) DeepCopyInto(out *RedisConfig) {
	*out = *in
	in.RedisConfig.DeepCopyInto(&out.RedisConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConfig.
func (in *RedisConfig) DeepCopy() *RedisConfig {
	if in == nil {
		return nil
	}
	out := new(RedisConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisExporter) DeepCopyInto(out *RedisExporter) {
	*out = *in
	in.RedisExporter.DeepCopyInto(&out.RedisExporter)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisExporter.
func (in *RedisExporter) DeepCopy() *RedisExporter {
	if in == nil {
		return nil
	}
	out := new(RedisExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFollower) DeepCopyInto(out *RedisFollower) {
	*out = *in
	in.RedisFollower.DeepCopyInto(&out.RedisFollower)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFollower.
func (in *RedisFollower) DeepCopy() *RedisFollower {
	if in == nil {
		return nil
	}
	out := new(RedisFollower)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLeader) DeepCopyInto(out *RedisLeader) {
	*out = *in
	in.RedisLeader.DeepCopyInto(&out.RedisLeader)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLeader.
func (in *RedisLeader) DeepCopy() *RedisLeader {
	if in == nil {
		return nil
	}
	out := new(RedisLeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisList) DeepCopyInto(out *RedisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Redis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisList.
func (in *RedisList) DeepCopy() *RedisList {
	if in == nil {
		return nil
	}
	out := new(RedisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisReplication) DeepCopyInto(out *RedisReplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisReplication.
func (in *RedisReplication) DeepCopy() *RedisReplication {
	if in == nil {
		return nil
	}
	out := new(RedisReplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisReplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisReplicationList) DeepCopyInto(out *RedisReplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisReplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisReplicationList.
func (in *RedisReplicationList) DeepCopy() *RedisReplicationList {
	if in == nil {
		return nil
	}
	out := new(RedisReplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisReplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisReplicationSpec) DeepCopyInto(out *RedisReplicationSpec) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
	in.KubernetesConfig.DeepCopyInto(&out.KubernetesConfig)
	if in.RedisExporter != nil {
		in, out := &in.RedisExporter, &out.RedisExporter
		*out = new(RedisExporter)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisConfig != nil {
		in, out := &in.RedisConfig, &out.RedisConfig
		*out = new(RedisConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = new([]v1.Toleration)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.Toleration, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(Probe)
		**out = **in
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(Probe)
		**out = **in
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = new([]Sidecar)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Sidecar, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisReplicationSpec.
func (in *RedisReplicationSpec) DeepCopy() *RedisReplicationSpec {
	if in == nil {
		return nil
	}
	out := new(RedisReplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisReplicationStatus) DeepCopyInto(out *RedisReplicationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisReplicationStatus.
func (in *RedisReplicationStatus) DeepCopy() *RedisReplicationStatus {
	if in == nil {
		return nil
	}
	out := new(RedisReplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSentinel) DeepCopyInto(out *RedisSentinel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSentinel.
func (in *RedisSentinel) DeepCopy() *RedisSentinel {
	if in == nil {
		return nil
	}
	out := new(RedisSentinel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisSentinel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSentinelConfig) DeepCopyInto(out *RedisSentinelConfig) {
	*out = *in
	in.RedisSentinelConfig.DeepCopyInto(&out.RedisSentinelConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSentinelConfig.
func (in *RedisSentinelConfig) DeepCopy() *RedisSentinelConfig {
	if in == nil {
		return nil
	}
	out := new(RedisSentinelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSentinelList) DeepCopyInto(out *RedisSentinelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisSentinel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSentinelList.
func (in *RedisSentinelList) DeepCopy() *RedisSentinelList {
	if in == nil {
		return nil
	}
	out := new(RedisSentinelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisSentinelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSentinelSpec) DeepCopyInto(out *RedisSentinelSpec) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
	in.KubernetesConfig.DeepCopyInto(&out.KubernetesConfig)
	if in.RedisSentinelConfig != nil {
		in, out := &in.RedisSentinelConfig, &out.RedisSentinelConfig
		*out = new(RedisSentinelConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = new([]v1.Toleration)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.Toleration, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(api.RedisPodDisruptionBudget)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(Probe)
		**out = **in
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(Probe)
		**out = **in
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = new([]Sidecar)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Sidecar, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSentinelSpec.
func (in *RedisSentinelSpec) DeepCopy() *RedisSentinelSpec {
	if in == nil {
		return nil
	}
	out := new(RedisSentinelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSentinelStatus) DeepCopyInto(out *RedisSentinelStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSentinelStatus.
func (in *RedisSentinelStatus) DeepCopy() *RedisSentinelStatus {
	if in == nil {
		return nil
	}
	out := new(RedisSentinelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSpec) DeepCopyInto(out *RedisSpec) {
	*out = *in
	in.KubernetesConfig.DeepCopyInto(&out.KubernetesConfig)
	if in.RedisExporter != nil {
		in, out := &in.RedisExporter, &out.RedisExporter
		*out = new(RedisExporter)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisConfig != nil {
		in, out := &in.RedisConfig, &out.RedisConfig
		*out = new(RedisConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = new([]v1.Toleration)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.Toleration, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(Probe)
		**out = **in
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(Probe)
		**out = **in
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = new([]Sidecar)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Sidecar, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSpec.
func (in *RedisSpec) DeepCopy() *RedisSpec {
	if in == nil {
		return nil
	}
	out := new(RedisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisStatus) DeepCopyInto(out *RedisStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisStatus.
func (in *RedisStatus) DeepCopy() *RedisStatus {
	if in == nil {
		return nil
	}
	out := new(RedisStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConfig) DeepCopyInto(out *ServiceConfig) {
	*out = *in
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConfig.
func (in *ServiceConfig) DeepCopy() *ServiceConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sidecar) DeepCopyInto(out *Sidecar) {
	*out = *in
	in.Sidecar.DeepCopyInto(&out.Sidecar)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sidecar.
func (in *Sidecar) DeepCopy() *Sidecar {
	if in == nil {
		return nil
	}
	out := new(Sidecar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	in.TLSConfig.DeepCopyInto(&out.TLSConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}