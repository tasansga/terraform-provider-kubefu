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
cp -R "$SCRIPT_DIR/schemas" "$TF_DIR/schemas"
mkdir -p "$TF_DIR/out"

export TF_CLI_CONFIG_FILE="$CLI_CONFIG"
export TOFU_CLI_CONFIG_FILE="$CLI_CONFIG"
export TF_IN_AUTOMATION=1
GOOS="$("$GO_CMD" env GOOS)"
case "$GOOS" in
  windows) PATHSEP=";" ;;
  *) PATHSEP=":" ;;
esac
export KUBEFU_SCHEMA_PATHS="$TF_DIR/schemas/ktopic.yaml${PATHSEP}$TF_DIR/schemas/kwidget.yaml"

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
  assert_contains() {
    local file="$1"
    local expected="$2"
    if ! grep -Fq "$expected" "$file"; then
      echo "expected content missing in $file: $expected" >&2
      exit 1
    fi
  }
  mkdir -p "$TF_DIR/out"
  for output_name in namespace_yaml config_map_yaml user_schema_yaml; do
    expected="$("$TF_BIN" output -raw "$output_name")"
    case "$output_name" in
      namespace_yaml) file="$TF_DIR/out/namespace.yaml" ;;
      config_map_yaml) file="$TF_DIR/out/config_map.yaml" ;;
      user_schema_yaml) file="$TF_DIR/out/user_schema.yaml" ;;
      *) echo "unknown output $output_name" >&2; exit 1 ;;
    esac
    printf '%s' "$expected" >"$file"
    if [[ ! -f "$file" ]]; then
      echo "expected output file missing: $file" >&2
      exit 1
    fi
    if [[ "$(cat "$file")" != "$expected" ]]; then
      echo "output mismatch for $output_name" >&2
      exit 1
    fi
  done
  assert_contains "$TF_DIR/out/namespace.yaml" "apiVersion: v1"
  assert_contains "$TF_DIR/out/namespace.yaml" "kind: Namespace"
  assert_contains "$TF_DIR/out/namespace.yaml" "name: kubefu-inttest"
  assert_contains "$TF_DIR/out/config_map.yaml" "kind: ConfigMap"
  assert_contains "$TF_DIR/out/config_map.yaml" "data:"
  assert_contains "$TF_DIR/out/config_map.yaml" "hello: world"
  assert_contains "$TF_DIR/out/config_map.yaml" "namespace: kubefu-inttest"
  assert_contains "$TF_DIR/out/user_schema.yaml" "apiVersion: crd.custom.test/v1beta1"
  assert_contains "$TF_DIR/out/user_schema.yaml" "kind: TstQueue"
  ;;
destroy)
  "$TF_BIN" destroy -auto-approve
  ;;
output)
  "$TF_BIN" output
  ;;
esac
