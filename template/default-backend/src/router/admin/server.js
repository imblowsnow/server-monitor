export default {
    path: 'server',
    children: [
        {
            path: ':id',
            name: 'server',
            component: () => import('@/views/admin/server/index.vue')
        },
        {
            path: 'add',
            name: 'server_add',
            component: () => import('@/views/admin/server/add.vue')
        },
        {
            path: 'edit/:id',
            name: 'server_edit',
            component: () => import('@/views/admin/server/edit.vue')
        }
    ]
}
