class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.4.3"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.3/gdrv-darwin-arm64"
      sha256 "7b076820bce8fd4176d166265709cb7576b4800c140a749ee17bc84c3d1af144"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.3/gdrv-darwin-amd64"
      sha256 "0a7ece0fc22892487f23e48788c20b3c8851132b2637655af40ccd718f60175b"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.3/gdrv-linux-arm64"
      sha256 "4a0de2016c15faea0deeda15d1b6679ce30ef48f8bee9107bf1487e61d862a9a"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.4.3/gdrv-linux-amd64"
      sha256 "0339122865bdb0b549e6854a5c0ed7e659a14457e8da7f9f6974a8a94716184c"
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
