# Contributing

When contributing to this project, please first discuss the change you wish to make via issue, email, or any other method with the maintainers before making a change.

Please note we have a [code of conduct](CODE_OF_CONDUCT.md), please follow it in all your interactions with the project.

## Pull Request Process

1. Fork
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull/Merge Request

## Sign Your Work

The sign-off is a simple line at the end of the explanation for a commit. All
commits needs to be signed. Your signature certifies that you wrote the patch or
otherwise have the right to contribute the material. The rules are pretty simple,
if you can certify the below (from [developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.

Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

Then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe.smith@example.com>

If you set your `user.name` and `user.email` git configs, you can sign your
commit automatically with `git commit -s`.

Note: If your git config information is set properly then viewing the
`git log` information for your commit will look something like this:

```
Author: Joe Smith <joe.smith@example.com>
Date:   Thu Feb 2 11:41:15 2018 -0800

    Update README

    Signed-off-by: Joe Smith <joe.smith@example.com>
```

Notice the `Author` and `Signed-off-by` lines match. If they don't
your PR will be rejected by the automated DCO check.

## Issues

Issues are used as the primary method for tracking anything to do with the project.

### Issue Types

There are 4 types of issues (each with their own corresponding [label](#labels)):

- Question: These are support or functionality inquiries that we want to have a record of for
  future reference. Generally these are questions that are too complex or large to store in the
  Slack channel or have particular interest to the community as a whole. Depending on the discussion,
  these can turn into "Feature" or "Bug" issues.
- Proposal: Used for items (like this one) that propose a new ideas or functionality that require
  a larger community discussion. This allows for feedback from others in the community before a
  feature is actually developed. This is not needed for small additions. Final word on whether or
  not a feature needs a proposal is up to the core maintainers. All issues that are proposals should
  both have a label and an issue title of "Proposal: [the rest of the title]." A proposal can become
  a "Feature" and does not require a milestone.
- Features: These track specific feature requests and ideas until they are complete. They can evolve
  from a "Proposal" or can be submitted individually depending on the size.
- Bugs: These track bugs with the code or problems with the documentation (i.e. missing or incomplete)

### Issue Lifecycle

The issue lifecycle is mainly driven by the core maintainers, but is good information for those
contributing to Helm. All issue types follow the same general lifecycle. Differences are noted below.

1. Issue creation
2. Triage
   - The maintainer in charge of triageing will apply the proper labels for the issue. This
     includes labels for priority, type, and metadata (such as "starter"). The only issue
     priority we will be tracking is whether or not the issue is "critical." If additional
     levels are needed in the future, we will add them.
   - (If needed) Clean up the title to succinctly and clearly state the issue. Also ensure
     that proposals are prefaced with "Proposal".
   - Add the issue to the correct milestone. If any questions come up, don't worry about
     adding the issue to a milestone until the questions are answered.
   - We attempt to do this process at least once per work day.
3. Discussion
   - "Feature" and "Bug" issues should be connected to the PR that resolves it.
   - Whoever is working on a "Feature" or "Bug" issue (whether a maintainer or someone from
     the community), should either assign the issue to them self or make a comment in the issue
     saying that they are taking it.
   - "Proposal" and "Question" issues should stay open until resolved or if they have not been
     active for more than 30 days. This will help keep the issue queue to a manageable size and
     reduce noise. Should the issue need to stay open, the `keep open` label can be added.
4. Issue closure
