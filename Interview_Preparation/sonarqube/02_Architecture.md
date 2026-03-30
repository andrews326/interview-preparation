📊 SonarQube – Lecture 2: Architecture & Components

🔹 High-level SonarQube Architecture

SonarQube has 3 main components:

Developer / CI (Scanner)
        ↓
   SonarQube Server
        ↓
     Database


👉 Remember this flow — interviewers love this.

🔹 1️⃣ SonarQube Scanner (MOST IMPORTANT)
What is Scanner?

A client-side tool

Runs on:

Developer machine

Jenkins

GitHub Actions

Scans the source code

What scanner does:

Reads source code

Applies rules

Sends analysis result to SonarQube server

Examples:

sonar-scanner

mvn sonar:sonar (Maven)

gradle sonar

📌 Scanner never stores data, it only sends data.

🔹 2️⃣ SonarQube Server
What it does:

Receives reports from scanner

Applies Quality Gates

Stores results in DB

Provides Web UI

Server responsibilities:

Rule engine

Issue management

Quality Gate evaluation

Authentication & authorization

📌 This is the brain 🧠 of SonarQube.

🔹 3️⃣ Database (PostgreSQL)
What is stored?

Analysis results

Issues (bugs, code smells, vulnerabilities)

Metrics

Quality Gate status

Users & permissions

📌 Source code is NOT stored, only metadata.

🔹 SonarQube Web UI (Part of Server)

From UI you can:

View issues

See code coverage

Configure rules

Create Quality Gates

Manage users & projects

🔹 Quality Gate (CRITICAL CONCEPT ⭐)
What is Quality Gate?

A set of conditions that code must pass.

Example conditions:

Code coverage ≥ 80%

No CRITICAL vulnerabilities

No blocker bugs

Result:

✅ PASSED → pipeline continues

❌ FAILED → pipeline stops

👉 This is how SonarQube blocks bad code.

🔹 How SonarQube works in CI/CD (Flow)
Git Push
   ↓
Jenkins Build
   ↓
SonarQube Scan
   ↓
Quality Gate
   ↓
Deploy OR Stop

🔹 SonarQube Editions (Basic idea)
| Edition     | Usage             |
| ----------- | ----------------- |
| Community   | Free, basic       |
| Developer   | PR analysis       |
| Enterprise  | Large org         |
| Data Center | High availability |


👉 For learning & interviews → Community edition is enough.

🔹 Interview MUST-REMEMBER ⭐

Q: Explain SonarQube architecture
👉 Answer:

SonarQube consists of a Scanner that analyzes code and sends results to the SonarQube Server, which processes the analysis, applies Quality Gates, and stores data in a database while exposing results through a web UI.


🔹 Does SonarQube have a database?
✅ YES, SonarQube uses a database.
👉 Database used:

PostgreSQL (recommended & production-ready)

📌 Notes:

Earlier versions supported MySQL, Oracle

Now PostgreSQL is the standard

🔹 What data does SonarQube store in the database?
❌ What it does NOT store

SonarQube does NOT store:

Your source code files

.go, .java, .js code

Docker images

👉 This is VERY IMPORTANT.

✅ What it DOES store (key part)

SonarQube stores analysis results and metadata, not code.

🔹 1️⃣ Issues found during scan

For each scan, SonarQube stores:

Bugs

Vulnerabilities

Code smells

Example stored data:

File name

Line number

Issue type

Severity (BLOCKER / CRITICAL / MAJOR)

Message

📌 Example:

“Null pointer possible at OrderService.go:42”

🔹 2️⃣ Metrics & measurements

SonarQube stores:

Code coverage %

Number of bugs

Cyclomatic complexity

Duplicated code %

Lines of code

📌 Used for dashboards & trends.

🔹 3️⃣ Quality Gate results

Stored data:

Which Quality Gate was applied

Pass or Fail status

Which condition failed

Example:

Coverage < 80% → FAIL

🔹 4️⃣ Project & branch info

Project key

Branch name (main, dev)

Pull request analysis data

🔹 5️⃣ User & permission data

Users

Roles

Tokens

Access control

🔹 6️⃣ Historical analysis (VERY IMPORTANT)

SonarQube keeps:

Previous scans

Trend over time

Example:

Last week: 20 bugs

Today: 5 bugs

🔹 Simple table (easy to remember)
| Stored in DB        | Yes / No |
| ------------------- | -------- |
| Source code         | ❌ No     |
| Bug list            | ✅ Yes    |
| Vulnerabilities     | ✅ Yes    |
| Coverage %          | ✅ Yes    |
| Quality Gate result | ✅ Yes    |
| Scan history        | ✅ Yes    |
| Users & tokens      | ✅ Yes    |


🔹 Why SonarQube must NOT store source code

Security reasons

Storage overhead

Source code stays in Git

👉 SonarQube only references file paths & line numbers.

🔹 Interview-ready answer ⭐

Q: Does SonarQube store source code in DB?
👉 Answer:

No. SonarQube stores only analysis results, metrics, issues, Quality Gate outcomes, and historical trends in its PostgreSQL database, not the source code itself.

🔹 One-line memory trick

Git stores code, SonarQube stores code quality results.


🔹 Interview-ready answer ⭐⭐⭐

Q: How does SonarQube know which code to analyze?
👉 Answer:

SonarQube does not pull code from GitHub. Jenkins checks out the repository and runs the Sonar Scanner on the workspace. The scanner analyzes the local source code and sends the analysis results to the SonarQube server.
