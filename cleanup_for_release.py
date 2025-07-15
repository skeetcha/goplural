#!/usr/bin/env python3
"""
Cleanup script to sanitize Plural Chat for GitHub release
Removes personal data while keeping the app structure intact
"""

import os
import sqlite3
import shutil
import glob

def cleanup_databases():
    """Delete databases completely - they'll be recreated fresh on first run"""
    print("🗄️ Deleting databases for fresh start...")
    
    # Delete app.db completely
    if os.path.exists('app.db'):
        os.remove('app.db')
        print("  ✅ app.db deleted (will be created fresh on first run)")
    else:
        print("  ℹ️ No app.db found")
    
    # Delete system.db completely  
    if os.path.exists('system.db'):
        os.remove('system.db')
        print("  ✅ system.db deleted (will be created fresh on first run)")
    else:
        print("  ℹ️ No system.db found")

def cleanup_avatars():
    """Remove personal avatar images but keep default"""
    print("🖼️ Cleaning avatars...")
    
    if os.path.exists('avatars'):
        # Remove all avatar files except default_avatar.png
        avatar_files = glob.glob('avatars/*')
        removed_count = 0
        for avatar_file in avatar_files:
            if os.path.isfile(avatar_file) and not avatar_file.endswith('default_avatar.png'):
                os.remove(avatar_file)
                removed_count += 1
        print(f"  ✅ Removed {removed_count} personal avatar files (kept default)")
    else:
        print("  ℹ️ No avatars directory found")

def cleanup_logs():
    """Remove log files"""
    print("📝 Cleaning logs...")
    
    log_files = glob.glob('*.log') + glob.glob('logs/*.log')
    for log_file in log_files:
        if os.path.exists(log_file):
            os.remove(log_file)
    
    if log_files:
        print(f"  ✅ Removed {len(log_files)} log files")
    else:
        print("  ℹ️ No log files found")

def cleanup_exports():
    """Remove any export files and chat history"""
    print("📤 Cleaning export files and chat history...")
    
    # Look for various file patterns that might contain personal data
    patterns = [
        '*.json',           # Export files
        'exports/*.json',   # Export directory
        'chat_history.*',   # Chat history files
        'messages.*',       # Message files
        'history.*',        # History files
        '*.log',           # Any remaining log files
        '*.txt'            # Text exports
    ]
    
    removed_files = []
    for pattern in patterns:
        files = glob.glob(pattern)
        for file_path in files:
            if os.path.exists(file_path) and os.path.isfile(file_path):
                # Skip our README and important files
                filename = os.path.basename(file_path)
                if filename not in ['README.md', 'requirements.txt', 'LICENSE']:
                    os.remove(file_path)
                    removed_files.append(file_path)
    
    if removed_files:
        print(f"  ✅ Removed {len(removed_files)} export/history files")
    else:
        print("  ℹ️ No export or history files found")

def cleanup_temp_files():
    """Remove temporary and cache files"""
    print("🧹 Cleaning temporary files...")
    
    temp_patterns = [
        '__pycache__/*',
        '*.pyc',
        '*.pyo', 
        '*.tmp',
        '.DS_Store',
        'Thumbs.db'
    ]
    
    removed_count = 0
    for pattern in temp_patterns:
        temp_files = glob.glob(pattern, recursive=True)
        for temp_file in temp_files:
            if os.path.exists(temp_file):
                if os.path.isfile(temp_file):
                    os.remove(temp_file)
                elif os.path.isdir(temp_file):
                    shutil.rmtree(temp_file)
                removed_count += 1
    
    if removed_count > 0:
        print(f"  ✅ Removed {removed_count} temporary files")
    else:
        print("  ℹ️ No temporary files found")

def create_sample_data():
    """Skip sample data creation - databases will be fresh"""
    print("👤 Skipping sample data creation...")
    print("  ℹ️ Databases deleted - will be created fresh on first app run")

def main():
    """Main cleanup function"""
    print("🧼 Starting Plural Chat cleanup for GitHub release...")
    print("=" * 50)
    
    # Make sure we're in the right directory
    if not os.path.exists('main.py'):
        print("❌ Error: Please run this script from the Plural Chat directory")
        return
    
    # Perform cleanup steps
    cleanup_databases()
    cleanup_avatars() 
    cleanup_logs()
    cleanup_exports()
    cleanup_temp_files()
    create_sample_data()
    
    print("=" * 50)
    print("✅ Cleanup complete! Your Plural Chat is now ready for GitHub.")
    print("📋 What was cleaned:")
    print("   • Complete database deletion (all members, chat history, settings)")
    print("   • Personal avatar images (kept default)")
    print("   • Log files and chat history exports")
    print("   • Export files and temporary data")
    print("   • Cache files and temporary files")
    print("🎉 Fresh databases ready for new users to start with!")
    print("\n💡 Tip: Test the app once more before committing to GitHub!")

if __name__ == "__main__":
    main()