
<script>
    import GameScorePage from "$lib/gameScorePage.svelte";
    import ScoreEntry from "$lib/scoreEntry.svelte";
    let showScoreEntry = false;
    let showGameScores = true;
    let title = "Video game high score server";
</script>



{#if showGameScores}
    <GameScorePage />
{/if}

<ScoreEntry bind:isOpen={showScoreEntry} />

<style>

</style>