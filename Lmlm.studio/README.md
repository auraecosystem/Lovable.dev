# LMLM Voice Studio — execution audit

Date: 2026\-10\-24
Project: `57aec9ea-b59e-4971-829c-e0e83f7de3e3`
Editor: https://lovable\.dev/projects/57aec9ea\-b59e\-4971\-829c\-e0e83f7de3e3
Preview: https://id\-preview–57aec9ea\-b59e\-4971\-829c\-e0e83f7de3e3\.lovable\.app

## What I verified by inspecting project metadata and source files

- Project scaffold reports ready; app is **not published**\.
- Stack is TanStack Start, React, TypeScript, Supabase, Drizzle, Vitest\.
- `package.json` defines `npm run build`, `npm run test`, and `npm run lint`\.
- `src/lib/voice/policy.ts` has consent/version/hash gates, text limits, quota checks, ownership helper, and deletion planning\.
- `src/lib/voice/provider.ts` fails closed if `VOICE_ENGINE_URL` is missing or not HTTPS; it does not fabricate audio\.
- `src/lib/voice/voice.test.ts` contains unit tests for consent gates, ownership logic, audio validation, missing provider configuration, quotas, and deletion planning\.
- `src/server/voice.functions.ts` uses authenticated server functions and signed URLs for audio access\.
- The README is still the generic Lovable starter README, not project\-specific setup documentation\.
- No MCP/agent\-integration configuration or MCP\-specific tools appeared in the inspected file list\.

## Findings that should be fixed before release

### P1 — Generation is synchronous

`generateSpeech` calls the configured speech provider inside the request and returns only after synthesis finishes\. This can time out when called through MCP and is not a durable asynchronous job flow\. Add a queued job/worker or external task runner, return a job ID promptly, and let `get_generation_status` poll status\. Do not expose long\-running synthesis as an MCP call\.

### P1 — Verify RLS and ownership in the actual Supabase project

The migration grants authenticated users SELECT on the app tables and defines owner\-read RLS policies, while privileged writes use the service\-role client\. This pattern can be safe only if every service\-role operation applies explicit ownership checks and storage paths remain private\. Verify deployed policies, bucket privacy, and that no client can write directly around the server functions\. The local migration alone does not prove deployed configuration\.

### P1 — Usage quota and profile limits need concurrency\-safe enforcement

The code checks counts before creating profiles/jobs\. Concurrent requests can race and exceed the configured limits\. Enforce quotas atomically in a database function/transaction or with a per\-user lock/unique time\-bucket strategy; add a per\-user concurrency limit and provider timeout/cost controls\.

### P1 — Validate input shape and provider response limits

The `uploadVoice` input validator checks the name and base64 string length, but the other fields need strict runtime validation too &#40;source enum, consent object schema, filename, description&#41;\. Provider output currently accepts WAV/MP3 based on content type but should also cap output size and validate audio bytes before storing\.

### P2 — Audit writes should not silently fail

The `audit` helper inserts events but ignores returned database errors\. Decide which actions require durable audit evidence and fail/alert if the audit insert fails; avoid storing excessive personal data in details\.

### P2 — Improve consent revocation and deletion integrity

Review all deletion and revocation operations for checked errors and transactional consistency\. Ensure revocation prevents new jobs immediately, and deletion reliably removes every derived output and associated metadata even if one storage operation fails\.

### P2 — Complete setup docs

Replace the generic README with required Node/Bun version, environment variables, Supabase setup/migrations, private storage bucket configuration, auth setup, provider API contract, local commands, test commands, and known limitations\. Never put service\-role keys or provider secrets in client variables\.

### P2 — Add MCP only after core safeguards

When Lovable credits are available, configure protected OAuth only\. Proposed tools: `list_my_voices`, `get_voice_consent_status`, `list_my_generation_jobs`, `get_generation_status`, and `generate_speech` &#40;job creation only&#41;\. Keep all tools authenticated; enforce user ownership and quotas in server code\. Do not expose raw voice samples/consent evidence, and do not add deletion tools until destructive behavior is fully tested\.

## Checks and blockers

- Source inspection: completed for `package.json`, README, provider, policy, voice server functions, consent/audio tests, voice/generate/history UI, auth middleware, and initial SQL migration\.
- Build/test/lint: **not run**; this environment has no checked\-out source tree or project runtime\. The Lovable project tool can inspect source but does not provide a local shell to execute these scripts\.
- Lovable agent edits: **blocked**\. Attempting to send a continuation task returned: “Your workspace is out of credits\.” Add credits or update the plan at https://lovable\.dev/settings/billing, then retry\.
- Publishing/deployment: not performed\.
- Real speech generation: not verified; requires a real HTTPS `VOICE_ENGINE_URL` provider implementing `POST /v1/synthesize` and compatible audio response headers/body\.
- MCP integration: not enabled or published\.

## Next execution commands &#40;once code is available locally&#41;

```sh
bun install --frozen-lockfile
bun run test
bun run lint
bun run build
```

Run these from the project root and record the actual outputs\. If using npm instead, use the lockfile\-compatible package manager; this project currently contains `bun.lock`\.
