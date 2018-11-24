<template>
    <div id="app" class="container-fluid">
        <div class="row">
            <div class="col-md-3">
                <nav-menu params="route: route"></nav-menu>
            </div>
            <div class="col-sm-9">
                <router-view></router-view>
                <footer class="my-5 pt-5 text-muted text-center text-small">
                    <p class="mb-1">&copy; 2017-2018 C-OOZ.com</p>
                    <ul class="list-inline">
                    <li class="list-inline-item"><a href="#">Privacy</a></li>
                    <li class="list-inline-item"><a href="#">Terms</a></li>
                    <li class="list-inline-item"><a href="#">Support</a></li>
                    </ul>
                </footer>
            </div>
        </div>
    </div>
</template>

<script>
import NavMenu from "./nav-menu";

export default {
  components: {
    "nav-menu": NavMenu
  },

  data() {
    return {};
  }
};
</script>

<style>
</style>
