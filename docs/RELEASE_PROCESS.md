# Release Process

Use this process for public releases.

## Prepare

1. Ensure `main` is clean.
2. Run:

   ```bash
   python3 scripts/validate_skills.py
   ```

3. Update [CHANGELOG.md](../CHANGELOG.md):
   - Move relevant `Unreleased` entries under the new version.
   - Add the release date.
   - Leave an empty `Unreleased` section at the top.
4. Update [VERSION](../VERSION).

## Tag

Create an annotated tag matching `VERSION`:

```bash
git tag -a "v$(cat VERSION)" -m "Release v$(cat VERSION)"
git push origin main --follow-tags
```

## GitHub Release

The `release.yml` workflow runs on `v*` tags. It validates the Skills, checks that the tag matches `VERSION`, builds archives, and creates a GitHub Release.

Release notes should summarize:

- Added, changed, fixed, or removed Skills.
- Compatibility or migration notes.
- Validation evidence.
- Known limitations.

## After Release

1. Confirm the GitHub Release exists and includes archives.
2. Open a pull request that starts the next `Unreleased` changelog section if needed.
3. Close or update any release tracking issues.
