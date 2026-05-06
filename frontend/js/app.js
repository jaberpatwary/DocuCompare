const API_BASE = '/api';

// --- Profile Loader ---
async function loadProfile() {
    try {
        const response = await fetch(`${API_BASE}/profile`);
        const profile = await response.json();
        if (!profile.name) return;

        // Hero Name
        document.querySelector('#hero h2').textContent = profile.name + ".";
        document.querySelector('#hero h3').textContent = profile.headline;
        document.querySelector('.hero-desc').textContent = profile.bio;

        // About Grid
        const aboutText = document.querySelector('.about-text');
        if (profile.about) {
            aboutText.innerHTML = profile.about.split('\n\n').map(p => `<p>${p.replace(/\n/g, '<br>')}</p>`).join('');
            // Skills List
            if (profile.skills) {
                const skillsArray = profile.skills.split(',').map(s => s.trim()).filter(s => s);
                if (skillsArray.length > 0) {
                    aboutText.innerHTML += `
                        <p>Here are a few technologies I've been working with recently:</p>
                        <ul style="display: grid; grid-template-columns: repeat(2, minmax(140px, 200px)); gap: 0 10px; padding: 0; margin: 20px 0 0 0; overflow: hidden; list-style: none;">
                            ${skillsArray.map(skill => `<li style="position: relative; margin-bottom: 10px; padding-left: 20px; font-family: var(--font-mono); font-size: 13px;">▹ ${skill}</li>`).join('')}
                        </ul>
                    `;
                }
            }
        }

        // About Image
        const aboutPic = document.querySelector('.about-pic .wrapper');
        const img = profile.image && profile.image.trim() !== '' ? profile.image : 'https://via.placeholder.com/300';
        aboutPic.innerHTML = `<img src="${img}" alt="${profile.name}">`;

        // Socials Sidebars
        const socialList = document.querySelector('.social-list');
        const mobileSocials = document.querySelector('.social-links-mobile'); // For footer if needed

        let socialHtml = '';
        if (profile.github) socialHtml += `<li><a href="${profile.github}" target="_blank" aria-label="GitHub"><svg role="img" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"></path></svg></a></li>`;
        if (profile.instagram) socialHtml += `<li><a href="${profile.instagram}" target="_blank" aria-label="Instagram"><svg role="img" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg></a></li>`;
        if (profile.twitter) socialHtml += `<li><a href="${profile.twitter}" target="_blank" aria-label="Twitter"><svg role="img" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M23 3a10.9 10.9 0 0 1-3.14 1.53 4.48 4.48 0 0 0-7.86 3v1A10.66 10.66 0 0 1 3 4s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5a4.5 4.5 0 0 0-.08-.83A7.72 7.72 0 0 0 23 3z"></path></svg></a></li>`;
        if (profile.linkedin) socialHtml += `<li><a href="${profile.linkedin}" target="_blank" aria-label="LinkedIn"><svg role="img" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"></path><rect x="2" y="9" width="4" height="12"></rect><circle cx="4" cy="4" r="2"></circle></svg></a></li>`;
        if (profile.facebook) socialHtml += `<li><a href="${profile.facebook}" target="_blank" aria-label="Facebook"><svg role="img" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"></path></svg></a></li>`;
        if (profile.youtube) socialHtml += `<li><a href="${profile.youtube}" target="_blank" aria-label="YouTube"><svg role="img" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22.54 6.42a2.78 2.78 0 0 0-1.94-2C18.88 4 12 4 12 4s-6.88 0-8.6.46a2.78 2.78 0 0 0-1.94 2A29 29 0 0 0 1 11.75a29 29 0 0 0 .46 5.33A2.78 2.78 0 0 0 3.4 19c1.72.46 8.6.46 8.6.46s6.88 0 8.6-.46a2.78 2.78 0 0 0 1.94-2 29 29 0 0 0 .46-5.33 29 29 0 0 0-.46-5.33z"></path><polygon points="9.75 15.02 15.5 11.75 9.75 8.48 9.75 15.02"></polygon></svg></a></li>`;

        socialList.innerHTML = socialHtml;

        // Email
        if (profile.email) {
            const emailLink = document.querySelector('.email-link');
            if (emailLink) {
                emailLink.textContent = profile.email;
                emailLink.href = `mailto:${profile.email}`;
            }
        }

    } catch (e) { console.error("Error loading profile", e); }
}

