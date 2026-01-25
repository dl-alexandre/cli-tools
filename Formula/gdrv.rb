class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.3.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.3.0/gdrv-darwin-arm64"
      sha256 "25ffc8f34d9faabb127f406bba81f7fd67da330e9ce722dbd15525b3afae67a2"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.3.0/gdrv-darwin-amd64"
      sha256 "3a707aa680ac2caa63afff06530afc9433cf5a58d54813c831c1e97838118dbb"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.3.0/gdrv-linux-arm64"
      sha256 "a252c2fe2c0402b3656971c60e5714b03b4ffee0617d223f42daf06ffb4b358f"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.3.0/gdrv-linux-amd64"
      sha256 "592536a510de39143313ba5c1824335a2d5f3cff8dc7f7ce0b80cb292499c056"
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
