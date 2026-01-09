/**
 * Navigation fixes for Material for MkDocs
 *
 * 1. Sidebar visibility fix: Ensures data-md-page attribute is correctly set
 *    during navigation.instant mode to prevent sidebar visibility issues.
 *
 * 2. Active tab fix: Ensures the correct navigation tab is marked as active
 *    based on the current URL path.
 */

// Only set up once - check if already initialized
if (!window.__navigationFixInitialized) {
  window.__navigationFixInitialized = true;

  function isHomePage(path) {
    return path.endsWith('/xcsh/') || path.endsWith('/xcsh/index.html') || path === '/';
  }

  function fixSidebar() {
    var path = window.location.pathname;
    var currentAttr = document.documentElement.getAttribute('data-md-page');
    var shouldBeHome = isHomePage(path);
    var isCurrentlyHome = currentAttr === 'home';

    // Only update if needed
    if (shouldBeHome && !isCurrentlyHome) {
      document.documentElement.setAttribute('data-md-page', 'home');
    } else if (!shouldBeHome && isCurrentlyHome) {
      document.documentElement.removeAttribute('data-md-page');
    }
  }

  function fixActiveTab() {
    var path = window.location.pathname;
    var tabItems = document.querySelectorAll('.md-tabs__item');

    tabItems.forEach(function(item) {
      var tab = item.querySelector('.md-tabs__link');
      if (!tab) return;

      var href = tab.getAttribute('href');
      if (!href) return;

      // Normalize the href for comparison
      var tabPath = href.replace(/^https?:\/\/[^\/]+/, '');

      // Remove existing active class from parent li
      item.classList.remove('md-tabs__item--active');

      // Check if this tab matches the current path
      // Match if the current path starts with the tab path (for section matching)
      // But not for home page (exact match only)
      var isHome = tabPath === '/xcsh/' || tabPath.endsWith('/xcsh/');
      var pathMatches = isHome
        ? isHomePage(path)
        : path.startsWith(tabPath) && !isHomePage(path);

      if (pathMatches) {
        item.classList.add('md-tabs__item--active');
      }
    });
  }

  function applyFixes() {
    fixSidebar();
    fixActiveTab();
  }

  // Run immediately
  applyFixes();

  // Subscribe to Material for MkDocs navigation events
  // This is the key - document$ fires after each instant navigation
  if (typeof document$ !== 'undefined') {
    document$.subscribe(applyFixes);
  }

  // Also handle browser back/forward
  window.addEventListener('popstate', applyFixes);

  /**
   * Accordion behavior for navigation
   * Ensures only one section is expanded at a time within each navigation level
   */
  function setupAccordion() {
    var sidebar = document.querySelector('.md-sidebar--primary');
    if (!sidebar) return;

    // Listen for toggle changes using event delegation
    sidebar.addEventListener('change', function(e) {
      if (!e.target.matches('.md-nav__toggle')) return;
      if (!e.target.checked) return; // Only act on expansion

      // Find sibling toggles at the same level
      var parentItem = e.target.closest('.md-nav__item');
      if (!parentItem) return;

      var parentList = parentItem.parentElement;
      if (!parentList) return;

      var siblingToggles = parentList.querySelectorAll(':scope > .md-nav__item > .md-nav__toggle');

      // Collapse all siblings except the one being expanded
      siblingToggles.forEach(function(toggle) {
        if (toggle !== e.target && toggle.checked) {
          toggle.checked = false;
        }
      });
    });
  }

  // Set up accordion behavior
  setupAccordion();

  // Re-setup accordion after instant navigation
  if (typeof document$ !== 'undefined') {
    document$.subscribe(setupAccordion);
  }
}
