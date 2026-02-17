class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.4.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.2/gdrv-darwin-arm64"
      sha256 "13c921686a7290474f3ee327ef4d575359ab6433c60faa20f4144dd424ec38a0"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.2/gdrv-darwin-amd64"
      sha256 "4c275753413704cf84962a2a919708a9d874198ec69cb0082750cfe5fbf6989a"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.2/gdrv-linux-arm64"
      sha256 "6c652497b01bd43bcd3496386be94320d524693e5e75cb7a0b6f711a854ec031"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.2/gdrv-linux-amd64"
      sha256 "378daaa8d1af172752d6a6f1278fffbeb0a8b46786061e997a6e8b29092ed59f"
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
