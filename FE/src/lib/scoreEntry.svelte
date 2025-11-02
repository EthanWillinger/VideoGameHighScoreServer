<script lang="ts">
    export let isOpen = false;

    function closeModal() {
        isOpen = false;
    }

    function handleOverlayClick(event: MouseEvent) {
        // Close only if clicking the overlay background, not the modal content
        if (event.target === event.currentTarget) {
            closeModal();
        }
    }

    function handleKeydown(event: KeyboardEvent) {
        if (event.key === 'Escape') {
            closeModal();
        }
    }
</script>

{#if isOpen}
    <div 
        class="modal-overlay" 
        on:click={handleOverlayClick}
        on:keydown={handleKeydown}
        role="dialog"
        aria-modal="true"
        tabindex="-1"
    >
        <div class="modal">
            <button class="close-button" on:click={closeModal}>&times;</button>
            <h2>Submit Score</h2>
            <form method="post" enctype="multipart/form-data">
                <div class="form-group">
                    <label for="scoreInput">Score:</label>
                    <input 
                        type="text" 
                        id="scoreInput" 
                        pattern="\d*"
                        inputmode="numeric"
                        placeholder="Enter your score"
                    >
                </div>
                <div class="form-group">
                    <label for="fileInput">Upload Evidence (Image/Video):</label>
                    <input 
                        type="file" 
                        id="fileInput" 
                        name="evidenceUpload" 
                        accept="image/png, image/jpeg, video/*"
                    />
                </div>
                <div class="form-group">
                    <button type="submit">Submit Score</button>
                </div>
            </form>
        </div>
    </div>
{/if}

<style>
    .modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;
    }

    .modal {
        background: white;
        padding: 2rem;
        border-radius: 8px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        position: relative;
        width: 90%;
        max-width: 500px;
    }

    .close-button {
        position: absolute;
        top: 10px;
        right: 10px;
        border: none;
        background: none;
        font-size: 24px;
        cursor: pointer;
        padding: 5px 10px;
    }

    .form-group {
        margin-bottom: 1rem;
    }

    label {
        display: block;
        margin-bottom: 0.5rem;
        font-weight: bold;
    }

    input {
        width: 100%;
        padding: 8px;
        font-size: 16px;
        box-sizing: border-box;
        border: 1px solid #ccc;
        border-radius: 4px;
    }

    button[type="submit"] {
        background: #4CAF50;
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 16px;
    }

    button[type="submit"]:hover {
        background: #45a049;
    }
</style>