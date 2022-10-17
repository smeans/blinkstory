document.addEventListener('DOMContentLoaded', () => {
	phInit();
})

document.addEventListener('ph-loaded', () => {
	initSite();
})

function initSite() {
	mainDate.date = new Date();
}