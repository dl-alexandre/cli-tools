class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.5.3"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.3/gdrv-darwin-arm64.tar.gz"
      sha256 "2af54157470e9344a8b5bbb2c307c4bff083f349757145c38e12468deae013ed"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.3/gdrv-darwin-amd64.tar.gz"
      sha256 "dd12eed6209400b140d9ca7bf1fb5f8f76d6797ff304bfae3f5b3d7b3b4f0b03"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.3/gdrv-linux-arm64.tar.gz"
      sha256 "fa9a708683241269d8140c834afafe10b8bcefe1b199defb7b0b9d8187ce6a95"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.3/gdrv-linux-amd64.tar.gz"
      sha256 "9edebd59dae3e0b584531a94c5b29587e2fefe4e021311f609ede6a5cc7fe857"
    end
  end

  def install
    bin.install "gdrv"
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
