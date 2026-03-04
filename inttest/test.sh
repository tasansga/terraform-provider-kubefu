#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

TEMPFILE="$SCRIPT_DIR/.tempdir"
if [[ -f "$TEMPFILE" ]]; then
  WORKDIR="$(cat "$TEMPFILE")"
  if [[ ! -d "$WORKDIR" ]]; then
    mkdir -p "$WORKDIR"
  fi
else
  WORKDIR="$(mktemp -d)"
  echo "$WORKDIR" >"$TEMPFILE"
fi

echo "using temporary workspace: $WORKDIR"

BIN_DIR="$WORKDIR/provider-bin"
DEV_DIR="$WORKDIR/provider-dev"
TF_DIR="$WORKDIR/terraform"

mkdir -p "$BIN_DIR" "$DEV_DIR" "$TF_DIR"

GO_CMD=${GO_CMD:-go}
"$GO_CMD" build -o "$BIN_DIR/terraform-provider-kubefu" ./cmd/terraform-provider-kubefu
cp "$BIN_DIR/terraform-provider-kubefu" "$DEV_DIR/terraform-provider-kubefu"

CLI_CONFIG="$WORKDIR/terraform.rc"
cat > "$CLI_CONFIG" <<EOF
provider_installation {
  dev_overrides {
    "tasansga/kubefu" = "$DEV_DIR"
  }
  direct {
    exclude = ["tasansga/kubefu"]
  }
}
EOF

cp "$SCRIPT_DIR/main.tf" "$TF_DIR/main.tf"

export TF_CLI_CONFIG_FILE="$CLI_CONFIG"
export TOFU_CLI_CONFIG_FILE="$CLI_CONFIG"
export TF_IN_AUTOMATION=1

TF_BIN="${TF_BIN:-$(command -v terraform || command -v tofu || true)}"
if [[ -z "$TF_BIN" ]]; then
  echo "neither tofu nor terraform is installed" >&2
  exit 1
fi

COMMAND="${1:-apply}"
case "$COMMAND" in
plan|apply|destroy|output) ;;
*)
  echo "subcommand must be one of plan|apply|destroy|output (default apply)" >&2
  exit 1
  ;;
esac

cd "$TF_DIR"

# With dev_overrides, init will attempt registry lookups and can fail.
# Only run init when explicitly requested via TF_FORCE_INIT=1.
if [[ "${TF_FORCE_INIT:-}" == "1" ]]; then
  INIT_ARGS=("-input=false" "-backend=false" "-get=false")
  "$TF_BIN" init "${INIT_ARGS[@]}"
fi

case "$COMMAND" in
plan)
  "$TF_BIN" plan
  ;;
apply)
  "$TF_BIN" apply -input=false -auto-approve
  ;;
destroy)
  "$TF_BIN" destroy -auto-approve
  ;;
output)
  "$TF_BIN" output
  ;;
esac
