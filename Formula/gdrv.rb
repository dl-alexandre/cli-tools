class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.4.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.1/gdrv-darwin-arm64"
      sha256 "143a84fc87b9c6876a89ff7e9d742b91f9599b365233597d599443cb92ed8440"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.1/gdrv-darwin-amd64"
      sha256 "99f4621ccebc932a36f245fe44dbb4c592ed69ff6f4bd78463576cb847db5d99"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.1/gdrv-linux-arm64"
      sha256 "8a97c00f246b916bbf7dc3838d27a73c413cd134b622d78ce0f0bdc185d222ef"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.1/gdrv-linux-amd64"
      sha256 "8ec59274e25654c78cbd3c84b74d47b66b52dc2f5f9d9f67e0d48fe41115be4b"
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
