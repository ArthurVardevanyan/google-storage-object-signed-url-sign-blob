default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: local
local:
	go build && mv terraform-provider-scaffolding-framework ~/.terraform.d/plugins/terraform.local/local/scaffolding/1.0.0/linux_amd64/terraform-provider-scaffolding_v1.0.0