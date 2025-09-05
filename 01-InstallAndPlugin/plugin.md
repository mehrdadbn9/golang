at terminal:

```bash
# Install or update extensions in smaller batches to avoid memory issues

# Batch 1: Core Go and AI extensions
code --install-extension golang.Go --force
code --install-extension GitHub.copilot --force
code --install-extension GitHub.copilot-chat --force
code --install-extension VisualStudioExptTeam.vscodeintellicode --force
code --install-extension ms-vscode.makefile-tools --force

# Batch 2: Code quality and formatting
code --install-extension usernamehw.errorlens --force
code --install-extension esbenp.prettier-vscode --force
code --install-extension HookyQR.beautify --force
code --install-extension aaron-bond.better-comments --force

# Batch 3: Git and version control
code --install-extension eamodio.gitlens --force
code --install-extension donjayamanne.githistory --force
code --install-extension mhutchie.git-graph --force

# Batch 4: Database and API tools
code --install-extension ckolkman.vscode-postgres --force
code --install-extension cweijan.vscode-redis-client --force
code --install-extension rangav.vscode-thunder-client --force

# Batch 5: Language support and utilities
code --install-extension redhat.vscode-yaml --force
code --install-extension christian-kohler.path-intellisense --force
code --install-extension formulahendry.code-runner --force

# Batch 6: UI and themes
code --install-extension PKief.material-icon-theme --force
code --install-extension vscode-icons-team.vscode-icons --force
code --install-extension CoenraadS.bracket-pair-colorizer-2 --force

# Batch 7: Remaining utilities
code --install-extension streetsidesoftware.code-spell-checker --force
code --install-extension IBM.output-colorizer --force
code --install-extension wmaurer.change-case --force
code --install-extension ms-azuretools.vscode-docker --force
```