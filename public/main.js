const update = document.querySelector('#update-button')


update.addEventListener('click', _=> {
    fetch('/movies', {
        method: 'put',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({
            title: 'Midsommar'
        })
    }).then(res => {
        if (res.ok) return res.json()
      })
      .then(response => {
        window.location.reload(true)
      })
      .then(response => {
        console.log(response)
      })
})