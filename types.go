package mesos_metrics

// Counter == int64
// Gauge   == float64

// Metrics from endpoint /{master}:5050/metrics/snapshot
// Mesos version 0.22.1
type MasterMetrics struct {
	CpuPercentage                                    float64 `json:"master/cpus_percent"`
	CpuTotal                                         int64 `json:"master/cpus_total"`
	CpuUsed                                          float64 `json:"master/cpus_used"`
	DiskPercentage                                   float64 `json:"master/disk_percent"`
	DiskTotal                                        int64 `json:"master/disk_total"`
	DiskUsed                                         int64 `json:"master/disk_used"`
	DroppedMessages                                  int64 `json:"master/dropped_messages"`

	Elected                                          float64 `json:"master/elected"`

	EventQueueDispatches                             float64 `json:"master/event_queue_dispatches"`
	EventQueueHttpRequests                           float64 `json:"master/event_queue_http_requests"`
	EventQueueMessages                               float64 `json:"master/event_queue_messages"`

	FrameworksActive                                 float64 `json:"master/frameworks_active"`
	FrameworksConnected                              float64 `json:"master/frameworks_connected"`
	FrameworksDisconnected                           float64 `json:"master/frameworks_disconnected"`
	FrameworksInactive                               float64 `json:"master/frameworks_inactive"`

	InvalidFrameworkToExecutorMessages               int64 `json:"master/invalid_framework_to_executor_messages"`
	InvalidStatusUpdateAcknowledgements              int64 `json:"master/invalid_status_update_acknowledgements"`
	InvalidStatusUpdates                             int64 `json:"master/invalid_status_updates"`

	MemoryPercentage                                 float64 `json:"master/mem_percent"`
	MemoryTotal                                      int64 `json:"master/mem_total"`
	MemoryUsed                                       int64 `json:"master/mem_used"`

	MessagesAuthenticate                             int64 `json:"master/messages_authenticate"`
	MessagesDeactivateFramework                      int64 `json:"master/messages_deactivate_framework"`
	MessagesDeclineOffers                            int64 `json:"master/messages_decline_offers"`
	MessagesExitedExecutor                           int64 `json:"master/messages_exited_executor"`
	MessagesFrameworkToExecutor                      int64 `json:"master/messages_framework_to_executor"`
	MessagesKillTask                                 int64 `json:"master/messages_kill_task"`
	MessagesLaunchTasks                              int64 `json:"master/messages_launch_tasks"`
	MessagesReconcileTasks                           int64 `json:"master/messages_reconcile_tasks"`
	MessagesRegisterFramework                        int64 `json:"master/messages_register_framework"`
	MessagesRegisterSlave                            int64 `json:"master/messages_register_slave"`
	MessagesReregisterFramework                      int64 `json:"master/messages_reregister_framework"`
	MessagesReregisterSlave                          int64 `json:"master/messages_reregister_slave"`
	MessagesResourceRequest                          int64 `json:"master/messages_resource_request"`
	MessagesReviveOffers                             int64 `json:"master/messages_revive_offers"`
	MessagesStatusUpdate                             int64 `json:"master/messages_status_update"`
	MessagesStatusUpdateAcknowledgement              int64 `json:"master/messages_status_update_acknowledgement"`
	MessagesUnregisterFramework                      int64 `json:"master/messages_unregister_framework"`
	MessagesUnregisterSlave                          int64 `json:"master/messages_unregister_slave"`

	OutstandingOffers                                float64 `json:"master/outstanding_offers"`
	RecoverySlaveRemovals                            int64 `json:"master/recovery_slave_removals"`

	SlaveRegistrations                               int64 `json:"master/slave_registrations"`
	SlaveRemovals                                    int64 `json:"master/slave_removals"`
	SlaveReregistrations                             int64 `json:"master/slave_reregistrations"`
	SlaveShutdownsCanceled                           int64 `json:"master/slave_shutdowns_canceled"`
	SlaveShutdownsScheduled                          int64 `json:"master/slave_shutdowns_scheduled"`
	SlavesActive                                     float64 `json:"master/slaves_active"`
	SlavesConnected                                  float64 `json:"master/slaves_connected"`
	SlavesDisconnected                               float64 `json:"master/slaves_disconnected"`
	SlavesInactive                                   float64 `json:"master/slaves_inactive"`

	TaskFailedSourceSlaveReasonCommandExecutorFailed int64 `json:"master/task_failed/source_slave/reason_command_executor_failed"`

	TasksError                                       int64   `json:"master/tasks_error"`
	TasksFailed                                      int64   `json:"master/tasks_failed"`
	TasksFinished                                    int64   `json:"master/tasks_finished"`
	TasksKilled                                      int64   `json:"master/tasks_killed"`
	TasksLost                                        int64   `json:"master/tasks_lost"`
	TasksRunning                                     float64 `json:"master/tasks_running"`
	TasksStaging                                     float64 `json:"master/tasks_staging"`
	TasksStarting                                    float64 `json:"master/tasks_starting"`

	UptimeSecs                                       float64 `json:"master/uptime_secs"`

	ValidFrameworkToExecutorMessages                 int64 `json:"master/valid_framework_to_executor_messages"`
	ValidStatusUpdateAcknowledgements                int64 `json:"master/valid_status_update_acknowledgements"`
	ValidStatusUpdates                               int64 `json:"master/valid_status_updates"`

	RegistrarQueuedOperations                        int64 `json:"registrar/queued_operations"`
	RegistrarSizeBytes                               int64 `json:"registrar/registry_size_bytes"`
	RegistrarStateFetchMilliSecs                     float64 `json:"registrar/state_fetch_ms"`
	RegistrarStateStoreMilliSecs                     float64 `json:"registrar/state_store_ms"`
	RegistrarStateStoreMilliSecsCount                int64 `json:"registrar/state_store_ms/count"`
	RegistrarStateStoreMilliSecsMax                  float64 `json:"registrar/state_store_ms/max"`
	RegistrarStateStoreMilliSecsMin                  float64 `json:"registrar/state_store_ms/min"`
	RegistrarStateStoreMilliSecsP50                  float64 `json:"registrar/state_store_ms/p50"`
	RegistrarStateStoreMilliSecsP90                  float64 `json:"registrar/state_store_ms/p90"`
	RegistrarStateStoreMilliSecsP95                  float64 `json:"registrar/state_store_ms/p95"`
	RegistrarStateStoreMilliSecsP99                  float64 `json:"registrar/state_store_ms/p99"`
	RegistrarStateStoreMilliSecsP999                 float64 `json:"registrar/state_store_ms/p999"`
	RegistrarStateStoreMilliSecsP9999                float64 `json:"registrar/state_store_ms/p9999"`

	SystemCpuTotal                                   int64 `json:"system/cpus_total"`
	SystemLoad15MinuteAvg                            float64 `json:"system/load_15min"`
	SystemLoad1MinuteAvg                             float64 `json:"system/load_1min"`
	SystemLoad5MinuteAvg                             float64 `json:"system/load_5min"`
	SystemMemoryFreeBytes                            int64 `json:"system/mem_free_bytes"`
	SystemMemoryTotalBytes                           int64 `json:"system/mem_total_bytes"`
}