// --- Experience (Tabs) ---
async function loadExperience() {
    try {
        const response = await fetch(`${API_BASE}/experience`);
        const experiences = await response.json();

        const tabsContainer = document.querySelector('.jobs-tabs');
        const contentContainer = document.querySelector('.jobs-content');

        tabsContainer.innerHTML = '';
        contentContainer.innerHTML = '';

        if (!experiences || experiences.length === 0) return;

        experiences.forEach((exp, index) => {
            // Create Tab
            const btn = document.createElement('button');
            btn.className = `tab-button ${index === 0 ? 'active' : ''}`;
            btn.textContent = exp.company;
            btn.id = `tab-${index}`;
            btn.setAttribute('role', 'tab');
            btn.setAttribute('aria-selected', index === 0);
            btn.setAttribute('aria-controls', `panel-${index}`);
            btn.onclick = () => switchTab(index);
            tabsContainer.appendChild(btn);

            // Create Panel
            const panel = document.createElement('div');
            panel.className = `job-panel ${index === 0 ? 'active' : ''}`;
            panel.id = `panel-${index}`;
            panel.setAttribute('role', 'tabpanel');
            panel.setAttribute('aria-labelledby', `tab-${index}`);

            panel.innerHTML = `
                <h3 class="job-title">
                    <span>${exp.role}</span>
                    <span class="job-company"> @ ${exp.company}</span>
                </h3>
                <p class="job-range">${exp.start_date} - ${exp.end_date}</p>
                <div class="job-desc">
                    <ul>
                        ${exp.description ? `<li>${exp.description}</li>` : ''}
                        ${exp.technologies ? `<li>Used: ${exp.technologies}</li>` : ''}
                    </ul>
                </div>
            `;
            contentContainer.appendChild(panel);
        });

    } catch (e) { console.error("Error loading experience", e); }
}

function switchTab(index) {
    const tabs = document.querySelectorAll('.tab-button');
    const panels = document.querySelectorAll('.job-panel');

    tabs.forEach(t => {
        t.classList.remove('active');
        t.setAttribute('aria-selected', 'false');
    });
    panels.forEach(p => p.classList.remove('active'));

    tabs[index].classList.add('active');
    tabs[index].setAttribute('aria-selected', 'true');
    panels[index].classList.add('active');
}

// --- Projects (Featured & Grid) ---
async function loadProjects() {
    try {
        const response = await fetch(`${API_BASE}/projects?limit=10`);
        const projects = await response.json();

        const featuredContainer = document.querySelector('.featured-projects');
        const gridContainer = document.querySelector('.projects-grid');

        featuredContainer.innerHTML = '';
        gridContainer.innerHTML = '';

        if (!projects || projects.length === 0) return;

        // Take first 3 as Featured, rest as Grid
        projects.forEach((proj, index) => {
            if (proj.featured) {
                renderFeaturedProject(proj, index, featuredContainer);
            } else {
                renderGridProject(proj, gridContainer);
            }
        });

    } catch (e) { console.error("Error loading projects", e); }
}

function renderFeaturedProject(proj, index, container) {
    const li = document.createElement('li');
    li.className = 'project-item';
    const isEven = index % 2 === 0; // odd/even logic for flipping

    const imgUrl = proj.image || 'https://via.placeholder.com/800x500';

    li.innerHTML = `
        <div class="project-content ${isEven ? '' : 'inverted'}">
            <p class="project-overline">Featured Project</p>
            <h3 class="project-title"><a href="#">${proj.title}</a></h3>
            <div class="project-desc-box">
                <p>${proj.content.substring(0, 150)}...</p>
            </div>
            <ul class="project-tech-list">
                ${proj.category ? proj.category.split(',').map(t => `<li>${t.trim()}</li>`).join('') : ''}
            </ul>
        </div>
        <div class="project-image ${isEven ? '' : 'inverted'}">
            <a href="#">
                <img src="${imgUrl}" alt="${proj.title}">
            </a>
        </div>
    `;
    container.appendChild(li);
}

