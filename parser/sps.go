package parser

import (
	"fmt"

	"github.com/ibbbpbbbp/H264analyzer/bits"
	"github.com/ibbbpbbbp/gobits"
)

const (
	Extended_SAR = 255
)

func parseHRD(hrd *hrd, bs *gobits.BitStream, prefix string) {
	prefix = prefix + "hrd."

	hrd.cpb_cnt_minus1 = bits.UE(bs)
	fmt.Printf("%scpb_cnt_minus1: %d\n", prefix, hrd.cpb_cnt_minus1)
	hrd.bit_rate_scale = bits.U(bs, 4)
	fmt.Printf("%sbit_rate_scale: %d\n", prefix, hrd.bit_rate_scale)
	hrd.cpb_size_scale = bits.U(bs, 4)
	fmt.Printf("%scpb_size_scale: %d\n", prefix, hrd.cpb_size_scale)
	hrd.bit_rate_value_minus1 = make([]int, 0, hrd.cpb_cnt_minus1+1)
	hrd.cpb_size_value_minus1 = make([]int, 0, hrd.cpb_cnt_minus1+1)
	hrd.cbr_flag = make([]int, 0, hrd.cpb_cnt_minus1+1)
	for schedShelIdx := 0; schedShelIdx <= hrd.cpb_cnt_minus1; schedShelIdx++ {
		hrd.bit_rate_value_minus1[schedShelIdx] = bits.UE(bs)
		fmt.Printf("%sbit_rate_value_minus1[%d]: %d\n", prefix, schedShelIdx, hrd.bit_rate_value_minus1)
		hrd.cpb_size_value_minus1[schedShelIdx] = bits.UE(bs)
		fmt.Printf("%scpb_size_value_minus1[%d]: %d\n", prefix, schedShelIdx, hrd.cpb_size_value_minus1)
		hrd.cbr_flag[schedShelIdx] = bits.U(bs, 1)
		fmt.Printf("%scbr_flag[%d]: %d\n", prefix, schedShelIdx, hrd.cbr_flag)
	}
	hrd.initial_cpb_removal_delay_length_minus1 = bits.U(bs, 5)
	fmt.Printf("%sinitial_cpb_removal_delay_length_minus1: %d\n", prefix, hrd.initial_cpb_removal_delay_length_minus1)
	hrd.cpb_removal_delay_length_minus1 = bits.U(bs, 5)
	fmt.Printf("%scpb_removal_delay_length_minus1: %d\n", prefix, hrd.cpb_removal_delay_length_minus1)
	hrd.dpb_output_delay_length_minus1 = bits.U(bs, 5)
	fmt.Printf("%sdpb_output_delay_length_minus1: %d\n", prefix, hrd.dpb_output_delay_length_minus1)
	hrd.time_offset_length = bits.U(bs, 5)
	fmt.Printf("%stime_offset_length: %d\n", prefix, hrd.time_offset_length)
}

