package parser

import (
	"fmt"

	"github.com/ibbbpbbbp/H264analyzer/bits"
	"github.com/ibbbpbbbp/gobits"
)

func parseRefPicListMVCModification(slice_type int, bs *gobits.BitStream, prefix string) {
	if slice_type%5 != 2 && slice_type%5 != 4 {
		ref_pic_list_modification_flag_l0 := bits.U(bs, 1)
		fmt.Printf("%sref_pic_list_modification_flag_l0: %d\n", prefix, ref_pic_list_modification_flag_l0)
		if ref_pic_list_modification_flag_l0 == 1 {
			for {
				modification_of_pic_nums_idc := bits.UE(bs)
				fmt.Printf("%smodification_of_pic_nums_idc: %d\n", prefix, modification_of_pic_nums_idc)
				if modification_of_pic_nums_idc == 0 || modification_of_pic_nums_idc == 1 {
					abs_diff_pic_num_minus1 := bits.UE(bs)
					fmt.Printf("%sabs_diff_pic_num_minus1: %d\n", prefix, abs_diff_pic_num_minus1)
				} else if modification_of_pic_nums_idc == 2 {
					long_term_pic_num := bits.UE(bs)
					fmt.Printf("%slong_term_pic_num: %d\n", prefix, long_term_pic_num)
				} else if modification_of_pic_nums_idc == 4 || modification_of_pic_nums_idc == 5 {
					abs_diff_view_idx_minus1 := bits.UE(bs)
					fmt.Printf("%sabs_diff_view_idx_minus1: %d\n", prefix, abs_diff_view_idx_minus1)
				}
				if modification_of_pic_nums_idc == 3 {
					break
				}
			}
		}
	}
	if slice_type%5 == 1 {
		ref_pic_list_modification_flag_l1 := bits.U(bs, 1)
		fmt.Printf("%sref_pic_list_modification_flag_l1: %d\n", prefix, ref_pic_list_modification_flag_l1)
		if ref_pic_list_modification_flag_l1 == 1 {
			for {
				modification_of_pic_nums_idc := bits.UE(bs)
				fmt.Printf("%smodification_of_pic_nums_idc: %d\n", prefix, modification_of_pic_nums_idc)
				if modification_of_pic_nums_idc == 0 || modification_of_pic_nums_idc == 1 {
					abs_diff_pic_num_minus1 := bits.UE(bs)
					fmt.Printf("%sabs_diff_pic_num_minus1: %d\n", prefix, abs_diff_pic_num_minus1)
				} else if modification_of_pic_nums_idc == 2 {
					long_term_pic_num := bits.UE(bs)
					fmt.Printf("%slong_term_pic_num: %d\n", prefix, long_term_pic_num)
				} else if modification_of_pic_nums_idc == 4 || modification_of_pic_nums_idc == 5 {
					abs_diff_view_idx_minus1 := bits.UE(bs)
					fmt.Printf("%sabs_diff_view_idx_minus1: %d\n", prefix, abs_diff_view_idx_minus1)
				}
				if modification_of_pic_nums_idc == 3 {
					break
				}
			}
		}
	}
}

