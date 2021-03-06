// Copyright 2015 bol.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	assert.InDelta(0.0546875, stats.CpuPercentage, 0.01)
	assert.EqualValues(96, stats.CpuTotal)
	assert.InDelta(5.25, stats.CpuUsed, 0.01)
	assert.InDelta(0, stats.DiskPercentage, 0.01)
	assert.EqualValues(6038166, stats.DiskTotal)
	assert.EqualValues(0, stats.DiskUsed)
	assert.EqualValues(1, stats.DroppedMessages)

	assert.EqualValues(1, stats.Elected)

	assert.EqualValues(2, stats.EventQueueDispatches)
	assert.EqualValues(0, stats.EventQueueHttpRequests)
	assert.EqualValues(0, stats.EventQueueMessages)

	assert.EqualValues(1, stats.FrameworksActive)
	assert.EqualValues(1, stats.FrameworksConnected)
	assert.EqualValues(0, stats.FrameworksDisconnected)
	assert.EqualValues(0, stats.FrameworksInactive)

	assert.EqualValues(0, stats.InvalidFrameworkToExecutorMessages)
	assert.EqualValues(0, stats.InvalidStatusUpdateAcknowledgements)
	assert.EqualValues(0, stats.InvalidStatusUpdates)

	assert.InDelta(0.0519961396422164, stats.MemoryPercentage, 0.01)
	assert.EqualValues(513942, stats.MemoryTotal)
	assert.EqualValues(26723, stats.MemoryUsed)

	assert.EqualValues(0, stats.MessagesAuthenticate)
	assert.EqualValues(0, stats.MessagesDeactivateFramework)
	assert.EqualValues(162748, stats.MessagesDeclineOffers)
	assert.EqualValues(0, stats.MessagesExitedExecutor)
	assert.EqualValues(0, stats.MessagesFrameworkToExecutor)
	assert.EqualValues(184, stats.MessagesKillTask)
	assert.EqualValues(403, stats.MessagesLaunchTasks)
	assert.EqualValues(3251, stats.MessagesReconcileTasks)
	assert.EqualValues(1, stats.MessagesRegisterFramework)
	assert.EqualValues(0, stats.MessagesRegisterSlave)
	assert.EqualValues(0, stats.MessagesReregisterFramework)
	assert.EqualValues(2, stats.MessagesReregisterSlave)
	assert.EqualValues(0, stats.MessagesResourceRequest)
	assert.EqualValues(0, stats.MessagesReviveOffers)
	assert.EqualValues(714, stats.MessagesStatusUpdate)
	assert.EqualValues(714, stats.MessagesStatusUpdateAcknowledgement)
	assert.EqualValues(0, stats.MessagesUnregisterFramework)
	assert.EqualValues(0, stats.MessagesUnregisterSlave)

	assert.EqualValues(0, stats.OutstandingOffers)
	assert.EqualValues(0, stats.RecoverySlaveRemovals)

	assert.EqualValues(0, stats.SlaveRegistrations)
	assert.EqualValues(0, stats.SlaveRemovals)
	assert.EqualValues(2, stats.SlaveReregistrations)
	assert.EqualValues(0, stats.SlaveShutdownsCanceled)
	assert.EqualValues(0, stats.SlaveShutdownsScheduled)
	assert.EqualValues(2, stats.SlavesActive)
	assert.EqualValues(2, stats.SlavesConnected)
	assert.EqualValues(0, stats.SlavesDisconnected)
	assert.EqualValues(0, stats.SlavesInactive)

	assert.EqualValues(57, stats.TaskFailedSourceSlaveReasonCommandExecutorFailed)

	assert.EqualValues(0, stats.TasksError)
	assert.EqualValues(139, stats.TasksFailed)
	assert.EqualValues(45, stats.TasksFinished)
	assert.EqualValues(184, stats.TasksKilled)
	assert.EqualValues(0, stats.TasksLost)
	assert.EqualValues(35, stats.TasksRunning)
	assert.EqualValues(0, stats.TasksStaging)
	assert.EqualValues(0, stats.TasksStarting)

	assert.InDelta(488714.649007872, stats.UptimeSecs, 0.01)

	assert.EqualValues(0, stats.ValidFrameworkToExecutorMessages)
	assert.EqualValues(714, stats.ValidStatusUpdateAcknowledgements)
	assert.EqualValues(714, stats.ValidStatusUpdates)

	assert.EqualValues(0, stats.RegistrarQueuedOperations)
	assert.EqualValues(451, stats.RegistrarSizeBytes)
	assert.InDelta(6.274048, stats.RegistrarStateFetchMilliSecs, 0.01)
	assert.InDelta(3.23968, stats.RegistrarStateStoreMilliSecs, 0.01)
	assert.EqualValues(3, stats.RegistrarStateStoreMilliSecsCount)
	assert.InDelta(9.825792, stats.RegistrarStateStoreMilliSecsMax, 0.01)
	assert.InDelta(2.58176, stats.RegistrarStateStoreMilliSecsMin, 0.01)
	assert.InDelta(3.23968, stats.RegistrarStateStoreMilliSecsP50, 0.01)
	assert.InDelta(8.5085696, stats.RegistrarStateStoreMilliSecsP90, 0.01)
	assert.InDelta(9.1671808, stats.RegistrarStateStoreMilliSecsP95, 0.01)
	assert.InDelta(9.69406976, stats.RegistrarStateStoreMilliSecsP99, 0.01)
	assert.InDelta(9.812619776, stats.RegistrarStateStoreMilliSecsP999, 0.01)
	assert.InDelta(9.8244747776, stats.RegistrarStateStoreMilliSecsP9999, 0.01)

	assert.EqualValues(48, stats.SystemCpuTotal)
	assert.InDelta(1.2, stats.SystemLoad15MinuteAvg, 0.01)
	assert.InDelta(1.67, stats.SystemLoad1MinuteAvg, 0.01)
	assert.InDelta(1.53, stats.SystemLoad5MinuteAvg, 0.01)
	assert.EqualValues(115047735296, stats.SystemMemoryFreeBytes)
	assert.EqualValues(270527713280, stats.SystemMemoryTotalBytes)
}

