package functions

import "math"


// function linear 
func Linear_Regression(numbers []float64) (float64, float64, float64) {
	lenght := float64(len(numbers))
	var SommeX, SommeY, SommeXY, SommeX2, SommeY2, Coefficient float64
    
	// variable indépendante i(0, 1, 2, .....)
	// variable dépendante y(contenu de data.txt)
	for i, y := range numbers {
		x := float64(i)
		SommeX = SommeX + x     // somme des Xi
		SommeY = SommeY + y     // somme des Yi
		SommeXY = SommeXY + (x * y)   // somme des (Xi * Yi)
		SommeX2 = SommeX2 + (x * x)   // somme des (Xi)²
		SommeY2 = SommeY2 + (y * y)   // somme des (Yi)²
	}
     
	// Regresssion linéaire (y(x) = ax + b)
	a := (lenght*SommeXY - SommeX*SommeY) / (lenght*SommeX2 - math.Pow(SommeX , 2))
	b := (SommeY - a*SommeX) / lenght
    
	// coefficient de corrélation de Pearson
    num := (lenght*SommeXY - SommeX*SommeY)   // LE NUMERATEUR
	deno := math.Sqrt((lenght*SommeX2 - math.Pow(SommeX , 2)) * (lenght*SommeY2 - math.Pow(SommeY , 2)))  // LE NUMERATEUR
    if deno == 0 {     //Eliminer la division par zero
		Coefficient = 0   
	}
	Coefficient = num / deno

	return a, b , Coefficient
}
