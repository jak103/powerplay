<script lang="ts" setup></script>

<template>
  <div
    class="position-fixed z-3"
    style="right: 50%; transform: translateX(50%)"
  >
    <button
      class="btn btn-danger btn-sm rounded-top-0"
      data-bs-toggle="offcanvas"
      data-bs-target="#offcanvasTop"
      aria-controls="offcanvasTop"
    >
      Dev <span class="d-none d-sm-inline">Mode</span>
      -
      <span class="d-sm-none">XS</span>
      <span class="d-none d-sm-inline d-md-none">SM</span>
      <span class="d-none d-md-inline d-lg-none">MD</span>
      <span class="d-none d-lg-inline d-xl-none">LG</span>
      <span class="d-none d-xl-inline d-xxl-none">XL</span>
      <span class="d-none d-xxl-inline">XXL</span>
    </button>
  </div>

  <div
    class="offcanvas offcanvas-top text-bg-danger"
    data-bs-theme="dark"
    tabindex="-1"
    id="offcanvasTop"
    aria-labelledby="offcanvasTopLabel"
  >
    <div class="offcanvas-header">
      <h5 class="offcanvas-title" id="offcanvasTopLabel">Dev Mode</h5>
      <button
        type="button"
        class="btn-close"
        data-bs-dismiss="offcanvas"
        aria-label="Close"
      ></button>
    </div>
    <div class="offcanvas-body">
      <p>
        You can put information here about editing in development mode, as well
        as add some helpful links.
      </p>
      <div class="hstack gap-3">
        <a
          class="btn btn-dark"
          href="https://vuejs.org/guide/introduction"
          target="_blank"
          >Vue 3</a
        >
        <a
          class="btn btn-dark"
          href="https://nuxt.com/docs/getting-started/introduction"
          target="_blank"
          >Nuxt 3</a
        >
        <a
          class="btn btn-dark"
          href="https://getbootstrap.com/docs/5.3/getting-started/introduction/"
          target="_blank"
          >Bootstrap 5</a
        >
      </div>
    </div>
  </div>
</template>
