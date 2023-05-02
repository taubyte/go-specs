package car

import "path"

func (d Dir) Uncompressed() string {
	return path.Join(d.String(), UncompressedCarFile)
}

func (d Dir) CompressedWebsite() string {
	return path.Join(d.String(), CompressedWebsiteCar)
}

func (d Dir) CompressedWasm() string {
	return path.Join(d.String(), CompressedArtifactCar)
}

func UncompressedOutput(outDir string) string {
	return path.Join(outDir, UncompressedCarFile)
}
