const modules = import.meta.glob('../views/**/*.vue')

export const asyncRouterHandle = (asyncRouter) => {
    asyncRouter.forEach(item => {
        if (item.component) {
            item.component = dynamicImport(modules, item.component)
        } else {
            delete item['component']
        }
        if (item.children) {
            asyncRouterHandle(item.children)
        }
    })
}

function dynamicImport(
    dynamicViewsModules,
    component
) {
    const keys = Object.keys(dynamicViewsModules)
    console.log(dynamicViewsModules)
    console.log(component)

    console.log("key:")
    console.log(keys)

    const matchKeys = keys.filter((key) => {
        const k = key.replace('../', '')
        return k === component
    })
    const matchKey = matchKeys[0]
    console.log(matchKey)
    console.log(dynamicViewsModules[matchKey])

    return dynamicViewsModules[matchKey]
}
