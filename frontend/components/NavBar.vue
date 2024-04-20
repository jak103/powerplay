<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <NuxtLink class="navbar-brand" to="/">Eccles Ice Hockey</NuxtLink>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <NuxtLink class="nav-link" to="/">Home</NuxtLink>
          </li>
          <li class="nav-item">
            <NuxtLink class="nav-link" to="/calendar">Calendar</NuxtLink>
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              Information
            </a>
            <ul class="dropdown-menu dropdown-menu-end">
              <li>
                <NuxtLink class="dropdown-item" to="/information">*How To Join</NuxtLink>
              </li>
              <li>
                <NuxtLink class="dropdown-item" to="/information"
                  >*How To Join</NuxtLink
                >
              </li>
              <li>
                <NuxtLink class="dropdown-item" to="/information"
                  >*For Players</NuxtLink
                >
              </li>
              <li>
                <NuxtLink class="dropdown-item" to="/information/managers"
                  >For Managers</NuxtLink
                >
              </li>
              <li>
                <NuxtLink class="dropdown-item" to="/information"
                  >*Our Staff</NuxtLink
                >
              </li>
              <li>
                <NuxtLink class="dropdown-item" to="/information"
                  >*Our Rink</NuxtLink
                >
              </li>
              <li>
                <NuxtLink class="dropdown-item" to="/information"
                  >*Substitutions</NuxtLink
                >
              </li>
              <li>
                <NuxtLink class="dropdown-item" to="/information"
                  >*Suspensions</NuxtLink
                >
              </li>
            </ul>
          </li>
        </ul>
        <NuxtLink to="/app/sign-in" class="btn btn-primary ms-2"
          >Sign In</NuxtLink
        >
      </div>
    </div>
  </nav>
</template>
