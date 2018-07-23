// +build !ignore_autogenerated

// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package crdexample

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrdExample) DeepCopyInto(out *CrdExample) {
	*out = *in
	if in.Repeats != nil {
		in, out := &in.Repeats, &out.Repeats
		*out = make([]*CrdExample_CrdExampleEmbed, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CrdExample_CrdExampleEmbed)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrdExample.
func (in *CrdExample) DeepCopy() *CrdExample {
	if in == nil {
		return nil
	}
	out := new(CrdExample)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrdExample_CrdExampleEmbed) DeepCopyInto(out *CrdExample_CrdExampleEmbed) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrdExample_CrdExampleEmbed.
func (in *CrdExample_CrdExampleEmbed) DeepCopy() *CrdExample_CrdExampleEmbed {
	if in == nil {
		return nil
	}
	out := new(CrdExample_CrdExampleEmbed)
	in.DeepCopyInto(out)
	return out
}
