class Cimis < Formula
  desc "CIMIS time-series database CLI tool with streaming and querying support"
  homepage "https://github.com/dl-alexandre/cimis-cli"
  version "v0.0.4"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.4/cimis-darwin-arm64.tar.gz"
      sha256 "a857456a87e4e5105fb0f477159a6b74bf4c9b11244d6f6d4db084147a29273f"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.4/cimis-darwin-amd64.tar.gz"
      sha256 "e2add4e97753cd694c267e413446824313e2e5e53f4f3ac48ab307a11b00364c"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.4/cimis-linux-arm64.tar.gz"
      sha256 "c3595eeb12a14aaed71bf874d743a05e70f5dd911e44365bf6d08a242b6406cb"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.4/cimis-linux-amd64.tar.gz"
      sha256 "56f2c401042a98564e87ee8b52af5cc9c82c8fd7e9e3a351525161df6e7b2831"
    end
  end

  def install
    bin.install "cimis"
  end

  test do
    system "#{bin}/cimis", "version"
  end
end
