package mesos_metrics

// Counter == int64
// Gauge   == float64

// Metrics from endpoint /{master}:5050/metrics/snapshot
type MasterMetrics struct {
	UptimeSecs                          float64 `json:"master/uptime_secs"`
	Elected                             float64 `json:"master/elected"`

	SlavesConnected                     float64 `json:"master/slaves_connected"`
	SlavesDisconnected                  float64 `json:"master/slaves_disconnected"`
	SlavesActive                        float64 `json:"master/slaves_active"`
	SlavesInactive                      float64 `json:"master/slaves_inactive"`

	FrameworksConnected                 float64 `json:"master/frameworks_connected"`
	FrameworksDisconnected              float64 `json:"master/frameworks_disconnected"`
	FrameworksActive                    float64 `json:"master/frameworks_active"`
	FrameworksInactive                  float64 `json:"master/frameworks_inactive"`

	OutstandingOffers                   float64 `json:"master/outstanding_offers"`

	TasksStaging                        float64 `json:"master/tasks_staging"`
	TasksStarting                       float64 `json:"master/tasks_starting"`
	TasksRunning                        float64 `json:"master/tasks_running"`
	TasksFinished                       int64   `json:"master/tasks_finished"`
	TasksFailed                         int64   `json:"master/tasks_failed"`
	TasksKilled                         int64   `json:"master/tasks_killed"`
	TasksLost                           int64   `json:"master/tasks_lost"`
	TasksError                          int64   `json:"master/tasks_error"`

	// Messages from schedulers.
	MessagesRegisterFramework           int64 `json:"master/messages_register_framework"`
	MessagesReregisterFramework         int64 `json:"master/messages_reregister_framework"`
	MessagesUnregisterFramework         int64 `json:"master/messages_unregister_framework"`
	MessagesDeactivateFramework         int64 `json:"master/messages_deactivate_framework"`
	MessagesKillTask                    int64 `json:"master/messages_kill_task"`
	MessagesStatusUpdateAcknowledgement int64 `json:"master/messages_status_update_acknowledgement"`
	MessagesResourceRequest             int64 `json:"master/messages_resource_request"`
	MessagesLaunchTasks                 int64 `json:"master/messages_launch_tasks"`
	MessagesDeclineOffers               int64 `json:"master/messages_decline_offers"`
	MessagesReviveOffers                int64 `json:"master/messages_revive_offers"`
	MessagesReconcileTasks              int64 `json:"master/messages_reconcile_tasks"`
	MessagesFrameworkToExecutor         int64 `json:"master/messages_framework_to_executor"`

	// Messages from executors.
	MessagesExecutorToFramework         int64 `json:"master/messages_executor_to_framework"`

	// Messages from slaves.
	MessagesRegisterSlave               int64 `json:"master/messages_register_slave"`
	MessagesReregisterSlave             int64 `json:"master/messages_reregister_slave"`
	MessagesUnregisterSlave             int64 `json:"master/messages_unregister_slave"`
	MessagesStatusUpdate                int64 `json:"master/messages_status_update"`
	MessagesExitedExecutor              int64 `json:"master/messages_exited_executor"`
	MessagesUpdateSlave                 int64 `json:"master/messages_update_slave"`

	// Messages from both schedulers and slaves.
	MessagesAuthenticate                int64 `json:"master/messages_authenticate"`

	ValidFrameworkToExecutorMessages    int64 `json:"master/valid_framework_to_executor_messages"`
	InvalidFrameworkToExecutorMessages  int64 `json:"master/invalid_framework_to_executor_messages"`
	ValidExecutorToFrameworkMessages    int64 `json:"master/valid_executor_to_framework_messages"`
	InvalidExecutorToFrameworkMessages  int64 `json:"master/invalid_executor_to_framework_messages"`

	ValidStatusUpdates                  int64 `json:"master/valid_status_updates"`
	InvalidStatusUpdates                int64 `json:"master/invalid_status_updates"`

	ValidStatusUpdateAcknowledgements   int64 `json:"master/valid_status_update_acknowledgements"`
	InvalidStatusUpdateAcknowledgements int64 `json:"master/invalid_status_update_acknowledgements"`

	// Recovery counters.
	RecoverySlaveRemovals               int64 `json:"master/recovery_slave_removals"`

	// Process metrics.
	EventQueueMessages                  float64 `json:"master/event_queue_messages"`
	EventQueueDispatches                float64 `json:"master/event_queue_dispatches"`
	EventQueueHttpRequests              float64 `json:"master/event_queue_http_requests"`

	// Successful registry operations.
	SlaveRegistrations                  int64 `json:"master/slave_registrations"`
	SlaveReregistrations                int64 `json:"master/slave_reregistrations"`
	SlaveRemovals                       int64 `json:"master/slave_removals"`
	SlaveRemovalsReasonUnhealthy        int64 `json:"master/slave_removals_reason_unhealthy"`
	SlaveRemovalsReasonUnregistered     int64 `json:"master/slave_removals_reason_unregistered"`
	SlaveRemovalsReasonRegistered       int64 `json:"master/slave_removals_reason_registered"`

	// Slave observer metrics.
	SlaveShutdownsScheduled             int64 `json:"master/slave_shutdowns_scheduled"`
	SlaveShutdownsCompleted             int64 `json:"master/slave_shutdowns_completed"`
	SlaveShutdownsCanceled              int64 `json:"master/slave_shutdowns_canceled"`
}

// Metrics from endpoint /{slave}:5051/metrics/snapshot
type SlaveMetrics struct {
	TasksStaging  int64 `json:"slave/tasks_staging"`
	TasksStarting int64 `json:"slave/tasks_starting"`
	TasksRunning  int64 `json:"slave/tasks_running"`
	TasksFinished int64   `json:"slave/tasks_finished"`
	TasksFailed   int64   `json:"slave/tasks_failed"`
	TasksKilled   int64   `json:"slave/tasks_killed"`
	TasksLost     int64   `json:"slave/tasks_lost"`
	TasksError    int64   `json:"slave/tasks_error"`
}
