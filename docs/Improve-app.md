> ## Documentation Index
> Fetch the complete documentation index at: https://docs.lovable.dev/llms.txt
> Use this file to discover all available pages before exploring further.

# Debug and improve your app

> Describe bugs so Lovable can fix them, investigate root causes in Plan mode, and improve a working app with codebase audits, performance checks, and careful changes.

Something breaking is a normal part of building. This guide covers how to describe bugs so Lovable can fix them, how to investigate persistent problems, and ready-made prompts for deeper work like codebase audits and performance checks.

## Quick workflow

When something breaks, work through these steps in order:

1. **Try to fix.** When an error appears, use **Try to fix** once or twice. Your account includes [10 free fixes](https://docs.lovable.dev/features/projects/chat.md#fix-errors-and-get-unstuck), and each one becomes available again 24 hours after you use it.
2. **Describe the bug clearly.** If the error survives, tell Lovable what is broken, where, what you expected, and what happened. Attach a screenshot or paste the error message when you have one. See [Describe the bug, not your frustration](#describe-the-bug-not-your-frustration).
3. **Switch to Plan mode.** Stop retrying blind fixes. Ask Lovable to find the root cause in [Plan mode](https://docs.lovable.dev/features/plan-mode.md) before changing more code. See [Investigate persistent problems in Plan mode](#investigate-persistent-problems-in-plan-mode).
4. **Revert if fixes have tangled the code.** Restore a working version from [version history](https://docs.lovable.dev/features/projects/history.md), or edit a past message and choose **Revert and resend**. See [Revert and take a different path](#revert-and-take-a-different-path).

Repeated blind fixes tend to pile up code that hides the real problem. Switching to Plan mode after one or two failed attempts is usually faster than trying again.

## Describe the bug, not your frustration

Lovable fixes what you can point at. Say what is broken, where, what you expected, and what happened instead.

> [!WARNING]
>
> Avoid generic and broad prompts:
>
> ```text wrap theme={null}
> Nothing works, fix it!
> ```

> [!TIP]
>
> Be specific about the symptom and the location:
>
> ```text wrap theme={null}
> The screen goes blank when I open the Projects page, and I can no longer make edits. Can you check what happened?
> ```

Three ways to give Lovable more to work with:

* **Attach a screenshot** of the broken state, especially for layout issues.
* **Paste the error.** Lovable reads your app's console logs itself, but pasting an exact error message from your browser's developer tools or your backend logs still helps when the error happens outside the preview, for example on your published site:

  ```text wrap theme={null}
  My published app shows a blank screen. Here's the error from the browser console:
  TypeError: Q9() is undefined at https://example.lovable.app/assets/index-DWQbrtrQQj.js:435
  ```
* **Point at the timeline.** If something used to work, say so, and let Lovable look at what changed:

  ```text wrap theme={null}
  The email notification worked yesterday and no longer sends. Review your recent changes to the notification flow, figure out which one broke it, and explain before fixing.
  ```

## Investigate persistent problems in Plan mode

When a fix does not stick, or the same error keeps returning in new variations, the root cause has not been found. Switch to Plan mode and investigate before building again. Prompts that work:

```text wrap theme={null}
What is the root cause of this build error? Show me the relevant code and explain what is going wrong before proposing a fix.
```

```text wrap theme={null}
What solutions have we already tried for this error? List them so we don't repeat ourselves.
```

```text wrap theme={null}
Explain in simple terms why this error occurs.
```

```text wrap theme={null}
This error keeps coming back. Can we take a different approach to achieve the same goal that avoids the problematic area?
```

Ask "why did this happen?", not just "what do we do now?". A quick fix might silence an error without addressing the logic behind it:

```text wrap theme={null}
You fixed the null pointer error by adding a check, but why was the value null in the first place? Can we address that cause?
```

When a fix in one place coincides with a new problem in another, ask about the connection instead of treating them separately:

```text wrap theme={null}
We fixed the booking form, and now the schedule page is acting up. Could the fix have caused it? What do the two share?
```

And when a single component is broken beyond patching, isolate it. Ask Lovable to build a fresh, minimal version of just that component, confirm it works, and then reintegrate it. Rebuilding one piece is often faster than patching an overly broken one.

## Revert and take a different path

If a series of fixes has tangled the code, rewinding is often faster than untangling. Restore a working state from [version history](https://docs.lovable.dev/features/projects/history.md), or edit a past message and choose **Revert and resend** to take a new direction from that point. After reverting, tell Lovable what you did so it has the context:

```text wrap theme={null}
I reverted the project to before the notifications feature. Let's implement it again, but in smaller steps this time.
```

After solving a hard bug, close the loop so future sessions benefit:

```text wrap theme={null}
Summarize what the issue was and how we fixed it, so we can add it to the project knowledge.
```

## Advanced prompts

Ready-made prompts for deeper work. Run the audits in Plan mode so you get analysis before anything changes.

### Codebase audit

Useful when the project has grown and you suspect structural issues:

```text wrap theme={null}
Perform a comprehensive audit of the codebase and report without changing any code:
- Identify files, components, or logic that are misplaced or could be better organized.
- Check the separation of concerns (data handling vs UI vs state) and point out overly coupled sections.
- Highlight code that is overly complex or not following best practices.
- Present the findings as an ordered list of recommended steps, from most critical to optional.
```

You can then implement the recommendations one at a time, verifying each in the preview.

### Performance audit

Useful when the app works but feels slow:

```text wrap theme={null}
The app works but feels sluggish. Analyze the project for performance bottlenecks and report without changing any code:
- Unnecessary or duplicated network and database calls.
- Components that re-render too often or do heavy work.
- Large assets or bundles that slow down loading.
- Recommend specific improvements, ordered by expected impact.
```

### Fragile changes

When you are about to touch a delicate area, such as authentication or payments, set the guardrails in the same prompt as the task:

```text wrap theme={null}
The next change is in a critical part of the app, so proceed with caution. Examine all related code before making changes, avoid modifying unrelated components, and pause and explain if anything is uncertain.

Task: add Google sign-in on top of the existing email and password login, without breaking either flow.
```

## Community debugging guidebook

These prompt snippets come from the Lovable community. Paste the ones that fit your situation into a prompt, or add them to your project [knowledge](https://docs.lovable.dev/features/knowledge.md) so they apply to every request.

> **Error fixing**
>
> ```text wrap theme={null}
> Fix only the code involved in this error. Trace the error message to its source, make a targeted change, and leave unrelated working code untouched. Before you report the fix as done, confirm it resolves the original problem without introducing new ones.
> ```

> **Code modification approach**
>
> ```text wrap theme={null}
> Change only what this task requires. Keep the existing variable names, patterns, and architecture, and check dependencies before editing so nothing else breaks. Show the change as a minimal diff rather than a rewrite. If you notice improvements outside the task, list them separately instead of implementing them.
> ```

> **Database integration**
>
> ```text wrap theme={null}
> Before you propose a new table or column, review the existing schema and reuse tables and fields that already serve the purpose. Keep any schema change compatible with current queries, plan a migration that preserves existing data, and check foreign keys and constraints before you change anything.
> ```

> **Thorough issue analysis**
>
> ```text wrap theme={null}
> Treat this as a diagnosis, not a quick patch. Gather the error messages, logs, and observed behavior first. Form several hypotheses about the cause, test them one at a time, and document what you ruled out. Propose a fix only once you have identified the root cause, and note any edge cases it affects.
> ```

> **Solution verification**
>
> ```text wrap theme={null}
> Before you call a fix complete, test it against the original issue, check related functionality for side effects, and make sure performance has not regressed. Walk through the edge cases. Only report the fix as confirmed after these checks pass.
> ```

> **Code consistency**
>
> ```text wrap theme={null}
> Match the style, naming conventions, and architectural patterns already in this codebase. Use the same error handling, logging, and testing approaches the project uses rather than introducing new ones.
> ```

> **Progressive enhancement**
>
> ```text wrap theme={null}
> Build new features on the existing architecture instead of introducing a new pattern. Use the extension points the current design already provides, keep existing features working, and explain how the addition fits into the current structure.
> ```

> **Documentation and explanation**
>
> ```text wrap theme={null}
> For every change, explain what you changed, why it was needed, and how it works. State any assumptions or dependencies. Add a code comment only where the logic is not obvious from the code itself.
> ```

> **Technical debt awareness**
>
> ```text wrap theme={null}
> Tell me when a fix adds technical debt. Distinguish a quick fix from a proper solution, recommend which one fits this situation, and note what should be refactored later if we take the shortcut.
> ```

> **Learning and adaptation**
>
> ```text wrap theme={null}
> Apply the patterns and preferences this project has already established, and take my feedback on earlier changes into account. Keep track of past issues and their fixes so we do not repeat them, and ask about the business requirement behind a request when it affects the technical approach.
> ```

> **Preventing duplicate components**
>
> ```text wrap theme={null}
> Before you create a new page, component, or flow, search the codebase for similar functionality. Reuse or extend what exists, and when two features are close, parameterize one component instead of duplicating it.
> ```

> **Dead code elimination**
>
> ```text wrap theme={null}
> When you replace functionality, remove the old implementation instead of commenting it out. Before deleting code, check for imports and references to confirm nothing still uses it. Remove unused imports, orphaned components, and unreachable branches you encounter, and explain why each is safe to delete.
> ```

> **Preserving working features**
>
> ```text wrap theme={null}
> Treat working features as locked. Do not remove or substantially change a functioning component unless I ask for it. When fixing an error in one area, do not make "just in case" changes elsewhere. If a change touches a shared component, confirm every feature that depends on it still works.
> ```

> **Deep problem-solving approach**
>
> ```text wrap theme={null}
> This error has resisted standard fixes, so step back before trying again. Question the current assumption about the cause, consider fundamentally different approaches rather than variations of the same one, and outline at least three options with their trade-offs before recommending one. Check less obvious sources such as environment configuration, external dependencies, and race conditions. Break the problem into parts that can be verified independently, and add logging or state tracing where the cause is still unclear.
> ```

> **Database query verification**
>
> ```text wrap theme={null}
> Before you write a query or change the schema, inspect the current tables, fields, and relationships. Reuse existing queries where one already does the job. If you propose a new table or field, confirm nothing equivalent exists under another name and explain why modifying an existing one is not enough. Note any performance implications of the query.
> ```

> **UI consistency and theming**
>
> ```text wrap theme={null}
> Follow the existing design system and color palette. Before you create a new UI component, study the existing ones and reuse their spacing, typography, interaction patterns, and theming. Take design tokens from the codebase rather than introducing new values, handle hover, active, disabled, and error states consistently, respect the existing responsive behavior, and keep accessibility standards such as color contrast and keyboard navigation.
> ```

> **Systematic debugging approach**
>
> ```text wrap theme={null}
> Debug this methodically rather than by trial and error. Reproduce the exact issue first, then gather console logs, network requests, component state, and error messages. Form hypotheses, test each one, and narrow down the affected components and trigger conditions. Document what you found, and confirm the final fix resolves the issue without regressions elsewhere.
> ```

> **Type safety and data validation**
>
> ```text wrap theme={null}
> Check the type definitions in both the database schema and the TypeScript interfaces before implementing anything, and keep strict type checking without falling back to the any type. Watch for common mismatches: numbers arriving as strings, date parsing, and nullable fields. Keep naming consistent between database columns and interfaces, and test with realistic data, including null and undefined cases.
> ```

> **Data flow management**
>
> ```text wrap theme={null}
> Treat data as one pipeline from the database through the API and application state to the UI, and track how it changes at each stage. Invalidate queries so the UI stays in sync with the database, watch for stale cache data, timing issues, and race conditions, and confirm the final data structure matches what each component expects. Add loading states and error boundaries so the UI stays stable when data is missing.
> ```

> **Performance optimization**
>
> ```text wrap theme={null}
> Review this project for performance problems before they become severe: unnecessary database calls, components that re-render too often, N+1 queries and request waterfalls, long lists without virtualization or pagination, and large bundles or unoptimized images. Prioritize fixes by their effect on load time and responsiveness, and avoid optimizations that do not address a measured bottleneck.
> ```

> **Error management and resilience**
>
> ```text wrap theme={null}
> Handle errors so the app stays usable: wrap risky operations in try/catch, use error boundaries so one failing component does not crash the whole page, and let components degrade gracefully with partial data. Show users clear, non-technical error messages, add retries and fallbacks where they help, and log enough context to debug without exposing private data. Fix the root cause rather than suppressing the symptom.
> ```

> **Component architecture**
>
> ```text wrap theme={null}
> Keep a clear component hierarchy with single-responsibility components and well-defined props. Avoid prop drilling by using context or state management where appropriate, and separate data-handling components from presentational ones. When debugging a component, review the whole tree: prop flow, where state lives, and how event handlers connect. Balance reuse against over-abstraction.
> ```

> **API integration and network management**
>
> For every API call, verify the authentication headers, parameters, and body format, and handle each type of network error explicitly. Keep request payloads, response types, and application state consistently typed. Check CORS configuration in every environment, add retries with exponential backoff for transient failures, respect rate limits, and cache requests where it reduces load. Test both the success path and the failure cases, and document each endpoint's purpose, parameters, and response format.
> ```text wrap theme={null}

> 
> **Extend the existing LMLM Voice Studio project with Lovable Agent Integrations (MCP)**
> ```text wrap theme={null}
>Important: plan and implement, but do not publish or deploy.
>1. Inspect the existing source code, backend, authentication, database, storage, consent flow, and speech-provider adapter before changing anything.
>2. Configure an MCP integration protected by OAuth. Never enable anonymous access.
>4. Add delete_my_voice only if ownership enforcement, explicit confirmation, and deletion auditing can be verified.
>5. Every operation involving user data must enforce authentication, record ownership, and applicable permissions on the backend. Never rely on frontend visibility as access control.
>6. Require verified consent before cloning or generating speech with a custom voice. Record the consent version and timestamp, and protect uploaded voice samples and consent evidence.
>7. Make generation asynchronous: return a job ID and provide a status tool. Do not expose long-running generation directly through an MCP request.
>8. Enforce per-user quotas, concurrency limits, and cost controls for generation. Make retries idempotent so the same request cannot unintentionally create duplicate jobs.
>9. Return only the minimum information needed. Do not return raw voice samples, private consent recordings, credentials, or internal provider secrets.
>10. Keep speech generation disabled with a clear error until a real inference provider is configured. Never fabricate audio or report a job as successful when it failed.
>11. Add or update automated tests for authentication, consent enforcement, cross-user access denial, quota enforcement, idempotency, and provider failures.
>12. Document the MCP tools, authentication, permissions, configuration, and test procedures in the README.
>Start with these narrowly scoped tools:
    * list_my_voices
    * get_voice_consent_status
    * list_my_generation_jobs
    * get_generation_status
    * generate_speech 
    Inspect the app’s supported stack and backend requirements first. Report compatibility blockers rather than silently migrating the architecture. Run available checks and report actual results.
    Do not publish, deploy, expose an MCP endpoint publicly, connect unrelated repositories, or claim any test passed unless it was run and verified.
``

## When to ask for help

If you are stuck despite all of this, ask the docs assistant on any page of these docs, or bring the details to the [Discord community](https://discord.gg/lovable-dev). Use Lovable to gather the specifics first (the error, what was tried, when it started), so others can help quickly.


## Related topics

- [AI features for your app](https://docs.lovable.dev/features/ai.md)
- [Lovable changelog](https://docs.lovable.dev/changelog.md)
- [Welcome to Lovable](https://docs.lovable.dev/introduction/welcome.md)
- [FAQ](https://docs.lovable.dev/introduction/faq.md)
- [Project settings](https://docs.lovable.dev/features/projects/settings.md)
