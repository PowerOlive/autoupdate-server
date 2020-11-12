package server

import (
	"fmt"
	"os"
	"path"
	"testing"
)

var (
	ghAccountOwner      = "getlantern"
	ghAccountRepository = "lantern"
)

func init() {
	os.Setenv(envSkipDownload, "true")
	if v := os.Getenv("GH_ACCOUNT_OWNER"); v != "" {
		ghAccountOwner = v
	}
	if v := os.Getenv("GH_ACCOUNT_REPOSITORY"); v != "" {
		ghAccountRepository = v
	}
}

func TestSplitUpdateAsset(t *testing.T) {
	var err error
	var info *AssetInfo

	if info, err = getAssetInfo("update_darwin_386.dmg"); err != nil {
		t.Fatalf("Failed to get asset info: %v", err)
	}
	if info.OS != OS.Darwin || info.Arch != Arch.X86 {
		t.Fatal("Failed to identify update asset.")
	}

	if info, err = getAssetInfo("update_darwin_amd64.v1"); err != nil {
		t.Fatalf("Failed to get asset info: %v", err)
	}
	if info.OS != OS.Darwin || info.Arch != Arch.X64 {
		t.Fatal("Failed to identify update asset.")
	}

	if info, err = getAssetInfo("update_linux_arm"); err != nil {
		t.Fatalf("Failed to get asset info: %v", err)
	}
	if info.OS != OS.Linux || info.Arch != Arch.ARM {
		t.Fatal("Failed to identify update asset.")
	}

	if info, err = getAssetInfo("update_windows_386"); err != nil {
		t.Fatalf("Failed to get asset info: %v", err)
	}
	if info.OS != OS.Windows || info.Arch != Arch.X86 {
		t.Fatal("Failed to identify update asset.")
	}

	if _, err = getAssetInfo("update_osx_386"); err == nil {
		t.Fatal("Should have ignored the release, \"osx\" is not a valid OS value.")
	}
}

var testClient *ReleaseManager

func getOrCreateTestClient(t *testing.T) *ReleaseManager {
	if testClient == nil {
		testClient = NewReleaseManager(ghAccountOwner, ghAccountRepository)
		if testClient == nil {
			t.Fatal("Failed to create new client.")
		}
		if err := testClient.UpdateAssetsMap(); err != nil {
			t.Fatalf("Failed to update assets map: %v", err)
		}
	}
	return testClient
}

func TestUpdateAssetsMap(t *testing.T) {
	testClient := getOrCreateTestClient(t)
	if testClient.updateAssetsMap == nil {
		t.Fatal("Assets map should not be nil at this point.")
	}
	if len(testClient.updateAssetsMap) == 0 {
		t.Fatal("Assets map is empty.")
	}
	if testClient.latestAssetsMap == nil {
		t.Fatal("Assets map should not be nil at this point.")
	}
	if len(testClient.latestAssetsMap) == 0 {
		t.Fatal("Assets map is empty.")
	}
}