func parseRefPicListModification(slice_type int, bs *gobits.BitStream, prefix string) {
	if slice_type%5 != 2 && slice_type%5 != 4 {
		ref_pic_list_modification_flag_l0 := bits.U(bs, 1)
		fmt.Printf("%sref_pic_list_modification_flag_l0: %d\n", prefix, ref_pic_list_modification_flag_l0)
		if ref_pic_list_modification_flag_l0 == 1 {
			for {
				modification_of_pic_nums_idc := bits.UE(bs)
				fmt.Printf("%smodification_of_pic_nums_idc: %d\n", prefix, modification_of_pic_nums_idc)
				if modification_of_pic_nums_idc == 0 || modification_of_pic_nums_idc == 1 {
					abs_diff_pic_num_minus1 := bits.UE(bs)
					fmt.Printf("%sabs_diff_pic_num_minus1: %d\n", prefix, abs_diff_pic_num_minus1)
				} else if modification_of_pic_nums_idc == 2 {
					long_term_pic_num := bits.UE(bs)
					fmt.Printf("%slong_term_pic_num: %d\n", prefix, long_term_pic_num)
				}
				if modification_of_pic_nums_idc == 3 {
					break
				}
			}
		}
	}
	if slice_type%5 == 1 {
		ref_pic_list_modification_flag_l1 := bits.U(bs, 1)
		fmt.Printf("%sref_pic_list_modification_flag_l1: %d\n", prefix, ref_pic_list_modification_flag_l1)
		if ref_pic_list_modification_flag_l1 == 1 {
			for {
				modification_of_pic_nums_idc := bits.UE(bs)
				fmt.Printf("%smodification_of_pic_nums_idc: %d\n", prefix, modification_of_pic_nums_idc)
				if modification_of_pic_nums_idc == 0 || modification_of_pic_nums_idc == 1 {
					abs_diff_pic_num_minus1 := bits.UE(bs)
					fmt.Printf("%sabs_diff_pic_num_minus1: %d\n", prefix, abs_diff_pic_num_minus1)
				} else if modification_of_pic_nums_idc == 2 {
					long_term_pic_num := bits.UE(bs)
					fmt.Printf("%slong_term_pic_num: %d\n", prefix, long_term_pic_num)
				}
				if modification_of_pic_nums_idc == 3 {
					break
				}
			}
		}
	}
}

func parsePredWeightTable(sps *sps, slice_type int, num_ref_idx_l0_active_minus1 int, num_ref_idx_l1_active_minus1 int, bs *gobits.BitStream, prefix string) {
	// Depending on the value of separate_colour_plane_flag, the value of the variable ChromaArrayType is assigned as follows:
	// – If separate_colour_plane_flag is equal to 0, ChromaArrayType is set equal to chroma_format_idc.
	// – Otherwise (separate_colour_plane_flag is equal to 1), ChromaArrayType is set equal to 0.
	chromaArrayType := 0
	if sps.separate_colour_plane_flag == 0 {
		chromaArrayType = sps.chroma_format_idc
	}

	luma_log2_weight_denom := bits.UE(bs)
	fmt.Printf("%sluma_log2_weight_denom: %d\n", prefix, luma_log2_weight_denom)
	if chromaArrayType != 0 {
		chroma_log2_weight_denom := bits.UE(bs)
		fmt.Printf("%schroma_log2_weight_denom: %d\n", prefix, chroma_log2_weight_denom)
	}
	for i := 0; i <= num_ref_idx_l0_active_minus1; i++ {
		luma_weight_l0_flag := bits.U(bs, 1)
		fmt.Printf("%sluma_weight_l0_flag: %d\n", prefix, luma_weight_l0_flag)
		if luma_weight_l0_flag == 1 {
			luma_weight_l0 := bits.SE(bs)
			fmt.Printf("%sluma_weight_l0[%d]: %d\n", prefix, i, luma_weight_l0)
			luma_offset_l0 := bits.SE(bs)
			fmt.Printf("%sluma_offset_l0[%d]: %d\n", prefix, i, luma_offset_l0)
		}
		if chromaArrayType != 0 {
			chroma_weight_l0_flag := bits.U(bs, 1)
			fmt.Printf("%schroma_weight_l0_flag: %d\n", prefix, chroma_weight_l0_flag)
			if chroma_weight_l0_flag == 1 {
				for j := 0; j < 2; j++ {
					chroma_weight_l0 := bits.SE(bs)
					fmt.Printf("%schroma_weight_l0[%d][%d]: %d\n", prefix, i, j, chroma_weight_l0)
					chroma_offset_l0 := bits.SE(bs)
					fmt.Printf("%schroma_offset_l0[%d][%d]: %d\n", prefix, i, j, chroma_offset_l0)
				}
			}
		}
	}
	if slice_type%5 == 1 {
		for i := 0; i <= num_ref_idx_l1_active_minus1; i++ {
			luma_weight_l1_flag := bits.U(bs, 1)
			fmt.Printf("%sluma_weight_l1_flag: %d\n", prefix, luma_weight_l1_flag)
			if luma_weight_l1_flag == 1 {
				luma_weight_l1 := bits.SE(bs)
				fmt.Printf("%sluma_weight_l1[%d]: %d\n", prefix, i, luma_weight_l1)
				luma_offset_l1 := bits.SE(bs)
				fmt.Printf("%sluma_offset_l1[%d]: %d\n", prefix, i, luma_offset_l1)
			}
			if chromaArrayType != 0 {
				chroma_weight_l1_flag := bits.U(bs, 1)
				fmt.Printf("%schroma_weight_l1_flag: %d\n", prefix, chroma_weight_l1_flag)
				if chroma_weight_l1_flag == 1 {
					for j := 0; j < 2; j++ {
						chroma_weight_l1 := bits.SE(bs)
						fmt.Printf("%schroma_weight_l1[%d][%d]: %d\n", prefix, i, j, chroma_weight_l1)
						chroma_offset_l1 := bits.SE(bs)
						fmt.Printf("%schroma_offset_l1[%d][%d]: %d\n", prefix, i, j, chroma_offset_l1)
					}
				}
			}
		}
	}
}

