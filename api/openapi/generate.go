package api

//go:generate rm -rf model restapi
//go:generate swagger generate server --api-package op --model-package model --strict-responders --strict-additional-properties --keep-spec-order --exclude-main
//go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;

