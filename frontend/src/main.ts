import './style.css';
import {defaultThemes} from './themes/default-themes';

import {GetFAIcon, GetCurrentTheme, GetCustomThemes, SetCurrentTheme} from '../wailsjs/go/main/App';

window.openSettings = function () {
    const dialog = document.getElementById("settings-dialog") as HTMLDialogElement;
    dialog.style.opacity = "0";
    dialog.style.pointerEvents = "auto";
    window.getFAIcon("xmark", dialog.querySelector(`button[data-action="close"]`)!);
    dialog.showModal();

    requestAnimationFrame(() => {
        dialog.style.opacity = "1";
    });
};

window.closeSettings = function () {
    const dialog = document.getElementById("settings-dialog") as HTMLDialogElement;
    dialog.style.opacity = "0";
    dialog.style.pointerEvents = "none";

    dialog.addEventListener("transitionend", function handler() {
        dialog.close();
        dialog.removeEventListener("transitionend", handler);
        dialog.classList.remove("transitioning");
        dialog.style.opacity = "";
        dialog.style.pointerEvents = "";
    });
};

window.getFAIcon = (id: string, el: HTMLElement): void => {
    if (id === "") return;

    try {
        GetFAIcon(id)
            .then((result: string) => {
                el.innerHTML = result;
            }).catch((err: any) => {
                console.error(err);
            });
    } catch (err: any) {
        console.error(err);
    }
};

window.changeTab = (event: Event): void => {
    var i: number;
    var tabcontent: HTMLCollectionOf<Element>;
    var tablinks: HTMLCollectionOf<Element>;

    tabcontent = document.getElementsByClassName("tabcontent");

    for (i = 0; i < tabcontent.length; i++) {
        let content = tabcontent[i] as HTMLElement;
        content.style.display = "none";
    }

    tablinks = document.getElementsByClassName("tablinks");

    for (i = 0; i < tablinks.length; i++) {
        let link = tablinks[i] as HTMLElement;
        link.classList.remove("active");
    }

    document.getElementById("settings-" + (event.target! as HTMLButtonElement).dataset["tabId"]!)!.style.display = "block";
    (event.currentTarget! as HTMLButtonElement).classList.add("active");
};

window.updateTheme = (): void => {
    const r = document.querySelector('html')! as HTMLElement;
    GetCurrentTheme()
        .then((theme: string) => {
            r.classList.add(theme);
            console.log("Theme has been updated");
        }).catch((err: any) => {
            console.error(err);
        });
}

document.querySelector('#app')!.innerHTML = `
<dialog id="settings-dialog">
    <div class="tab">
        <button class="tablinks" onclick="window.changeTab(event)" data-tab-id="appearance">Appearance</button>
        <button class="tablinks" onclick="window.changeTab(event)" data-tab-id="members">Members</button>
    </div>

    <div id="settings-appearance" class="tabcontent">
        <label for="theme-selector">Theme:</label>
        <select id="theme-selector" onchange="window.changeCurrentTheme(event);">
	</select>
    </div>

    <div id="settings-members" class="tabcontent">

    </div>

    <button onclick="window.closeSettings()" data-action="close"></button>
</dialog>
`;

window.loadedThemes = defaultThemes;
const themeSelector: HTMLSelectElement = document.getElementById('theme-selector') as HTMLSelectElement;

GetCustomThemes().then((res: Record<string, string>) => {
    Object.entries(res).forEach((theme: [string, string]) => {
        const newStyle = document.createElement('style');
        newStyle.innerText = theme[1];
        window.loadedThemes.push(theme[0]);
        console.log('Loaded theme', theme[0]);
    });
}).catch((err: any) => {
    console.error(err);
});

window.updateTheme();

window.loadedThemes.forEach((theme) => {
    const newOption = document.createElement('option') as HTMLOptionElement;
    newOption.value = theme;
    newOption.innerText = theme;

    GetCurrentTheme()
        .then((currentTheme: string) => {
            if (theme == currentTheme) {
	        newOption.selected = true;
	    }
        }).catch((err: any) => {
            console.error(err);
        });

    themeSelector.appendChild(newOption);
});

window.changeCurrentTheme = (ev: Event) => {
    const themeTarget = ev.target as HTMLSelectElement;
    
    SetCurrentTheme(themeTarget.value)
        .then((oldTheme: string) => {
	    const html = document.querySelector('html') as HTMLHtmlElement;
	    html.classList.remove(oldTheme);
	    html.classList.add(themeTarget.value);
	});
};

declare global {
    interface Window {
        openSettings: () => void;
        closeSettings: () => void;
        getFAIcon: (id: string, el: HTMLElement) => void;
        changeTab: (event: Event) => void;
        updateTheme: () => void;
        loadedThemes: Array<string>;
	changeCurrentTheme: (ev: Event) => void;
    }
}