func parseDecRefPicMarking(nal *nal, bs *gobits.BitStream, prefix string) {
	if nal.nal_unit_type == naluTypeOfIDR {
		no_output_of_prior_pics_flag := bits.U(bs, 1)
		fmt.Printf("%sno_output_of_prior_pics_flag: %d\n", prefix, no_output_of_prior_pics_flag)
		long_term_reference_flag := bits.U(bs, 1)
		fmt.Printf("%slong_term_reference_flag: %d\n", prefix, long_term_reference_flag)
	} else {
		adaptive_ref_pic_marking_mode_flag := bits.U(bs, 1)
		fmt.Printf("%sadaptive_ref_pic_marking_mode_flag: %d\n", prefix, adaptive_ref_pic_marking_mode_flag)
		if adaptive_ref_pic_marking_mode_flag == 1 {
			for {
				memory_management_control_operation := bits.UE(bs)
				fmt.Printf("%smemory_management_control_operation: %d\n", prefix, memory_management_control_operation)
				if memory_management_control_operation == 1 || memory_management_control_operation == 3 {
					difference_of_pic_nums_minus1 := bits.UE(bs)
					fmt.Printf("%sdifference_of_pic_nums_minus1: %d\n", prefix, difference_of_pic_nums_minus1)
				}
				if memory_management_control_operation == 2 {
					long_term_pic_num := bits.UE(bs)
					fmt.Printf("%slong_term_pic_num: %d\n", prefix, long_term_pic_num)
				}
				if memory_management_control_operation == 3 || memory_management_control_operation == 6 {
					long_term_frame_idx := bits.UE(bs)
					fmt.Printf("%slong_term_frame_idx: %d\n", prefix, long_term_frame_idx)
				}
				if memory_management_control_operation == 4 {
					max_long_term_frame_idx_plus1 := bits.UE(bs)
					fmt.Printf("%smax_long_term_frame_idx_plus1: %d\n", prefix, max_long_term_frame_idx_plus1)
				}
				if memory_management_control_operation == 0 {
					break
				}
			}
		}
	}
}

