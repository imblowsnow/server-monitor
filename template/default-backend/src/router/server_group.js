export default {
    path: 'server_group',
    children: [
        {
            path: '',
            name: 'server_group',
            component: () => import('@/views/server_group/index.vue')
        },
        {
            path: 'add',
            name: 'server_group_add',
            component: () => import('@/views/server_group/add.vue')
        },
        {
            path: 'edit/:id',
            name: 'server_group_edit',
            component: () => import('@/views/server_group/edit.vue')
        }
    ]
}
