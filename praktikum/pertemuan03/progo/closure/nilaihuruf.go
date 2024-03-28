package closure

func GetNilaiHuruf(nilaiUjian int) string {
    batasNilai := map[string]int{
        "A": 80,
        "B": 70,
        "C": 60,
        "D": 50,
    }

    for huruf, batas := range batasNilai {
        if nilaiUjian >= batas {
            return huruf
        }
    }
    return "E"
}
