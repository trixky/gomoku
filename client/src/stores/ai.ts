import type OptionsModel from '../models/options';
import type AiModel from '../models/ai';
import Config from '../config';
import generateOptions from '../utils/generate_options';
import { SHAPES } from '../models/algo_options';

export default {
	threepio: <AiModel>{
		options: generateOptions(),
		description:
			'Threepio is a humanoid robot, his playing style is very classic.<br/>He will try not to keep you waiting too long.<br/>He is a good opponent to train.'
	},
	ripley: <AiModel>{
		options: <OptionsModel>{
			time_out: 1500,
			ai: 'ripley',
			depth: {
				max: 6,
				min: 3,
				pruning: true,
				pruning_percentage: 70
			},
			width: {
				max: Config.options.width.max.default,
				pruning: false,
				pruning_percentage: Config.options.width.pruning_percentage.default,
				multi_threading: Config.options.width.multi_threading.default
			},
			proximity: {
				radius: 1,
				threshold: 1,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.neighbour
			},
			heuristics: {
				alignement: 0,
				capture: 10,
				random: 2,
				show: Config.options.heuristics.show.default
			},
			suspicion: {
				active: Config.options.suspicion.active.default,
				radius: Config.options.suspicion.radius.default
			},
			analyzer: {
				layered: Config.options.analyzer.layered.default,
				percentage: Config.options.analyzer.percentage.default,
				rounded: Config.options.analyzer.rounded.default
			}
		},
		description:
			"Chased by a xenomorph, Ripley won't take the time to think about all the possibilities. She'll slip into the first air duct she finds and go as far as she can without looking back."
	},
	morty: <AiModel>{
		options: <OptionsModel>{
			time_out: 3000,
			ai: 'morty',
			depth: {
				max: 1,
				min: 0,
				pruning: false,
				pruning_percentage: Config.options.depth.pruning_percentage.default
			},
			width: {
				max: Config.options.width.max.default,
				min: 10,
				pruning: true,
				pruning_percentage: 70,
				multi_threading: Config.options.width.multi_threading.default
			},
			proximity: {
				radius: 1,
				threshold: 1,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.neighbour
			},
			heuristics: {
				alignement: 3,
				capture: 3,
				random: 3,
				show: Config.options.heuristics.show.default
			},
			suspicion: {
				active: true,
				radius: 6
			},
			analyzer: {
				layered: Config.options.analyzer.layered.default,
				percentage: Config.options.analyzer.percentage.default,
				rounded: Config.options.analyzer.rounded.default
			}
		},
		description:
			"Morty really sucks, he can't project himself beyond those feet.<br/>Sometimes it feels like he's playing randomly.<br/>He knows the rules."
	},
	deep: <AiModel>{
		options: <OptionsModel>{
			time_out: Config.options.time_out.max,
			ai: 'deep',
			depth: {
				max: 10,
				min: 2,
				pruning: false,
				pruning_percentage: Config.options.depth.pruning_percentage.default
			},
			width: {
				max: Config.options.width.max.default,
				pruning: false,
				pruning_percentage: Config.options.width.pruning_percentage.default,
				multi_threading: true
			},
			proximity: {
				radius: 3,
				threshold: 3,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.star
			},
			heuristics: {
				alignement: 2,
				capture: 10,
				random: 1,
				show: Config.options.heuristics.show.default
			},
			suspicion: {
				active: false,
				radius: Config.options.suspicion.radius.min
			},
			analyzer: {
				layered: Config.options.analyzer.layered.default,
				percentage: Config.options.analyzer.percentage.default,
				rounded: Config.options.analyzer.rounded.default
			}
		},
		description:
			'Deep is an enormous supercomputer, it will think deep and wide.<br/>You will can play your second move.. in another life, when your species may have disappeared. Be patient.'
	},
	gary: <AiModel>{
		options: <OptionsModel>{
			time_out: 5000,
			ai: 'gary',
			depth: {
				max: 5,
				min: 1,
				pruning: true,
				pruning_percentage: 30
			},
			width: {
				max: Config.options.width.max.default,
				pruning: true,
				pruning_percentage: 60,
				multi_threading: true
			},
			proximity: {
				radius: 2,
				threshold: 2,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.star
			},
			heuristics: {
				alignement: 2,
				capture: 8,
				random: 2,
				show: Config.options.heuristics.show.default
			},
			suspicion: {
				active: true,
				radius: 7
			},
			analyzer: {
				layered: Config.options.analyzer.layered.default,
				percentage: Config.options.analyzer.percentage.default,
				rounded: Config.options.analyzer.rounded.default
			}
		},
		description:
			'Gary is a smart, caring and calm snail.<br/>He is quite an intellectual, spending most of his days levitating, teleporting, and reading poetry. He sees far, but moves slowly.'
	},
	joker: <AiModel>{
		options: <OptionsModel>{
			time_out: 4000,
			ai: 'joker',
			depth: {
				max: 4,
				min: 2,
				pruning: true,
				pruning_percentage: 20
			},
			width: {
				max: Config.options.width.max.default,
				min: 20,
				pruning: true,
				pruning_percentage: 40,
				multi_threading: true
			},
			proximity: {
				radius: 3,
				threshold: 2,
				show: Config.options.proximity.show.default,
				evolution: true,
				shape: SHAPES.star
			},
			heuristics: {
				alignement: 2,
				capture: 10,
				random: 3,
				show: Config.options.heuristics.show.default
			},
			suspicion: {
				active: false,
				radius: Config.options.suspicion.radius.min
			},
			analyzer: {
				layered: Config.options.analyzer.layered.default,
				percentage: Config.options.analyzer.percentage.default,
				rounded: Config.options.analyzer.rounded.default
			}
		},
		description:
			'With Joker, everything is planned in advance.<br/>He will take the time to look for moves that may seem innocent at first sight, but which can be devastating later.'
	},
	custom: <AiModel>{
		options: generateOptions(),
		description:
			'Customize your own artificial intelligence.<br/>You can take the time to read the description of each parameter by hovering over them with your mouse.'
	}
};
