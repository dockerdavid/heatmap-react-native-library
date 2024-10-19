package heatmapRepositoryAdapter

import (
	"encoding/json"
	heatmapDomain "heatmap_mobile/domain"
)

func (r Repository) SaveDevice(device string) error {
	stmt, err := r.DB.DB.Prepare(`
		INSERT INTO device 
			(id, name, brand, model, manufacturer, screen_width, screen_height, build_id, battery_level, build_number, bundle_id, free_disk_storage, power_state, readable_version, system_name, system_version) 
		VALUES 
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			name = excluded.name,
			brand = excluded.brand,
			model = excluded.model,
			manufacturer = excluded.manufacturer,
			screen_width = excluded.screen_width,
			screen_height = excluded.screen_height,
			build_id = excluded.build_id,
			battery_level = excluded.battery_level,
			build_number = excluded.build_number,
			bundle_id = excluded.bundle_id,
			free_disk_storage = excluded.free_disk_storage,
			power_state = excluded.power_state,
			readable_version = excluded.readable_version,
			system_name = excluded.system_name,
			system_version = excluded.system_version`)

	if err != nil {
		return err
	}

	var d heatmapDomain.Device

	err = json.Unmarshal([]byte(device), &d)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(d.ID, d.Name, d.Brand, d.Model, d.Manufacturer, d.ScreenWidth, d.ScreenHeight, d.BuildID, d.BatteryLevel, d.BuildNumber, d.BundleID, d.FreeDiskStorage, d.PowerState, d.ReadableVersion, d.SystemName, d.SystemVersion)

	if err != nil {
		return err
	}

	return nil
}
