import request from "../request.js";

export const getRecentArticle = (params) => {
    return request('get', `/article/recent`, params)
}