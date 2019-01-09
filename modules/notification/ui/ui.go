// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ui

import (
	"code.gitea.io/git"
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/notification/base"
)

type (
	notificationService struct {
		issueQueue chan issueNotificationOpts
	}

	issueNotificationOpts struct {
		issue                *models.Issue
		notificationAuthorID int64
	}
)

var (
	_ base.Notifier = &notificationService{}
)

// NewNotifier create a new notificationService notifier
func NewNotifier() base.Notifier {
	return &notificationService{
		issueQueue: make(chan issueNotificationOpts, 100),
	}
}

func (ns *notificationService) Run() {
	for {
		select {
		case opts := <-ns.issueQueue:
			if err := models.CreateOrUpdateIssueNotifications(opts.issue, opts.notificationAuthorID); err != nil {
				log.Error(4, "Was unable to create issue notification: %v", err)
			}
		}
	}
}

func (ns *notificationService) NotifyCreateIssueComment(doer *models.User, repo *models.Repository,
	issue *models.Issue, comment *models.Comment) {
	ns.issueQueue <- issueNotificationOpts{
		issue,
		doer.ID,
	}
}

func (ns *notificationService) NotifyNewIssue(issue *models.Issue) {
	ns.issueQueue <- issueNotificationOpts{
		issue,
		issue.Poster.ID,
	}
}

func (ns *notificationService) NotifyIssueChangeStatus(doer *models.User, issue *models.Issue, isClosed bool) {
	ns.issueQueue <- issueNotificationOpts{
		issue,
		doer.ID,
	}
}

func (ns *notificationService) NotifyMergePullRequest(pr *models.PullRequest, doer *models.User, gitRepo *git.Repository) {
	ns.issueQueue <- issueNotificationOpts{
		pr.Issue,
		doer.ID,
	}
}

func (ns *notificationService) NotifyNewPullRequest(pr *models.PullRequest) {
	ns.issueQueue <- issueNotificationOpts{
		pr.Issue,
		pr.Issue.PosterID,
	}
}

func (ns *notificationService) NotifyPullRequestReview(pr *models.PullRequest, r *models.Review, c *models.Comment) {
	ns.issueQueue <- issueNotificationOpts{
		pr.Issue,
		r.Reviewer.ID,
	}
}

func (ns *notificationService) NotifyUpdateComment(doer *models.User, c *models.Comment, oldContent string) {
}

func (ns *notificationService) NotifyDeleteComment(doer *models.User, c *models.Comment) {
}

func (ns *notificationService) NotifyDeleteRepository(doer *models.User, repo *models.Repository) {
}

func (ns *notificationService) NotifyForkRepository(doer *models.User, oldRepo, repo *models.Repository) {
}

func (ns *notificationService) NotifyNewRelease(rel *models.Release) {
}

func (ns *notificationService) NotifyUpdateRelease(doer *models.User, rel *models.Release) {
}

func (ns *notificationService) NotifyDeleteRelease(doer *models.User, rel *models.Release) {
}

func (ns *notificationService) NotifyIssueChangeMilestone(doer *models.User, issue *models.Issue) {
}

func (ns *notificationService) NotifyIssueChangeContent(doer *models.User, issue *models.Issue, oldContent string) {
}

func (ns *notificationService) NotifyIssueChangeAssignee(doer *models.User, issue *models.Issue, removed bool) {
}

func (ns *notificationService) NotifyIssueClearLabels(doer *models.User, issue *models.Issue) {
}

func (ns *notificationService) NotifyIssueChangeTitle(doer *models.User, issue *models.Issue, oldTitle string) {
}

func (ns *notificationService) NotifyIssueChangeLabels(doer *models.User, issue *models.Issue,
	addedLabels []*models.Label, removedLabels []*models.Label) {
}

func (ns *notificationService) NotifyCreateRepository(doer *models.User, u *models.User, repo *models.Repository) {
}

func (ns *notificationService) NotifyMigrateRepository(doer *models.User, u *models.User, repo *models.Repository) {
}
