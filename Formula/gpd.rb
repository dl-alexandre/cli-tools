class Gpd < Formula
  desc "Google Play Developer CLI for managing Play Console operations"
  homepage "https://github.com/dl-alexandre/Google-Play-Developer-CLI"
  license "MIT"
  head "https://github.com/dl-alexandre/Google-Play-Developer-CLI.git", branch: "master"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/gpd"
  end

  test do
    system "#{bin}/gpd", "version"
  end
end
