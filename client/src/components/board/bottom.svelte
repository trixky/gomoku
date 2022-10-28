<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import OptionsStore from '../../stores/options';
	import { SHAPES } from '../../models/algo_options';
	import optionsStore from '../../stores/options';

	function handleProximityVisibility(e: any) {
		e.target.checked ? OptionsStore.showProximity() : OptionsStore.hideProximity();
	}

	function handleProximityEvolution(e: any) {
		OptionsStore.setProximityEvolution(e.target.checked);
	}

	function handleProximityShape(e: any) {
		optionsStore.setProximityShape(e.target.value);
	}

	function handleProximityRadius(e: any) {
		OptionsStore.setProximityRadius(e.target.value);
	}

	function handleProximityThreshold(e: any) {
		OptionsStore.setProximityThreshold(e.target.value);
	}

	function handleDepthMax(e: any) {
		OptionsStore.setDepthMax(e.target.value);
	}

	function handleDepthMin(e: any) {
		OptionsStore.setDepthMin(e.target.value);
	}

	function handleDepthPruning(e: any) {
		OptionsStore.setDepthPruning(e.target.checked);
	}
</script>

<!-- ========================= HTML -->
<div class="bottom-container">
	<div class="options-containers">
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
					<p>evolution</p>
					<input
						type="checkbox"
						checked={$OptionsStore.proximity.evolution}
						on:change={handleProximityEvolution}
						disabled
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
						value={$OptionsStore.proximity.radius}
					/>
				</div>
				<div class="option">
					<p>threshold</p>
					<input
						type="number"
						on:change={handleProximityThreshold}
						min={Config.options.proximity.threshold.min}
						max={Config.options.proximity.threshold.max}
						value={$OptionsStore.proximity.threshold}
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
					<p>reduction</p>
					<input
						type="checkbox"
						checked={$OptionsStore.depth.reduction}
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
	</div>
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.bottom-container {
		@apply flex justify-center mt-6;
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

	h3 {
		@apply inline-block opacity-30;
	}

	.option {
		@apply mr-3;
	}

	.option {
		@apply mr-3;
	}

	.option:last-child {
		@apply mr-0;
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
</style>
