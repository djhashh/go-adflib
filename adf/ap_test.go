package adf

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/gonum/floats"
	"github.com/tetsuzawa/go-research/go-adf/misc"
)

/*
func TestFiltAP_Adapt(t *testing.T) {
	type fields struct {
		filtBase filtBase
		kind           string
		order          int
		eps            float64
		wHistory       *mat.Dense
		xMem           *mat.Dense
		dMem           *mat.Dense
		yMem           *mat.Dense
		eMem           *mat.Dense
		epsIDE         *mat.Dense
		ide            *mat.Dense
	}
	type args struct {
		d float64
		x []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &FiltAP{
				filtBase: tt.fields.filtBase,
				kind:           tt.fields.kind,
				order:          tt.fields.order,
				eps:            tt.fields.eps,
				wHistory:       tt.fields.wHistory,
				xMem:           tt.fields.xMem,
				dMem:           tt.fields.dMem,
				yMem:           tt.fields.yMem,
				eMem:           tt.fields.eMem,
				epsIDE:         tt.fields.epsIDE,
				ide:            tt.fields.ide,
			}
		})
	}
}
*/

func TestFiltAP_Run(t *testing.T) {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 64
	L := 4
	//input value
	var x = make([][]float64, n)
	//noise
	var v = make([]float64, n)
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, rand.NormFloat64())
		x[i] = append([]float64{}, xRow...)
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][0]
	}
	//type fields struct {
	//	filtBase filtBase
	//	kind           string
	//	order          int
	//	eps            float64
	//	wHistory       *mat.Dense
	//	xMem           *mat.Dense
	//	dMem           *mat.Dense
	//	yMem           *mat.Dense
	//	eMem           *mat.Dense
	//	epsIDE         *mat.Dense
	//	ide            *mat.Dense
	//}
	type args struct {
		d []float64
		x [][]float64
	}
	tests := []struct {
		name string
		//fields  fields
		args    args
		want    []float64
		want1   []float64
		want2   [][]float64
		wantErr bool
	}{
		{
			name: "Run AP Filter",
			args: args{
				d: d,
				x: x,
			},
			want:    []float64{0, 0, 0, 0.802869025666775, -0.5051833624336448, 0.19321069483247766, 0.08414012901448248, 0.01819087707636731, -0.9724044742768908, 1.1255386871921087, -0.36697072877339, 0.10881438257027298, 0.41350563005357366, 0.48222559828493095, -0.3603423163193269, 1.2986354516902474, -0.2056719500878501, -1.053566789626397, 0.6924544712619379, -0.8418523862734272, 0.5962156133434275, 0.2807937437818256, -0.3120005270751306, -0.6349031312027253, -0.05426734892564805, -0.07208582230856853, 1.6349713868205522, 0.4587083081312851, 0.1313159496409251, 0.14729434732832772, -1.6640170895463562, 0.350033044755764, 0.2497236323965607, -1.8144550964150357, 1.900613977489524, 0.4908281195343719, -1.14661723624259, 2.7201801762765543, 0.49096219558646437, -2.507320700623442, -1.2945311566423394, -3.563219427347747, 2.9957575883101875, 1.7601427591950753, 1.125806760369593, -1.1906331266649801, -4.020766638316268, 2.776103088046718, -3.7632122126985834, -0.6408886381008105, 2.57501205305168, 5.987541854432293, -2.42522363768885, -3.6267231847981707, 0.9478114098410717, 6.890816162619367, 1.206938123671592, -1.9346312152318679, -1.5918673467118551, -1.4672750901335518, 8.509546253637483, 1.1743583484609474, -0.4545588870766662, 2.704348798958019,},
			want1:   []float64{0, 0, 0, -2.0366272032647217, -0.015811208719505476, 0.12959455777910223, 0.07466761116195314, -0.7494738932538464, 2.5578084365575138, 0.17330216032532553, 1.099412654577903, 0.5913065198697118, 0.5861204909576888, -0.7987628411790192, 1.4610715100693477, -0.3089250314817158, -1.229374972044378, 1.1909171631997477, -1.538548844717535, 0.9979788815562418, -0.31647189468025605, 0.42195651778035564, -0.7575448752778898, 0.9651925168153871, -1.0664467434309852, 0.9993842930726518, -0.661937004977367, 0.051407802354731025, 0.1202240944182241, 0.019396371589969846, 0.025376011689921762, 0.8041977689542732, -1.0202845802463545, 1.0382418865412686, -0.480334480700644, -0.8212848820093415, 0.8995825194123539, -0.9845469196219927, -0.6967961228433243, 1.455655403019498, 0.6170076103792002, 1.6015815162628386, -1.0098383063290552, -1.8012808672640679, 0.07792485411682426, 0.35204036414563, 2.617673909942603, -1.8018199241580963, 1.770249842707101, 1.022122378400761, -2.1654408584922304, -3.4573034110798866, 1.1549809191933569, 3.206095372176847, -0.9456507412384341, -4.507955817187426, -1.0156720058709843, 1.9446374132167068, 1.1335544985056158, 0.9908175111673628, -5.778424382613037, -1.121887709813678, 1.2544352892302317, -1.9790113110143075,},
			want2:   [][]float64{{-0.6507507226658572, 0.657011901126719, -0.5912859617953181, 0.31741294852014135, }, {-0.6507507226658572, 0.657011901126719, -0.5912859617953181, 0.31741294852014135, }, {-0.6507507226658572, 0.657011901126719, -0.5912859617953181, 0.31741294852014135, }, {-0.6507507226658572, 0.6444100956878449, -0.5505421916919521, 0.3033659059943217, }, {-0.5521970683546433, 0.642650279163117, -0.5629341418005288, 0.30898805185485523, }, {-0.49138855369063095, 0.6574268586087346, -0.5627920753921952, 0.30146563509516294, }, {-0.420887983585906, 0.6142060538012653, -0.6004091097127641, 0.2918711583578221, }, {-0.34024150895557875, 0.668466651494392, -0.5875380690135541, 0.28238156688679084, }, {-0.3308304749020782, 0.6122506874232411, -0.5481390222199736, 0.2691180504347391, }, {-0.19675194743602448, 0.5927055146980833, -0.5435832557702132, 0.2764177552726182, }, {-0.1255797820767666, 0.6022092600421253, -0.5450525448764298, 0.2639694187808881, }, {-0.0344821062456471, 0.5927540263596934, -0.536957064602162, 0.24991731042233067, }, {0.014292986732229257, 0.6288456226032885, -0.5226064607321647, 0.26310298568607476, }, {0.03435531059213745, 0.65055355731077, -0.48354614718178773, 0.24757016648249708, }, {0.07414509349131344, 0.6032228807833149, -0.438341236721251, 0.267330787762371, }, {0.1313875510632423, 0.620273334105952, -0.466787297321897, 0.338217744758321, }, {0.16449521623071361, 0.6686166941647445, -0.47954399562788086, 0.32698341923136115, }, {0.22467108639905353, 0.6377277517343437, -0.4628261522460664, 0.26237849170627114, }, {0.33099658414116045, 0.5856145444533772, -0.43971709085596067, 0.26377390710845305, }, {0.42609962500856774, 0.5663162682263965, -0.4174814098872975, 0.25914106771095274, }, {0.49253833562881705, 0.5629010799939641, -0.4053659981562781, 0.20072347815728334, }, {0.5202457848221002, 0.5178399397880908, -0.40560623484714686, 0.1966042008011857, }, {0.5592387548809309, 0.5414408737677309, -0.4225282395286183, 0.15264940961772946, }, {0.5946523197013959, 0.5214066870301344, -0.43164598873458576, 0.10615779316357202, }, {0.6019877251183502, 0.4767243040032946, -0.3749921939923649, 0.08802646821362557, }, {0.6928675847547843, 0.47774359092767266, -0.33073477606851726, 0.06538171550802999, }, {0.8072797161514413, 0.4587700845051254, -0.3685760701394982, 0.03322926504452952, }, {0.8366518765126616, 0.4301979135379306, -0.388091806660845, 0.023914854971403575, }, {0.8697740299482034, 0.4627989003317823, -0.3681629585919799, 0.03740650370567282, }, {0.9593207669558845, 0.4882720037427491, -0.35392314267927777, 0.04635644482160593, }, {1.017211297774383, 0.46661775341861134, -0.4118199987669563, 0.0561352584532345, }, {1.0734122862469098, 0.49575295353057236, -0.4320122075811327, -0.018124001986969243, }, {1.1205221505000338, 0.4349037367170888, -0.37470855564997657, -0.017024839023229987, }, {1.2251750542623578, 0.4093399288516509, -0.3806367239376972, 0.06633181432491886, }, {1.297891126929455, 0.4818645740766877, -0.432012704474937, 0.08523715395912285, }, {1.3160094756821608, 0.45147327660541836, -0.4005027107886914, 0.034236594307225685, }, {1.4271520692269368, 0.35761326022321516, -0.43946275743066154, 0.06663821758880388, }, {1.4370431827675811, 0.4208485934015055, -0.5605416760586179, 0.10190224032689446, }, {1.4870546481980638, 0.4304837086871396, -0.44988410088691466, 0.18535053840339344, }, {1.69060394333388, 0.4642159271929022, -0.3476518690467093, 0.12315543650020042, }, {1.8570362364198045, 0.3540978125173727, -0.45114009363108315, 0.14011431449036585, }, {1.9073686929607359, 0.33206705615858284, -0.438358008777701, 0.28020329719538006, }, {1.9581749987998502, 0.4768593953885656, -0.44842164280875846, 0.24856829763337634, }, {1.98120600533486, 0.4270980737297785, -0.5547889733145839, 0.13996791284419668, }, {1.9810080320460788, 0.36228696888802364, -0.474323968272799, 0.15391338007530164, }, {2.1454415258133146, 0.3770258621633636, -0.40153849210074444, 0.06957086669996387, }, {2.3789923901271384, 0.23119327360355377, -0.4038037603817205, 0.06979797079408727, }, {2.44953783932032, 0.1927630238993015, -0.5776663458205635, 0.14587736182375077, }, {2.375826377345893, 0.2468965843270776, -0.6105478945023629, 0.1496341812796585, }, {2.560195513638454, 0.31898364533890466, -0.5234391675236515, 0.3358446452567905, }, {2.5935103846333294, 0.47598128922804367, -0.5765765593522174, 0.1870394563469767, }, {2.6453919373715626, 0.32622244816590884, -0.7153678586897085, 0.28441052631161057, }, {2.598864971383949, 0.43979698347144025, -0.6800391645040362, 0.10937879379477058, }, {2.722776225076692, 0.32979688715758965, -0.8477437981054685, 0.20135269931845245, }, {2.623212367408185, 0.362298502669838, -0.729601028217503, 0.06630440339542448, }, {2.7314615353417366, 0.3410561723889439, -0.9572937457486621, 0.01674994737860297, }, {2.5702972486193856, 0.291721584560639, -0.914402379558519, -0.05270995265857674, }, {2.725525687226232, 0.3419169922677623, -0.850748746061584, -0.039362665766004945, }, {2.8847675917283517, 0.0736532752898652, -0.933046907136647, -0.03861662522987369, }, {2.938614125067058, 0.1583298068340384, -0.8340734500676892, 0.07194358504947579, }, {2.971205262221316, 0.22273598198312505, -1.092636680440552, 0.01772864570601724, }, {2.8163124432840787, 0.1671841316092459, -1.1228308035637653, -0.07637392555692735, }, {2.885667060658634, 0.07543207934105499, -1.014831934556014, -0.0103594723874961, }, {3.0745760988315642, 0.16594698539632707, -1.048212899719334, 0.1451809316186104, },},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//af := &FiltAP{
			//	filtBase: tt.fields.filtBase,
			//	kind:           tt.fields.kind,
			//	order:          tt.fields.order,
			//	eps:            tt.fields.eps,
			//	wHistory:       tt.fields.wHistory,
			//	xMem:           tt.fields.xMem,
			//	dMem:           tt.fields.dMem,
			//	yMem:           tt.fields.yMem,
			//	eMem:           tt.fields.eMem,
			//	epsIDE:         tt.fields.epsIDE,
			//	ide:            tt.fields.ide,
			//}
			af := Must(NewFiltAP(L, 0.1, 4, 1e-5, "random"))
			got, got1, got2, err := af.Run(tt.args.d, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
				for i := 0; i < n; i++ {
					fmt.Printf("%g, ", got[i])
				}
				fmt.Println("")
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Run() got1 = %v, want %v", got1, tt.want1)
				for i := 0; i < n; i++ {
					fmt.Printf("%g, ", got1[i])
				}
				fmt.Println("")
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Run() got2 = %v, want %v", got2, tt.want2)
				fmt.Println(len(got2), len(got2[0]))
				for i := 0; i < n; i++ {
					fmt.Print("{")
					for k := 0; k < L; k++ {
						fmt.Printf("%g, ", got2[i][k])
					}
					fmt.Print("}, ")
				}
				fmt.Println("")
			}
		})
	}
}

