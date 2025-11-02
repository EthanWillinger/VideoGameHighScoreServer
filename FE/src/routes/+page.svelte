<script>
	import ScoreEntry from "$lib/scoreEntry.svelte";
    let showScoreEntry = false;
</script>


<button class="submit-score-btn" on:click={() => showScoreEntry = true}>
    Submit Score
</button>

<ScoreEntry bind:isOpen={showScoreEntry} />

<style>
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
</style>