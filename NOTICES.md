# Third-Party Notices

This file contains notices and license information for third-party software and assets used in Plural Chat.

## Software Dependencies

### ttkbootstrap
- **License**: MIT License
- **Copyright**: © 2021 Israel Dryer
- **Project**: https://github.com/israel-dryer/ttkbootstrap
- **Usage**: Modern UI themes and styling framework
- **License Text**: 
  ```
  MIT License
  
  Copyright (c) 2021 Israel Dryer
  
  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the "Software"), to deal
  in the Software without restriction, including without limitation the rights
  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  copies of the Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:
  
  The above copyright notice and this permission notice shall be included in all
  copies or substantial portions of the Software.
  
  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
  SOFTWARE.
  ```

### Pillow (PIL Fork)
- **License**: HPND License (Historical Permission Notice and Disclaimer)
- **Copyright**: © 1997-2011 by Secret Labs AB, © 1995-2011 by Fredrik Lundh, © 2010-2024 by Jeffrey A. Clark (Alex) and contributors
- **Project**: https://python-pillow.org/
- **Usage**: Image processing, resizing, and format conversion for avatars
- **License**: HPND - https://github.com/python-pillow/Pillow/blob/main/LICENSE

### Requests
- **License**: Apache License 2.0
- **Copyright**: © 2019 Kenneth Reitz
- **Project**: https://github.com/psf/requests
- **Usage**: HTTP client for PluralKit API communication and avatar downloading
- **License**: Apache 2.0 - https://github.com/psf/requests/blob/main/LICENSE

### Cryptography
- **License**: Apache License 2.0 and BSD License (dual licensed)
- **Copyright**: © Individual contributors
- **Project**: https://github.com/pyca/cryptography
- **Usage**: Secure storage of PluralKit API tokens
- **License**: Apache 2.0/BSD - https://github.com/pyca/cryptography/blob/main/LICENSE

### aiohttp
- **License**: Apache License 2.0
- **Copyright**: © 2013-2024 Nikolay Kim and Andrew Svetlov
- **Project**: https://github.com/aio-libs/aiohttp
- **Usage**: Asynchronous HTTP client for improved API performance
- **License**: Apache 2.0 - https://github.com/aio-libs/aiohttp/blob/master/LICENSE.txt

### aria2p
- **License**: ISC License
- **Copyright**: © 2018-2024 Timothée Mazzucotelli
- **Project**: https://github.com/pawamoy/aria2p
- **Usage**: Python wrapper for aria2 download manager
- **License**: ISC - https://github.com/pawamoy/aria2p/blob/main/LICENSE

## External Tools (Optional Dependencies)

### aria2
- **License**: GPL v2+ with OpenSSL Exception
- **Copyright**: © 2006-2023 Tatsuhiro Tsujikawa and contributors
- **Project**: https://aria2.github.io/
- **Usage**: Ultra-fast parallel avatar downloading (optional enhancement)
- **Installation**: User-installed via package manager (e.g., `brew install aria2`)
- **Note**: Falls back to built-in downloading if aria2 is not available
- **License**: GPL v2+ - https://github.com/aria2/aria2/blob/master/COPYING

## Assets and Media

### Default Avatar Image
- **Source**: Vecteezy
- **License**: Free License with Attribution Required
- **Attribution**: "Default avatar vectors created by Vecteezy - www.vecteezy.com"
- **URL**: https://www.vecteezy.com/free-png/default-avatar
- **Usage**: Fallback avatar for system members without custom avatars
- **File**: `avatars/default_avatar.png`

## APIs and Services

### PluralKit API
- **Service**: PluralKit
- **Website**: https://pluralkit.me/
- **Usage**: System member data synchronization and management
- **Terms**: https://pluralkit.me/terms
- **Privacy**: https://pluralkit.me/privacy
- **Note**: PluralKit is a free service for plural systems. Plural Chat integrates with but is not affiliated with PluralKit.

## Fonts

The application uses system fonts and does not bundle any proprietary typefaces. Font selection includes:
- **Consolas** (Windows/Microsoft) - Used as default monospace font where available
- **Monaco** (macOS/Apple) - Alternative monospace font for macOS systems
- **Courier New** - Cross-platform fallback monospace font
- System default fonts as configured by the user's operating system

## Python Standard Library

This application uses Python's standard library components including:
- **tkinter** - GUI framework (included with Python)
- **sqlite3** - Database management (included with Python)
- **json** - Data serialization (included with Python)
- **datetime** - Date and time handling (included with Python)
- **logging** - Application logging (included with Python)
- **threading** - Concurrent operations (included with Python)
- **subprocess** - External process management (included with Python)
- **asyncio** - Asynchronous programming support (included with Python 3.8+)

---

## Compliance Notes

- All third-party components are used in compliance with their respective licenses
- Original license texts are preserved and referenced where required
- Attribution requirements are met in both the application and this notice file
- No modifications have been made to third-party code that would violate license terms

## Reporting Issues

If you believe any attribution is missing or incorrect, please report it via:
- GitHub Issues: [Your Repository URL]
- Email: [Your Contact Email]

---

**Last Updated**: July 16, 2025
**Plural Chat Version**: 0.1.0