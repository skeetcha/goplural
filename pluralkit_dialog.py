import tkinter as tk
from tkinter import messagebox
import ttkbootstrap as ttk
from ttkbootstrap.constants import *
import threading
from datetime import datetime
import logging
from logging.handlers import RotatingFileHandler


class PluralKitDialog:
    """Dialog for PluralKit integration setup and sync"""
    
    def __init__(self, parent, pk_sync, refresh_callback):
        self.parent = parent
        self.pk_sync = pk_sync
        self.refresh_callback = refresh_callback
        self.dialog = None
        self.logger = logging.getLogger('plural_chat.pluralkit_dialog')
        
    def show(self):
        """Show the PluralKit dialog"""
        self.dialog = tk.Toplevel(self.parent)
        self.dialog.title("PluralKit Integration")
        self.dialog.geometry("500x400")
        self.dialog.resizable(False, False)
        self.dialog.transient(self.parent)
        self.dialog.grab_set()
        
        # Center the dialog
        self.dialog.update_idletasks()
        x = (self.dialog.winfo_screenwidth() // 2) - (500 // 2)
        y = (self.dialog.winfo_screenheight() // 2) - (400 // 2)
        self.dialog.geometry(f"500x400+{x}+{y}")
        
        self.setup_ui()
        
    def setup_ui(self):
        """Setup the dialog UI"""
        main_frame = ttk.Frame(self.dialog, padding=20)
        main_frame.pack(fill=BOTH, expand=True)
        
        # Title
        title_label = ttk.Label(main_frame, text="PluralKit Integration", 
                               font=("Arial", 16, "bold"))
        title_label.pack(pady=(0, 20))
        
        # Token section
        token_frame = ttk.LabelFrame(main_frame, text="API Token", padding=10)
        token_frame.pack(fill=X, pady=(0, 15))
        
        ttk.Label(token_frame, text="Enter your PluralKit token:").pack(anchor=W)
        
        token_info = ttk.Label(token_frame, 
                              text="Get your token from: https://dash.pluralkit.me/dash/token",
                              font=("Arial", 8))
        token_info.pack(anchor=W, pady=(0, 5))
        
        self.token_entry = ttk.Entry(token_frame, width=50, show="*")
        self.token_entry.pack(fill=X, pady=(0, 10))
        
        # Check if token already exists
        existing_token = self.pk_sync.app_db.get_api_token("pluralkit")
        if existing_token:
            self.token_entry.insert(0, existing_token)
        
        token_button_frame = ttk.Frame(token_frame)
        token_button_frame.pack(fill=X)
        
        self.test_button = ttk.Button(token_button_frame, text="Test Connection", 
                                     command=self.test_connection, bootstyle="info")
        self.test_button.pack(side=LEFT, padx=(0, 10))
        
        self.save_token_button = ttk.Button(token_button_frame, text="Save Token", 
                                           command=self.save_token, bootstyle="success")
        self.save_token_button.pack(side=LEFT)
        
        # Status label
        self.status_label = ttk.Label(main_frame, text="", font=("Arial", 10))
        self.status_label.pack(pady=(0, 15))
        
        # Sync section
        sync_frame = ttk.LabelFrame(main_frame, text="Sync Options", padding=10)
        sync_frame.pack(fill=X, pady=(0, 15))
        
        self.download_avatars_var = tk.BooleanVar(value=True)
        ttk.Checkbutton(sync_frame, text="Download avatars", 
                       variable=self.download_avatars_var).pack(anchor=W, pady=(0, 10))
        
        sync_button_frame = ttk.Frame(sync_frame)
        sync_button_frame.pack(fill=X)
        
        self.sync_button = ttk.Button(sync_button_frame, text="Sync Members", 
                                     command=self.sync_members, bootstyle="primary")
        self.sync_button.pack(side=LEFT, padx=(0, 10))
        
        self.import_button = ttk.Button(sync_button_frame, text="Full Import", 
                                       command=self.full_import, bootstyle="warning")
        self.import_button.pack(side=LEFT)
        
        # Progress section
        progress_frame = ttk.Frame(main_frame)
        progress_frame.pack(fill=X, pady=(0, 15))
        
        self.progress_label = ttk.Label(progress_frame, text="")
        self.progress_label.pack(anchor=W)
        
        self.progress_bar = ttk.Progressbar(progress_frame, mode='indeterminate')
        self.progress_bar.pack(fill=X, pady=(5, 0))
        
        # Close button
        close_frame = ttk.Frame(main_frame)
        close_frame.pack(fill=X)
        
        ttk.Button(close_frame, text="Close", command=self.close_dialog, 
                  bootstyle="secondary").pack(side=RIGHT)
        
        # Initial status check
        if existing_token:
            self.check_connection_status()
    
    def test_connection(self):
        """Test the PluralKit API connection"""
        token = self.token_entry.get().strip()
        if not token:
            self.status_label.config(text="Please enter a token", foreground="red")
            return
        
        self.pk_sync.api.set_token(token)
        success, message = self.pk_sync.api.test_connection()
        
        if success:
            self.status_label.config(text=f"✓ {message}", foreground="green")
        else:
            self.status_label.config(text=f"✗ {message}", foreground="red")
    
    def save_token(self):
        """Save the PluralKit token"""
        token = self.token_entry.get().strip()
        if not token:
            messagebox.showerror("Error", "Please enter a token")
            return
        
        success, message = self.pk_sync.setup_token(token)
        
        if success:
            self.status_label.config(text=f"✓ Token saved: {message}", foreground="green")
            messagebox.showinfo("Success", "Token saved successfully!")
        else:
            self.status_label.config(text=f"✗ {message}", foreground="red")
            messagebox.showerror("Error", f"Failed to save token: {message}")
    
    def check_connection_status(self):
        """Check if saved token still works"""
        if self.pk_sync.load_saved_token():
            success, message = self.pk_sync.api.test_connection()
            if success:
                self.status_label.config(text=f"✓ Connected: {message}", foreground="green")
            else:
                self.status_label.config(text=f"✗ Connection issue: {message}", foreground="orange")
        else:
            self.status_label.config(text="No token configured", foreground="gray")
    
    def sync_members(self):
        """Sync members from PluralKit"""
        if not self.pk_sync.load_saved_token():
            messagebox.showerror("Error", "Please save a valid token first")
            return
        
        self.progress_label.config(text="Syncing members...")
        self.progress_bar.start()
        self.sync_button.config(state=DISABLED)
        
        def sync_thread():
            try:
                download_avatars = self.download_avatars_var.get()
                new_count, updated_count, errors = self.pk_sync.sync_members(download_avatars)
                
                # Update UI in main thread
                self.dialog.after(0, self.sync_complete, new_count, updated_count, errors)
            except Exception as e:
                self.dialog.after(0, self.sync_error, str(e))
        
        threading.Thread(target=sync_thread, daemon=True).start()
    
    def full_import(self):
        """Perform full import from PluralKit"""
        if not messagebox.askyesno("Full Import Warning", 
                                  "This will replace all existing members. Continue?"):
            return
        
        if not self.pk_sync.load_saved_token():
            messagebox.showerror("Error", "Please save a valid token first")
            return
        
        self.progress_label.config(text="Importing system...")
        self.progress_bar.start()
        self.import_button.config(state=DISABLED)
        
        def import_thread():
            try:
                download_avatars = self.download_avatars_var.get()
                success, message, stats = self.pk_sync.import_full_system(download_avatars)
                
                # Update UI in main thread
                self.dialog.after(0, self.import_complete, success, message, stats)
            except Exception as e:
                self.dialog.after(0, self.import_error, str(e))
        
        threading.Thread(target=import_thread, daemon=True).start()
    
    def sync_complete(self, new_count, updated_count, errors):
        """Handle sync completion"""
        self.progress_bar.stop()
        self.sync_button.config(state=NORMAL)
        
        if errors:
            error_msg = "\n".join(errors[:5])  # Show first 5 errors
            if len(errors) > 5:
                error_msg += f"\n... and {len(errors) - 5} more errors"
            
            self.progress_label.config(text=f"Sync completed with errors")
            messagebox.showwarning("Sync Warning", 
                                 f"Added {new_count}, updated {updated_count} members.\n\nErrors:\n{error_msg}")
        else:
            self.progress_label.config(text=f"Sync complete: +{new_count}, ~{updated_count}")
            messagebox.showinfo("Sync Complete", 
                               f"Successfully synced members!\n\nNew: {new_count}\nUpdated: {updated_count}")
        
        # Refresh the main app
        self.refresh_callback()
    
    def import_complete(self, success, message, stats):
        """Handle import completion"""
        self.progress_bar.stop()
        self.import_button.config(state=NORMAL)
        
        if success:
            self.progress_label.config(text="Import complete")
            details = f"Members imported: {stats['members_imported']}\n"
            details += f"Avatars downloaded: {stats['avatars_downloaded']}"
            
            if stats['errors']:
                details += f"\nErrors: {len(stats['errors'])}"
            
            messagebox.showinfo("Import Complete", f"{message}\n\n{details}")
        else:
            self.progress_label.config(text="Import failed")
            messagebox.showerror("Import Error", message)
        
        # Refresh the main app
        self.refresh_callback()
    
    def sync_error(self, error):
        """Handle sync error"""
        self.progress_bar.stop()
        self.sync_button.config(state=NORMAL)
        self.progress_label.config(text="Sync failed")
        messagebox.showerror("Sync Error", f"Failed to sync: {error}")
    
    def import_error(self, error):
        """Handle import error"""
        self.progress_bar.stop()
        self.import_button.config(state=NORMAL)
        self.progress_label.config(text="Import failed")
        messagebox.showerror("Import Error", f"Failed to import: {error}")
    
    def close_dialog(self):
        """Close the dialog"""
        if self.dialog:
            self.dialog.destroy()