func TestSlaveMetricsJsonMapping(t *testing.T) {
	assert := assert.New(t)
	var stats SlaveMetrics
	if err := json.NewDecoder(strings.NewReader(slaveMetricsSnapshotJson)).Decode(&stats); err != nil {
		t.Errorf("failed to deserialize response: ", err)
	}
	t.Log(stats)

	assert.InDelta(0.0729166666666667, stats.CpuPercentage, 0.01)
	assert.EqualValues(48, stats.CpuTotal)
	assert.InDelta(3.5, stats.CpuUsed, 0.01)
	assert.InDelta(0, stats.DiskPercentage, 0.01)
	assert.EqualValues(3019083, stats.DiskTotal)
	assert.EqualValues(0, stats.DiskUsed)

	assert.EqualValues(0, stats.ExecutorsRegistering)
	assert.EqualValues(14, stats.ExecutorsRunning)
	assert.EqualValues(203, stats.ExecutorsTerminated)
	assert.EqualValues(0, stats.ExecutorsTerminating)

	assert.EqualValues(1, stats.FrameworksActive)
	assert.EqualValues(0, stats.InvalidFrameworkMessages)
	assert.EqualValues(0, stats.InvalidStatusUpdates)

	assert.InDelta(0.045795050803398, stats.MemoryPercentage, 0.01)
	assert.EqualValues(256971, stats.MemoryTotal)
	assert.EqualValues(11768, stats.MemoryUsed)

	assert.EqualValues(0, stats.RecoveryErrors)
	assert.EqualValues(1, stats.Registered)

	assert.EqualValues(76, stats.TasksFailed)
	assert.EqualValues(29, stats.TasksFinished)
	assert.EqualValues(98, stats.TasksKilled)
	assert.EqualValues(0, stats.TasksLost)
	assert.EqualValues(14, stats.TasksRunning)
	assert.EqualValues(0, stats.TasksStaging)
	assert.EqualValues(0, stats.TasksStarting)

	assert.InDelta(488886.147593984, stats.UptimeSecs, 0.01)

	assert.EqualValues(0, stats.ValidFrameworkMessages)
	assert.EqualValues(391, stats.ValidStatusUpdates)

	assert.EqualValues(48, stats.SystemCpuTotal)
	assert.InDelta(0.22, stats.SystemLoad15MinuteAvg, 0.01)
	assert.InDelta(0.18, stats.SystemLoad1MinuteAvg, 0.01)
	assert.InDelta(0.2, stats.SystemLoad5MinuteAvg, 0.01)
	assert.EqualValues(253197996032, stats.SystemMemoryFreeBytes)
	assert.EqualValues(270527713280, stats.SystemMemoryTotalBytes)
}