func parseSliceHeader(h264 *H264Stream, bs *gobits.BitStream, prefix string) {
	prefix = prefix + "sh."

	first_mb_in_slice := bits.UE(bs)
	fmt.Printf("%sfirst_mb_in_slice: %d\n", prefix, first_mb_in_slice)
	slice_type := bits.UE(bs)
	fmt.Printf("%sslice_type: %d\n", prefix, slice_type)
	pic_parameter_set_id := bits.UE(bs)
	fmt.Printf("%spic_parameter_set_id: %d\n", prefix, pic_parameter_set_id)

	nal := h264.nal
	pps, ok := h264.ppsTable[pic_parameter_set_id]
	if !ok {
		return
	}

	sps, ok := h264.spsTable[pps.seq_parameter_set_id]
	if !ok {
		return
	}

	if sps.separate_colour_plane_flag == 1 {
		colour_plane_id := bits.U(bs, 2)
		fmt.Printf("%scolour_plane_id: %d\n", prefix, colour_plane_id)
	}
	// frame_num is used as an identifier for pictures and shall be represented by log2_max_frame_num_minus4 + 4 bits in the bitstream.
	frame_num := bits.U(bs, byte(sps.log2_max_frame_num_minus4+4))
	fmt.Printf("%sframe_num: %d\n", prefix, frame_num)
	// field_pic_flag equal to 1 specifies that the slice is a slice of a coded field.
	// field_pic_flag equal to 0 specifies that the slice is a slice of a coded frame.
	// When field_pic_flag is not present it shall be inferred to be equal to 0.
	field_pic_flag := 0
	if sps.frame_mbs_only_flag == 0 {
		field_pic_flag = bits.U(bs, 1)
		fmt.Printf("%sfield_pic_flag: %d\n", prefix, field_pic_flag)
		if field_pic_flag == 1 {
			bottom_field_flag := bits.U(bs, 1)
			fmt.Printf("%sbottom_field_flag: %d\n", prefix, bottom_field_flag)
		}
	}
	if nal.nal_unit_type == naluTypeOfIDR {
		idr_pic_id := bits.UE(bs)
		fmt.Printf("%sidr_pic_id: %d\n", prefix, idr_pic_id)
	}
	if sps.pic_order_cnt_type == 0 {
		// The length of the pic_order_cnt_lsb syntax element is log2_max_pic_order_cnt_lsb_minus4 + 4 bits.
		pic_order_cnt_lsb := bits.U(bs, byte(sps.log2_max_pic_order_cnt_lsb_minus4+4))
		fmt.Printf("%spic_order_cnt_lsb: %d\n", prefix, pic_order_cnt_lsb)
		if pps.bottom_field_pic_order_in_frame_present_flag == 1 && field_pic_flag == 0 {
			delta_pic_order_cnt_bottom := bits.SE(bs)
			fmt.Printf("%sdelta_pic_order_cnt_bottom: %d\n", prefix, delta_pic_order_cnt_bottom)
		}
	}
	if sps.pic_order_cnt_type == 1 && sps.delta_pic_order_always_zero_flag == 0 {
		delta_pic_order_cnt := [2]int{}
		delta_pic_order_cnt[0] = bits.SE(bs)
		fmt.Printf("%sdelta_pic_order_cnt[0]: %d\n", prefix, delta_pic_order_cnt[0])
		if pps.bottom_field_pic_order_in_frame_present_flag == 1 && field_pic_flag == 0 {
			delta_pic_order_cnt[1] = bits.SE(bs)
			fmt.Printf("%sdelta_pic_order_cnt[1]: %d\n", prefix, delta_pic_order_cnt[1])
		}
	}
	if pps.redundant_pic_cnt_present_flag == 1 {
		redundant_pic_cnt := bits.UE(bs)
		fmt.Printf("%sredundant_pic_cnt: %d\n", prefix, redundant_pic_cnt)
	}
	if slice_type == sliceTypeOfB {
		direct_spatial_mv_pred_flag := bits.U(bs, 1)
		fmt.Printf("%sdirect_spatial_mv_pred_flag: %d\n", prefix, direct_spatial_mv_pred_flag)
	}
	// When the current slice is a P, SP, or B slice and num_ref_idx_l0_active_minus1 is not present, num_ref_idx_l0_active_minus1 shall be inferred to be equal to num_ref_idx_l0_default_active_minus1.
	// When the current slice is a B slice and num_ref_idx_l1_active_minus1 is not present, num_ref_idx_l1_active_minus1 shall be inferred to be equal to num_ref_idx_l1_default_active_minus1.
	num_ref_idx_l0_active_minus1 := pps.num_ref_idx_l0_default_active_minus1
	num_ref_idx_l1_active_minus1 := pps.num_ref_idx_l1_default_active_minus1
	if slice_type == sliceTypeOfP || slice_type == sliceTypeOfSP || slice_type == sliceTypeOfB {
		num_ref_idx_active_override_flag := bits.U(bs, 1)
		fmt.Printf("%snum_ref_idx_active_override_flag: %d\n", prefix, num_ref_idx_active_override_flag)
		if num_ref_idx_active_override_flag == 1 {
			num_ref_idx_l0_active_minus1 = bits.UE(bs)
			fmt.Printf("%snum_ref_idx_l0_active_minus1: %d\n", prefix, num_ref_idx_l0_active_minus1)
			if slice_type == sliceTypeOfB {
				num_ref_idx_l1_active_minus1 = bits.UE(bs)
				fmt.Printf("%snum_ref_idx_l1_active_minus1: %d\n", prefix, num_ref_idx_l1_active_minus1)
			}
		}
	}
	if nal.nal_unit_type == 20 || nal.nal_unit_type == 21 {
		parseRefPicListMVCModification(slice_type, bs, prefix)
	} else {
		parseRefPicListModification(slice_type, bs, prefix)
	}
	if pps.weighted_pred_flag == 1 && (slice_type == sliceTypeOfP || slice_type == sliceTypeOfB) ||
		(pps.weighted_bipred_idc == 1 && slice_type == sliceTypeOfB) {
		parsePredWeightTable(sps, slice_type, num_ref_idx_l0_active_minus1, num_ref_idx_l1_active_minus1, bs, prefix)
	}
	if nal.nal_ref_idc != 0 {
		parseDecRefPicMarking(nal, bs, prefix)
	}
	if pps.entropy_coding_mode_flag == 1 && slice_type != sliceTypeOfI && slice_type != sliceTypeOfSI {
		cabac_init_idc := bits.UE(bs)
		fmt.Printf("%scabac_init_idc: %d\n", prefix, cabac_init_idc)
	}
	slice_qp_delta := bits.SE(bs)
	fmt.Printf("%sslice_qp_delta: %d\n", prefix, slice_qp_delta)
	if slice_type == sliceTypeOfSP || slice_type == sliceTypeOfSI {
		if slice_type == sliceTypeOfSP {
			sp_for_switch_flag := bits.U(bs, 1)
			fmt.Printf("%ssp_for_switch_flag: %d\n", prefix, sp_for_switch_flag)
		}
		slice_qs_delta := bits.SE(bs)
		fmt.Printf("%sslice_qs_delta: %d\n", prefix, slice_qs_delta)
	}
	if pps.deblocking_filter_control_present_flag == 1 {
		disable_deblocking_filter_idc := bits.UE(bs)
		fmt.Printf("%sdisable_deblocking_filter_idc: %d\n", prefix, disable_deblocking_filter_idc)
		if disable_deblocking_filter_idc != 1 {
			slice_alpha_c0_offset_div2 := bits.SE(bs)
			fmt.Printf("%sslice_alpha_c0_offset_div2: %d\n", prefix, slice_alpha_c0_offset_div2)
			slice_beta_offset_div2 := bits.SE(bs)
			fmt.Printf("%sslice_beta_offset_div2: %d\n", prefix, slice_beta_offset_div2)
		}
	}
	if pps.num_slice_groups_minus1 > 0 && pps.slice_group_map_type >= 3 && pps.slice_group_map_type <= 5 {
		picWidthInMbs := sps.pic_width_in_mbs_minus1 + 1
		picHeightMapUnits := sps.pic_height_in_map_units_minus1 + 1
		picSizeInMapUnits := picWidthInMbs * picHeightMapUnits
		sliceGroupChangeRate := pps.slice_group_change_rate_minus1 + 1
		slice_group_change_cycle := bits.CeilLog2(picSizeInMapUnits/sliceGroupChangeRate + 1)
		fmt.Printf("%sslice_group_change_cycle: %d\n", prefix, slice_group_change_cycle)
	}
}
