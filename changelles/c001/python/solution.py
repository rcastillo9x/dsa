def solution(arr, target):
    pairs= {}
    
    for index, number in enumerate(arr):
        complement = target - number
        if complement in pairs:
            print(f"Found: number { number } complement { complement }")
            return pairs[complement], index
        pairs[number] = index
 
            

if __name__ == '__main__':
    arr = [7, 6, 3, 8, 2]

    print(solution(arr, 9))