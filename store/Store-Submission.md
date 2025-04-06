Error:
Validation failed (409)
Missing or invalid signature. The bundle 'com.focusedforsuccess.Tip-Calculator' at bundle path 'Payload/tip.app' is not signed using an Apple submission certificate. (ID: 3e0f6b04-aa2a-4317-84ca-062ac0015d49)

Fixed By:
re-signing (see makefile)

Error:
Validation failed (409)
Missing Code Signing Entitlements. No entitlements found in bundle 'com.focusedforsuccess.Tip-Calculator' for executable 'Payload/tip.app/main'." (ID: 1e6f1d09-cb80-4ad7-bd76-92772f412582)

Fixed By:
Creating an entitlements file in Xcode storing it in store/tip.entitlements, then including the --entitlements store/tip.entitlements option to codesign

Validation failed (409)
Missing or invalid signature. The bundle 'com.focusedforsuccess.Tip-Calculator' at bundle path 'Payload/tip.app' is not signed using an Apple submission certificate. (ID: 472d6e77-93d9-4c79-a100-eb703e0df01f)

it appears --deep has been deprecated... removing it.

# Tip_App_Store.mobileprovision,distribution.cer

These are the original files used to submit to the app store using Transmission.app.

# Tip_iPhone_App_Store.mobileprovision,ios_distribution.cer

These are new files used to submit to the app store using Transmission.app. They were created in an attempt to fix verification errors. These files are equivalent the previous ones.

Note: the .cer file must be installed using Keychain Access.

xcodebuild -target "${PROJECT_NAME}" -sdk "${TARGET_SDK}" -configuration Release
/usr/bin/xcrun -sdk iphoneos PackageApplication -v "${RELEASE_BUILDDIR}/${APPLICATION_NAME}.app" -o "${BUILD_HISTORY_DIR}/${APPLICATION_NAME}.ipa" --sign "${DEVELOPER_NAME}" --embed "${PROVISONING_PROFILE}”

TODO: Clean up the Distribution certificates, create a new one, check that it has the submission extension, package and check with Transmission app.
