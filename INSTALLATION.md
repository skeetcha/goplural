# 🚀 Installation Guide

This guide covers installation of Plural Chat on Windows, macOS, and Linux systems.

## 📋 Prerequisites

**All platforms need:**
- Python 3.8 or higher
- Git (for cloning the repository)
- Internet connection (for downloading dependencies)

---

## 🪟 Windows Installation

### Option 1: Using Python from Microsoft Store (Recommended)

1. **Install Python:**
   - Open Microsoft Store
   - Search for "Python 3.12" (or latest version)
   - Click "Install"

2. **Install Git:**
   - Download from [git-scm.com](https://git-scm.com/download/win)
   - Run installer with default settings

3. **Install Plural Chat:**
   ```powershell
   # Open PowerShell or Command Prompt
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip install -e .
   ```

4. **Run the application:**
   ```powershell
   plural-chat
   ```

### Option 2: Using Python.org Download

1. **Install Python:**
   - Download from [python.org](https://www.python.org/downloads/)
   - ⚠️ **IMPORTANT**: Check "Add Python to PATH" during installation
   - Choose "Install for all users" if you want

2. **Install Git:**
   - Download from [git-scm.com](https://git-scm.com/download/win)
   - Run installer with default settings

3. **Install Plural Chat:**
   ```cmd
   # Open Command Prompt or PowerShell
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip install -e .
   ```

4. **Run the application:**
   ```cmd
   plural-chat
   ```

### Option 3: Using Package Manager (Advanced)

**Using Chocolatey:**
```powershell
# Install Chocolatey first: https://chocolatey.org/install
choco install python git
git clone https://github.com/Ktiseos-Nyx/plural_chat.git
cd plural_chat
pip install -e .
plural-chat
```

**Using Scoop:**
```powershell
# Install Scoop first: https://scoop.sh/
scoop install python git
git clone https://github.com/Ktiseos-Nyx/plural_chat.git
cd plural_chat
pip install -e .
plural-chat
```

### Windows Troubleshooting

**"'python' is not recognized":**
- Reinstall Python with "Add to PATH" checked
- Or use `py` instead of `python`

**"'pip' is not recognized":**
- Use `python -m pip` instead of `pip`

**Permission errors:**
- Run PowerShell/Command Prompt as Administrator
- Or install in user directory: `pip install -e . --user`

---

## 🍎 macOS Installation

### Option 1: Using Homebrew (Recommended)

1. **Install Homebrew:**
   ```bash
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```

2. **Install Python and Git:**
   ```bash
   brew install python git
   ```

3. **Install Plural Chat:**
   ```bash
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip3 install -e .
   ```

4. **Run the application:**
   ```bash
   plural-chat
   ```

### Option 2: Using Python.org Download

1. **Install Python:**
   - Download from [python.org](https://www.python.org/downloads/)
   - Run the installer package

2. **Install Git:**
   - Download from [git-scm.com](https://git-scm.com/download/mac)
   - Or install Xcode Command Line Tools: `xcode-select --install`

3. **Install Plural Chat:**
   ```bash
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip3 install -e .
   ```

4. **Run the application:**
   ```bash
   plural-chat
   ```

### Option 3: Using MacPorts

1. **Install MacPorts:**
   - Download from [macports.org](https://www.macports.org/install.php)

2. **Install Python and Git:**
   ```bash
   sudo port install python312 git
   sudo port select --set python3 python312
   ```

3. **Install Plural Chat:**
   ```bash
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip3 install -e .
   ```

4. **Run the application:**
   ```bash
   plural-chat
   ```

### macOS Troubleshooting

**"command not found: python":**
- Use `python3` instead of `python`
- Or create alias: `alias python=python3`

**Permission denied:**
- Use `pip3 install -e . --user` for user installation

**Tkinter missing:**
- Install Python from python.org (includes tkinter)
- Or: `brew install python-tk`

---

## 🐧 Linux Installation

### Ubuntu/Debian

1. **Update package list:**
   ```bash
   sudo apt update
   ```

2. **Install Python and Git:**
   ```bash
   sudo apt install python3 python3-pip python3-venv git python3-tk
   ```

3. **Install Plural Chat:**
   ```bash
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip3 install -e .
   ```

4. **Run the application:**
   ```bash
   plural-chat
   ```

### Fedora/CentOS/RHEL

1. **Install Python and Git:**
   ```bash
   # Fedora
   sudo dnf install python3 python3-pip git python3-tkinter
   
   # CentOS/RHEL
   sudo yum install python3 python3-pip git tkinter
   ```

2. **Install Plural Chat:**
   ```bash
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip3 install -e .
   ```

3. **Run the application:**
   ```bash
   plural-chat
   ```

### Arch Linux

1. **Install Python and Git:**
   ```bash
   sudo pacman -S python python-pip git tk
   ```

2. **Install Plural Chat:**
   ```bash
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip install -e .
   ```

3. **Run the application:**
   ```bash
   plural-chat
   ```

### openSUSE

1. **Install Python and Git:**
   ```bash
   sudo zypper install python3 python3-pip git python3-tk
   ```

2. **Install Plural Chat:**
   ```bash
   git clone https://github.com/Ktiseos-Nyx/plural_chat.git
   cd plural_chat
   pip3 install -e .
   ```

3. **Run the application:**
   ```bash
   plural-chat
   ```

### Linux Troubleshooting

**"No module named '_tkinter'":**
- Install tkinter: `sudo apt install python3-tk` (Ubuntu/Debian)
- Or: `sudo dnf install python3-tkinter` (Fedora)

**"Permission denied":**
- Use `pip3 install -e . --user` for user installation
- Or use virtual environment (see below)

**"python3-distutils not found":**
- Install distutils: `sudo apt install python3-distutils`

---

## 🔧 Optional Enhancements

### Virtual Environment (Recommended)

Using a virtual environment keeps your system Python clean:

```bash
# Create virtual environment
python3 -m venv venv

# Activate it
# Windows:
venv\Scripts\activate
# macOS/Linux:
source venv/bin/activate

# Install Plural Chat
pip install -e .

# Run the application
plural-chat

# Deactivate when done
deactivate
```

### Aria2 for Fast Downloads (Optional)

For ultra-fast avatar downloads:

**Windows:**
```powershell
# Using Chocolatey
choco install aria2

# Using Scoop
scoop install aria2
```

**macOS:**
```bash
brew install aria2
```

**Linux:**
```bash
# Ubuntu/Debian
sudo apt install aria2

# Fedora
sudo dnf install aria2

# Arch
sudo pacman -S aria2
```

### Desktop Shortcut (Optional)

**Windows:**
1. Right-click on Desktop → New → Shortcut
2. Target: `C:\path\to\python.exe -m plural_chat`
3. Name: "Plural Chat"

**macOS:**
1. Open Automator → New Document → Application
2. Add "Run Shell Script" action
3. Script: `/usr/local/bin/plural-chat`
4. Save as "Plural Chat.app"

**Linux:**
Create `~/.local/share/applications/plural-chat.desktop`:
```ini
[Desktop Entry]
Name=Plural Chat
Exec=plural-chat
Icon=applications-chat
Type=Application
Categories=Network;Chat;
```

---

## 🆘 Common Issues

### "ModuleNotFoundError: No module named 'tkinter'"

**Solution:** Install tkinter for your system:
- **Windows**: Usually included with Python
- **macOS**: `brew install python-tk` or use python.org installer
- **Linux**: `sudo apt install python3-tk` (Ubuntu/Debian)

### "Permission denied" errors

**Solution:** Use user installation:
```bash
pip install -e . --user
```

### "Python not found" or "pip not found"

**Solution:** 
- **Windows**: Reinstall Python with "Add to PATH" checked
- **macOS/Linux**: Use `python3` and `pip3` instead

### UI looks broken or themes don't work

**Solution:**
1. Update to latest Python version
2. Reinstall dependencies: `pip install -r requirements.txt --upgrade`
3. Check that tkinter is properly installed

### App won't start

**Solution:**
1. Check Python version: `python --version` (should be 3.8+)
2. Verify installation: `pip show Plural-Chat`
3. Try running directly: `python main.py`

---

## 🔄 Updating

To update Plural Chat to the latest version:

```bash
cd plural_chat
git pull origin main
pip install -e . --upgrade
```

---

## 🗑️ Uninstalling

To remove Plural Chat:

```bash
pip uninstall Plural-Chat
rm -rf plural_chat  # Remove the directory
```

---

## 💬 Need Help?

If you're still having trouble:

- **GitHub Issues**: [Report installation problems](https://github.com/Ktiseos-Nyx/plural_chat/issues)
- **GitHub Discussions**: [Ask for help](https://github.com/Ktiseos-Nyx/plural_chat/discussions)
- **Discord**: [Join our community](https://discord.gg/HhBSvM9gBY)

---

## 🎉 Success!

Once installed, you should be able to:
- Run `plural-chat` from any terminal/command prompt
- See the Plural Chat window open
- Start adding members and chatting!

**Next steps:**
1. Read the [README](README.md) for usage instructions
2. Check out the [Contributing Guide](CONTRIBUTING.md) if you want to help
3. Join our [Discord](https://discord.gg/HhBSvM9gBY) community!

---

*Made with 💜 by and for the plural community*