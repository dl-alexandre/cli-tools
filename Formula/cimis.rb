class Cimis < Formula
  desc "CIMIS time-series database CLI tool with streaming and querying support"
  homepage "https://github.com/dl-alexandre/cimis-cli"
  version "v0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.1/cimis-darwin-arm64"
      sha256 "ad8d94c999098034505a6f2b02912c8626b3e20d0cab803c3b42cb440054ebdc"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.1/cimis-darwin-amd64"
      sha256 "d8f968f8796ba64cf4b70d55daa3efe32b0c9ff6d466f4443ed0cc627073505e"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.1/cimis-linux-arm64"
      sha256 "ad9ac256907205b00ed2084eabfbc2131e2a53437ba28e694095760e502c3e86"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cimis-cli/releases/download/v0.0.1/cimis-linux-amd64"
      sha256 "5faf5fb400867f56db759e16d1e8d2cd41fa923dcc9181eb37cedfe079835fce"
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
