package fileapi

import "syscall"

func getFixedDriveMounts() ([]fixedVolumeMounts, error) {
	var out []fixedVolumeMounts

	err := enumVolumes(func(guid []uint16) error {
		mounts, err := maybeGetFixedVolumeMounts(guid)
		if err != nil {
			return err
		}
		if len(mounts) > 0 {
			out = append(out, fixedVolumeMounts{
				volume: syscall.UTF16ToString(guid),
				mounts: LPSTRsToStrings(mounts),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return out, nil
}
