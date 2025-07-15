# 🎨 Custom Theme Guide for Plural Chat

## Current Available Themes

Your Plural Chat currently supports all 19 built-in ttkbootstrap themes:

- **cerculean** - Clean blue theme
- **cosmo** - Modern flat theme  
- **cyborg** - Dark cyberpunk theme
- **darkly** - Dark bootstrap theme
- **flatly** - Minimal flat theme
- **journal** - Clean newspaper style
- **kn-bootstrap-vapor-dark** - Dark vapor variant
- **litera** - Typography-focused theme
- **lumen** - Bright clean theme
- **minty** - Fresh green theme
- **morph** - Morphing gradient theme
- **pulse** - Purple accent theme
- **sandstone** - Warm earth tones
- **simplex** - Simple clean theme
- **solar** - Solarized theme
- **superhero** - Dark hero theme (default)
- **united** - Professional theme
- **vapor** - 🔥 **VAPOR THEME** - Synthwave/retrowave theme
- **yeti** - Clean white theme

## 🔥 Using the VAPOR Theme

The **vapor** theme is perfect for that synthwave/retrowave aesthetic! To use it:

1. Open Plural Chat
2. Click **Settings** button
3. Select **vapor** from the theme dropdown
4. Theme changes instantly! ✨

## 🎨 Creating Custom Themes

If you found the **TTK Theme Creator** for Bootstrap, here's how to add custom themes:

### Method 1: Using ttkbootstrap Theme Creator
```python
# If you have the theme creator tool
from ttkbootstrap import Style

# Create a new style
style = Style()

# Define your custom theme
custom_theme = {
    'type': 'dark',  # or 'light'
    'colors': {
        'primary': '#ff6b6b',
        'secondary': '#4ecdc4', 
        'success': '#51cf66',
        'info': '#339af0',
        'warning': '#ffd43b',
        'danger': '#ff8787',
        'light': '#f8f9fa',
        'dark': '#343a40'
    }
}

# Register the theme
style.theme_create('my_custom_theme', custom_theme)
```

### Method 2: JSON Theme Files
Create a JSON file with your theme definition:

```json
{
    "name": "my_synthwave_theme",
    "type": "dark",
    "colors": {
        "primary": "#ff0080",
        "secondary": "#00ffff", 
        "success": "#00ff00",
        "info": "#0080ff",
        "warning": "#ffff00",
        "danger": "#ff4040",
        "light": "#2a2a2a",
        "dark": "#1a1a1a"
    }
}
```

### Method 3: Direct Style Creation
```python
import ttkbootstrap as ttk

# Create a new window with custom theme
root = ttk.Window(
    title="Custom Theme",
    themename="vapor",  # Start with vapor as base
    size=(900, 700)
)

# Modify styles programmatically
style = ttk.Style()
style.configure('Custom.TButton', 
                background='#ff6b6b',
                foreground='white',
                borderwidth=0,
                focuscolor='none')
```

## 🚀 Adding Custom Themes to Plural Chat

To add custom themes to your Plural Chat:

1. **Create the theme** using one of the methods above
2. **Add it to the theme list** in `settings_manager.py`:

```python
# In settings_manager.py, update the theme selector
custom_themes = ['my_custom_theme', 'my_synthwave_theme']
all_themes = list(ttk_bs.Style().theme_names()) + custom_themes
self.theme_selector['values'] = sorted(all_themes)
```

3. **Handle theme loading** in `main.py`:

```python
def change_theme(self, theme_name):
    try:
        if theme_name in custom_themes:
            # Load custom theme
            self.load_custom_theme(theme_name)
        else:
            # Use built-in theme
            self.root.style.theme_use(theme_name)
    except Exception as e:
        self.logger.error(f"Theme error: {e}")
```

## 🎯 Pro Tips

1. **Vapor Theme** is already available - just select it in settings!
2. **Dark themes** work best with chat applications
3. **Test your themes** with different UI elements
4. **Save theme preferences** in the database (already implemented)
5. **Use consistent color schemes** for better UX

## 🔥 Vapor Theme Features

The built-in **vapor** theme includes:
- 🌃 Dark synthwave background
- 💙 Cyan/blue accent colors  
- 🎨 Retro-futuristic styling
- 👾 Perfect for plural systems with a cyberpunk aesthetic

Just select it in Settings > Theme and enjoy the synthwave vibes! 🚀