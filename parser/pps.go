package parser

import (
	"fmt"

	"github.com/ibbbpbbbp/H264analyzer/bits"
	"github.com/ibbbpbbbp/gobits"
)

func parsePPS(h264 *H264Stream, bs *gobits.BitStream, prefix string) {
	prefix = prefix + "pps."
	pps := &pps{}

	pps.pic_parameter_set_id = bits.UE(bs)
	fmt.Printf("%spic_parameter_set_id: %d\n", prefix, pps.pic_parameter_set_id)
	pps.seq_parameter_set_id = bits.UE(bs)
	fmt.Printf("%sseq_parameter_set_id: %d\n", prefix, pps.seq_parameter_set_id)
	pps.entropy_coding_mode_flag = bits.U(bs, 1)
	fmt.Printf("%sentropy_coding_mode_flag: %d\n", prefix, pps.entropy_coding_mode_flag)
	pps.bottom_field_pic_order_in_frame_present_flag = bits.U(bs, 1)
	fmt.Printf("%sbottom_field_pic_order_in_frame_present_flag: %d\n", prefix, pps.bottom_field_pic_order_in_frame_present_flag)
	pps.num_slice_groups_minus1 = bits.UE(bs)
	fmt.Printf("%snum_slice_groups_minus1: %d\n", prefix, pps.num_slice_groups_minus1)
	if pps.num_slice_groups_minus1 > 0 {
		pps.slice_group_map_type = bits.UE(bs)
		fmt.Printf("%sslice_group_map_type: %d\n", prefix, pps.slice_group_map_type)
		if pps.slice_group_map_type == 0 {
			pps.run_length_minus1 = make([]int, 0, pps.num_slice_groups_minus1+1)
			for iGroup := 0; iGroup <= pps.num_slice_groups_minus1; iGroup++ {
				pps.run_length_minus1[iGroup] = bits.UE(bs)
				fmt.Printf("%srun_length_minus1[%d]: %d\n", prefix, iGroup, pps.run_length_minus1)
			}
		} else if pps.slice_group_map_type == 2 {
			pps.top_left = make([]int, 0, pps.num_slice_groups_minus1+1)
			pps.bottom_right = make([]int, 0, pps.num_slice_groups_minus1+1)
			for iGroup := 0; iGroup <= pps.num_slice_groups_minus1; iGroup++ {
				pps.top_left[iGroup] = bits.UE(bs)
				fmt.Printf("%stop_left[%d]: %d\n", prefix, iGroup, pps.top_left)
				pps.bottom_right[iGroup] = bits.UE(bs)
				fmt.Printf("%sbottom_right[%d]: %d\n", prefix, iGroup, pps.bottom_right)
			}
		} else if pps.slice_group_map_type == 3 || pps.slice_group_map_type == 4 || pps.slice_group_map_type == 5 {
			pps.slice_group_change_direction_flag = bits.U(bs, 1)
			fmt.Printf("%sslice_group_change_direction_flag: %d\n", prefix, pps.slice_group_change_direction_flag)
			pps.slice_group_change_rate_minus1 = bits.UE(bs)
			fmt.Printf("%sslice_group_change_rate_minus1: %d\n", prefix, pps.slice_group_change_rate_minus1)
		} else if pps.slice_group_map_type == 6 {
			// slice_group_id[i] identifies a slice group of the i-th slice group map unit in raster scan order.
			// The length of the slice_group_id[i] syntax element is Ceil( Log2( num_slice_groups_minus1 + 1 ) ) gobits.
			// The value of slice_group_id[i] shall be in the range of 0 to num_slice_groups_minus1, inclusive.
			slice_group_id_bits := bits.CeilLog2(pps.num_slice_groups_minus1 + 1)
			pps.pic_size_in_map_units_minus1 = bits.UE(bs)
			fmt.Printf("%spic_size_in_map_units_minus1: %d\n", prefix, pps.pic_size_in_map_units_minus1)
			pps.slice_group_id = make([]int, 0, pps.pic_size_in_map_units_minus1+1)
			for i := 0; i <= pps.pic_size_in_map_units_minus1; i++ {
				pps.slice_group_id[i] = bits.U(bs, byte(slice_group_id_bits))
				fmt.Printf("%sslice_group_id[%d]: %d\n", prefix, i, pps.slice_group_id)
			}
		}
	}
	pps.num_ref_idx_l0_default_active_minus1 = bits.UE(bs)
	fmt.Printf("%snum_ref_idx_l0_default_active_minus1: %d\n", prefix, pps.num_ref_idx_l0_default_active_minus1)
	pps.num_ref_idx_l1_default_active_minus1 = bits.UE(bs)
	fmt.Printf("%snum_ref_idx_l1_default_active_minus1: %d\n", prefix, pps.num_ref_idx_l1_default_active_minus1)
	pps.weighted_pred_flag = bits.U(bs, 1)
	fmt.Printf("%sweighted_pred_flag: %d\n", prefix, pps.weighted_pred_flag)
	pps.weighted_bipred_idc = bits.U(bs, 2)
	fmt.Printf("%sweighted_bipred_idc: %d\n", prefix, pps.weighted_bipred_idc)
	pps.pic_init_qp_minus26 = bits.SE(bs)
	fmt.Printf("%spic_init_qp_minus26: %d\n", prefix, pps.pic_init_qp_minus26)
	pps.pic_init_qs_minus26 = bits.SE(bs)
	fmt.Printf("%spic_init_qs_minus26: %d\n", prefix, pps.pic_init_qs_minus26)
	pps.chroma_qp_index_offset = bits.SE(bs)
	fmt.Printf("%schroma_qp_index_offset: %d\n", prefix, pps.chroma_qp_index_offset)
	pps.deblocking_filter_control_present_flag = bits.U(bs, 1)
	fmt.Printf("%sdeblocking_filter_control_present_flag: %d\n", prefix, pps.deblocking_filter_control_present_flag)
	pps.constrained_intra_pred_flag = bits.U(bs, 1)
	fmt.Printf("%sconstrained_intra_pred_flag: %d\n", prefix, pps.constrained_intra_pred_flag)
	pps.redundant_pic_cnt_present_flag = bits.U(bs, 1)
	fmt.Printf("%sredundant_pic_cnt_present_flag: %d\n", prefix, pps.redundant_pic_cnt_present_flag)

	sps, ok := h264.spsTable[pps.seq_parameter_set_id]
	if bits.MoreRBSP(bs) && ok {
		pps.transform_8x8_mode_flag = bits.U(bs, 1)
		fmt.Printf("%stransform_8x8_mode_flag: %d\n", prefix, pps.transform_8x8_mode_flag)
		pps.pic_scaling_matrix_present_flag = bits.U(bs, 1)
		fmt.Printf("%spic_scaling_matrix_present_flag: %d\n", prefix, pps.pic_scaling_matrix_present_flag)
		if pps.pic_scaling_matrix_present_flag == 1 {
			count := 6
			if sps.chroma_format_idc != 3 {
				count += 2 * pps.transform_8x8_mode_flag
			} else {
				count += 6 * pps.transform_8x8_mode_flag
			}
			pps.pic_scaling_list_present_flag = make([]int, 0, count)
			for i := 0; i < count; i++ {
				pps.pic_scaling_list_present_flag[i] = bits.U(bs, 1)
				fmt.Printf("%spic_scaling_list_present_flag[%d]: %d\n", prefix, i, pps.pic_scaling_list_present_flag)
				if pps.pic_scaling_list_present_flag[i] == 1 {
					if i < 6 {
						parseScalingList(bs, 16, prefix)
					} else {
						parseScalingList(bs, 64, prefix)
					}
				}
			}
		}
		pps.second_chroma_qp_index_offset = bits.SE(bs)
		fmt.Printf("%ssecond_chroma_qp_index_offset: %d\n", prefix, pps.second_chroma_qp_index_offset)
	}

	h264.ppsTable[pps.pic_parameter_set_id] = pps
}
