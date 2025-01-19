
def run_cases(test, cases, solution, sort=False):
    for case in cases:
        params = case[:-1]
        expected = case[-1]
        found = solution(*params)
        if sort:
            found = list(sorted(found))
            expected = list(sorted(expected))
        try:
            test.assertEqual(found, expected)
        except:
            print("Failed Case: ", params)
            print("Found: ", found)
            print("Expected : ", expected)
            raise