func parseVUI(sps *sps, bs *gobits.BitStream, prefix string) {
	prefix = prefix + "vui."

	sps.vui.aspect_ratio_info_present_flag = bits.U(bs, 1)
	fmt.Printf("%saspect_ratio_info_present_flag: %d\n", prefix, sps.vui.aspect_ratio_info_present_flag)
	if sps.vui.aspect_ratio_info_present_flag == 1 {
		sps.vui.aspect_ratio_idc = bits.U(bs, 8)
		fmt.Printf("%saspect_ratio_idc: %d\n", prefix, sps.vui.aspect_ratio_idc)
		if sps.vui.aspect_ratio_idc == Extended_SAR {
			sps.vui.sar_width = bits.U(bs, 16)
			fmt.Printf("%ssar_width: %d\n", prefix, sps.vui.sar_width)
			sps.vui.sar_height = bits.U(bs, 16)
			fmt.Printf("%ssar_height: %d\n", prefix, sps.vui.sar_height)
		}
		sps.vui.overscan_info_present_flag = bits.U(bs, 1)
		fmt.Printf("%soverscan_info_present_flag: %d\n", prefix, sps.vui.overscan_info_present_flag)
		if sps.vui.overscan_info_present_flag == 1 {
			sps.vui.overscan_appropriate_flag = bits.U(bs, 1)
			fmt.Printf("%soverscan_appropriate_flag: %d\n", prefix, sps.vui.overscan_appropriate_flag)
		}
		sps.vui.video_signal_type_present_flag = bits.U(bs, 1)
		fmt.Printf("%svideo_signal_type_present_flag: %d\n", prefix, sps.vui.video_signal_type_present_flag)
		if sps.vui.video_signal_type_present_flag == 1 {
			sps.vui.video_format = bits.U(bs, 3)
			fmt.Printf("%svideo_format: %d\n", prefix, sps.vui.video_format)
			sps.vui.video_full_range_flag = bits.U(bs, 1)
			fmt.Printf("%svideo_full_range_flag: %d\n", prefix, sps.vui.video_full_range_flag)
			sps.vui.colour_description_present_flag = bits.U(bs, 1)
			fmt.Printf("%scolour_description_present_flag: %d\n", prefix, sps.vui.colour_description_present_flag)
			if sps.vui.colour_description_present_flag == 1 {
				sps.vui.colour_primaries = bits.U(bs, 8)
				fmt.Printf("%scolour_primaries: %d\n", prefix, sps.vui.colour_primaries)
				sps.vui.transfer_characteristics = bits.U(bs, 8)
				fmt.Printf("%stransfer_characteristics: %d\n", prefix, sps.vui.transfer_characteristics)
				sps.vui.matrix_coefficients = bits.U(bs, 8)
				fmt.Printf("%smatrix_coefficients: %d\n", prefix, sps.vui.matrix_coefficients)
			}
		}
		sps.vui.chroma_loc_info_present_flag = bits.U(bs, 1)
		fmt.Printf("%schroma_loc_info_present_flag: %d\n", prefix, sps.vui.chroma_loc_info_present_flag)
		if sps.vui.chroma_loc_info_present_flag == 1 {
			sps.vui.chroma_sample_loc_type_top_field = bits.UE(bs)
			fmt.Printf("%schroma_sample_loc_type_top_field: %d\n", prefix, sps.vui.chroma_sample_loc_type_top_field)
			sps.vui.chroma_sample_loc_type_bottom_field = bits.UE(bs)
			fmt.Printf("%schroma_sample_loc_type_bottom_field: %d\n", prefix, sps.vui.chroma_sample_loc_type_bottom_field)
		}
		sps.vui.timing_info_present_flag = bits.U(bs, 1)
		fmt.Printf("%stiming_info_present_flag: %d\n", prefix, sps.vui.timing_info_present_flag)
		if sps.vui.timing_info_present_flag == 1 {
			sps.vui.num_units_in_tick = bits.U(bs, 32)
			fmt.Printf("%snum_units_in_tick: %d\n", prefix, sps.vui.num_units_in_tick)
			sps.vui.time_scale = bits.U(bs, 32)
			fmt.Printf("%stime_scale: %d\n", prefix, sps.vui.time_scale)
			sps.vui.fixed_frame_rate_flag = bits.U(bs, 1)
			fmt.Printf("%sfixed_frame_rate_flag: %d\n", prefix, sps.vui.fixed_frame_rate_flag)
		}
		sps.vui.nal_hrd_parameters_present_flag = bits.U(bs, 1)
		fmt.Printf("%snal_hrd_parameters_present_flag: %d\n", prefix, sps.vui.nal_hrd_parameters_present_flag)
		sps.vui.vcl_hrd_parameters_present_flag = bits.U(bs, 1)
		fmt.Printf("%svcl_hrd_parameters_present_flag: %d\n", prefix, sps.vui.vcl_hrd_parameters_present_flag)
		if sps.vui.nal_hrd_parameters_present_flag == 1 {
			parseHRD(&sps.nal_hrd, bs, prefix+"nal_")
		}
		if sps.vui.vcl_hrd_parameters_present_flag == 1 {
			parseHRD(&sps.vcl_hrd, bs, prefix+"vcl_")
		}
		if sps.vui.nal_hrd_parameters_present_flag == 1 || sps.vui.vcl_hrd_parameters_present_flag == 1 {
			sps.vui.low_delay_hrd_flag = bits.U(bs, 1)
			fmt.Printf("%slow_delay_hrd_flag: %d\n", prefix, sps.vui.low_delay_hrd_flag)
		}
		sps.vui.pic_struct_present_flag = bits.U(bs, 1)
		fmt.Printf("%spic_struct_present_flag: %d\n", prefix, sps.vui.pic_struct_present_flag)
		sps.vui.bitstream_restriction_flag = bits.U(bs, 1)
		fmt.Printf("%sbitstream_restriction_flag: %d\n", prefix, sps.vui.bitstream_restriction_flag)
		if sps.vui.bitstream_restriction_flag == 1 {
			sps.vui.motion_vectors_over_pic_boundaries_flag = bits.U(bs, 1)
			fmt.Printf("%smotion_vectors_over_pic_boundaries_flag: %d\n", prefix, sps.vui.motion_vectors_over_pic_boundaries_flag)
			sps.vui.max_bytes_per_pic_denom = bits.UE(bs)
			fmt.Printf("%smax_bytes_per_pic_denom: %d\n", prefix, sps.vui.max_bytes_per_pic_denom)
			sps.vui.max_bits_per_mb_denom = bits.UE(bs)
			fmt.Printf("%smax_bits_per_mb_denom: %d\n", prefix, sps.vui.max_bits_per_mb_denom)
			sps.vui.log2_max_mv_length_horizontal = bits.UE(bs)
			fmt.Printf("%slog2_max_mv_length_horizontal: %d\n", prefix, sps.vui.log2_max_mv_length_horizontal)
			sps.vui.log2_max_mv_length_vertical = bits.UE(bs)
			fmt.Printf("%slog2_max_mv_length_vertical: %d\n", prefix, sps.vui.log2_max_mv_length_vertical)
			sps.vui.max_num_reorder_frames = bits.UE(bs)
			fmt.Printf("%smax_num_reorder_frames: %d\n", prefix, sps.vui.max_num_reorder_frames)
			sps.vui.max_dec_frame_buffering = bits.UE(bs)
			fmt.Printf("%smax_dec_frame_buffering: %d\n", prefix, sps.vui.max_dec_frame_buffering)
		}
	}
}