func TestDownloadOldestVersionAndUpgradeIt(t *testing.T) {
	testClient := getOrCreateTestClient(t)
	if len(testClient.updateAssetsMap) == 0 {
		t.Fatal("Assets map is empty.")
	}

	oldestVersionMap := make(map[string]map[string]*Asset)

	// Using the updateAssetsMap to look for the oldest version of each release.
	for os := range testClient.updateAssetsMap {
		for arch := range testClient.updateAssetsMap[os] {
			var oldestAsset *Asset

			for i := range testClient.updateAssetsMap[os][arch] {
				asset := testClient.updateAssetsMap[os][arch][i]
				if oldestAsset == nil {
					oldestAsset = asset
				} else {
					if asset.v.LT(oldestAsset.v) {
						oldestAsset = asset
					}
				}
			}
			if oldestAsset != nil {
				if oldestVersionMap[os] == nil {
					oldestVersionMap[os] = make(map[string]*Asset)
				}

				oldestVersionMap[os][arch] = oldestAsset
			}
		}
	}

	// Let's download each one of the oldest versions.
	var err error
	var p *Patch

	if len(oldestVersionMap) == 0 {
		t.Fatal("No older software versions to test with.")
	}

	tests := 0

	for os := range oldestVersionMap {
		for arch := range oldestVersionMap[os] {
			asset := oldestVersionMap[os][arch]
			newAsset := testClient.latestAssetsMap[os][arch]

			t.Logf("Upgrading %v to %v (%s/%s)", asset.v, newAsset.v, os, arch)

			if asset == newAsset {
				t.Logf("Skipping version %s %s %s", os, arch, asset.v)
				// Skipping
				continue
			}

			// Generate a binary diff of the two assets.
			if p, err = generatePatch(asset.URL, newAsset.URL); err != nil {
				t.Fatalf("Unable to generate patch: %v", err)
			}

			// Apply patch.
			var oldAssetFile string
			if oldAssetFile, err = downloadAsset(asset.URL); err != nil {
				t.Fatal(err)
			}

			var newAssetFile string
			if newAssetFile, err = downloadAsset(newAsset.URL); err != nil {
				t.Fatal(err)
			}

			patchedFile := "_tests/" + path.Base(asset.URL)

			if err = bspatch(oldAssetFile, patchedFile, p.File); err != nil {
				t.Fatalf("Failed to apply binary diff: %v", err)
			}

			// Compare the two versions.
			if fileHash(oldAssetFile) == fileHash(newAssetFile) {
				t.Fatal("Nothing to update, probably not a good test case.")
			}

			if fileHash(patchedFile) != fileHash(newAssetFile) {
				t.Fatal("File hashes after patch must be equal.")
			}

			var cs string
			if cs, err = checksumForFile(patchedFile); err != nil {
				t.Fatalf("Could not get checksum for %s: %v", patchedFile, err)
			}

			if cs == asset.Checksum {
				t.Fatal("Computed checksum for patchedFile must be different than the stored older asset checksum.")
			}

			if cs != newAsset.Checksum {
				t.Fatal("Computed checksum for patchedFile must be equal to the stored newer asset checksum.")
			}

			var ss string
			if ss, err = signatureForFile(patchedFile); err != nil {
				t.Fatalf("Could not get signature for %s: %v", patchedFile, err)
			}

			if ss == asset.Signature {
				t.Fatal("Computed signature for patchedFile must be different than the stored older asset signature.")
			}

			if ss != newAsset.Signature {
				t.Fatal("Computed signature for patchedFile must be equal to the stored newer asset signature.")
			}

			tests++

		}
	}

	if tests == 0 {
		t.Fatal("Seems like there is not any newer software version to test with.")
	}

	// Let's walk over the array again but using CheckForUpdate instead.
	for os := range oldestVersionMap {
		for arch := range oldestVersionMap[os] {
			asset := oldestVersionMap[os][arch]
			params := Params{
				AppVersion: asset.v.String(),
				OS:         asset.OS,
				Arch:       asset.Arch,
				Checksum:   asset.Checksum,
			}

			r, err := testClient.CheckForUpdate(&params, true)
			if err != nil {
				if err == ErrNoUpdateAvailable {
					// That's OK, let's make sure.
					newAsset := testClient.latestAssetsMap[os][arch]
					if asset != newAsset {
						t.Fatal("CheckForUpdate said no update was available!")
					}
				} else {
					t.Fatal("CheckForUpdate: ", err)
				}
			}

			if os != "android" && r.PatchType != PATCHTYPE_BSDIFF {
				t.Fatal("Expecting a patch.")
			}

			if r.Version != testClient.latestAssetsMap[os][arch].v.String() {
				t.Fatalf("Expecting %v, got %v.", testClient.latestAssetsMap[os][arch].v, r.Version)
			}
		}
	}

	// Let's walk again using an odd checksum.
	for os := range oldestVersionMap {
		for arch := range oldestVersionMap[os] {
			asset := oldestVersionMap[os][arch]
			params := Params{
				AppVersion: asset.v.String(),
				OS:         asset.OS,
				Arch:       asset.Arch,
				Checksum:   "?",
			}

			r, err := testClient.CheckForUpdate(&params, true)
			if err != nil {
				if err == ErrNoUpdateAvailable {
					// That's OK, let's make sure.
					newAsset := testClient.latestAssetsMap[os][arch]
					if asset != newAsset {
						t.Fatal("CheckForUpdate said no update was available!")
					}
				} else {
					t.Fatal("CheckForUpdate: ", err)
				}
			}

			if r.PatchType != PATCHTYPE_NONE {
				t.Fatal("Expecting no patch.")
			}

			if r.Version != testClient.latestAssetsMap[os][arch].v.String() {
				t.Fatalf("Expecting %v, got %v.", testClient.latestAssetsMap[os][arch].v, r.Version)
			}
		}
	}
}

func TestCheckOsVersion(t *testing.T) {
	testClient := getOrCreateTestClient(t)
	// change to a version present on GitHub
	lastLanternVersionForWindowsXP = "5.3.4"
	lastLanternVersionForOSXYosemite = "5.3.1"
	osVersionWindowsXP := "5.1.0"
	osVersionOSXYosemite := "14.4.1"
	for _, fields := range [][]string{
		[]string{osVersionWindowsXP, "windows", "386", "3.7.0", lastLanternVersionForWindowsXP},
		[]string{osVersionWindowsXP, "windows", "386", "9.9.9", ""},
		[]string{osVersionOSXYosemite, "darwin", "amd64", "3.7.0", lastLanternVersionForOSXYosemite},
		[]string{osVersionOSXYosemite, "darwin", "amd64", "9.9.9", ""},
		[]string{"", "android", "arm", "9.9.9", ""},
		[]string{"", "android", "arm64", "9.9.9", ""},
	} {
		params := Params{
			OSVersion:  fields[0],
			OS:         fields[1],
			Arch:       fields[2],
			AppVersion: fields[3],
			Checksum:   "fake",
		}
		versionString := fmt.Sprintf("%s(%s%s/%s)", params.AppVersion, params.OS, params.OSVersion, params.Arch)
		r, err := testClient.CheckForUpdate(&params, true)
		if err == nil {
			if r.Version != fields[4] {
				t.Fatalf("Expecting %s for %s, got %s", fields[4], versionString, r.Version)
			}
		} else if err != ErrNoUpdateAvailable {
			t.Fatalf("Expect error '%v' for %s, got '%v'", ErrNoUpdateAvailable, versionString, err)
		}
	}
}
