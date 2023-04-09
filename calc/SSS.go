package calc

// 2022 SSS Contribution Table
// https://web.z.com/ph/blog/sss-monthly-contribution/
// https://cashmart.ph/sss-unemployment-benefit/
func CalcSSS(compensation float64) float64 {
	var result float64

	table := [][3]float64{
		{1000, 3249.99, 135},
		{3250, 3749.99, 157.50},
		{3750, 4249.99, 180},
		{4250, 4749.99, 202.5},
		{4750, 5249.99, 225},
		{5250, 5749.99, 247.5},
		{5750, 6249.99, 270},
		{6250, 6749.99, 292.5},
		{6750, 7249.99, 315},
		{7250, 7749.99, 337.5},
		{7750, 8249.99, 360},
		{8250, 8749.99, 382.5},
		{8750, 9249.99, 405},
		{9250, 9749.99, 427.5},
		{9750, 10249.99, 450},
		{10250, 10749.99, 472.5},
		{10750, 11249.99, 495},
		{11250, 11749.99, 517.5},
		{11750, 12249.99, 540},
		{12250, 12749.99, 562.5},
		{12750, 13249.99, 585},
		{13250, 13749.99, 607.5},
		{13750, 14249.99, 630},
		{14250, 14749.99, 652.5},
		{14750, 15249.99, 675},
		{15250, 15749.99, 697.5},
		{15750, 16249.99, 720},
		{16250, 16749.99, 742.5},
		{16750, 17249.99, 765},
		{17250, 17749.99, 787.5},
		{17750, 18249.99, 810},
		{18250, 18749.99, 832.5},
		{18750, 19249.99, 855},
		{19250, 19749.99, 877.5},
		{19750, 20249.99, 900},
		{20250, 20749.99, 922.5},
		{20750, 21249.99, 945},
		{21250, 21749.99, 967.5},
		{21750, 22249.99, 990},
		{22250, 22749.99, 1012.5},
		{22270, 23249.99, 1035},
		{23250, 23749.99, 1057.5},
		{23750, 24249.99, 1080},
		{24250, 24749.99, 1102.5},
	}

	switch {
	// Note: 2022 SSS Contribution Table has no value for 999.99 and below. This func will return 0 instead.
	case compensation >= 24750:
		result = 1125
	default:
		for _, row := range table {
			if compensation >= row[0] && compensation <= row[1] {
				result = row[2]
			}
		}
	}

	return result
}
