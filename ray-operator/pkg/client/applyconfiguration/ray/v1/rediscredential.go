// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// RedisCredentialApplyConfiguration represents a declarative configuration of the RedisCredential type for use
// with apply.
type RedisCredentialApplyConfiguration struct {
	ValueFrom *corev1.EnvVarSource `json:"valueFrom,omitempty"`
	Value     *string              `json:"value,omitempty"`
}

// RedisCredentialApplyConfiguration constructs a declarative configuration of the RedisCredential type for use with
// apply.
func RedisCredential() *RedisCredentialApplyConfiguration {
	return &RedisCredentialApplyConfiguration{}
}

// WithValueFrom sets the ValueFrom field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ValueFrom field is set to the value of the last call.
func (b *RedisCredentialApplyConfiguration) WithValueFrom(value corev1.EnvVarSource) *RedisCredentialApplyConfiguration {
	b.ValueFrom = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *RedisCredentialApplyConfiguration) WithValue(value string) *RedisCredentialApplyConfiguration {
	b.Value = &value
	return b
}
