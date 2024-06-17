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
            name: 'serverAdd',
            component: () => import('@/views/server/add.vue')
        },
        {
            path: 'edit/:id',
            name: 'serverEdit',
            component: () => import('@/views/server/edit.vue')
        }
    ]
}
