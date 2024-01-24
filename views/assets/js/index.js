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
            const listItem = $('<div>').addClass('note-item__content');
            const title = $('<h3>').addClass('note-item__title').text(item.name);
            const date = $('<p>').addClass('note-item__date').text(item.quote);
            const body = $('<p>').addClass('note-item__body').text(item.created_at);

            listItem.append(title, date, body)

            const listAction = $('<div>').addClass('note-item__action');
            const buttonDelete = $('<button>').addClass('note-item__delete-button').text('Delete');
            const buttonEdit = $('<button>').addClass('note-item__archive-button').text('Edit');

            buttonDelete.on('click', function () {
                deleteData(item.id)
            })

            listAction.append(buttonDelete, buttonEdit)

            const wrapperItem = $('<div>').addClass('note-item');

            wrapperItem.append(listItem, listAction)

            dataContainer.append(wrapperItem)
        })
    } catch (error) {
        console.error("Error:", error);
    }
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