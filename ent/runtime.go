// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/suyuan32/simple-admin-job/ent/schema"
	"github.com/suyuan32/simple-admin-job/ent/task"
	"github.com/suyuan32/simple-admin-job/ent/tasklog"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	taskMixin := schema.Task{}.Mixin()
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskMixinFields1 := taskMixin[1].Fields()
	_ = taskMixinFields1
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskMixinFields0[1].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskMixinFields0[2].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	// taskDescStatus is the schema descriptor for status field.
	taskDescStatus := taskMixinFields1[0].Descriptor()
	// task.DefaultStatus holds the default value on creation for the status field.
	task.DefaultStatus = taskDescStatus.Default.(uint8)
	tasklogFields := schema.TaskLog{}.Fields()
	_ = tasklogFields
	// tasklogDescStartedAt is the schema descriptor for started_at field.
	tasklogDescStartedAt := tasklogFields[1].Descriptor()
	// tasklog.DefaultStartedAt holds the default value on creation for the started_at field.
	tasklog.DefaultStartedAt = tasklogDescStartedAt.Default.(func() time.Time)
}
