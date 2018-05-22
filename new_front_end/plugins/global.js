import Vue from 'vue'

Vue.filter('log', (d) => {
    console.log(d)
})

Vue.filter('formatHashrate', formatHashrate)

function formatHashrate(
    hashrate,
    fixed = 2,
    del = ' ',
    prefix = 'h'
) {
    let i = 0
    const units = [
        prefix,
        'K' + prefix,
        'M' + prefix,
        'G' + prefix,
        'T' + prefix,
        'P' + prefix
    ]
    while (hashrate >= 1000) {
        hashrate = hashrate / 1000
        i++
    }

    if (!hashrate) {
        hashrate = 0
    }

    return hashrate.toFixed(fixed) + del + units[i]
}


Vue.filter('formattime', formattime)

function formattime(
    time
) {
    var mins = time % 60;
    var hours = (time - mins) / 60;
    if (mins < 10) mins = '0' + mins;
    if (hours < 10) hours = '0' + hours;
    var rezult = hours + ':' + mins; // нужный вам формат
    return rezult
}