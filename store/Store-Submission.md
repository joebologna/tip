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

