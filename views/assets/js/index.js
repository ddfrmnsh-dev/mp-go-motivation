const apiUrl = "http://localhost:8000"

async function fetchRandomData() {
    try {
        const response = await fetch(`${apiUrl}/random`);

        if (!response.ok) {
            throw new Error("Network response was not ok");
        }

        const data = await response.json();
        const resData = data.motivation

        displayRandomData(resData)
    } catch (error) {
        console.error("Error:", error);
    }
}

async function fetchAllData() {
    try {
        const res = await fetch(`${apiUrl}/getAllMotivation`);

        if (!res.ok) {
            throw new Error("Network response was not ok");
        }

        const resData = await res.json();

        const data = resData.data

        displayAllData(data)
    } catch (error) {
        console.error("Error :", error);
    }
}

function displayAllData(data) {
    try {
        const dataContainer = $('#data-container');

        $.each(data, function (index, item) {
            const listItem = $('<div>').addClass('quote-item__content');
            const title = $('<h3>').addClass('quote-item__title').text(item.name);
            const date = $('<p>').addClass('quote-item__date').text(item.quote);
            const body = $('<p>').addClass('quote-item__body').text(item.created_at);

            listItem.append(title, date, body)

            const listAction = $('<div>').addClass('quote-item__action');
            const buttonDelete = $('<button>').addClass('quote-item__delete-button').text('Delete');
            const buttonEdit = $('<button>').addClass('quote-item__archive-button').text('Edit');

            buttonDelete.on('click', function () {
                deleteData(item.id)
            })

            buttonEdit.on('click', function () {
                openEditModal(item)
            })

            listAction.append(buttonDelete, buttonEdit)

            const wrapperItem = $('<div>').addClass('quote-item');

            wrapperItem.append(listItem, listAction)

            dataContainer.append(wrapperItem)
        })
    } catch (error) {
        console.error("Error:", error);
    }
}

function openEditModal(data) {
    console.log("cek modal", data)
    $('#editName').val(data.name);
    $('#editMotivation').val(data.quote);

    $('#editModal').modal('show');

    $('#editForm').submit(function (e) {
        e.preventDefault()
        const newName = $('#editName').val();
        const newQuote = $('#editMotivation').val();

        if (newName.trim() !== '' && newQuote.trim() !== '') {
            $('#editModal').modal('hide');

            updateData(data.id, { name: newName, quote: newQuote })
        }
    })
}

async function deleteData(quoteId) {
    try {
        const deleteUrl = `${apiUrl}/motivation/${quoteId}`
        const deleteRes = await fetch(deleteUrl, {
            method: 'Delete',
        })

        // if (!deleteRes.ok) {
        //     throw new Error('Failed to delete data')
        // }
        console.log(deleteRes)
        window.location.reload();

    } catch (error) {
        console.error("Error:", error);
    }
}

async function updateData(itemId, newData) {
    try {
        const updateUrl = `${apiUrl}/motivation/${itemId}`
        const editRes = await fetch(updateUrl, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(newData)
        })

        console.log(editRes)

        // window.location.reload();
    } catch (error) {
        console.error("Error:", error);
    }
}

function displayRandomData(data) {
    try {
        const randomContainer = $('#random-container');

        const blockQuote = $('<blockquote>').addClass('blockquote')
        const dataQuote = $('<h2>').text(data.quote)

        blockQuote.append(dataQuote)

        const figCaption = $('<figcaption>').addClass('blockquote-footer text-end').text(data.name)

        randomContainer.append(blockQuote, figCaption)
    } catch (error) {
        console.error("Error:", error);
    }
}

fetchAllData();
fetchRandomData();