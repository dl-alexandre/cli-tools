class Abc < Formula
  desc "CLI for Apple Business Connect API - manage business presence on Apple Maps"
  homepage "https://github.com/dl-alexandre/Apple-Business-Connect-CLI"
  version "v0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Apple-Business-Connect-CLI/releases/download/v0.0.1/abc-darwin-arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_DARWIN_ARM64"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Apple-Business-Connect-CLI/releases/download/v0.0.1/abc-darwin-amd64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_DARWIN_AMD64"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/Apple-Business-Connect-CLI/releases/download/v0.0.1/abc-linux-arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_LINUX_ARM64"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/Apple-Business-Connect-CLI/releases/download/v0.0.1/abc-linux-amd64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_LINUX_AMD64"
    end
  end

  def install
    bin.install "abc"
    # Install shell completions if available
    bash_completion.install "completions/abc.bash" if File.exist?("completions/abc.bash")
    zsh_completion.install "completions/_abc" if File.exist?("completions/_abc")
    fish_completion.install "completions/abc.fish" if File.exist?("completions/abc.fish")
  end

  test do
    system "#{bin}/abc", "version"
  end
end