/*
func TestNewFiltAP(t *testing.T) {
	type args struct {
		n     int
		mu    float64
		order int
		eps   float64
		w     interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    AdaptiveFilter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFiltAP(tt.args.n, tt.args.mu, tt.args.order, tt.args.eps, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFiltAP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFiltAP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

func ExampleExploreLearning_ap() {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 64
	L := 4
	order := 4
	mu := 1.0
	eps := 0.001
	//input value
	var x = make([][]float64, n)
	//noise
	var v = make([]float64, n)
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, rand.NormFloat64())
		x[i] = append([]float64{}, xRow...)
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][L-1]
	}

	af, err := NewFiltAP(L, mu, order, eps, "random")
	check(err)
	es, mus, err := ExploreLearning(af, d, x, 0.001, 2.0, 100, 0.5, 100, "MSE", nil)
	check(err)

	res := make(map[float64]float64, len(es))
	for i := 0; i < len(es); i++ {
		res[es[i]] = mus[i]
	}
	//for i := 0; i < len(es); i++ {
	//	fmt.Println(es[i], mus[i])
	//}
	eMin := floats.Min(es)
	fmt.Printf("the step size mu with the smallest error is %.3f\n", res[eMin])
	//output:
	//the step size mu with the smallest error is 0.001
}
