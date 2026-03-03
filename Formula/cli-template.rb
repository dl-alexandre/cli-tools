class CliTemplate < Formula
  desc "Production-ready Go CLI template with Kong, Viper, caching, and GoReleaser"
  homepage "https://github.com/dl-alexandre/cli-template"
  version "v1.1.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cli-template/releases/download/v1.1.0/cli-template-darwin-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cli-template/releases/download/v1.1.0/cli-template-darwin-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cli-template/releases/download/v1.1.0/cli-template-linux-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cli-template/releases/download/v1.1.0/cli-template-linux-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  def install
    bin.install "cli-template"
  end

  test do
    system "#{bin}/cli-template", "version"
  end
end
