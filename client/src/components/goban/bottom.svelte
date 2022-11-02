<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import OptionsStore from '../../stores/options';
	import { SHAPES } from '../../models/algo_options';
	import optionsStore from '../../stores/options';
	import AiStore from '../../stores/ai';

	let selected_ai = Config.options.ai.default;
	let advanced_mode = false;
	$: shape_neighboor = $OptionsStore.proximity.shape == SHAPES.neighbour;

	// ----------------------- handle ai
	function handleAi(e: any) {
		let ai: 'custom' | 'threepio' | 'ripley' | 'morty' | 'threepio' | 'deep' | 'gary' | 'joker' =
			e.target.value;

		if (ai != 'custom') OptionsStore.setAi(AiStore[ai].options);
		selected_ai = ai;
	}

	// ----------------------- handle time out
	function handleTimeOut(e: any) {
		selected_ai = 'custom';
		OptionsStore.setTimeOut(e.target.value);
	}

	// ----------------------- handle proximity
	function handleProximityVisibility(e: any) {
		selected_ai = 'custom';
		if (e.target.checked) {
			OptionsStore.showProximity();
			optionsStore.hideHeuristics();
		} else {
			OptionsStore.hideProximity();
		}
	}

	function handleProximityShape(e: any) {
		selected_ai = 'custom';
		optionsStore.setProximityShape(e.target.value);
	}

	function handleProximityRadius(e: any) {
		selected_ai = 'custom';
		OptionsStore.setProximityRadius(e.target.value);
	}

	function handleProximityThreshold(e: any) {
		selected_ai = 'custom';
		OptionsStore.setProximityThreshold(e.target.value);
	}

	function handleProximityEvolution(e: any) {
		OptionsStore.setProximityEvolution(e.target.checked);
	}

	function handleProximityReduction(e: any) {
		selected_ai = 'custom';
		OptionsStore.setProximityReduction(e.target.checked);
	}

	// ----------------------- handle depth
	function handleDepthMax(e: any) {
		selected_ai = 'custom';
		OptionsStore.setDepthMax(e.target.value);
	}

	function handleDepthMin(e: any) {
		selected_ai = 'custom';
		OptionsStore.setDepthMin(e.target.value);
	}

	function handleDepthPruning(e: any) {
		selected_ai = 'custom';
		OptionsStore.setDepthPruning(e.target.checked);
	}
	// ----------------------- handle width
	function handleWidthMultiThreading(e: any) {
		selected_ai = 'custom';
		OptionsStore.setWidthMultiThreading(e.target.checked);
	}
	function handleWidthPruning(e: any) {
		OptionsStore.setWidthPruning(e.target.checked);
	}
	function handleWidthMax(e: any) {
		OptionsStore.setWidthMax(e.target.value);
	}
	// ----------------------- handle heuristics
	function handleHeuristicsVisibility(e: any) {
		selected_ai = 'custom';
		if (e.target.checked) {
			OptionsStore.showHeuristics();
			optionsStore.hideProximity();
		} else {
			OptionsStore.hideHeuristics();
		}
	}

	function handleHeuristicsPotentialAlignement(e: any) {
		selected_ai = 'custom';
		OptionsStore.setHeuristicsPotentialAlignement(e.target.value);
	}

	function handleHeuristicsAlignement(e: any) {
		selected_ai = 'custom';
		OptionsStore.setHeuristicsAlignement(e.target.value);
	}

	function handleHeuristicsPotentialCapture(e: any) {
		selected_ai = 'custom';
		OptionsStore.setHeuristicsPotentialCapture(e.target.value);
	}

	function handleHeuristicsCapture(e: any) {
		selected_ai = 'custom';
		OptionsStore.setHeuristicsCapture(e.target.value);
	}

	function handleHeuristicsRandom(e: any) {
		selected_ai = 'custom';
		OptionsStore.setHeuristicsRandom(e.target.value);
	}
</script>

