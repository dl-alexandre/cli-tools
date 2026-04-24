#!/bin/bash
# Cleanup stale update-formulas branches in homebrew-tap
# Requires: gh CLI authenticated with repo access

set -e

REPO="dl-alexandre/homebrew-tap"

echo "Fetching stale update-formulas branches..."
branches=$(gh api repos/$REPO/branches --paginate --jq '.[] | select(.name | startswith("update-formulas-")) | .name')

count=$(echo "$branches" | grep -c '^update-formulas-' || echo 0)
echo "Found $count stale branches"

for branch in $branches; do
    # Check PR state
    pr_state=$(gh api repos/$REPO/pulls --paginate --jq ".[] | select(.head.ref == \"$branch\") | .state" | head -1 || echo "none")
    
    if [ "$pr_state" != "open" ]; then
        echo "Deleting: $branch (PR state: ${pr_state:-none})"
        gh api -X DELETE repos/$REPO/git/refs/heads/$branch || echo "  Failed to delete $branch"
    else
        echo "Keeping: $branch (open PR)"
    fi
done

echo "Done."
