<script lang="ts">
    export let imageFile: File | null = null;
    let imageUrl: string | null = null;

    $: {
        if (imageFile) {
            // Clean up previous URL if it exists
            if (imageUrl) {
                URL.revokeObjectURL(imageUrl);
            }
            imageUrl = URL.createObjectURL(imageFile);
        }
    }

    // Cleanup on component destruction
    import { onDestroy } from 'svelte';
    onDestroy(() => {
        if (imageUrl) {
            URL.revokeObjectURL(imageUrl);
        }
    });
</script>

<div class="preview-container">
    {#if imageUrl}
        <div class="image-preview">
            <img src={imageUrl} alt="Screenshot preview" />
        </div>
    {:else}
        <div class="placeholder">
            <span>No image selected</span>
        </div>
    {/if}
</div>

<style>
    .preview-container {
        margin-top: 1rem;
        width: 100%;
        border: 2px dashed #ccc;
        border-radius: 4px;
        overflow: hidden;
    }

    .image-preview {
        width: 100%;
        height: 200px;
        display: flex;
        justify-content: center;
        align-items: center;
        background: #f5f5f5;
    }

    .image-preview img {
        max-width: 100%;
        max-height: 100%;
        object-fit: contain;
    }

    .placeholder {
        width: 100%;
        height: 200px;
        display: flex;
        justify-content: center;
        align-items: center;
        background: #f5f5f5;
        color: #666;
    }
</style>
