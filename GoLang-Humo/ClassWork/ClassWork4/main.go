package main

import "fmt"

/*for init; cond; operations{
	body
}

for i = 0; i <= 9; i++ {
	fmt.Println(i)
}

=
==
!=

init = 0
for ;cond;operations{
	body
}
*/

func main() {
	/*	var i int
											for i = 0; i <= 9; i++ {
												fmt.Println(i)
											}
											fmt.Println(i)

											var num = 546
											for num > 0 {
												a := num % 10
												fmt.Println(a)
												num = num / 10
											}
											//n = n / 10   n /= 10
											for n := 546; n > 0; n /= 10 {
												a := n % 10
												fmt.Println(a)
											}

											for {
												time.Sleep(1 * time.Second)
												fmt.Println("Hello World")
											}

											//Continue
											//break

										for i := 0; i < 100; i++ {
											if i%3 == 0 {
												continue
											}
											fmt.Println(i)
										}
									for i := 1; i < 100; i++ {
										if i%3 == 0 {
											break
										}
										fmt.Println(i)
									}

								for i := 1; i < 10; i++ {
									for j := 1; j < 10; j++ {
										fmt.Printf("%d * %d = %d\n", i, j, i*j)
									}
									fmt.Println("**********")
								}

							//for 10
							var n int
							var Sum float64 = 0
							fmt.Scan(&n)
							for i := 1; i <= n; i++ {
								Sum = Sum + 1/float64(i)
							}
							fmt.Printf("%.2f", Sum)
						//for 12
						var n int
						var p float64 = 1
						fmt.Scan(&n)
						for i := 1; i <= n; i++ {
							p = p * (1 + float64(i)*0.1)
						}
						fmt.Println(p)
						//for 13
						var m int
						var proz float64 = 0
						fmt.Scan(&m)
						var k float64 = 1
						for i := 1; i <= m; i++ {
							proz = proz + (1+float64(i)*0.1)*k
							k = k * (-1)
						}
						fmt.Println(proz)


					//for 14
					var n int
					fmt.Scan(&n)
					var N int = 0
					for i := 1; i <= n; i++ {
						N = N + (i*2 - 1)
						fmt.Println(N)
					}

				//for 15
				var A float64
				var N int
				var M float64 = 1
				fmt.Scan(&A, &N)
				for i := 0; i < N; i++ {
					M = M * A
				}
				fmt.Println(M)

			//series2
			var proiz float64 = 1
			var l float64
			for i := 0; i < 10; i++ {
				fmt.Scan(&l)
				proiz = proiz * l
			}
			fmt.Println(proiz)

		//series 3
		var n float64 = 0
		var l float64
		for i := 0; i < 10; i++ {
			fmt.Scan(&l)
			n = n + l
		}
		fmt.Println(n / 2)
	*/
	//series 6
	var N int
	fmt.Scan(&N)
	var M float64
	var proiz float64 = 1
	for i := 0; i < N; i++ {
		fmt.Scan(&M)
		l := int64(M)
		k := M - float64(l)
		fmt.Printf("%.2f", k)
		proiz = proiz * k
	}
	fmt.Printf("\n%.5f", proiz)
}
