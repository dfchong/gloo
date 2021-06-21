// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/matcher/v3/string.proto

package v3

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *StringMatcher) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.matcher.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3.StringMatcher")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetIgnoreCase())
	if err != nil {
		return 0, err
	}

	switch m.MatchPattern.(type) {

	case *StringMatcher_Exact:

		if _, err = hasher.Write([]byte(m.GetExact())); err != nil {
			return 0, err
		}

	case *StringMatcher_Prefix:

		if _, err = hasher.Write([]byte(m.GetPrefix())); err != nil {
			return 0, err
		}

	case *StringMatcher_Suffix:

		if _, err = hasher.Write([]byte(m.GetSuffix())); err != nil {
			return 0, err
		}

	case *StringMatcher_SafeRegex:

		if h, ok := interface{}(m.GetSafeRegex()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SafeRegex")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSafeRegex(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SafeRegex")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListStringMatcher) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.matcher.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3.ListStringMatcher")); err != nil {
		return 0, err
	}

	for _, v := range m.GetPatterns() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}