<!-- ========================= HTML -->
<div class="bottom-container">
	<div class="options-containers">
		<div class="options-form">
			<h3>Parameters</h3>
			<div class="options">
				<div class="option">
					<p>advanced</p>
					<input type="checkbox" bind:checked={advanced_mode} />
				</div>
				<div class="option">
					<p>time out</p>
					<input
						type="number"
						on:change={handleTimeOut}
						min={Config.options.time_out.min}
						max={Config.options.time_out.max}
						value={$OptionsStore.time_out}
					/>
				</div>
				<div class="option">
					<p>AI</p>
					<select value={selected_ai} on:change={handleAi}>
						{#each Object.entries(AiStore) as ai}
							<option value={ai[0]}>
								{ai[0]}
							</option>
						{/each}
					</select>
				</div>
			</div>
		</div>
		{#if advanced_mode}
			<div class="options-form">
				<h3>Proximity</h3>
				<div class="options">
					<div class="option">
						<p>show</p>
						<input
							type="checkbox"
							checked={$OptionsStore.proximity.show}
							on:change={handleProximityVisibility}
						/>
					</div>
					<div class="option">
						<p>reduction</p>
						<input
							type="checkbox"
							checked={$OptionsStore.proximity.reduction}
							on:change={handleProximityReduction}
							disabled
						/>
					</div>
					<div class="option">
						<p>evolution</p>
						<input
							type="checkbox"
							checked={$OptionsStore.proximity.evolution}
							on:change={handleProximityEvolution}
						/>
					</div>
				</div>
			</div>
			<div class="options-form">
				<h3>Shape</h3>
				<div class="options">
					<div class="option">
						<p>type</p>
						<select value={$OptionsStore.proximity.shape} on:change={handleProximityShape}>
							{#each Object.entries(SHAPES) as shape}
								<option value={shape[1]}>
									{shape[0]}
								</option>
							{/each}
						</select>
					</div>
					<div class="option">
						<p>radius</p>
						<input
							type="number"
							on:change={handleProximityRadius}
							min={Config.options.proximity.radius.min}
							max={Config.options.proximity.radius.max}
							value={shape_neighboor ? 1 : $OptionsStore.proximity.radius}
							disabled={shape_neighboor}
						/>
					</div>
					<div class="option">
						<p>threshold</p>
						<input
							type="number"
							on:change={handleProximityThreshold}
							min={Config.options.proximity.threshold.min}
							max={Config.options.proximity.threshold.max}
							value={shape_neighboor ? 1 : $OptionsStore.proximity.threshold}
							disabled={shape_neighboor}
						/>
					</div>
				</div>
			</div>
			<div class="options-form">
				<h3>Width</h3>
				<div class="options">
					<div class="option">
						<p>multi-threading</p>
						<input
							type="checkbox"
							checked={$OptionsStore.width.multi_threading}
							on:change={handleWidthMultiThreading}
						/>
					</div>
					<div class="option">
						<p>pruning</p>
						<input
							type="checkbox"
							checked={$OptionsStore.width.pruning}
							on:change={handleWidthPruning}
							disabled
						/>
					</div>
					<div class="option">
						<p>max</p>
						<input
							type="number"
							on:change={handleWidthMax}
							min={Config.options.width.max.min}
							max={Config.options.width.max.max}
							value={$OptionsStore.width.max}
						/>
					</div>
				</div>
			</div>
			<div class="options-form">
				<h3>Depth</h3>
				<div class="options">
					<div class="option">
						<p>pruning</p>
						<input
							type="checkbox"
							checked={$OptionsStore.depth.pruning}
							on:change={handleDepthPruning}
							disabled
						/>
					</div>
					<div class="option">
						<p>max</p>
						<input
							type="number"
							on:change={handleDepthMax}
							min={Config.options.depth.max.min}
							max={Config.options.depth.max.max}
							value={$OptionsStore.depth.max}
						/>
					</div>
					<div class="option">
						<p>min</p>
						<input
							type="number"
							on:change={handleDepthMin}
							min={Config.options.depth.min.min}
							max={Config.options.depth.min.max}
							value={$OptionsStore.depth.min}
							disabled={!$OptionsStore.depth.pruning}
						/>
					</div>
				</div>
			</div>
			<div class="options-form">
				<h3>Heuristics</h3>
				<div class="range-container">
					<div class="option">
						<p>show</p>
						<input
							type="checkbox"
							checked={$OptionsStore.heuristics.show}
							on:change={handleHeuristicsVisibility}
						/>
					</div>
					<div class="option">
						<p>potential alignement</p>
						<input
							type="range"
							min={Config.options.heuristics.values.min}
							max={Config.options.heuristics.values.max}
							on:input={handleHeuristicsPotentialAlignement}
						/>
						<p class="range-value">
							{$OptionsStore.heuristics.potential_alignement}/{Config.options.heuristics.values.max}
						</p>
					</div>
					<div class="option">
						<p>alignement</p>
						<input
							type="range"
							min={Config.options.heuristics.values.min}
							max={Config.options.heuristics.values.max}
							on:input={handleHeuristicsAlignement}
						/>
						<p class="range-value">
							{$OptionsStore.heuristics.alignement}/{Config.options.heuristics.values.max}
						</p>
					</div>
					<div class="option">
						<p>potential capture</p>
						<input
							type="range"
							min={Config.options.heuristics.values.min}
							max={Config.options.heuristics.values.max}
							on:input={handleHeuristicsPotentialCapture}
						/>
						<p class="range-value">
							{$OptionsStore.heuristics.potential_capture}/{Config.options.heuristics.values.max}
						</p>
					</div>
					<div class="option">
						<p>capture</p>
						<input
							type="range"
							min={Config.options.heuristics.values.min}
							max={Config.options.heuristics.values.max}
							on:input={handleHeuristicsCapture}
						/>
						<p class="range-value">
							{$OptionsStore.heuristics.capture}/{Config.options.heuristics.values.max}
						</p>
					</div>
					<div class="option">
						<p>random</p>
						<input
							type="range"
							min={Config.options.heuristics.values.min}
							max={Config.options.heuristics.values.max}
							on:input={handleHeuristicsRandom}
						/>
						<p class="range-value">
							{$OptionsStore.heuristics.random}/{Config.options.heuristics.values.max}
						</p>
					</div>
				</div>
			</div>
		{/if}
	</div>
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.bottom-container {
		@apply flex justify-center mt-6 mb-10;
	}

	.options-containers {
		@apply w-full;
	}

	.options-form {
		@apply flex justify-between w-full mt-3;
	}

	.options {
		@apply inline-flex justify-end w-full;
	}

	.range-container {
		@apply flex flex-col justify-end content-end align-bottom items-end;
	}

	h3 {
		@apply inline-block w-[110px] opacity-30;
	}

	.option {
		@apply mr-3;
	}

	.options > .option:last-child {
		@apply mr-0;
	}

	.range-container > .option {
		@apply mr-1;
	}

	.option > p {
		@apply inline-block;
	}

	input {
		@apply mx-1;
	}

	input[type='checkbox'] {
		@apply align-middle;
	}

	input[type='number'] {
		@apply align-middle rounded-sm pl-1;
		border: solid 1px black;
	}

	select {
		@apply bg-white rounded-sm mx-1;
		border: solid 1px black;
	}

	input[type='range'] {
		@apply w-48 my-0 mx-2 translate-y-[6px];
	}

	.range-value {
		@apply w-10 text-right;
	}
</style>
