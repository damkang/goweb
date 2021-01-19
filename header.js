const toggleBtn = document.querySelector('.header_toggleBtn');
const menu = document.querySelector('.header_menu');
const link = document.querySelector('.header_link');

toggleBtn.addEventListener('click',()=>{
   menu.classList.toggle('active');
   link.classList.toggle('active');
})