function renderGridProject(proj, container) {
    const li = document.createElement('li');
    li.innerHTML = `
        <div class="inner-card">
            <div style="width: 100%;">
                <div class="project-top">
                    <div class="folder">
                        <svg xmlns="http://www.w3.org/2000/svg" role="img" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>
                    </div>
                </div>
                <h3 class="project-card-title">${proj.title}</h3>
                <div class="project-card-desc">
                    <p>${proj.content.substring(0, 100)}...</p>
                </div>
            </div>
            <ul class="project-card-tech">
                 ${proj.category ? proj.category.split(',').map(t => `<li>${t.trim()}</li>`).join('') : ''}
            </ul>
        </div>
    `;
    container.appendChild(li);
}

// --- Navigation ---
async function loadNavigation() {
    try {
        const res = await fetch(`${API_BASE}/nav-items`);
        const items = await res.json();

        const navList = document.querySelector('.nav-menu ol');
        if (navList && items.length > 0) {
            navList.innerHTML = items.map((item, index) => {
                const num = (index + 1).toString().padStart(2, '0');
                return `<li><a href="${item.link}" class="nav-link"><span>${num}.</span> ${item.label}</a></li>`;
            }).join('');
        }
    } catch (e) { console.error("Error loading nav", e); }
}

// --- New Sections Loaders ---

async function loadAchievements() {
    try {
        const res = await fetch(`${API_BASE}/achievements`);
        const data = await res.json();
        const container = document.getElementById('achievements-container');
        if (!container) return;

        container.innerHTML = data.map(i => `
            <div style="background: var(--light-navy); padding: 20px; border-radius: 4px; border: 1px solid var(--lightest-navy);">
                <h3 style="color: var(--lightest-slate); margin-bottom: 5px;">${i.title}</h3>
                <span style="font-family: var(--font-mono); font-size: 13px; color: var(--green);">${i.date}</span>
                <p style="margin-top: 10px; font-size: 14px;">${i.description}</p>
            </div>
        `).join('');
    } catch (e) { console.error("Error loading achievements", e); }
}

async function loadPhotos() {
    try {
        const res = await fetch(`${API_BASE}/photos`);
        const data = await res.json();
        const container = document.getElementById('photos-container');
        if (!container) return;

        container.innerHTML = data.map(i => `
            <div style="width: 300px; overflow: hidden; border-radius: 8px; position: relative; group">
                <img src="${i.image}" alt="${i.caption}" style="width: 100%; height: 200px; object-fit: cover; transition: transform 0.3s ease;">
                ${i.caption ? `<div style="padding: 10px; background: rgba(0,0,0,0.7); position: absolute; bottom: 0; width: 100%; color: white; font-size: 14px;">${i.caption}</div>` : ''}
            </div>
        `).join('');
    } catch (e) { console.error("Error loading photos", e); }
}

async function loadVideos() {
    try {
        const res = await fetch(`${API_BASE}/videos`);
        const data = await res.json();
        const container = document.getElementById('videos-container');
        if (!container) return;

        container.innerHTML = data.map(i => {
            // Extract Video ID
            let videoId = '';
            if (i.youtube_link.includes('v=')) {
                videoId = i.youtube_link.split('v=')[1].split('&')[0];
            } else if (i.youtube_link.includes('youtu.be/')) {
                videoId = i.youtube_link.split('youtu.be/')[1];
            }

            return `
            <div style="background: var(--light-navy); border-radius: 8px; overflow: hidden;">
                <iframe width="100%" height="220" src="https://www.youtube.com/embed/${videoId}" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
                <div style="padding: 15px;">
                    <h3 style="color: var(--lightest-slate); margin-bottom: 5px;">${i.title}</h3>
                    <p style="font-size: 14px;">${i.description}</p>
                </div>
            </div>
            `;
        }).join('');
    } catch (e) { console.error("Error loading videos", e); }
}

// Init
document.addEventListener('DOMContentLoaded', () => {
    loadProfile();
    loadExperience();
    loadProjects();
    // loadNavigation(); // Disabled to show static menu items
    loadAchievements();
    loadPhotos();
    loadVideos();
});
