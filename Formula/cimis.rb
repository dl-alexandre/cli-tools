class Cimis < Formula
  desc "CIMIS time-series database CLI tool with streaming and querying support"
  homepage "https://github.com/dl-alexandre/cimis-cli"
  version "v0.0.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.2/cimis-darwin-arm64"
      sha256 "575df414bcb5e538e40349bd068ff3590ce488d4b62cf510dc5f9ec2b1ffa000"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.2/cimis-darwin-amd64"
      sha256 "eb20ed2d2dc3769dda66c919ef852cfa011b3e36c9cb80485b27d7b71c9b531a"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.2/cimis-linux-arm64"
      sha256 "c0d7a29a58ffc4458bf64edb0c5a4e5207ce91fdb75ddd69592baada308ad4ec"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.2/cimis-linux-amd64"
      sha256 "d947c92e1674330ec375d8d81092427467186419d31da57a9a5b092be604504e"
    end
  end

  def install
    bin.install "cimis-darwin-arm64" => "cimis" if OS.mac? && Hardware::CPU.arm?
    bin.install "cimis-darwin-amd64" => "cimis" if OS.mac? && Hardware::CPU.intel?
    bin.install "cimis-linux-arm64" => "cimis" if OS.linux? && Hardware::CPU.arm?
    bin.install "cimis-linux-amd64" => "cimis" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/cimis", "version"
  end
end
