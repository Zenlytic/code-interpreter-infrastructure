// Code generated by ent, DO NOT EDIT.

package models

import (
	"time"

	"github.com/e2b-dev/infra/packages/shared/pkg/models/env"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/envalias"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/envbuild"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/snapshot"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/team"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/teamapikey"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/user"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/usersteams"
	"github.com/e2b-dev/infra/packages/shared/pkg/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	envFields := schema.Env{}.Fields()
	_ = envFields
	// envDescCreatedAt is the schema descriptor for created_at field.
	envDescCreatedAt := envFields[1].Descriptor()
	// env.DefaultCreatedAt holds the default value on creation for the created_at field.
	env.DefaultCreatedAt = envDescCreatedAt.Default.(func() time.Time)
	// envDescUpdatedAt is the schema descriptor for updated_at field.
	envDescUpdatedAt := envFields[2].Descriptor()
	// env.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	env.DefaultUpdatedAt = envDescUpdatedAt.Default.(func() time.Time)
	// envDescBuildCount is the schema descriptor for build_count field.
	envDescBuildCount := envFields[5].Descriptor()
	// env.DefaultBuildCount holds the default value on creation for the build_count field.
	env.DefaultBuildCount = envDescBuildCount.Default.(int32)
	// envDescSpawnCount is the schema descriptor for spawn_count field.
	envDescSpawnCount := envFields[6].Descriptor()
	// env.DefaultSpawnCount holds the default value on creation for the spawn_count field.
	env.DefaultSpawnCount = envDescSpawnCount.Default.(int64)
	envaliasFields := schema.EnvAlias{}.Fields()
	_ = envaliasFields
	// envaliasDescIsRenamable is the schema descriptor for is_renamable field.
	envaliasDescIsRenamable := envaliasFields[2].Descriptor()
	// envalias.DefaultIsRenamable holds the default value on creation for the is_renamable field.
	envalias.DefaultIsRenamable = envaliasDescIsRenamable.Default.(bool)
	envbuildFields := schema.EnvBuild{}.Fields()
	_ = envbuildFields
	// envbuildDescCreatedAt is the schema descriptor for created_at field.
	envbuildDescCreatedAt := envbuildFields[1].Descriptor()
	// envbuild.DefaultCreatedAt holds the default value on creation for the created_at field.
	envbuild.DefaultCreatedAt = envbuildDescCreatedAt.Default.(func() time.Time)
	// envbuildDescUpdatedAt is the schema descriptor for updated_at field.
	envbuildDescUpdatedAt := envbuildFields[2].Descriptor()
	// envbuild.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	envbuild.DefaultUpdatedAt = envbuildDescUpdatedAt.Default.(func() time.Time)
	// envbuildDescKernelVersion is the schema descriptor for kernel_version field.
	envbuildDescKernelVersion := envbuildFields[12].Descriptor()
	// envbuild.DefaultKernelVersion holds the default value on creation for the kernel_version field.
	envbuild.DefaultKernelVersion = envbuildDescKernelVersion.Default.(string)
	// envbuildDescFirecrackerVersion is the schema descriptor for firecracker_version field.
	envbuildDescFirecrackerVersion := envbuildFields[13].Descriptor()
	// envbuild.DefaultFirecrackerVersion holds the default value on creation for the firecracker_version field.
	envbuild.DefaultFirecrackerVersion = envbuildDescFirecrackerVersion.Default.(string)
	snapshotFields := schema.Snapshot{}.Fields()
	_ = snapshotFields
	// snapshotDescCreatedAt is the schema descriptor for created_at field.
	snapshotDescCreatedAt := snapshotFields[1].Descriptor()
	// snapshot.DefaultCreatedAt holds the default value on creation for the created_at field.
	snapshot.DefaultCreatedAt = snapshotDescCreatedAt.Default.(func() time.Time)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescCreatedAt is the schema descriptor for created_at field.
	teamDescCreatedAt := teamFields[1].Descriptor()
	// team.DefaultCreatedAt holds the default value on creation for the created_at field.
	team.DefaultCreatedAt = teamDescCreatedAt.Default.(func() time.Time)
	// teamDescEmail is the schema descriptor for email field.
	teamDescEmail := teamFields[7].Descriptor()
	// team.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	team.EmailValidator = teamDescEmail.Validators[0].(func(string) error)
	teamapikeyFields := schema.TeamAPIKey{}.Fields()
	_ = teamapikeyFields
	// teamapikeyDescCreatedAt is the schema descriptor for created_at field.
	teamapikeyDescCreatedAt := teamapikeyFields[1].Descriptor()
	// teamapikey.DefaultCreatedAt holds the default value on creation for the created_at field.
	teamapikey.DefaultCreatedAt = teamapikeyDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	usersteamsFields := schema.UsersTeams{}.Fields()
	_ = usersteamsFields
	// usersteamsDescIsDefault is the schema descriptor for is_default field.
	usersteamsDescIsDefault := usersteamsFields[2].Descriptor()
	// usersteams.DefaultIsDefault holds the default value on creation for the is_default field.
	usersteams.DefaultIsDefault = usersteamsDescIsDefault.Default.(bool)
}
