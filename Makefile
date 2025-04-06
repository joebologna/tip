sim:
	fyne package -os iossimulator -certificate "Apple Development: Joe Bologna (Q5BFX4NKWR)" -profile "Tip"
	-xcrun simctl boot "iPhone 16 Pro Max"
	xcrun simctl install "iPhone 16 Pro Max" tip.app

phone:
	fyne package -os ios -certificate "Apple Development: Joe Bologna (Q5BFX4NKWR)" -profile "Tip"

# this is unused
dist:
	fyne package -os ios -certificate "Apple Distribution: Focused for Success, Inc. (2GC862GT48)" -profile "Tip App Store"

app_store:
	../fyne/cmd/fyne/fyne package -os ios -certificate "Apple Distribution: Focused for Success, Inc. (2GC862GT48)" -profile "Tip App Store" -release
	find . -name .DS_Store | xargs rm
	rm -fr store/Payload/tip.app
	mv tip.app store/Payload/tip.app
	codesign --remove-signature store/Payload/tip.app
	codesign --sign "Apple Distribution: Focused for Success, Inc. (2GC862GT48)" --entitlements store/tip.entitlements store/Payload/tip.app
	codesign --verify --deep --strict -vvv store/Payload/tip.app
	cd store; zip -r tip.ipa Payload
