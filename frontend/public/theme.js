const theme = localStorage.getItem('theme');
if(theme === 'dark'){
  document.body.setAttribute('data-bs-theme', 'dark')
}