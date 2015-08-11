package mesos_metrics

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var masterMetricsSnapshotJson string = `
{
  "master/cpus_percent": 0.0546875,
  "master/cpus_total": 96,
  "master/cpus_used": 5.25,
  "master/disk_percent": 0,
  "master/disk_total": 6038166,
  "master/disk_used": 0,
  "master/dropped_messages": 1,
  "master/elected": 1,
  "master/event_queue_dispatches": 2,
  "master/event_queue_http_requests": 0,
  "master/event_queue_messages": 0,
  "master/frameworks_active": 1,
  "master/frameworks_connected": 1,
  "master/frameworks_disconnected": 0,
  "master/frameworks_inactive": 0,
  "master/invalid_framework_to_executor_messages": 0,
  "master/invalid_status_update_acknowledgements": 0,
  "master/invalid_status_updates": 0,
  "master/mem_percent": 0.0519961396422164,
  "master/mem_total": 513942,
  "master/mem_used": 26723,
  "master/messages_authenticate": 0,
  "master/messages_deactivate_framework": 0,
  "master/messages_decline_offers": 162748,
  "master/messages_exited_executor": 0,
  "master/messages_framework_to_executor": 0,
  "master/messages_kill_task": 184,
  "master/messages_launch_tasks": 403,
  "master/messages_reconcile_tasks": 3251,
  "master/messages_register_framework": 1,
  "master/messages_register_slave": 0,
  "master/messages_reregister_framework": 0,
  "master/messages_reregister_slave": 2,
  "master/messages_resource_request": 0,
  "master/messages_revive_offers": 0,
  "master/messages_status_update": 714,
  "master/messages_status_update_acknowledgement": 714,
  "master/messages_unregister_framework": 0,
  "master/messages_unregister_slave": 0,
  "master/outstanding_offers": 0,
  "master/recovery_slave_removals": 0,
  "master/slave_registrations": 0,
  "master/slave_removals": 0,
  "master/slave_reregistrations": 2,
  "master/slave_shutdowns_canceled": 0,
  "master/slave_shutdowns_scheduled": 0,
  "master/slaves_active": 2,
  "master/slaves_connected": 2,
  "master/slaves_disconnected": 0,
  "master/slaves_inactive": 0,
  "master/task_failed/source_slave/reason_command_executor_failed": 57,
  "master/tasks_error": 0,
  "master/tasks_failed": 139,
  "master/tasks_finished": 45,
  "master/tasks_killed": 184,
  "master/tasks_lost": 0,
  "master/tasks_running": 35,
  "master/tasks_staging": 0,
  "master/tasks_starting": 0,
  "master/uptime_secs": 488714.649007872,
  "master/valid_framework_to_executor_messages": 0,
  "master/valid_status_update_acknowledgements": 714,
  "master/valid_status_updates": 714,
  "registrar/queued_operations": 0,
  "registrar/registry_size_bytes": 451,
  "registrar/state_fetch_ms": 6.274048,
  "registrar/state_store_ms": 3.23968,
  "registrar/state_store_ms/count": 3,
  "registrar/state_store_ms/max": 9.825792,
  "registrar/state_store_ms/min": 2.58176,
  "registrar/state_store_ms/p50": 3.23968,
  "registrar/state_store_ms/p90": 8.5085696,
  "registrar/state_store_ms/p95": 9.1671808,
  "registrar/state_store_ms/p99": 9.69406976,
  "registrar/state_store_ms/p999": 9.812619776,
  "registrar/state_store_ms/p9999": 9.8244747776,
  "system/cpus_total": 48,
  "system/load_15min": 1.2,
  "system/load_1min": 1.67,
  "system/load_5min": 1.53,
  "system/mem_free_bytes": 115047735296,
  "system/mem_total_bytes": 270527713280
}
`

var slaveMetricsSnapshotJson = `
{
  "slave/cpus_percent": 0.0729166666666667,
  "slave/cpus_total": 48,
  "slave/cpus_used": 3.5,
  "slave/disk_percent": 0,
  "slave/disk_total": 3019083,
  "slave/disk_used": 0,
  "slave/executors_registering": 0,
  "slave/executors_running": 14,
  "slave/executors_terminated": 203,
  "slave/executors_terminating": 0,
  "slave/frameworks_active": 1,
  "slave/invalid_framework_messages": 0,
  "slave/invalid_status_updates": 0,
  "slave/mem_percent": 0.045795050803398,
  "slave/mem_total": 256971,
  "slave/mem_used": 11768,
  "slave/recovery_errors": 0,
  "slave/registered": 1,
  "slave/tasks_failed": 76,
  "slave/tasks_finished": 29,
  "slave/tasks_killed": 98,
  "slave/tasks_lost": 0,
  "slave/tasks_running": 14,
  "slave/tasks_staging": 0,
  "slave/tasks_starting": 0,
  "slave/uptime_secs": 488886.147593984,
  "slave/valid_framework_messages": 0,
  "slave/valid_status_updates": 391,
  "system/cpus_total": 48,
  "system/load_15min": 0.22,
  "system/load_1min": 0.18,
  "system/load_5min": 0.2,
  "system/mem_free_bytes": 253197996032,
  "system/mem_total_bytes": 270527713280
}`

