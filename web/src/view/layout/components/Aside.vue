<script setup>
import { MENU_CONFIG } from '../../../config/menu';
import * as Icons from '@element-plus/icons-vue';
import { useIsCollapse } from '../../../store/index.js'
import { storeToRefs } from 'pinia';

// 获取菜单栏折叠状态
const {isCollapse} = storeToRefs(useIsCollapse())

// 创建一个获取图标组件的辅助函数
const getIcon = (iconName) => {
  if (!iconName) return null;
  return Icons[iconName] || null;
}
</script>

<template>
    <!-- 左侧边栏:style是动态样式 -->
<el-aside 
    class='el-aside' 
    :style="{width:isCollapse?'65px':'320px'}"
> 
    <!-- 平台名称 -->
    <router-link to="/cluster/dashboard">
        <el-button text>
            <el-icon><Icons.HomeFilled /></el-icon>
            <span v-show="!isCollapse">Ks</span>
        </el-button>
    </router-link>
    <!-- 菜单 -->
     <div id="menu">
        <el-menu
            :default-active="$route.path"
            class="el-menu-vertical-demo"
            router
            :collapse="isCollapse"
            style="border: none;"
            :collapse-transition="false"
        >
            <!-- 自动生成菜单 -->
            <el-sub-menu v-for="menu in MENU_CONFIG" :key="menu.index" :index="menu.index">
                <template #title>
                    <el-icon> 
                        <component :is="getIcon(menu.icon)" />
                    </el-icon>
                    <span style="margin-left: 5px;">{{ menu.title }}</span>
                </template>
                <!-- 如果有子菜单 -->
                <el-sub-menu v-if="menu.subMenu" v-for="subMenu in menu.subMenu" :key="subMenu.index" :index="subMenu.index">
                    <template #title>
                        <el-icon> 
                            <component :is="getIcon(subMenu.icon)" />
                        </el-icon>
                        {{ subMenu.title }}
                    </template>
                    <el-menu-item v-for="subItem in subMenu.items" :key="subItem.index" :index="subItem.index">
                        <template #title>{{subItem.title}}</template>
                    </el-menu-item>
                </el-sub-menu>
                <!-- 没有子菜单 -->
                <el-menu-item v-for="item in menu.items" :key="item.index" :index="item.index">
                    <template #title>
                        <el-icon> 
                            <component :is="getIcon(item.icon)" />
                        </el-icon>
                        {{item.title}}
                    </template>
                </el-menu-item>                 
            </el-sub-menu>
        </el-menu>
     </div>
      
</el-aside>
</template> 

<style lang="less" scoped>
.el-aside{
    border-right: 1px solid #cccccc;
    button{
        width: 100%;
        font-size: 25px;
        margin:30px auto;
    }
    // 折叠动画更加丝滑
    transition:width 0.15s;
    -webkit-transition: width 0.15s;
    -moz-transition: width 0.15s;
    -webkit-transition: width 0.15s;
    -o-transition: width 0.15s;
}
</style>