func parseSPS(h264 *H264Stream, bs *gobits.BitStream, prefix string) {
	prefix = prefix + "sps."
	sps := &sps{}

	sps.profile_idc = bits.U(bs, 8)
	fmt.Printf("%sprofile_idc: %d\n", prefix, sps.profile_idc)
	sps.constraint_set0_flag = bits.U(bs, 1)
	fmt.Printf("%sconstraint_set0_flag: %d\n", prefix, sps.constraint_set0_flag)
	sps.constraint_set1_flag = bits.U(bs, 1)
	fmt.Printf("%sconstraint_set1_flag: %d\n", prefix, sps.constraint_set1_flag)
	sps.constraint_set2_flag = bits.U(bs, 1)
	fmt.Printf("%sconstraint_set2_flag: %d\n", prefix, sps.constraint_set2_flag)
	sps.constraint_set3_flag = bits.U(bs, 1)
	fmt.Printf("%sconstraint_set3_flag: %d\n", prefix, sps.constraint_set3_flag)
	sps.constraint_set4_flag = bits.U(bs, 1)
	fmt.Printf("%sconstraint_set4_flag: %d\n", prefix, sps.constraint_set4_flag)
	sps.constraint_set5_flag = bits.U(bs, 1)
	fmt.Printf("%sconstraint_set5_flag: %d\n", prefix, sps.constraint_set5_flag)
	sps.reserved_zero_2bits = bits.U(bs, 2)
	fmt.Printf("%sreserved_zero_2bits: %d\n", prefix, sps.reserved_zero_2bits)
	sps.level_idc = bits.U(bs, 8)
	fmt.Printf("%slevel_idc: %d\n", prefix, sps.level_idc)
	sps.seq_parameter_set_id = bits.UE(bs)
	fmt.Printf("%sseq_parameter_set_id: %d\n", prefix, sps.seq_parameter_set_id)
	if sps.profile_idc == 100 || sps.profile_idc == 110 || sps.profile_idc == 122 || sps.profile_idc == 244 ||
		sps.profile_idc == 44 || sps.profile_idc == 83 || sps.profile_idc == 86 || sps.profile_idc == 118 ||
		sps.profile_idc == 128 || sps.profile_idc == 138 || sps.profile_idc == 139 || sps.profile_idc == 134 || sps.profile_idc == 135 {
		sps.chroma_format_idc = bits.UE(bs)
		fmt.Printf("%schroma_format_idc: %d\n", prefix, sps.chroma_format_idc)
		if sps.chroma_format_idc == 3 {
			sps.separate_colour_plane_flag = bits.U(bs, 1)
			fmt.Printf("%sseparate_colour_plane_flag: %d\n", prefix, sps.separate_colour_plane_flag)
		}
		sps.bit_depth_luma_minus8 = bits.UE(bs)
		fmt.Printf("%sbit_depth_luma_minus8: %d\n", prefix, sps.bit_depth_luma_minus8)
		sps.bit_depth_chroma_minus8 = bits.UE(bs)
		fmt.Printf("%sbit_depth_chroma_minus8: %d\n", prefix, sps.bit_depth_chroma_minus8)
		sps.qpprime_y_zero_transform_bypass_flag = bits.U(bs, 1)
		fmt.Printf("%sqpprime_y_zero_transform_bypass_flag: %d\n", prefix, sps.qpprime_y_zero_transform_bypass_flag)
		sps.seq_scaling_matrix_present_flag = bits.U(bs, 1)
		fmt.Printf("%sseq_scaling_matrix_present_flag: %d\n", prefix, sps.seq_scaling_matrix_present_flag)
		if sps.seq_scaling_matrix_present_flag == 1 {
			scaling_list_count := 12
			if sps.chroma_format_idc != 3 {
				scaling_list_count = 8
			}
			sps.seq_scaling_list_present_flag = make([]int, 0, scaling_list_count)
			for i := 0; i < scaling_list_count; i++ {
				sps.seq_scaling_list_present_flag[i] = bits.U(bs, 1)
				fmt.Printf("%sseq_scaling_list_present_flag[%d]: %d\n", prefix, i, sps.seq_scaling_list_present_flag[i])
				if sps.seq_scaling_list_present_flag[i] == 1 {
					if i < 6 {
						parseScalingList(bs, 16, prefix)
					} else {
						parseScalingList(bs, 64, prefix)
					}
				}
			}
		}
	}
	sps.log2_max_frame_num_minus4 = bits.UE(bs)
	fmt.Printf("%slog2_max_frame_num_minus4: %d\n", prefix, sps.log2_max_frame_num_minus4)
	sps.pic_order_cnt_type = bits.UE(bs)
	fmt.Printf("%spic_order_cnt_type: %d\n", prefix, sps.pic_order_cnt_type)
	if sps.pic_order_cnt_type == 0 {
		sps.log2_max_pic_order_cnt_lsb_minus4 = bits.UE(bs)
		fmt.Printf("%slog2_max_pic_order_cnt_lsb_minus4: %d\n", prefix, sps.log2_max_pic_order_cnt_lsb_minus4)
	} else if sps.pic_order_cnt_type == 1 {
		sps.delta_pic_order_always_zero_flag = bits.U(bs, 1)
		fmt.Printf("%sdelta_pic_order_always_zero_flag: %d\n", prefix, sps.delta_pic_order_always_zero_flag)
		sps.offset_for_non_ref_pic = bits.SE(bs)
		fmt.Printf("%soffset_for_non_ref_pic: %d\n", prefix, sps.offset_for_non_ref_pic)
		sps.offset_for_top_to_bottom_field = bits.SE(bs)
		fmt.Printf("%soffset_for_top_to_bottom_field: %d\n", prefix, sps.offset_for_top_to_bottom_field)
		sps.num_ref_frames_in_pic_order_cnt_cycle = bits.UE(bs)
		fmt.Printf("%snum_ref_frames_in_pic_order_cnt_cycle: %d\n", prefix, sps.num_ref_frames_in_pic_order_cnt_cycle)
		sps.offset_for_ref_frame = make([]int, 0, sps.num_ref_frames_in_pic_order_cnt_cycle)
		for i := 0; i < sps.num_ref_frames_in_pic_order_cnt_cycle; i++ {
			sps.offset_for_ref_frame[i] = bits.SE(bs)
			fmt.Printf("%soffset_for_ref_frame[%d]: %d\n", prefix, i, sps.offset_for_ref_frame[i])
		}
	}
	sps.max_num_ref_frames = bits.UE(bs)
	fmt.Printf("%smax_num_ref_frames: %d\n", prefix, sps.max_num_ref_frames)
	sps.gaps_in_frame_num_value_allowed_flag = bits.U(bs, 1)
	fmt.Printf("%sgaps_in_frame_num_value_allowed_flag: %d\n", prefix, sps.gaps_in_frame_num_value_allowed_flag)
	sps.pic_width_in_mbs_minus1 = bits.UE(bs)
	fmt.Printf("%spic_width_in_mbs_minus1: %d\n", prefix, sps.pic_width_in_mbs_minus1)
	sps.pic_height_in_map_units_minus1 = bits.UE(bs)
	fmt.Printf("%spic_height_in_map_units_minus1: %d\n", prefix, sps.pic_height_in_map_units_minus1)
	sps.frame_mbs_only_flag = bits.U(bs, 1)
	fmt.Printf("%sframe_mbs_only_flag: %d\n", prefix, sps.frame_mbs_only_flag)
	if sps.frame_mbs_only_flag == 0 {
		sps.mb_adaptive_frame_field_flag = bits.U(bs, 1)
		fmt.Printf("%smb_adaptive_frame_field_flag: %d\n", prefix, sps.mb_adaptive_frame_field_flag)
	}
	sps.direct_8x8_inference_flag = bits.U(bs, 1)
	fmt.Printf("%sdirect_8x8_inference_flag: %d\n", prefix, sps.direct_8x8_inference_flag)
	sps.frame_cropping_flag = bits.U(bs, 1)
	fmt.Printf("%sframe_cropping_flag: %d\n", prefix, sps.frame_cropping_flag)
	if sps.frame_cropping_flag == 1 {
		sps.frame_crop_left_offset = bits.UE(bs)
		fmt.Printf("%sframe_crop_left_offset: %d\n", prefix, sps.frame_crop_left_offset)
		sps.frame_crop_right_offset = bits.UE(bs)
		fmt.Printf("%sframe_crop_right_offset: %d\n", prefix, sps.frame_crop_right_offset)
		sps.frame_crop_top_offset = bits.UE(bs)
		fmt.Printf("%sframe_crop_top_offset: %d\n", prefix, sps.frame_crop_top_offset)
		sps.frame_crop_bottom_offset = bits.UE(bs)
		fmt.Printf("%sframe_crop_bottom_offset: %d\n", prefix, sps.frame_crop_bottom_offset)
	}
	sps.vui_parameters_present_flag = bits.U(bs, 1)
	fmt.Printf("%svui_parameters_present_flag: %d\n", prefix, sps.vui_parameters_present_flag)
	if sps.vui_parameters_present_flag == 1 {
		parseVUI(sps, bs, prefix)
	}

	h264.spsTable[sps.seq_parameter_set_id] = sps
}
