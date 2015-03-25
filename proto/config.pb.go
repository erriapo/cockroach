// Code generated by protoc-gen-gogo.
// source: cockroach/proto/config.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

// Attributes specifies a list of arbitrary strings describing
// node topology, store type, and machine capabilities.
type Attributes struct {
	Attrs            []string `protobuf:"bytes,1,rep,name=attrs" json:"attrs" yaml:"attrs,flow"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Attributes) Reset()         { *m = Attributes{} }
func (m *Attributes) String() string { return proto1.CompactTextString(m) }
func (*Attributes) ProtoMessage()    {}

func (m *Attributes) GetAttrs() []string {
	if m != nil {
		return m.Attrs
	}
	return nil
}

// Replica describes a replica location by node ID (corresponds to a
// host:port via lookup on gossip network), store ID (identifies the
// device) and associated attributes. Replicas are stored in Range
// lookup records (meta1, meta2).
type Replica struct {
	NodeID  NodeID  `protobuf:"varint,1,opt,name=node_id,customtype=NodeID" json:"node_id"`
	StoreID StoreID `protobuf:"varint,2,opt,name=store_id,customtype=StoreID" json:"store_id"`
	// Combination of node & store attributes.
	Attrs            Attributes `protobuf:"bytes,3,opt,name=attrs" json:"attrs"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Replica) Reset()         { *m = Replica{} }
func (m *Replica) String() string { return proto1.CompactTextString(m) }
func (*Replica) ProtoMessage()    {}

func (m *Replica) GetAttrs() Attributes {
	if m != nil {
		return m.Attrs
	}
	return Attributes{}
}

// RangeDescriptor is the value stored in a range metadata key.
// A range is described using an inclusive start key, a non-inclusive end key,
// and a list of replicas where the range is stored.
type RangeDescriptor struct {
	RaftID int64 `protobuf:"varint,1,opt,name=raft_id" json:"raft_id"`
	// StartKey is the first key which may be contained by this range.
	StartKey Key `protobuf:"bytes,2,opt,name=start_key,customtype=Key" json:"start_key"`
	// EndKey marks the end of the range's possible keys.  EndKey itself is not
	// contained in this range - it will be contained in the immediately
	// subsequent range.
	EndKey Key `protobuf:"bytes,3,opt,name=end_key,customtype=Key" json:"end_key"`
	// Replicas is the set of replicas on which this range is stored, the
	// ordering being arbitrary and subject to permutation.
	Replicas         []Replica `protobuf:"bytes,4,rep,name=replicas" json:"replicas"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *RangeDescriptor) Reset()         { *m = RangeDescriptor{} }
func (m *RangeDescriptor) String() string { return proto1.CompactTextString(m) }
func (*RangeDescriptor) ProtoMessage()    {}

func (m *RangeDescriptor) GetRaftID() int64 {
	if m != nil {
		return m.RaftID
	}
	return 0
}

func (m *RangeDescriptor) GetReplicas() []Replica {
	if m != nil {
		return m.Replicas
	}
	return nil
}

// GCPolicy defines garbage collection policies which apply to MVCC
// values within a zone.
//
// TODO(spencer): flesh this out to include maximum number of values
//   as well as whether there's an intersection between max values
//   and TTL or a union.
type GCPolicy struct {
	// TTLSeconds specifies the maximum age of a value before it's
	// garbage collected. Only older versions of values are garbage
	// collected. Specifying <=0 mean older versions are never GC'd.
	TTLSeconds       int32  `protobuf:"varint,1,opt,name=ttl_seconds" json:"ttl_seconds"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GCPolicy) Reset()         { *m = GCPolicy{} }
func (m *GCPolicy) String() string { return proto1.CompactTextString(m) }
func (*GCPolicy) ProtoMessage()    {}

func (m *GCPolicy) GetTTLSeconds() int32 {
	if m != nil {
		return m.TTLSeconds
	}
	return 0
}

