#!/bin/bash
# Cleanup straggler branches across all CLI repos
# Requires: gh CLI authenticated with repo access

set -e

delete_if_stale() {
    local repo=$1
    local branch=$2
    local reason=$3
    
    echo "[$repo] Checking $branch..."
    pr_state=$(gh api repos/$repo/pulls --paginate --jq ".[] | select(.head.ref == \"$branch\") | .state" 2>/dev/null | head -1 || echo "none")
    
    if [ "$pr_state" != "open" ]; then
        echo "  -> Deleting ($reason)"
        gh api -X DELETE repos/$repo/git/refs/heads/$branch 2>/dev/null || echo "  -> Failed (may be protected)"
    else
        echo "  -> Keeping (open PR)"
    fi
}

echo "=== Cleaning up stale branches ==="

# homebrew-tap stale formula branches (already covered by other script)
echo ""
echo "Run cleanup-stale-branches.sh for homebrew-tap formula branches"

# cimis-cli old workflow branch
delete_if_stale "dl-alexandre/cimis-cli" "codex/act-workflow-fixes" "old branch from Feb 2026, no open PR"

# Google-Drive-CLI api-drift auto branch - keep if active
echo ""
echo "[dl-alexandre/Google-Drive-CLI] Checking auto/api-drift..."
last_update=$(gh api repos/dl-alexandre/Google-Drive-CLI/branches/auto/api-drift --jq '.commit.commit.committer.date' 2>/dev/null || echo "unknown")
echo "  Last update: $last_update"
echo "  -> Keep (automated drift detection branch)"

# Apple-Map-Server-CLI dependabot branch - check if PR is open
echo ""
echo "[dl-alexandre/Apple-Map-Server-CLI] Checking dependabot branch..."
dep_pr=$(gh pr list --repo dl-alexandre/Apple-Map-Server-CLI --state open --json headRefName,number --jq '.[] | select(.headRefName | startswith("dependabot/")) | .number' 2>/dev/null | head -1)
if [ -n "$dep_pr" ]; then
    echo "  -> Keep (open PR #$dep_pr)"
else
    echo "  -> Orphaned dependabot branch - can delete after PR merged"
fi

echo ""
echo "=== Done ==="
