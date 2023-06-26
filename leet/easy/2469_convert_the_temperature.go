// https://leetcode.com/problems/convert-the-temperature/

// Uses a straight-forward conversion

func convertTemperature(celsius float64) []float64 {
    const kelvinAdd = 273.15
    const fhRate = 1.80
    const fhAdd = 32.00
    
    kelvin := celsius + kelvinAdd
    fh := celsius * fhRate + fhAdd
    return []float64{kelvin, fh}
}
