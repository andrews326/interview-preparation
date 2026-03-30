🔹 1️⃣ What does SonarQube actually find?

SonarQube mainly finds 4 types of issues:
| Type                 | Meaning                                   |
| -------------------- | ----------------------------------------- |
| **Bug**              | Logic error that may cause wrong behavior |
| **Vulnerability**    | Security weakness                         |
| **Code Smell**       | Poor coding practice                      |
| **Security Hotspot** | Sensitive code that needs review          |


🔹 2️⃣ Bugs 🐞
👉 What is a Bug?

A bug is code that can break at runtime or behave incorrectly.

Example (Go):
var x *int
fmt.Println(*x) // nil pointer dereference


📌 SonarQube detects:

Possible panic

Nil dereference

Infinite loops

Incorrect conditions

🔴 Severity levels:

Blocker

Critical

Major

Minor

Info

🔹 3️⃣ Vulnerabilities 🔐
👉 What is a Vulnerability?

A security flaw that can be exploited by attackers.

Example:
password := "admin123"


SonarQube flags:

Hardcoded credentials

SQL Injection

Command Injection

Weak cryptography

📌 These are real security risks, not style issues.

🔹 4️⃣ Code Smells 😷
👉 What is Code Smell?

Code that works, but is:

Hard to maintain

Hard to read

Poorly designed

Example:
func process(a int, b int, c int, d int, e int) {}


Problems:

Too many parameters

High complexity

Duplicate code

📌 Code smells affect long-term project health.

🔹 5️⃣ Security Hotspots 🔥
👉 What is a Security Hotspot?

Code that may be safe or unsafe, but needs human review.

Example:

http.ListenAndServe(":8080", nil)


SonarQube asks:

“Is this intended to be public?”

📌 It does NOT mark it as vulnerable automatically.

🔹 6️⃣ Rules & Rule Engine

SonarQube works using rules.

Example rules:

Avoid hardcoded passwords

Avoid empty error handling

Avoid high cognitive complexity

Each language has:

Its own rule set

Customizable rules

📌 Rules are evaluated during scan.

🔹 7️⃣ Quality Profiles
👉 What is Quality Profile?

A collection of rules.

Examples:

Default Go profile

Custom company profile

📌 You can:

Enable/disable rules

Change severity

🔹 8️⃣ Quality Gates (VERY IMPORTANT) 🚦
👉 What is Quality Gate?

A pass/fail condition for your project.

Example:

No new bugs
No new vulnerabilities
Coverage > 80%
Duplications < 3%


If gate fails:

Jenkins pipeline fails

Merge blocked

📌 This is how SonarQube enforces quality.

🔹 9️⃣ New Code vs Overall Code

SonarQube focuses on:

New Code (most important)

Not old legacy issues

📌 This is realistic for big companies.


🔹 Simple memory trick 🧠

Bug → App crash 😡

Vulnerability → Hacker 😈

Code smell → Future pain 😩

Hotspot → Review needed 👀