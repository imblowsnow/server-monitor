import backendRouter from './backend.js'
import frontendRouter from './frontend.js'

let router = [];
if (import.meta.env.MODE === 'frontend') {
    router = frontendRouter
}else{
    router = backendRouter
}

export default router
