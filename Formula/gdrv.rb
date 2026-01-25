class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.2.5"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.5/gdrv-darwin-arm64"
      sha256 "687726214d7c34646d8681b562fe3fed630183908e5f657ac5dd28a82fe2828a"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.5/gdrv-darwin-amd64"
      sha256 "c0ffccac9b756a1a993408fce9e29f793b3e6fcbf8d9970e643ae04fc0d5afc5"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.5/gdrv-linux-arm64"
      sha256 "c5c86f5f8c0d3b85915a1f6051ff5619d99e41ed1ab39eeb640504475da468d0"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.5/gdrv-linux-amd64"
      sha256 "bff083c8198756a590a2a0a433a41a145f88b4c14702d010ac93752fe6503c16"
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
