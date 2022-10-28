<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import OptionsStore from '../../stores/options';
	import { SHAPES } from '../../models/algo_options';
	import optionsStore from '../../stores/options';

	function handleProximityVisibility(e: any) {
		e.target.checked ? OptionsStore.show() : OptionsStore.hide();
	}

	function handleProximityShape(e: any) {
		const shape_id = parseInt(e.target.value);

		if (!isNaN(shape_id)) {
			optionsStore.setShape(shape_id);
		}
	}

	function handleProximityRadius(e: any) {
		OptionsStore.setRadius(e.target.value);
	}

	function handleProximityThreshold(e: any) {
		OptionsStore.setThreshold(e.target.value);
	}
</script>

<!-- ========================= HTML -->
<div class="bottom-container">
	<div class="w-full">
		<div class="options-containers">
			<h3>Proximity</h3>
			<div class="options">
				<div class="option">
					<p>show</p>
					<input type="checkbox" on:change={handleProximityVisibility} />
				</div>
				<div class="option">
					<p>shape</p>
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
	</div>
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.bottom-container {
		@apply flex justify-center mt-6;
	}

	.options-containers {
		@apply flex justify-between w-full;
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
