class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.6.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.6.0/gdrv-darwin-arm64.tar.gz"
      sha256 "37d941aab04afbd51b876cb10958feec4ec20844c781ea8608248f51d0082e5a"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.6.0/gdrv-darwin-amd64.tar.gz"
      sha256 "1fc8da6973b4847c75ff505d3b589999139d8350be1b87ab3aa646393c47ffe2"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.6.0/gdrv-linux-arm64.tar.gz"
      sha256 "84341163607d752ce483af556930b67bacb135f3eae871046d3b19b852e09293"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.6.0/gdrv-linux-amd64.tar.gz"
      sha256 "da4b1d68dadd36c6f168a3b0223615c49592d195c7c2e772f7ab514e8914228a"
    end
  end

  def install
    bin.install "gdrv"
  end

  test do
    system "#{bin}/gdrv", "version"
  end
end
