import request from "@/api/index";

export const GetProducts = async function () {
    try {
        const resp = await request('GET', {
            path: '/v1/products',
        })
        const prs: Products[] = resp.data.products
        return prs
    } catch (e) {
        console.log(e)
        return Promise.reject(e)
    }
}



