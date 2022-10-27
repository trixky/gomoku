<!-- ========================= SCRIPT -->
<script lang="ts">
	import Config from '../../config';
	import OptionsStore from '../../stores/options';
	function handleProximityVisibility(e: any) {
		e.target.checked ? OptionsStore.show() : OptionsStore.hide();
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
	<div class="options">
		<div>
			<h3>Proximity</h3>
			<input type="checkbox" on:change={handleProximityVisibility} />
			<p>radius</p>
			<input
				type="number"
				on:change={handleProximityRadius}
				min={Config.options.proximity.radius.min}
				max={Config.options.proximity.radius.max}
				value={$OptionsStore.proximity.radius}
			/>
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

<!-- ========================= CSS -->
<style lang="postcss">
	.bottom-container {
		@apply flex justify-center mt-4;
	}

	.options {
		@apply flex;
	}

	.options > div > h3,
	p {
		@apply inline-block;
	}

	input[type='checkbox'] {
		@apply mx-2 align-middle;
	}

	input[type='number'] {
		@apply mx-2 align-middle rounded-sm pl-1;
		border: solid 1px black;
	}
</style>
