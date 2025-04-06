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
	../fyne/cmd/fyne/fyne package -os ios -certificate "iPhone Distribution: Focused for Success, Inc. (2GC862GT48)" -profile "Tip iPhone App Store"
	find . -name .DS_Store | xargs rm
	rm -fr store/Payload/tip.app
	mv tip.app store/Payload/tip.app
	codesign --remove-signature store/Payload/tip.app
	codesign --sign "iPhone Distribution: Focused for Success, Inc. (2GC862GT48)" store/Payload/tip.app
	codesign --verify --deep --strict store/Payload/tip.app
	cd store; zip -r tip.ipa Payload