func TestMasterMetricsJsonMapping(t *testing.T) {
	assert := assert.New(t)
	var stats MasterMetrics
	if err := json.NewDecoder(strings.NewReader(masterMetricsSnapshotJson)).Decode(&stats); err != nil {
		t.Errorf("failed to deserialize response: ", err)
	}
	t.Log(stats)
	assert.InDelta(488714.649007872, stats.UptimeSecs, 0.01)
	assert.InDelta(1, stats.Elected, 0.01)

	assert.InDelta(2, stats.SlavesConnected, 0.01)
	assert.InDelta(0, stats.SlavesDisconnected, 0.01)
	assert.InDelta(2, stats.SlavesActive, 0.01)
	assert.InDelta(0, stats.SlavesInactive, 0.01)

	assert.InDelta(1, stats.FrameworksConnected, 0.01)
	assert.InDelta(0, stats.FrameworksDisconnected, 0.01)
	assert.InDelta(1, stats.FrameworksActive, 0.01)
	assert.InDelta(0, stats.FrameworksInactive, 0.01)

	assert.InDelta(0, stats.OutstandingOffers, 0.01)

	assert.InDelta(0, stats.TasksStaging, 0.01)
	assert.InDelta(0, stats.TasksStarting, 0.01)
	assert.InDelta(35, stats.TasksRunning, 0.01)

	assert.EqualValues(45, stats.TasksFinished)
	assert.EqualValues(139, stats.TasksFailed)
	assert.EqualValues(184, stats.TasksKilled)
	assert.EqualValues(0, stats.TasksLost)
	assert.EqualValues(0, stats.TasksError)

	assert.EqualValues(1, stats.MessagesRegisterFramework)
	assert.EqualValues(0, stats.MessagesReregisterFramework)
	assert.EqualValues(0, stats.MessagesUnregisterFramework)
	assert.EqualValues(0, stats.MessagesDeactivateFramework)
	assert.EqualValues(184, stats.MessagesKillTask)
	assert.EqualValues(714, stats.MessagesStatusUpdateAcknowledgement)
	assert.EqualValues(0, stats.MessagesResourceRequest)
	assert.EqualValues(403, stats.MessagesLaunchTasks)
	assert.EqualValues(162748, stats.MessagesDeclineOffers)
	assert.EqualValues(0, stats.MessagesReviveOffers)
	assert.EqualValues(3251, stats.MessagesReconcileTasks)
	assert.EqualValues(0, stats.MessagesFrameworkToExecutor)

	assert.EqualValues(0, stats.MessagesExecutorToFramework)

	assert.EqualValues(0, stats.MessagesRegisterSlave)
	assert.EqualValues(2, stats.MessagesReregisterSlave)
	assert.EqualValues(0, stats.MessagesUnregisterSlave)
	assert.EqualValues(714, stats.MessagesStatusUpdate)
	assert.EqualValues(0, stats.MessagesExitedExecutor)
	assert.EqualValues(0, stats.MessagesUpdateSlave)

	assert.EqualValues(0, stats.MessagesAuthenticate)

	assert.EqualValues(0, stats.ValidFrameworkToExecutorMessages)
	assert.EqualValues(0, stats.InvalidFrameworkToExecutorMessages)
	assert.EqualValues(0, stats.ValidExecutorToFrameworkMessages)
	assert.EqualValues(0, stats.InvalidExecutorToFrameworkMessages)

	assert.EqualValues(714, stats.ValidStatusUpdates)
	assert.EqualValues(0, stats.InvalidStatusUpdates)

	assert.EqualValues(714, stats.ValidStatusUpdateAcknowledgements)
	assert.EqualValues(0, stats.InvalidStatusUpdateAcknowledgements)

	assert.EqualValues(0, stats.RecoverySlaveRemovals)

	assert.InDelta(0, stats.EventQueueMessages, 0.01)
	assert.InDelta(2, stats.EventQueueDispatches, 0.01)
	assert.InDelta(0, stats.EventQueueHttpRequests, 0.01)

	assert.EqualValues(0, stats.SlaveRegistrations)
	assert.EqualValues(2, stats.SlaveReregistrations)
	assert.EqualValues(0, stats.SlaveRemovals)
	assert.EqualValues(0, stats.SlaveRemovalsReasonUnhealthy)
	assert.EqualValues(0, stats.SlaveRemovalsReasonUnregistered)
	assert.EqualValues(0, stats.SlaveRemovalsReasonRegistered)

	assert.EqualValues(0, stats.SlaveShutdownsScheduled)
	assert.EqualValues(0, stats.SlaveShutdownsCompleted)
	assert.EqualValues(0, stats.SlaveShutdownsCanceled)
}

func TestSlaveMetricsJsonMapping(t *testing.T) {
	assert := assert.New(t)
	var stats SlaveMetrics
	if err := json.NewDecoder(strings.NewReader(slaveMetricsSnapshotJson)).Decode(&stats); err != nil {
		t.Errorf("failed to deserialize response: ", err)
	}
	t.Log(stats)
	assert.EqualValues(0, stats.TasksStaging)
	assert.EqualValues(0, stats.TasksStarting)
	assert.EqualValues(14, stats.TasksRunning)

	assert.EqualValues(29, stats.TasksFinished)
	assert.EqualValues(76, stats.TasksFailed)
	assert.EqualValues(98, stats.TasksKilled)
	assert.EqualValues(0, stats.TasksLost)
}