// AcctConfig holds accounting configuration.
type AcctConfig struct {
	ClusterId        string `protobuf:"bytes,1,opt,name=cluster_id" json:"cluster_id" yaml:"cluster_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AcctConfig) Reset()         { *m = AcctConfig{} }
func (m *AcctConfig) String() string { return proto1.CompactTextString(m) }
func (*AcctConfig) ProtoMessage()    {}

func (m *AcctConfig) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

// PermConfig holds permission configuration, specifying read/write ACLs.
type PermConfig struct {
	// ACL lists users with read permissions.
	Read []string `protobuf:"bytes,1,rep,name=read" json:"read" yaml:"read,omitempty"`
	// ACL lists users with write permissions.
	Write            []string `protobuf:"bytes,2,rep,name=write" json:"write" yaml:"write,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PermConfig) Reset()         { *m = PermConfig{} }
func (m *PermConfig) String() string { return proto1.CompactTextString(m) }
func (*PermConfig) ProtoMessage()    {}

func (m *PermConfig) GetRead() []string {
	if m != nil {
		return m.Read
	}
	return nil
}

func (m *PermConfig) GetWrite() []string {
	if m != nil {
		return m.Write
	}
	return nil
}

// ZoneConfig holds configuration that is needed for a range of KV pairs.
type ZoneConfig struct {
	// ReplicaAttrs is a slice of Attributes, each describing required attributes
	// for each replica in the zone. The order in which the attributes are stored
	// in ReplicaAttrs is arbitrary and may change.
	ReplicaAttrs  []Attributes `protobuf:"bytes,1,rep,name=replica_attrs" json:"replica_attrs" yaml:"replicas,omitempty"`
	RangeMinBytes int64        `protobuf:"varint,2,opt,name=range_min_bytes" json:"range_min_bytes" yaml:"range_min_bytes,omitempty"`
	RangeMaxBytes int64        `protobuf:"varint,3,opt,name=range_max_bytes" json:"range_max_bytes" yaml:"range_max_bytes,omitempty"`
	// If GC policy is not set, uses the next highest, non-null policy
	// in the zone config hierarchy, up to the default policy if necessary.
	GC               *GCPolicy `protobuf:"bytes,4,opt,name=gc" json:"gc,omitempty" yaml:"gc,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ZoneConfig) Reset()         { *m = ZoneConfig{} }
func (m *ZoneConfig) String() string { return proto1.CompactTextString(m) }
func (*ZoneConfig) ProtoMessage()    {}

func (m *ZoneConfig) GetReplicaAttrs() []Attributes {
	if m != nil {
		return m.ReplicaAttrs
	}
	return nil
}

func (m *ZoneConfig) GetRangeMinBytes() int64 {
	if m != nil {
		return m.RangeMinBytes
	}
	return 0
}

func (m *ZoneConfig) GetRangeMaxBytes() int64 {
	if m != nil {
		return m.RangeMaxBytes
	}
	return 0
}

func (m *ZoneConfig) GetGC() *GCPolicy {
	if m != nil {
		return m.GC
	}
	return nil
}

// RangeTree holds the root node and size of the range tree.
type RangeTree struct {
	RootKey          Key    `protobuf:"bytes,1,opt,name=root_key,customtype=Key" json:"root_key"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RangeTree) Reset()         { *m = RangeTree{} }
func (m *RangeTree) String() string { return proto1.CompactTextString(m) }
func (*RangeTree) ProtoMessage()    {}

// RangeTreeNode holds the configuration for each node of the Red-Black Tree that references all ranges.
type RangeTreeNode struct {
	Key Key `protobuf:"bytes,1,opt,name=key,customtype=Key" json:"key"`
	// Color is black if true, red if false.
	Black bool `protobuf:"varint,2,opt,name=black" json:"black"`
	// If the parent key is null, this is the root node.
	ParentKey        Key    `protobuf:"bytes,3,opt,name=parent_key,customtype=Key" json:"parent_key"`
	LeftKey          *Key   `protobuf:"bytes,4,opt,name=left_key,customtype=Key" json:"left_key,omitempty"`
	RightKey         *Key   `protobuf:"bytes,5,opt,name=right_key,customtype=Key" json:"right_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RangeTreeNode) Reset()         { *m = RangeTreeNode{} }
func (m *RangeTreeNode) String() string { return proto1.CompactTextString(m) }
func (*RangeTreeNode) ProtoMessage()    {}

func (m *RangeTreeNode) GetBlack() bool {
	if m != nil {
		return m.Black
	}
	return false
}

func init() {
}
