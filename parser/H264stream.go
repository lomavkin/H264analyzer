package parser

import (
	"fmt"

	"github.com/ibbbpbbbp/gobits"
	"github.com/ibbbpbbbp/H264analyzer/bits"
)

const (
	naluShortStartSequenceSize = 3
	naluTypeSize               = 1

	naluTypeOfSlice         = 1
	naluTypeOfIDR           = 5
	naluTypeOfSEI           = 6
	naluTypeOfSPS           = 7
	naluTypeOfPPS           = 8
	naluTypeOfAUD           = 9
	naluTypeOfEndOfSequence = 10
	naluTypeOfEndOfStream   = 11
	naluTypeOfFiller        = 12
	naluTypeOfAUX           = 19

	sliceTypeOfP  = 0
	sliceTypeOfB  = 1
	sliceTypeOfI  = 2
	sliceTypeOfSP = 3
	sliceTypeOfSI = 4
)

type nalindex struct {
	offset int64
	size   int64
}

type nal struct {
	forbidden_zero_bit int
	nal_ref_idc        int
	nal_unit_type      int
}

type slice_header struct {
	first_mb_in_slice                int
	slice_type                       int
	pic_parameter_set_id             int
	colour_plane_id                  int
	frame_num                        int
	field_pic_flag                   int
	bottom_field_flag                int
	idr_pic_id                       int
	pic_order_cnt_lsb                int
	delta_pic_order_cnt_bottom       int
	delta_pic_order_cnt              [2]int
	redundant_pic_cnt                int
	direct_spatial_mv_pred_flag      int
	num_ref_idx_active_override_flag int
	num_ref_idx_l0_active_minus1     int
	num_ref_idx_l1_active_minus1     int
	cabac_init_idc                   int
	slice_qp_delta                   int
	sp_for_switch_flag               int
	slice_qs_delta                   int
	disable_deblocking_filter_idc    int
	slice_alpha_c0_offset_div2       int
	slice_beta_offset_div2           int
	slice_group_change_cycle         int

	pwt struct {
		luma_log2_weight_denom   int
		chroma_log2_weight_denom int
		luma_weight_l0_flag      []int
		luma_weight_l0           []int
		luma_offset_l0           []int
		chroma_weight_l0_flag    []int
		chroma_weight_l0         [][2]int
		chroma_offset_l0         [][2]int
		luma_weight_l1_flag      []int
		luma_weight_l1           []int
		luma_offset_l1           []int
		chroma_weight_l1_flag    []int
		chroma_weight_l1         [][2]int
		chroma_offset_l1         [][2]int
	}

	rplmm struct {
		ref_pic_list_modification_flag_l0 int

		moditication_l0 struct {
			modification_of_pic_nums_idc []int
			abs_diff_pic_num_minus1      []int
			long_term_pic_num            []int
			abs_diff_view_idx_minus1     []int
		}

		ref_pic_list_modification_flag_l1 int

		modification_l1 struct {
			modification_of_pic_nums_idc []int
			abs_diff_pic_num_minus1      []int
			long_term_pic_num            []int
			abs_diff_view_idx_minus1     []int
		}
	}

	rplm struct {
		ref_pic_list_modification_flag_l0 int

		moditication_l0 struct {
			modification_of_pic_nums_idc []int
			abs_diff_pic_num_minus1      []int
			long_term_pic_num            []int
		}

		ref_pic_list_modification_flag_l1 int

		modification_l1 struct {
			modification_of_pic_nums_idc []int
			abs_diff_pic_num_minus1      []int
			long_term_pic_num            []int
		}
	}

	drpm struct {
		no_output_of_prior_pics_flag        int
		long_term_reference_flag            int
		adaptive_ref_pic_marking_mode_flag  int
		memory_management_control_operation []int
		difference_of_pic_nums_minus1       []int
		long_term_pic_num                   []int
		long_term_frame_idx                 []int
		max_long_term_frame_idx_plus1       []int
	}
}

type hrd struct {
	cpb_cnt_minus1                          int
	bit_rate_scale                          int
	cpb_size_scale                          int
	bit_rate_value_minus1                   []int
	cpb_size_value_minus1                   []int
	cbr_flag                                []int
	initial_cpb_removal_delay_length_minus1 int
	cpb_removal_delay_length_minus1         int
	dpb_output_delay_length_minus1          int
	time_offset_length                      int
}

