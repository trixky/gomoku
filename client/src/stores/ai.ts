import type OptionsModel from '../models/options';
import type AiModel from '../models/ai';
import Config from '../config';
import generateOptions from '../utils/generate_options';
import { SHAPES } from '../models/algo_options';

export default {
	custom: <AiModel>{
		options: generateOptions(),
		description:
			'Threepio is a humanoid robot, his playing style is very classic. He is a good opponent to train.'
	},
	threepio: <AiModel>{
		options: generateOptions(),
		description:
			'Threepio is a humanoid robot, his playing style is very classic. He is a good opponent to train.'
	},
	ripley: <AiModel>{
		options: <OptionsModel>{
			time_out: 1500,
			ai: 'ripley',
			depth: {
				max: 6,
				min: 3,
				pruning: true
			},
			width: {
				max: Config.options.width.max.default,
				pruning: true,
				multi_threading: Config.options.width.multi_threading.default
			},
			proximity: {
				radius: 1,
				threshold: 1,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.neighbour,
				reduction: Config.options.proximity.reduction.default
			},
			heuristics: {
				potential_alignement: 0,
				alignement: 0,
				potential_capture: 5,
				capture: 10,
				random: 2,
				show: Config.options.heuristics.show.default
			}
		},
		description:
			"Chased by a xenomorph, Ripley won't take the time to think about all the possibilities. she'll slip into the first air duct she finds and go as far as she can without looking back."
	},
	morty: <AiModel>{
		options: <OptionsModel>{
			time_out: 3000,
			ai: 'morty',
			depth: {
				max: 1,
				min: 0,
				pruning: false
			},
			width: {
				max: Config.options.width.max.default,
				pruning: false,
				multi_threading: Config.options.width.multi_threading.default
			},
			proximity: {
				radius: 1,
				threshold: 1,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.neighbour,
				reduction: Config.options.proximity.reduction.default
			},
			heuristics: {
				potential_alignement: 2,
				alignement: 3,
				potential_capture: 2,
				capture: 3,
				random: 3,
				show: Config.options.heuristics.show.default
			}
		},
		description:
			"Morty really sucks, he can't project himself beyond those feet. Sometimes it feels like he's playing randomly. He knows the rules."
	},
	deep: <AiModel>{
		options: <OptionsModel>{
			time_out: Config.options.time_out.max,
			ai: 'deep',
			depth: {
				max: 10,
				min: 2,
				pruning: true
			},
			width: {
				max: Config.options.width.max.default,
				pruning: true,
				multi_threading: true
			},
			proximity: {
				radius: 3,
				threshold: 3,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.star,
				reduction: true
			},
			heuristics: {
				potential_alignement: 1,
				alignement: 2,
				potential_capture: 4,
				capture: 10,
				random: 1,
				show: Config.options.heuristics.show.default
			}
		},
		description:
			'Deep is an enormous supercomputer, it will think deep and wide. you will can play your second move.. in another life, when your species may have disappeared. Be patient.'
	},
	gary: <AiModel>{
		options: <OptionsModel>{
			time_out: 5000,
			ai: 'gary',
			depth: {
				max: 5,
				min: 1,
				pruning: true
			},
			width: {
				max: Config.options.width.max.default,
				pruning: true,
				multi_threading: true
			},
			proximity: {
				radius: 2,
				threshold: 2,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.star,
				reduction: true
			},
			heuristics: {
				potential_alignement: 1,
				alignement: 2,
				potential_capture: 4,
				capture: 8,
				random: 2,
				show: Config.options.heuristics.show.default
			}
		},
		description:
			'Gary is a smart, caring and calm snail. He is quite an intellectual, spending most of his days levitating, teleporting, and reading poetry. He sees far, but moves slowly.'
	},
	joker: <AiModel>{
		options: <OptionsModel>{
			time_out: 4000,
			ai: 'joker',
			depth: {
				max: 4,
				min: 2,
				pruning: true
			},
			width: {
				max: Config.options.width.max.default,
				pruning: true,
				multi_threading: true
			},
			proximity: {
				radius: 3,
				threshold: 2,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.star,
				reduction: true
			},
			heuristics: {
				potential_alignement: 1,
				alignement: 2,
				potential_capture: 5,
				capture: 10,
				random: 3,
				show: Config.options.heuristics.show.default
			}
		},
		description:
			'Joker will take the time to look for moves that may seem innocent at first sight, but which can be devastating later.'
	}
};
