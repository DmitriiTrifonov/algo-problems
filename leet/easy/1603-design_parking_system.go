// https://leetcode.com/problems/design-parking-system/

// Use switch case

type ParkingSystem struct {
    big int
    medium int
    small int
    currentBig int
    currentMedium int
    currentSmall int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{
        big: big,
        medium: medium,
        small: small,
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType {
        case 1: 
        if this.currentBig + 1 <= this.big {
            this.currentBig++
            return true
        }
        return false
        case 2: 
        if this.currentMedium + 1 <= this.medium {
            this.currentMedium++
            return true
        }
        return false
        case 3: 
        if this.currentSmall + 1 <= this.small {
            this.currentSmall++
            return true
        }
        return false
    }
    return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
