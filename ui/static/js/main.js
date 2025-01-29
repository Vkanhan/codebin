document.addEventListener('DOMContentLoaded', function () {
    console.log('Codebin is ready!');

    const navLinks = document.querySelectorAll('nav a');
    navLinks.forEach(link => {
        link.addEventListener('click', function (event) {
            console.log(`Navigating to: ${link.getAttribute('href')}`);
        });
    });

    const latestGistSection = document.querySelector('.gist-container');
    if (latestGistSection) {
        console.log('Latest Gist section found.');
    }
});