# Adapter


## Clash API
// This script is designed to initiate the sing-box process
// with the necessary configuration to enable clash_api for node switching capabilities.
// According to the documentation available at
// https://sing-box.sagernet.org/configuration/experimental/clash-api/
// it's essential to ensure that clash_api is activated within
// the configuration JSON. If clash_api is not already enabled,
// the script will modify the JSON file under
// the 'experimental' section by appending:
// {
//   "external_controller": "127.0.0.1:9090",
//   "secret": "",
// }
// This adjustment allows the sing-box process to start with support
// for functionalities accessible through clash_api, such as node
// switching and proxy mode toggling. Given the program's design
// requirement to support the clash kernel, this approach enables
// control over functionalities of both kernels exclusively through clash_api.