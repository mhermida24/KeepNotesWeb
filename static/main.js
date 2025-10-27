// -------- Event listeners ----------//

document.addEventListener('DOMContentLoaded', () => {

    /*if we dont wait for the DOM to be built we will get a js error in the event listeners */

    //create note card  button event listener
    const takeNoteButton = document.querySelector('.take-note-button');
    takeNoteButton.addEventListener('click', () => {
        createNote();
    });
    
    //create folder button event listener
    const addNewFolderButton = document.getElementById('add-folder-btn');
    if (addNewFolderButton) {
        addNewFolderButton.addEventListener('click', () => {
            createFolder();
        });
    }

    //save note button event listener
    const cardContainer = document.querySelector('.cards-container');
    cardContainer.addEventListener('click', (e) => {
        console.log('Click detected on:', e.target);
        if (e.target.classList.contains('save-btn')) {
            console.log('Save button clicked!');
            const noteCard = e.target.closest('.note-card');
            console.log('Note card found:', noteCard);
            createOrUpdateNote(noteCard);
        }
        if (e.target.classList.contains('delete-note-btn')) {
            const noteCard = e.target.closest('.note-card');
            deleteNoteCard(noteCard);
        }
    });
});


//----- Event listeners handler functions ---------//

// create note card HTML
function createNote(){
    const cardContainer = document.querySelector('.cards-container');
    const newNoteCard = document.createElement('div');
    newNoteCard.classList.add('note-card');
    newNoteCard.innerHTML = `
        <div class="note-header">
            <div class="note-title" contenteditable="true" data-placeholder="Title"></div>
            <button class="delete-note-btn">×</button>
        </div>
        <div class="note-content">
            <div class="note-body" contenteditable="true" data-placeholder="Take a note..."></div>
        </div>
        <button class="save-btn">Save</button>
    `;
    cardContainer.appendChild(newNoteCard);
    // No es necesario inicializar noteId, undefined es suficiente para la verificación
}

async function createOrUpdateNote(noteCard){
    const title = noteCard.querySelector('.note-title').textContent.trim();
    const body = noteCard.querySelector('.note-body').textContent.trim();
    
    console.log('Saving note...', { title, body, hasId: !!noteCard.dataset.noteId });
    
    try {
        // Si no tiene noteId, es nueva
        if (!noteCard.dataset.noteId) {
            console.log('Creating new note...');
            // CREATE
            const response = await fetch('/api/notes', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title, body })
            });
            
            console.log('Response status:', response.status);
            const data = await response.json();
            console.log('Full response data:', data);  // Ver qué devuelve exactamente
            
            noteCard.dataset.noteId = data.ID;  // ID en mayúscula
            console.log('Note created in DB with ID:', data.ID);
            
        } else {
            console.log('Updating existing note...');
            // UPDATE
            const noteId = noteCard.dataset.noteId;
            const response = await fetch(`/api/notes/${noteId}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title, body })
            });
            console.log('Note updated, status:', response.status);
        }
    } catch (error) {
        console.error('Error saving note:', error);
    }
}

async function deleteNoteCard(noteCard){
    const noteId = noteCard.dataset.noteId;
    
    // Si tiene ID, eliminar de la BD
    if (noteId) {
        try {
            const response = await fetch(`/api/notes/${noteId}`, {
                method: 'DELETE'
            });
            console.log('Note deleted from DB, status:', response.status);
        } catch (error) {
            console.error('Error deleting note:', error);
        }
    }
    
    // Eliminar del DOM
    noteCard.remove();
}

function createFolder(){

    console.log('Folder created');
}