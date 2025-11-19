<script>
	import ScoreEntry from "$lib/scoreEntry.svelte";
    let showScoreEntry = false;
    let title = "Video game high score server";
</script>

<h1 class = "title">{title}</h1>

<button class="submit-score-btn" on:click={() => showScoreEntry = true}>
    Submit Score
</button>

<ScoreEntry bind:isOpen={showScoreEntry} />

<style>
        .title {
        text-align: center;
        margin-top: 2rem;
        font-size: 2.5rem;
        color: #333;
        font: sans-serif;
    }
    .submit-score-btn {
        background: #005503;
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 16px;
        margin-top: 20px;
    }

    

    .submit-score-btn:hover {
        background: #45a049;
    }

    .button-container {
        display: flex;
        gap: 10px;
        justify-content: center;
        margin-top: 20px;
    }

    .view-scores-btn {
        background: #0066cc;
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 16px;
    }

    .view-scores-btn:hover {
        background: #0052a3;
    }
</style>