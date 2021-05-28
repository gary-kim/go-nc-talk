# Go Library for Nextcloud Talk

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.3.0](https://github.com/gary-kim/go-nc-talk/tree/v0.3.0) - 2021-05-28

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.2.2...v0.3.0)

### Added

- enh: add sending complex messages [\#50](https://github.com/gary-kim/go-nc-talk/pull/50) ([@gary-kim](https://github.com/gary-kim))


## [v0.2.2](https://github.com/gary-kim/go-nc-talk/tree/v0.2.2) - 2021-05-27

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.2.1...v0.2.2)

### Fixed

- fix: use sent response for delete message request [\#52](https://github.com/gary-kim/go-nc-talk/pull/52) ([@gary-kim](https://github.com/gary-kim))

## [v0.2.1](https://github.com/gary-kim/go-nc-talk/tree/v0.2.1) - 2021-05-27

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.2.0...v0.2.1)

### Added

- enh: support StatusForbidden response during message fetch [\#47](https://github.com/gary-kim/go-nc-talk/pull/47) ([@gary-kim](https://github.com/gary-kim))

### Fixed

- fix: ensure program doesn't crash when delete fails [\#51](https://github.com/gary-kim/go-nc-talk/pull/51) ([@gary-kim](https://github.com/gary-kim))
- fix: remove conflict with built-in function [\#48](https://github.com/gary-kim/go-nc-talk/pull/48) ([@gary-kim](https://github.com/gary-kim))
- fix: add annotation for ChatMaxLength capability [\#46](https://github.com/gary-kim/go-nc-talk/pull/46) ([@gary-kim](https://github.com/gary-kim))

## [v0.2.0](https://github.com/gary-kim/go-nc-talk/tree/v0.2.0) - 2021-05-25

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.1.7...v0.2.0)

### Added

- Add capabilities and make compatible with Talk 12.0 [\#45](https://github.com/gary-kim/go-nc-talk/pull/45) ([@gary-kim](https://github.com/gary-kim))
- feature: add delete support [\#41](https://github.com/gary-kim/go-nc-talk/pull/41) ([@gary-kim](https://github.com/gary-kim))

### Fixed

- fix: Renovate config [\#43](https://github.com/gary-kim/go-nc-talk/pull/43) ([@gary-kim](https://github.com/gary-kim))

### Dependencies

- chore\(deps\): update golang docker tag to v1.16 [\#40](https://github.com/gary-kim/go-nc-talk/pull/40) ([@gary-kim-bot](https://github.com/gary-kim-bot))
- chore\(deps\): update module stretchr/testify to v1.7.0 [\#38](https://github.com/gary-kim/go-nc-talk/pull/38) ([@gary-kim-bot](https://github.com/gary-kim-bot))

## [v0.1.7](https://github.com/gary-kim/go-nc-talk/tree/v0.1.7) - 2020-12-05

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.1.5...v0.1.7)

### Fixed

- Trim slash from end of NextcloudURL [\#31](https://github.com/gary-kim/go-nc-talk/pull/31) ([@gary-kim](https://github.com/gary-kim))
- Last message can be empty [\#32](https://github.com/gary-kim/go-nc-talk/pull/32) ([@gary-kim](https://github.com/gary-kim))

## [v0.1.5](https://github.com/gary-kim/go-nc-talk/tree/v0.1.5) - 2020-10-12

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.1.4...v0.1.5)

### Fixed

- fix: RemoteDavEndpoint incorrect [\#29](https://github.com/gary-kim/go-nc-talk/pull/29) ([@gary-kim](https://github.com/gary-kim))
- fix: type in deprecation message [\#26](https://github.com/gary-kim/go-nc-talk/pull/26) ([@gary-kim](https://github.com/gary-kim))
- Fix entrypoint URL double slashes \#22 [\#25](https://github.com/gary-kim/go-nc-talk/pull/25) ([@u5surf](https://github.com/u5surf))

### Dependencies

- chore\(deps\): update golang docker tag to v1.15 [\#24](https://github.com/gary-kim/go-nc-talk/pull/24) ([@gary-kim-bot](https://github.com/gary-kim-bot))

## [v0.1.4](https://github.com/gary-kim/go-nc-talk/tree/v0.1.4) - 2020-09-22

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.1.3...v0.1.4)

### Fixed

- Add ActorType for message data [\#18](https://github.com/gary-kim/go-nc-talk/pull/18) ([@gary-kim](https://github.com/gary-kim))

## [v0.1.3](https://github.com/gary-kim/go-nc-talk/tree/v0.1.3) - 2020-09-03

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.1.2...v0.1.3)

### Fixed

- Close response bodies [\#15](https://github.com/gary-kim/go-nc-talk/pull/15) ([@gary-kim](https://github.com/gary-kim))

## [v0.1.2](https://github.com/gary-kim/go-nc-talk/tree/v0.1.2) - 2020-08-28

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.1.1...v0.1.2)

### Fixed

- Use lastReadMessage for first lastKnownMessageId [\#14](https://github.com/gary-kim/go-nc-talk/pull/14) ([@tilosp](https://github.com/tilosp))

## [v0.1.1](https://github.com/gary-kim/go-nc-talk/tree/v0.1.1) - 2020-08-24

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.1.0...v0.1.1)

### Fixed

- ROS type should be of ROST type [\#12](https://github.com/gary-kim/go-nc-talk/pull/12) ([@gary-kim](https://github.com/gary-kim))
- Fix error when sending a message with no RCS data [\#10](https://github.com/gary-kim/go-nc-talk/pull/10) ([@gary-kim](https://github.com/gary-kim))

## [v0.1.0](https://github.com/gary-kim/go-nc-talk/tree/v0.1.0) - 2020-08-13

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.0.2...v0.1.0)

### Added

- Add TLSConfig [\#9](https://github.com/gary-kim/go-nc-talk/pull/9) ([@gary-kim](https://github.com/gary-kim))
- Add Software using this library in README [\#8](https://github.com/gary-kim/go-nc-talk/pull/8) ([@gary-kim](https://github.com/gary-kim))
- Add some basic tests [\#7](https://github.com/gary-kim/go-nc-talk/pull/7) ([@gary-kim](https://github.com/gary-kim))
- Add support for downloading files [\#1](https://github.com/gary-kim/go-nc-talk/pull/1) ([@gary-kim](https://github.com/gary-kim))

### Fixed

- Return error on blank token [\#6](https://github.com/gary-kim/go-nc-talk/pull/6) ([@gary-kim](https://github.com/gary-kim))
- Add v0.0.2 to changelog [\#4](https://github.com/gary-kim/go-nc-talk/pull/4) ([@gary-kim](https://github.com/gary-kim))

## [v0.0.2](https://github.com/gary-kim/go-nc-talk/tree/v0.0.2) - 2020-07-26

[Full Changelog](https://github.com/gary-kim/go-nc-talk/compare/v0.0.1...v0.0.2)

### Changed

- Add installation instructions to README.md [\#2](https://github.com/gary-kim/go-nc-talk/pull/2) ([@gary-kim](https://github.com/gary-kim))

### Fixed

- Fix Capabilities Request [\#3](https://github.com/gary-kim/go-nc-talk/pull/3) ([@gary-kim](https://github.com/gary-kim))

## [v0.0.1](https://github.com/gary-kim/riotchat/tree/v0.0.1) - 2020-07-10

* First release
