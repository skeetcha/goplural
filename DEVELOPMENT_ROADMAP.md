# 🗺️ Plural Chat Development Roadmap

## 🎯 **Project Vision**
Transform Plural Chat from a basic chat app into the **premier free, open-source plural communication tool** with premium features at $0 cost.

**Core Philosophy:** Everything Discord charges for, we do for free. Everything PluralKit lacks, we provide.

**🤝 Non-Exclusive Mission:** Plural Chat is **not a replacement** for PluralKit, Simply Plural, or any existing tools. We're here for **everyone** - whether you use other tools or not. Our goal is to **complement and enhance** the plural ecosystem, providing additional options and features that support **anyone who requires them**. The plural community thrives on choice and accessibility, and we're proud to be part of that diverse toolkit.

**🔓 Open Source Reality:** This is open source software - meaning anyone can use it, period. This isn't a "personal belief" or political choice, it's literally how open source works. The source code is public, the software is free, and users decide what tools work for them. Some people need features that other tools don't provide, some prefer offline/local solutions, some have accessibility requirements, and some just want different approaches. **Software doesn't have opinions - it serves whoever needs it.** If providing accessible tools to people who need them is controversial, that's a community problem, not a software problem.

---

## 🏗️ **Current Architecture (Post-Refactor)**

### ✅ **What's Working (Don't Break This!)**
- **Modular UI System** - ThemeManager + Component architecture
- **Aria2 Avatar Downloads** - Ultra-fast bulk downloading for large systems
- **Subprocess PK Sync** - No more UI freezing during imports
- **ttkbootstrap Themes** - Consistent theming across all dialogs
- **SQLite Backend** - Local data storage with proper schema
- **Security** - Avatar URL validation, secure file handling

