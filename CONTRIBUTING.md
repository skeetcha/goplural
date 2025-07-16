# 🤝 Contributing to Plural Chat

Thank you for your interest in contributing to Plural Chat! This project is built **by and for the plural community**, and we welcome contributions from systems of all sizes and experiences.

## 🌟 Quick Start

1. **Check our [Code of Conduct](CODE_OF_CONDUCT.md)** - We're committed to a safe, inclusive community
2. **Browse the [Development Roadmap](DEVELOPMENT_ROADMAP.md)** - See what we're working on
3. **Join our [Discord](https://discord.gg/HhBSvM9gBY)** - Chat with other contributors
4. **Fork the repo** and start coding!

## 🎯 Ways to Contribute

### 🐛 Bug Reports
Found something broken? Help us fix it!

**Before reporting:**
- Check [existing issues](https://github.com/Ktiseos-Nyx/plural_chat/issues) first
- Try the latest version
- Include steps to reproduce

**Good bug reports include:**
- Clear description of what happened vs. what should happen
- Step-by-step reproduction instructions
- Your OS, Python version, and relevant system info
- Screenshots if relevant
- Error messages/logs if available

### 💡 Feature Requests
Have an idea? We'd love to hear it!

**Before suggesting:**
- Check our [Development Roadmap](DEVELOPMENT_ROADMAP.md)
- Search existing issues for similar requests
- Consider if it fits our [Core Philosophy](#-core-philosophy)

**Good feature requests include:**
- Clear problem statement: "As a plural system, I want..."
- Specific use case: "This would help because..."
- Rough implementation ideas (if you have them)

### 🔧 Code Contributions
Ready to dive in? Here's how:

## 🚀 Development Setup

### Prerequisites
- Python 3.8 or higher
- Git
- A text editor or IDE

### Getting Started
```bash
# Fork the repo on GitHub, then clone your fork
git clone https://github.com/YOUR-USERNAME/plural_chat.git
cd plural_chat

# Create a virtual environment
python -m venv venv

# Activate it
# On Windows:
venv\Scripts\activate
# On macOS/Linux:
source venv/bin/activate

# Install dependencies
pip install -e .

# Run the app to make sure everything works
python main.py
```

## 📋 Contribution Guidelines

### 🎨 Code Style
- **Follow existing patterns** - Look at similar code in the project
- **Use meaningful names** - Variables, functions, classes should be self-documenting
- **Comment sparingly** - Good code explains itself, comments explain why
- **Keep functions small** - One responsibility per function
- **Error handling** - Always handle potential errors gracefully

### 🗂️ Project Structure
```
plural_chat/
├── main.py                 # Main application entry point
├── database_manager.py     # Database operations
├── pluralkit_api.py       # PluralKit integration
├── ui/                    # UI components
│   ├── themes/            # Theme management
│   └── components/        # Reusable UI components
├── dialogs/               # Dialog windows
└── docs/                  # Documentation
```

### 🔄 Development Workflow

1. **Create a branch** for your feature/fix:
   ```bash
   git checkout -b feature/amazing-new-feature
   ```

2. **Make your changes** following our guidelines

3. **Test thoroughly**:
   - Run the app and test your changes
   - Try different scenarios
   - Test with different themes
   - Check that existing features still work

4. **Commit with good messages**:
   ```bash
   git commit -m "Add member quick-switch hotkeys

   - Ctrl+1-9 switches to members 1-9
   - Visual feedback shows current selection
   - Fallback to member dropdown for 10+ members
   
   Fixes #123"
   ```

5. **Push and create a Pull Request**:
   ```bash
   git push origin feature/amazing-new-feature
   ```

### 📝 Pull Request Guidelines

**Good PRs include:**
- Clear title and description
- Link to related issues
- Screenshots/GIFs for UI changes
- Test instructions
- Breaking changes clearly marked

**Before submitting:**
- Test your changes thoroughly
- Update documentation if needed
- Check that existing tests still pass
- Make sure the app still starts and works

## 🎯 Good First Issues

New to the project? Start here:

- **Add new themes** - Create new ttkbootstrap theme configurations
- **Improve error messages** - Make error messages more helpful
- **Add export formats** - Support for new export file formats
- **Documentation** - Improve README, add code comments
- **UI polish** - Small improvements to existing interfaces
- **Bug fixes** - Fix reported issues

Look for issues labeled `good first issue` or `help wanted`.

## 🧠 Advanced Contributions

Ready for bigger challenges?

- **New UI components** - Member list improvements, chat enhancements
- **Database migrations** - Schema changes and upgrade paths
- **Performance optimizations** - Make the app faster and more responsive
- **Security improvements** - Enhance avatar validation, data protection
- **Accessibility features** - Screen reader support, keyboard navigation
- **Testing infrastructure** - Add automated tests

## 🌟 Core Philosophy

When contributing, keep these principles in mind:

### ✅ **What We Want**
- **Local-first** - User owns their data, works offline
- **Privacy-focused** - No tracking, no data collection
- **Accessible** - Works for systems with disabilities
- **Community-driven** - Built by plural folks, for plural folks
- **Free forever** - Core functionality never paywalled
- **Complementary** - Enhances existing tools, doesn't replace them

### ❌ **What We Don't Want**
- **Cloud-dependent features** - App should work fully offline
- **Tracking or analytics** - Respect user privacy completely
- **Paid features** - Core functionality must stay free
- **Controversial takes** - Software serves everyone, no politics
- **Breaking changes** - Don't break existing installations
- **Scope creep** - Stay focused on chat/communication

## 🔧 Technical Guidelines

### 🗄️ Database Changes
- **Always provide migration path** - Don't break existing data
- **Test with real data** - Use sample systems for testing
- **Backwards compatibility** - Old databases should still work

### 🎨 UI Changes
- **Respect all themes** - Test with different themes
- **Keyboard navigation** - Everything should work without mouse
- **Consistent spacing** - Follow existing UI patterns
- **Responsive design** - Work at different window sizes

### 🔒 Security
- **Validate all inputs** - Never trust user data
- **Secure file handling** - Prevent path traversal attacks
- **Avatar URL validation** - Only allow trusted domains
- **SQL injection prevention** - Use parameterized queries

## 🤝 Community

### 💬 Getting Help
- **Discord**: [https://discord.gg/HhBSvM9gBY](https://discord.gg/HhBSvM9gBY)
- **Issues**: GitHub issues for bugs and features
- **Discussions**: GitHub discussions for general questions

### 🌈 Plural-Friendly Environment
- **System privacy respected** - No one has to "prove" their plurality
- **Multiple contributors welcome** - Systems can have multiple members contributing
- **Communication styles vary** - Be patient with different communication preferences
- **No system policing** - We don't dictate how systems should work
- **Inclusive language** - Avoid singlet-centric assumptions

### 📋 Review Process
1. **Maintainer review** - Core team reviews all PRs
2. **Community feedback** - Other contributors can comment
3. **Testing period** - Significant changes get testing time
4. **Documentation updates** - Update relevant docs
5. **Merge and celebrate** - Your contribution is live!

## 🎉 Recognition

Contributors get:
- **GitHub contributor badge** - Automatic recognition
- **Discord contributor role** - Special role in our server
- **Changelog mention** - Credit in release notes
- **Community appreciation** - Thanks from the plural community!

## 📚 Resources

- **Python tkinter docs**: [https://docs.python.org/3/library/tkinter.html](https://docs.python.org/3/library/tkinter.html)
- **ttkbootstrap docs**: [https://ttkbootstrap.readthedocs.io/](https://ttkbootstrap.readthedocs.io/)
- **PluralKit API docs**: [https://pluralkit.me/api/](https://pluralkit.me/api/)
- **SQLite docs**: [https://sqlite.org/docs.html](https://sqlite.org/docs.html)

## ❓ Questions?

**Not sure about something?** Ask! We're here to help:
- Open a [GitHub discussion](https://github.com/Ktiseos-Nyx/plural_chat/discussions)
- Join our [Discord](https://discord.gg/HhBSvM9gBY)
- Create an issue with the `question` label

**Remember**: There are no stupid questions, and we were all beginners once!

---

## 🙏 Thank You!

Every contribution, no matter how small, helps make Plural Chat better for the entire community. Whether you're fixing a typo, adding a feature, or just reporting a bug - **you're making a difference**.

The plural community thrives on mutual support and collaboration. Let's build something amazing together! 🚀

---

*Made with 💜 by and for the plural community*