// Metrics from endpoint /{slave}:5051/metrics/snapshot
// Mesos version 0.22.1
type SlaveMetrics struct {
	CpuPercentage            float64 `json:"slave/cpus_percent"`
	CpuTotal                 int64 `json:"slave/cpus_total"`
	CpuUsed                  float64 `json:"slave/cpus_used"`
	DiskPercentage           float64 `json:"slave/disk_percent"`
	DiskTotal                int64 `json:"slave/disk_total"`
	DiskUsed                 int64 `json:"slave/disk_used"`

	ExecutorsRegistering     int64 `json:"slave/executors_registering"`
	ExecutorsRunning         int64 `json:"slave/executors_running"`
	ExecutorsTerminated      int64 `json:"slave/executors_terminated"`
	ExecutorsTerminating     int64 `json:"slave/executors_terminating"`

	FrameworksActive         int64 `json:"slave/frameworks_active"`
	InvalidFrameworkMessages int64 `json:"slave/invalid_framework_messages"`
	InvalidStatusUpdates     int64 `json:"slave/invalid_status_updates"`

	MemoryPercentage         float64 `json:"slave/mem_percent"`
	MemoryTotal              int64 `json:"slave/mem_total"`
	MemoryUsed               int64 `json:"slave/mem_used"`

	RecoveryErrors           int64 `json:"slave/recovery_errors"`
	Registered               int64 `json:"slave/registered"`

	TasksFailed              int64 `json:"slave/tasks_failed"`
	TasksFinished            int64 `json:"slave/tasks_finished"`
	TasksKilled              int64 `json:"slave/tasks_killed"`
	TasksLost                int64 `json:"slave/tasks_lost"`
	TasksRunning             int64 `json:"slave/tasks_running"`
	TasksStaging             int64 `json:"slave/tasks_staging"`
	TasksStarting            int64 `json:"slave/tasks_starting"`

	UptimeSecs               float64 `json:"slave/uptime_secs"`

	ValidFrameworkMessages   int64 `json:"slave/valid_framework_messages"`
	ValidStatusUpdates       int64 `json:"slave/valid_status_updates"`

	SystemCpuTotal           int64 `json:"system/cpus_total"`
	SystemLoad15MinuteAvg    float64 `json:"system/load_15min"`
	SystemLoad1MinuteAvg     float64 `json:"system/load_1min"`
	SystemLoad5MinuteAvg     float64 `json:"system/load_5min"`
	SystemMemoryFreeBytes    int64 `json:"system/mem_free_bytes"`
	SystemMemoryTotalBytes   int64 `json:"system/mem_total_bytes"`
}