### 🔧 **Core Files (Critical - Touch Carefully)**
- `main.py` - Application entry point and main window coordination
- `ui/themes/manager.py` - Theme system (don't break theme switching!)
- `ui/components/member_list.py` - Member list with avatar thumbnails
- `database_manager.py` - All database operations
- `pk_sync_worker.py` - Background PluralKit operations
- `aria2_avatar_downloader.py` - Bulk avatar downloading

---

## 🚀 **Phase 1: Visual Polish & UX** *(Next 2-3 weeks)*

### 🎨 **1.1 Message Enhancements**
**Goal:** Make conversations feel more natural and informative

**Features to Add:**
- [ ] **Message timestamps** - Relative (2m ago) and absolute (3:45 PM)
- [ ] **Member color coding** - Each member gets consistent colors in chat
- [ ] **Avatar improvements** - Hover tooltips showing member info
- [ ] **Message status indicators** - Sent/edited/system messages
- [ ] **Display name vs full name setting** - Toggle between showing display names or full member names in chat and UI
- [ ] **Help section for diary** - Add guidance and tips for using the diary system effectively
- [ ] **Polish About dialog** - Fine-tune Ko-fi button sizing, spacing, and overall layout

**Technical Notes:**
- Store message timestamp in existing SQLite schema
- Use member color from database for consistent styling
- Implement tooltip component extending ttkbootstrap

**Files to Touch:**
- `main.py` - Chat display logic
- `database_manager.py` - Add timestamp queries
- `ui/components/` - New tooltip component

### 🔄 **1.2 Interaction Improvements**
**Goal:** Faster, smoother member switching and message management

**Features to Add:**
- [ ] **Quick member switch** - Hotkeys (Ctrl+1, Ctrl+2, etc.)
- [ ] **Typing indicators** - Show current member selection clearly
- [ ] **Message editing** - Right-click context menu
- [ ] **Draft messages** - Auto-save unfinished messages per member

**Technical Notes:**
- Bind keyboard shortcuts to main window
- Store drafts in temporary storage (not database)
- Use ttkbootstrap context menus

**Files to Touch:**
- `main.py` - Keyboard handling, draft storage
- `ui/components/message_input.py` - New component for message area

---

## 🚀 **Phase 2: System Info & Analytics** *(Weeks 4-5)*

### 📊 **2.1 System Statistics Panel**
**Goal:** Provide insights into system communication patterns

**Features to Add:**
- [ ] **Message count per member** - Real-time statistics
- [ ] **Activity timeline** - When system is most active
- [ ] **Member debut tracking** - First message dates
- [ ] **Fun stats** - Longest messages, conversation streaks
- [ ] **Export reports** - Save stats as text/JSON

**Technical Notes:**
- Create new database views for efficient stat queries
- Use matplotlib or similar for simple charts
- New UI panel in main window or separate dialog

**Files to Create:**
- `ui/components/stats_panel.py` - Statistics display component
- `analytics/stats_calculator.py` - Statistics calculation logic
- `database_manager.py` - Add analytics queries

### 🔍 **2.2 Search & History**
**Goal:** Make old conversations easily discoverable

**Features to Add:**
- [ ] **Full-text search** - Search messages by content
- [ ] **Member filter** - Show only messages from specific members
- [ ] **Date range filter** - Find messages from specific time periods
- [ ] **Search highlighting** - Highlight search terms in results

**Technical Notes:**
- SQLite FTS (Full-Text Search) for message content
- New search UI component
- Efficient pagination for large result sets

---

## 🚀 **Phase 3: Customization & Themes** *(Weeks 6-8)*

### 🎨 **3.1 Advanced Theming**
**Goal:** "Premium Discord features for $0" - Ultimate customization

**Features to Add:**
- [ ] **TTK Creator integration** - Button in Settings to launch `python -m ttkcreator` for visual theme creation
- [ ] **Custom color schemes** - User-defined color palettes
- [ ] **Chat layouts** - Compact, spacious, bubble modes
- [ ] **Background customization** - Images, patterns, gradients
- [ ] **Font customization** - Per-component font settings
- [ ] **Theme sharing** - Export/import theme files

**Technical Notes:**
- TTK Creator button: Launch subprocess with `python -m ttkcreator` command
- Extend current theme system without breaking existing themes
- Store custom themes as JSON files
- Theme preview system
- Auto-refresh theme list after TTK Creator closes (detect new themes)

**Files to Expand:**
- `ui/themes/` - Add custom theme support
- `ui/themes/manager.py` - Theme import/export
- New: `ui/themes/customizer.py` - Theme creation UI

### 🏷️ **3.2 Member Customization**
**Goal:** Let each member express their individuality

**Features to Add:**
- [ ] **Member information editor** - Section to customize existing members beyond basic PK data
- [ ] **System hierarchy support** - "Is Sub/Side" system type choices for larger/complex systems
- [ ] **Custom member badges** - Icons, symbols, flair
- [ ] **Avatar frames** - Decorative borders around avatars
- [ ] **Member-specific fonts** - Different fonts per member
- [ ] **Custom member colors** - Override theme colors per member
- [ ] **Member mood indicators** - Optional mood/status per message

**Technical Notes:**
- Extend member database schema for customization data
- Backwards-compatible with existing members
- Asset management for badges/frames

---

## 🚀 **Phase 4: Power User Features** *(Weeks 9-12)*

### 📁 **4.1 Data Management**
**Goal:** Give users complete control over their data

**Features to Add:**
- [ ] **Advanced export options** - Multiple formats (JSON, TXT, Markdown, HTML)
- [ ] **Conversation backups** - Automated local backups
- [ ] **Data import** - Import from other chat platforms
- [ ] **Selective sync** - Choose which PK data to sync
- [ ] **Offline mode** - Full functionality without internet

### 🔧 **4.2 Automation & Templates**
**Goal:** Streamline repetitive tasks

**Features to Add:**
- [ ] **Message templates** - Quick phrases and responses
- [ ] **Auto-responses** - Simple bot-like responses
- [ ] **Scheduled messages** - Send messages at specific times
- [ ] **Quick actions** - Keyboard shortcuts for common tasks
- [ ] **Bulk operations** - Mass edit/delete messages

---

## 🚀 **Phase 5: Internationalization & Accessibility** *(Future Release)*

### 🌍 **5.1 Localization Framework**
**Goal:** Make Plural Chat accessible to the global plural community

**Features to Add:**
- [ ] **Internationalization (i18n) system** - Using Python gettext or JSON-based approach
- [ ] **Language file extraction** - Extract all UI strings to translatable files
- [ ] **Multi-language support** - Support for major plural community languages
- [ ] **Community translation system** - Easy way for community to contribute translations
- [ ] **RTL language support** - Right-to-left text for Arabic, Hebrew, etc.

**Priority Languages:**
- English (default)
- Spanish - Large plural community
- German - Active plural community  
- French - Growing community
- Japanese - Emerging community
- Portuguese - South American community

**Technical Implementation:**
```python
# String wrapping approach
from gettext import gettext as _

# Instead of: 
ttk.Label(text="Settings")

# Use:
ttk.Label(text=_("Settings"))
```

**File Structure:**
```
locales/
├── en_US/
│   └── messages.json
├── es_ES/
│   └── messages.json
├── de_DE/
│   └── messages.json
└── template.json  # For translators
```

### ♿ **5.2 Accessibility Improvements**
**Goal:** Make Plural Chat usable for systems with disabilities

**Features to Add:**
- [ ] **Screen reader compatibility** - Proper ARIA labels and navigation
- [ ] **High contrast themes** - Better visibility for vision impairments
- [ ] **Keyboard-only navigation** - Full app usable without mouse
- [ ] **Text scaling options** - Support for larger fonts/UI elements
- [ ] **Voice activation** - Optional voice commands for member switching
- [ ] **Dyslexia-friendly fonts** - Optional dyslexia-friendly typography

**Technical Notes:**
- Test with NVDA/JAWS screen readers
- Follow WCAG 2.1 accessibility guidelines
- Implement proper tab order and focus management
- Add alt text for all images and icons

---

## 📝 **Development Notes & Known Issues**

### 🔒 **PluralKit API Limitations**
- **Private Members**: Private members cannot have their avatars downloaded via API
  - This is expected behavior for privacy protection
  - Avatar downloads will fail/skip for private members
  - Only affects avatar downloading, not member data sync
  - **Status**: Working as intended (privacy feature)

### ⚠️ **Minor Issues to Investigate**
- **"FINAL STATUS" errors**: Minor errors reported during sync operations
  - Needs investigation to determine root cause
  - May be related to sync completion reporting
  - **Status**: Under investigation

### 📋 **Documentation Tasks**
- [ ] Add private member limitations to README
- [ ] Improve error messages to distinguish private vs download failed
- [ ] Document expected vs unexpected avatar download failures

---

## 🛡️ **Critical Design Principles**

### 🚫 **Never Break These**
1. **PluralKit compatibility** - Always sync properly with PK
2. **Local-first data** - User owns their data, works offline
3. **Zero cost** - No paid features, no premium tiers
4. **Backwards compatibility** - Don't break existing installations
5. **Security first** - Validate all external data (avatars, imports)

### 📐 **Architecture Guidelines**
1. **Component isolation** - New features as separate components
2. **Database migrations** - Always provide upgrade paths
3. **Error handling** - Graceful degradation, never crash
4. **Testing strategy** - Test critical paths before releases
5. **Performance** - Keep UI responsive, use background workers

### 🎨 **UI/UX Standards**
1. **Theme consistency** - All new UI respects current theme
2. **Accessibility** - Keyboard navigation, screen reader support
3. **Progressive disclosure** - Advanced features don't clutter basic UI
4. **User choice** - Everything customizable, nothing forced
5. **Familiar patterns** - Follow established UI conventions

---

## 🤝 **Contribution Guidelines**

### 🎯 **Good First Issues**
- Add new themes to the theme manager
- Create new message export formats
- Implement new statistics calculations
- Add new badge icons or avatar frames
- Improve error messages and user feedback

### 🧠 **Advanced Features**
- Custom theme creation UI
- Advanced search functionality
- Data import from other platforms
- Performance optimizations
- Security enhancements

### 📋 **Before Contributing**
1. **Check this roadmap** - Make sure your idea fits
2. **File an issue** - Discuss before implementing
3. **Follow architecture** - Use existing patterns
4. **Test thoroughly** - Don't break existing features
5. **Document changes** - Update relevant docs

---

## 🎉 **Success Metrics**

### 📊 **Technical Goals**
- [ ] **Sub-second startup time** - Fast app launch
- [ ] **Handles 1000+ members** - Performance at scale
- [ ] **100% offline capable** - Works without internet
- [ ] **Zero data loss** - Bulletproof data handling
- [ ] **Cross-platform** - Works on Windows/Mac/Linux

### 🎮 **User Experience Goals**
- [ ] **Intuitive for new users** - No manual required
- [ ] **Powerful for experts** - Advanced features available
- [ ] **Customizable everything** - Users make it their own
- [ ] **Faster than Discord** - Quicker member switching
- [ ] **More features than PK** - Beyond basic chat

---

## 📅 **Release Strategy**

### 🔄 **Development Cycle**
- **Weekly mini-releases** - Small feature additions
- **Monthly feature releases** - Complete new functionality  
- **Quarterly major releases** - Architecture improvements

### 🚀 **Version Numbering**
- **1.x.x** - Current stable (post-refactor)
- **2.x.x** - Visual polish & UX complete
- **3.x.x** - Full customization system
- **4.x.x** - Power user features complete

---

## 🚀 **Phase 6: Plugin System** *(Future Release)*

### 🧩 **6.1 Plugin Architecture**
**Goal:** Allow community-contributed features without bloating core app

**Features to Add:**
- [ ] **Plugin API framework** - Safe, sandboxed plugin system
- [ ] **Plugin manager UI** - Install, enable/disable, configure plugins
- [ ] **Plugin marketplace/registry** - Community plugin discovery
- [ ] **Plugin development kit** - Documentation and tools for plugin creators
- [ ] **Plugin sandboxing** - Security isolation for third-party code

**Example Plugin Categories:**
- **Export plugins** - New export formats (Markdown, HTML, PDF, etc.)
- **Social sharing plugins** - Optional sharing to social media (screenshot conversations, system stats)
- **Integration plugins** - Connect to other plural tools and services
- **Theme plugins** - Community-created themes and customizations
- **Analytics plugins** - Advanced statistics and reporting
- **Backup plugins** - Cloud backup providers, sync services

**Technical Notes:**
- Plugin system should be **optional** - core app works fully without any plugins
- Use secure plugin API boundaries - plugins can't access sensitive data without permission
- Plugin installation from local files or community registry
- Plugins can extend UI, add menu items, create new dialogs
- Plugin data storage separate from core app database

**Security Considerations:**
- All plugins run in restricted environment
- User approval required for data access
- No direct database access - only through API
- Plugin code review for community registry
- Easy plugin removal without affecting core app

**Files to Create:**
- `plugin_system/` - Core plugin architecture
- `plugin_system/api.py` - Plugin API interface
- `plugin_system/manager.py` - Plugin loading and management
- `plugin_system/ui/` - Plugin management UI
- `docs/plugin_development.md` - Plugin development guide

---

## 📱 **Phase 7: Mobile Companion App** *(Long-term Future)*

### 🚀 **7.1 Mobile Strategy**
**Goal:** Provide mobile access without competing with existing tools

**Philosophy:**
- **Desktop-first, mobile-companion** - Mobile app syncs with desktop, doesn't replace it
- **Freemium model** - Basic features free, premium features cover development costs
- **$25/year Android publishing** - Way better than ads, keeps app clean
- **iOS later** - $99/year is steeper, but if Android succeeds

**Core Free Features:**
- [ ] **Read-only chat history** - View conversations from desktop
- [ ] **Basic member switching** - Send messages as different members
- [ ] **PluralKit sync** - Same data as desktop app
- [ ] **Simple themes** - Basic theming support
- [ ] **Offline reading** - Cached conversations available offline

**Premium Features ($1.99 one-time purchase):**
- [ ] **Full chat editing** - Edit/delete messages from mobile
- [ ] **File attachments** - Send images, voice notes from mobile
- [ ] **Advanced themes** - Custom themes, plugin support
- [ ] **Cloud sync** - Sync between multiple devices
- [ ] **Voice messages** - Record and send voice notes
- [ ] **Export features** - Generate reports, backups from mobile

**Why This Works:**
- **Not competing with SP** - We're focused on chat/conversation, not full plural management
- **Covers costs** - App store fees, server costs for cloud sync
- **No ads** - Clean experience, no tracking
- **Desktop users get mobile access** - Natural extension of existing workflow

**Technical Approach:**
- **Flutter/React Native** - Cross-platform development
- **Local SQLite + Cloud sync** - Same database structure as desktop
- **API backend** - Secure sync between desktop and mobile
- **Selective sync** - Choose what data to sync to mobile

**Revenue Model:**
- **Freemium with one-time purchase** - Basic functionality always free
- **$1.99 forever** - No subscriptions, no recurring fees, no bullshit
- **Premium features are convenience** - Everything works without paying
- **Transparent pricing** - Buy once, own forever
- **Desktop always free** - Mobile premium helps fund overall development

**Files to Create:**
- `mobile_app/` - Flutter/React Native mobile app
- `sync_server/` - Cloud sync backend (optional)
- `api/` - Mobile API endpoints
- `mobile_app/docs/` - Mobile development documentation

### 🛠️ **7.3 Community Add-On System**
**Goal:** Let community developers create mobile add-ons with fair pricing

**Community Add-On Philosophy:**
- [ ] **$5 maximum price** - One-time purchase, covers development costs only
- [ ] **GitHub distribution** - Load add-ons directly from GitHub repositories
- [ ] **Cost-covering only** - Pricing to cover app store fees, hosting, development time
- [ ] **Community-driven** - Anyone can create and distribute add-ons
- [ ] **No gatekeeping** - Open ecosystem, not curated app store

**Example Add-On Ideas:**
- **Custom export formats** - Export to specific journal apps, blog platforms
- **Social integrations** - Post system stats, share favorite conversations
- **Advanced themes** - Community-created theme packs
- **Analytics dashboards** - Detailed system activity reports
- **Backup providers** - Integration with specific cloud storage services
- **Voice processing** - Advanced voice note features, transcription

**Technical Implementation:**
- **GitHub as marketplace** - Add-ons distributed via GitHub releases
- **Plugin manifest system** - Standard format for add-on metadata
- **In-app GitHub integration** - Browse, install, update add-ons from within app
- **Secure sandboxing** - Add-ons run in isolated environment
- **Easy installation** - QR code or link to install add-ons
- **Version management** - Automatic updates from GitHub

**Pricing Guidelines for Community:**
- **Free add-ons encouraged** - Most should be free, community-driven
- **$1-5 range for paid** - Only to cover legitimate development costs
- **No recurring fees** - One-time purchase maximum
- **Transparent cost breakdown** - Developers explain what the fee covers
- **Pay-what-you-want option** - Suggested price with option to pay more

**Why This Works:**
- **Supports community developers** - Fair compensation for work
- **Keeps costs reasonable** - $5 max prevents exploitation
- **Uses existing infrastructure** - GitHub is free, familiar
- **No app store fees for add-ons** - Distributed directly
- **Encourages innovation** - Low barrier to entry for developers

### 🤝 **7.2 Desktop-Mobile Integration**
**Goal:** Seamless experience across devices

**Features to Add:**
- [ ] **QR code pairing** - Easy device linking
- [ ] **Selective sync** - Choose conversations to sync
- [ ] **Push notifications** - Optional mobile notifications
- [ ] **Handoff support** - Continue conversations across devices
- [ ] **Conflict resolution** - Smart merge when both devices are used

**Technical Notes:**
- End-to-end encryption for cloud sync
- Local-first architecture maintained
- Mobile can work fully offline with synced data
- Desktop remains primary, mobile is convenience

---

**Ready to build something awesome?** 🚀✨

*"Because every system deserves premium features without premium prices."*