type sps struct {
	profile_idc                           int
	constraint_set0_flag                  int
	constraint_set1_flag                  int
	constraint_set2_flag                  int
	constraint_set3_flag                  int
	constraint_set4_flag                  int
	constraint_set5_flag                  int
	reserved_zero_2bits                   int
	level_idc                             int
	seq_parameter_set_id                  int
	chroma_format_idc                     int
	separate_colour_plane_flag            int
	bit_depth_luma_minus8                 int
	bit_depth_chroma_minus8               int
	qpprime_y_zero_transform_bypass_flag  int
	seq_scaling_matrix_present_flag       int
	seq_scaling_list_present_flag         []int
	log2_max_frame_num_minus4             int
	pic_order_cnt_type                    int
	log2_max_pic_order_cnt_lsb_minus4     int
	delta_pic_order_always_zero_flag      int
	offset_for_non_ref_pic                int
	offset_for_top_to_bottom_field        int
	num_ref_frames_in_pic_order_cnt_cycle int
	offset_for_ref_frame                  []int
	max_num_ref_frames                    int
	gaps_in_frame_num_value_allowed_flag  int
	pic_width_in_mbs_minus1               int
	pic_height_in_map_units_minus1        int
	frame_mbs_only_flag                   int
	mb_adaptive_frame_field_flag          int
	direct_8x8_inference_flag             int
	frame_cropping_flag                   int
	frame_crop_left_offset                int
	frame_crop_right_offset               int
	frame_crop_top_offset                 int
	frame_crop_bottom_offset              int
	vui_parameters_present_flag           int

	vui struct {
		aspect_ratio_info_present_flag          int
		aspect_ratio_idc                        int
		sar_width                               int
		sar_height                              int
		overscan_info_present_flag              int
		overscan_appropriate_flag               int
		video_signal_type_present_flag          int
		video_format                            int
		video_full_range_flag                   int
		colour_description_present_flag         int
		colour_primaries                        int
		transfer_characteristics                int
		matrix_coefficients                     int
		chroma_loc_info_present_flag            int
		chroma_sample_loc_type_top_field        int
		chroma_sample_loc_type_bottom_field     int
		timing_info_present_flag                int
		num_units_in_tick                       int
		time_scale                              int
		fixed_frame_rate_flag                   int
		nal_hrd_parameters_present_flag         int
		vcl_hrd_parameters_present_flag         int
		low_delay_hrd_flag                      int
		pic_struct_present_flag                 int
		bitstream_restriction_flag              int
		motion_vectors_over_pic_boundaries_flag int
		max_bytes_per_pic_denom                 int
		max_bits_per_mb_denom                   int
		log2_max_mv_length_horizontal           int
		log2_max_mv_length_vertical             int
		max_num_reorder_frames                  int
		max_dec_frame_buffering                 int
	}

	nal_hrd hrd
	vcl_hrd hrd
}

type pps struct {
	pic_parameter_set_id                         int
	seq_parameter_set_id                         int
	entropy_coding_mode_flag                     int
	bottom_field_pic_order_in_frame_present_flag int
	num_slice_groups_minus1                      int
	slice_group_map_type                         int
	run_length_minus1                            []int
	top_left                                     []int
	bottom_right                                 []int
	slice_group_change_direction_flag            int
	slice_group_change_rate_minus1               int
	pic_size_in_map_units_minus1                 int
	slice_group_id                               []int
	num_ref_idx_l0_default_active_minus1         int
	num_ref_idx_l1_default_active_minus1         int
	weighted_pred_flag                           int
	weighted_bipred_idc                          int
	pic_init_qp_minus26                          int
	pic_init_qs_minus26                          int
	chroma_qp_index_offset                       int
	deblocking_filter_control_present_flag       int
	constrained_intra_pred_flag                  int
	redundant_pic_cnt_present_flag               int
	_more_rbsp_data_present                      int
	transform_8x8_mode_flag                      int
	pic_scaling_matrix_present_flag              int
	pic_scaling_list_present_flag                []int
	second_chroma_qp_index_offset                int
}

type H264Stream struct {
	nal *nal

	spsTable map[int]*sps
	ppsTable map[int]*pps
}

func ParseH264Stream(ba gobits.ByteAccessor) *H264Stream {
	h264 := &H264Stream{
		spsTable: map[int]*sps{},
		ppsTable: map[int]*pps{},
	}
	for pos := int64(0); ; {
		var index *nalindex
		index, pos = findNAL(ba, pos)
		fmt.Printf("== NAL at offset %d, size %d ==\n", index.offset, index.size)
		parseNAL(h264, bits.RBSP(ba.Slice(index.offset, index.size)), "  ")
		if pos == 0 {
			break
		}
	}
	return nil
}
