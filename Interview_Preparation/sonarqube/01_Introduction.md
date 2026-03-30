🔹 What is SonarQube?

SonarQube is a code quality and security analysis tool.

👉 It analyzes source code, not running applications.

SonarQube checks:

🐞 Bugs

🔐 Security vulnerabilities

🧹 Code smells

📉 Technical debt

🧪 Test coverage

🔹 Why SonarQube exists (Real-world problem)

Before SonarQube ❌:

Code works but is hard to maintain

Bugs found late in production

Security issues unnoticed

No standard quality rules

“It works on my machine” mindset

After SonarQube ✅:

Enforced code quality standards

Early detection of bugs & vulnerabilities

Clean, readable, maintainable code

Same quality rules for entire team

🔹 What SonarQube does NOT do (IMPORTANT ❗)

SonarQube:

❌ Does NOT scan Docker images

❌ Does NOT scan OS vulnerabilities

❌ Does NOT replace Trivy

👉 That’s why Trivy + SonarQube are used together.

🔹 What exactly SonarQube analyzes
1️⃣ Bugs

Code that may:

Cause crashes

Produce wrong results

Example:

if (a = b) { }

2️⃣ Vulnerabilities

Security issues like:

SQL Injection

Hardcoded passwords

Weak cryptography

3️⃣ Code Smells

Code that:

Works

But is badly written

Example:

Duplicate code

Very long methods

Unused variables

4️⃣ Technical Debt

👉 Time needed to fix poor code.

Example:

“If we clean this code now, it’ll take 10 minutes
If we delay, it’ll take 2 hours later”

🔹 Where SonarQube is used

During CI/CD pipelines

On every:

Commit

Pull request

Merge to main branch

Typical Flow::
Developer → Git Push → Jenkins → SonarQube → Quality Gate → Deploy

🔹 SonarQube vs Trivy (VERY IMPORTANT)

| Feature         | SonarQube | Trivy |
| --------------- | --------- | ----- |
| Code quality    | ✅ Yes    | ❌ No |
| Security (code) | ✅ Yes    | ❌ No |
| Image scan      | ❌ No     | ✅ Yes|
| K8s scan        | ❌ No     | ✅ Yes|
| CI/CD           | ✅ Yes    | ✅ Yes|


🔹 Interview MUST-REMEMBER ⭐

SonarQube is a static code analysis tool used to detect bugs, vulnerabilities, code smells, and technical debt to improve code quality and security.

