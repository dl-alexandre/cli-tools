class CliTemplate < Formula
  desc "Production-ready Go CLI template with Kong, Viper, caching, and GoReleaser"
  homepage "https://github.com/dl-alexandre/cli-template"
  version "v1.1.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cli-template/releases/download/v1.1.0/cli-template-darwin-arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_DARWIN_ARM64"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cli-template/releases/download/v1.1.0/cli-template-darwin-amd64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_DARWIN_AMD64"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dl-alexandre/cli-template/releases/download/v1.1.0/cli-template-linux-arm64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_LINUX_ARM64"
    end
    if Hardware::CPU.intel?
      url "https://github.com/dl-alexandre/cli-template/releases/download/v1.1.0/cli-template-linux-amd64.tar.gz"
      sha256 "PLACEHOLDER_SHA256_LINUX_AMD64"
    end
  end

  def install
    bin.install "cli-template"
    # Install shell completions if available
    bash_completion.install "completions/cli-template.bash" if File.exist?("completions/cli-template.bash")
    zsh_completion.install "completions/_cli-template" if File.exist?("completions/_cli-template")
    fish_completion.install "completions/cli-template.fish" if File.exist?("completions/cli-template.fish")
  end

  test do
    system "#{bin}/cli-template", "version"
  end
end
