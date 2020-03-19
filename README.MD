# VibesEmulator

This is a server emulator for the excellent game "The Test: Fun for friends".

The goal of this project is not to build a fully featured software but mainly a PoC to play with friends without relying on the newer, worst Facebook Messenger's version of this game. It also helps me to learn Android reverse engineering (I really like my setup so I'll list tools below)

This uses sqlite as a local storage.

## Tools used

With these few software you can get pretty far along (That's how I made this whole server)

- JEB Decompiler (Software that decompiles APK and have a lot of IDE-like feature [Like cross-references, Renaming variables, ...])
- Fiddler 4 (Proxy that lets you analyze the requests that your client does on a server, also include HTTPs decryption for apps that does not have their certs pinned [This can be solved with Frida if i'm not mistaken])
- Android Studio (Just for logcat, there is no other great alternative just to filter logcat and have a great display)


- apktool (Patch the APK to use my server instead of the original one)
- apksigner (Re-sign the APK to be installed after patching it)
- XposedFramework (The alternative to building a patched APK is patching it dynamically, I've done this before but I no longer have source-code for this)


- VSCode (IDE to program in go)
- DBeaver (Accessing the database to check what's going on)