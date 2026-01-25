class Gdrv < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  version "v0.2.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.2/gdrv-darwin-arm64"
      sha256 "d50cf972977fd8edc8f9a492ec2a81bbdd00028df467b21039bba638c497bda2"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.2/gdrv-darwin-amd64"
      sha256 "43c50e2d75f23bfe4fb91ddae88ba6567b48398035440f670f28e98fbff90db3"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.2/gdrv-linux-arm64"
      sha256 "a789359441881be04699e4ab3c66174bb1727dfcaf6e1c057c6a0559a2edbad1"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Google-Drive-CLI/releases/download/v0.2.2/gdrv-linux-amd64"
      sha256 "e7ea2561e8e57c0a7270e6e608f4bf78014b8821bde9cd2bc016b9fd98d6c640"
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
