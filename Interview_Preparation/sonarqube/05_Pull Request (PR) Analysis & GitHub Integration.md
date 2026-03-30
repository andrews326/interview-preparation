🔁 SonarQube – Lecture 6: Pull Request (PR) Analysis & GitHub Integration

This is very important in real companies.

🔹 1️⃣ What is PR Analysis?

When a developer raises a Pull Request:

SonarQube analyzes only the changed code

Not the entire repository

📌 This is called “New Code Analysis”

🔹 2️⃣ Why PR Analysis is needed?

Without PR analysis:

Bad code enters main branch

Bugs & vulnerabilities go live

With PR analysis:

Code is checked before merge

Quality Gate blocks bad PRs


Developer creates PR
        ↓
GitHub Webhook
        ↓
Jenkins PR Pipeline
        ↓
Sonar Scanner (PR mode)
        ↓
SonarQube analyzes NEW code
        ↓
Quality Gate evaluated
        ↓
PR PASS / FAIL


Developer creates PR
        ↓
GitHub Webhook
        ↓
Jenkins PR Pipeline
        ↓
Sonar Scanner (PR mode)
        ↓
SonarQube analyzes NEW code
        ↓
Quality Gate evaluated
        ↓
PR PASS / FAIL
