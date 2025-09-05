#!/bin/bash

# Array of extensions
extensions=(
    "golang.Go"
    "GitHub.copilot"
    "GitHub.copilot-chat"
    "VisualStudioExptTeam.vscodeintellicode"
    "ms-vscode.makefile-tools"
    "usernamehw.errorlens"
    "esbenp.prettier-vscode"
    "HookyQR.beautify"
    "aaron-bond.better-comments"
    "eamodio.gitlens"
    "donjayamanne.githistory"
    "mhutchie.git-graph"
    "ckolkman.vscode-postgres"
    "cweijan.vscode-redis-client"
    "rangav.vscode-thunder-client"
    "redhat.vscode-yaml"
    "christian-kohler.path-intellisense"
    "formulahendry.code-runner"
    "PKief.material-icon-theme"
    "vscode-icons-team.vscode-icons"
    "CoenraadS.bracket-pair-colorizer-2"
    "streetsidesoftware.code-spell-checker"
    "IBM.output-colorizer"
    "wmaurer.change-case"
    "ms-azuretools.vscode-docker"
)

# Install each extension
for ext in "${extensions[@]}"; do
    echo "Installing $ext..."
    code --install-extension "$ext" --force
    sleep 1  # Small delay between installations
done

echo "All extensions installed/updated successfully!"