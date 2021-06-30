const toBRL = (num)=>{
    let currency = num.toString().replace(/\D+/g, '');
    let cur = currency.length;
    let ren = currency.slice(0, cur-2);
    let cy = currency.slice(cur-2);
    currency = `R$${ren},${cy}`
    return currency;
}

console.log(toBRL(10000));
console.log(toBRL("1bb4mi5"));
console.log(toBRL(20));
console.log(toBRL(200));
