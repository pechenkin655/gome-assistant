package gomeassistant

import "time"

type EventZWaveJSValueNotification struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Event struct {
		EventType string `json:"event_type"`
		Data      struct {
			Domain           string `json:"domain"`
			NodeID           int    `json:"node_id"`
			HomeID           int64  `json:"home_id"`
			Endpoint         int    `json:"endpoint"`
			DeviceID         string `json:"device_id"`
			CommandClass     int    `json:"command_class"`
			CommandClassName string `json:"command_class_name"`
			Label            string `json:"label"`
			Property         string `json:"property"`
			PropertyName     string `json:"property_name"`
			PropertyKey      string `json:"property_key"`
			PropertyKeyName  string `json:"property_key_name"`
			Value            string `json:"value"`
			ValueRaw         int    `json:"value_raw"`
		} `json:"data"`
		Origin    string    `json:"origin"`
		TimeFired time.Time `json:"time_fired"`
	} `json:"event"`
}

type EventTimerEvent struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Event struct {
		EventType string `json:"event_type"`
		Data      struct {
			EntityID   string    `json:"entity_id"`
			FinishedAt time.Time `json:"finished_at"`
		} `json:"data"`
		Origin    string `json:"origin"`
		TimeFired time.Time `json:"time_fired"`
		Context   struct {
			ID       string  `json:"id"`
			ParentID *string `json:"parent_id"` // nullable
			UserID   *string `json:"user_id"`   // nullable
		} `json:"context"`
	} `json:"event"`
}
