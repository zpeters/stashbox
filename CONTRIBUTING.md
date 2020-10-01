# Contributing

When contributing to this repository, please first discuss the change you wish to make via [Issues](https://github.com/zpeters/stashbox/issues).

Please note we have a [code of conduct](https://github.com/zpeters/stashbox/blob/main/CODE_OF_CONDUCT.md), please follow it in all your interactions with the project.

## Development Process

We follow [Git Flow](https://guides.github.com/introduction/flow/) for changes to this repository.

_GitHub desktop is much easier if you're unfamiliar with using git / GitHub on the command line._

* Fork this repo to create a copy hosted on your github account. The Fork button is in the top right of the page.
    * If you're a collaborator on the repo you can instead just create a branch.
* Clone down your copy of this repo onto your local machine: `git clone <YOUR GITHUB REPO URL>`
* Navigate to the new directory git created. `cd stashbox`
* Create a new branch for your work `git checkout -b <YOUR BRANCH NAME>` Your branch name should be something descriptive of the changes you wish to make, and can include the issue number this change is associated with. Example: `feature/45-update-docs`
* Make your changes. 
* When you're ready to apply your changes, push your changed files to your forked repo
    * `git add <FILENAMES OF CHANGED FILES>`
    * `git commit -m "<YOUR COMMIT MESSAGE>"` Your commit message should be descriptive of the changes you made
    * `git push -u origin HEAD` This will push your changes to the branch you created on your forked repo
* Open a Pull Request to the stashbox repository
    * Navigate to the [stashbox](https://github.com/zpeters/stashbox) repository
    * Click `New pull request`
    * Click `Compare across forks`
    * Select `base repository: zpeters/stashbox`
    * Select `head repository: <YOUR FORKED REPOSITORY>`
    * Select `head branch: <YOUR BRANCH NAME>`
    * Click `Create pull request`

Your pull request will be reviewed and we'll get back to you!
