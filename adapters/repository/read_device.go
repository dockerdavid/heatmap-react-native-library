package heatmapRepositoryAdapter

import (
	"encoding/json"
	heatmapDomain "heatmap_mobile/domain"
)

func (r Repository) ReadDevice() (string, error) {
	const query = `SELECT id, name, brand, model, manufacturer, screen_width, screen_height, build_id, battery_level, build_number, bundle_id, free_disk_storage, power_state, readable_version, system_name, system_version FROM device`

	prepare, err := r.DB.DB.Prepare(query)

	if err != nil {
		return "", err
	}

	defer prepare.Close()

	var device heatmapDomain.Device

	if err = prepare.QueryRow().Scan(&device.ID, &device.Name, &device.Brand, &device.Model, &device.Manufacturer, &device.ScreenWidth, &device.ScreenHeight, &device.BuildID, &device.BatteryLevel, &device.BuildNumber, &device.BundleID, &device.FreeDiskStorage, &device.PowerState, &device.ReadableVersion, &device.SystemName, &device.SystemVersion); err != nil {
		return "", err
	}

	marshal, err := json.Marshal(device)

	return string(marshal), nil
}
