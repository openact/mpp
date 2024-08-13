package layout

type Input struct {
	SPCODE                             int     `csv:"SPCODE,omitempty"`
	POL_NUMBER                         int     `csv:"POL_NUMBER,omitempty"`
	AGE_AT_ENTRY                       int     `csv:"AGE_AT_ENTRY,omitempty"`
	SEX                                int     `csv:"SEX,omitempty"`
	POL_TERM_Y                         int     `csv:"POL_TERM_Y,omitempty"`
	ANNUAL_PREM                        float64 `csv:"ANNUAL_PREM,omitempty"`
	PREM_PAYBL_Y                       int     `csv:"PREM_PAYBL_Y,omitempty"`
	PREM_FREQ                          int     `csv:"PREM_FREQ,omitempty"`
	SUM_ASSURED                        float64 `csv:"SUM_ASSURED,omitempty"`
	SMOKER_STAT                        int     `csv:"SMOKER_STAT,omitempty"`
	DURATIONIF_M                       int     `csv:"DURATIONIF_M,omitempty"`
	INIT_POLS_IF                       int     `csv:"INIT_POLS_IF,omitempty"`
	INIT_UNFDU_A                       float64 `csv:"INIT_UNFDU_A,omitempty"`
	DIST_IND                           int     `csv:"DIST_IND,omitempty"`
	OCC_CLASS                          int     `csv:"OCC_CLASS,omitempty"`
	ACDTH_BEN                          float64 `csv:"ACDTH_BEN,omitempty"`
	ACDMB_BEN                          float64 `csv:"ACDMB_BEN,omitempty"`
	LIFE2_AGE_AE                       int     `csv:"LIFE2_AGE_AE,omitempty"`
	LIFE2_TERM_AGE_AE                  int     `csv:"LIFE2_TERM_AGE_AE,omitempty"`
	CESS_AGE_WP                        int     `csv:"CESS_AGE_WP,omitempty"`
	PWAIV_BEN                          float64 `csv:"PWAIV_BEN,omitempty"`
	PWAI2_BEN                          float64 `csv:"PWAI2_BEN,omitempty"`
	CESS_AGE_SP                        int     `csv:"CESS_AGE_SP,omitempty"`
	PRU_MED_BEN                        float64 `csv:"PRU_MED_BEN,omitempty"`
	CC_BEN                             float64 `csv:"CC_BEN,omitempty"`
	CC_EXP                             int     `csv:"CC_EXP,omitempty"`
	ADD_CC_BEN                         float64 `csv:"ADD_CC_BEN,omitempty"`
	PLAN_CC                            int     `csv:"PLAN_CC,omitempty"`
	PLAN_CCP                           int     `csv:"PLAN_CCP,omitempty"`
	LIFE3_AGE_AE                       int     `csv:"LIFE3_AGE_AE,omitempty"`
	LIFE3_TERM_AGE_AE                  int     `csv:"LIFE3_TERM_AGE_AE,omitempty"`
	TERM_BEN                           float64 `csv:"TERM_BEN,omitempty"`
	TERM_EXP                           int     `csv:"TERM_EXP,omitempty"`
	MEDICAL_BEN                        float64 `csv:"MEDICAL_BEN,omitempty"`
	PLAN_MED                           int     `csv:"PLAN_MED,omitempty"`
	HSPLUS_IND                         int     `csv:"HSPLUS_IND,omitempty"`
	MCC_BEN                            float64 `csv:"MCC_BEN,omitempty"`
	MCC_EXP                            int     `csv:"MCC_EXP,omitempty"`
	CINCOME_BEN                        float64 `csv:"CINCOME_BEN,omitempty"`
	CINCOME_EXP                        int     `csv:"CINCOME_EXP,omitempty"`
	ESCC_BEN                           float64 `csv:"ESCC_BEN,omitempty"`
	ESCC_EXP                           int     `csv:"ESCC_EXP,omitempty"`
	CURR                               int     `csv:"CURR,omitempty"`
	REDU_ADM                           int     `csv:"REDU_ADM,omitempty"`
	ANN_PREM_BAS                       float64 `csv:"ANN_PREM_BAS,omitempty"`
	ANN_PREM_REC                       float64 `csv:"ANN_PREM_REC,omitempty"`
	PRM_B_ALL_F1                       float64 `csv:"PRM_B_ALL_F1,omitempty"`
	PRM_R_ALL_F1                       float64 `csv:"PRM_R_ALL_F1,omitempty"`
	I_UNFDU_B_F1                       float64 `csv:"I_UNFDU_B_F1,omitempty"`
	I_UNFDU_R_F1                       float64 `csv:"I_UNFDU_R_F1,omitempty"`
	I_UNFDU_T_F1                       float64 `csv:"I_UNFDU_T_F1,omitempty"`
	INIT_EXCS_DD                       float64 `csv:"INIT_EXCS_DD,omitempty"`
	I_EXCS_DD_B                        float64 `csv:"I_EXCS_DD_B,omitempty"`
	I_EXCS_DD_R                        float64 `csv:"I_EXCS_DD_R,omitempty"`
	NO_LAP_GTEE                        int     `csv:"NO_LAP_GTEE,omitempty"`
	COMM_IND_B                         int     `csv:"COMM_IND_B,omitempty"`
	COMM_IND_R                         int     `csv:"COMM_IND_R,omitempty"`
	EXP_IND_B                          int     `csv:"EXP_IND_B,omitempty"`
	EXP_IND_R                          int     `csv:"EXP_IND_R,omitempty"`
	V_EXP_IND_B                        int     `csv:"V_EXP_IND_B,omitempty"`
	V_EXP_IND_R                        int     `csv:"V_EXP_IND_R,omitempty"`
	FUND_IND_B                         int     `csv:"FUND_IND_B,omitempty"`
	FUND_IND_R                         int     `csv:"FUND_IND_R,omitempty"`
	RI_ACDMB_BEN_PC                    float64 `csv:"RI_ACDMB_BEN_PC,omitempty"`
	RI_ACDTH_BEN_PC                    float64 `csv:"RI_ACDTH_BEN_PC,omitempty"`
	RI_ADD_CC_BEN_PC                   float64 `csv:"RI_ADD_CC_BEN_PC,omitempty"`
	RI_BEN_CI                          float64 `csv:"RI_BEN_CI,omitempty"`
	RI_BEN_WAIV                        float64 `csv:"RI_BEN_WAIV,omitempty"`
	RI_BEN_WAIV2                       float64 `csv:"RI_BEN_WAIV2,omitempty"`
	RI_CC_BEN_PC                       float64 `csv:"RI_CC_BEN_PC,omitempty"`
	RI_CINCOME_BEN_PC                  float64 `csv:"RI_CINCOME_BEN_PC,omitempty"`
	RI_ESCC_BEN_PC                     float64 `csv:"RI_ESCC_BEN_PC,omitempty"`
	RI_MCC_BEN_PC                      float64 `csv:"RI_MCC_BEN_PC,omitempty"`
	RI_PRU_HS_BEN_PC                   float64 `csv:"RI_PRU_HS_BEN_PC,omitempty"`
	RI_PRU_MED_BEN_PC                  float64 `csv:"RI_PRU_MED_BEN_PC,omitempty"`
	RI_PWAI2_BEN_PC                    float64 `csv:"RI_PWAI2_BEN_PC,omitempty"`
	RI_PWAIV_BEN_PC                    float64 `csv:"RI_PWAIV_BEN_PC,omitempty"`
	RI_SUMASSD_PC                      float64 `csv:"RI_SUMASSD_PC,omitempty"`
	RI_TERM_BEN_PC                     float64 `csv:"RI_TERM_BEN_PC,omitempty"`
	REINS_IND                          float64 `csv:"REINS_IND,omitempty"`
	MAX_AGE_AD                         int     `csv:"MAX_AGE_AD,omitempty"`
	MAX_AGE_ADD                        int     `csv:"MAX_AGE_ADD,omitempty"`
	MAX_AGE_WAIV                       int     `csv:"MAX_AGE_WAIV,omitempty"`
	MAX_AGE_WAIV2                      int     `csv:"MAX_AGE_WAIV2,omitempty"`
	MAX_AGE_PMED                       int     `csv:"MAX_AGE_PMED,omitempty"`
	MAX_AGE_ADD_CC                     int     `csv:"MAX_AGE_ADD_CC,omitempty"`
	MAX_AGE_HS                         int     `csv:"MAX_AGE_HS,omitempty"`
	MAX_AGE_TERM                       int     `csv:"MAX_AGE_TERM,omitempty"`
	MAX_AGE_MCC                        int     `csv:"MAX_AGE_MCC,omitempty"`
	MAX_AGE_CINCOME                    int     `csv:"MAX_AGE_CINCOME,omitempty"`
	MAX_AGE_ESCC                       int     `csv:"MAX_AGE_ESCC,omitempty"`
	COI_TERM_DISCOUNT_TYPE             int     `csv:"COI_TERM_DISCOUNT_TYPE,omitempty"`
	RI_TABLE_IND                       int     `csv:"RI_TABLE_IND,omitempty"`
	RI_AD_IND                          int     `csv:"RI_AD_IND,omitempty"`
	RI_ADDB_IND                        int     `csv:"RI_ADDB_IND,omitempty"`
	RI_WAI_IND                         int     `csv:"RI_WAI_IND,omitempty"`
	RI_WAI2_IND                        int     `csv:"RI_WAI2_IND,omitempty"`
	RI_MED_IND                         int     `csv:"RI_MED_IND,omitempty"`
	RI_ADD_CC_IND                      int     `csv:"RI_ADD_CC_IND,omitempty"`
	RI_HS_IND                          int     `csv:"RI_HS_IND,omitempty"`
	RI_TERM_IND                        int     `csv:"RI_TERM_IND,omitempty"`
	RI_MCC_IND                         int     `csv:"RI_MCC_IND,omitempty"`
	RI_CI_IND                          int     `csv:"RI_CI_IND,omitempty"`
	RI_ESCC_IND                        int     `csv:"RI_ESCC_IND,omitempty"`
	RI_CC_IND                          int     `csv:"RI_CC_IND,omitempty"`
	PW_IND_USED                        int     `csv:"PW_IND_USED,omitempty"`
	PH_IND_PP                          int     `csv:"PH_IND_PP,omitempty"`
	PLAN_PRNTL                         int     `csv:"PLAN_PRNTL,omitempty"`
	AGE_MT_PRNTL                       int     `csv:"AGE_MT_PRNTL,omitempty"`
	MAX_AGE_PNTL                       int     `csv:"MAX_AGE_PNTL,omitempty"`
	RI_PRU_PRNTL_BEN_PC                float64 `csv:"RI_PRU_PRNTL_BEN_PC,omitempty"`
	JUVCI_EXP                          int     `csv:"JUVCI_EXP,omitempty"`
	JUVCI_BEN                          float64 `csv:"JUVCI_BEN,omitempty"`
	RI_PRU_JUVCI_BEN_PC                float64 `csv:"RI_PRU_JUVCI_BEN_PC,omitempty"`
	UPR_FAC                            float64 `csv:"UPR_FAC,omitempty"`
	ACDMB_IND                          int     `csv:"ACDMB_IND,omitempty"`
	ACDTH_IND                          int     `csv:"ACDTH_IND,omitempty"`
	ESCC_PLUS_IND                      int     `csv:"ESCC_PLUS_IND,omitempty"`
	PROD_NAME                          string  `csv:"PROD_NAME,omitempty"`
	ALLOC_IND                          int     `csv:"ALLOC_IND,omitempty"`
	WP                                 int     `csv:"WP,omitempty"`
	SP_TERM                            float64 `csv:"SP_TERM,omitempty"`
	REGION                             int     `csv:"REGION,omitempty"`
	SA_MULT_GRP                        int     `csv:"SA_MULT_GRP,omitempty"`
	SEX_2                              int     `csv:"SEX_2,omitempty"`
	SEX_3                              int     `csv:"SEX_3,omitempty"`
	SMOKER_STAT_2                      int     `csv:"SMOKER_STAT_2,omitempty"`
	SMOKER_STAT_3                      int     `csv:"SMOKER_STAT_3,omitempty"`
	TERM_BEN_3                         float64 `csv:"TERM_BEN_3,omitempty"`
	PEP_BEN_FIX                        float64 `csv:"PEP_BEN_FIX,omitempty"`
	PEP_BEN_FIX2                       float64 `csv:"PEP_BEN_FIX2,omitempty"`
	MAX_AGE_PEP1                       int     `csv:"MAX_AGE_PEP1,omitempty"`
	MAX_AGE_PEP2                       int     `csv:"MAX_AGE_PEP2,omitempty"`
	MAX_AGE_TERM_3                     int     `csv:"MAX_AGE_TERM_3,omitempty"`
	RI_PEP1_BEN_PC                     float64 `csv:"RI_PEP1_BEN_PC,omitempty"`
	RI_PEP2_BEN_PC                     float64 `csv:"RI_PEP2_BEN_PC,omitempty"`
	WOP_ESCC_IND                       int     `csv:"WOP_ESCC_IND,omitempty"`
	WOP2_ESCC_IND                      int     `csv:"WOP2_ESCC_IND,omitempty"`
	CCB61_BEN                          float64 `csv:"CCB61_BEN,omitempty"`
	CCB61_EXP                          int     `csv:"CCB61_EXP,omitempty"`
	MAX_AGE_CCB61                      int     `csv:"MAX_AGE_CCB61,omitempty"`
	RI_CCB61_BEN_PC                    float64 `csv:"RI_CCB61_BEN_PC,omitempty"`
	HS_MONTH                           int     `csv:"HS_MONTH,omitempty"`
	INIT_BON_F1                        float64 `csv:"INIT_BON_F1,omitempty"`
	INIT_BONP_F1                       float64 `csv:"INIT_BONP_F1,omitempty"`
	BILL_CHNL                          int     `csv:"BILL_CHNL,omitempty"`
	BILL_CHNL2                         int     `csv:"BILL_CHNL2,omitempty"`
	ANN_BON_PREM_BAS                   float64 `csv:"ANN_BON_PREM_BAS,omitempty"`
	CI_BEN                             float64 `csv:"CI_BEN,omitempty"`
	MAX_AGE_CI                         float64 `csv:"MAX_AGE_CI,omitempty"`
	PLAN_CI                            float64 `csv:"PLAN_CI,omitempty"`
	RI_CI_BEN_PC                       float64 `csv:"RI_CI_BEN_PC,omitempty"`
	TPD_BEN                            float64 `csv:"TPD_BEN,omitempty"`
	TPD_EXP                            float64 `csv:"TPD_EXP,omitempty"`
	MAX_AGE_TPD                        float64 `csv:"MAX_AGE_TPD,omitempty"`
	RI_TPD_BEN_PC                      float64 `csv:"RI_TPD_BEN_PC,omitempty"`
	PPH_PLUS_BEN                       float64 `csv:"PPH_PLUS_BEN,omitempty"`
	PPH2_PLAN_MED                      int     `csv:"PPH2_PLAN_MED,omitempty"`
	PPH_PLUS_IND                       int     `csv:"PPH_PLUS_IND,omitempty"`
	PPH2_HFA_IND                       int     `csv:"PPH2_HFA_IND,omitempty"`
	MAX_AGE_PPH_PLUS                   int     `csv:"MAX_AGE_PPH_PLUS,omitempty"`
	RI_PPH_PLUS_BEN_PC                 float64 `csv:"RI_PPH_PLUS_BEN_PC,omitempty"`
	SEL_FAC_IND                        int     `csv:"SEL_FAC_IND,omitempty"`
	CICA_CCB34_IND                     int     `csv:"CICA_CCB34_IND,omitempty"`
	CICA_CCB61_IND                     int     `csv:"CICA_CCB61_IND,omitempty"`
	CICA_ESCC_IND                      int     `csv:"CICA_ESCC_IND,omitempty"`
	DUR_PEP_AT_ENTRY                   int     `csv:"DUR_PEP_AT_ENTRY,omitempty"`
	MUNRE_2020                         int     `csv:"MUNRE_2020,omitempty"`
	WVR_CLAIM_STAT                     int     `csv:"WVR_CLAIM_STAT,omitempty"`
	CINC_CLAIM_STAT                    int     `csv:"CINC_CLAIM_STAT,omitempty"`
	PEP_CLAIM_STAT                     int     `csv:"PEP_CLAIM_STAT,omitempty"`
	WVR_IND                            int     `csv:"WVR_IND,omitempty"`
	CINC_IND                           int     `csv:"CINC_IND,omitempty"`
	PEP_IND                            int     `csv:"PEP_IND,omitempty"`
	PROD_CD                            string  `csv:"PROD_CD,omitempty"`
	IFRS_CCF_EXTERNAL_1                int     `csv:"IFRS_CCF_EXTERNAL_1,omitempty"`
	IFRS_CCF_EXTERNAL_2                int     `csv:"IFRS_CCF_EXTERNAL_2,omitempty"`
	IFRS_CCF_EXTERNAL_3                int     `csv:"IFRS_CCF_EXTERNAL_3,omitempty"`
	IFRS_CCF_EXTERNAL_4                int     `csv:"IFRS_CCF_EXTERNAL_4,omitempty"`
	IFRS_CCF_EXTERNAL_5                int     `csv:"IFRS_CCF_EXTERNAL_5,omitempty"`
	IFRS_CCF_EXTERNAL_6                int     `csv:"IFRS_CCF_EXTERNAL_6,omitempty"`
	IFRS_CCF_EXTERNAL_7                int     `csv:"IFRS_CCF_EXTERNAL_7,omitempty"`
	IFRS_CCF_EXTERNAL_8                int     `csv:"IFRS_CCF_EXTERNAL_8,omitempty"`
	IFRS_CCF_EXTERNAL_9                int     `csv:"IFRS_CCF_EXTERNAL_9,omitempty"`
	IFRS_CCF_EXTERNAL_10               int     `csv:"IFRS_CCF_EXTERNAL_10,omitempty"`
	IFRS_CCF_EXTERNAL_11               int     `csv:"IFRS_CCF_EXTERNAL_11,omitempty"`
	IFRS_CCF_EXTERNAL_12               int     `csv:"IFRS_CCF_EXTERNAL_12,omitempty"`
	IFRS_CCF_EXTERNAL_13               int     `csv:"IFRS_CCF_EXTERNAL_13,omitempty"`
	IFRS_CCF_EXTERNAL_14               int     `csv:"IFRS_CCF_EXTERNAL_14,omitempty"`
	IFRS_CCF_EXTERNAL_15               int     `csv:"IFRS_CCF_EXTERNAL_15,omitempty"`
	IFRS_CCF_EXTERNAL_16               int     `csv:"IFRS_CCF_EXTERNAL_16,omitempty"`
	IFRS_CCF_EXTERNAL_17               int     `csv:"IFRS_CCF_EXTERNAL_17,omitempty"`
	IFRS_CCF_EXTERNAL_18               int     `csv:"IFRS_CCF_EXTERNAL_18,omitempty"`
	IFRS_CCF_EXTERNAL_19               int     `csv:"IFRS_CCF_EXTERNAL_19,omitempty"`
	IFRS_CCF_EXTERNAL_20               int     `csv:"IFRS_CCF_EXTERNAL_20,omitempty"`
	IFRS_CCF_EXTERNAL_21               int     `csv:"IFRS_CCF_EXTERNAL_21,omitempty"`
	IFRS_CCF_EXTERNAL_22               int     `csv:"IFRS_CCF_EXTERNAL_22,omitempty"`
	IFRS_CCF_EXTERNAL_23               int     `csv:"IFRS_CCF_EXTERNAL_23,omitempty"`
	IFRS_CCF_EXTERNAL_24               int     `csv:"IFRS_CCF_EXTERNAL_24,omitempty"`
	IFRS_CCF_EXTERNAL_25               int     `csv:"IFRS_CCF_EXTERNAL_25,omitempty"`
	IFRS_CCF_EXTERNAL_26               int     `csv:"IFRS_CCF_EXTERNAL_26,omitempty"`
	IFRS_CCF_EXTERNAL_27               int     `csv:"IFRS_CCF_EXTERNAL_27,omitempty"`
	IFRS_CCF_EXTERNAL_28               int     `csv:"IFRS_CCF_EXTERNAL_28,omitempty"`
	IFRS_CURRENCY                      string  `csv:"IFRS_CURRENCY,omitempty"`
	IFRS_ICG_ID                        string  `csv:"IFRS_ICG_ID,omitempty"`
	IFRS_ICG_ID_PROPHET                int     `csv:"IFRS_ICG_ID_PROPHET,omitempty"`
	IF_NB                              string  `csv:"IF_NB,omitempty"`
	IFRS_MEASURE_MODEL                 int     `csv:"IFRS_MEASURE_MODEL,omitempty"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY1 string  `csv:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY1,omitempty"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY2 string  `csv:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY2,omitempty"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY3 string  `csv:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY3,omitempty"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY4 string  `csv:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY4,omitempty"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY5 string  `csv:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY5,omitempty"`
	IFRS_RI_EXT_CCF_IND                int     `csv:"IFRS_RI_EXT_CCF_IND,omitempty"`
	IFRS_RI_ICG_ID_TREATY1             string  `csv:"IFRS_RI_ICG_ID_TREATY1,omitempty"`
	IFRS_RI_ICG_ID_TREATY2             string  `csv:"IFRS_RI_ICG_ID_TREATY2,omitempty"`
	IFRS_RI_ICG_ID_TREATY3             string  `csv:"IFRS_RI_ICG_ID_TREATY3,omitempty"`
	IFRS_RI_ICG_ID_TREATY4             string  `csv:"IFRS_RI_ICG_ID_TREATY4,omitempty"`
	IFRS_RI_ICG_ID_TREATY5             string  `csv:"IFRS_RI_ICG_ID_TREATY5,omitempty"`
	IFRS_RI_ICG_ID_TREATY6             string  `csv:"IFRS_RI_ICG_ID_TREATY6,omitempty"`
	IFRS_RI_ICG_ID_TREATY7             string  `csv:"IFRS_RI_ICG_ID_TREATY7,omitempty"`
	IFRS_RI_ICG_ID_TREATY8             string  `csv:"IFRS_RI_ICG_ID_TREATY8,omitempty"`
	IFRS_RI_ICG_ID_TREATY9             string  `csv:"IFRS_RI_ICG_ID_TREATY9,omitempty"`
	IFRS_RI_ICG_ID_TREATY10            string  `csv:"IFRS_RI_ICG_ID_TREATY10,omitempty"`
	IFRS_RI_ICG_ID_TREATY11            string  `csv:"IFRS_RI_ICG_ID_TREATY11,omitempty"`
	IFRS_RI_ICG_ID_TREATY12            string  `csv:"IFRS_RI_ICG_ID_TREATY12,omitempty"`
	IFRS_RI_ICG_ID_TREATY13            string  `csv:"IFRS_RI_ICG_ID_TREATY13,omitempty"`
	IFRS_RI_ICG_ID_TREATY14            string  `csv:"IFRS_RI_ICG_ID_TREATY14,omitempty"`
	IFRS_RI_ICG_ID_TREATY15            string  `csv:"IFRS_RI_ICG_ID_TREATY15,omitempty"`
	IFRS_RI_ICG_ID_TREATY16            string  `csv:"IFRS_RI_ICG_ID_TREATY16,omitempty"`
	IFRS_RI_ICG_ID_TREATY17            string  `csv:"IFRS_RI_ICG_ID_TREATY17,omitempty"`
	IFRS_RI_ICG_ID_TREATY18            string  `csv:"IFRS_RI_ICG_ID_TREATY18,omitempty"`
	IFRS_RI_ICG_ID_TREATY19            string  `csv:"IFRS_RI_ICG_ID_TREATY19,omitempty"`
	IFRS_RI_ICG_ID_TREATY20            string  `csv:"IFRS_RI_ICG_ID_TREATY20,omitempty"`
	IFRS_RI_ICG_ID_TREATY21            string  `csv:"IFRS_RI_ICG_ID_TREATY21,omitempty"`
	IFRS_RI_ICG_ID_TREATY22            string  `csv:"IFRS_RI_ICG_ID_TREATY22,omitempty"`
	IFRS_RI_ICG_ID_TREATY23            string  `csv:"IFRS_RI_ICG_ID_TREATY23,omitempty"`
	IFRS_RI_ICG_ID_TREATY24            string  `csv:"IFRS_RI_ICG_ID_TREATY24,omitempty"`
	IFRS_RI_ICG_ID_TREATY25            string  `csv:"IFRS_RI_ICG_ID_TREATY25,omitempty"`
	TREATY_ID_TREATY1                  string  `csv:"TREATY_ID_TREATY1,omitempty"`
	TREATY_ID_TREATY2                  string  `csv:"TREATY_ID_TREATY2,omitempty"`
	TREATY_ID_TREATY3                  string  `csv:"TREATY_ID_TREATY3,omitempty"`
	TREATY_ID_TREATY4                  string  `csv:"TREATY_ID_TREATY4,omitempty"`
	TREATY_ID_TREATY5                  string  `csv:"TREATY_ID_TREATY5,omitempty"`
	TREATY_ID_TREATY6                  string  `csv:"TREATY_ID_TREATY6,omitempty"`
	TREATY_ID_TREATY7                  string  `csv:"TREATY_ID_TREATY7,omitempty"`
	TREATY_ID_TREATY8                  string  `csv:"TREATY_ID_TREATY8,omitempty"`
	TREATY_ID_TREATY9                  string  `csv:"TREATY_ID_TREATY9,omitempty"`
	TREATY_ID_TREATY10                 string  `csv:"TREATY_ID_TREATY10,omitempty"`
	TREATY_ID_TREATY11                 string  `csv:"TREATY_ID_TREATY11,omitempty"`
	TREATY_ID_TREATY12                 string  `csv:"TREATY_ID_TREATY12,omitempty"`
	TREATY_ID_TREATY13                 string  `csv:"TREATY_ID_TREATY13,omitempty"`
	TREATY_ID_TREATY14                 string  `csv:"TREATY_ID_TREATY14,omitempty"`
	TREATY_ID_TREATY15                 string  `csv:"TREATY_ID_TREATY15,omitempty"`
	TREATY_ID_TREATY16                 string  `csv:"TREATY_ID_TREATY16,omitempty"`
	TREATY_ID_TREATY17                 string  `csv:"TREATY_ID_TREATY17,omitempty"`
	TREATY_ID_TREATY18                 string  `csv:"TREATY_ID_TREATY18,omitempty"`
	TREATY_ID_TREATY19                 string  `csv:"TREATY_ID_TREATY19,omitempty"`
	TREATY_ID_TREATY20                 string  `csv:"TREATY_ID_TREATY20,omitempty"`
	TREATY_ID_TREATY21                 string  `csv:"TREATY_ID_TREATY21,omitempty"`
	TREATY_ID_TREATY22                 string  `csv:"TREATY_ID_TREATY22,omitempty"`
	TREATY_ID_TREATY23                 string  `csv:"TREATY_ID_TREATY23,omitempty"`
	TREATY_ID_TREATY24                 string  `csv:"TREATY_ID_TREATY24,omitempty"`
	TREATY_ID_TREATY25                 string  `csv:"TREATY_ID_TREATY25,omitempty"`
	EXT_CCF_IND                        int     `csv:"EXT_CCF_IND,omitempty"`
	BASIC_ENTRY_MONTH                  int     `csv:"BASIC_ENTRY_MONTH,omitempty"`
	BASIC_ENTRY_YEAR                   int     `csv:"BASIC_ENTRY_YEAR,omitempty"`
	BENEFIT_CD                         string  `csv:"BENEFIT_CD,omitempty"`
	IFRS_CB_TERM_M                     int     `csv:"IFRS_CB_TERM_M,omitempty"`
	IFRS_CY_GRP                        int     `csv:"IFRS_CY_GRP,omitempty"`
	IFRS_ONEROUS_GRP                   int     `csv:"IFRS_ONEROUS_GRP,omitempty"`
	IFRS_PORT_GRP                      int     `csv:"IFRS_PORT_GRP,omitempty"`
	MTHS_TO_SALE                       int     `csv:"MTHS_TO_SALE,omitempty"`
	ENTITY_ID                          string  `csv:"ENTITY_ID,omitempty"`
	TOP_UP_MONTH                       int     `csv:"TOP_UP_MONTH,omitempty"`
	TOP_UP_YEAR                        int     `csv:"TOP_UP_YEAR,omitempty"`
	TOPUP_DATE                         string  `csv:"TOPUP_DATE,omitempty"`
	TOPUPMONTH                         int     `csv:"TOPUPMONTH,omitempty"`
	TOPUPYEAR                          int     `csv:"TOPUPYEAR,omitempty"`
	RI_CCF_IFRS_COV_UNITS_TREATY1      int     `csv:"RI_CCF_IFRS_COV_UNITS_TREATY1,omitempty"`
	RI_CCF_IFRS_COV_UNITS_TREATY2      int     `csv:"RI_CCF_IFRS_COV_UNITS_TREATY2,omitempty"`
	RI_CCF_IFRS_COV_UNITS_TREATY3      int     `csv:"RI_CCF_IFRS_COV_UNITS_TREATY3,omitempty"`
	RI_CCF_IFRS_COV_UNITS_TREATY4      int     `csv:"RI_CCF_IFRS_COV_UNITS_TREATY4,omitempty"`
	RI_CCF_IFRS_COV_UNITS_TREATY5      int     `csv:"RI_CCF_IFRS_COV_UNITS_TREATY5,omitempty"`
	CC_CODE                            string  `csv:"CC_CODE,omitempty"`
	AD_CODE                            string  `csv:"AD_CODE,omitempty"`
	ADD_CODE                           string  `csv:"ADD_CODE,omitempty"`
	MED_CODE                           string  `csv:"MED_CODE,omitempty"`
	WAI_CODE                           string  `csv:"WAI_CODE,omitempty"`
	SPO_CODE                           string  `csv:"SPO_CODE,omitempty"`
	CCP_CODE                           string  `csv:"CCP_CODE,omitempty"`
	HS_CODE                            string  `csv:"HS_CODE,omitempty"`
	TERM_CODE                          string  `csv:"TERM_CODE,omitempty"`
	CINC_CODE                          string  `csv:"CINC_CODE,omitempty"`
	MCC_CODE                           string  `csv:"MCC_CODE,omitempty"`
	ESCC_CODE                          string  `csv:"ESCC_CODE,omitempty"`
	EDU_PAR1_CODE                      string  `csv:"EDU_PAR1_CODE,omitempty"`
	EDU_PAR2_CODE                      string  `csv:"EDU_PAR2_CODE,omitempty"`
	CC61_CODE                          string  `csv:"CC61_CODE,omitempty"`
	CIR_CODE                           string  `csv:"CIR_CODE,omitempty"`
	TPDD_CODE                          string  `csv:"TPDD_CODE,omitempty"`
	PAR2_TERM_CODE                     string  `csv:"PAR2_TERM_CODE,omitempty"`
	CURR_TYPE                          string  `csv:"CURR_TYPE,omitempty"`
	POL_NO_IFRS17                      string  `csv:"POL_NO_IFRS17,omitempty"`
	JCC_CODE                           string  `csv:"JCC_CODE,omitempty"`
	SPO1_CODE                          string  `csv:"SPO1_CODE,omitempty"`
	SPO2_CODE                          string  `csv:"SPO2_CODE,omitempty"`
	CICA_CCB34_CODE                    string  `csv:"CICA_CCB34_CODE,omitempty"`
	CICA_ESCC_CODE                     string  `csv:"CICA_ESCC_CODE,omitempty"`
	CICA_CCB61_CODE                    string  `csv:"CICA_CCB61_CODE,omitempty"`
	IFRS_APE_NBMARGIN                  float64 `csv:"IFRS_APE_NBMARGIN,omitempty"`
	IFRS_MI_RID_ST_IND                 int     `csv:"IFRS_MI_RID_ST_IND,omitempty"`
	ANNUAL_PREM_TOPUP                  float64 `csv:"ANNUAL_PREM_TOPUP,omitempty"`
	COMM_IND_B_TOPUP                   int     `csv:"COMM_IND_B_TOPUP,omitempty"`
	EXP_IND_B_TOPUP                    int     `csv:"EXP_IND_B_TOPUP,omitempty"`
	IFRS_TOPUP_IND                     int     `csv:"IFRS_TOPUP_IND,omitempty"`
	RCD                                string  `csv:"RCD,omitempty"`
	HO_ISS_DATE                        string  `csv:"HO_ISS_DATE,omitempty"`
	RI_CICA_ESCC_BEN_PC                float64 `csv:"RI_CICA_ESCC_BEN_PC,omitempty"`
	RI_CICA_CCB34_BEN_PC               float64 `csv:"RI_CICA_CCB34_BEN_PC,omitempty"`
	RI_CICA_CCB61_BEN_PC               float64 `csv:"RI_CICA_CCB61_BEN_PC,omitempty"`
	CICA_ESCC_ALT_IND                  int     `csv:"CICA_ESCC_ALT_IND,omitempty"`
	CICA_CCB61_ALT_IND                 int     `csv:"CICA_CCB61_ALT_IND,omitempty"`
	BASIC_LOAD_PERC                    float64 `csv:"BASIC_LOAD_PERC,omitempty"`
	BASIC_LOAD_ADJ                     float64 `csv:"BASIC_LOAD_ADJ,omitempty"`
	CC_LOAD_PERC                       float64 `csv:"CC_LOAD_PERC,omitempty"`
	CC_LOAD_ADJ                        float64 `csv:"CC_LOAD_ADJ,omitempty"`
	AD_LOAD_PERC                       float64 `csv:"AD_LOAD_PERC,omitempty"`
	AD_LOAD_ADJ                        float64 `csv:"AD_LOAD_ADJ,omitempty"`
	ADD_LOAD_PERC                      float64 `csv:"ADD_LOAD_PERC,omitempty"`
	ADD_LOAD_ADJ                       float64 `csv:"ADD_LOAD_ADJ,omitempty"`
	MED_LOAD_PERC                      float64 `csv:"MED_LOAD_PERC,omitempty"`
	MED_LOAD_ADJ                       float64 `csv:"MED_LOAD_ADJ,omitempty"`
	WAI_LOAD_PERC                      float64 `csv:"WAI_LOAD_PERC,omitempty"`
	WAI_LOAD_ADJ                       float64 `csv:"WAI_LOAD_ADJ,omitempty"`
	SPO_LOAD_PERC                      float64 `csv:"SPO_LOAD_PERC,omitempty"`
	SPO_LOAD_ADJ                       float64 `csv:"SPO_LOAD_ADJ,omitempty"`
	CCP_LOAD_PERC                      float64 `csv:"CCP_LOAD_PERC,omitempty"`
	CCP_LOAD_ADJ                       float64 `csv:"CCP_LOAD_ADJ,omitempty"`
	HS_LOAD_PERC                       float64 `csv:"HS_LOAD_PERC,omitempty"`
	HS_LOAD_ADJ                        float64 `csv:"HS_LOAD_ADJ,omitempty"`
	TERM_LOAD_PERC                     float64 `csv:"TERM_LOAD_PERC,omitempty"`
	TERM_LOAD_ADJ                      float64 `csv:"TERM_LOAD_ADJ,omitempty"`
	PARENT_LOAD_PERC                   float64 `csv:"PARENT_LOAD_PERC,omitempty"`
	PARENT_LOAD_ADJ                    float64 `csv:"PARENT_LOAD_ADJ,omitempty"`
	CINC_LOAD_PERC                     float64 `csv:"CINC_LOAD_PERC,omitempty"`
	CINC_LOAD_ADJ                      float64 `csv:"CINC_LOAD_ADJ,omitempty"`
	MCC_LOAD_PERC                      float64 `csv:"MCC_LOAD_PERC,omitempty"`
	MCC_LOAD_ADJ                       float64 `csv:"MCC_LOAD_ADJ,omitempty"`
	ESCC_LOAD_PERC                     float64 `csv:"ESCC_LOAD_PERC,omitempty"`
	ESCC_LOAD_ADJ                      float64 `csv:"ESCC_LOAD_ADJ,omitempty"`
	JCC_LOAD_PERC                      float64 `csv:"JCC_LOAD_PERC,omitempty"`
	JCC_LOAD_ADJ                       float64 `csv:"JCC_LOAD_ADJ,omitempty"`
	EDU_PAR1_LOAD_PERC                 float64 `csv:"EDU_PAR1_LOAD_PERC,omitempty"`
	EDU_PAR1_LOAD_ADJ                  float64 `csv:"EDU_PAR1_LOAD_ADJ,omitempty"`
	EDU_PAR2_LOAD_PERC                 float64 `csv:"EDU_PAR2_LOAD_PERC,omitempty"`
	EDU_PAR2_LOAD_ADJ                  float64 `csv:"EDU_PAR2_LOAD_ADJ,omitempty"`
	PAR2_TERM_LOAD_PERC                float64 `csv:"PAR2_TERM_LOAD_PERC,omitempty"`
	PAR2_TERM_LOAD_ADJ                 float64 `csv:"PAR2_TERM_LOAD_ADJ,omitempty"`
	CC61_LOAD_PERC                     float64 `csv:"CC61_LOAD_PERC,omitempty"`
	CC61_LOAD_ADJ                      float64 `csv:"CC61_LOAD_ADJ,omitempty"`
	CIR_LOAD_PERC                      float64 `csv:"CIR_LOAD_PERC,omitempty"`
	CIR_LOAD_ADJ                       float64 `csv:"CIR_LOAD_ADJ,omitempty"`
	TPDD_LOAD_PERC                     float64 `csv:"TPDD_LOAD_PERC,omitempty"`
	TPDD_LOAD_ADJ                      float64 `csv:"TPDD_LOAD_ADJ,omitempty"`
	CICA_CCB61_LOAD_PERC               float64 `csv:"CICA_CCB61_LOAD_PERC,omitempty"`
	CICA_CCB61_LOAD_ADJ                float64 `csv:"CICA_CCB61_LOAD_ADJ,omitempty"`
	CICA_CCB34_LOAD_PERC               float64 `csv:"CICA_CCB34_LOAD_PERC,omitempty"`
	CICA_CCB34_LOAD_ADJ                float64 `csv:"CICA_CCB34_LOAD_ADJ,omitempty"`
	CICA_ESCC_LOAD_PERC                float64 `csv:"CICA_ESCC_LOAD_PERC,omitempty"`
	CICA_ESCC_LOAD_ADJ                 float64 `csv:"CICA_ESCC_LOAD_ADJ,omitempty"`
	DIG_OFF_IND                        int     `csv:"DIG_OFF_IND,omitempty"`
	PPH_PLUS_PRO_BEN                   float64 `csv:"PPH_PLUS_PRO_BEN,omitempty"`
	ADM_IND_CHG                        float64 `csv:"ADM_IND_CHG,omitempty"`
	CAMP_IND_IF                        float64 `csv:"CAMP_IND_IF,omitempty"`
	CAMP_IND_NB                        float64 `csv:"CAMP_IND_NB,omitempty"`
	LOY_BONUS_IND                      int     `csv:"LOY_BONUS_IND,omitempty"`
}

type Output struct {
	PROD_NAME                          string    `type:"input" def:"PROD_NAME"`
	SPCODE                             int       `type:"input" def:"SPCODE"`
	POL_NUMBER                         int       `type:"input" def:"POL_NUMBER"`
	AGE_AT_ENTRY                       int       `type:"input" def:"AGE_AT_ENTRY"`
	SEX                                int       `type:"input" def:"SEX"`
	SMOKER_STAT                        int       `type:"input" def:"SMOKER_STAT"`
	POL_TERM_Y                         int       `type:"input" def:"POL_TERM_Y"`
	ANNUAL_PREM                        float64   `type:"input" def:"ANNUAL_PREM"`
	PREM_PAYBL_Y                       int       `type:"input" def:"PREM_PAYBL_Y"`
	PREM_FREQ                          int       `type:"input" def:"PREM_FREQ"`
	SUM_ASSURED                        float64   `type:"input" def:"SUM_ASSURED"`
	DURATIONIF_M                       int       `type:"input" def:"DURATIONIF_M"`
	IFRS_CB_TERM_M                     int       `type:"input" def:"IFRS_CB_TERM_M"`
	REGION                             int       `type:"input" def:"REGION"`
	BILL_CHNL_IDX                      int       `type:"input" def:"BILL_CHNL"`
	BILL_CHNL2_IDX                     int       `type:"input" def:"BILL_CHNL2"`
	DIST_IDX                           int       `type:"input" def:"DIST_IND"`
	IF_NB                              string    `type:"input" def:"IF_NB"`
	MTHS_TO_SALE                       int       `type:"input" def:"MTHS_TO_SALE"`
	INIT_POLS_IF                       int       `type:"input" def:"INIT_POLS_IF"`
	NO_LAP_GTEE_IND                    int       `type:"input" def:"NO_LAP_GTEE"`
	DTH_LOAD_AMT                       float64   `type:"input" def:"BASIC_LOAD_ADJ"`
	DTH_LOAD_PC                        float64   `type:"code" def:""`
	CC_LOAD_AMT                        float64   `type:"input" def:"CC_LOAD_ADJ"`
	CC_LOAD_PC                         float64   `type:"code" def:""`
	PHOL_IND                           int       `type:"input" def:"PH_IND_PP"`
	ANN_PREM_A                         float64   `type:"input" def:"ANN_PREM_BAS"`
	ANN_PREM_C                         float64   `type:"input" def:"ANN_PREM_REC"`
	ANNUAL_PREM_TOPUP                  float64   `type:"input" def:"ANNUAL_PREM_TOPUP"`
	ALLOC_PREM_A_IDX                   int       `type:"input" def:"FUND_IND_B"`
	ALLOC_PREM_C_IDX                   int       `type:"input" def:"FUND_IND_R"`
	FUND_UNIT_A_1                      float64   `type:"input" def:"I_UNFDU_B_F1"`
	FUND_UNIT_C_1                      float64   `type:"input" def:"I_UNFDU_R_F1"`
	SUBFUND_A_PC_1                     float64   `type:"input" def:"PRM_B_ALL_F1"`
	SUBFUND_C_PC_1                     float64   `type:"input" def:"PRM_R_ALL_F1"`
	SUBFUND_PC_1                       float64   `type:"code" def:""`
	DEDN_CFWD_A_0_1                    float64   `type:"input" def:"I_EXCS_DD_B"`
	DEDN_CFWD_C_0_1                    float64   `type:"input" def:"I_EXCS_DD_R"`
	DEDN_CFWD_0_1                      float64   `type:"input" def:"INIT_EXCS_DD"`
	DEDN_LBU_A_PP                      float64   `type:"code" def:""`
	DEDN_LBU_C_PP                      float64   `type:"code" def:""`
	DEDN_DISC_TIER                     int       `type:"code" def:""`
	BON_UNIT_IND                       int       `type:"code" def:""`
	ANNUAL_PREM_BON                    float64   `type:"input" def:"ANN_BON_PREM_BAS"`
	INIT_VOUG_1                        float64   `type:"input" def:"INIT_BONP_F1"`
	INIT_VOUP_1                        float64   `type:"input" def:"INIT_BONP_F1"`
	PART_SURR_IDX                      int       `type:"code" def:""`
	ALLOC_IDX                          int       `type:"input" def:"ALLOC_IND"`
	EXP_A_IDX                          int       `type:"input" def:"EXP_IND_B"`
	EXP_C_IDX                          int       `type:"input" def:"EXP_IND_R"`
	EXP_A_TOPUP_IDX                    int       `type:"input" def:"EXP_IND_B_TOPUP"`
	HO_ISS_DATE                        string    `type:"input" def:"HO_ISS_DATE"`
	COMM_A_IDX                         int       `type:"input" def:"COMM_IND_B"`
	COMM_C_IDX                         int       `type:"input" def:"COMM_IND_R"`
	COMM_A_TOPUP_IDX                   int       `type:"input" def:"COMM_IND_B_TOPUP"`
	SA_MULT_GRP                        int       `type:"input" def:"SA_MULT_GRP"`
	RI_BEN_CI                          float64   `type:"input" def:"RI_BEN_CI"`
	RI_BEN_WAIV                        float64   `type:"input" def:"RI_BEN_WAIV"`
	RI_BEN_WAIV2                       float64   `type:"input" def:"RI_BEN_WAIV2"`
	ACDB_POL_TERM_M                    float64   `type:"code" def:""`
	DUR_PEP_AT_ENTRY                   int       `type:"input" def:"DUR_PEP_AT_ENTRY"`
	RI_SUMASSD_PC                      float64   `type:"input" def:"RI_SUMASSD_PC"`
	RI_BASE_IDX                        []int     `type:"code" def:""`
	RI_RIDER_IDX                       []int     `type:"code" def:""`
	RI_RIDER_CED_PC                    []float64 `type:"code" def:""`
	NUM_OF_RIDERS                      int       `type:"code" def:""`
	SEL_FAC_IND                        int       `type:"input" def:"SEL_FAC_IND"`
	AGE_AT_ENTRY2                      int       `type:"input" def:"LIFE2_AGE_AE"`
	AGE_AT_ENTRY3                      int       `type:"input" def:"LIFE3_AGE_AE"`
	AGE_AT_ENTRY_MOTHER                int       `type:"input" def:"AGE_MT_PRNTL"`
	CESS_AGE_SP                        int       `type:"input" def:"CESS_AGE_SP"`
	CESS_AGE_WP                        int       `type:"input" def:"CESS_AGE_WP"`
	CC_EXPIRE_AGE                      int       `type:"input" def:"CC_EXP"`
	SEX2                               int       `type:"code" def:"SEX_2"`
	SEX3                               int       `type:"input" def:"SEX_3"`
	HS_POLICY_IND                      int       `type:"code" def:""`
	HSPLUS_IND                         int       `type:"input" def:"HSPLUS_IND"`
	CINC_CLAIMED_IND                   int       `type:"input" def:"CINC_CLAIM_STAT"`
	PEP_CLAIMED_IND                    int       `type:"input" def:"PEP_CLAIM_STAT"`
	WOP_CLAIMED_IND                    int       `type:"input" def:"WVR_CLAIM_STAT"`
	RIDER_IDX                          []int     `type:"code" def:""`
	RIDER_SA                           []float64 `type:"code" def:""`
	RIDER_PREM                         []float64 `type:"code" def:""`
	RIDER_POL_TERM_M                   []int     `type:"code" def:""`
	RIDER_PREM_TERM_M                  []int     `type:"code" def:""`
	RIDER_LOAD_AMT                     []float64 `type:"code" def:""`
	RIDER_LOAD_PC                      []float64 `type:"code" def:""`
	RIDER_VERSION_IDX                  []int     `type:"code" def:""`
	RIDER_BEN_IDX                      []int     `type:"code" def:""`
	REDU_ADM                           int       `type:"input" def:"REDU_ADM"`
	DIG_OFF_IND                        int       `type:"input" def:"DIG_OFF_IND"`
	MUNRE_2020                         int       `type:"input" def:"MUNRE_2020"`
	SUM_ASSURED_GIO                    float64   `type:"code" def:""`
	SUM_ASSURED_XSA                    float64   `type:"code" def:""`
	GIO_BONUS_POL_TERM_M               int       `type:"code" def:""`
	XSA_BONUS_POL_TERM_M               int       `type:"code" def:""`
	RIDER_SA_GIO                       []float64 `type:"code" def:""`
	RIDER_SA_XSA                       []float64 `type:"code" def:""`
	RIDER_POL_TERM_M_GIO               []float64 `type:"code" def:""`
	RIDER_POL_TERM_M_XSA               []float64 `type:"code" def:""`
	DURATIONIF_PHOL                    []float64 `type:"code" def:""`
	ALT_ANN_PREM_A                     []float64 `type:"code" def:""`
	ALT_PREM_COUNTER                   []float64 `type:"code" def:""`
	CAMP_IND_IF                        float64   `type:"input" def:"CAMP_IND_IF"`
	CAMP_IND_NB                        float64   `type:"input" def:"CAMP_IND_NB"`
	TOP_UP_YEAR                        int       `type:"input" def:"TOP_UP_YEAR"`
	UPR_FAC                            float64   `type:"input" def:"UPR_FAC"`
	BASIC_ENTRY_YEAR                   int       `type:"input" def:"BASIC_ENTRY_YEAR"`
	BASIC_ENTRY_MONTH                  int       `type:"input" def:"BASIC_ENTRY_MONTH"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY1 string    `type:"input" def:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY1"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY2 string    `type:"input" def:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY2"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY3 string    `type:"input" def:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY3"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY4 string    `type:"input" def:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY4"`
	IFRS_RI_CSM_GROUP_CURRENCY_TREATY5 string    `type:"input" def:"IFRS_RI_CSM_GROUP_CURRENCY_TREATY5"`
	IFRS_RI_ICG_ID_TREATY1             string    `type:"input" def:"IFRS_RI_ICG_ID_TREATY1"`
	IFRS_RI_ICG_ID_TREATY_RID          []string  `type:"code" def:""`
	TREATY_ID_TREATY1                  string    `type:"input" def:"TREATY_ID_TREATY1"`
	TREATY_ID_TREATY_RID               []string  `type:"code" def:""`
	IFRS_RIDER_CODE                    []string  `type:"code" def:""`
	ENTITY_ID                          string    `type:"input" def:"ENTITY_ID"`
	PROD_CD                            string    `type:"input" def:"PROD_CD"`
	IFRS_CY_GRP                        int       `type:"input" def:"IFRS_CY_GRP"`
	IFRS_PORT_GRP                      int       `type:"input" def:"IFRS_PORT_GRP"`
	IFRS_ICG_ID                        string    `type:"input" def:"IFRS_ICG_ID"`
	IFRS_ICG_ID_PROPHET                int       `type:"input" def:"IFRS_ICG_ID_PROPHET"`
	IFRS_ONEROUS_GRP                   int       `type:"input" def:"IFRS_ONEROUS_GRP"`
	IFRS_CCF_EXTERNAL                  []int     `type:"code" def:""`
	EXT_CCF_IND                        int       `type:"input" def:"EXT_CCF_IND"`
	IFRS_RI_EXT_CCF_IND                int       `type:"input" def:"IFRS_RI_EXT_CCF_IND"`
	IFRS_MEASURE_MODEL                 int       `type:"input" def:"IFRS_MEASURE_MODEL"`
	ADM_IND_CHG                        float64   `type:"input" def:"ADM_IND_CHG"`
	WOP_IND                            int       `type:"input" def:"WVR_IND"`
	WOP2_ESCC_IND                      int       `type:"input" def:"WOP2_ESCC_IND"`
	CINC_IND                           int       `type:"input" def:"CINC_IND"`
	CINCOME_BEN                        float64   `type:"input" def:"CINCOME_BEN"`
	MEDICAL_BEN                        float64   `type:"input" def:"MEDICAL_BEN"`
	PPH_PLUS_BEN2                      int       `type:"code" def:""`
	PPH_PLUS_PRO_BEN2                  int       `type:"code" def:""`
	CI_BEN                             float64   `type:"input" def:"CI_BEN"`
	PLAN_MED                           int       `type:"input" def:"PLAN_MED"`
	PPH_PLUS_BEN                       float64   `type:"input" def:"PPH_PLUS_BEN"`
	PPH_PLUS_PRO_BEN                   float64   `type:"input" def:"PPH_PLUS_PRO_BEN"`
	MEDICAL_BEN2                       float64   `type:"code" def:""`
	MAX_AGE_CINCOME                    int       `type:"input" def:"MAX_AGE_CINCOME"`
	MAX_AGE_PEP2                       int       `type:"input" def:"MAX_AGE_PEP2"`
	MAX_AGE_PEP3                       int       `type:"code" def:""`
	MAX_AGE_WAIV                       int       `type:"input" def:"MAX_AGE_WAIV"`
	MAX_AGE_WAIV2                      int       `type:"input" def:"MAX_AGE_WAIV2"`
	PEP_BEN_FIX                        float64   `type:"input" def:"PEP_BEN_FIX"`
	PEP_BEN_FIX2                       float64   `type:"input" def:"PEP_BEN_FIX2"`
	PEP_IND                            int       `type:"input" def:"PEP_IND"`
	PLAN_CC                            int       `type:"input" def:"PLAN_CC"`
	PW_IND_USED                        int       `type:"input" def:"PW_IND_USED"`
	PWAI2_BEN                          float64   `type:"input" def:"PWAI2_BEN"`
	PWAIV_BEN                          float64   `type:"input" def:"PWAIV_BEN"`
	CC_BEN                             float64   `type:"input" def:"CC_BEN"`
	TEMP_PLAN_PRNTL                    int       `type:"code" def:""`
	TEMP_PRNTL_BEN                     float64   `type:"code" def:""`
	TEMP_TERM_BEN                      float64   `type:"code" def:""`
	TEMP_PW_FLAG                       int       `type:"code" def:""`
	LIFE3_AGE_AE                       int       `type:"input" def:"LIFE3_AGE_AE"`
	TEST_ID                            int       `type:"code" def:""`
}
