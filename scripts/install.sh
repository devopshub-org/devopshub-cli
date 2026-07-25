#!/usr/bin/env bash
# ─────────────────────────────────────────────────────────────────────────────
# DevOpsHub CLI — Universal Install Script
# https://github.com/devopshub-org/devopshub-cli
#
# Usage:
#   curl -sSfL https://raw.githubusercontent.com/devopshub-org/devopshub-cli/main/scripts/install.sh | bash
#   curl -sSfL ... | bash -s -- --version v0.1.0
#   curl -sSfL ... | bash -s -- --dir /usr/local/bin
# ─────────────────────────────────────────────────────────────────────────────
set -euo pipefail

REPO="devopshub-org/devopshub-cli"
BINARY="devopshub"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
VERSION="${VERSION:-}"

# ─── Color helpers ────────────────────────────────────────────────────────────
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
BOLD='\033[1m'
RESET='\033[0m'

log_info()    { echo -e "${CYAN}ℹ ${RESET}$*"; }
log_success() { echo -e "${GREEN}✓ ${RESET}$*"; }
log_warn()    { echo -e "${YELLOW}⚠ ${RESET}$*"; }
log_error()   { echo -e "${RED}✗ ${RESET}$*" >&2; exit 1; }

# ─── Parse arguments ─────────────────────────────────────────────────────────
while [[ $# -gt 0 ]]; do
  case "$1" in
    --version) VERSION="$2"; shift 2 ;;
    --dir)     INSTALL_DIR="$2"; shift 2 ;;
    --help)
      echo "Usage: install.sh [--version <tag>] [--dir <install-dir>]"
      exit 0
      ;;
    *) log_error "Unknown argument: $1" ;;
  esac
done

# ─── Detect OS and architecture ───────────────────────────────────────────────
detect_platform() {
  local os arch

  case "$(uname -s)" in
    Linux*)  os="linux"  ;;
    Darwin*) os="darwin" ;;
    MINGW*|MSYS*|CYGWIN*) os="windows" ;;
    *) log_error "Unsupported OS: $(uname -s)" ;;
  esac

  case "$(uname -m)" in
    x86_64|amd64) arch="x86_64"   ;;
    aarch64|arm64) arch="aarch64" ;;
    *) log_error "Unsupported architecture: $(uname -m)" ;;
  esac

  echo "${os}_${arch}"
}

# ─── Resolve latest version if not pinned ────────────────────────────────────
resolve_version() {
  if [[ -z "${VERSION}" ]]; then
    log_info "Resolving latest release..."
    VERSION=$(curl -sSfL "https://api.github.com/repos/${REPO}/releases/latest" \
      | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')
    [[ -z "${VERSION}" ]] && log_error "Could not determine latest version. Use --version to pin."
    log_info "Latest version: ${BOLD}${VERSION}${RESET}"
  fi
}

# ─── Download and install ────────────────────────────────────────────────────
main() {
  echo -e "\n${BOLD}DevOpsHub CLI Installer${RESET}"
  echo -e "────────────────────────\n"

  # Check required tools.
  for cmd in curl tar; do
    command -v "$cmd" &>/dev/null || log_error "Required tool not found: $cmd"
  done

  resolve_version

  local platform
  platform=$(detect_platform)

  local archive_name="${BINARY}_${VERSION#v}_${platform}.tar.gz"
  local download_url="https://github.com/${REPO}/releases/download/${VERSION}/${archive_name}"
  local checksum_url="https://github.com/${REPO}/releases/download/${VERSION}/checksums.txt"
  local tmpdir
  tmpdir=$(mktemp -d)
  trap 'rm -rf "${tmpdir}"' EXIT

  log_info "Downloading ${archive_name}..."
  curl -sSfL "${download_url}" -o "${tmpdir}/${archive_name}"
  curl -sSfL "${checksum_url}" -o "${tmpdir}/checksums.txt"

  log_info "Verifying checksum..."
  (cd "${tmpdir}" && grep "${archive_name}" checksums.txt | sha256sum -c -) \
    || log_error "Checksum verification failed. Aborting installation."

  log_info "Extracting..."
  tar -xzf "${tmpdir}/${archive_name}" -C "${tmpdir}"

  log_info "Installing to ${INSTALL_DIR}..."
  install -d "${INSTALL_DIR}"
  install -m 0755 "${tmpdir}/${BINARY}" "${INSTALL_DIR}/${BINARY}"

  log_success "DevOpsHub CLI ${VERSION} installed to ${INSTALL_DIR}/${BINARY}"
  echo -e "\nRun: ${CYAN}${BINARY} --help${RESET}"
  echo ""
}

main "$@"
