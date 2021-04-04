<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <sapn>{{collapsed ? 'Blog':'My Blog'}}</sapn>
    </div>
    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index"><a-icon type="dashboard"></a-icon><span>仪表盘</span></a-menu-item>
      <a-sub-menu>
        <span slot="title"><a-icon type="file"></a-icon>文章管理 <span></span></span>
        <a-menu-item key="addart"><a-icon type="form"></a-icon><span>添加文章</span></a-menu-item>
        <a-menu-item key="artlist"><a-icon type="snippets"></a-icon><span>文章列表</span></a-menu-item>
      </a-sub-menu>

      <a-menu-item key="catelist"><a-icon type="book"></a-icon><span>分类列表</span></a-menu-item>
      <a-menu-item key="userlist"><a-icon type="user"></a-icon><span>用户列表</span></a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data () {
    return {
      collapsed: false
    }
  },
  methods: {
    goToPage (item) {
      this.$router.push('/admin/' + item.key).catch((err) => err)
    }
  }
}
</script>
  <style scoped>
.log {
  height: 32px;
  margin: 16px;
  background-color: #ffff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>
