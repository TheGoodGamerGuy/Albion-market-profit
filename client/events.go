package client

//EventType used to identify event types
//go:generate stringer -type=EventType
type EventType uint16

const (
	evUnused EventType = iota
	evLeave
	evJoinFinished
	evMove
	evTeleport
	evChangeEquipment
	evHealthUpdate
	evEnergyUpdate
	evDamageShieldUpdate
	evCraftingFocusUpdate
	evActiveSpellEffectsUpdate
	evResetCooldowns
	evAttack
	evCastStart
	evChannelingUpdate
	evCastCancel
	evCastTimeUpdate
	evCastFinished
	evCastSpell
	evCastSpells
	evCastHit
	evCastHits
	evStoredTargetsUpdate
	evChannelingEnded
	evAttackBuilding
	evInventoryPutItem
	evInventoryDeleteItem
	evNewCharacter
	evNewEquipmentItem
	evNewSimpleItem
	evNewFurnitureItem
	evNewKillTrophyItem
	evNewJournalItem
	evNewLaborerItem
	evNewEquipmentItemLegendarySoul
	evNewSimpleHarvestableObject
	evNewSimpleHarvestableObjectList
	evNewHarvestableObject
	evNewTreasureDestinationObject
	evTreasureDestinationObjectStatus
	evCloseTreasureDestinationObject
	evNewSilverObject
	evNewBuilding
	evHarvestableChangeState
	evMobChangeState
	evFactionBuildingInfo
	evCraftBuildingInfo
	evRepairBuildingInfo
	evMeldBuildingInfo
	evConstructionSiteInfo
	evPlayerBuildingInfo
	evFarmBuildingInfo
	evTutorialBuildingInfo
	evLaborerObjectInfo
	evLaborerObjectJobInfo
	evMarketPlaceBuildingInfo
	evHarvestStart
	evHarvestCancel
	evHarvestFinished
	evTakeSilver
	evActionOnBuildingStart
	evActionOnBuildingCancel
	evActionOnBuildingFinished
	evItemRerollQualityFinished
	evInstallResourceStart
	evInstallResourceCancel
	evInstallResourceFinished
	evCraftItemFinished
	evLogoutCancel
	evChatMessage
	evChatSay
	evChatWhisper
	evChatMuted
	evPlayEmote
	evStopEmote
	evSystemMessage
	evUtilityTextMessage
	evUpdateMoney
	evUpdateFame
	evUpdateLearningPoints
	evUpdateReSpecPoints
	evUpdateCurrency
	evUpdateFactionStanding
	evUpdateMistCityStanding
	evRespawn
	evServerDebugLog
	evCharacterEquipmentChanged
	evRegenerationHealthChanged
	evRegenerationEnergyChanged
	evRegenerationMountHealthChanged
	evRegenerationCraftingChanged
	evRegenerationHealthEnergyComboChanged
	evRegenerationPlayerComboChanged
	evDurabilityChanged
	evNewLoot
	evAttachItemContainer
	evDetachItemContainer
	evInvalidateItemContainer
	evLockItemContainer
	evGuildUpdate
	evGuildPlayerUpdated
	evInvitedToGuild
	evGuildMemberWorldUpdate
	evUpdateMatchDetails
	evObjectEvent
	evNewMonolithObject
	evNewOrbObject
	evNewCastleObject
	evNewSpellEffectArea
	evUpdateSpellEffectArea
	evNewChainSpell
	evUpdateChainSpell
	evNewTreasureChest
	evStartMatch
	evStartArenaMatchInfos
	evEndArenaMatch
	evMatchUpdate
	evActiveMatchUpdate
	evNewMob
	evDebugMobInfo
	evDebugVariablesInfo
	evDebugReputationInfo
	evDebugDiminishingReturnInfo
	evDebugSmartClusterQueueInfo
	evClaimOrbStart
	evClaimOrbFinished
	evClaimOrbCancel
	evOrbUpdate
	evOrbClaimed
	evOrbReset
	evNewWarCampObject
	evNewMatchLootChestObject
	evNewArenaExit
	evGuildMemberTerritoryUpdate
	evInvitedMercenaryToMatch
	evClusterInfoUpdate
	evForcedMovement
	evForcedMovementCancel
	evCharacterStats
	evCharacterStatsKillHistory
	evCharacterStatsDeathHistory
	evGuildStats
	evKillHistoryDetails
	evFullAchievementInfo
	evFinishedAchievement
	evAchievementProgressInfo
	evFullAchievementProgressInfo
	evFullTrackedAchievementInfo
	evFullAutoLearnAchievementInfo
	evQuestGiverQuestOffered
	evQuestGiverDebugInfo
	evConsoleEvent
	evTimeSync
	evChangeAvatar
	evChangeMountSkin
	evGameEvent
	evKilledPlayer
	evDied
	evKnockedDown
	evUnconcious
	evMatchPlayerJoinedEvent
	evMatchPlayerStatsEvent
	evMatchPlayerStatsCompleteEvent
	evMatchTimeLineEventEvent
	evMatchPlayerMainGearStatsEvent
	evMatchPlayerChangedAvatarEvent
	evInvitationPlayerTrade
	evPlayerTradeStart
	evPlayerTradeCancel
	evPlayerTradeUpdate
	evPlayerTradeFinished
	evPlayerTradeAcceptChange
	evMiniMapPing
	evMarketPlaceNotification
	evDuellingChallengePlayer
	evNewDuellingPost
	evDuelStarted
	evDuelEnded
	evDuelDenied
	evDuelRequestCanceled
	evDuelLeftArea
	evDuelReEnteredArea
	evNewRealEstate
	evMiniMapOwnedBuildingsPositions
	evRealEstateListUpdate
	evGuildLogoUpdate
	evGuildLogoChanged
	evPlaceableObjectPlace
	evPlaceableObjectPlaceCancel
	evFurnitureObjectBuffProviderInfo
	evFurnitureObjectCheatProviderInfo
	evFarmableObjectInfo
	evNewUnreadMails
	evMailOperationPossible
	evGuildLogoObjectUpdate
	evStartLogout
	evNewChatChannels
	evJoinedChatChannel
	evLeftChatChannel
	evRemovedChatChannel
	evAccessStatus
	evMounted
	evMountStart
	evMountCancel
	evNewTravelpoint
	evNewIslandAccessPoint
	evNewExit
	evUpdateHome
	evUpdateChatSettings
	evResurrectionOffer
	evResurrectionReply
	evLootEquipmentChanged
	evUpdateUnlockedGuildLogos
	evUpdateUnlockedAvatars
	evUpdateUnlockedAvatarRings
	evUpdateUnlockedBuildings
	evNewIslandManagement
	evNewTeleportStone
	evCloak
	evPartyInvitation
	evPartyJoinRequest
	evPartyJoined
	evPartyDisbanded
	evPartyPlayerJoined
	evPartyChangedOrder
	evPartyPlayerLeft
	evPartyLeaderChanged
	evPartyLootSettingChangedPlayer
	evPartySilverGained
	evPartyPlayerUpdated
	evPartyInvitationAnswer
	evPartyJoinRequestAnswer
	evPartyMarkedObjectsUpdated
	evPartyOnClusterPartyJoined
	evPartySetRoleFlag
	evPartyInviteOrJoinPlayerEquipmentInfo
	evSpellCooldownUpdate
	evNewHellgateExitPortal
	evNewExpeditionExit
	evNewExpeditionNarrator
	evExitEnterStart
	evExitEnterCancel
	evExitEnterFinished
	evNewQuestGiverObject
	evFullQuestInfo
	evQuestProgressInfo
	evQuestGiverInfoForPlayer
	evFullExpeditionInfo
	evExpeditionQuestProgressInfo
	evInvitedToExpedition
	evExpeditionRegistrationInfo
	evEnteringExpeditionStart
	evEnteringExpeditionCancel
	evRewardGranted
	evArenaRegistrationInfo
	evEnteringArenaStart
	evEnteringArenaCancel
	evEnteringArenaLockStart
	evEnteringArenaLockCancel
	evInvitedToArenaMatch
	evUsingHellgateShrine
	evEnteringHellgateLockStart
	evEnteringHellgateLockCancel
	evPlayerCounts
	evInCombatStateUpdate
	evOtherGrabbedLoot
	evTreasureChestUsingStart
	evTreasureChestUsingFinished
	evTreasureChestUsingCancel
	evTreasureChestUsingOpeningComplete
	evTreasureChestForceCloseInventory
	evLocalTreasuresUpdate
	evLootChestSpawnpointsUpdate
	evPremiumChanged
	evPremiumExtended
	evPremiumLifeTimeRewardGained
	evGoldPurchased
	evLaborerGotUpgraded
	evJournalGotFull
	evJournalFillError
	evFriendRequest
	evFriendRequestInfos
	evFriendInfos
	evFriendRequestAnswered
	evFriendOnlineStatus
	evFriendRequestCanceled
	evFriendRemoved
	evFriendUpdated
	evPartyLootItems
	evPartyLootItemsRemoved
	evReputationUpdate
	evDefenseUnitAttackBegin
	evDefenseUnitAttackEnd
	evDefenseUnitAttackDamage
	evUnrestrictedPvpZoneUpdate
	evReputationImplicationUpdate
	evNewMountObject
	evMountHealthUpdate
	evMountCooldownUpdate
	evNewExpeditionAgent
	evNewExpeditionCheckPoint
	evExpeditionStartEvent
	evVoteEvent
	evRatingEvent
	evNewArenaAgent
	evBoostFarmable
	evUseFunction
	evNewPortalEntrance
	evNewPortalExit
	evNewRandomDungeonExit
	evWaitingQueueUpdate
	evPlayerMovementRateUpdate
	evObserveStart
	evMinimapZergs
	evMinimapSmartClusterZergs
	evPaymentTransactions
	evPerformanceStatsUpdate
	evOverloadModeUpdate
	evDebugDrawEvent
	evRecordCameraMove
	evRecordStart
	evClaimPowerCrystalStart
	evClaimPowerCrystalCancel
	evClaimPowerCrystalReset
	evClaimPowerCrystalFinished
	evTerritoryClaimStart
	evTerritoryClaimCancel
	evTerritoryClaimFinished
	evTerritoryScheduleResult
	evTerritoryUpgradeWithPowerCrystalResult
	evReturningPowerCrystalStart
	evReturningPowerCrystalFinished
	evUpdateAccountState
	evStartDeterministicRoam
	evGuildFullAccessTagsUpdated
	evGuildAccessTagUpdated
	evGvgSeasonUpdate
	evGvgSeasonCheatCommand
	evSeasonPointsByKillingBooster
	evFishingStart
	evFishingCast
	evFishingCatch
	evFishingFinished
	evFishingCancel
	evNewFloatObject
	evNewFishingZoneObject
	evFishingMiniGame
	evSteamAchievementCompleted
	evUpdatePuppet
	evChangeFlaggingFinished
	evNewOutpostObject
	evOutpostUpdate
	evOutpostClaimed
	evOverChargeEnd
	evOverChargeStatus
	evPartyFinderFullUpdate
	evPartyFinderUpdate
	evPartyFinderApplicantsUpdate
	evPartyFinderEquipmentSnapshot
	evPartyFinderJoinRequestDeclined
	evNewUnlockedPersonalSeasonRewards
	evPersonalSeasonPointsGained
	evPersonalSeasonPastSeasonDataEvent
	evEasyAntiCheatMessageToClient
	evMatchLootChestOpeningStart
	evMatchLootChestOpeningFinished
	evMatchLootChestOpeningCancel
	evNotifyCrystalMatchReward
	evCrystalRealmFeedback
	evNewLocationMarker
	evNewTutorialBlocker
	evNewTileSwitch
	evNewInformationProvider
	evNewDynamicGuildLogo
	evNewDecoration
	evTutorialUpdate
	evTriggerHintBox
	evRandomDungeonPositionInfo
	evNewLootChest
	evUpdateLootChest
	evLootChestOpened
	evUpdateLootProtectedByMobsWithMinimapDisplay
	evNewShrine
	evUpdateShrine
	evUpdateRoom
	evNewMistDungeonRoomMobSoul
	evNewHellgateShrine
	evUpdateHellgateShrine
	evActivateHellgateExit
	evMutePlayerUpdate
	evShopTileUpdate
	evShopUpdate
	evAntiCheatKick
	evBattlEyeServerMessage
	evUnlockVanityUnlock
	evAvatarUnlocked
	evCustomizationChanged
	evBaseVaultInfo
	evGuildVaultInfo
	evBankVaultInfo
	evRecoveryVaultPlayerInfo
	evRecoveryVaultGuildInfo
	evUpdateWardrobe
	evCastlePhaseChanged
	evGuildAccountLogEvent
	evNewHideoutObject
	evNewHideoutManagement
	evNewHideoutExit
	evInitHideoutAttackStart
	evInitHideoutAttackCancel
	evInitHideoutAttackFinished
	evHideoutManagementUpdate
	evHideoutUpgradeWithPowerCrystalResult
	evIpChanged
	evSmartClusterQueueUpdateInfo
	evSmartClusterQueueActiveInfo
	evSmartClusterQueueKickWarning
	evSmartClusterQueueInvite
	evReceivedGvgSeasonPoints
	evTowerPowerPointUpdate
	evOpenWorldAttackScheduleStart
	evOpenWorldAttackScheduleFinished
	evOpenWorldAttackScheduleCancel
	evOpenWorldAttackConquerStart
	evOpenWorldAttackConquerFinished
	evOpenWorldAttackConquerCancel
	evOpenWorldAttackConquerStatus
	evOpenWorldAttackStart
	evOpenWorldAttackEnd
	evNewRandomResourceBlocker
	evNewHomeObject
	evHideoutObjectUpdate
	evUpdateInfamy
	evMinimapPositionMarkers
	evNewTunnelExit
	evCorruptedDungeonUpdate
	evCorruptedDungeonStatus
	evCorruptedDungeonInfamy
	evHellgateRestrictedAreaUpdate
	evHellgateInfamy
	evHellgateStatus
	evHellgateStatusUpdate
	evHellgateSuspense
	evReplaceSpellSlotWithMultiSpell
	evNewCorruptedShrine
	evUpdateCorruptedShrine
	evCorruptedShrineUsageStart
	evCorruptedShrineUsageCancel
	evExitUsed
	evLinkedToObject
	evLinkToObjectBroken
	evEstimatedMarketValueUpdate
	evStuckCancel
	evDungonEscapeReady
	evFactionWarfareClusterState
	evFactionWarfareHasUnclaimedWeeklyReportsEvent
	evSimpleFeedback
	evSmartClusterQueueSkipClusterError
	evXignCodeEvent
	evBatchUseItemStart
	evBatchUseItemEnd
	evRedZoneEventClusterStatus
	evRedZonePlayerNotification
	evRedZoneWorldEvent
	evFactionWarfareStats
	evUpdateFactionBalanceFactors
	evFactionEnlistmentChanged
	evUpdateFactionRank
	evFactionWarfareCampaignRewardsUnlocked
	evFeaturedFeatureUpdate
	evNewPowerCrystalObject
	evMinimapCrystalPositionMarker
	evCarryPowerCrystalUpdate
	evPickupPowerCrystalStart
	evPickupPowerCrystalCancel
	evPickupPowerCrystalFinished
	evDoSimpleActionStart
	evDoSimpleActionCancel
	evDoSimpleActionFinished
	evNotifyGuestAccountVerified
	evMightAndFavorReceivedEvent
	evWeeklyPvpChallengeRewardStateUpdate
	evNewUnlockedPvpSeasonChallengeRewards
	evStaticDungeonEntrancesDungeonEventStatusUpdates
	evStaticDungeonDungeonValueUpdate
	evStaticDungeonEntranceDungeonEventsAborted
	evInAppPurchaseConfirmedGooglePlay
	evFeatureSwitchInfo
	evPartyJoinRequestAborted
	evPartyInviteAborted
	evPartyStartHuntRequest
	evPartyStartHuntRequested
	evPartyStartHuntRequestAnswer
	evGuildInviteDeclined
	evCancelMultiSpellSlots
	evNewVisualEventObject
	evCastleClaimProgress
	evCastleClaimProgressLogo
	evTownPortalUpdateState
	evTownPortalFailed
	evConsumableVanityChargesAdded
	evFestivitiesUpdate
	evNewBannerObject
	evNewMistsImmediateReturnExit
	evMistsPlayerJoinedInfo
	evNewMistsStaticEntrance
	evNewMistsOpenWorldExit
	evNewTunnelExitTemp
	evNewMistsWispSpawn
	evMistsWispSpawnStateChange
	evNewMistsCityEntrance
	evNewMistsCityRoadsEntrance
	evMistsCityRoadsEntrancePartyStateUpdate
	evMistsCityRoadsEntranceClearStateForParty
	evMistsEntranceDataChanged
	evNewMistsCagedWisp
	evMistsWispCageOpened
	evMistsEntrancePartyBindingCreated
	evMistsEntrancePartyBindingCleared
	evMistsEntrancePartyBindingInfos
	evNewMistsBorderExit
	evNewMistsDungeonExit
	evLocalQuestInfos
	evLocalQuestStarted
	evLocalQuestActive
	evLocalQuestInactive
	evLocalQuestProgressUpdate
	evNewUnrestrictedPvpZone
	evTemporaryFlaggingStatusUpdate
	evSpellTestPerformanceUpdate
	evTransformation
	evTransformationEnd
	evUpdateTrustlevel
	evRevealHiddenTimeStamps
	evModifyItemTraitFinished
	evRerollItemTraitValueFinished
	evHuntQuestProgressInfo
	evHuntStarted
	evHuntFinished
	evHuntAborted
	evHuntMissionStepStateUpdate
	evNewHuntTrack
	evHuntMissionUpdate
	evHuntQuestMissionProgressUpdate
	evHuntTrackUsed
	evHuntTrackUseableAgain
	evMinimapHuntTrackMarkers
	evNoTracksFound
	evHuntQuestAborted
	evInteractWithTrackStart
	evInteractWithTrackCancel
	evInteractWithTrackFinished
	evNewDynamicCompound
	evLegendaryItemDestroyed
	evAttunementInfo
	evTerritoryClaimRaidedRawEnergyCrystalResult
	evCarryPowerCrystalExpiryWarning
	evPowerCrystalExpired
	evTerritoryRaidStart
	evTerritoryRaidCancel
	evTerritoryRaidFinished
	evTerritoryRaidResult
	evTerritoryMonolithActiveRaidStatus
	evTerritoryMonolithActiveRaidCancelled
	evMonolithEnergyStorageUpdate
	evMonolithNextScheduledOpenWorldAttackUpdate
)