package block

import (
	"fmt"
)

type StorageOverlay struct {
	overlay *Overlay
	cache   *MmapCache
	storage ReadonlyDevice
}

func NewStorageOverlay(device ReadonlyDevice, blockSize int64, cachePath string) (*StorageOverlay, error) {
	size, err := device.Size()
	if err != nil {
		return nil, fmt.Errorf("error getting device size: %w", err)
	}

	cache, err := NewMmapCache(size, blockSize, cachePath)
	if err != nil {
		return nil, fmt.Errorf("error creating cache: %w", err)
	}

	overlay := newOverlay(device, cache)

	return &StorageOverlay{
		overlay: overlay,
		cache:   cache,
		storage: device,
	}, nil
}

func (o *StorageOverlay) ReadAt(p []byte, off int64) (n int, err error) {
	return o.overlay.ReadAt(p, off)
}

func (o *StorageOverlay) WriteAt(p []byte, off int64) (n int, err error) {
	return o.overlay.WriteAt(p, off)
}

func (o *StorageOverlay) Size() (int64, error) {
	return o.storage.Size()
}

func (o *StorageOverlay) Close() error {
	return o.cache.Close()
}

func (o *StorageOverlay) Slice(offset, length int64) ([]byte, error) {
	return o.cache.Slice(offset, length)
}
