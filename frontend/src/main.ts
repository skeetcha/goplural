import './style.css';
import './app.css';

window.openSettings = function () {
    const dialog = document.getElementById("test-dialog") as HTMLDialogElement;
    dialog.style.opacity = "0";
    dialog.style.pointerEvents = "auto";
    dialog.showModal();

    requestAnimationFrame(() => {
        dialog.style.opacity = "1";
    });
};

window.closeSettings = function () {
    const dialog = document.getElementById("test-dialog") as HTMLDialogElement;
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

document.querySelector('#app')!.innerHTML = `
<dialog id="test-dialog">
    <p>Greetings</p>
    <button onclick="window.closeSettings()">OK</button>
</dialog>
`;

declare global {
    interface Window {
        openSettings: () => void;
        closeSettings: () => void;
    }
}
