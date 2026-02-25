class Cimis < Formula
  desc "CIMIS time-series database CLI tool with streaming and querying support"
  homepage "https://github.com/dl-alexandre/cimis-cli"
  version "v0.0.3"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.3/cimis-darwin-arm64.tar.gz"
      sha256 "9c5efe3fe64d7d07fe2286c8e8f785e9a0e9ba3ab069c1d151ae7ff39d73c2a0"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.3/cimis-darwin-amd64.tar.gz"
      sha256 "77e17c0c913aef24c397029116afbfd4c095a1a56cd99e48e9106e7e84047281"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.3/cimis-linux-arm64.tar.gz"
      sha256 "63aaca37285a58f491755b6a8309be0192984f1c51364a3916fdb01a180fd65a"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.3/cimis-linux-amd64.tar.gz"
      sha256 "77358fbea8a90f740dd0003c34d35d06c3808d92ea1a430d5aa22f792ee47d52"
    end
  end

  def install
    bin.install "cimis"
  end

  test do
    system "#{bin}/cimis", "version"
  end
end
