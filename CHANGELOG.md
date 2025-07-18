# Changelog

## [0.6.0](https://github.com/edkadigital/plugin-barman-cloud/compare/v0.5.0...v0.6.0) (2025-06-24)


### Features

* Add `liveness` and `readiness` probe support ([#69](https://github.com/edkadigital/plugin-barman-cloud/issues/69)) ([5fd9449](https://github.com/edkadigital/plugin-barman-cloud/commit/5fd9449b27394756e0baf76b1356900850f687a6))
* Additional environment variables ([#81](https://github.com/edkadigital/plugin-barman-cloud/issues/81)) ([be40375](https://github.com/edkadigital/plugin-barman-cloud/commit/be4037529c44858278dd80e3eb32f39f3f68c5c6))
* **deps:** Update dependency barman to v3.14.0 ([#368](https://github.com/edkadigital/plugin-barman-cloud/issues/368)) ([3550013](https://github.com/edkadigital/plugin-barman-cloud/commit/35500130bf0fe25eb3a191bc78f4818c318acf26))
* Forbid usage of `.spec.configuration.serverName` in ObjectStore ([#336](https://github.com/edkadigital/plugin-barman-cloud/issues/336)) ([3420f43](https://github.com/edkadigital/plugin-barman-cloud/commit/3420f430739ac8518c83cd3b23bf6a8e42b411f7)), closes [#334](https://github.com/edkadigital/plugin-barman-cloud/issues/334)
* Generate apidoc using genref ([#228](https://github.com/edkadigital/plugin-barman-cloud/issues/228)) ([74bdb9a](https://github.com/edkadigital/plugin-barman-cloud/commit/74bdb9a590f169eade4eea27caa85fc3b1809e41)), closes [#206](https://github.com/edkadigital/plugin-barman-cloud/issues/206)
* Grant permissions to read secrets ([#25](https://github.com/edkadigital/plugin-barman-cloud/issues/25)) ([76383a3](https://github.com/edkadigital/plugin-barman-cloud/commit/76383a30afd3bd829f01936dc3dfc81f1d189d2d))
* Implement evaluate lifecycle hook ([#222](https://github.com/edkadigital/plugin-barman-cloud/issues/222)) ([a7ef56b](https://github.com/edkadigital/plugin-barman-cloud/commit/a7ef56b6e7a8abfcf312f42190b5c3828f9b2a79))
* Lenient decoding of CNPG resources ([#192](https://github.com/edkadigital/plugin-barman-cloud/issues/192)) ([13e3fab](https://github.com/edkadigital/plugin-barman-cloud/commit/13e3fab2688ec6ea342ed7304680025f98e6af27))
* Log the downloaded backup catalog before restore ([#323](https://github.com/edkadigital/plugin-barman-cloud/issues/323)) ([9db184f](https://github.com/edkadigital/plugin-barman-cloud/commit/9db184f5d4c325ed18aeb4fba6c57c28b0e3ae40)), closes [#319](https://github.com/edkadigital/plugin-barman-cloud/issues/319)
* Operator plugin and manifests ([#18](https://github.com/edkadigital/plugin-barman-cloud/issues/18)) ([dd6548c](https://github.com/edkadigital/plugin-barman-cloud/commit/dd6548c4a26031324975d97aee345e4e6a2e7efa))
* Release-please cleanup ([#115](https://github.com/edkadigital/plugin-barman-cloud/issues/115)) ([cd03c55](https://github.com/edkadigital/plugin-barman-cloud/commit/cd03c556ef86c429b8699961eb24e1361b5759ff)), closes [#114](https://github.com/edkadigital/plugin-barman-cloud/issues/114)
* Retention policy ([#191](https://github.com/edkadigital/plugin-barman-cloud/issues/191)) ([fecd1e9](https://github.com/edkadigital/plugin-barman-cloud/commit/fecd1e9513ce1748a289840f735a2f23a0ce5218))
* Separate recovery and cluster object store ([#76](https://github.com/edkadigital/plugin-barman-cloud/issues/76)) ([e30edd2](https://github.com/edkadigital/plugin-barman-cloud/commit/e30edd2318d76e10fd7af344c0e4326f1e5033ec))
* Separate recovery object store from replica source ([#83](https://github.com/edkadigital/plugin-barman-cloud/issues/83)) ([e4735a2](https://github.com/edkadigital/plugin-barman-cloud/commit/e4735a2f85724cf8493f513658783e5330c3efcf))
* Sidecar injection and loading ([#22](https://github.com/edkadigital/plugin-barman-cloud/issues/22)) ([ea6ee30](https://github.com/edkadigital/plugin-barman-cloud/commit/ea6ee30d2ea30f9e9df22002ce5f5a68fcb37ade))
* Sidecar role and rolebinding ([#23](https://github.com/edkadigital/plugin-barman-cloud/issues/23)) ([2f62d53](https://github.com/edkadigital/plugin-barman-cloud/commit/2f62d539c949f344cb5534b7ffbb90860663a106))
* **sidecar:** Add resource requirements and limits ([#307](https://github.com/edkadigital/plugin-barman-cloud/issues/307)) ([4bb3471](https://github.com/edkadigital/plugin-barman-cloud/commit/4bb347121d3328783ca9eceb656863cde37cb8aa)), closes [#253](https://github.com/edkadigital/plugin-barman-cloud/issues/253)
* **spike:** Backup method ([#20](https://github.com/edkadigital/plugin-barman-cloud/issues/20)) ([9fa1c0b](https://github.com/edkadigital/plugin-barman-cloud/commit/9fa1c0beab4882af3f4c737d049b5bafcf7e28a6))
* **spike:** Restore ([#29](https://github.com/edkadigital/plugin-barman-cloud/issues/29)) ([240077c](https://github.com/edkadigital/plugin-barman-cloud/commit/240077c77192d9572767d7ec76d02e578b94faca))
* **spike:** Wal-archive and wal-restore methods ([#4](https://github.com/edkadigital/plugin-barman-cloud/issues/4)) ([1c86ff6](https://github.com/edkadigital/plugin-barman-cloud/commit/1c86ff65747b5b348fb1ed2b0e5b0594fd156116))
* Support additional compression methods in the sidecar image ([#158](https://github.com/edkadigital/plugin-barman-cloud/issues/158)) ([ee5fd84](https://github.com/edkadigital/plugin-barman-cloud/commit/ee5fd840924c0997f301764af32a684aa8424b22)), closes [#127](https://github.com/edkadigital/plugin-barman-cloud/issues/127)
* Support custom CA certificates ([#198](https://github.com/edkadigital/plugin-barman-cloud/issues/198)) ([fcbc472](https://github.com/edkadigital/plugin-barman-cloud/commit/fcbc47209222f712178ba422020c88eef7d50c08))
* Support lz4, xz, and zstandard compressions ([#201](https://github.com/edkadigital/plugin-barman-cloud/issues/201)) ([795313f](https://github.com/edkadigital/plugin-barman-cloud/commit/795313f4aa2f4888fdf2cb711de74aaea7b045a7)), closes [#200](https://github.com/edkadigital/plugin-barman-cloud/issues/200)
* Support snapshot recovery job ([#258](https://github.com/edkadigital/plugin-barman-cloud/issues/258)) ([e00024f](https://github.com/edkadigital/plugin-barman-cloud/commit/e00024f136996305999c0440ae9b48861828e160))
* Upgrade Barman to 3.13.0 ([#209](https://github.com/edkadigital/plugin-barman-cloud/issues/209)) ([56d8cce](https://github.com/edkadigital/plugin-barman-cloud/commit/56d8cceb3b8c7a17f3dcdd2dc14b48a725aaea9f)), closes [#208](https://github.com/edkadigital/plugin-barman-cloud/issues/208)
* **wal:** Parallel WAL archiving ([#262](https://github.com/edkadigital/plugin-barman-cloud/issues/262)) ([88fd3e5](https://github.com/edkadigital/plugin-barman-cloud/commit/88fd3e504f35e004fab47ca33a2e67dd40120e2c)), closes [#260](https://github.com/edkadigital/plugin-barman-cloud/issues/260) [#266](https://github.com/edkadigital/plugin-barman-cloud/issues/266)


### Bug Fixes

* Avoid injecting the plugin environment into the PG container ([#62](https://github.com/edkadigital/plugin-barman-cloud/issues/62)) ([9c77e3d](https://github.com/edkadigital/plugin-barman-cloud/commit/9c77e3de9f05a56c567c9fa6b0f75ca55a05ddf8))
* Controller and sidecar containers run as non-root ([#225](https://github.com/edkadigital/plugin-barman-cloud/issues/225)) ([5788c1f](https://github.com/edkadigital/plugin-barman-cloud/commit/5788c1f72794a331e9176dabc625a5937abff010)), closes [#177](https://github.com/edkadigital/plugin-barman-cloud/issues/177)
* Custom CA support for retention policies ([#224](https://github.com/edkadigital/plugin-barman-cloud/issues/224)) ([bac7b67](https://github.com/edkadigital/plugin-barman-cloud/commit/bac7b673a2ef239dd28bd2d1eced083009ad8ba6)), closes [#220](https://github.com/edkadigital/plugin-barman-cloud/issues/220)
* **deps:** Lock file maintenance documentation dependencies ([#379](https://github.com/edkadigital/plugin-barman-cloud/issues/379)) ([a0327ea](https://github.com/edkadigital/plugin-barman-cloud/commit/a0327ea574558d6c1a913e13a12bb454818900a7))
* **deps:** Lock file maintenance documentation dependencies ([#399](https://github.com/edkadigital/plugin-barman-cloud/issues/399)) ([7146c51](https://github.com/edkadigital/plugin-barman-cloud/commit/7146c51de11a5d673aef23e36e07a2b0c528d3b7))
* **deps:** Lock file maintenance documentation dependencies ([#407](https://github.com/edkadigital/plugin-barman-cloud/issues/407)) ([4d323c2](https://github.com/edkadigital/plugin-barman-cloud/commit/4d323c2d3df2bcd52c126b369922bec67db68a2c))
* **deps:** Update all non-major go dependencies ([#103](https://github.com/edkadigital/plugin-barman-cloud/issues/103)) ([55258f6](https://github.com/edkadigital/plugin-barman-cloud/commit/55258f69008d1475f65d549d47a6c87485624e28))
* **deps:** Update all non-major go dependencies ([#15](https://github.com/edkadigital/plugin-barman-cloud/issues/15)) ([3289d91](https://github.com/edkadigital/plugin-barman-cloud/commit/3289d91db4f924bad2f7f6dc8901f4544616233e))
* **deps:** Update all non-major go dependencies ([#152](https://github.com/edkadigital/plugin-barman-cloud/issues/152)) ([e77799a](https://github.com/edkadigital/plugin-barman-cloud/commit/e77799af028ba892ed8f3261554682c1b540a7f5))
* **deps:** Update all non-major go dependencies ([#19](https://github.com/edkadigital/plugin-barman-cloud/issues/19)) ([3785162](https://github.com/edkadigital/plugin-barman-cloud/commit/378516225d6441dcba32bcd7533d54501d91cf08))
* **deps:** Update all non-major go dependencies ([#213](https://github.com/edkadigital/plugin-barman-cloud/issues/213)) ([a5b8649](https://github.com/edkadigital/plugin-barman-cloud/commit/a5b8649bd0eac1df6e51291ff197a6a548d0f479))
* **deps:** Update all non-major go dependencies ([#219](https://github.com/edkadigital/plugin-barman-cloud/issues/219)) ([0d4a3d3](https://github.com/edkadigital/plugin-barman-cloud/commit/0d4a3d38f77e9d51a3f627fa768673e3c4b5e650))
* **deps:** Update all non-major go dependencies ([#246](https://github.com/edkadigital/plugin-barman-cloud/issues/246)) ([ed1feaa](https://github.com/edkadigital/plugin-barman-cloud/commit/ed1feaaddcddfabd48a2d9a28013e7585d8babd6))
* **deps:** Update all non-major go dependencies ([#278](https://github.com/edkadigital/plugin-barman-cloud/issues/278)) ([010c9b9](https://github.com/edkadigital/plugin-barman-cloud/commit/010c9b93d4e2d06eb89ba49219f15144c98515cf))
* **deps:** Update all non-major go dependencies ([#366](https://github.com/edkadigital/plugin-barman-cloud/issues/366)) ([1097abb](https://github.com/edkadigital/plugin-barman-cloud/commit/1097abbd1d26502a3cfc81f932bffd5bef2377a4))
* **deps:** Update all non-major go dependencies ([#9](https://github.com/edkadigital/plugin-barman-cloud/issues/9)) ([435986b](https://github.com/edkadigital/plugin-barman-cloud/commit/435986b7a1e7bf9e5d4d1c018c37fd6e28f2aaa7))
* **deps:** Update github.com/cloudnative-pg/cloudnative-pg digest to 34ab236 ([#180](https://github.com/edkadigital/plugin-barman-cloud/issues/180)) ([e9e636a](https://github.com/edkadigital/plugin-barman-cloud/commit/e9e636ada08de4a1f6db0a31e2f133e703580394))
* **deps:** Update golang.org/x/net ([#188](https://github.com/edkadigital/plugin-barman-cloud/issues/188)) ([aba0748](https://github.com/edkadigital/plugin-barman-cloud/commit/aba07487891b731b6439429c7b30da21bc260d5f))
* **deps:** Update k8s.io/utils digest to 0f33e8f ([#301](https://github.com/edkadigital/plugin-barman-cloud/issues/301)) ([ab398d7](https://github.com/edkadigital/plugin-barman-cloud/commit/ab398d7d30ebe241b2b682c42c4b129254955b24))
* **deps:** Update k8s.io/utils digest to 1f6e0b7 ([#237](https://github.com/edkadigital/plugin-barman-cloud/issues/237)) ([792679f](https://github.com/edkadigital/plugin-barman-cloud/commit/792679ff673f60deeac3293d4bfb3e5182a09bef))
* **deps:** Update k8s.io/utils digest to 24370be ([#90](https://github.com/edkadigital/plugin-barman-cloud/issues/90)) ([1bc5fac](https://github.com/edkadigital/plugin-barman-cloud/commit/1bc5facc9acadbcdb88e76ec294f6c4c050c4daa))
* **deps:** Update k8s.io/utils digest to 4c0f3b2 ([#392](https://github.com/edkadigital/plugin-barman-cloud/issues/392)) ([e58973c](https://github.com/edkadigital/plugin-barman-cloud/commit/e58973cd55b89c2e4615cf67c85b08627590aae1))
* **deps:** Update kubernetes packages to v0.31.1 ([#10](https://github.com/edkadigital/plugin-barman-cloud/issues/10)) ([76486c2](https://github.com/edkadigital/plugin-barman-cloud/commit/76486c28637fa10be3b8b5f260d5b626ac142ca4))
* **deps:** Update kubernetes packages to v0.31.3 ([#64](https://github.com/edkadigital/plugin-barman-cloud/issues/64)) ([c639af1](https://github.com/edkadigital/plugin-barman-cloud/commit/c639af1295123c12d462d52b769ac0c973c22c93))
* **deps:** Update kubernetes packages to v0.32.1 ([#147](https://github.com/edkadigital/plugin-barman-cloud/issues/147)) ([dbc5550](https://github.com/edkadigital/plugin-barman-cloud/commit/dbc5550c9c503dfb0a6206a244995cdda9d28c1d))
* **deps:** Update kubernetes packages to v0.32.2 ([#172](https://github.com/edkadigital/plugin-barman-cloud/issues/172)) ([bb9658b](https://github.com/edkadigital/plugin-barman-cloud/commit/bb9658b28c95f9b7e1f202dcf2be76bff7756960))
* **deps:** Update kubernetes packages to v0.32.3 ([#216](https://github.com/edkadigital/plugin-barman-cloud/issues/216)) ([9d22676](https://github.com/edkadigital/plugin-barman-cloud/commit/9d22676f2a5667b516a4f496ab6188a2333e5333))
* **deps:** Update kubernetes packages to v0.33.0 ([#281](https://github.com/edkadigital/plugin-barman-cloud/issues/281)) ([c6f36d5](https://github.com/edkadigital/plugin-barman-cloud/commit/c6f36d57562a99175e2d3d446ca2d7e7c36b09c3))
* **deps:** Update kubernetes packages to v0.33.1 ([#361](https://github.com/edkadigital/plugin-barman-cloud/issues/361)) ([9d4bc45](https://github.com/edkadigital/plugin-barman-cloud/commit/9d4bc456b09b9d79c1ad58f686c8201885ffe4ce))
* **deps:** Update module github.com/cert-manager/cert-manager to v1.16.2 [security] ([#63](https://github.com/edkadigital/plugin-barman-cloud/issues/63)) ([53d2c09](https://github.com/edkadigital/plugin-barman-cloud/commit/53d2c0999313b1447d873b27b1f87e1dd93c6e6a))
* **deps:** Update module github.com/cloudnative-pg/api to v1 ([#131](https://github.com/edkadigital/plugin-barman-cloud/issues/131)) ([0c8ff74](https://github.com/edkadigital/plugin-barman-cloud/commit/0c8ff7426ff15623deba0c9603ba76dece3cb6a5))
* **deps:** Update module github.com/cloudnative-pg/cnpg-i-machinery to v0.1.2 ([#182](https://github.com/edkadigital/plugin-barman-cloud/issues/182)) ([12cd519](https://github.com/edkadigital/plugin-barman-cloud/commit/12cd5195234ee17ca0b09c2448cc9dc50c614149))
* **deps:** Update module google.golang.org/grpc to v1.71.0 ([#187](https://github.com/edkadigital/plugin-barman-cloud/issues/187)) ([e1f1660](https://github.com/edkadigital/plugin-barman-cloud/commit/e1f166023f55fb02d987ac011e3580af1f9d273a))
* **deps:** Update module google.golang.org/grpc to v1.72.1 ([#345](https://github.com/edkadigital/plugin-barman-cloud/issues/345)) ([d9fd8dd](https://github.com/edkadigital/plugin-barman-cloud/commit/d9fd8dd8681e33ec64c911eade3516a73f793ac5))
* **deps:** Update module google.golang.org/grpc to v1.73.0 ([#394](https://github.com/edkadigital/plugin-barman-cloud/issues/394)) ([1365906](https://github.com/edkadigital/plugin-barman-cloud/commit/1365906204d895cac78ef93d5753d0b5f717c9ac))
* **deps:** Update module k8s.io/client-go to v0.31.1 ([#16](https://github.com/edkadigital/plugin-barman-cloud/issues/16)) ([cbefe26](https://github.com/edkadigital/plugin-barman-cloud/commit/cbefe26440203e88f8d60335b64f32b01249ba0d))
* **deps:** Update module sigs.k8s.io/controller-runtime to v0.19.2 ([#67](https://github.com/edkadigital/plugin-barman-cloud/issues/67)) ([74d4f5d](https://github.com/edkadigital/plugin-barman-cloud/commit/74d4f5d1902ed557375adff3e775b35dd662d2fc))
* **deps:** Update module sigs.k8s.io/controller-runtime to v0.19.3 ([#78](https://github.com/edkadigital/plugin-barman-cloud/issues/78)) ([497022f](https://github.com/edkadigital/plugin-barman-cloud/commit/497022f1967c90598285573bc54341d22308f2f0))
* **deps:** Update module sigs.k8s.io/controller-runtime to v0.21.0 ([#367](https://github.com/edkadigital/plugin-barman-cloud/issues/367)) ([fecc2f7](https://github.com/edkadigital/plugin-barman-cloud/commit/fecc2f7d28e5ad58c6370f0a26014908ce4caaaf))
* **deps:** Update module sigs.k8s.io/kustomize/api to v0.18.0 ([#51](https://github.com/edkadigital/plugin-barman-cloud/issues/51)) ([b2d3032](https://github.com/edkadigital/plugin-barman-cloud/commit/b2d303205499ccca426fe2b72964eeefa6556fdd))
* **deps:** Update module sigs.k8s.io/kustomize/api to v0.19.0 ([#148](https://github.com/edkadigital/plugin-barman-cloud/issues/148)) ([9ba6351](https://github.com/edkadigital/plugin-barman-cloud/commit/9ba63518f929748f4a422eaa58293c8125b7a5f1))
* **deps:** Update react monorepo to v19.1.0 ([#286](https://github.com/edkadigital/plugin-barman-cloud/issues/286)) ([99f31a1](https://github.com/edkadigital/plugin-barman-cloud/commit/99f31a1e5e0313534699c49393edc6beabac60ec))
* **deps:** Use latest commit from CNPG 1.25 branch ([#178](https://github.com/edkadigital/plugin-barman-cloud/issues/178)) ([dfbeaf8](https://github.com/edkadigital/plugin-barman-cloud/commit/dfbeaf802ec98357fdbb92b5fcefc38a29939cfe))
* Do not add barman-certificates projection if not needed ([#354](https://github.com/edkadigital/plugin-barman-cloud/issues/354)) ([918823d](https://github.com/edkadigital/plugin-barman-cloud/commit/918823dbf1c78e5460f83af50bf85be6c1aefafe))
* **docs:** Fix TOC links ([#261](https://github.com/edkadigital/plugin-barman-cloud/issues/261)) ([2bb5e90](https://github.com/edkadigital/plugin-barman-cloud/commit/2bb5e90357b2defd6fdaa8ff9982e21f58bc5ecc))
* **docs:** Replace "no downtime" with "without data loss" ([#349](https://github.com/edkadigital/plugin-barman-cloud/issues/349)) ([5e1b845](https://github.com/edkadigital/plugin-barman-cloud/commit/5e1b845caedb67cf79173af3a319d55260b21627))
* Duplicate certificate projections ([#331](https://github.com/edkadigital/plugin-barman-cloud/issues/331)) ([8c20e4f](https://github.com/edkadigital/plugin-barman-cloud/commit/8c20e4fe8578b5b18277ce2ae8ba11783b1cac84)), closes [#329](https://github.com/edkadigital/plugin-barman-cloud/issues/329)
* Ensure restore configuration points to manager `wal-restore` ([#68](https://github.com/edkadigital/plugin-barman-cloud/issues/68)) ([afd4603](https://github.com/edkadigital/plugin-barman-cloud/commit/afd4603023ce0f245687856eb05d9a30875b8bac))
* Exit code 0 on clean shutdown ([#70](https://github.com/edkadigital/plugin-barman-cloud/issues/70)) ([9d8fa07](https://github.com/edkadigital/plugin-barman-cloud/commit/9d8fa079fec6b82c5aef6397b4b6318fbe9ebb0b))
* Obsolete deepcopy ([1e6c69b](https://github.com/edkadigital/plugin-barman-cloud/commit/1e6c69bac022914732fbaabb5bae0969893f5049))
* Remove lifecycle `Pod` `Patch` subscription ([#378](https://github.com/edkadigital/plugin-barman-cloud/issues/378)) ([40316b5](https://github.com/edkadigital/plugin-barman-cloud/commit/40316b5f2d72deac0f042ceecd271a97b369a62f))
* Replica source object store on replica clusters being promoted ([#96](https://github.com/edkadigital/plugin-barman-cloud/issues/96)) ([4656d44](https://github.com/edkadigital/plugin-barman-cloud/commit/4656d44c85a3d05e38cb536b2db69aa47c575619))
* Role patching ([#325](https://github.com/edkadigital/plugin-barman-cloud/issues/325)) ([f484b9e](https://github.com/edkadigital/plugin-barman-cloud/commit/f484b9e748ad776f7ecec0ed83a2b2424fde2dfc)), closes [#318](https://github.com/edkadigital/plugin-barman-cloud/issues/318)
* Use a fixed golangci-lint version ([#230](https://github.com/edkadigital/plugin-barman-cloud/issues/230)) ([78fe21b](https://github.com/edkadigital/plugin-barman-cloud/commit/78fe21b24dc9366c34260babe6b049a310abe9f0))

## [0.5.0](https://github.com/edkadigital/plugin-barman-cloud/compare/v0.4.1...v0.5.0) (2025-06-03)


### Features

* **deps:** Update dependency barman to v3.14.0 ([#368](https://github.com/edkadigital/plugin-barman-cloud/issues/368)) ([3550013](https://github.com/edkadigital/plugin-barman-cloud/commit/35500130bf0fe25eb3a191bc78f4818c318acf26))


### Bug Fixes

* Remove lifecycle `Pod` `Patch` subscription ([#378](https://github.com/edkadigital/plugin-barman-cloud/issues/378)) ([40316b5](https://github.com/edkadigital/plugin-barman-cloud/commit/40316b5f2d72deac0f042ceecd271a97b369a62f))

## [0.4.1](https://github.com/edkadigital/plugin-barman-cloud/compare/v0.4.0...v0.4.1) (2025-05-29)


### Bug Fixes

* **deps:** Update all non-major go dependencies ([#366](https://github.com/edkadigital/plugin-barman-cloud/issues/366)) ([1097abb](https://github.com/edkadigital/plugin-barman-cloud/commit/1097abbd1d26502a3cfc81f932bffd5bef2377a4))
* **deps:** Update kubernetes packages to v0.33.1 ([#361](https://github.com/edkadigital/plugin-barman-cloud/issues/361)) ([9d4bc45](https://github.com/edkadigital/plugin-barman-cloud/commit/9d4bc456b09b9d79c1ad58f686c8201885ffe4ce))
* **deps:** Update module google.golang.org/grpc to v1.72.1 ([#345](https://github.com/edkadigital/plugin-barman-cloud/issues/345)) ([d9fd8dd](https://github.com/edkadigital/plugin-barman-cloud/commit/d9fd8dd8681e33ec64c911eade3516a73f793ac5))
* **deps:** Update module sigs.k8s.io/controller-runtime to v0.21.0 ([#367](https://github.com/edkadigital/plugin-barman-cloud/issues/367)) ([fecc2f7](https://github.com/edkadigital/plugin-barman-cloud/commit/fecc2f7d28e5ad58c6370f0a26014908ce4caaaf))
* Do not add barman-certificates projection if not needed ([#354](https://github.com/edkadigital/plugin-barman-cloud/issues/354)) ([918823d](https://github.com/edkadigital/plugin-barman-cloud/commit/918823dbf1c78e5460f83af50bf85be6c1aefafe))
* **docs:** Replace "no downtime" with "without data loss" ([#349](https://github.com/edkadigital/plugin-barman-cloud/issues/349)) ([5e1b845](https://github.com/edkadigital/plugin-barman-cloud/commit/5e1b845caedb67cf79173af3a319d55260b21627))

## [0.4.0](https://github.com/edkadigital/plugin-barman-cloud/compare/v0.3.0...v0.4.0) (2025-05-12)


### Features

* Forbid usage of `.spec.configuration.serverName` in ObjectStore ([#336](https://github.com/edkadigital/plugin-barman-cloud/issues/336)) ([3420f43](https://github.com/edkadigital/plugin-barman-cloud/commit/3420f430739ac8518c83cd3b23bf6a8e42b411f7)), closes [#334](https://github.com/edkadigital/plugin-barman-cloud/issues/334)
* Log the downloaded backup catalog before restore ([#323](https://github.com/edkadigital/plugin-barman-cloud/issues/323)) ([9db184f](https://github.com/edkadigital/plugin-barman-cloud/commit/9db184f5d4c325ed18aeb4fba6c57c28b0e3ae40)), closes [#319](https://github.com/edkadigital/plugin-barman-cloud/issues/319)
* **sidecar:** Add resource requirements and limits ([#307](https://github.com/edkadigital/plugin-barman-cloud/issues/307)) ([4bb3471](https://github.com/edkadigital/plugin-barman-cloud/commit/4bb347121d3328783ca9eceb656863cde37cb8aa)), closes [#253](https://github.com/edkadigital/plugin-barman-cloud/issues/253)
* Support snapshot recovery job ([#258](https://github.com/edkadigital/plugin-barman-cloud/issues/258)) ([e00024f](https://github.com/edkadigital/plugin-barman-cloud/commit/e00024f136996305999c0440ae9b48861828e160))
* **wal:** Parallel WAL archiving ([#262](https://github.com/edkadigital/plugin-barman-cloud/issues/262)) ([88fd3e5](https://github.com/edkadigital/plugin-barman-cloud/commit/88fd3e504f35e004fab47ca33a2e67dd40120e2c)), closes [#260](https://github.com/edkadigital/plugin-barman-cloud/issues/260) [#266](https://github.com/edkadigital/plugin-barman-cloud/issues/266)


### Bug Fixes

* [#260](https://github.com/edkadigital/plugin-barman-cloud/issues/260) ([88fd3e5](https://github.com/edkadigital/plugin-barman-cloud/commit/88fd3e504f35e004fab47ca33a2e67dd40120e2c))
* [#266](https://github.com/edkadigital/plugin-barman-cloud/issues/266) ([88fd3e5](https://github.com/edkadigital/plugin-barman-cloud/commit/88fd3e504f35e004fab47ca33a2e67dd40120e2c))
* **deps:** Update all non-major go dependencies ([#246](https://github.com/edkadigital/plugin-barman-cloud/issues/246)) ([ed1feaa](https://github.com/edkadigital/plugin-barman-cloud/commit/ed1feaaddcddfabd48a2d9a28013e7585d8babd6))
* **deps:** Update all non-major go dependencies ([#278](https://github.com/edkadigital/plugin-barman-cloud/issues/278)) ([010c9b9](https://github.com/edkadigital/plugin-barman-cloud/commit/010c9b93d4e2d06eb89ba49219f15144c98515cf))
* **deps:** Update k8s.io/utils digest to 0f33e8f ([#301](https://github.com/edkadigital/plugin-barman-cloud/issues/301)) ([ab398d7](https://github.com/edkadigital/plugin-barman-cloud/commit/ab398d7d30ebe241b2b682c42c4b129254955b24))
* **deps:** Update kubernetes packages to v0.33.0 ([#281](https://github.com/edkadigital/plugin-barman-cloud/issues/281)) ([c6f36d5](https://github.com/edkadigital/plugin-barman-cloud/commit/c6f36d57562a99175e2d3d446ca2d7e7c36b09c3))
* **deps:** Update react monorepo to v19.1.0 ([#286](https://github.com/edkadigital/plugin-barman-cloud/issues/286)) ([99f31a1](https://github.com/edkadigital/plugin-barman-cloud/commit/99f31a1e5e0313534699c49393edc6beabac60ec))
* **docs:** Fix TOC links ([#261](https://github.com/edkadigital/plugin-barman-cloud/issues/261)) ([2bb5e90](https://github.com/edkadigital/plugin-barman-cloud/commit/2bb5e90357b2defd6fdaa8ff9982e21f58bc5ecc))
* Duplicate certificate projections ([#331](https://github.com/edkadigital/plugin-barman-cloud/issues/331)) ([8c20e4f](https://github.com/edkadigital/plugin-barman-cloud/commit/8c20e4fe8578b5b18277ce2ae8ba11783b1cac84)), closes [#329](https://github.com/edkadigital/plugin-barman-cloud/issues/329)
* Role patching ([#325](https://github.com/edkadigital/plugin-barman-cloud/issues/325)) ([f484b9e](https://github.com/edkadigital/plugin-barman-cloud/commit/f484b9e748ad776f7ecec0ed83a2b2424fde2dfc)), closes [#318](https://github.com/edkadigital/plugin-barman-cloud/issues/318)

## [0.3.0](https://github.com/edkadigital/plugin-barman-cloud/compare/v0.2.0...v0.3.0) (2025-03-28)


### Features

* Generate apidoc using genref ([#228](https://github.com/edkadigital/plugin-barman-cloud/issues/228)) ([74bdb9a](https://github.com/edkadigital/plugin-barman-cloud/commit/74bdb9a590f169eade4eea27caa85fc3b1809e41)), closes [#206](https://github.com/edkadigital/plugin-barman-cloud/issues/206)
* Implement evaluate lifecycle hook ([#222](https://github.com/edkadigital/plugin-barman-cloud/issues/222)) ([a7ef56b](https://github.com/edkadigital/plugin-barman-cloud/commit/a7ef56b6e7a8abfcf312f42190b5c3828f9b2a79))
* Lenient decoding of CNPG resources ([#192](https://github.com/edkadigital/plugin-barman-cloud/issues/192)) ([13e3fab](https://github.com/edkadigital/plugin-barman-cloud/commit/13e3fab2688ec6ea342ed7304680025f98e6af27))
* Retention policy ([#191](https://github.com/edkadigital/plugin-barman-cloud/issues/191)) ([fecd1e9](https://github.com/edkadigital/plugin-barman-cloud/commit/fecd1e9513ce1748a289840f735a2f23a0ce5218))
* Support custom CA certificates ([#198](https://github.com/edkadigital/plugin-barman-cloud/issues/198)) ([fcbc472](https://github.com/edkadigital/plugin-barman-cloud/commit/fcbc47209222f712178ba422020c88eef7d50c08))
* Support lz4, xz, and zstandard compressions ([#201](https://github.com/edkadigital/plugin-barman-cloud/issues/201)) ([795313f](https://github.com/edkadigital/plugin-barman-cloud/commit/795313f4aa2f4888fdf2cb711de74aaea7b045a7)), closes [#200](https://github.com/edkadigital/plugin-barman-cloud/issues/200)
* Upgrade Barman to 3.13.0 ([#209](https://github.com/edkadigital/plugin-barman-cloud/issues/209)) ([56d8cce](https://github.com/edkadigital/plugin-barman-cloud/commit/56d8cceb3b8c7a17f3dcdd2dc14b48a725aaea9f)), closes [#208](https://github.com/edkadigital/plugin-barman-cloud/issues/208)


### Bug Fixes

* Controller and sidecar containers run as non-root ([#225](https://github.com/edkadigital/plugin-barman-cloud/issues/225)) ([5788c1f](https://github.com/edkadigital/plugin-barman-cloud/commit/5788c1f72794a331e9176dabc625a5937abff010)), closes [#177](https://github.com/edkadigital/plugin-barman-cloud/issues/177)
* Custom CA support for retention policies ([#224](https://github.com/edkadigital/plugin-barman-cloud/issues/224)) ([bac7b67](https://github.com/edkadigital/plugin-barman-cloud/commit/bac7b673a2ef239dd28bd2d1eced083009ad8ba6)), closes [#220](https://github.com/edkadigital/plugin-barman-cloud/issues/220)
* **deps:** Update all non-major go dependencies ([#213](https://github.com/edkadigital/plugin-barman-cloud/issues/213)) ([a5b8649](https://github.com/edkadigital/plugin-barman-cloud/commit/a5b8649bd0eac1df6e51291ff197a6a548d0f479))
* **deps:** Update all non-major go dependencies ([#219](https://github.com/edkadigital/plugin-barman-cloud/issues/219)) ([0d4a3d3](https://github.com/edkadigital/plugin-barman-cloud/commit/0d4a3d38f77e9d51a3f627fa768673e3c4b5e650))
* **deps:** Update k8s.io/utils digest to 1f6e0b7 ([#237](https://github.com/edkadigital/plugin-barman-cloud/issues/237)) ([792679f](https://github.com/edkadigital/plugin-barman-cloud/commit/792679ff673f60deeac3293d4bfb3e5182a09bef))
* **deps:** Update kubernetes packages to v0.32.3 ([#216](https://github.com/edkadigital/plugin-barman-cloud/issues/216)) ([9d22676](https://github.com/edkadigital/plugin-barman-cloud/commit/9d22676f2a5667b516a4f496ab6188a2333e5333))
* Use a fixed golangci-lint version ([#230](https://github.com/edkadigital/plugin-barman-cloud/issues/230)) ([78fe21b](https://github.com/edkadigital/plugin-barman-cloud/commit/78fe21b24dc9366c34260babe6b049a310abe9f0))

## [0.2.0](https://github.com/edkadigital/plugin-barman-cloud/compare/v0.1.0...v0.2.0) (2025-03-05)


### Features

* Release-please cleanup ([#115](https://github.com/edkadigital/plugin-barman-cloud/issues/115)) ([cd03c55](https://github.com/edkadigital/plugin-barman-cloud/commit/cd03c556ef86c429b8699961eb24e1361b5759ff)), closes [#114](https://github.com/edkadigital/plugin-barman-cloud/issues/114)
* Support additional compression methods in the sidecar image ([#158](https://github.com/edkadigital/plugin-barman-cloud/issues/158)) ([ee5fd84](https://github.com/edkadigital/plugin-barman-cloud/commit/ee5fd840924c0997f301764af32a684aa8424b22)), closes [#127](https://github.com/edkadigital/plugin-barman-cloud/issues/127)


### Bug Fixes

* **deps:** Update all non-major go dependencies ([#103](https://github.com/edkadigital/plugin-barman-cloud/issues/103)) ([55258f6](https://github.com/edkadigital/plugin-barman-cloud/commit/55258f69008d1475f65d549d47a6c87485624e28))
* **deps:** Update all non-major go dependencies ([#152](https://github.com/edkadigital/plugin-barman-cloud/issues/152)) ([e77799a](https://github.com/edkadigital/plugin-barman-cloud/commit/e77799af028ba892ed8f3261554682c1b540a7f5))
* **deps:** Update github.com/cloudnative-pg/cloudnative-pg digest to 34ab236 ([#180](https://github.com/edkadigital/plugin-barman-cloud/issues/180)) ([e9e636a](https://github.com/edkadigital/plugin-barman-cloud/commit/e9e636ada08de4a1f6db0a31e2f133e703580394))
* **deps:** Update golang.org/x/net ([#188](https://github.com/edkadigital/plugin-barman-cloud/issues/188)) ([aba0748](https://github.com/edkadigital/plugin-barman-cloud/commit/aba07487891b731b6439429c7b30da21bc260d5f))
* **deps:** Update kubernetes packages to v0.32.1 ([#147](https://github.com/edkadigital/plugin-barman-cloud/issues/147)) ([dbc5550](https://github.com/edkadigital/plugin-barman-cloud/commit/dbc5550c9c503dfb0a6206a244995cdda9d28c1d))
* **deps:** Update kubernetes packages to v0.32.2 ([#172](https://github.com/edkadigital/plugin-barman-cloud/issues/172)) ([bb9658b](https://github.com/edkadigital/plugin-barman-cloud/commit/bb9658b28c95f9b7e1f202dcf2be76bff7756960))
* **deps:** Update module github.com/cloudnative-pg/api to v1 ([#131](https://github.com/edkadigital/plugin-barman-cloud/issues/131)) ([0c8ff74](https://github.com/edkadigital/plugin-barman-cloud/commit/0c8ff7426ff15623deba0c9603ba76dece3cb6a5))
* **deps:** Update module github.com/cloudnative-pg/cnpg-i-machinery to v0.1.2 ([#182](https://github.com/edkadigital/plugin-barman-cloud/issues/182)) ([12cd519](https://github.com/edkadigital/plugin-barman-cloud/commit/12cd5195234ee17ca0b09c2448cc9dc50c614149))
* **deps:** Update module google.golang.org/grpc to v1.71.0 ([#187](https://github.com/edkadigital/plugin-barman-cloud/issues/187)) ([e1f1660](https://github.com/edkadigital/plugin-barman-cloud/commit/e1f166023f55fb02d987ac011e3580af1f9d273a))
* **deps:** Update module sigs.k8s.io/kustomize/api to v0.19.0 ([#148](https://github.com/edkadigital/plugin-barman-cloud/issues/148)) ([9ba6351](https://github.com/edkadigital/plugin-barman-cloud/commit/9ba63518f929748f4a422eaa58293c8125b7a5f1))
* **deps:** Use latest commit from CNPG 1.25 branch ([#178](https://github.com/edkadigital/plugin-barman-cloud/issues/178)) ([dfbeaf8](https://github.com/edkadigital/plugin-barman-cloud/commit/dfbeaf802ec98357fdbb92b5fcefc38a29939cfe))

## 0.1.0 (2024-12-12)


### Features

* Add `liveness` and `readiness` probe support ([#69](https://github.com/edkadigital/plugin-barman-cloud/issues/69)) ([5fd9449](https://github.com/edkadigital/plugin-barman-cloud/commit/5fd9449b27394756e0baf76b1356900850f687a6))
* Additional environment variables ([#81](https://github.com/edkadigital/plugin-barman-cloud/issues/81)) ([be40375](https://github.com/edkadigital/plugin-barman-cloud/commit/be4037529c44858278dd80e3eb32f39f3f68c5c6))
* Backup method ([#20](https://github.com/edkadigital/plugin-barman-cloud/issues/20)) ([9fa1c0b](https://github.com/edkadigital/plugin-barman-cloud/commit/9fa1c0beab4882af3f4c737d049b5bafcf7e28a6))
* Grant permissions to read secrets ([#25](https://github.com/edkadigital/plugin-barman-cloud/issues/25)) ([76383a3](https://github.com/edkadigital/plugin-barman-cloud/commit/76383a30afd3bd829f01936dc3dfc81f1d189d2d))
* Operator plugin and manifests ([#18](https://github.com/edkadigital/plugin-barman-cloud/issues/18)) ([dd6548c](https://github.com/edkadigital/plugin-barman-cloud/commit/dd6548c4a26031324975d97aee345e4e6a2e7efa))
* Separate recovery and cluster object store ([#76](https://github.com/edkadigital/plugin-barman-cloud/issues/76)) ([e30edd2](https://github.com/edkadigital/plugin-barman-cloud/commit/e30edd2318d76e10fd7af344c0e4326f1e5033ec))
* Separate recovery object store from replica source ([#83](https://github.com/edkadigital/plugin-barman-cloud/issues/83)) ([e4735a2](https://github.com/edkadigital/plugin-barman-cloud/commit/e4735a2f85724cf8493f513658783e5330c3efcf))
* Sidecar injection and loading ([#22](https://github.com/edkadigital/plugin-barman-cloud/issues/22)) ([ea6ee30](https://github.com/edkadigital/plugin-barman-cloud/commit/ea6ee30d2ea30f9e9df22002ce5f5a68fcb37ade))
* Sidecar role and rolebinding ([#23](https://github.com/edkadigital/plugin-barman-cloud/issues/23)) ([2f62d53](https://github.com/edkadigital/plugin-barman-cloud/commit/2f62d539c949f344cb5534b7ffbb90860663a106))
* Restore ([#29](https://github.com/edkadigital/plugin-barman-cloud/issues/29)) ([240077c](https://github.com/edkadigital/plugin-barman-cloud/commit/240077c77192d9572767d7ec76d02e578b94faca))
* Wal-archive and wal-restore methods ([#4](https://github.com/edkadigital/plugin-barman-cloud/issues/4)) ([1c86ff6](https://github.com/edkadigital/plugin-barman-cloud/commit/1c86ff65747b5b348fb1ed2b0e5b0594fd156116))
