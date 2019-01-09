---
date: "2018-05-10T16:00:00+02:00"
title: "Usage: Issue and Pull Request templates"
slug: "issue-pull-request-templates"
weight: 15
toc: true
draft: false
menu:
  sidebar:
    parent: "usage"
    name: "Issue and Pull Request templates"
    weight: 15
    identifier: "issue-pull-request-templates"
---

# Issue and Pull Request Templates

For some projects there are a standard list of questions that users need to be asked
for creating an issue, or adding a pull request. Gitea supports adding templates to the
main branch of the repository so that they can autopopulate the form when users are 
creating issues, and pull requests. This will cut down on the initial back and forth
of getting some clarifiying details.

Possible file names for issue templates:

* ISSUE_TEMPLATE.md
* issue_template.md
* .gitea/ISSUE_TEMPLATE.md
* .gitea/issue_template.md
* .github/ISSUE_TEMPLATE.md
* .github/issue_template.md


Possible file names for PR templates:

* PULL_REQUEST_TEMPLATE.md
* pull_request_template.md
* .gitea/PULL_REQUEST_TEMPLATE.md
* .gitea/pull_request_template.md
* .github/PULL_REQUEST_TEMPLATE.md
* .github/pull_request_template.md
