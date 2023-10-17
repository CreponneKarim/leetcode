package leetcode

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    var isParent = make([]bool,n)
	var isChild  = make([]bool,n)

	for i := 0; i < n; i++ {
	
		r := rightChild	[i]
		l := leftChild	[i]

		if l!=-1 {
			isParent[i]	=true

			if !isChild[l]{
				isChild[l]	=true
			}else{
				return false
			}

		}
		if r!=-1 {
			isParent[i]	=true

			if !isChild[r]{
				isChild[r]	=true
			}else{
				return false
			}
		}
	}

	var nbElemsIsNotChild int = 0
	
	for index, elem:= range isChild {
		if !elem {
			if nbElemsIsNotChild==0{
				nbElemsIsNotChild++
				if !isParent[index] {
					return false
				} 
			}else {
				return false
			}
		}
	}
    
    return true
}