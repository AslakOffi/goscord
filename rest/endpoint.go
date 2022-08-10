package rest

const (
	BaseUrl    = "https://discord.com/api/v9"
	GatewayUrl = "wss://gateway.discord.gg/?v=9&encoding=json"

	// Audit Log
	EndpointGetGuildsAuditLog = "/guilds/%s/audit-logs"

	// Auto Moderation
	EndpointListAutoModerationRulesForGuild = "/guilds/%s/auto-moderation/rules"
	EndpointGetAutoModerationRule           = "/guilds/%s/auto-moderation/rules/%s"
	EndpointCreateAutoModerationRule        = "/guilds/%s/auto-moderation/rules"
	EndpointModifyAutoModerationRule        = "/guilds/%s/auto-moderation/rules/%s"
	EndpointDeleteAutoModerationRule        = "/guilds/%s/auto-moderation/rules/%s"

	// Channel
	EndpointGetChannel                       = "/channels/%s"
	EndpointModifyChannel                    = "/channels/%s"
	EndpointDeleteChannel                    = "/channels/%s"
	EndpointGetChannelMessages               = "/channels/%s/messages"
	EndpointGetChannelMessage                = "/channels/%s/messages/%s"
	EndpointCreateMessage                    = "/channels/%s/messages"
	EndpointCrosspostMessage                 = "/channels/%s/messages/%s/crosspost"
	EndpointOwnReaction                      = "/channels/%s/messages/%s/reactions/%s/@me"
	EndpointDeleteUserReaction               = "/channels/%s/messages/%s/reactions/%s/%s"
	EndpointGetReactions                     = "/channels/%s/messages/%s/reactions/%s"
	EndpointDeleteAllReactions               = "/channels/%s/messages/%s/reactions"
	EndpointDeleteAllReactionsForEmoji       = "/channels/%s/messages/%s/reactions/%s"
	EndpointEditMessage                      = "/channels/%s/messages/%s"
	EndpointDeleteMessage                    = "/channels/%s/messages/%s"
	EndpointBulkDeleteMessages               = "/channels/%s/messages/bulk-delete"
	EndpointEditChannelPermissions           = "/channels/%s/permissions/%s"
	EndpointGetChannelInvites                = "/channels/%s/invites"
	EndpointCreateChannelInvite              = "/channels/%s/invites"
	EndpointDeleteChannelPermission          = "/channels/%s/permissions/%s"
	EndpointFollowNewsChannel                = "/channels/%s/followers"
	EndpointTriggerTypingIndicator           = "/channels/%s/typing"
	EndpointGetPinnedMessages                = "/channels/%s/pins"
	EndpointPinMessage                       = "/channels/%s/pins/%s"
	EndpointUnpinMessage                     = "/channels/%s/pins/%s"
	EndpointGroupDMAddRecipient              = "/channels/%s/recipients/%s"
	EndpointGroupDMRemoveRecipient           = "/channels/%s/recipients/%s"
	EndpointStartThreadFromMessage           = "/channels/%s/messages/%s/threads"
	EndpointStartThreadWithoutMessage        = "/channels/%s/threads"
	EndpointStartThreadInForumChannel        = "/channels/%s/threads"
	EndpointJoinThread                       = "/channels/%s/thread-members/@me"
	EndpointAddThreadMember                  = "/channels/%s/thread-members/%s"
	EndpointLeaveThread                      = "/channels/%s/thread-members/@me"
	EndpointRemoveThreadMember               = "/channels/%s/thread-members/%s"
	EndpointGetThreadMember                  = "/channels/%s/thread-members/%s"
	EndpointListThreadMembers                = "/channels/%s/thread-members"
	EndpointListPublicArchivedThreads        = "/channels/%s/threads/archived/public"
	EndpointListPrivateArchivedThreads       = "/channels/%s/threads/archived/private"
	EndpointListJoinedPrivateArchivedThreads = "/channels/%s/users/@me/threads/archived/private"

	// Emoji
	EndpointListGuildEmojis  = "/guilds/%s/emojis"
	EndpointGetGuildEmoji    = "/guilds/%s/emojis/%s"
	EndpointCreateGuildEmoji = "/guilds/%s/emojis"
	EndpointModifyGuildEmoji = "/guilds/%s/emojis/%s"
	EndpointDeleteGuildEmoji = "/guilds/%s/emojis/%s"

	// Guild
	EndpointCreateGuild                 = "/guilds"
	EndpointGetGuild                    = "/guilds/%s"
	EndpointGetGuildPreview             = "/guilds/%s/preview"
	EndpointModify                      = "/guilds/%s"
	EndpointDeleteGuild                 = "/guilds/%s"
	EndpointGetGuildChannels            = "/guilds/%s/channels"
	EndpointCreateGuildChannel          = "/guilds/%s/channels"
	EndpointModifyChannelPositions      = "/guilds/%s/channels"
	EndpointListActiveGuildThreads      = "/guilds/%s/threads/active"
	EndpointGetGuildMember              = "/guilds/%s/members/%s"
	EndpointListGuildMembers            = "/guilds/%s/members"
	EndpointSearchGuildMembers          = "/guilds/%s/members/search"
	EndpointAddGuildMember              = "/guilds/%s/members/%s"
	EndpointModifyGuildMember           = "/guilds/%s/members/%s"
	EndpointModifyCurrentMember         = "/guilds/%s/members/@me"
	EndpointModifyCurrentNick           = "/guilds/%s/members/@me/nick" // Deprecated
	EndpointAddGuildMemberRole          = "/guilds/%s/members/%s/roles/%s"
	EndpointRemoveGuildMemberRole       = "/guilds/%s/members/%s/roles/%s"
	EndpointRemoveGuildMember           = "/guilds/%s/members/%s"
	EndpointGetGuildBans                = "/guilds/%s/bans"
	EndpointGetGuildBan                 = "/guilds/%s/bans/%s"
	EndpointCreateGuildBan              = "/guilds/%s/bans/%s"
	EndpointDeleteGuildBan              = "/guilds/%s/bans/%s"
	EndpointGetGuildRoles               = "/guilds/%s/roles"
	EndpointCreateGuildRole             = "/guilds/%s/roles"
	EndpointModifyGuildRolePositions    = "/guilds/%s/roles"
	EndpointModifyGuildRole             = "/guilds/%s/roles/%s"
	EndpointModifyGuildMfaLevel         = "/guilds/%s/mfa"
	EndpointDeleteGuildRole             = "/guilds/%s/roles/%s"
	EndpointGetGuildPruneCount          = "/guilds/%s/prune"
	EndpointBeginGuildPrune             = "/guilds/%s/prune"
	EndpointGetGuildVoiceRegions        = "/guilds/%s/regions"
	EndpointGetGuildInvites             = "/guilds/%s/invites"
	EndpointGetGuildIntegrations        = "/guilds/%s/integrations"
	EndpointDeleteGuildIntegration      = "/guilds/%s/integrations/%s"
	EndpointGetGuildWidgetSettings      = "/guilds/%s/widget"
	EndpointModifyGuildWidgetSettings   = "/guilds/%s/widget"
	EndpointGetGuildWidget              = "/guilds/%s/widget.json"
	EndpointGetGuildVanityURL           = "/guilds/%s/vanity-url"
	EndpointModifyGuildWidgetImage      = "/guilds/%s/widget.png"
	EndpointGetGuildWelcomeScreen       = "/guilds/%s/welcome-screen"
	EndpointModifyGuildWelcomeScreen    = "/guilds/%s/welcome-screen"
	EndpointModifyCurrentUserVoiceState = "/guilds/%s/voice/@me/state"
	EndpointModifyUserVoiceState        = "/guilds/%s/voice/%s/state"

	// Guild Scheduled Event
	EndpointListScheduledGuildEvents    = "/guilds/%s/scheduled-events"
	EndpointCreateScheduledGuildEvent   = "/guilds/%s/scheduled-events"
	EndpointGetScheduledGuildEvent      = "/guilds/%s/scheduled-events/%s"
	EndpointModifyScheduledGuildEvent   = "/guilds/%s/scheduled-events/%s"
	EndpointDeleteScheduledGuildEvent   = "/guilds/%s/scheduled-events/%s"
	EndpointGetGuildScheduledEventUsers = "/guilds/%s/scheduled-events/%s/users"

	// Guild Template
	EndpointGetGuildTemplate     = "/guilds/templates/%s"
	EndpointCreateGuildTemplate  = "/guilds/templates/%s"
	EndpointGetGuildTemplates    = "/guilds/%s/templates"
	EndpointCreateGuildTemplates = "/guilds/%s/templates"
	EndpointSyncGuildTemplates   = "/guilds/%s/templates/%s"
	EndpointModifyGuildTemplate  = "/guilds/%s/templates/%s"
	EndpointDeleteGuildTemplate  = "/guilds/%s/templates/%s"

	// Interaction
	EndpointRegisterGlobalCommand = "/applications/%s/commands"
	EndpointRegisterGuildCommand  = "/applications/%s/guilds/%s/commands"

	// Invite
	EndpointGetInvite    = "/invites/%s"
	EndpointDeleteInvite = "/invites/%s"

	// Stage Instance
	EndpointCreateStageInstance = "/stage-instances"
	EndpointGetStageInstance    = "/stage-instances/%s"
	EndpointModifyStageInstance = "/stage-instances/%s"
	EndpointDeleteStageInstance = "/stage-instances/%s"

	// Sticker
	EndpointGetSticker            = "/stickers/%s"
	EndpointListNitroStickerPacks = "/stickers/packs"
	EndpointListGuildStickers     = "/guilds/%s/stickers"
	EndpointGetGuildSticker       = "/guilds/%s/stickers/%s"
	EndpointCreateGuildSticker    = "/guilds/%s/stickers"
	EndpointModifyGuildSticker    = "/guilds/%s/stickers/%s"
	EndpointDeleteGuildSticker    = "/guilds/%s/stickers/%s"

	// User
	EndpointGetCurrentUser             = "/users/@me"
	EndpointGetUser                    = "/users/%s"
	EndpointModifyCurrentUser          = "/users/@me"
	EndpointGetCurrentUserGuilds       = "/users/@me/guilds"
	EndpointGetCurrentUserGuildMembers = "/users/@me/guilds/%s/members"
	EndpointLeaveGuild                 = "/users/@me/guilds/%s"
	EndpointCreateDmChannel            = "/users/@me/channels"
	EndpointCreateGroupDm              = "/users/@me/channels"
	EndpointGetUserConnections         = "/users/@me/connections"

	// Voice
	EndpointListVoiceRegions = "/voice/regions"

	// Webhook
	EndpointCreateWebhook                  = "/channels/%s/webhooks"
	EndpointGetChannelWebhooks             = "/channels/%s/webhooks"
	EndpointGetGuildWebhooks               = "/guilds/%s/webhooks"
	EndpointGetWebhook                     = "/webhooks/%s"
	EndpointGetWebhookWithToken            = "/webhooks/%s/%s"
	EndpointModifyWebhook                  = "/webhooks/%s"
	EndpointModifyWebhookWithToken         = "/webhooks/%s/%s"
	EndpointDeleteWebhook                  = "/webhooks/%s"
	EndpointDeleteWebhookWithToken         = "/webhooks/%s/%s"
	EndpointExecuteWebhook                 = "/webhooks/%s/%s"
	EndpointExecuteSlackCompatibleWebhook  = "/webhooks/%s/%s/slack"
	EndpointExecuteGithubCompatibleWebhook = "/webhooks/%s/%s/github"
	EndpointGetWebhookMessageHistory       = "/webhooks/%s/%s/messages/%s"
	EndpointGetWebhookMessage              = "/webhooks/%s/%s/messages/%s"
	EndpointEditWebhookMessage             = "/webhooks/%s/%s/messages/%s"
	EndpointDeleteWebhookMessage           = "/webhooks/%s/%s/messages/%s"
)
