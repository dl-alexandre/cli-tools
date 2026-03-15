class Grokipedia < Formula
  desc "Grokipedia CLI - Wikipedia and knowledge base search tool"
  homepage "https://github.com/dl-alexandre/Grokipedia-CLI"
  version "v0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Grokipedia-CLI/releases/download/v0.0.1/grokipedia-darwin-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Grokipedia-CLI/releases/download/v0.0.1/grokipedia-darwin-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Grokipedia-CLI/releases/download/v0.0.1/grokipedia-linux-arm64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Grokipedia-CLI/releases/download/v0.0.1/grokipedia-linux-amd64.tar.gz"
      sha256 "TO_BE_UPDATED"
    end
  end

  def install
    bin.install "grokipedia"
  end

  test do
    system "#{bin}/grokipedia", "version"
  end
end
