<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import OptionsStore from '../../stores/options';
	import { SHAPES } from '../../models/algo_options';
	import optionsStore from '../../stores/options';
	import AiStore from '../../stores/ai';
	import LayersStore from '../../stores/layers';
	import LayersPercentageStore from '../../stores/layers_percentage';
	import LoadingStore from '../../stores/loading';

	let selected_ai = Config.options.ai.default;
	let advanced_mode = false;

	$: layers_percentage_unit = $OptionsStore.analyzer.percentage ? ' %' : '';

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
		selected_ai = 'custom';
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

	function handleDepthPruningPercentage(e: any) {
		selected_ai = 'custom';
		OptionsStore.setDepthPruningPercentage(e.target.value);
	}
	// ----------------------- handle width
	function handleWidthMultiThreading(e: any) {
		selected_ai = 'custom';
		OptionsStore.setWidthMultiThreading(e.target.checked);
	}
	function handleWidthPruning(e: any) {
		OptionsStore.setWidthPruning(e.target.checked);
	}
	function handleWidthPruningPercentage(e: any) {
		selected_ai = 'custom';
		OptionsStore.setWidthPruningPercentage(e.target.value);
	}
	function handleWidthMax(e: any) {
		selected_ai = 'custom';
		OptionsStore.setWidthMax(e.target.value);
	}
	// ----------------------- handle heuristics
	function handleHeuristicsVisibility(e: any) {
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
	// ----------------------- handle heuristics
	function handleSuspicionActivation(e: any) {
		selected_ai = 'custom';
		if (e.target.checked) {
			OptionsStore.activeSuspicion();
		} else {
			OptionsStore.disableSuspicion();
		}
	}
	function handleSuspicionRadius(e: any) {
		selected_ai = 'custom';
		OptionsStore.setSuspicionRadius(e.target.value);
	}
	// ----------------------- handle analyzer
	function handleAnalyzerLayered(e: any) {
		if (e.target.checked) {
			OptionsStore.activeAnalyzerLayered();
		} else {
			OptionsStore.disableAnalyzerLayered();
		}
	}
	function handleAnalyzerPercentage(e: any) {
		if (e.target.checked) {
			OptionsStore.activeAnalyzerPercentage();
		} else {
			OptionsStore.disableAnalyzerPercentage();
		}
	}
	function handleAnalyzerRounded(e: any) {
		if (e.target.checked) {
			OptionsStore.activeAnalyzerRounded();
		} else {
			OptionsStore.disableAnalyzerRounded();
		}
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
						step={Config.options.time_out.step}
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
		<div class="options-form">
			<div class="options">
				<div class="option">
					<p class="max-w-[400px] text-right italic text-sm">
						{@html Object.entries(AiStore).find((ai) => ai[0] === selected_ai)?.[1].description}
					</p>
				</div>
			</div>
		</div>
		{#if advanced_mode}
			<div class="options-form">
				<div class="options">
					<div class="option">
						<p>multi-threading</p>
						<input
							type="checkbox"
							checked={$OptionsStore.width.multi_threading}
							on:change={handleWidthMultiThreading}
						/>
					</div>
				</div>
			</div>
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
							class="max-w-[56px]"
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
						<p>pruning</p>
						<input
							type="checkbox"
							checked={$OptionsStore.width.pruning}
							on:change={handleWidthPruning}
						/>
						<input
							type="number"
							style={$OptionsStore.width.pruning
								? $OptionsStore.width.pruning_percentage >= 100
									? 'color: crimson;'
									: $OptionsStore.width.pruning_percentage >= 80
									? 'color: orangered;'
									: ''
								: ''}
							step={Config.options.width.pruning_percentage.step}
							on:change={handleWidthPruningPercentage}
							min={Config.options.width.pruning_percentage.min}
							max={Config.options.width.pruning_percentage.max}
							value={$OptionsStore.width.pruning_percentage}
							disabled={!$OptionsStore.width.pruning}
						/>
						<span>%</span>
					</div>
					<div class="option">
						<p>max</p>
						<input
							type="number"
							step={Config.options.width.max.step}
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
						<p>min</p>
						<input
							type="number"
							on:change={handleDepthMin}
							min={Config.options.depth.min.min}
							max={$OptionsStore.depth.max}
							value={$OptionsStore.depth.min}
							disabled={!$OptionsStore.depth.pruning && !$OptionsStore.width.pruning}
						/>
					</div>
					<div class="option">
						<p>pruning</p>
						<input
							type="checkbox"
							checked={$OptionsStore.depth.pruning}
							on:change={handleDepthPruning}
						/>
						<input
							type="number"
							style={$OptionsStore.depth.pruning
								? $OptionsStore.depth.pruning_percentage >= 300
									? 'color: crimson;'
									: $OptionsStore.depth.pruning_percentage >= 280
									? 'color: orangered;'
									: ''
								: ''}
							step={Config.options.depth.pruning_percentage.step}
							on:change={handleDepthPruningPercentage}
							min={Config.options.depth.pruning_percentage.min}
							max={Config.options.depth.pruning_percentage.max}
							value={$OptionsStore.depth.pruning_percentage}
							disabled={!$OptionsStore.depth.pruning}
						/>
						<span>%</span>
					</div>
					<div class="option">
						<p>max</p>
						<input
							type="number"
							on:change={handleDepthMax}
							min={$OptionsStore.depth.min}
							max={Config.options.depth.max.max}
							value={$OptionsStore.depth.max}
						/>
					</div>
				</div>
			</div>
			<div class="options-form">
				<h3>Heuristics</h3>
				<div class="options">
					<div class="option">
						<p>show</p>
						<input
							type="checkbox"
							checked={$OptionsStore.heuristics.show}
							on:change={handleHeuristicsVisibility}
						/>
					</div>
					<div class="option">
						<p>suspicion radius</p>
						<input
							type="checkbox"
							checked={$OptionsStore.suspicion.active}
							on:change={handleSuspicionActivation}
						/>
						<input
							type="number"
							on:change={handleSuspicionRadius}
							min={Config.options.suspicion.radius.min}
							max={Config.options.suspicion.radius.max}
							value={$OptionsStore.suspicion.radius}
							disabled={!$OptionsStore.suspicion.active}
						/>
					</div>
				</div>
			</div>
			<div class="options-form">
				<div class="range-container">
					<div class="option">
						<p>potential alignement</p>
						<input
							type="range"
							min={Config.options.heuristics.values.min}
							max={Config.options.heuristics.values.max}
							on:input={handleHeuristicsPotentialAlignement}
							value={$OptionsStore.heuristics.potential_alignement}
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
							value={$OptionsStore.heuristics.alignement}
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
							value={$OptionsStore.heuristics.potential_capture}
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
							value={$OptionsStore.heuristics.capture}
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
							value={$OptionsStore.heuristics.random}
						/>
						<p class="range-value">
							{$OptionsStore.heuristics.random}/{Config.options.heuristics.values.max}
						</p>
					</div>
				</div>
			</div>
			<div class="options-form">
				<h3>Analyzer</h3>
				<div class="options">
					<div class="option">
						<p>layered</p>
						<input
							type="checkbox"
							checked={$OptionsStore.analyzer.layered}
							on:change={handleAnalyzerLayered}
						/>
					</div>
					<div class="option">
						<p>percentage</p>
						<input
							type="checkbox"
							checked={$OptionsStore.analyzer.percentage}
							on:change={handleAnalyzerPercentage}
						/>
					</div>
					<div class="option">
						<p>rounded</p>
						<input
							type="checkbox"
							checked={$OptionsStore.analyzer.rounded}
							on:change={handleAnalyzerRounded}
							disabled={!$OptionsStore.analyzer.percentage}
						/>
					</div>
				</div>
			</div>
			<div class:loading={$LoadingStore} class="analyzer-grid">
				<p />
				<p class="label top">selected</p>
				<p class="label top">depth<br />pruning</p>
				<p class="label top">width<br />pruning</p>
				<p class="label top">cutted by<br />max width</p>
				<p class="label top">cutted by<br />time out</p>
				<p class="label top">saved by<br />min depth</p>
				<p class="label top total">total</p>
				{#each $OptionsStore.analyzer.percentage ? $LayersPercentageStore : $LayersStore as layer, index (index.toString() + layer.selected.toString())}
					{#if index === $LayersStore.length - 1 || $OptionsStore.analyzer.layered}
						<p class:total={index === $LayersStore.length - 1} class="label left">
							{index === $LayersStore.length - 1 ? 'total' : `layer ${index}`}
						</p>
						<p class:blue={layer.selected}>
							{layer.selected || '--'}{layers_percentage_unit}
						</p>
						<p class:orange={layer.pruned_in_depth}>
							{layer.pruned_in_depth || '--'}{layers_percentage_unit}
						</p>
						<p class:orange={layer.pruned_in_width}>
							{layer.pruned_in_width || '--'}{layers_percentage_unit}
						</p>
						<p class:red={layer.cutted_by_max_width}>
							{layer.cutted_by_max_width || '--'}{layers_percentage_unit}
						</p>
						<p class:red={layer.cutted_by_time_out}>
							{layer.cutted_by_time_out || '--'}{layers_percentage_unit}
						</p>
						<p class:green={layer.saved_by_min_depth}>
							{layer.saved_by_min_depth || '--'}{layers_percentage_unit}
						</p>
						<p class:yellow={layer.total}>{layer.total || '--'}</p>
					{/if}
				{/each}
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
		@apply flex justify-end w-full mt-3;
	}

	.options {
		@apply inline-flex justify-end w-full;
	}

	.range-container {
		@apply flex flex-col justify-end content-end items-end;
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
		@apply w-48 m-0 ml-3 translate-y-[6px];
	}

	.range-value {
		@apply w-10 text-right;
	}

	.analyzer-grid {
		@apply grid mt-[14px] transition-all duration-300;
		grid-template-columns: repeat(8, 1fr);
		grid-auto-rows: repeat(3, 1fr);
	}

	.analyzer-grid.loading {
		@apply opacity-40;
	}

	.analyzer-grid > p {
		@apply text-right px-[7px] py-[2px] text-xs;
		border-right: solid 1px black;
		border-bottom: solid 1px black;
	}

	.analyzer-grid > .blue {
		@apply bg-blue-300 bg-opacity-10;
	}

	.analyzer-grid > .orange {
		@apply bg-orange-300 bg-opacity-10;
	}

	.analyzer-grid > .red {
		@apply bg-red-300 bg-opacity-10;
	}

	.analyzer-grid > .green {
		@apply bg-green-300 bg-opacity-10;
	}

	.analyzer-grid > .yellow {
		@apply bg-yellow-300 bg-opacity-10;
	}

	.analyzer-grid > p.label {
		@apply italic bg-neutral-100;
	}

	.analyzer-grid > p.label.left {
		@apply text-center;
		border-left: solid 1px black;
	}

	.analyzer-grid > p.label.top {
		border-top: solid 1px black;
	}

	.analyzer-grid > p.label.total {
		@apply bg-neutral-500 bg-opacity-20;
	}
</style>
