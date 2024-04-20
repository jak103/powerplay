// This will initialize the theme without a quick flash of light theme, then changing to dark theme on the first load.

const theme = localStorage.getItem('theme');
if(theme === 'dark'){
  document.body.setAttribute('data-bs-theme', 'dark')
}