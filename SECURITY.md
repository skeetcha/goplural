# 🔒 Security Policy

## 🎯 Our Security Commitment

Plural Chat is built with **privacy and security as core principles**. We take security seriously because:

- **Your data is yours** - We don't collect, store, or transmit user data
- **Local-first architecture** - Everything stays on your device
- **Plural community trust** - Systems need safe spaces to communicate
- **Open source transparency** - All code is public and auditable

## 🛡️ Supported Versions

We actively maintain security for:

| Version | Supported          |
| ------- | ------------------ |
| 0.x.x   | ✅ Active support  |

**Always use the latest version** for the best security and features.

## 🚨 Reporting Security Vulnerabilities

**Please DO NOT report security vulnerabilities through public GitHub issues.**

### 📧 Private Reporting

**For security issues, contact us privately:**

1. **Discord DM**: Message `duskfallcrew` on [our Discord](https://discord.gg/HhBSvM9gBY)
2. **GitHub Security**: Use [GitHub's security advisory feature](https://github.com/skeetcha/goplural/security/advisories)
3. **Email**: If you have our email, that works too

### 📋 What to Include

Please include as much of the following as possible:

- **Vulnerability description** - What's the issue?
- **Impact assessment** - What could an attacker do?
- **Reproduction steps** - How to reproduce the issue
- **Affected versions** - Which versions are vulnerable
- **Proposed solution** - If you have ideas for fixing it
- **Your contact info** - So we can follow up

### ⏰ Response Timeline

- **24 hours** - Initial acknowledgment
- **72 hours** - Initial assessment and severity rating
- **1 week** - Detailed investigation and fix timeline
- **2 weeks** - Fix developed and tested (for high severity)
- **Public disclosure** - After fix is released and users have time to update

## 🔐 Security Features

### 🏠 **Local-First Architecture**
- **No cloud storage** - All data stays on your device
- **No user accounts** - No authentication servers to compromise
- **No telemetry** - We don't track usage or collect analytics
- **Offline capable** - Works without internet connection

### 🖼️ **Avatar Security**
- **URL validation** - Only trusted domains allowed
- **File type checking** - Only image files accepted
- **Size limits** - Prevents resource exhaustion
- **Path traversal prevention** - Secure local file storage
- **Automatic compression** - Reduces file size and removes metadata

### 🗄️ **Database Security**
- **SQLite local storage** - No remote database connections
- **Parameterized queries** - SQL injection prevention
- **File permissions** - Restrictive access controls
- **No sensitive data** - PluralKit tokens stored securely

### 🔗 **Network Security**
- **HTTPS only** - All external connections use TLS
- **Trusted domains** - Whitelist approach for external resources
- **Rate limiting** - Prevents abuse of external APIs
- **Timeout handling** - Prevents hanging connections

### 💻 **Code Security**
- **Input validation** - All user inputs sanitized
- **Error handling** - Graceful failure without information leakage
- **Dependency management** - Regular updates of third-party libraries
- **Code review** - All changes reviewed before merging

## 🚫 Attack Surface Analysis

### ✅ **Low Risk Areas**
- **Local database** - No network exposure
- **Chat history** - Stored locally only
- **Member data** - Private to your system
- **Themes** - Static configuration files

### ⚠️ **Medium Risk Areas**
- **Avatar downloads** - External image fetching
- **PluralKit API** - Third-party API integration
- **File imports** - JSON/export file parsing
- **Plugin system** - When implemented, will need sandboxing

### 🔴 **High Risk Areas**
- **Token storage** - PluralKit tokens need secure storage
- **External images** - Avatar URLs from untrusted sources
- **Export parsing** - Malicious export files
- **Future web features** - Any future web integration

## 🛠️ Security Best Practices

### 👥 **For Users**
- **Keep updated** - Always use the latest version
- **Trusted sources** - Only download from official GitHub releases
- **Review imports** - Be cautious with export files from unknown sources
- **Secure your device** - Use device encryption and strong passwords
- **PluralKit tokens** - Don't share your PK token with anyone

### 🔧 **For Developers**
- **Follow secure coding practices** - See [CONTRIBUTING.md](CONTRIBUTING.md)
- **Validate all inputs** - Never trust user-provided data
- **Use parameterized queries** - Prevent SQL injection
- **Review dependencies** - Keep third-party libraries updated
- **Test security features** - Verify validation and sanitization

## 📋 Security Checklist

Before releasing, we verify:

- [ ] **All user inputs validated** - No direct database queries
- [ ] **External URLs whitelisted** - Only trusted domains allowed
- [ ] **File uploads sanitized** - Images processed safely
- [ ] **Error handling complete** - No sensitive info in error messages
- [ ] **Dependencies updated** - All libraries are current
- [ ] **Code reviewed** - Security-focused review completed

## 🔍 Common Security Concerns

### **"Is my data safe?"**
✅ **Yes** - All data stays on your device, we never see it

### **"Can others access my PluralKit token?"**
✅ **Protected** - Tokens are stored securely and never transmitted except to PluralKit

### **"Are avatar downloads safe?"**
✅ **Validated** - Only trusted domains, file type checking, size limits

### **"What about malicious export files?"**
⚠️ **Validated** - We parse carefully, but always review imports from unknown sources

### **"Can plugins access my data?"**
🔮 **Future feature** - Will be sandboxed with explicit permission system

## 🆘 Security Incident Response

**If a security issue is discovered:**

1. **Immediate containment** - Assess and contain the issue
2. **User notification** - Alert users via GitHub and Discord
3. **Emergency patch** - Develop and test fix rapidly
4. **Release update** - Push fix to users immediately
5. **Post-incident review** - Analyze how to prevent similar issues

## 📊 Security Metrics

We track:
- **Vulnerability disclosure time** - How quickly we respond
- **Fix deployment time** - How quickly fixes reach users
- **Security-related issues** - GitHub issues labeled `security`
- **Dependency updates** - Frequency of security updates

## 🔮 Future Security Enhancements

**Planned improvements:**
- **Code signing** - Verify download authenticity
- **Automatic updates** - Security patches delivered automatically
- **Enhanced token encryption** - Stronger protection for stored tokens
- **Plugin sandboxing** - Secure plugin execution environment
- **Security audits** - Regular third-party security reviews

## 📚 Security Resources

**Learn more about security:**
- [OWASP Top 10](https://owasp.org/www-project-top-ten/) - Common web vulnerabilities
- [Python Security Guide](https://python-security.readthedocs.io/) - Python-specific security
- [SQLite Security](https://sqlite.org/security.html) - Database security best practices
- [Secure Coding Practices](https://owasp.org/www-project-secure-coding-practices-quick-reference-guide/) - General secure coding

## 🏆 Security Hall of Fame

**Responsible disclosure contributors:**
- *None yet - be the first!*

**We appreciate:**
- Responsible disclosure of vulnerabilities
- Detailed reports with reproduction steps
- Patience during investigation and fix process
- Suggestions for security improvements

## 💬 Questions?

**Security questions welcome:**
- **Discord**: [https://discord.gg/HhBSvM9gBY](https://discord.gg/HhBSvM9gBY)
- **GitHub Issues**: For non-sensitive security discussions
- **Security Advisory**: For private vulnerability reports

---

## 🙏 Thank You

Security is a community effort. Thank you for helping keep Plural Chat safe for everyone in the plural community!

---

*Security is not a product, but a process. Let's build it together.* 🔒