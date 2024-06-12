# Hello World
## A small project used in our series ("Self-hosting Git with CI/CD, using Gitea, Actions and its Container registry")[https://thehomelabber.com/guides/self-hosted-git-ci-cd]

To use this, clone the repository locally and change the remote origin to your new repository:
```bash
git clone git@github.com:TheHomelabber/self-hosted-git-hello-world.git -o hello-world
cd hello-world
git remote remove git@github.com:TheHomelabber/self-hosted-git-hello-world.git

# Change the origin to your repository.
git remote add origin git@git.example.com:Example/self-hosted-git-hello-world.git
```
