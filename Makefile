sim:
	fyne package -os iossimulator -certificate "Apple Development: Joe Bologna (Q5BFX4NKWR)" -profile "Tip"
	-xcrun simctl boot "iPhone 16 Pro Max"
	xcrun simctl install "iPhone 16 Pro Max" tip.app

phone:
	fyne package -os ios -certificate "Apple Development: Joe Bologna (Q5BFX4NKWR)" -profile "Tip"
