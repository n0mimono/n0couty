export function toArray(obj: any) {
    return Object.keys(obj).map(key => {return obj[key]});    
}
