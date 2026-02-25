class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.5.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-darwin-arm64.tar.gz"
      sha256 "b77aa9f611237ca3082a1efa6abddca9b2076099a54f60f53ad7335bde96bb09"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-darwin-amd64.tar.gz"
      sha256 "83a0f65e2404cfb0f4f9ec66703563071227dc63ccd4a07bf7ff890d014f8053"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-linux-arm64.tar.gz"
      sha256 "482d97e1e9bbf8bbafc1c6812123a2e90965e8cf16aac913d3f20a69691a7d31"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.5.1/gdrv-linux-amd64.tar.gz"
      sha256 "14ec9b8fa76f045ce5370184d9d4d7f157ada96b716bcce6f7614a257d2b3155"
    end
  end

  def install
    bin.install "gdrv"
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
