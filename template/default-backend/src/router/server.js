export default {
    path: 'server',
    children: [
        {
            path: ':id',
            name: 'server',
            component: () => import('@/views/server/index.vue')
        },
        {
            path: 'add',
            name: 'server_add',
            component: () => import('@/views/server/add.vue')
        },
        {
            path: 'edit/:id',
            name: 'server_edit',
            component: () => import('@/views/server/edit.vue')
        }
    ]
}
