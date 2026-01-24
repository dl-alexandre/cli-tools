class Gdrive < Formula
  desc "Google Drive CLI with full read/write support"
  homepage "https://github.com/dl-alexandre/Google-Drive-CLI"
  license "MIT"
  head "https://github.com/dl-alexandre/Google-Drive-CLI.git", branch: "master"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/gdrive"
  end

  test do
    system "#{bin}/gdrive", "version"
  end
end
