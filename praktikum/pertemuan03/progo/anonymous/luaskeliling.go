package anonymous

import (
    "math"
)

func HitungLuasKeliling(jariJari float64) (luas, keliling float64) {
    // Menghitung luas dan keliling dengan anonymous function 
    luas = func(r float64) float64 { 
        return math.Pi * r * r 
    }(jariJari) 
 
    keliling = func(r float64) float64 { 
        return 2 * math.Pi * r 
    }(jariJari) 

    return luas, keliling
}
