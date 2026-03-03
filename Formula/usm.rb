class Usm < Formula
  desc "UniFi Site Manager CLI for cloud-based site management"
  homepage "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI"
  version "v0.0.3"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v0.0.3/usm-darwin-arm64"
      sha256 "146686a543830f5f147e124a6878b1d47cde3e23356c8fff960056d63425801f"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v0.0.3/usm-darwin-amd64"
      sha256 "e8710efe87e60faedc5c9285d6df5c4a5898a79c9ad2ba8886aae3780fff2031"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v0.0.3/usm-linux-arm64"
      sha256 "ee96a1e08fd8404f7b7cc1888b8aefb276f862a7a8c71b56c2e41a5da369e9de"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/UniFi-Site-Manager-CLI/releases/download/v0.0.3/usm-linux-amd64"
      sha256 "e6c7d03d0bc508df23ff36a802723e4c6ed2cebc58a89ed3c40ec48aa7b6599d"
    end
  end

  def install
    bin.install "usm-darwin-arm64" => "usm" if OS.mac? && Hardware::CPU.arm?
    bin.install "usm-darwin-amd64" => "usm" if OS.mac? && Hardware::CPU.intel?
    bin.install "usm-linux-arm64" => "usm" if OS.linux? && Hardware::CPU.arm?
    bin.install "usm-linux-amd64" => "usm" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/usm", "version"
  end
end
