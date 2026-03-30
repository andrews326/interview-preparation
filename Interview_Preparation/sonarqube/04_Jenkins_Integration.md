🔹 3️⃣ End-to-End Flow (VERY IMPORTANT)
🔄 Real production flow

Developer → Git Push
      ↓
GitHub Webhook
      ↓
Jenkins Pipeline Triggered
      ↓
Jenkins pulls code (git clone)
      ↓
Build/Test (Maven / Go / npm)
      ↓
Sonar Scanner runs
      ↓
Results sent to SonarQube
      ↓
Quality Gate evaluated
      ↓
Jenkins PASS / FAIL
