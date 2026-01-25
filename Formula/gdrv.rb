class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.2.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.1/gdrv-darwin-arm64"
      sha256 "61cc368af018ee42949ce184e25d1f68915642e517df19882941d52fde8f2838"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.1/gdrv-darwin-amd64"
      sha256 "6e0841cc73b788629cbdecca79b34d17b53de9dce92e823b892b499d21512fd9"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.1/gdrv-linux-arm64"
      sha256 "1c465d627a4d94750561bb9d26a3394f0e8376d7b36f04fac28d4834f51cafe3"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.1/gdrv-linux-amd64"
      sha256 "45b4e054ed925f62a4694edc61a88d5ef8416a3262b08dafba31d338591b6c12"
    end
  end

  def install
    bin.install "gdrv-darwin-arm64" => "gdrv" if OS.mac? && Hardware::CPU.arm?
    bin.install "gdrv-darwin-amd64" => "gdrv" if OS.mac? && Hardware::CPU.intel?
    bin.install "gdrv-linux-arm64" => "gdrv" if OS.linux? && Hardware::CPU.arm?
    bin.install "gdrv-linux-amd64" => "gdrv" if OS.linux? && Hardware::CPU.intel?
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
