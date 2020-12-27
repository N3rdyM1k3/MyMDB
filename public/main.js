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

const deleteButton = document.querySelector('#delete-button')
const messageDiv = document.querySelector('#message')


deleteButton.addEventListener('click', _=> {
  fetch('/movies', {
    method: 'delete',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({title: 'Parasite'})
  }).then(res => {
    if (res.ok) return res.json()
  }).then(response => {
    if (response == 'Nothing to delete'){
      messageDiv.textContent = 'Nothing to delete'
    } else {
      window.location.reload(true)
  }})
})



