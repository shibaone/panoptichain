// Code generated by "stringer -type=ObservableTopic"; DO NOT EDIT.

package topics

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NewEVMBlock-0]
	_ = x[BorStateSync-1]
	_ = x[BlockInterval-2]
	_ = x[CheckpointSignatures-3]
	_ = x[ValidatorWallet-4]
	_ = x[HeimdallBlockInterval-5]
	_ = x[NewHeimdallBlock-6]
	_ = x[Milestone-7]
	_ = x[Reorg-8]
	_ = x[SensorBlocks-9]
	_ = x[SensorBlockEvents-10]
	_ = x[BorMissedBlockProposal-11]
	_ = x[HeimdallMissedBlockProposal-12]
	_ = x[Checkpoint-13]
	_ = x[MissedCheckpointProposal-14]
	_ = x[MissedMilestoneProposal-15]
	_ = x[TransactionPool-16]
	_ = x[StolenBlock-17]
	_ = x[HashDivergence-18]
	_ = x[System-19]
	_ = x[RefreshStateTime-20]
	_ = x[ZkEVMBatches-21]
	_ = x[ExitRoots-22]
	_ = x[BridgeEvent-23]
	_ = x[ClaimEvent-24]
	_ = x[DepositCounts-25]
	_ = x[BridgeEventTimes-26]
	_ = x[ClaimEventTimes-27]
	_ = x[RollupManager-28]
	_ = x[Span-29]
	_ = x[TimeToMine-30]
	_ = x[AccountBalances-31]
	_ = x[TrustedBatch-32]
	_ = x[ExchangeRate-33]
	_ = x[TimeToFinalized-34]
}

const _ObservableTopic_name = "NewEVMBlockBorStateSyncBlockIntervalCheckpointSignaturesValidatorWalletHeimdallBlockIntervalNewHeimdallBlockMilestoneReorgSensorBlocksSensorBlockEventsBorMissedBlockProposalHeimdallMissedBlockProposalCheckpointMissedCheckpointProposalMissedMilestoneProposalTransactionPoolStolenBlockHashDivergenceSystemRefreshStateTimeZkEVMBatchesExitRootsBridgeEventClaimEventDepositCountsBridgeEventTimesClaimEventTimesRollupManagerSpanTimeToMineAccountBalancesTrustedBatchExchangeRateTimeToFinalized"

var _ObservableTopic_index = [...]uint16{0, 11, 23, 36, 56, 71, 92, 108, 117, 122, 134, 151, 173, 200, 210, 234, 257, 272, 283, 297, 303, 319, 331, 340, 351, 361, 374, 390, 405, 418, 422, 432, 447, 459, 471, 486}

func (i ObservableTopic) String() string {
	if i < 0 || i >= ObservableTopic(len(_ObservableTopic_index)-1) {
		return "ObservableTopic(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ObservableTopic_name[_ObservableTopic_index[i]:_ObservableTopic_index[i+1]]
}