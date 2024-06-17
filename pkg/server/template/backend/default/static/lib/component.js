const { defineAsyncComponent } = Vue
export function registerComponent(name) {
    if (window.asyncComponents === undefined) {
        window.asyncComponents = {}
    }
    if (window.asyncComponents[name]) {
        return window.asyncComponents[name]
    }
    return window.asyncComponents[name] = defineAsyncComponent(async () => {
        let componentRes = await fetch('./component/' + name + '/index.html')
        let componentJs = await import('../../component/' + name + '/index.js')
        return new Promise(async (resolve) => {
            resolve({
                template: await componentRes.text(),
                ...componentJs.default
            })
        })
    })
}
